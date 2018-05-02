package easydb

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	// mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/widuu/goini"
)

var (
	mysqlInst *dbMysql
)

const (
	defaultHost   = "localhost"
	defaultUser   = "root"
	defaultPwd    = "123456"
	defaultPort   = "3306"
	defaultDbName = "testdb"
)

type dbMysql struct {
	easydb
}

func initMysql() *dbMysql {
	mu := sync.Mutex{}
	mu.Lock()
	if mysqlInst == nil {
		mysqlConfig()
	} else {
		if mysqlInst.Ping() != nil {
			mysqlConfig()
		}
	}
	mu.Unlock()
	return mysqlInst
}

func mysqlConfig() {
	mysqlInst = &dbMysql{}
	mysqlInst.dbType = MYSQL
	conf := goini.SetConfig(getCurrentPath() + "db_config.ini")
	mysqlInst.user = conf.GetValue(db_mysql, "username")
	mysqlInst.host = conf.GetValue(db_mysql, "host")
	mysqlInst.port = conf.GetValue(db_mysql, "port")
	mysqlInst.pwd = conf.GetValue(db_mysql, "password")
	mysqlInst.defaultDbName = conf.GetValue(db_mysql, "database")
	if mysqlInst.user == no_value {
		mysqlInst.user = defaultUser
	}
	if mysqlInst.host == no_value {
		mysqlInst.host = defaultHost
	}
	if mysqlInst.port == no_value {
		mysqlInst.port = defaultPort
	}
	if mysqlInst.pwd == no_value {
		mysqlInst.pwd = defaultPwd
	}
	if mysqlInst.defaultDbName == no_value {
		mysqlInst.defaultDbName = defaultDbName
	}
	linkStr := "%s:%s@tcp(%s:%s)/%s"
	mysqlInst.DB, _ = sql.Open("mysql", fmt.Sprintf(linkStr, mysqlInst.user, mysqlInst.pwd, mysqlInst.host, mysqlInst.port, mysqlInst.defaultDbName))
	if nil != mysqlInst.Ping() {
		log.Fatal(mysqlInst.Ping())
	}
}
