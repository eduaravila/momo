package session

import (
	"errors"
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
) (*Session, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}

	if userUUID == "" {
		return nil, errors.New("userUUID is empty")
	}

	if expiredAt.IsZero() {
		return nil, errors.New("expiredAt is empty")
	}

	if sessionToken == nil {
		return nil, errors.New("sessionToken is empty")
	}

	return &Session{
		ID:           id,
		UserID:       userUUID,
		ExpiredAt:    expiredAt,
		SessionToken: sessionToken,
		IsValid:      sessionToken.Valid,
		Metadata:     NewSessionMetadata(userAgent, ipAddress),
	}, nil
}

func UnmarshalSessionFromDb(
	id string,
	userUUID string,
	createdAt time.Time,
	expiredAt time.Time,
	ipAddress string,
	userAgent string,
	isValid bool,
	sessionToken *Token,
) (*Session, error) {
	return NewSession(
		id,
		userUUID,
		expiredAt,
		sessionToken,
		userAgent,
		ipAddress,
	)
}
