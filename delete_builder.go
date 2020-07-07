package easydb

// deleteBuilder build delete
type deleteBuilder struct {
	queryBuilder
}

// BuildDelete begin
func BuildDelete(dbName ...string) iBuildDeleteReturn {
	p := &deleteBuilder{}
	p.queryValue = newQueryValue(Delete)
	if len(dbName) > 0 && len(dbName[0]) > 0 {
		p.dbName = dbName[0] + "."
	}
	return p
}

// Table set table
func (p *deleteBuilder) Table(table Table) iDTableReturn {
	p.queryBuilder.Tables(table)
	return p
}

// Where add where
func (p *deleteBuilder) Where(where Where) iDWheresReturn {
	p.queryBuilder.Where(where)
	return p
}

// And and
func (p *deleteBuilder) And() iDAndReturn {
	p.queryBuilder.And()
	return p
}

// Or or
func (p *deleteBuilder) Or() iDOrReturn {
	p.queryBuilder.Or()
	return p
}

// StartGroup start new group
func (p *deleteBuilder) StartGroup() iDStartGroup {
	p.queryBuilder.StartGroup()
	return p
}

// EndGroup end group
func (p *deleteBuilder) EndGroup() iDEndGroup {
	p.queryBuilder.EndGroup()
	return p
}

func (p *deleteBuilder) __delete() {
}
