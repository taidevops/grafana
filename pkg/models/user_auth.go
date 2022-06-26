package models

import (
	"time"

	"golang.org/x/oauth2"
)

const (
	AuthModuleLDAP = "ldap"
)

type UserAuth struct {
	Id                int64
	UserId            int64
	AuthModule        string
	AuthId            string
	Created           time.Time
	OAuthAccessToken  string
	OAuthRefreshToken string
	OAuthIdToken      string
	OAuthTokenType    string
	OAuthExpiry       time.Time
}

type ExternalUserInfo struct {
	OAuthToken     *oauth2.Token
	AuthModule     string
	AuthId         string
	UserId         int64
	Email          string
	Login          string
	Name           string
	Groups         []string
	OrgRoles       map[int64]RoleType
	IsGrafanaAdmin *bool // This is a pointer to know if we should sync this or not (nil = ignore sync)
	IsDisabled     bool
}

type LoginInfo struct {
	AuthModule    string
	User          *User
	ExternalUser  ExternalUserInfo
	LoginUsername string
	HTTPStatus    int
	Error         error
}

// RequestURIKey is used as key to save request URI in contexts
// (used for the Enterprise auditing feature)
type RequestURIKey struct{}
