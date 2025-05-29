package event

import (
	"context"
	"time"
)

type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Location    string    `json:"location"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Service interface {
	Create(ctx context.Context, event *Event) error
	GetByID(ctx context.Context, id string) (*Event, error)
	List(ctx context.Context, filter EventFilter) ([]*Event, error)
	Update(ctx context.Context, event *Event) error
	Delete(ctx context.Context, id string) error
}

type EventFilter struct {
	StartDate *time.Time
	EndDate   *time.Time
	Location  string
	CreatedBy string
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, event *Event) error {
	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()
	return s.repo.Create(ctx, event)
}

func (s *service) GetByID(ctx context.Context, id string) (*Event, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service) List(ctx context.Context, filter EventFilter) ([]*Event, error) {
	return s.repo.List(ctx, filter)
}

func (s *service) Update(ctx context.Context, event *Event) error {
	event.UpdatedAt = time.Now()
	return s.repo.Update(ctx, event)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
