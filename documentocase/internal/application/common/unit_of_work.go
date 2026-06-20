package application

import (
	"context"
	"database/sql"
)

type UnitOfWork interface {
	Do(
		ctx context.Context,
		fn func(tx *sql.Tx) error,
	) error
}
