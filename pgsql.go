package easydb

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	// pgsql
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

func initPgsql(config *DBConfig) easydb {
	return pgsqlConfig(config)
}

func pgsqlConfig(config *DBConfig) easydb {
	var openErr error
	pgsqlInst := &Pgsql{}
	pgsqlInst.dbType = PGSQL
	pgsqlInst.DBConfig = config
	linkStr := "postgres://%s:%s@%s:%s/%s"
	if pgsqlInst.ConnParams != "" {
		linkStr += "?" + strings.Replace(pgsqlInst.ConnParams, "/", "%2F", -1)
	}
	pgsqlInst.DB, openErr = sql.Open("postgres", fmt.Sprintf(linkStr, pgsqlInst.UserName, pgsqlInst.Password, pgsqlInst.Host, pgsqlInst.Port, pgsqlInst.Schema))
	if openErr != nil {
		panic(openErr)
	}
	if err := pgsqlInst.Ping(); err != nil {
		panic(err)
	}
	if pgsqlInst.MaxOpenConns > 0 {
		pgsqlInst.DB.SetMaxOpenConns(pgsqlInst.MaxOpenConns)
	}
	if pgsqlInst.MaxIdleConns > 0 {
		pgsqlInst.DB.SetMaxIdleConns(pgsqlInst.MaxIdleConns)
	}
	if pgsqlInst.ConnMaxLifetime > 0 {
		pgsqlInst.DB.SetConnMaxLifetime(time.Duration(pgsqlInst.ConnMaxLifetime) * time.Millisecond)
	}
	if pgsqlInst.EnableMigrate {
		driver, err := postgres.WithInstance(pgsqlInst.DB, &postgres.Config{})
		if err != nil {
			panic(err)
		}
		if m, err := migrate.NewWithDatabaseInstance(
			"file://"+pgsqlInst.MigrateDir,
			"postgres", driver); err == nil {
			if err = m.Up(); err != nil {
				if err != migrate.ErrNoChange {
					panic(err)
				}
			}
		} else {
			panic(err)
		}
	}
	return pgsqlInst
}

type Pgsql struct {
	Mysql
}

func (p *Pgsql) SqlDB() *sql.DB {
	return p.SqlDB()
}
func (p *Pgsql) Before(optType dbOptType, sqlBuilder iSQLBuilder) (string, []interface{}) {
	sql, val := p.Mysql.Before(optType, sqlBuilder)
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
	return sql, val
}
func (p *Pgsql) Do(optType dbOptType, sql string, val []interface{}) (result []map[string]interface{}, err error) {
	switch optType {
	case Insert:
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
		resMap := map[string]interface{}{}

		var lastInsertID int
		execErr := stmt.QueryRow(val...).Scan(&lastInsertID)
		if execErr != nil {
			err = execErr
			log.Printf("execute error: %s sql: %s val: %v", execErr.Error(), sql, val)
			return
		}
		resMap["id"] = lastInsertID
		resultArr = append(resultArr, resMap)
		result = resultArr
	default:
		result, err = p.Mysql.Do(optType, sql, val)
	}
	return
}

func (p *Pgsql) After(optType dbOptType, rawResult []map[string]interface{}) (result interface{}, err error) {
	return p.Mysql.After(optType, rawResult)
}
