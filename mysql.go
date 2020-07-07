package easydb

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	// mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func initMysql(config *DBConfig) easydb {
	return mysqlConfig(config)
}

func mysqlConfig(config *DBConfig) easydb {
	var openErr error
	mysqlInst := &Mysql{}
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
	return mysqlInst
}

type Mysql struct {
	*sql.DB
	*DBConfig
	dbType dbType
}

func (p *Mysql) SqlDB() *sql.DB {
	return p.DB
}
func (p *Mysql) Before(optType dbOptType, sqlBuilder iSQLBuilder) (string, []interface{}) {
	sql, _ := sqlBuilder.Gen()
	val := sqlBuilder.Val()
	return sql, val
}
func (p *Mysql) Do(optType dbOptType, sql string, val []interface{}) (result []map[string]interface{}, err error) {
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
		rows, queryErr := stmt.Query(val...)
		defer func() {
			rows.Close()
		}()
		if queryErr != nil {
			err = queryErr
			log.Printf("execute query error: %s sql: %s val: %v", queryErr.Error(), sql, val)
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
		r, execErr := stmt.Exec(val...)
		if execErr != nil {
			err = execErr
			log.Printf("execute error: %s sql: %s val: %v", execErr.Error(), sql, val)
			return
		}
		if id, idErr := r.LastInsertId(); idErr == nil {
			resMap["id"] = id
		}
		if ra, raErr := r.RowsAffected(); raErr == nil {
			resMap["ra"] = ra
		}
		resultArr = append(resultArr, resMap)
		break
	}
	result = resultArr
	return
}

func (p *Mysql) After(optType dbOptType, rawResult []map[string]interface{}) (result interface{}, err error) {
	return rawResult, nil
}
