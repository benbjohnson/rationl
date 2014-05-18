package handlers

import (
	"errors"
	"net/http"
	"path"

	"github.com/benbjohnson/rationl"
	"github.com/benbjohnson/rationl/assets"
	"github.com/benbjohnson/rationl/handlers/templates"
	"github.com/gorilla/mux"
)

// AuthorizeURL is the path to the authorization URL.
const AuthorizeURL = "/authorize"

var (
	// ErrUnauthorized is returned when the user is not logged in or is not
	// authorized to access a given resource.
	ErrUnauthorized = errors.New("unauthorized")

	// ErrValidation is returned when a saved resource is invalid. It is used
	// as a marker since validation errors are handled by rerendering the
	// template instead of displaying an error.
	ErrInvalid = errors.New("invalid resource")
)

// NewHandler returns a new root HTTP handler.
func NewHandler(db *rationl.DB, clientID, secret string) http.Handler {
	var authorizeHandler = newAuthorizeHandler(db, clientID, secret)

	r := mux.NewRouter()
	r.Handle("/", &indexHandler{db})
	r.HandleFunc("/assets/{filename}", assetsHandleFunc)

	r.Handle("/authorize", authorizeHandler)
	r.Handle("/authorize/callback", authorizeHandler)

	r.Handle("/investigations", &InvestigationsHandler{db}).Methods("GET")
	r.Handle("/investigations", &CreateInvestigationHandler{db}).Methods("POST")
	r.Handle("/investigations/new", &NewInvestigationHandler{db}).Methods("GET")
	r.Handle("/investigations/{id}", &InvestigationHandler{db}).Methods("GET")
	// r.Handle("/investigations/{id}", &UpdateInvestigationHandler{db}).Methods("PATCH")
	r.Handle("/investigations/{id}/edit", &EditInvestigationHandler{db}).Methods("GET")
	// r.Handle("/investigations/{id}", &DeleteInvestigationHandler{db}).Methods("DELETE")
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

// Error translates errors into HTTP responses.
func Error(w http.ResponseWriter, err error) {
	switch err {
	case ErrInvalid: // do nothing.
	default:
		templates.Error(w, err)
	}
}
