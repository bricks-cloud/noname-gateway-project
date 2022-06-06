package importer

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type user struct {
	Comment          sql.NullString `db:"comment"`
	DefaultNamespace sql.NullString `db:"default_namespace"`
	DefaultRole      sql.NullString `db:"default_role"`
	DefaultWarehouse sql.NullString `db:"default_warehouse"`
	Disabled         bool           `db:"disabled"`
	DisplayName      sql.NullString `db:"display_name"`
	Email            sql.NullString `db:"email"`
	FirstName        sql.NullString `db:"first_name"`
	HasRsaPublicKey  bool           `db:"has_rsa_public_key"`
	LastName         sql.NullString `db:"last_name"`
	LoginName        sql.NullString `db:"login_name"`
	Name             sql.NullString `db:"name"`
}

func ListUsers(pattern string, db *sql.DB) ([]user, error) {
	stmt := fmt.Sprintf(`SHOW USERS like '%s'`, pattern)
	rows, err := Query(db, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dbs := []user{}
	err = sqlx.StructScan(rows, &dbs)
	if err == sql.ErrNoRows {
		log.Printf("[DEBUG] no users found")
		return nil, nil
	}
	return dbs, errors.Wrapf(err, "unable to scan row for %s", stmt)
}
