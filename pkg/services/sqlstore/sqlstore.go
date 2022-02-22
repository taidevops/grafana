package sqlstore

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"xorm.io/xorm"

	"github.com/grafana/grafana/pkg/services/sqlstore/migrator"
)

var (
	x	*xorm.Engine
	dialect migrator.Dialect


)

// ContextSessionKey is used as key to save values in `context.Context`
type ContextSessionKey struct{}

type SQLStore struct {
	engine *xorm.Engine
	Dialect migrator.Dialect
}

func ProvideService() (*SQLStore, error) {
	// This change will make xorm use an empty default schema for postgres and
	// by that mimic the functionality of how it was functioning before
	// xorm's changes above.
	xorm.DefaultPostgresSchema = ""
	x, err :=
}

func (ss *SQLStore) Migrate(isDatabaseLockingEnabled bool) error {

	migrator :=
}

func (ss *SQLStore) initEngine(engine *xorm.Engine) error {
	if ss.engine != nil {
		return nil
	}

	connectionString, err := ss.buildConnectionString()
	if err != nil {
		return err
	}

	if engine == nil {
		var err error
		engine, err = xorm.NewEngine("", connectionString)
		if err != nil {
			return err
		}
	}

	engine.SetMaxOpenConns(10)
	engine.SetMaxIdleConns(5)
	engine.SetConnMaxLifetime(time.Second * time.Duration(10))

	// configure sql logging

	ss.engine = engine
	return nil
}
