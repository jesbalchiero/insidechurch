package events

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"insidechurch/backend/internal/domain"

	"github.com/google/uuid"
)

type PostgresEventStore struct {
	db *sql.DB
}

func NewPostgresEventStore(db *sql.DB) *PostgresEventStore {
	return &PostgresEventStore{db: db}
}

func (es *PostgresEventStore) Save(event domain.Event) error {
	query := `
		INSERT INTO events (
			id,
			aggregate_id,
			event_type,
			event_data,
			version,
			created_at
		) VALUES ($1, $2, $3, $4, $5, $6)
	`

	eventData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("erro ao serializar evento: %w", err)
	}

	_, err = es.db.Exec(
		query,
		uuid.New().String(),
		event.GetAggregateID(),
		event.GetEventType(),
		eventData,
		event.GetVersion(),
		time.Now(),
	)

	if err != nil {
		return fmt.Errorf("erro ao salvar evento: %w", err)
	}

	return nil
}

func (es *PostgresEventStore) GetEvents(aggregateID string) ([]domain.Event, error) {
	query := `
		SELECT event_type, event_data, version, created_at
		FROM events
		WHERE aggregate_id = $1
		ORDER BY version ASC
	`

	rows, err := es.db.Query(query, aggregateID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar eventos: %w", err)
	}
	defer rows.Close()

	var events []domain.Event
	for rows.Next() {
		var eventType string
		var eventData []byte
		var version int
		var createdAt time.Time

		err := rows.Scan(&eventType, &eventData, &version, &createdAt)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler evento: %w", err)
		}

		// Aqui você precisará implementar a lógica para desserializar o evento
		// baseado no eventType. Isso dependerá da sua implementação específica
		// de domain.Event e seus tipos concretos.

		// Exemplo de como poderia ser:
		// event, err := deserializeEvent(eventType, eventData)
		// if err != nil {
		//     return nil, fmt.Errorf("erro ao desserializar evento: %w", err)
		// }
		// events = append(events, event)
	}

	return events, nil
}
