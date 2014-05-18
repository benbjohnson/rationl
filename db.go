package rationl

import (
	"encoding/binary"
	"fmt"
	"log"
	"net/http"
	"os"

	"code.google.com/p/goauth2/oauth"
	"github.com/boltdb/bolt"
	"github.com/google/go-github/github"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

// DB represents the primary data storage.
type DB struct {
	*bolt.DB
	store sessions.Store
}

// Open opens and initializes the database.
func Open(path string, mode os.FileMode) (*DB, error) {
	var db = &DB{}
	var err error
	db.DB, err = bolt.Open(path, mode)
	if err != nil {
		return nil, err
	}

	// Initialize schema.
	err = db.Update(func(tx *Tx) error {
		var meta, _ = tx.CreateBucketIfNotExists([]byte("Meta"))
		tx.CreateBucketIfNotExists([]byte("User"))
		tx.CreateBucketIfNotExists([]byte("Investigation"))
		tx.CreateBucketIfNotExists([]byte("Experiment"))

		// Initialize secure cookie store.
		var secret = meta.Get([]byte("secret"))
		if secret == nil {
			secret = securecookie.GenerateRandomKey(64)
			if err := meta.Put([]byte("secret"), secret); err != nil {
				return fmt.Errorf("secret: %s", err)
			}
		}
		db.store = sessions.NewCookieStore(secret)

		return nil
	})
	if err != nil {
		_ = db.DB.Close()
		return nil, err
	}
	return db, nil
}

// View executes a function in the context of a read-only transaction.
func (db *DB) View(fn func(*Tx) error) error {
	return db.DB.View(func(tx *bolt.Tx) error {
		return fn(&Tx{tx, db})
	})
}

// Update executes a function in the context of a writable transaction.
func (db *DB) Update(fn func(*Tx) error) error {
	return db.DB.Update(func(tx *bolt.Tx) error {
		return fn(&Tx{tx, db})
	})
}

// Tx represents a transaction.
type Tx struct {
	*bolt.Tx
	db *DB
}

// GitHubClient returns an instance of the GitHub client for a given access token.
func (tx *Tx) GitHubClient(token string) *github.Client {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: token},
	}
	return github.NewClient(t.Client())
}

func (tx *Tx) userBucket() *bolt.Bucket          { return tx.Bucket([]byte("User")) }
func (tx *Tx) investigationBucket() *bolt.Bucket { return tx.Bucket([]byte("Investigation")) }
func (tx *Tx) experimentBucket() *bolt.Bucket    { return tx.Bucket([]byte("Experiment")) }

// User retrieves an user from the database by ID.
func (tx *Tx) User(id int64) *User {
	v := tx.userBucket().Get(i64tob(id))
	if v == nil {
		return nil
	}

	var u = &User{}
	if err := u.Unmarshal(v); err != nil {
		log.Println("user: unmarshal:", err)
		return nil
	}
	return u
}

// SaveUser stores an user in the database.
func (tx *Tx) SaveUser(u *User) error {
	if *u.ID == 0 {
		return fmt.Errorf("invalid user id: %d", *u.ID)
	}
	b, err := u.Marshal()
	if err != nil {
		return fmt.Errorf("marshal: %s", err)
	}
	return tx.userBucket().Put(i64tob(*u.ID), b)
}

// FindOrCreateUserByAccessToken retrieves or creates a new user based on an
// access token provided by GitHub.
func (tx *Tx) FindOrCreateUserByAccessToken(token string) (*User, error) {
	// Retrieve user info from GitHub.
	client := tx.GitHubClient(token)
	user, _, err := client.Users.Get("")
	if err != nil {
		return nil, fmt.Errorf("get user: %s", err)
	}

	// Create our own user record in the database.
	var u *User
	if u = tx.User(int64(*user.ID)); u == nil {
		u = &User{}
		u.SetID(int64(*user.ID))
	}
	u.SetEmail(*user.Email)
	u.SetAccessToken(token)

	// Save updated record.
	if err := tx.SaveUser(u); err != nil {
		return nil, fmt.Errorf("save user: %s", err)
	}
	return u, nil
}

// Investigation retrieves an investigation from the database by ID.
func (tx *Tx) Investigation(id string) *Investigation {
	v := tx.investigationBucket().Get([]byte(id))
	if v == nil {
		return nil
	}

	var i Investigation
	if err := i.Unmarshal(v); err != nil {
		log.Println("investigation: unmarshal:", err)
		return nil
	}
	return &i
}

// Parses a session for a given HTTP request.
func (tx *Tx) Session(r *http.Request) *Session {
	var s = &Session{}
	s.Session, _ = tx.db.store.Get(r, "default")
	if userID, ok := s.Values["UserID"].(int); ok {
		s.User = tx.User(int64(userID))
	}
	return s
}

// Session represents an authenticated session.
type Session struct {
	*sessions.Session
	User *User
}

// Converts an integer to a big-endian encoded byte slice.
func i64tob(v int64) []byte {
	var b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
