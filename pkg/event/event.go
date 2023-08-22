package event

import (
	"time"

	"github.com/gocql/gocql"
)

type Metadata map[string]string

type Event struct {
	ID               gocql.UUID `json:"id"`
	ActorID          string     `json:"actor_id"`
	TrackableOwnerID string     `json:"trackable_owner_id"`
	EventType        string     `json:"event_type"`
	TrackableType    string     `json:"trackable_type"`
	TrackableID      string     `json:"trackable_id"`
	CreatedAt        time.Time  `json:"created_at"`
	Metadata         Metadata   `json:"metadata"`
}
