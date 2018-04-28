package easydb

type iSQLBuilder interface {
	Gen() (sql string, err error)
	Val() []string
}

type iAdapter interface {
	Insert(sqlBuilder iSQLBuilder) error
	Delete(sqlBuilder iSQLBuilder) error
	Update(sqlBuilder iSQLBuilder) error
	Select(sqlBuilder iSQLBuilder) ([]map[string]interface{}, error)
}

type iDBOperate interface {
	Do(optType DBOptType, sqlBuilder iSQLBuilder) ([]map[string]interface{}, error)
}

// dbAdapter adapter
type dbAdapter struct {
	iDBOperate
}

//Insert Insert
func (p *dbAdapter) Insert(sqlBuilder iSQLBuilder) error {
	_, err := p.Do(Insert, sqlBuilder)
	return err
}

//Update Update
func (p *dbAdapter) Update(sqlBuilder iSQLBuilder) error {
	_, err := p.Do(Update, sqlBuilder)
	return err
}

//Delete Delete
func (p *dbAdapter) Delete(sqlBuilder iSQLBuilder) error {
	_, err := p.Do(Delete, sqlBuilder)
	return err
}

//Select Select
func (p *dbAdapter) Select(sqlBuilder iSQLBuilder) ([]map[string]interface{}, error) {
	return p.Do(Select, sqlBuilder)
}
