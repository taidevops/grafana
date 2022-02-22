package sqlstore

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mattn/go-sqlite3"
	"xorm.io/xorm"

	"github.com/grafana/grafana/pkg/bus"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/util/errutil"
)

var tsclogger = log.New("sqlstore.transactions")

// WithTransactionalDbSession calls the callback with a session within a transaction.
func (ss *SQLStore) WithTransactionalDbSession(ctx context.Context, callback DBTransactionFunc) error {
	
}