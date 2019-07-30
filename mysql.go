package easydb

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"
	"time"

	// mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
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
	var openErr error
	mysqlInst := &DBMysql{}
	mysqlInst.dbType = MYSQL
	mysqlInst.DBConfig = config
	linkStr := "%s:%s@tcp(%s:%s)/%s"
	linkStr = fmt.Sprintf(linkStr, mysqlInst.UserName, mysqlInst.Password, mysqlInst.Host, mysqlInst.Port, mysqlInst.Schema)
	if mysqlInst.ConnParams != "" {
		linkStr += "?" + strings.Replace(mysqlInst.ConnParams, "/", "%2F", -1)
	}
	mysqlInst.DB, openErr = sql.Open("mysql", linkStr)
	if openErr != nil {
		panic(openErr)
	}
	if err := mysqlInst.Ping(); err != nil {
		panic(err)
	}
	if mysqlInst.MaxOpenConns > 0 {
		mysqlInst.DB.SetMaxOpenConns(mysqlInst.MaxOpenConns)
	}
	if mysqlInst.MaxIdleConns > 0 {
		mysqlInst.DB.SetMaxIdleConns(mysqlInst.MaxIdleConns)
	}
	if mysqlInst.ConnMaxLifetime > 0 {
		mysqlInst.DB.SetConnMaxLifetime(time.Duration(mysqlInst.ConnMaxLifetime) * time.Millisecond)
	}
	if mysqlInst.EnableMigrate {
		driver, err := mysql.WithInstance(mysqlInst.DB, &mysql.Config{})
		if err != nil {
			panic(err)
		}
		if m, err := migrate.NewWithDatabaseInstance(
			"file://"+mysqlInst.MigrateDir,
			"mysql", driver); err == nil {
			if err = m.Up(); err != nil {
				if err != migrate.ErrNoChange {
					panic(err)
				}
			}
		} else {
			panic(err)
		}
	}
	mysqlInsts[config.DataSource] = mysqlInst
}
