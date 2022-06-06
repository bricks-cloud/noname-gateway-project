package importer

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type role struct {
	Name    sql.NullString `db:"name"`
	Comment sql.NullString `db:"comment"`
}

func ScanRole(row *sqlx.Row) (*role, error) {
	r := &role{}
	err := row.StructScan(r)
	return r, err
}
