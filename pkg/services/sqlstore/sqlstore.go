package sqlstore

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

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
	Cfg *setting.Cfg
	Bus bus.Bus
	CacheService *localcache.CacheService

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

func newSQLStore(cfg *setting.Cfg, cacheService *localcache.CacheService, b bus.Bus, engine *xorm.Engine,
	migrations registry.DatabaseMigrator, tracer tracing.Tracer, opts ...InitTestDBOpt) (*SQLStore, error) {
	ss := &SQLStore{
		Cfg:                         cfg,
		Bus:                         b,
		CacheService:                cacheService,
		log:                         log.New("sqlstore"),
		skipEnsureDefaultOrgAndUser: false,
		migrations:                  migrations,
		tracer:                      tracer,
	}

	if err := ss.initEngine(engine); err != nil {
		return nil, errutil.Wrap("failed to connect to database", err)
	}

	ss.Dialect = migrator.NewDialect(ss.engine)

	// temporarily still set global var
	x = ss.engine
	dialect = ss.Dialect

	// Init repo instances

	return ss, nil
}

func (ss *SQLStore) Migrate(isDatabaseLockingEnabled bool) error {
	if ss.dbCfg.SkipMigrations {
		return nil
	}

	migrator := migrator.NewMigrator(ss.engine, ss.Cfg)
	ss.migrations.AddMigration(migrator)

	return migrator.Start()
}

// Sync syncs changes to the database.
func (ss *SQLStore) Sync() error {
	return ss.engine.Sync2()
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

type DatabaseConfig struct {
	Type                        string
	Host                        string
	Name                        string
	User                        string
	Pwd                         string
	Path                        string
	SslMode                     string
	CaCertPath                  string
	ClientKeyPath               string
	ClientCertPath              string
	ServerCertName              string
	ConnectionString            string
	IsolationLevel              string
	MaxOpenConn                 int
	MaxIdleConn                 int
	ConnMaxLifetime             int
	CacheMode                   string
	UrlQueryParams              map[string][]string
	SkipMigrations              bool
	MigrationLockAttemptTimeout int
}
