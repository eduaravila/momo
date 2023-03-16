package session

import (
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
