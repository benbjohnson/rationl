package handlers

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"

	"code.google.com/p/goauth2/oauth"
	"github.com/benbjohnson/rationl"
)

// authorizeHandler represents the handler used for authorizing GitHub accounts.
type authorizeHandler struct {
	db     *rationl.DB
	Config *oauth.Config
}

func newAuthorizeHandler(db *rationl.DB, clientID, secret string) http.Handler {
	return &authorizeHandler{
		db: db,
		Config: &oauth.Config{
			ClientId:     clientID,
			ClientSecret: secret,
			Scope:        "user:email",
			AuthURL:      "https://github.com/login/oauth/authorize",
			TokenURL:     "https://github.com/login/oauth/access_token",
		},
	}
}

// ServeHTTP handles authorization requests.
func (h *authorizeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/authorize":
		h.redirect(w, r)
	case "/authorize/callback":
		h.complete(w, r)
	default:
		log.Println("not found:", r.URL.Path)
		http.NotFound(w, r)
	}
}

// Handles the initial request for authorization. Verifies that the user is
// logged in and then redirects to GitHub for authorization.
func (h *authorizeHandler) redirect(w http.ResponseWriter, r *http.Request) {
	// Generate auth state.
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var state = fmt.Sprintf("%x", b)

	// Save state to the session.
	err := h.db.View(func(tx *rationl.Tx) error {
		var session = tx.Session(r)
		session.Values["AuthState"] = state
		return session.Save(r, w)
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect user to GitHub for OAuth authorization.
	http.Redirect(w, r, h.Config.AuthCodeURL(state), http.StatusFound)
}

// Handles the redirect back from GitHub OAuth authorization and saves the
// access token to the account.
func (h *authorizeHandler) complete(w http.ResponseWriter, r *http.Request) {
	err := h.db.Update(func(tx *rationl.Tx) error {
		var session = tx.Session(r)
		var state, _ = session.Values["AuthState"].(string)

		// Verify that the auth code was not tampered with.
		if r.FormValue("state") != state {
			return fmt.Errorf("invalid oauth state")
		}

		// Extract the access token.
		var t = &oauth.Transport{Config: h.Config}
		token, err := t.Exchange(r.FormValue("code"))
		if err != nil {
			return fmt.Errorf("exchange: %s", err)
		}

		// Create user from token.
		u, err := tx.FindOrCreateUserByAccessToken(token.AccessToken)
		if err != nil {
			return fmt.Errorf("create user: %s", err)
		}
		session.Values["UserID"] = int(u.GetID())
		session.Save(r, w)

		return nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to home page.
	http.Redirect(w, r, "/investigations", http.StatusFound)
}
