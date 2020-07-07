package easydb

import (
	"database/sql"
	"errors"
	"os"
	"path/filepath"
	"sync"
)

type iSQLBuilder interface {
	Gen() (sql string, err error)
	Val() []interface{}
}

// DBExec db executor
type DBExec struct {
	db easydb
}

//Exec Exec
func (p *DBExec) Exec(sqlBuilder iSQLBuilder) ([]map[string]interface{}, error) {
	if _, ok := sqlBuilder.(interface{ __insert() }); ok {
		return p.Insert(sqlBuilder)
	}
	if _, ok := sqlBuilder.(interface{ __update() }); ok {
		return p.Update(sqlBuilder)
	}
	if _, ok := sqlBuilder.(interface{ __delete() }); ok {
		return p.Delete(sqlBuilder)
	}
	if _, ok := sqlBuilder.(interface{ __query() }); ok {
		return p.Select(sqlBuilder)
	}
	return nil, errors.New("method not support")
}

func (p *DBExec) exec(opt dbOptType, sqlBuilder iSQLBuilder) (result []map[string]interface{}, err error) {
	sql, val := p.db.Before(opt, sqlBuilder)
	if raw, e := p.db.Do(opt, sql, val); e == nil {
		if obj, e := p.db.After(opt, raw); e == nil {
			switch obj.(type) {
			case []map[string]interface{}:
				result = obj.([]map[string]interface{})
			default:
				err = errors.New("result type invalid")
			}
		} else {
			err = e
		}
	} else {
		err = e
	}
	return
}

//Insert Insert
func (p *DBExec) Insert(sqlBuilder iSQLBuilder) ([]map[string]interface{}, error) {
	return p.exec(Insert, sqlBuilder)
}

//Update Update
func (p *DBExec) Update(sqlBuilder iSQLBuilder) ([]map[string]interface{}, error) {
	return p.exec(Update, sqlBuilder)
}

//Delete Delete
func (p *DBExec) Delete(sqlBuilder iSQLBuilder) ([]map[string]interface{}, error) {
	return p.exec(Delete, sqlBuilder)
}

//Select Select
func (p *DBExec) Select(sqlBuilder iSQLBuilder) ([]map[string]interface{}, error) {
	return p.exec(Select, sqlBuilder)
}

type easydb interface {
	SqlDB() *sql.DB
	Before(optType dbOptType, sqlBuilder iSQLBuilder) (sql string, val []interface{})
	Do(optType dbOptType, sql string, val []interface{}) (result []map[string]interface{}, err error)
	After(optType dbOptType, rawResult []map[string]interface{}) (result interface{}, err error)
}

var dbExecs = make(map[string]*DBExec, 10)
var execsLock sync.Mutex

// GetExec get executor
func GetExec(dbType dbType, config *DBConfig) *DBExec {
	return GetInst(dbType, config)
}

// GetInst GetInst
func GetInst(dbType dbType, config *DBConfig) *DBExec {
	var exec *DBExec
	// set default value of config.DataSource
	if config.DataSource == "" {
		config.DataSource = config.Host + config.Port + config.Schema
	}
	execsLock.Lock()
	defer execsLock.Unlock()
	// if executor in cache
	if _, ok := dbExecs[config.DataSource]; ok {
		exec = dbExecs[config.DataSource]
	} else { // if executor not in cache
		var db easydb
		switch dbType {
		case MYSQL:
			db = initMysql(config)
			break
		case PGSQL:
			db = initPgsql(config)
			break
		}
		exec = &DBExec{db}
		dbExecs[config.DataSource] = exec
	}
	return exec
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
