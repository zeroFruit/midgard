package eventsource

type EventStore interface {
	Save(aggregateID string, events ...Event) error
	Load(aggregateID string) ([]Event, error)
}
