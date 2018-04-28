package mysql

import (
	"github.com/jeffzhangme/easydb"
	"github.com/jeffzhangme/easydb/restrict"
	"strings"
)

// UpdateBuilder build update
type UpdateBuilder struct {
	DeleteBuilder
}

// BuildUpdate begin
func BuildUpdate(dbName ...string) restrict.IBuildUpdateReturn {
	p := &UpdateBuilder{}
	p.sqlBuilder = build(easydb.Update)
	p.values = map[string][]string{}
	p.values["value"] = []string{}
	if len(dbName) > 0 && len(dbName[0]) > 0 {
		p.dbName = dbName[0] + "."
	}
	return p
}

// Table set table
func (p *UpdateBuilder) Table(table easydb.Table) restrict.IUTableReturn {
	p.DeleteBuilder.Table(table)
	return p
}

// Set set
func (p *UpdateBuilder) Set(columns ...easydb.Column) restrict.IUSetReturn {
	for _, column := range columns {
		if strings.Contains(column.Value, "`") {
			p.sqlBuilder.columns(column.Name + " = " + column.Value)
		} else {
			p.sqlBuilder.columns(column.Name + " = ? ")
			p.values["value"] = append(p.values["value"], column.Value)
		}
	}
	return p
}
