package app

import "github.com/eduaravila/momo/apps/auth/internal/app/command"

type Application struct {
	Queries  Queries
	Commands Commands
}

type Commands struct {
	AuthenticateWithOIDC command.AuthenticateWithOIDCHandler
}

type Queries struct {
}
