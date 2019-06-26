package easydb

import (
	"bytes"
	"strconv"
	"strings"
	"unsafe"
)

// queryValue query value
type queryValue struct {
	sqlBuilder
	values map[string][]interface{}
	dbName string
}

// newQueryValue init query value
func newQueryValue(opts ...dbOptType) queryValue {
	opt := Select
	if len(opts) > 0 {
		opt = opts[0]
	}
	p := queryValue{sqlBuilder: build(opt)}
	p.values = map[string][]interface{}{}
	return p
}

// queryBuilder build query
type queryBuilder struct {
	queryValue
}

// BuildQuery begin
func BuildQuery(dbName ...string) iBuildQueryReturn {
	p := &queryBuilder{queryValue: newQueryValue()}
	if len(dbName) > 0 && len(dbName[0]) > 0 {
		p.dbName = dbName[0] + "."
	}
	return p
}

// QueryFuncs set query func
func (p *queryBuilder) QueryFuncs(funccs ...QueryFunc) iQFuncReturn {
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
func (p *queryBuilder) Columns(columns ...Column) iColumnsReturn {
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
func (p *queryBuilder) Tables(tables ...Table) iFromsReturn {
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
func (p *queryBuilder) LeftJoin(table Table) iJoinReturn {
	as := " "
	if "" != table.As {
		as += table.As
	}
	if !strings.Contains(table.Name, ".") {
		table.Name = p.dbName + table.Name
	}
	p.sqlBuilder.join(string(LeftJoin) + table.Name + as)
	return p
}

// RightJoin right join
func (p *queryBuilder) RightJoin(table Table) iJoinReturn {
	as := " "
	if "" != table.As {
		as += table.As
	}
	if !strings.Contains(table.Name, ".") {
		table.Name = p.dbName + table.Name
	}
	p.sqlBuilder.join(string(RightJoin) + table.Name + as)
	return p
}

// On on
func (p *queryBuilder) On(on On) iOnReturn {
	where := (*Where)(unsafe.Pointer(&on))
	p.sqlBuilder.on(likeWhere(p, "join", where))
	return p
}

// Where set where
func (p *queryBuilder) Where(where Where) iWheresReturn {
	p.sqlBuilder.wheres(likeWhere(p, "where", &where))
	return p
}

// StartGroup start new group
func (p *queryBuilder) StartGroup() iStartGroupReturn {
	sql, _ := p.Gen()
	if strings.Contains(sql, "HAVING") {
		p.sqlBuilder.havings(string(GroupStart))
	} else if strings.Contains(sql, "WHERE") {
		p.sqlBuilder.wheres(string(GroupStart))
	} else {
		p.sqlBuilder.join(string(GroupStart))
	}
	return p
}

// EndGroup end group
func (p *queryBuilder) EndGroup() iEndGroupReturn {
	sql, _ := p.Gen()
	if strings.Contains(sql, "HAVING") {
		p.sqlBuilder.havings(string(GroupEnd))
	} else if strings.Contains(sql, "WHERE") {
		p.sqlBuilder.wheres(string(GroupEnd))
	} else {
		p.sqlBuilder.join(string(GroupEnd))
	}
	return p
}

// And and
func (p *queryBuilder) And() iAndReturn {
	sql, _ := p.Gen()
	if strings.Contains(sql, "HAVING") {
		p.sqlBuilder.havings(string(AND))
	} else if strings.Contains(sql, "WHERE") {
		p.sqlBuilder.wheres(string(AND))
	} else {
		p.sqlBuilder.join(string(AND))
	}
	return p
}

// Or or
func (p *queryBuilder) Or() iOrReturn {
	sql, _ := p.Gen()
	if strings.Contains(sql, "HAVING") {
		p.sqlBuilder.havings(string(OR))
	} else if strings.Contains(sql, "WHERE") {
		p.sqlBuilder.wheres(string(OR))
	} else {
		p.sqlBuilder.join(string(OR))
	}
	return p
}

// GroupBy group by
func (p *queryBuilder) GroupBy(group ...string) iGroupByReturn {
	p.sqlBuilder.group(" ( " + strings.Join(group, ", ") + " ) ")
	return p
}

// Having set having
func (p *queryBuilder) Having(having Having) iHavingsReturn {
	where := (*Where)(unsafe.Pointer(&having))
	p.sqlBuilder.havings(likeWhere(p, "having", where))
	return p
}

// OrderBy order by
func (p *queryBuilder) OrderBy(orders ...Order) iOrderByReturn {
	orderStr := bytes.NewBufferString("")
	for _, order := range orders {
		orderStr.WriteString(order.Key + " " + string(order.Type) + ", ")
	}
	orderStr.Truncate(orderStr.Len() - 2)
	p.sqlBuilder.order(orderStr.String())
	return p
}

// Limit set limit
func (p *queryBuilder) Limit(limit int) iLimitReturn {
	limitStr := strconv.Itoa(limit)
	p.values["limit"] = append(p.values["limit"], limitStr)
	p.sqlBuilder.limit(" ? ")
	return p
}

// Offset set offset
func (p *queryBuilder) Offset(offset int) iOffsetReturn {
	offsetStr := strconv.Itoa(offset)
	p.values["offset"] = append(p.values["offset"], offsetStr)
	p.sqlBuilder.offset(" ? ")
	return p
}

// Gen get sql
func (p *queryBuilder) Gen() (sql string, err error) {
	return p.sqlBuilder.gen()
}

// Val get values
func (p *queryBuilder) Val() []interface{} {
	values := []interface{}{}
	values = append(values, p.values["join"]...)
	values = append(values, p.values["value"]...)
	values = append(values, p.values["where"]...)
	values = append(values, p.values["having"]...)
	values = append(values, p.values["limit"]...)
	values = append(values, p.values["offset"]...)
	return values
}
func likeWhere(p *queryBuilder, t string, where *Where) string {
	whereStr := bytes.NewBufferString(where.Key)
	switch where.Opt {
	case IN:
		placeholder := []string{}
		for _, iv := range where.Ins {
			if v, ok := iv.(string); ok {
				if strings.HasPrefix(v, "`") && strings.HasSuffix(v, "`") {
					if strings.Contains(v, ".") {
						placeholder = append(placeholder, v[1:len(v)-1])
					} else {
						placeholder = append(placeholder, v)
					}
				} else {
					placeholder = append(placeholder, "?")
					p.values[t] = append(p.values[t], iv)
				}
			} else {
				placeholder = append(placeholder, "?")
				p.values[t] = append(p.values[t], iv)
			}
		}
		whereStr.WriteString(" IN ( ")
		whereStr.WriteString(strings.Join(placeholder, ", "))
		whereStr.WriteString(" ) ")
		break
	case IsNULL:
		whereStr.WriteString(string(IsNULL))
		break
	case NotNULL:
		whereStr.WriteString(string(NotNULL))
		break
	default:
		whereStr.WriteString(string(where.Opt))
		if v, ok := where.Value.(string); ok {
			if strings.HasPrefix(v, "`") && strings.HasSuffix(v, "`") {
				if strings.Contains(v, ".") {
					whereStr.WriteString(v[1 : len(v)-1])
				} else {
					whereStr.WriteString(v)
				}
			} else {
				whereStr.WriteString(" ? ")
				p.values[t] = append(p.values[t], where.Value)
			}
		} else {
			whereStr.WriteString(" ? ")
			p.values[t] = append(p.values[t], where.Value)
		}
	}
	return whereStr.String()
}
