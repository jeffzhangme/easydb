package mysql

import (
	"bytes"
	"github.com/jeffzhangme/easydb"
	"strings"
)

type sqlBuilder struct {
	sql *sql_
}

func newSQLBuilder() *sqlBuilder {
	sqlBuilder := &sqlBuilder{sql: newSQL()}
	return sqlBuilder
}

func build(optType easydb.DBOptType) sqlBuilder {
	sqlBuilder := sqlBuilder{sql: newSQL()}
	sqlBuilder.operate(optType)
	return sqlBuilder
}

func (p *sqlBuilder) queryType(queryType easydb.QueryType) {
	p.sql.queryType = queryType
}

func (p *sqlBuilder) operate(operate easydb.DBOptType) {
	p.sql.operate = operate
}

func (p *sqlBuilder) columns(column string) {
	if 0 == p.sql.columns.Len() {
		p.sql.columns.WriteString(column)
	} else {
		_, err := p.sql.columns.WriteString(", " + column)
		if nil != err {
		}
	}
}

func (p *sqlBuilder) tables(table string) {
	if 0 == p.sql.tables.Len() {
		p.sql.tables.WriteString(table)
	} else {
		_, err := p.sql.tables.WriteString(", " + table)
		if nil != err {
		}
	}
}

func (p *sqlBuilder) join(joinStr string) {
	p.sql.joins.WriteString(joinStr)
}

func (p *sqlBuilder) on(on string) {
	if !strings.HasSuffix(p.sql.joins.String(), " ) ") {
		p.sql.joins.WriteString(" ON ( ")
	} else {
		p.sql.joins.Truncate(p.sql.joins.Len() - 3)
	}
	_, err := p.sql.joins.WriteString(on + " ) ")
	if nil != err {

	}
}

func (p *sqlBuilder) wheres(where string) {
	if 0 == p.sql.wheres.Len() {
		p.sql.wheres.WriteString(" WHERE ( ")
	} else {
		p.sql.wheres.Truncate(p.sql.wheres.Len() - 3)
	}
	_, err := p.sql.wheres.WriteString(where + " ) ")
	if nil != err {

	}
}

func (p *sqlBuilder) order(order string) {
	p.sql.order = " ORDER BY " + order
}

func (p *sqlBuilder) limit(limit string) {
	p.sql.limit = " LIMIT " + limit
}

func (p *sqlBuilder) offset(offset string) {
	p.sql.offset = " OFFSET " + offset
}

func (p *sqlBuilder) group(group string) {
	p.sql.group = " GROUP BY " + group
}

func (p *sqlBuilder) havings(having string) {
	if 0 == p.sql.havings.Len() {
		p.sql.havings = bytes.NewBufferString(" HAVING ( ")
	} else {
		p.sql.havings.Truncate(p.sql.havings.Len() - 3)
	}
	_, err := p.sql.havings.WriteString(having + " ) ")
	if nil != err {

	}
}

func (p *sqlBuilder) gen() (sql string, err error) {
	sqlBuffers := &bytes.Buffer{}
	tables := p.sql.tables.String()
	joins := p.sql.joins.String()
	wheres := p.sql.wheres.String()
	columns := p.sql.columns.String()
	if "" != p.sql.queryType {
		columns = string(p.sql.queryType) + "( " + columns + " )"
	}
	switch p.sql.operate {
	case easydb.Insert:
		columns = " (" + columns + ") "
		values := []string{}
		for range strings.Split(columns, ",") {
			values = append(values, "?")
		}
		valueStr := " (" + strings.Join(values, ",") + ") "
		sqlBuffers.WriteString(string(easydb.Insert))
		sqlBuffers.WriteString(tables)
		sqlBuffers.WriteString(columns)
		sqlBuffers.WriteString(" VALUES ")
		sqlBuffers.WriteString(valueStr)
		break
	case easydb.Delete:
		sqlBuffers.WriteString(string(easydb.Delete))
		sqlBuffers.WriteString(tables)
		sqlBuffers.WriteString(joins)
		sqlBuffers.WriteString(wheres)
		break
	case easydb.Update:
		sqlBuffers.WriteString(string(easydb.Update))
		sqlBuffers.WriteString(tables)
		sqlBuffers.WriteString(joins)
		sqlBuffers.WriteString(" SET ")
		sqlBuffers.WriteString(columns)
		sqlBuffers.WriteString(wheres)
		break
	case easydb.Select:
		havings := p.sql.havings.String()
		sqlBuffers.WriteString(string(easydb.Select))
		sqlBuffers.WriteString(columns)
		sqlBuffers.WriteString(" FROM ")
		sqlBuffers.WriteString(tables)
		sqlBuffers.WriteString(joins)
		sqlBuffers.WriteString(wheres)
		sqlBuffers.WriteString(p.sql.group)
		sqlBuffers.WriteString(havings)
		sqlBuffers.WriteString(p.sql.order)
		sqlBuffers.WriteString(p.sql.limit)
		sqlBuffers.WriteString(p.sql.offset)
		break

	default:
		break
	}
	return sqlBuffers.String(), nil
}
