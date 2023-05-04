package service

import (
	"context"

	"github.com/eduaravila/momo/apps/auth/internal/domain/session"
)

type OAuthServiceMock struct {
}

func (o *OAuthServiceMock) GetAccount(ctx context.Context, accessToken, accountUUID, userUUID string) (*session.Account, error) {
	return nil, nil
}

type TokenServiceMock struct {
}

func (t *TokenServiceMock) CreateSessionToken(ctx context.Context, subject string) (*session.Token, error) {
	return nil, nil
}

func (t *TokenServiceMock) VerifyToken(ctx context.Context, subject string) (*session.Token, error) {
	return nil, nil
}
