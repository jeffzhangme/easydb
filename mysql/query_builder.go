package mysql

import (
	"bytes"
	"strconv"
	"strings"
	"unsafe"

	"github.com/jeffzhangme/easydb"
	"github.com/jeffzhangme/easydb/restrict"
)

// QueryValue query value
type QueryValue struct {
	sqlBuilder
	values map[string][]string
	dbName string
}

// NewQueryValue init query value
func NewQueryValue(opts ...easydb.DBOptType) QueryValue {
	opt := easydb.Select
	if len(opts) > 0 {
		opt = opts[0]
	}
	p := QueryValue{sqlBuilder: build(opt)}
	p.values = map[string][]string{}
	return p
}

// QueryBuilder build query
type QueryBuilder struct {
	QueryValue
}

// BuildQuery begin
func BuildQuery(dbName ...string) restrict.IBuildQueryReturn {
	p := &QueryBuilder{}
	p.sqlBuilder = build(easydb.Select)
	if len(dbName) > 0 && len(dbName[0]) > 0 {
		p.dbName = dbName[0] + "."
	}
	return p
}

// QueryFuncs set query func
func (p *QueryBuilder) QueryFuncs(funccs ...easydb.QueryFunc) restrict.IQFuncReturn {
	for _, funcc := range funccs {
		as := ""
		if "" != funcc.As {
			as = " AS " + funcc.As
		}
		p.sqlBuilder.columns(string(funcc.Type) + "(" + strings.Join(funcc.Names, ", ") + ")" + as)
	}

	return p
}

// Columns set columns
func (p *QueryBuilder) Columns(columns ...easydb.Column) restrict.IColumnsReturn {
	for _, column := range columns {
		as := ""
		if "" != column.As {
			as = " AS " + column.As
		}
		p.sqlBuilder.columns(column.Name + as)
	}
	return p
}

// Tables set table
func (p *QueryBuilder) Tables(tables ...easydb.Table) restrict.IFromsReturn {
	for _, table := range tables {
		as := " "
		if "" != table.As {
			as += table.As
		}
		if !strings.Contains(table.Name, ".") {
			table.Name = p.dbName + table.Name
		}
		p.sqlBuilder.tables(table.Name + as)
	}
	return p
}

// LeftJoin left join
func (p *QueryBuilder) LeftJoin(table easydb.Table) restrict.IJoinReturn {
	as := " "
	if "" != table.As {
		as += table.As
	}
	if !strings.Contains(table.Name, ".") {
		table.Name = p.dbName + table.Name
	}
	p.sqlBuilder.join(string(easydb.LeftJoin) + table.Name + as)
	return p
}

// RightJoin right join
func (p *QueryBuilder) RightJoin(table easydb.Table) restrict.IJoinReturn {
	as := " "
	if "" != table.As {
		as += table.As
	}
	if !strings.Contains(table.Name, ".") {
		table.Name = p.dbName + table.Name
	}
	p.sqlBuilder.join(string(easydb.RightJoin) + table.Name + as)
	return p
}

// On on
func (p *QueryBuilder) On(on easydb.On) restrict.IOnReturn {
	where := (*easydb.Where)(unsafe.Pointer(&on))
	p.sqlBuilder.on(likeWhere(p, "join", where))
	return p
}

// Where set where
func (p *QueryBuilder) Where(where easydb.Where) restrict.IWheresReturn {
	p.sqlBuilder.wheres(likeWhere(p, "where", &where))
	return p
}

// StartGroup start new group
func (p *QueryBuilder) StartGroup() restrict.IStartGroupReturn {
	sql, _ := p.Gen()
	if strings.Contains(sql, "HAVING") {
		p.sqlBuilder.havings(string(easydb.GroupStart))
	} else if strings.Contains(sql, "WHERE") {
		p.sqlBuilder.wheres(string(easydb.GroupStart))
	} else {
		p.sqlBuilder.join(string(easydb.GroupStart))
	}
	return p
}

// EndGroup end group
func (p *QueryBuilder) EndGroup() restrict.IEndGroupReturn {
	sql, _ := p.Gen()
	if strings.Contains(sql, "HAVING") {
		p.sqlBuilder.havings(string(easydb.GroupEnd))
	} else if strings.Contains(sql, "WHERE") {
		p.sqlBuilder.wheres(string(easydb.GroupEnd))
	} else {
		p.sqlBuilder.join(string(easydb.GroupEnd))
	}
	return p
}

// And and
func (p *QueryBuilder) And() restrict.IAndReturn {
	sql, _ := p.Gen()
	if strings.Contains(sql, "HAVING") {
		p.sqlBuilder.havings(string(easydb.AND))
	} else if strings.Contains(sql, "WHERE") {
		p.sqlBuilder.wheres(string(easydb.AND))
	} else {
		p.sqlBuilder.join(string(easydb.AND))
	}
	return p
}

// Or or
func (p *QueryBuilder) Or() restrict.IOrReturn {
	sql, _ := p.Gen()
	if strings.Contains(sql, "HAVING") {
		p.sqlBuilder.havings(string(easydb.OR))
	} else if strings.Contains(sql, "WHERE") {
		p.sqlBuilder.wheres(string(easydb.OR))
	} else {
		p.sqlBuilder.join(string(easydb.OR))
	}
	return p
}

// GroupBy group by
func (p *QueryBuilder) GroupBy(group ...string) restrict.IGroupByReturn {
	p.sqlBuilder.group(" ( " + strings.Join(group, ", ") + " ) ")
	return p
}

// Having set having
func (p *QueryBuilder) Having(having easydb.Having) restrict.IHavingsReturn {
	where := (*easydb.Where)(unsafe.Pointer(&having))
	p.sqlBuilder.havings(likeWhere(p, "having", where))
	return p
}

// OrderBy order by
func (p *QueryBuilder) OrderBy(orders ...easydb.Order) restrict.IOrderByReturn {
	orderStr := bytes.NewBufferString("")
	for _, order := range orders {
		orderStr.WriteString(order.Key + " " + string(order.Type) + ", ")
	}
	orderStr.Truncate(orderStr.Len() - 2)
	p.sqlBuilder.order(orderStr.String())
	return p
}

// Limit set limit
func (p *QueryBuilder) Limit(limit int) restrict.ILimitReturn {
	limitStr := strconv.Itoa(limit)
	p.values["limit"] = append(p.values["limit"], limitStr)
	p.sqlBuilder.limit(" ? ")
	return p
}

// Offset set offset
func (p *QueryBuilder) Offset(offset int) restrict.IOffsetReturn {
	offsetStr := strconv.Itoa(offset)
	p.values["offset"] = append(p.values["offset"], offsetStr)
	p.sqlBuilder.offset(" ? ")
	return p
}

// Gen get sql
func (p *QueryBuilder) Gen() (sql string, err error) {
	return p.sqlBuilder.gen()
}

// Val get values
func (p *QueryBuilder) Val() []string {
	values := []string{}
	values = append(values, p.values["join"]...)
	values = append(values, p.values["value"]...)
	values = append(values, p.values["where"]...)
	values = append(values, p.values["having"]...)
	values = append(values, p.values["limit"]...)
	values = append(values, p.values["offset"]...)
	return values
}
func likeWhere(p *QueryBuilder, t string, where *easydb.Where) string {
	whereStr := bytes.NewBufferString(where.Key)
	switch where.Opt {
	case easydb.IN:
		placeholder := []string{}
		for _, v := range where.Ins {
			if strings.Contains(v, "`") {
				if strings.Contains(v, ".") {
					placeholder = append(placeholder, v[1:len(v)-1])
				} else {
					placeholder = append(placeholder, v)
				}
			} else {
				placeholder = append(placeholder, "?")
				p.values[t] = append(p.values[t], v)
			}
		}
		whereStr.WriteString(" IN ( ")
		whereStr.WriteString(strings.Join(placeholder, ", "))
		whereStr.WriteString(" ) ")
		break
	case easydb.IsNULL:
		whereStr.WriteString(string(easydb.IsNULL))
		break
	case easydb.NotNULL:
		whereStr.WriteString(string(easydb.NotNULL))
		break
	default:
		whereStr.WriteString(string(where.Opt))
		if strings.Contains(where.Value, "`") {
			if strings.Contains(where.Value, ".") {
				whereStr.WriteString(where.Value[1 : len(where.Value)-1])
			} else {
				whereStr.WriteString(where.Value)
			}
		} else {
			whereStr.WriteString(" ? ")
			p.values[t] = append(p.values[t], where.Value)
		}
	}
	return whereStr.String()
}
