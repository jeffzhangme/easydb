package easydb

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"
	"time"

	// pgsql
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

var (
	pgsqlInsts = map[string]*DBPgsql{}
)

// DBPgsql postgresql database
type DBPgsql struct {
	easydb
}

func initPgsql(config *DBConfig) *DBPgsql {
	mu := sync.Mutex{}
	mu.Lock()
	if pgsqlInsts[config.DataSource] == nil {
		pgsqlConfig(config)
	} else {
		if pgsqlInsts[config.DataSource].Ping() != nil {
			pgsqlConfig(config)
		}
	}
	mu.Unlock()
	return pgsqlInsts[config.DataSource]
}

func pgsqlConfig(config *DBConfig) {
	var openErr error
	pgsqlInst := &DBPgsql{}
	pgsqlInst.dbType = PGSQL
	pgsqlInst.DBConfig = config
	linkStr := "postgres://%s:%s@%s:%s/%s"
	if pgsqlInst.ConnParams != "" {
		linkStr += "?" + strings.Replace(pgsqlInst.ConnParams, "/", "%2F", -1)
	}
	pgsqlInst.DB, openErr = sql.Open("postgres", fmt.Sprintf(linkStr, pgsqlInst.UserName, pgsqlInst.Password, pgsqlInst.Host, pgsqlInst.Port, pgsqlInst.Schema))
	if openErr != nil {
		panic(openErr)
	}
	if err := pgsqlInst.Ping(); err != nil {
		panic(err)
	}
	if pgsqlInst.MaxOpenConns > 0 {
		pgsqlInst.DB.SetMaxOpenConns(pgsqlInst.MaxOpenConns)
	}
	if pgsqlInst.MaxIdleConns > 0 {
		pgsqlInst.DB.SetMaxIdleConns(pgsqlInst.MaxIdleConns)
	}
	if pgsqlInst.ConnMaxLifetime > 0 {
		pgsqlInst.DB.SetConnMaxLifetime(time.Duration(pgsqlInst.ConnMaxLifetime) * time.Millisecond)
	}
	if pgsqlInst.EnableMigrate {
		driver, err := postgres.WithInstance(pgsqlInst.DB, &postgres.Config{})
		if err != nil {
			panic(err)
		}
		if m, err := migrate.NewWithDatabaseInstance(
			"file://"+pgsqlInst.MigrateDir,
			"postgres", driver); err == nil {
			if err = m.Up(); err != nil {
				if err != migrate.ErrNoChange {
					panic(err)
				}
			}
		} else {
			panic(err)
		}
	}
	pgsqlInsts[config.DataSource] = pgsqlInst
}
