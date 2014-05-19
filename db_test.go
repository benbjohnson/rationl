package rationl_test

import (
	"io/ioutil"
	"os"
	"testing"

	"code.google.com/p/gogoprotobuf/proto"
	. "github.com/benbjohnson/rationl"
)

// Ensures that a user can be saved and retrieved.
func TestDB_SaveUser(t *testing.T) {
	var db = open()
	defer close(db)
	db.Update(func(tx *Tx) error {
		equals(t, nil, tx.SaveUser(&User{ID: proto.Int64(1), Email: proto.String("a@b.com"), AccessToken: proto.String("foo")}))
		return nil
	})
	db.View(func(tx *Tx) error {
		var u = tx.User(1)
		equals(t, int64(1), u.GetID())
		equals(t, "a@b.com", u.GetEmail())
		equals(t, "foo", u.GetAccessToken())
		return nil
	})
}

// open creates a temporary database.
func open() *DB {
	db, err := Open(tempfile(), 0600)
	if err != nil {
		panic("open: " + err.Error())
	}
	return db
}

// close closes a database and removes it from the filesystem.
func close(db *DB) {
	defer os.Remove(db.Path())
	db.Close()
}

// tempfile returns a path to a temporary file.
func tempfile() string {
	f, _ := ioutil.TempFile("", "rationl-")
	f.Close()
	os.Remove(f.Name())
	return f.Name()
}

func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		tb.Fatalf(msg, v...)
	}
}

func equals(tb testing.TB, exp, act interface{}) {
	assert(tb, exp == act, "exp: %#v, got: %#v", exp, act)
}
