package easydb

import (
	"database/sql"
	"fmt"
	"sync"

	// pgsql
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
	pgsqlInst := &DBPgsql{}
	pgsqlInst.dbType = PGSQL
	pgsqlInst.DBConfig = config
	linkStr := "postgres://%s:%s@%s:%s/%s?sslmode=require"
	pgsqlInst.DB, _ = sql.Open("postgres", fmt.Sprintf(linkStr, pgsqlInst.UserName, pgsqlInst.Password, pgsqlInst.Host, pgsqlInst.Port, pgsqlInst.Schema))
	if err := pgsqlInst.Ping(); err != nil {
		panic(err)
	}
	pgsqlInsts[config.DataSource] = pgsqlInst
}
