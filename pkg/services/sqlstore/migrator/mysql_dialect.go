package migrator

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/grafana/grafana/pkg/util/errutil"
	"xorm.io/xorm"
)

type MySQLDialect struct {
	BaseDialect
}

func NewMySqlDialect(engine *xorm.Engine) Dialect {
	d := MySQLDialect{}
	return &d
}

func (db *MySqlDialect) SupportEngine() bool {
	return true
}

func (db *MySqlDialect) Quote(name string) string {
	return "`" + name + "`"
}

