package event

import (
	"github.com/mitubaEX/graphQL_sample/domain/model/event"
)

type EventRepository interface {
	Store(event *event.Event) EventRepository
	FindById(eventId string) (*event.Event, error)
	EventList() []*event.Event
}
