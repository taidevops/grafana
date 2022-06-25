package models

import (
	"errors"
	"time"
)

// Typed errors
var (
	ErrCaseInsensitive   = errors.New("case insensitive conflict")
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrLastGrafanaAdmin  = errors.New("cannot remove last grafana admin")
	ErrProtectedUser     = errors.New("cannot adopt protected user")
)

type Password string

func (p Password) IsWeak() bool {
	return len(p) < 4
}

type User struct {
	Id            int64
	Version       int
	Email         string
	Name          string
	Login         string
	Password      string
	Salt          string
	Rands         string
	Company       string
	EmailVerified bool
	Theme         string
	HelpFlags1    HelpFlags1
	IsDisabled    bool

	IsAdmin          bool
	IsServiceAccount bool
	OrgId            int64

	Created    time.Time
	Updated    time.Time
	LastSeenAt time.Time
}

// ---------------------
// COMMANDS

type CreateUserCommand struct {
	Email            string
	Login            string
	Name             string
	Company          string
	OrgId            int64
	OrgName          string
	Password         string
	EmailVerified    bool
	IsAdmin          bool
	IsDisabled       bool
	SkipOrgSetup     bool
	DefaultOrgRole   string
	IsServiceAccount bool

	Result User
}

// ------------------------
// DTO & Projections

type SignedInUser struct {
	UserId int64
	OrgId int64
	OrgName string
}
