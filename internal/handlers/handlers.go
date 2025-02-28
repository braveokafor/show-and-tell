package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

// Handler contains dependencies for our HTTP handlers
type Handler struct {
	db        *sql.DB
	templates map[string]*template.Template
}

// NewHandler creates a new handler with the given database connection
func NewHandler(db *sql.DB) *Handler {
	h := &Handler{
		db:        db,
		templates: make(map[string]*template.Template),
	}

	// Parse templates
	baseLayout := "ui/templates/base.html"
	templateFiles := []string{"index", "session_form"}

	for _, tmpl := range templateFiles {
		h.templates[tmpl] = template.Must(template.ParseFiles(
			baseLayout,
			filepath.Join("ui/templates", tmpl+".html"),
		))
	}

	return h
}

// Routes sets up the HTTP routes for the application
func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	// Static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Session CRUD routes
	mux.HandleFunc("/", h.index)
	mux.HandleFunc("/sessions", h.listSessions)
	mux.HandleFunc("/sessions/new", h.newSessionForm)
	mux.HandleFunc("/sessions/create", h.createSession)
	mux.HandleFunc("/sessions/edit/", h.editSessionForm)
	mux.HandleFunc("/sessions/update/", h.updateSession)
	mux.HandleFunc("/sessions/delete/", h.deleteSession)

	// Demo endpoints
	mux.HandleFunc("/demo/panic", h.triggerPanic)
	mux.HandleFunc("/demo/memory", h.allocateMemory)
	mux.HandleFunc("/health", h.healthCheck)

	return mux
}

// renderTemplate renders a template with the given data
func (h *Handler) renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	if t, ok := h.templates[tmpl]; ok {
		err := t.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Template not found", http.StatusInternalServerError)
	}
}
