package models

import "errors"

// Typed errors
var (
	ErrTeamNotFound  = errors.New("team not found")
	ErrTeamNameTaken = errors.New("team name is taken")
)

// Team model
type Team struct {
	Id    int64 `json:"id"`
	OrgId int64 `json:"orgId"`
}

// --------------------
// COMMANDS

type CreateTeamCommand struct {
	Name  string `json:"name" binding:"Required"`
	Email string `json:"email"`
	OrgId int64  `json:"-"`

	Result Team `json:"-"s`
}

type UpdateTeamCommand struct {
	Id    int64
	Name  string
	Email string
	OrgId int64 `json:"-"`
}



type IsAdminOfTeamsQuery struct {
	SignedInUser *SignedInUser
	Result       bool
}
