package easydb

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlInsts = map[string]*dbMysql{}
)

type dbMysql struct {
	easydb
}

func initMysql(config *dbConfig) *dbMysql {
	mu := sync.Mutex{}
	mu.Lock()
	if mysqlInsts[config.DataSource] == nil {
		mysqlConfig(config)
	} else {
		if mysqlInsts[config.DataSource].Ping() != nil {
			mysqlConfig(config)
		}
	}
	mu.Unlock()
	return mysqlInsts[config.DataSource]
}

func mysqlConfig(config *dbConfig) {
	mysqlInst := &dbMysql{}
	mysqlInst.dbType = MYSQL
	mysqlInst.dbConfig = config
	linkStr := "%s:%s@tcp(%s:%s)/%s"
	mysqlInst.DB, _ = sql.Open("mysql", fmt.Sprintf(linkStr, mysqlInst.UserName, mysqlInst.Password, mysqlInst.Host, mysqlInst.Port, mysqlInst.Schema))
	if nil != mysqlInst.Ping() {
		log.Fatal(mysqlInst.Ping())
	}
	mysqlInsts[config.DataSource] = mysqlInst
}
