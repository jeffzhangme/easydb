# easydb: DB adapter and SQL builder for golang


## Features

* Easy to use and expand
* Package management using [glide](https://github.com/Masterminds/glide)
* Current support for mysql and postgresql
* Support for Multi Data Sources
## Usage

It is recommended to use [glide](https://github.com/Masterminds/glide)


Add the configuration in the file glide.yaml: 

```
- package: github.com/jeffzhangme/easydb
```

Install the package with the command from shell:

```
$ glide install
```

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
	db := easydb.GetInst(easydb.MYSQL, mysqlConf)
	db2 := easydb.GetInst(easydb.MYSQL, mysqlConf2)
	err := db.Insert(
		easydb.BuildInsert("testdb").
			Table(easydb.Table{Name: "test_table"}).
			Values(easydb.Column{Name: "name", Value: "name"}))
	fmt.Println(err)
	result2, err2 := db2.Select(
		easydb.BuildQuery("testdb").
			Columns(easydb.Column{Name: "*"}).
			Tables(easydb.Table{Name: "test_table"}).
			Where(easydb.Where{Key: "name", Opt: easydb.EQ, Value: "name"}))
	fmt.Println(result2, err2)
}
```