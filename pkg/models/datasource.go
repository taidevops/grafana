package models

import (
	"errors"

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

	JsonData *simplejson.Json `json:"jsonData"`
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

type AddDataSourceComamnd struct {
	Name string `json:"name" binding:"Required"`
	Type string `json:"type" binding:"Required"`
}

type UpdateDataSourceCommand struct {
	Name string `json:"name" binding:"Required"`
}

// -----------
// QUERIES

type GetDataSourcesQuery struct {
	OrgId int64
	DataSourceLimit int
	User *
}
