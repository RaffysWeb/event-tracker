package models

import (
	"kafka_events/pkg/database"
	"time"

	"github.com/gocql/gocql"
)

type Metadata map[string]string

type Event struct {
	ID               gocql.UUID `json:"id"`
	ActorID          string     `json:"actorId"`
	TrackableOwnerID string     `json:"trackableOwnerId"`
	EventType        string     `json:"eventType"`
	TrackableType    string     `json:"trackableType"`
	TrackableID      string     `json:"trackableId"`
	CreatedAt        time.Time  `json:"createdAt"`
	Metadata         Metadata   `json:"metadata"`
}

// SaveEvent saves the Event data to the Cassandra database.
func (e *Event) SaveEvent() error {
	session := database.GetSession()

	query := session.Query(`
	INSERT INTO events_by_trackable (
			id,
			actor_id, 
			trackable_owner_id,
			event_type,
			trackable_type,
			trackable_id, 
			created_at,
			metadata) 
	VALUES (
			:id,
			:actor_id,
			:trackable_owner_id,
			:event_type,
			:trackable_type,
			:trackable_id,
			:created_at,
			:metadata
	)`).
		Bind(
			gocql.TimeUUID(),
			e.ActorID,
			e.TrackableOwnerID,
			e.EventType,
			e.TrackableType,
			e.TrackableID,
			time.Now(),
			e.Metadata,
		)

	return query.Exec()
}
