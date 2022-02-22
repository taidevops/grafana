package models

import (
	"errors"
	"time"
)

// Typed errors
var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrLastGrafanaAdmin  = errors.New("cannot remove last grafana admin")
	ErrProtectedUser     = errors.New("cannot adopt protected user")
)

type Password string

func (p *Password) IsWeak() bool {
	return len(p) < 4
}

type User struct {
	Id            int64
	Version       int
	Email         string
	Name          string
}

// ---------------------
// COMMANDS

type CreateUserCommand struct {
	Email            string
	Login            string

	Result User
}

