package migrator

import (
	"fmt"
	"strings"

	"xorm.io/xorm"
)

var (
	ErrLockDB        = fmt.Errorf("failed to obtain lock")
	ErrReleaseLockDB = fmt.Errorf("failed to release lock")
)

type Dialect interface {
	DriverName() string
	Quote(string) string
	AndStr() string
	AutoIncrStr() string
	OrStr() string
	EqStr() string
	ShowCreateNull() bool
	SQLType(col *Column) string
	SupportEngine() bool
	LikeStr() string
}

type LockCfg struct {
	Session *xorm.Session
	Timeout int
}

type dialectFunc func(*xorm.Engine) Dialect

var supportedDialects = map[string]dialectFunc{
	MySQL:                  NewMysqlDialect,
}

func NewDialect(engine *xorm.Engine) Dialect {
	name := engine.DriverName()
	if fn, exist := supportedDialects[name]; exist {
		return fn(engine)
	}

	panic("Unsupported database type: " + name)
}

type BaseDialect struct {
	dialect    Dialect
	engine     *xorm.Engine
	driverName string
}

func (b *BaseDialect) DriverName() string {
	return b.driverName
}

