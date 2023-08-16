package models

import (
	"kafka_events/pkg/database"
	"time"
)

type Metadata map[string]string

type Event struct {
	ID               string    `json:"id" avro:"id"`
	ActorID          string    `json:"actor_id" avro:"actor_id"`
	TrackableOwnerID string    `json:"trackable_owner_id" avro:"trackable_owner_id"`
	EventType        string    `json:"event_type" avro:"event_type"`
	TrackableType    string    `json:"trackable_type" avro:"trackable_type"`
	TrackableID      string    `json:"trackable_id" avro:"trackable_id"`
	CreatedAt        time.Time `json:"created_at" avro:"created_at"`
	Metadata         Metadata  `json:"metadata" avro:"metadata"`
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
		)`).Bind(
		e.ID,
		e.ActorID,
		e.TrackableOwnerID,
		e.EventType,
		e.TrackableType,
		e.TrackableID,
		e.CreatedAt,
		e.Metadata,
	)

	return query.Exec()
}
