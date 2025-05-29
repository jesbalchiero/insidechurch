package events

import "sync"

type EventHandler func(event Event)

type EventDispatcher struct {
	handlers map[string][]EventHandler
	mu       sync.RWMutex
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (d *EventDispatcher) Register(eventType string, handler EventHandler) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, exists := d.handlers[eventType]; !exists {
		d.handlers[eventType] = make([]EventHandler, 0)
	}
	d.handlers[eventType] = append(d.handlers[eventType], handler)
}

func (d *EventDispatcher) Dispatch(event Event) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if handlers, exists := d.handlers[event.GetEventType()]; exists {
		for _, handler := range handlers {
			handler(event)
		}
	}
}
