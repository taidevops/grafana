package migrator

import (
	"fmt"
	"strings"

	"xorm.io/xorm"
)

const (
	Postgres = "postgres"
	SQLite   = "sqlite3"
	MySQL    = "mysql"
	MSSQL    = "mssql"
)

type Migration interface {
	SQL(dialect Dialect) string
	Id() string
	SetId(string)
	GetCondition()
}

type CodeMigration interface {
	Migration
	Exec(sess *xorm.Session, migrator *Migrator) error
}

type SQLType string

type ColumnType string

const (
	DB_TYPE_STRING ColumnType = "String"
)

type Table struct {
	Name string
	Columns []*Column
	PrimaryKeys []string
	Indices []*Index
}

var (
	DB_Bit = "BIT"
	DB_TinyInt = "TINYINT"
	DB_SmallInt = "SMALLINT"

)
