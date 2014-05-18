package handlers

import (
	"net/http"
	"path"

	"github.com/benbjohnson/rationl"
	"github.com/benbjohnson/rationl/assets"
	"github.com/benbjohnson/rationl/handlers/templates"
	"github.com/gorilla/mux"
)

const authorizeUrl = "/authorize"

// NewHandler returns a new root HTTP handler.
func NewHandler(db *rationl.DB, clientID, secret string) http.Handler {
	var authorizeHandler = newAuthorizeHandler(db, clientID, secret)

	r := mux.NewRouter()
	r.Handle("/", &indexHandler{db})
	r.HandleFunc("/assets/{filename}", assetsHandleFunc)

	r.Handle("/authorize", authorizeHandler)
	r.Handle("/authorize/callback", authorizeHandler)

	r.Handle("/investigations", &investigationsHandler{db})
	return r
}

// indexHandler handles the rendering of the home page.
type indexHandler struct {
	db *rationl.DB
}

func (h *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.db.View(func(tx *rationl.Tx) error {
		return templates.Index(w, tx.Session(r))
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
