package easydb

import (
	"fmt"
	"testing"
	// . "github.com/jeffzhangme/easydb"
)

var dbExec *DBExec
var mysqlConf *DBConfig
var mysqlConf2 = NewMysqlConfig(WithPassword("password"), WithSchema("testdb"), WithHost("127.0.0.1"))
var pgsqlConf = NewPgsqlConfig(WithSchema("testdb"))

func Test_mysql_insert(t *testing.T) {

	dbExec = GetInst(MYSQL, mysqlConf2)
	r, _ := dbExec.Insert(
		BuildInsert("testdb").
			Table(Table{Name: "test_table"}).
			Values(Column{Name: "name", Value: "name"}))
	fmt.Println(r)
}
func Test_mysql_delete(t *testing.T) {
	mysqlConf = NewMysqlConfig(WithPassword("password"), WithSchema("testdb"))
	db := GetInst(MYSQL, mysqlConf2)
	r, _ := db.Delete(
		BuildDelete("testdb").
			Table(Table{Name: "test_table"}).
			Where(Where{Key: "id", Opt: EQ, Value: 1}))
	fmt.Println(r)
}
func Test_mysql_update(t *testing.T) {
	mysqlConf = NewMysqlConfig(WithPassword("password"), WithSchema("testdb"))
	db := GetInst(MYSQL, mysqlConf2)
	r, _ := db.Update(
		BuildUpdate("testdb").
			Table(Table{Name: "test_table"}).
			Set(Column{Name: "name", Value: "name1"}).
			Where(Where{Key: "email", Opt: EQ, Value: ""}))
	fmt.Println(r)
}
func Test_mysql_select(t *testing.T) {
	db := GetInst(MYSQL, mysqlConf2)
	r, _ := db.Select(
		BuildQuery("testdb").
			Columns(Column{Name: "*"}).
			Tables(Table{Name: "test_table"}).
			Where(Where{Key: "name", Opt: NE, Value: "name"}))
	fmt.Println(r)
}

func Test_pgsql_insert(t *testing.T) {
	db := GetInst(PGSQL, pgsqlConf)
	r, _ := db.Insert(
		BuildInsert("public").
			Table(Table{Name: "test_table"}).
			Values(Column{Name: "name", Value: "name"}))
	fmt.Println(r)
}
func Test_pgsql_delete(t *testing.T) {
	db := GetInst(PGSQL, pgsqlConf)
	r, _ := db.Delete(
		BuildDelete("public").
			Table(Table{Name: "test_table"}).
			Where(Where{Key: "name", Opt: NE, Value: "name"}))
	fmt.Println(r)
}
func Test_pgsql_update(t *testing.T) {
	db := GetInst(PGSQL, pgsqlConf)
	r, _ := db.Update(
		BuildUpdate("public").
			Table(Table{Name: "test_table"}).
			Set(Column{Name: "email", Value: "email"}).
			Where(Where{Key: "name", Opt: EQ, Value: "name"}))
	fmt.Println(r)
}

func Test_pgsql_select(t *testing.T) {
	db := GetInst(PGSQL, pgsqlConf)
	r, _ := db.Select(
		BuildQuery("public").
			Columns(Column{Name: "*"}).
			Tables(Table{Name: "test_table"}).
			Where(Where{Key: "name", Opt: EQ, Value: "name"}))
	fmt.Println(r)
}
