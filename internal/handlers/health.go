package handlers

import (
	"encoding/json"
	"net/http"
)

// healthCheck responds with application health information
func (h *Handler) healthCheck(w http.ResponseWriter, r *http.Request) {
	// Check database connection
	err := h.db.Ping()
	dbStatus := "up"
	if err != nil {
		dbStatus = "down"
	}

	// Create health response
	health := struct {
		Status   string `json:"status"`
		Database string `json:"database"`
	}{
		Status:   "ok",
		Database: dbStatus,
	}

	// If the database is down, set a non-200 status code
	if dbStatus == "down" {
		health.Status = "degraded"
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	// Write the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}
