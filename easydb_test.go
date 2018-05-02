package easydb

import (
	"fmt"
	"testing"
)

func Test_mysql_insert(t *testing.T) {
	defer Close()
	db := GetInst(MYSQL)
	err := db.Insert(
		BuildInsert("testdb").
			Table(Table{Name: "test_table"}).
			Values(Column{Name: "name", Value: "name"}))
	fmt.Println(err)
}
func Test_mysql_delete(t *testing.T) {
	defer Close()
	db := GetInst(MYSQL)
	err := db.Delete(
		BuildDelete("testdb").
			Table(Table{Name: "test_table"}).
			Where(Where{Key: "name", Opt: NE, Value: "name"}))
	fmt.Println(err)
}
func Test_mysql_update(t *testing.T) {
	defer Close()
	db := GetInst(MYSQL)
	err := db.Update(
		BuildUpdate("testdb").
			Table(Table{Name: "test_table"}).
			Set(Column{Name: "name", Value: "name"}).
			Where(Where{Key: "email", Opt: EQ, Value: ""}))
	fmt.Println(err)
}
func Test_mysql_select(t *testing.T) {
	defer Close()
	db := GetInst(MYSQL)
	result, _ := db.Select(
		BuildQuery("testdb").
			Columns(Column{Name: "*"}).
			Tables(Table{Name: "test_table"}).
			Where(Where{Key: "name", Opt: NE, Value: "name"}))
	fmt.Println(result)
}

func Test_pgsql_insert(t *testing.T) {
	defer Close()
	db := GetInst(PGSQL)
	err := db.Insert(
		BuildInsert("public").
			Table(Table{Name: "test_table"}).
			Values(Column{Name: "name", Value: "name"}))
	fmt.Println(err)
}
func Test_pgsql_delete(t *testing.T) {
	defer Close()
	db := GetInst(PGSQL)
	err := db.Delete(
		BuildDelete("public").
			Table(Table{Name: "test_table"}).
			Where(Where{Key: "name", Opt: NE, Value: "name"}))
	fmt.Println(err)
}
func Test_pgsql_update(t *testing.T) {
	defer Close()
	db := GetInst(PGSQL)
	err := db.Update(
		BuildUpdate("public").
			Table(Table{Name: "test_table"}).
			Set(Column{Name: "email", Value: "email"}).
			Where(Where{Key: "name", Opt: EQ, Value: "name"}))
	fmt.Println(err)
}

func Test_pgsql_select(t *testing.T) {
	defer Close()
	db := GetInst(PGSQL)
	result, _ := db.Select(
		BuildQuery("public").
			Columns(Column{Name: "*"}).
			Tables(Table{Name: "test_table"}).
			Where(Where{Key: "name", Opt: EQ, Value: "name"}))
	fmt.Println(result)
}
