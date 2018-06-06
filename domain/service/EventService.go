package service

import (
	"github.com/mitubaEX/graphQL_sample/domain/model/event"
	"github.com/mitubaEX/graphQL_sample/infrastructure"
)

func FindEventById(userId string) (*event.Event, error) {
	repository := infrastructure.NewEventRepository()
	return repository.FindById(userId)
}
