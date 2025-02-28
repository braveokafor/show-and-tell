package models

import "time"

// Session represents a show and tell session
type Session struct {
	ID        int       `json:"id"`
	TimeSlot  time.Time `json:"time_slot"`
	Team      string    `json:"team"`
	Topic     string    `json:"topic"`
	Presenter string    `json:"presenter"`
}
