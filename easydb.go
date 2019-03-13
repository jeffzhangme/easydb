package easydb

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type easydb struct {
	*sql.DB
	*DBConfig
	dbType dbType
}

var mysqlExecs, pgsqlExecs = make(map[string]*DBExec, 10), make(map[string]*DBExec, 10)

// GetInst GetInst
func GetInst(dbType dbType, config *DBConfig) *DBExec {
	var exec *DBExec
	mu := sync.Mutex{}
	mu.Lock()
	switch dbType {
	case MYSQL:
		if nil == mysqlExecs[config.DataSource] {
			mysqlExecs[config.DataSource] = &DBExec{initMysql(config)}
		}
		exec = mysqlExecs[config.DataSource]
		break
	case PGSQL:
		if nil == pgsqlExecs[config.DataSource] {
			pgsqlExecs[config.DataSource] = &DBExec{initPgsql(config)}
		}
		exec = pgsqlExecs[config.DataSource]
		break
	}
	mu.Unlock()
	return exec
}

// Do Do
func (p *easydb) Do(optType dbOptType, sqlBuilder iSQLBuilder) (result []map[string]interface{}, err error) {
	sql, _ := sqlBuilder.Gen()
	switch p.dbType {
	case PGSQL:
		placeholder := make([]interface{}, len(sqlBuilder.Val()))
		for i := range sqlBuilder.Val() {
			placeholder[i] = "$" + strconv.Itoa(i+1)
		}
		sql = strings.Replace(sql, "?", "%s", -1)
		if len(placeholder) > 0 {
			sql = fmt.Sprintf(sql, placeholder...)
		}
		break
	}
	stmt, err := p.Prepare(sql)
	switch optType {
	case Select:
		rows, queryErr := stmt.Query(convertToInterfaceSlice(sqlBuilder.Val())...)
		err = queryErr
		columns, _ := rows.Columns()
		dest := make([]interface{}, len(columns))
		destPointers := make([]interface{}, len(columns))
		for i := range columns {
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

// Close Close
func Close() {
	for _, inst := range mysqlInsts {
		inst.Close()
	}
	for _, inst := range pgsqlInsts {
		inst.Close()
	}
	mysqlExecs, pgsqlExecs = nil, nil
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
