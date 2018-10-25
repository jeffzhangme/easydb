package easydb

import (
	"fmt"
	"testing"
)

var mysqlConf = NewMysqlConfig(WithPassword("password"), WithSchema("testdb"))
var mysqlConf2 = NewMysqlConfig(WithPassword("password"), WithSchema("testdb"), WithHost("127.0.0.1"))
var pgsqlConf = NewPgsqlConfig(WithSchema("testdb"))

func Test_mysql_insert(t *testing.T) {
	db := GetInst(MYSQL, mysqlConf2)
	err := db.Insert(
		BuildInsert("testdb").
			Table(Table{Name: "test_table"}).
			Values(Column{Name: "name", Value: "name"}))
	fmt.Println(err)
}
func Test_mysql_delete(t *testing.T) {
	db := GetInst(MYSQL, mysqlConf)
	err := db.Delete(
		BuildDelete("testdb").
			Table(Table{Name: "test_table"}).
			Where(Where{Key: "name", Opt: NE, Value: "name"}))
	fmt.Println(err)
}
func Test_mysql_update(t *testing.T) {
	db := GetInst(MYSQL, mysqlConf)
	err := db.Update(
		BuildUpdate("testdb").
			Table(Table{Name: "test_table"}).
			Set(Column{Name: "name", Value: "name"}).
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
