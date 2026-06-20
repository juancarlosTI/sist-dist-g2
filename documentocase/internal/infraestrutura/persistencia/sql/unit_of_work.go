package sql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
)

type SQLUnitOfWork struct {
	db         *sql.DB
	maxRetries int
}

func NewSQLUnitOfWork(db *sql.DB) *SQLUnitOfWork {
	return &SQLUnitOfWork{
		db:         db,
		maxRetries: 3,
	}
}

func IsRetryable(err error) bool {

	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {

		switch pgErr.Code {

		// serialization_failure
		case "40001":

			return true

		// deadlock_detected
		case "40P01":

			return true

		// lock_not_available
		case "55P03":

			return true
		}
	}

	return false
}

func (u *SQLUnitOfWork) Do(
	ctx context.Context,
	fn func(tx *sql.Tx) error,
) error {

	var lastErr error

	for attempt := 1; attempt <= u.maxRetries; attempt++ {

		tx, err := u.db.BeginTx(ctx, nil)
		if err != nil {
			return fmt.Errorf(
				"erro iniciando transação: %w",
				err,
			)
		}

		committed := false

		defer func() {
			if !committed {
				_ = tx.Rollback()
			}
		}()

		err = fn(tx)

		if err != nil {

			lastErr = err

			if IsRetryable(err) {
				continue
			}

			return err
		}

		err = tx.Commit()

		if err != nil {

			lastErr = err

			if IsRetryable(err) {
				continue
			}

			return fmt.Errorf(
				"erro realizando commit: %w",
				err,
			)
		}

		committed = true

		return nil
	}

	return fmt.Errorf(
		"transação falhou após %d tentativas: %w",
		u.maxRetries,
		lastErr,
	)
}
