package easydb

import (
	"database/sql"
	"fmt"
	"log"
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

var dbExecs = make(map[string]*DBExec, 10)

// GetExec get executor
func GetExec(dbType dbType, config *DBConfig) *DBExec {
	return GetInst(dbType, config)
}

// GetInst GetInst
func GetInst(dbType dbType, config *DBConfig) *DBExec {
	var exec *DBExec
	mu := sync.Mutex{}
	// set default value of config.DataSource
	if config.DataSource == "" {
		config.DataSource = config.Host + config.Port + config.Schema
	}
	mu.Lock()
	// if executor in cache
	if _, ok := dbExecs[config.DataSource]; ok {
		exec = dbExecs[config.DataSource]
	} else { // if executor not in cache
		var dbOperator iDBOperate
		switch dbType {
		case MYSQL:
			dbOperator = initMysql(config)
			break
		case PGSQL:
			dbOperator = initPgsql(config)
			break
		}
		exec = &DBExec{dbOperator}
		dbExecs[config.DataSource] = exec
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
		if optType == Insert {
			sql += " RETURNING id"
		}
		break
	}
	stmt, sErr := p.Prepare(sql)
	if sErr != nil {
		err = sErr
		log.Printf("create statement error: %s", sErr.Error())
		return
	}
	defer func() {
		stmt.Close()
	}()
	resultArr := []map[string]interface{}{}
	switch optType {
	case Select:
		rows, queryErr := stmt.Query(sqlBuilder.Val()...)
		defer func() {
			rows.Close()
		}()
		if queryErr != nil {
			err = queryErr
			log.Printf("execute query error: %s sql: %s val: %v", queryErr.Error(), sql, sqlBuilder.Val())
			return
		}
		columns, _ := rows.Columns()
		dest := make([]interface{}, len(columns))
		destPointers := make([]interface{}, len(columns))
		for i := range columns {
			destPointers[i] = &dest[i]
		}
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
		break
	default:
		resMap := map[string]interface{}{}
		if p.dbType == PGSQL && optType == Insert {
			var lastInsertID int
			execErr := stmt.QueryRow(sqlBuilder.Val()...).Scan(&lastInsertID)
			if execErr != nil {
				err = execErr
				log.Printf("execute error: %s sql: %s val: %v", execErr.Error(), sql, sqlBuilder.Val())
				return
			}
			resMap["id"] = lastInsertID
		} else {
			r, execErr := stmt.Exec(sqlBuilder.Val()...)
			if execErr != nil {
				err = execErr
				log.Printf("execute error: %s sql: %s val: %v", execErr.Error(), sql, sqlBuilder.Val())
				return
			}
			if id, idErr := r.LastInsertId(); idErr == nil {
				resMap["id"] = id
			}
			if ra, raErr := r.RowsAffected(); raErr == nil {
				resMap["ra"] = ra
			}
		}
		resultArr = append(resultArr, resMap)
		break
	}
	result = resultArr
	return
}

// Close Close
func Close() {
	for key := range dbExecs {
		delete(dbExecs, key)
	}
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