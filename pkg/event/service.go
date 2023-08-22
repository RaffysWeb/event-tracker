package event

import (
	"context"
)

type EventService interface {
	CreateEvent(ctx context.Context, event *Event) (*Event, error)
}

type EventServiceImpl struct {
	eventRepository EventRepository
}

func NewEventService(eventRepository EventRepository) *EventServiceImpl {
	return &EventServiceImpl{eventRepository}
}

// CreateEvent creates a new Event.
func (s *EventServiceImpl) CreateEvent(ctx context.Context, event *Event) (*Event, error) {
	err := s.eventRepository.CreateEvent(ctx, event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
