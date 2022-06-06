package importer

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	sf "github.com/snowflakedb/gosnowflake"
)

type Config struct {
	User      string
	Password  string
	Database  string
	Schema    string
	Account   string
	Warehouse string
}

// Run on show databases and create for each resource.
func CreateDatabases(ctx context.Context, c *Config) error {
	cfg := &sf.Config{
		User:      c.User,
		Password:  c.Password,
		Database:  c.Database,
		Schema:    c.Schema,
		Account:   c.Account,
		Warehouse: c.Warehouse,
	}
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
