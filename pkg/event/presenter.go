package event

import (
	"time"
)

type EventsByActionType struct {
	EventType string    `json:"event_type"`
	Time      time.Time `json:"_time"`
	Value     int       `json:"_value"`
}
