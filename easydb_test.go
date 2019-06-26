package easydb

import (
	"fmt"
	"testing"
)

var dbExec *DBExec
var mysqlConf *DBConfig
var mysqlConf2 = NewMysqlConfig(WithPassword("password"), WithSchema("testdb"), WithHost("127.0.0.1"))
var pgsqlConf = NewPgsqlConfig(WithSchema("testdb"))

func Test_mysql_insert(t *testing.T) {
	dbExec = GetInst(MYSQL, mysqlConf2)
	err := dbExec.Insert(
		BuildInsert("testdb").
			Table(Table{Name: "test_table"}).
			Values(Column{Name: "name", Value: "name"}))
	fmt.Println(err)
}
func Test_mysql_delete(t *testing.T) {
	mysqlConf = NewMysqlConfig(WithPassword("password"), WithSchema("testdb"))
	db := GetInst(MYSQL, mysqlConf2)
	err := db.Delete(
		BuildDelete("testdb").
			Table(Table{Name: "test_table"}).
			Where(Where{Key: "id", Opt: EQ, Value: 1}))
	fmt.Println(err)
}
func Test_mysql_update(t *testing.T) {
	mysqlConf = NewMysqlConfig(WithPassword("password"), WithSchema("testdb"))
	db := GetInst(MYSQL, mysqlConf2)
	err := db.Update(
		BuildUpdate("testdb").
			Table(Table{Name: "test_table"}).
			Set(Column{Name: "name", Value: "name1"}).
			Where(Where{Key: "email", Opt: EQ, Value: ""}))
	fmt.Println(err)
}
func Test_mysql_select(t *testing.T) {
	db := GetInst(MYSQL, mysqlConf2)
	result, _ := db.Select(
		BuildQuery("testdb").
			Columns(Column{Name: "*"}).
			Tables(Table{Name: "test_table"}).
			Where(Where{Key: "name", Opt: NE, Value: "name"}))
	fmt.Println(result)
}

func Test_pgsql_insert(t *testing.T) {
	db := GetInst(PGSQL, pgsqlConf)
	err := db.Insert(
		BuildInsert("public").
			Table(Table{Name: "test_table"}).
			Values(Column{Name: "name", Value: "name"}))
	fmt.Println(err)
}
func Test_pgsql_delete(t *testing.T) {
	db := GetInst(PGSQL, pgsqlConf)
	err := db.Delete(
		BuildDelete("public").
			Table(Table{Name: "test_table"}).
			Where(Where{Key: "name", Opt: NE, Value: "name"}))
	fmt.Println(err)
}
func Test_pgsql_update(t *testing.T) {
	db := GetInst(PGSQL, pgsqlConf)
	err := db.Update(
		BuildUpdate("public").
			Table(Table{Name: "test_table"}).
			Set(Column{Name: "email", Value: "email"}).
			Where(Where{Key: "name", Opt: EQ, Value: "name"}))
	fmt.Println(err)
}

func Test_pgsql_select(t *testing.T) {
	db := GetInst(PGSQL, pgsqlConf)
	result, _ := db.Select(
		BuildQuery("public").
			Columns(Column{Name: "*"}).
			Tables(Table{Name: "test_table"}).
			Where(Where{Key: "name", Opt: EQ, Value: "name"}))
	fmt.Println(result)
}
