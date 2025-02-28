package handlers

import (
	"net/http"
	"os"
)

// triggerPanic intentionally triggers a panic for demo purposes
func (h *Handler) triggerPanic(w http.ResponseWriter, r *http.Request) {
	os.Exit(1)
}
