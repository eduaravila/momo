package ports

import (
	"context"
	"strings"

	"github.com/eduaravila/momo/apps/auth/internal/app"
	"github.com/eduaravila/momo/apps/auth/internal/app/query"
	"github.com/eduaravila/momo/packages/genproto/auth"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GRPCServer struct {
	app app.Application
}

func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{app}
}

func (s *GRPCServer) IsValidSession(
	ctx context.Context,
	request *auth.IsValidSessionTokenRequest,
) (*auth.IsValidSessionTokenResponse, error) {
	query := query.SessionToken{Token: request.Token}

	token, err := s.app.Queries.VerifySessionToken.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	if token == nil {
		return &auth.IsValidSessionTokenResponse{
			Valid: false,
		}, nil
	}

	return &auth.IsValidSessionTokenResponse{
		Valid: token.Valid,
		Claims: &auth.IsValidSessionTokenResponse_Claims{
			Subject:   token.Claims.Subject,
			Issuer:    token.Claims.Issuer,
			Audience:  strings.Join(token.Claims.Audience, ","),
			ExpiresAt: timestamppb.New(token.Claims.ExpiresAt),
			IssuedAt:  timestamppb.New(token.Claims.IssuedAt),
		},
	}, nil
}
