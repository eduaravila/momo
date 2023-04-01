package session

import (
	"time"
)

type Session struct {
	ID           string
	UserID       string
	ExpiredAt    time.Time
	SessionToken *Token
	Metadata     *ClientMetadata
	IsValid      bool
}

func NewSessionMetadata(
	userAgent string,
	ipAddress string,
) *ClientMetadata {
	return &ClientMetadata{
		UserAgent: userAgent,
		IPAddress: ipAddress,
	}
}

func NewSession(
	id string,
	userUUID string,
	expiredAt time.Time,
	sessionToken *Token,
	userAgent string,
	ipAddress string,
) *Session {
	return &Session{
		ID:           id,
		UserID:       userUUID,
		ExpiredAt:    expiredAt,
		SessionToken: sessionToken,
		IsValid:      sessionToken.Valid,
		Metadata:     NewSessionMetadata(userAgent, ipAddress),
	}
}
