package event

import (
	"errors"
	"github.com/mitubaEX/graphQL_sample/domain/model/event"
)

type EventRepositoryMem struct {
	events []*event.Event
}

func NewEventRepositoryMem() EventRepository {
	return &EventRepositoryMem{[]*event.Event{}}
}

// store event to repository
func (self *EventRepositoryMem) Store(event *event.Event) EventRepository {
	self.events = append(self.events, event)
	return self
}

func (self EventRepositoryMem) FindById(userId string) (*event.Event, error) {
	for _, val := range self.events {
		if val.EventId == userId {
			return val, nil
		}
	}

	return nil, errors.New("user not found")
}

func (self EventRepositoryMem) EventList() []*event.Event {
	return self.events
}
