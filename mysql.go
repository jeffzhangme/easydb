package easydb

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
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
	*sql.DB
	host, user, pwd, port, defaultDbName string
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
}

// Do Do
func (p *dbMysql) Do(optType DBOptType, sqlBuilder iSQLBuilder) (result []map[string]interface{}, err error) {
	sql, _ := sqlBuilder.Gen()
	stmt, err := p.Prepare(sql)
	switch optType {
	case Select:
		rows, queryErr := stmt.Query(convertToInterfaceSlice(sqlBuilder.Val())...)
		err = queryErr
		columns, _ := rows.Columns()
		dest := make([]interface{}, len(columns))
		destPointers := make([]interface{}, len(columns))
		for i, _ := range columns {
			destPointers[i] = &dest[i]
		}
		resultArr := []map[string]interface{}{}
		for rows.Next() {
			err = rows.Scan(destPointers...)
			resultMap := map[string]interface{}{}
			for i, val := range dest {
				resultMap[columns[i]] = val
				if v, ok := (val).([]byte); ok {
					resultMap[columns[i]] = string(v)
				}
			}
			resultArr = append(resultArr, resultMap)
		}
		result = resultArr
		break
	default:
		_, execErr := stmt.Exec(convertToInterfaceSlice(sqlBuilder.Val())...)
		err = execErr
		break
	}
	return
}

// convertToInterfaceSlice []string to []interface{}
func convertToInterfaceSlice(strSlice []string) []interface{} {
	interSlice := make([]interface{}, len(strSlice))
	for index, value := range strSlice {
		interSlice[index] = value
	}
	return interSlice
}

func getCurrentPath() string {
	currentPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	currentPath += "/"
	return currentPath
}
