package sqlstore

import "context"

type Store interface {
	GetAdminStats(ctx context.Context, query *)
}
