package pgsql

import (
	"github.com/jeffzhangme/easydb"
	"github.com/jeffzhangme/easydb/mysql"
	"github.com/jeffzhangme/easydb/restrict"
	"fmt"
	"strconv"
	"strings"
)

// InsertBuilder build insert
type InsertBuilder struct {
	mysql.InsertBuilder
	schema string
}

// BuildInsert begin
func BuildInsert(schema ...string) restrict.IBuildInsertReturn {
	p := &InsertBuilder{}
	p.QueryValue = mysql.NewQueryValue(easydb.Insert)
	if len(schema) > 0 && len(schema[0]) > 0 {
		p.schema = schema[0] + "."
	}
	return p
}

// Table set table
func (p *InsertBuilder) Table(table easydb.Table) restrict.IITableReturn {
	if !strings.Contains(table.Name, ".") {
		table.Name = p.schema + table.Name
	}
	p.InsertBuilder.Table(table)
	return p
}

// Values set values
func (p *InsertBuilder) Values(columns ...easydb.Column) restrict.IIValuesReturn {
	p.InsertBuilder.Values(columns...)
	return p
}

// Gen get sql
func (p *InsertBuilder) Gen() (sql string, err error) {
	placeholder := make([]interface{}, len(p.InsertBuilder.Val()))
	for i, _ := range p.InsertBuilder.Val() {
		placeholder[i] = "$" + strconv.Itoa(i+1)
	}
	sql, err = p.InsertBuilder.Gen()
	sql = strings.Replace(sql, "?", "%s", -1)
	if len(placeholder) > 0 {
		sql = fmt.Sprintf(sql, placeholder...)
	}
	return
}
