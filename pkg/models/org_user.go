package models

import "errors"

// Typed errors
var (
	ErrLastOrgAdmin        = errors.New("cannot remove last organization admin")
	ErrOrgUserNotFound     = errors.New("cannot find the organization user")
	ErrOrgUserAlreadyAdded = errors.New("user is already added to organization")
)

// swagger:enum RoleType
type RoleType string
