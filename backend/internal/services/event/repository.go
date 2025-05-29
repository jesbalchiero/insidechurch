package event

import "context"

type Repository interface {
	Create(ctx context.Context, event *Event) error
	GetByID(ctx context.Context, id string) (*Event, error)
	List(ctx context.Context, filter EventFilter) ([]*Event, error)
	Update(ctx context.Context, event *Event) error
	Delete(ctx context.Context, id string) error
}
