package sqlite

import (
	"database/sql"
)

func MigrateDB(db *sql.DB) error {
	q := `CREATE TABLE employee (name text, supervisor text)`
	_, err := db.Exec(q)
	return err
}
