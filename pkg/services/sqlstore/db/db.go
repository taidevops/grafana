package db

import "context"

type DB interface {
	WithTransactionDbSession(ctx context.Context, callback )
}
