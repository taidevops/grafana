package serverlock

import (
	"context"
	"time"

	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/services/sqlstore"
)

func ProvideService(sqlStore *sqlstore.SQLStore) *ServerLockService {
	return &ServerLockService{

	}
}

// ServerLockService allows servers in HA mode to claim a lock
// and execute an function if the server was granted the lock
type ServerLockService struct {
	SQLStore *sqlstore.SQLStore
	log      log.Logger
}

func (sl *ServerLockService) LockAndExecute() error {
	// gets or creates a lockable row

}

func (sl *ServerLockService) getOrCreate(ctx context.Context) (*serverLock, error) {
	var result *serverLock

	err := sl.SQLStore.WithTransactionalDBSession(ctx, func(dbSession *sqlstore.DBSession) error {
		lockRows := []*serverLock{}
		err := dbSession.Where("")
	})
}