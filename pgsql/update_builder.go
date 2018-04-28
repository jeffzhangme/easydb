package pgsql

import (
	"github.com/jeffzhangme/easydb"
	"github.com/jeffzhangme/easydb/mysql"
	"github.com/jeffzhangme/easydb/restrict"
	"fmt"
	"strconv"
	"strings"
)

// UpdateBuilder build update
type UpdateBuilder struct {
	mysql.UpdateBuilder
	schema string
}

// BuildUpdate begin
func BuildUpdate(schema ...string) restrict.IBuildUpdateReturn {
	p := &UpdateBuilder{}
	p.QueryValue = mysql.NewQueryValue(easydb.Update)
	if len(schema) > 0 && len(schema[0]) > 0 {
		p.schema = schema[0] + "."
	}
	return p
}

// Table set table
func (p *UpdateBuilder) Table(table easydb.Table) restrict.IUTableReturn {
	if !strings.Contains(table.Name, ".") {
		table.Name = p.schema + table.Name
	}
	p.UpdateBuilder.Table(table)
	return p
}

// Set set
func (p *UpdateBuilder) Set(columns ...easydb.Column) restrict.IUSetReturn {
	p.UpdateBuilder.Set(columns...)
	return p
}

// Gen get sql
func (p *UpdateBuilder) Gen() (sql string, err error) {
	placeholder := make([]interface{}, len(p.UpdateBuilder.Val()))
	for i, _ := range p.UpdateBuilder.Val() {
		placeholder[i] = "$" + strconv.Itoa(i+1)
	}
	sql, err = p.UpdateBuilder.Gen()
	sql = strings.Replace(sql, "?", "%s", -1)
	if len(placeholder) > 0 {
		sql = fmt.Sprintf(sql, placeholder...)
	}
	return
}
// Where set where
func (p *UpdateBuilder) Where(where easydb.Where) restrict.IDWheresReturn {
	p.UpdateBuilder.Where(where)
	return p
}

// And and
func (p *UpdateBuilder) And() restrict.IDAndReturn {
	p.UpdateBuilder.And()
	return p
}

// Or or
func (p *UpdateBuilder) Or() restrict.IDOrReturn {
	p.UpdateBuilder.Or()
	return p
}

// StartGroup start new group
func (p *UpdateBuilder) StartGroup() restrict.IDStartGroup {
	p.UpdateBuilder.StartGroup()
	return p
}

// EndGroup end group
func (p *UpdateBuilder) EndGroup() restrict.IDEndGroup {
	p.UpdateBuilder.EndGroup()
	return p
}
