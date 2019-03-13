package easydb

type iSQLBuilder interface {
	Gen() (sql string, err error)
	Val() []string
}

type iDBOperate interface {
	Do(optType dbOptType, sqlBuilder iSQLBuilder) ([]map[string]interface{}, error)
}

// DBExec db executor
type DBExec struct {
	iDBOperate
}

//Insert Insert
func (p *DBExec) Insert(sqlBuilder iSQLBuilder) error {
	_, err := p.Do(Insert, sqlBuilder)
	return err
}

//Update Update
func (p *DBExec) Update(sqlBuilder iSQLBuilder) error {
	_, err := p.Do(Update, sqlBuilder)
	return err
}

//Delete Delete
func (p *DBExec) Delete(sqlBuilder iSQLBuilder) error {
	_, err := p.Do(Delete, sqlBuilder)
	return err
}

//Select Select
func (p *DBExec) Select(sqlBuilder iSQLBuilder) ([]map[string]interface{}, error) {
	return p.Do(Select, sqlBuilder)
}
