package common

import (
	"errors"

	"github.com/google/uuid"
)

type UserID struct {
	value string
}

func (id UserID) String() string {
	return string(id.value)
}

func NovoUserID() (UserID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return UserID{}, err
	}
	return UserID{value: id.String()}, nil
}

func NovoExternalUserID(value string) (UserID, error) {
	if value == "" {
		return UserID{}, errors.New("user id vazio")
	}
	return UserID{value: value}, nil
}
