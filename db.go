package rationl

import (
	"encoding/binary"
	"fmt"
	"net/http"
	"os"

	"github.com/boltdb/bolt"
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

// UserBucket returns the bucket used to store Users.
func (tx *Tx) UserBucket() *bolt.Bucket { return tx.Bucket([]byte("User")) }

// User retrieves an user from the database by ID.
func (tx *Tx) User(id int) *User {
	if v := tx.UserBucket().Get(itob(id)); v != nil {
		var u = &User{}
		if err := u.Unmarshal(v); err != nil {
			warn("user: unmarshal:", err)
			return nil
		}
		return u
	}
	return nil
}

// Parses a session for a given HTTP request.
func (tx *Tx) Session(r *http.Request) *Session {
	var session = &Session{}
	s, _ := tx.db.store.Get(r, "default")
	if userID, ok := s.Values["UserID"].(int); ok {
		session.User = tx.User(userID)
	}
	return session
}

// Session represents an authenticated session.
type Session struct {
	User *User
}

// Converts an integer to a big-endian encoded byte slice.
func itob(v int) []byte {
	var b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
