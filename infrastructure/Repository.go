package infrastructure

import (
	"github.com/mitubaEX/graphQL_sample/infrastructure/event"
	"github.com/mitubaEX/graphQL_sample/infrastructure/user"
)

var NewUserRepository func() user.UserRepository = user.NewUserRepositoryMem

func UseUserRepositoryMem() {
	NewUserRepository = user.NewUserRepositoryMem
}

var NewEventRepository func() event.EventRepository = event.NewEventRepositoryMem

func UseEventRepositoryMem() {
	NewEventRepository = event.NewEventRepositoryMem
}
