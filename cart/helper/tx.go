package helper

import "database/sql"

func Commit0rRollback(tx *sql.Tx) {
	err := recover()

	if err != nil {
		errRollback := tx.Rollback()
		PanicIfError(errRollback)
	} else {
		errCommit := tx.Commit()

		PanicIfError(errCommit)
	}
}
