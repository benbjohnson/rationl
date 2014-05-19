package rationl

import (
	"encoding/binary"
	"fmt"
	"log"
	"net/http"
	"os"

	"code.google.com/p/goauth2/oauth"
	"code.google.com/p/gogoprotobuf/proto"
	"github.com/boltdb/bolt"
	"github.com/google/go-github/github"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/nu7hatch/gouuid"
)

var (
	usrbkt    = []byte("User")
	usrinvbkt = []byte("User.[]Investigation")
	invbkt    = []byte("Investigation")
	invexpbkt = []byte("Investigation.[]Experiment")
	expbkt    = []byte("Experiment")
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
		tx.CreateBucketIfNotExists(usrbkt)
		tx.CreateBucketIfNotExists(usrinvbkt)
		tx.CreateBucketIfNotExists(invbkt)
		tx.CreateBucketIfNotExists(invexpbkt)
		tx.CreateBucketIfNotExists(expbkt)

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

// User retrieves an user from the database by ID.
func (tx *Tx) User(id int64) *User {
	v := tx.Bucket(usrbkt).Get(i64tob(id))
	if v == nil {
		return nil
	}

	var u = &User{}
	if err := u.Unmarshal(v); err != nil {
		log.Println("usr: unmarshal:", err)
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
	return tx.Bucket(usrbkt).Put(i64tob(u.GetID()), b)
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
		u.ID = proto.Int64(int64(*user.ID))
	}
	u.Email = proto.String(*user.Email)
	u.AccessToken = proto.String(token)

	// Save updated record.
	if err := tx.SaveUser(u); err != nil {
		return nil, fmt.Errorf("save user: %s", err)
	}
	return u, nil
}

// Investigation retrieves an investigation from the database by ID.
func (tx *Tx) Investigation(id string) *Investigation {
	v := tx.Bucket(invbkt).Get([]byte(id))
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

// InvestigationsByUserID retrieves a list of investigations by user id.
func (tx *Tx) InvestigationsByUserID(id int64) []*Investigation {
	b := tx.Bucket(usrinvbkt).Bucket(i64tob(id))
	if b == nil {
		return nil
	}
	var a []*Investigation
	b.ForEach(func(k, _ []byte) error {
		a = append(a, tx.Investigation(string(k)))
		return nil
	})
	return a
}

// CreateInvestigation creates a new investigation.
func (tx *Tx) CreateInvestigation(i *Investigation) error {
	if tx.User(i.GetUserID()) == nil {
		return fmt.Errorf("invalid user: %d", i.GetUserID())
	} else if i.GetName() == "" {
		return fmt.Errorf("name required")
	}

	// Create a UUID.
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	i.ID = proto.String(id.String())

	// Save investigation
	v, err := i.Marshal()
	if err != nil {
		return fmt.Errorf("marshal: %s", err)
	}
	if err := tx.Bucket(invbkt).Put([]byte(i.GetID()), v); err != nil {
		return fmt.Errorf("inv: put: %s", err)
	}

	// Add to index.
	b, err := tx.Bucket(usrinvbkt).CreateBucketIfNotExists(i64tob(i.GetUserID()))
	if err != nil {
		return fmt.Errorf("usrinv: create: %s", err)
	}
	if err := b.Put([]byte(i.GetID()), []byte{}); err != nil {
		return fmt.Errorf("usrinv: put: %s", err)
	}

	return nil
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
	User  *User
	Error error
}

// Authenticated returns true if there is a user attached to the session.
func (s *Session) Authenticated() bool {
	return s.User != nil
}

// Converts an integer to a big-endian encoded byte slice.
func i64tob(v int64) []byte {
	var b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
