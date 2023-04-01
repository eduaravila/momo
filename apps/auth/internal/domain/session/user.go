package session

import (
	"errors"
	"time"
)

type User struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ClientMetadata struct {
	UserAgent string
	IPAddress string
}

func UnmarshalUserFromDatabase(
	id string,
	createdAt time.Time,
	updatedAt time.Time,
) (*User, error) {
	return NewUser(id, createdAt, updatedAt)
}

func NewUser(
	id string,
	createdAt time.Time,
	updatedAt time.Time,
) (*User, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}

	if createdAt.IsZero() {
		return nil, errors.New("createdAt is empty")
	}

	if updatedAt.IsZero() {
		return nil, errors.New("updatedAt is empty")
	}

	return &User{
		ID:        id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
