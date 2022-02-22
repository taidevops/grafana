package sqlstore

import (
	"bytes"
	"strings"

	"github.com/grafana/grafana/pkg/models"
)

type SQLBuilder struct {
	sql    bytes.Buffer
	params []interface{}
}

func (sb *SQLBuilder) Write(sql string, params ...interface{}) {
	sb.sql.WriteString(sql)

	if len(params) > 0 {
		sb.params = append(sb.params, params...)
	}
}

func (sb *SQLBuilder) GetSQLString() string {
	return sb.sql.String()
}

func (sb *SQLBuilder) GetParams() []interface{} {
	return sb.params
}

func (sb *SQLBuilder) AddParams(params ...interface{}) {
	sb.params = append(sb.params, params...)
}

func (sb *SQLBuilder) WriteDashboardPermissionFilter() {
	sb.sql.WriteString(` AND
	(
		dashboard.id IN (
			SELECT from (
				SELECT

				UNION
				SELECT
			) AS a
		)
	)`)
}
