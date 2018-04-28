package test

import (
	"fmt"
	"testing"

	"github.com/jeffzhangme/easydb"
	"github.com/jeffzhangme/easydb/mysql"
	"github.com/jeffzhangme/easydb/pgsql"
)

func Test_mysql_insert(t *testing.T) {
	defer easydb.Close()
	db := easydb.GetInst(easydb.MYSQL)
	err := db.Insert(
		mysql.BuildInsert("testdb").
			Table(easydb.Table{Name: "test_table"}).
			Values(easydb.Column{Name: "name", Value: "name"}))
	fmt.Println(err)
}
func Test_mysql_delete(t *testing.T) {
	defer easydb.Close()
	db := easydb.GetInst(easydb.MYSQL)
	err := db.Delete(
		mysql.BuildDelete("testdb").
			Table(easydb.Table{Name: "test_table"}).
			Where(easydb.Where{Key: "name", Opt: easydb.EQ, Value: "name"}))
	fmt.Println(err)
}
func Test_mysql_update(t *testing.T) {
	defer easydb.Close()
	db := easydb.GetInst(easydb.MYSQL)
	err := db.Update(
		mysql.BuildUpdate("testdb").
			Table(easydb.Table{Name: "test_table"}).
			Set(easydb.Column{Name: "name", Value: "name"}).
			Where(easydb.Where{Key: "email", Opt: easydb.EQ, Value: ""}))
	fmt.Println(err)
}
func Test_mysql_select(t *testing.T) {
	defer easydb.Close()
	db := easydb.GetInst(easydb.MYSQL)
	result, _ := db.Select(
		mysql.BuildQuery("testdb").
			Columns(easydb.Column{Name: "*"}).
			Tables(easydb.Table{Name: "test_table"}))
	fmt.Println(result)
}

func Test_pgsql_insert(t *testing.T) {
	defer easydb.Close()
	db := easydb.GetInst(easydb.PGSQL)
	err := db.Insert(
		pgsql.BuildInsert("public").
			Table(easydb.Table{Name: "test_table"}).
			Values(easydb.Column{Name: "name", Value: "name"}))
	fmt.Println(err)
}
func Test_pgsql_delete(t *testing.T) {
	defer easydb.Close()
	db := easydb.GetInst(easydb.PGSQL)
	err := db.Delete(
		pgsql.BuildDelete("public").
			Table(easydb.Table{Name: "test_table"}).
			Where(easydb.Where{Key: "name", Opt: easydb.NE, Value: "name"}))
	fmt.Println(err)
}
func Test_pgsql_update(t *testing.T) {
	defer easydb.Close()
	db := easydb.GetInst(easydb.PGSQL)
	err := db.Update(
		pgsql.BuildUpdate("public").
			Table(easydb.Table{Name: "test_table"}).
			Set(easydb.Column{Name: "email", Value: "email"}).
			Where(easydb.Where{Key: "name", Opt: easydb.EQ, Value: "name"}))
	fmt.Println(err)
}

func Test_pgsql_select(t *testing.T) {
	defer easydb.Close()
	db := easydb.GetInst(easydb.PGSQL)
	result, _ := db.Select(
		pgsql.BuildQuery("public").
			Columns(easydb.Column{Name: "*"}).
			Tables(easydb.Table{Name: "test_table"}).
			Where(easydb.Where{Key: "name", Opt: easydb.EQ, Value: "name"}))
	fmt.Println(result)
}
