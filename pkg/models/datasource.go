package models

import (
	"errors"
	"time"

	"github.com/grafana/grafana/pkg/components/simplejson"
)

const (
	DS_GRAPHITE = "graphite"
)

var (
	ErrDataSourceNotFound = errors.New("data source not found")
)

type DsAccess string

type DataSource struct {
	Id int64 `json:"id"`
}

