package annotations

import (
	"context"
	"errors"

	"github.com/grafana/grafana/pkg/components/simplejson"
	"github.com/grafana/grafana/pkg/setting"
)

var (
	ErrTimerangeMissing = errors.New("missing timerange")
)

type Repository interface {
	Save(item *Item) error
}

type Item struct {
	Id          int64            `json:"id"`
	OrgId       int64            `json:"orgId"`
	UserId      int64            `json:"userId"`
	DashboardId int64            `json:"dashboardId"`
	PanelId     int64            `json:"panelId"`
	Text        string           `json:"text"`
	AlertId     int64            `json:"alertId"`

	// needed until we remove it from db
	Type  string
	Title string
}

func (i Item) TableName() string {
	return "annotation"
}

type ItemDTO struct {
	Id          int64            `json:"id"`
	AlertId     int64            `json:"alertId"`
}
