package pgsql

import (
	"github.com/jeffzhangme/easydb"
	"github.com/jeffzhangme/easydb/mysql"
	"github.com/jeffzhangme/easydb/restrict"
)

// DeleteBuilder build delete
type DeleteBuilder struct {
	QueryBuilder
}

// BuildDelete begin
func BuildDelete(schema ...string) restrict.IBuildDeleteReturn {
	p := &DeleteBuilder{}
	p.QueryValue = mysql.NewQueryValue(easydb.Delete)
	if len(schema) > 0 && len(schema[0]) > 0 {
		p.schema = schema[0] + "."
	}
	return p
}

// Table set table
func (p *DeleteBuilder) Table(table easydb.Table) restrict.IDTableReturn {
	p.Tables(table)
	return p
}

// Where set where
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
