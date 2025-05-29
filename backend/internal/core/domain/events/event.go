package events

import "time"

// Event é a interface base para todos os eventos de domínio
type Event interface {
	GetTimestamp() time.Time
	GetEventType() string
}

// BaseEvent implementa a interface Event com campos comuns
type BaseEvent struct {
	Timestamp time.Time
	EventType string
}

func NewBaseEvent(eventType string) BaseEvent {
	return BaseEvent{
		Timestamp: time.Now(),
		EventType: eventType,
	}
}

func (e BaseEvent) GetTimestamp() time.Time {
	return e.Timestamp
}

func (e BaseEvent) GetEventType() string {
	return e.EventType
}
