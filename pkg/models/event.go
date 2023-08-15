package models

import (
	"kafka_events/pkg/database"
)

type Metadata map[string]string

type Event struct {
	ID        string   `json:"id"`
	EventType string   `json:"eventType"`
	UserID    string   `json:"userId"`
	Metadata  Metadata `json:"metadata"`
}

// SaveEvent saves the Event data to the Cassandra database.
func (e *Event) SaveEvent() error {
	session := database.GetSession()

	query := session.Query("INSERT INTO events (id, eventType, userId, metadata) VALUES (?, ?, ?, ?)",
		e.ID, e.EventType, e.UserID, e.Metadata)

	return query.Exec()
}
