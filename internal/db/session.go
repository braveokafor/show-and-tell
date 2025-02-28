package db

import (
	"database/sql"
	"time"

	"github.com/braveokafor/show-and-tell/internal/models"
)

// GetAllSessions retrieves all sessions from the database
func GetAllSessions(db *sql.DB) ([]models.Session, error) {
	rows, err := db.Query("SELECT id, time_slot, team, topic, presenter FROM sessions ORDER BY time_slot")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []models.Session
	for rows.Next() {
		var s models.Session
		if err := rows.Scan(&s.ID, &s.TimeSlot, &s.Team, &s.Topic, &s.Presenter); err != nil {
			return nil, err
		}
		sessions = append(sessions, s)
	}

	return sessions, nil
}

// GetSession retrieves a single session by ID
func GetSession(db *sql.DB, id int) (models.Session, error) {
	var s models.Session
	err := db.QueryRow("SELECT id, time_slot, team, topic, presenter FROM sessions WHERE id = $1", id).
		Scan(&s.ID, &s.TimeSlot, &s.Team, &s.Topic, &s.Presenter)
	return s, err
}

// CreateSession adds a new session to the database
func CreateSession(db *sql.DB, timeSlot time.Time, team, topic, presenter string) (int, error) {
	var id int
	err := db.QueryRow(
		"INSERT INTO sessions (time_slot, team, topic, presenter) VALUES ($1, $2, $3, $4) RETURNING id",
		timeSlot, team, topic, presenter,
	).Scan(&id)
	return id, err
}

// UpdateSession modifies an existing session
func UpdateSession(db *sql.DB, id int, timeSlot time.Time, team, topic, presenter string) error {
	_, err := db.Exec(
		"UPDATE sessions SET time_slot = $1, team = $2, topic = $3, presenter = $4 WHERE id = $5",
		timeSlot, team, topic, presenter, id,
	)
	return err
}

// DeleteSession removes a session from the database
func DeleteSession(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM sessions WHERE id = $1", id)
	return err
}
