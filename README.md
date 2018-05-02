# easydb: DB adapter and SQL builder for golang


## Features

* Easy to use and expand
* Package management using [glide](https://github.com/Masterminds/glide)
* Current support for mysql and postgresql

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

Add your configuration to `db_config.ini` and refer to the unit test method in the `easydb_test.go` file for use.

Your code looks like this:

```
package main

import (
	"fmt"

	"github.com/jeffzhangme/easydb"
)

func main() {
	defer easydb.Close()
	db := easydb.GetInst(easydb.MYSQL)
	result, _ := db.Select(
		easydb.BuildQuery("testdb").
			Columns(easydb.Column{Name: "*"}).
			Tables(easydb.Table{Name: "test_table"}).
			Where(easydb.Where{Key: "name", Opt: easydb.EQ, Value: "name"}))
	fmt.Println(result)
}
```