package models

import "errors"

var (
	ErrApiKeyNotFound          = errors.New("API key not found")
	ErrInvalidApiKey           = errors.New("invalid API key")
	ErrInvalidApiKeyExpiration = errors.New("negative value for SecondsToLive")
	ErrDuplicateApiKey         = errors.New("API key, organization ID and name must be unique")
)

type ApiKey struct {
	Id int64
	OrgId int64
	Name string
	Key string
	Role 
}
