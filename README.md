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
  version: master
```

Install the package with the command from shell:

```
$ glide install
```

Add your configuration to `db_config.ini` and refer to the unit test method in the `easydb_test.go` file in the test directory for use.

Your code looks like this:

```
import (
  "github.com/jeffzhangme/easydb"
  "github.com/jeffzhangme/easydb/mysql"
)

func main() {
	defer easydb.Close()
	db := easydb.GetInst(easydb.MYSQL)
	result, _ := db.Select(
		mysql.BuildQuery("testdb").
			Columns(easydb.Column{Name: "*"}).
			Tables(easydb.Table{Name: "test_table"}))
	fmt.Println(result)
}
```