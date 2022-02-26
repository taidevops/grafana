package sqlstore

import (
	"context"

	"github.com/grafana/grafana/pkg/bus"
)

func (ss *SQLStore) addUserQueryAndCommandHandlers() {
	ss.Bus.AddHandler(ss.GetSignedInUserWithCacheCtx)

	bus.AddHandler("sql", ss.GetUserById)
}
