package easydb

import (
	"strings"
)

// updateBuilder build update
type updateBuilder struct {
	deleteBuilder
}

// BuildUpdate begin
func BuildUpdate(dbName ...string) iBuildUpdateReturn {
	p := &updateBuilder{}
	p.queryValue = newQueryValue(Update)
	if len(dbName) > 0 && len(dbName[0]) > 0 {
		p.dbName = dbName[0] + "."
	}
	return p
}

// Table set table
func (p *updateBuilder) Table(table Table) iUTableReturn {
	p.deleteBuilder.Table(table)
	return p
}

// Set set
func (p *updateBuilder) Set(columns ...Column) iUSetReturn {
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
