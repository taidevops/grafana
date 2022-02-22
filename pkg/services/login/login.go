package login

import (
	"context"
	"errors"

	"github.com/grafana/grafana/pkg/models"
)

var (
	ErrInvalidCredentials = errors.New("Invalid user or password")
)

type TeamSyncFunc func(user *models.User, externalUser *models.ExternalUserInfo) error

type Service interface {
	CreateUser(cmd models.CreateUserCommand) (*models.User, error)
	UpsertUser(ctx context.Context, cmd *models.UpsertUserCommand) error
	SetTeamSyncFunc(TeamSyncFunc)
}
