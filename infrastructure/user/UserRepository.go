package user

import (
	"github.com/mitubaEX/graphQL_sample/domain/model/user"
)

type UserRepository interface {
	Store(user *user.User) UserRepository
	FindById(userId string) (*user.User, error)
	UserList() []*user.User
}
