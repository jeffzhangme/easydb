package easydb

// insertBuilder build insert
type insertBuilder struct {
	deleteBuilder
}

// BuildInsert begin
func BuildInsert(dbName ...string) iBuildInsertReturn {
	p := &insertBuilder{}
	p.queryValue = newQueryValue(Insert)
	if len(dbName) > 0 && len(dbName[0]) > 0 {
		p.dbName = dbName[0] + "."
	}
	return p
}

// Table set table
func (p *insertBuilder) Table(table Table) iITableReturn {
	p.deleteBuilder.Table(table)
	return p
}

// Values set values
func (p *insertBuilder) Values(columns ...Column) iIValuesReturn {
	for _, column := range columns {
		p.sqlBuilder.columns(column.Name)
		p.values["value"] = append(p.values["value"], column.Value)
	}
	return p
}

func (p *insertBuilder) __insert() {
}
