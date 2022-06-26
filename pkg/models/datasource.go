package models

import (
	"errors"
	"time"

	"github.com/taidevops/grafana/pkg/components/simplejson"
)

const (
	DS_GRAPHITE    = "graphite"
	DS_INFLUXDB    = "influxdb"
	DS_INFLUXDB_08 = "influxdb_08"
	DS_ES          = "elasticsearch"
	DS_PROMETHEUS  = "prometheus"
)

var (
	ErrDataSourceNotFound   = errors.New("data source not found")
	ErrDataSourceNameExists = errors.New("data source with the same name already exists")
)

type DsAccess string

type DataSource struct {
	Id      int64 `json:"id"`
	OrgId   int64 `json:"orgId"`
	Version int   `json:"version"`

	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Access DsAccess `json:"access"`
	Url    string   `json:"url"`
	// swagger:ignore
	Password      string `json:"-"`
	User          string `json:"user"`
	Database      string `json:"database"`
	BasicAuth     bool   `json:"basicAuth"`
	BasicAuthUser string `json:"basicAuthUser"`
	// swagger:ignore
	BasicAuthPassword string            `json:"-"`
	WithCredentials   bool              `json:"withCredentials"`
	IsDefault         bool              `json:"isDefault"`
	JsonData          *simplejson.Json  `json:"jsonData"`
	SecureJsonData    map[string][]byte `json:"secureJsonData"`
	ReadOnly          bool              `json:"readOnly"`
	Uid               string            `json:"uid"`

	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// AllowedCookies parses the jsondata.keepCookies and returns a list of
// allowed cookies, otherwise an empty list.
func (ds DataSource) AllowedCookies() []string {
	if ds.JsonData != nil {
		if keepCookies := ds.JsonData.Get("keepCookies"); keepCookies != nil {
			return keepCookies.MustStringArray()
		}
	}

	return []string{}
}

// ----------------------
// COMMANDS

// Also acts as api DTO
type AddDataSourceCommand struct {
	Name            string            `json:"name" binding:"Required"`
	Type            string            `json:"type" binding:"Required"`
	Access          DsAccess          `json:"access" binding:"Required"`
	Url             string            `json:"url"`
	Database        string            `json:"database"`
	User            string            `json:"user"`
	BasicAuth       bool              `json:"basicAuth"`
	BasicAuthUser   string            `json:"basicAuthUser"`
	WithCredentials bool              `json:"withCredentials"`
	IsDefault       bool              `json:"isDefault"`
	JsonData        *simplejson.Json  `json:"jsonData"`
	SecureJsonData  map[string]string `json:"secureJsonData"`
	Uid             string            `json:"uid"`

	OrgId                   int64             `json:"-"`
	UserId                  int64             `json:"-"`
	ReadOnly                bool              `json:"-"`
	EncryptedSecureJsonData map[string][]byte `json:"-"`
	UpdateSecretFn          UpdateSecretFn    `json:"-"`

	Result *DataSource `json:"-"`
}

type UpdateDataSourceCommand struct {
	Name string `json:"name" binding:"Required"`
}

// Function for updating secrets along with datasources, to ensure atomicity
type UpdateSecretFn func() error

// -----------
// QUERIES

type GetDataSourcesQuery struct {
	OrgId           int64
	DataSourceLimit int
	User            *SignedInUser
	Result          []*DataSource
}
