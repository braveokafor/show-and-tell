package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/braveokafor/show-and-tell/internal/db"
)

// index redirects to the sessions listing
func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, "/sessions", http.StatusSeeOther)
}

// listSessions displays all sessions
func (h *Handler) listSessions(w http.ResponseWriter, r *http.Request) {
	sessions, err := db.GetAllSessions(h.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Sessions interface{}
	}{
		Sessions: sessions,
	}

	h.renderTemplate(w, "index", data)
}

// newSessionForm displays the form to create a new session
func (h *Handler) newSessionForm(w http.ResponseWriter, r *http.Request) {
	h.renderTemplate(w, "session_form", nil)
}

// createSession processes the form submission to create a new session
func (h *Handler) createSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	timeSlot, err := time.Parse("2006-01-02T15:04", r.FormValue("time_slot"))
	if err != nil {
		http.Error(w, "Invalid time format", http.StatusBadRequest)
		return
	}

	team := r.FormValue("team")
	if team == "" {
		http.Error(w, "Team is required", http.StatusBadRequest)
		return
	}

	topic := r.FormValue("topic")
	presenter := r.FormValue("presenter")

	_, err = db.CreateSession(h.db, timeSlot, team, topic, presenter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/sessions", http.StatusSeeOther)
}

// editSessionForm displays the form to edit an existing session
func (h *Handler) editSessionForm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/sessions/edit/"))
	if err != nil {
		http.Error(w, "Invalid session ID", http.StatusBadRequest)
		return
	}

	session, err := db.GetSession(h.db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Session interface{}
	}{
		Session: session,
	}

	h.renderTemplate(w, "session_form", data)
}

// updateSession processes the form submission to update an existing session
func (h *Handler) updateSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/sessions/update/"))
	if err != nil {
		http.Error(w, "Invalid session ID", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	timeSlot, err := time.Parse("2006-01-02T15:04", r.FormValue("time_slot"))
	if err != nil {
		http.Error(w, "Invalid time format", http.StatusBadRequest)
		return
	}

	team := r.FormValue("team")
	if team == "" {
		http.Error(w, "Team is required", http.StatusBadRequest)
		return
	}

	topic := r.FormValue("topic")
	presenter := r.FormValue("presenter")

	err = db.UpdateSession(h.db, id, timeSlot, team, topic, presenter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/sessions", http.StatusSeeOther)
}

// deleteSession removes a session
func (h *Handler) deleteSession(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/sessions/delete/"))
	if err != nil {
		http.Error(w, "Invalid session ID", http.StatusBadRequest)
		return
	}

	err = db.DeleteSession(h.db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/sessions", http.StatusSeeOther)
}
