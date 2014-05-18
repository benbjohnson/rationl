package handlers

import (
	"net/http"

	"github.com/benbjohnson/rationl"
	"github.com/benbjohnson/rationl/handlers/templates"
)

// investigationsHandler renders a list of investigations for the user.
type investigationsHandler struct {
	db *rationl.DB
}

func (h *investigationsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.db.View(func(tx *rationl.Tx) error {
		var session = tx.Session(r)
		if session.User == nil {
			http.Redirect(w, r, authorizeUrl, http.StatusFound)
			return nil
		}

		var a []*rationl.Investigation
		for _, id := range session.User.GetInvestigationIDs() {
			a = append(a, tx.Investigation(id))
		}
		return templates.Investigations(w, session, a)
	})
}
