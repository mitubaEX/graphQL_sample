package event

import (
	"errors"
	"github.com/mitubaEX/graphQL_sample/domain/model/user"
	"github.com/google/uuid"
	"time"
)

type Event struct {
	EventId        string       `json:"eventId"`
	User           *user.User   `json:"user"`
	EventName      string       `json:"eventName"`
	Description    string       `json:"description"`
	Location       string       `json:"location"`
	StartTime      string       `json:"startTime"`
	EndTime        string       `json:"endTime"`
	Participants   []*user.User `json:"participants"`
	RegisteredTime int64        `json:"registeredTime"`
}

// constructor
func NewEvent(
	givenUser *user.User,
	eventName string,
	description string,
	location string,
	startTime string,
	endTime string) (*Event, error) {

	// validate
	if eventName == "" {
		return nil, errors.New("eventName is empty")
	}
	if description == "" {
		return nil, errors.New("description is empty")
	}
	if location == "" {
		return nil, errors.New("location is empty")
	}
	if startTime == "" {
		return nil, errors.New("startTime is empty")
	}
	if endTime == "" {
		return nil, errors.New("endTime is empty")
	}
	if givenUser == nil {
		return nil, errors.New("user is nil")
	}

	return &Event{
		EventId:        uuid.New().String(),
		User:           givenUser,
		EventName:      eventName,
		Description:    description,
		Location:       location,
		StartTime:      startTime,
		EndTime:        endTime,
		Participants:   []*user.User{},
		RegisteredTime: time.Now().Unix()}, nil
}

func (event *Event) Equals(other *Event) bool {
	if event.EventId == other.EventId {
		return true
	}
	return false
}
