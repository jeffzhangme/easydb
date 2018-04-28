package pgsql

import (
	"github.com/jeffzhangme/easydb"
	"github.com/jeffzhangme/easydb/mysql"
	"github.com/jeffzhangme/easydb/restrict"
	"fmt"
	"strconv"
	"strings"
)

// QueryBuilder build query
type QueryBuilder struct {
	mysql.QueryBuilder
	schema string
}

// BuildQuery begin
func BuildQuery(schema ...string) restrict.IBuildQueryReturn {
	p := &QueryBuilder{}
	p.QueryValue = mysql.NewQueryValue()
	if len(schema) > 0 && len(schema[0]) > 0 {
		p.schema = schema[0] + "."
	}
	return p
}

// QueryFuncs set query func
func (p *QueryBuilder) QueryFuncs(funccs ...easydb.QueryFunc) restrict.IQFuncReturn {
	p.QueryBuilder.QueryFuncs(funccs...)
	return p
}

// Columns set columns
func (p *QueryBuilder) Columns(columns ...easydb.Column) restrict.IColumnsReturn {
	p.QueryBuilder.Columns(columns...)
	return p
}

// Tables set table
func (p *QueryBuilder) Tables(tables ...easydb.Table) restrict.IFromsReturn {
	fullTables := []easydb.Table{}
	for _, table := range tables {
		if !strings.Contains(table.Name, ".") {
			table.Name = p.schema + table.Name
		}
		fullTables = append(fullTables, table)
	}
	p.QueryBuilder.Tables(fullTables...)
	return p
}

// LeftJoin left join
func (p *QueryBuilder) LeftJoin(table easydb.Table) restrict.IJoinReturn {
	p.QueryBuilder.LeftJoin(table)
	return p
}

// RightJoin right join
func (p *QueryBuilder) RightJoin(table easydb.Table) restrict.IJoinReturn {
	p.QueryBuilder.RightJoin(table)
	return p
}

// On on
func (p *QueryBuilder) On(on easydb.On) restrict.IOnReturn {
	p.QueryBuilder.On(on)
	return p
}

// Where set where
func (p *QueryBuilder) Where(where easydb.Where) restrict.IWheresReturn {
	p.QueryBuilder.Where(where)
	return p
}

// StartGroup start new group
func (p *QueryBuilder) StartGroup() restrict.IStartGroupReturn {
	p.QueryBuilder.StartGroup()
	return p
}

// EndGroup end group
func (p *QueryBuilder) EndGroup() restrict.IEndGroupReturn {
	p.QueryBuilder.EndGroup()
	return p
}

// And and
func (p *QueryBuilder) And() restrict.IAndReturn {
	p.QueryBuilder.And()
	return p
}

// Or or
func (p *QueryBuilder) Or() restrict.IOrReturn {
	p.QueryBuilder.Or()
	return p
}

// GroupBy group by
func (p *QueryBuilder) GroupBy(group ...string) restrict.IGroupByReturn {
	p.QueryBuilder.GroupBy(group...)
	return p
}

// Having set having
func (p *QueryBuilder) Having(having easydb.Having) restrict.IHavingsReturn {
	p.QueryBuilder.Having(having)
	return p
}

// OrderBy order by
func (p *QueryBuilder) OrderBy(orders ...easydb.Order) restrict.IOrderByReturn {
	p.QueryBuilder.OrderBy(orders...)
	return p
}

// Limit set limit
func (p *QueryBuilder) Limit(limit int) restrict.ILimitReturn {
	p.QueryBuilder.Limit(limit)
	return p
}

// Offset set offset
func (p *QueryBuilder) Offset(offset int) restrict.IOffsetReturn {
	p.QueryBuilder.Offset(offset)
	return p
}

// Gen get sql
func (p *QueryBuilder) Gen() (sql string, err error) {
	placeholder := make([]interface{}, len(p.QueryBuilder.Val()))
	for i, _ := range p.QueryBuilder.Val() {
		placeholder[i] = "$" + strconv.Itoa(i+1)
	}
	sql, err = p.QueryBuilder.Gen()
	sql = strings.Replace(sql, "?", "%s", -1)
	if len(placeholder) > 0 {
		sql = fmt.Sprintf(sql, placeholder...)
	}
	return
}

// Val get values
func (p *QueryBuilder) Val() []string {
	return p.QueryBuilder.Val()
}
