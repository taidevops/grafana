package migrations

import (
	"fmt"

	. "github.com/grafana/grafana/pkg/services/sqlstore/migrator"
	"github.com/grafana/grafana/pkg/util"
	"xorm.io/xorm"
)

func addUserMigrations(mg *Migrator) {
	userV1 := Table{
		Name: "user",
		Columns: []*Column{
			{Name: "id", Type: DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "version", Type: DB_Int, Nullable: false},
		},
		Indices: []*Index{
			{Cols: []string{"login"}, Type: UniqueIndex},
			{Cols: []string{"email"}, Type: UniqueIndex},
		}
	}
}
