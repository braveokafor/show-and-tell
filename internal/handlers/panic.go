package handlers

import (
	"net/http"
)

// triggerPanic intentionally triggers a panic for demo purposes
func (h *Handler) triggerPanic(w http.ResponseWriter, r *http.Request) {
	panic("This is a demo panic for failover testing")
}
