package app

import "github.com/eduaravila/momo/apps/auth/internal/oauth/app/command"

type Application struct {
	Queries Queries
	Command Commands
}

type Commands struct {
	AuthenticateWithOIDC command.AuthenticateWithOIDCHandler
}

type Queries struct {
}
