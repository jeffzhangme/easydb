package mysql

import (
	"github.com/jeffzhangme/easydb"
	"github.com/jeffzhangme/easydb/restrict"
)

// DeleteBuilder build delete
type DeleteBuilder struct {
	QueryBuilder
}

// BuildDelete begin
func BuildDelete(dbName ...string) restrict.IBuildDeleteReturn {
	p := &DeleteBuilder{}
	p.sqlBuilder = build(easydb.Delete)
	p.values = map[string][]string{}
	if len(dbName) > 0 && len(dbName[0]) > 0 {
		p.dbName = dbName[0] + "."
	}
	return p
}

// Table set table
func (p *DeleteBuilder) Table(table easydb.Table) restrict.IDTableReturn {
	p.QueryBuilder.Tables(table)
	return p
}

// Where add where
func (p *DeleteBuilder) Where(where easydb.Where) restrict.IDWheresReturn {
	p.QueryBuilder.Where(where)
	return p
}

// And and
func (p *DeleteBuilder) And() restrict.IDAndReturn {
	p.QueryBuilder.And()
	return p
}

// Or or
func (p *DeleteBuilder) Or() restrict.IDOrReturn {
	p.QueryBuilder.Or()
	return p
}

// StartGroup start new group
func (p *DeleteBuilder) StartGroup() restrict.IDStartGroup {
	p.QueryBuilder.StartGroup()
	return p
}

// EndGroup end group
func (p *DeleteBuilder) EndGroup() restrict.IDEndGroup {
	p.QueryBuilder.EndGroup()
	return p
}
