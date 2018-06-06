package user

import (
	"errors"
	"github.com/mitubaEX/graphQL_sample/domain/model/user"
)

type UserRepositoryMem struct {
	users []*user.User
}

func NewUserRepositoryMem() UserRepository {
	return &UserRepositoryMem{[]*user.User{}}
}

// store event to repository
func (self *UserRepositoryMem) Store(user *user.User) UserRepository {
	self.users = append(self.users, user)
	return self
}

func (self UserRepositoryMem) FindById(userId string) (*user.User, error) {
	for _, val := range self.users {
		if val.UserId == userId {
			return val, nil
		}
	}

	return nil, errors.New("user not found")
}

func (self UserRepositoryMem) UserList() []*user.User {
	return self.users
}
