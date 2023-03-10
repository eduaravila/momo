package session

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	CreatedAt    time.Time
	ExpiredAt    time.Time
	SessionToken string
	IpAddress    string
	UserAgent    string
	IsValid      bool
}
