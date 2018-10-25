package easydb

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	// pgsql
	_ "github.com/lib/pq"
)

var (
	pgsqlInsts = map[string]*dbPgsql{}
)

type dbPgsql struct {
	easydb
}

func initPgsql(config *dbConfig) *dbPgsql {
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

func pgsqlConfig(config *dbConfig) {
	pgsqlInst := &dbPgsql{}
	pgsqlInst.dbType = PGSQL
	pgsqlInst.dbConfig = config
	linkStr := "postgres://%s:%s@%s:%s/%s?sslmode=require"
	pgsqlInst.DB, _ = sql.Open("postgres", fmt.Sprintf(linkStr, pgsqlInst.UserName, pgsqlInst.Password, pgsqlInst.Host, pgsqlInst.Port, pgsqlInst.Schema))
	if nil != pgsqlInst.Ping() {
		log.Fatal(pgsqlInst.Ping())
	}
	pgsqlInsts[config.DataSource] = pgsqlInst
}
