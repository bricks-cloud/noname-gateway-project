package importer

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

func Query(db *sql.DB, stmt string) (*sqlx.Rows, error) {
	log.Print("[DEBUG] query stmt ", stmt)
	sdb := sqlx.NewDb(db, "snowflake").Unsafe()
	return sdb.Queryx(stmt)
}
