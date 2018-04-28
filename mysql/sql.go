package mysql

import (
	"bytes"
	"github.com/jeffzhangme/easydb"
)

type sql_ struct {
	// opt
	operate easydb.DBOptType
	// query type
	queryType easydb.QueryType
	// table
	tables *bytes.Buffer
	// join
	joins *bytes.Buffer
	// where
	wheres *bytes.Buffer
	// columns
	columns *bytes.Buffer
	// order by
	order string
	// limit
	limit string
	// offset
	offset string
	// group by
	group string
	// having
	havings *bytes.Buffer
}

func newSQL() *sql_ {
	sql := &sql_{}
	sql.tables = bytes.NewBufferString("")
	sql.joins = bytes.NewBufferString("")
	sql.wheres = bytes.NewBufferString("")
	sql.columns = bytes.NewBufferString("")
	sql.havings = bytes.NewBufferString("")
	sql.order = ""
	sql.limit = ""
	sql.offset = ""
	sql.group = ""
	return sql
}
