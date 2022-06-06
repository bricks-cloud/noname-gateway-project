package importer

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	sf "github.com/snowflakedb/gosnowflake"
)

type SnowflakeGenerator struct {
}

// Run on show databases and create for each resource
func (g SnowflakeGenerator) createDatabases(ctx context.Context) error {
	cfg := &sf.Config{}
	dsn, err := sf.DSN(cfg)
	if err != nil {
		return err
	}
	db, err := sql.Open("snowflake", dsn)
	if err != nil {
		return err
	}
	dbx := sqlx.NewDb(db, "snowflake")
	dbs, err := ListDatabases(dbx)
	if err != nil {
		return err
	}
	fmt.Println(dbs)
	return nil
}
