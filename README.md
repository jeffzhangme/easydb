# easydb: DB adapter and SQL builder for golang


## Features

* Easy to use and expand
* Package management using [go modules](https://blog.golang.org/using-go-modules)
* Current support for mysql and postgresql
* Support for Multi Data Sources
## Usage

It is recommended to use [go module](https://blog.golang.org/using-go-modules)

Your code looks like this:

```
package main

import (
	"fmt"

	"github.com/jeffzhangme/easydb"
)

var mysqlConf = easydb.NewMysqlConfig(easydb.WithPassword("password"), easydb.WithSchema("testdb"))
var mysqlConf2 = easydb.NewMysqlConfig(easydb.WithHost("127.0.0.1"), easydb.WithPassword("password"), easydb.WithSchema("testdb"))

func main() {
	defer easydb.Close()
	db := easydb.GetExec(easydb.MYSQL, mysqlConf)
	db2 := easydb.GetExec(easydb.MYSQL, mysqlConf2)
	r, _ := db.Exec(
		easydb.BuildInsert("testdb").
			Table(easydb.Table{Name: "test_table"}).
			Values(easydb.Column{Name: "name", Value: "name"}))
	fmt.Println(r)
	r2, _ := db2.Exec(
		easydb.BuildQuery("testdb").
			Columns(easydb.Column{Name: "*"}).
			Tables(easydb.Table{Name: "test_table"}).
			Where(easydb.Where{Key: "name", Opt: easydb.EQ, Value: "name"}))
	fmt.Println(r2)
}
```