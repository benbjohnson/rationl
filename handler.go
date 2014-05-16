package rationl

import (
	"net/http"
	"path"

	"github.com/benbjohnson/rationl/assets"
	"github.com/gorilla/mux"
)

// NewHandler returns a new root HTTP handler.
func NewHandler(db *DB) http.Handler {
	r := mux.NewRouter()
	r.Handle("/", &indexHandler{db})
	r.HandleFunc("/assets/{filename}", assetsHandleFunc)
	return r
}

// indexHandler handles the rendering of the home page.
type indexHandler struct {
	db *DB
}

func (h *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.db.View(func(tx *Tx) error {
		return index(w, tx.Session(r))
	})
}

// Returns static file content to the client.
func assetsHandleFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	b, err := assets.Asset(filename)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	switch path.Ext(filename) {
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":
		w.Header().Set("Content-Type", "application/javascript")
	}
	w.Write(b)
}
