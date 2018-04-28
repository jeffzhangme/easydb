package mysql

import (
	"github.com/jeffzhangme/easydb"
	"github.com/jeffzhangme/easydb/restrict"
)

// InsertBuilder build insert
type InsertBuilder struct {
	DeleteBuilder
}

// BuildInsert begin
func BuildInsert(dbName ...string) restrict.IBuildInsertReturn {
	p := &InsertBuilder{}
	p.sqlBuilder = build(easydb.Insert)
	p.values = map[string][]string{}
	if len(dbName) > 0 && len(dbName[0]) > 0 {
		p.dbName = dbName[0] + "."
	}
	return p
}

// Table set table
func (p *InsertBuilder) Table(table easydb.Table) restrict.IITableReturn {
	p.DeleteBuilder.Table(table)
	return p
}

// Values set values
func (p *InsertBuilder) Values(columns ...easydb.Column) restrict.IIValuesReturn {
	for _, column := range columns {
		p.sqlBuilder.columns(column.Name)
		p.values["value"] = append(p.values["value"], column.Value)
	}
	return p
}
