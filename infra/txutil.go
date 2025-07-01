package infra

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Transaction .
func Transaction(fn func(tx *sqlx.Tx) error) error {
	tx, err := SetupDB().Beginx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		}
	}()

	err = fn(tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
