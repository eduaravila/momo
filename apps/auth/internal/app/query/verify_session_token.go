package query

import (
	"context"

	"github.com/eduaravila/momo/apps/auth/internal/domain/session"
	"github.com/eduaravila/momo/packages/decorators"
	"golang.org/x/exp/slog"
)

type SessionToken struct {
	Token string
}

type SessionTokenVerifierHandlers decorators.QueryHandler[SessionToken, *session.Token]

type TokenVeriferService interface {
	VerifyToken(ctx context.Context, token string) (*session.Token, error)
}
type sessionTokenVerifierHandler struct {
	tokenService TokenVeriferService
}

func NewSessionTokenVerifierHandler(
	tokenService TokenVeriferService,
	logger *slog.Logger,
	metrics decorators.MetricsClient,
) SessionTokenVerifierHandlers {
	return decorators.ApplyQueryDecorators[SessionToken, *session.Token](
		&sessionTokenVerifierHandler{
			tokenService: tokenService,
		},
		logger,
		metrics,
	)
}

func (s *sessionTokenVerifierHandler) Handle(
	ctx context.Context,
	query SessionToken,
) (*session.Token, error) {
	return s.tokenService.VerifyToken(ctx, query.Token)
}
