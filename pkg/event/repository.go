package event

import (
	"context"
	"kafka_events/pkg/database/cassandra"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, event *Event) error
}

type EventRepositoryImpl struct{}

func NewEventRepository() EventRepository {
	return &EventRepositoryImpl{}
}

// SaveEvent saves the Event data to the Cassandra database.
func (r *EventRepositoryImpl) CreateEvent(ctx context.Context, e *Event) error {
	session := cassandra.GetSession()

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

	return query.WithContext(ctx).Exec()
}
