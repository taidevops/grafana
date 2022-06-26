package models

import "errors"

// Typed errors
var (
	ErrTeamMemberAlreadyAdded = errors.New("User is already added to this team")
)

// TeamMember model
type TeamMember struct {
	Id int64
	OrgId int64
	TeamId int64
	UserId int64
	External bool
	Permission Per
}
