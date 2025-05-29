package domain

// Event representa um evento de domínio
type Event interface {
	GetAggregateID() string
	GetEventType() string
	GetVersion() int
}
