package user

import (
	"errors"
	"github.com/google/uuid"
)

type User struct {
	UserId      string `json:"userId"`
	UserName    string `json:"userName"`
	Description string `json:"description"`
	PhotoURL    string `json:"photoURL"`
	Email       string `json:"email"`
}

// constructor
func NewUser(
	userName string,
	description string,
	photoURL string,
	email string) (*User, error) {

	// validate
	if userName == "" {
		return &User{"", "", "", "", ""}, errors.New("eventName is empty")
	}
	if description == "" {
		return &User{"", "", "", "", ""}, errors.New("description is empty")
	}
	if photoURL == "" {
		return &User{"", "", "", "", ""}, errors.New("location is empty")
	}
	if email == "" {
		return &User{"", "", "", "", ""}, errors.New("startTime is empty")
	}

	return &User{
		UserId:      uuid.New().String(),
		UserName:    userName,
		Description: description,
		PhotoURL:    photoURL,
		Email:       email}, nil
}

func (user *User) Equals(other User) bool {
	if user.UserId == other.UserId {
		return true
	}
	return false
}
