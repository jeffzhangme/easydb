package easydb

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	// pgsql
	_ "github.com/lib/pq"
	"github.com/widuu/goini"
)

var (
	pgsqlInst *dbPgsql
)

const (
	defaultPgHost   = "localhost"
	defaultPgUser   = "postgres"
	defaultPgPwd    = "postgres"
	defaultPgPort   = "5432"
	defaultPgDbName = "testdb"
)

type dbPgsql struct {
	easydb
}

func initPgsql() *dbPgsql {
	mu := sync.Mutex{}
	mu.Lock()
	if pgsqlInst == nil {
		pgsqlConfig()
	} else {
		if pgsqlInst.Ping() != nil {
			pgsqlConfig()
		}
	}
	mu.Unlock()
	return pgsqlInst
}

func pgsqlConfig() {
	pgsqlInst = &dbPgsql{}
	pgsqlInst.dbType = PGSQL
	conf := goini.SetConfig(getCurrentPath() + "db_config.ini")
	pgsqlInst.user = conf.GetValue(db_pgsql, "username")
	pgsqlInst.host = conf.GetValue(db_pgsql, "host")
	pgsqlInst.port = conf.GetValue(db_pgsql, "port")
	pgsqlInst.pwd = conf.GetValue(db_pgsql, "password")
	pgsqlInst.defaultDbName = conf.GetValue(db_pgsql, "database")
	if pgsqlInst.user == no_value {
		pgsqlInst.user = defaultPgUser
	}
	if pgsqlInst.host == no_value {
		pgsqlInst.host = defaultPgHost
	}
	if pgsqlInst.port == no_value {
		pgsqlInst.port = defaultPgPort
	}
	if pgsqlInst.pwd == no_value {
		pgsqlInst.pwd = defaultPgPwd
	}
	if pgsqlInst.defaultDbName == no_value {
		pgsqlInst.defaultDbName = defaultPgDbName
	}
	linkStr := "postgres://%s:%s@%s:%s/%s?sslmode=require"
	pgsqlInst.DB, _ = sql.Open("postgres", fmt.Sprintf(linkStr, pgsqlInst.user, pgsqlInst.pwd, pgsqlInst.host, pgsqlInst.port, pgsqlInst.defaultDbName))
	if nil != pgsqlInst.Ping() {
		log.Fatal(pgsqlInst.Ping())
	}
}
