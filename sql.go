package easydb

import (
	"bytes"
)

type sql_ struct {
	// opt
	operate dbOptType
	// query type
	queryType queryType
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
