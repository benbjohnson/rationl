package handlers

import (
	"net/http"

	"code.google.com/p/gogoprotobuf/proto"
	"github.com/benbjohnson/rationl"
	"github.com/benbjohnson/rationl/handlers/templates"
	"github.com/gorilla/mux"
)

// InvestigationHandler renders a single investigation.
type InvestigationHandler struct {
	db *rationl.DB
}

func (h *InvestigationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.db.View(func(tx *rationl.Tx) error {
		var vars = mux.Vars(r)
		var session = tx.Session(r)
		if !session.Authenticated() {
			return ErrUnauthorized
		}
		return templates.Investigation(w, session, tx.Investigation(vars["id"]))
	})
	if err != nil {
		Error(w, err)
	}
}

// InvestigationsHandler renders a list of investigations for the current user.
type InvestigationsHandler struct {
	db *rationl.DB
}

func (h *InvestigationsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.db.View(func(tx *rationl.Tx) error {
		var session = tx.Session(r)
		if !session.Authenticated() {
			return ErrUnauthorized
		}
		return templates.Investigations(w, session, tx.InvestigationsByUserID(session.User.GetID()))
	})
	if err != nil {
		Error(w, err)
	}
}

// NewInvestigationHandler renders an editable form for a new investigation.
type NewInvestigationHandler struct {
	db *rationl.DB
}

func (h *NewInvestigationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.db.View(func(tx *rationl.Tx) error {
		var session = tx.Session(r)
		if !session.Authenticated() {
			return ErrUnauthorized
		}

		var i = &rationl.Investigation{UserID: proto.Int64(session.User.GetID())}
		return templates.NewInvestigation(w, session, i)
	})
	if err != nil {
		Error(w, err)
	}
}

// EditInvestigationHandler renders an editable form for an existing investigation.
type EditInvestigationHandler struct {
	db *rationl.DB
}

func (h *EditInvestigationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.db.View(func(tx *rationl.Tx) error {
		var vars = mux.Vars(r)
		var session = tx.Session(r)
		if !session.Authenticated() {
			return ErrUnauthorized
		}

		// Find and authorize investigation.
		var i = tx.Investigation(vars["id"])
		if i.GetUserID() != session.User.GetID() {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return nil
		}

		return templates.EditInvestigation(w, session, i)
	})
	if err != nil {
		Error(w, err)
	}
}

// CreateInvestigationHandler creates a new investigation.
type CreateInvestigationHandler struct {
	db *rationl.DB
}

func (h *CreateInvestigationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.db.Update(func(tx *rationl.Tx) error {
		var session = tx.Session(r)
		if !session.Authenticated() {
			return ErrUnauthorized
		}

		// Create investigation and save.
		var i = &rationl.Investigation{}
		i.UserID = proto.Int64(session.User.GetID())
		i.Name = proto.String(r.FormValue("name"))
		if err := tx.CreateInvestigation(i); err != nil {
			session.Error = err
			templates.NewInvestigation(w, session, i)
			return ErrInvalid
		}

		http.Redirect(w, r, "/investigations/"+i.GetID(), http.StatusFound)
		return nil
	})
	if err != nil {
		Error(w, err)
		return
	}
}
