package easydb

import (
	"database/sql"
	"fmt"
	"sync"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlInsts = map[string]*DBMysql{}
)

// DBMysql mysql database
type DBMysql struct {
	easydb
}

func initMysql(config *DBConfig) *DBMysql {
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

func mysqlConfig(config *DBConfig) {
	mysqlInst := &DBMysql{}
	mysqlInst.dbType = MYSQL
	mysqlInst.DBConfig = config
	linkStr := "%s:%s@tcp(%s:%s)/%s"
	mysqlInst.DB, _ = sql.Open("mysql", fmt.Sprintf(linkStr, mysqlInst.UserName, mysqlInst.Password, mysqlInst.Host, mysqlInst.Port, mysqlInst.Schema))
	if err := mysqlInst.Ping(); err != nil {
		panic(err)
	}
	mysqlInsts[config.DataSource] = mysqlInst
}
