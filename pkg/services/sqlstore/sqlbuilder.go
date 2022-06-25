package sqlstore

import "bytes"

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
