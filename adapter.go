package easydb

type iSQLBuilder interface {
	Gen() (sql string, err error)
	Val() []interface{}
}

type iDBOperate interface {
	Do(optType dbOptType, sqlBuilder iSQLBuilder) ([]map[string]interface{}, error)
}

// DBExec db executor
type DBExec struct {
	iDBOperate
}

//Insert Insert
func (p *DBExec) Insert(sqlBuilder iSQLBuilder) ([]map[string]interface{}, error) {
	return p.Do(Insert, sqlBuilder)
}

//Update Update
func (p *DBExec) Update(sqlBuilder iSQLBuilder) ([]map[string]interface{}, error) {
	return p.Do(Update, sqlBuilder)
}

//Delete Delete
func (p *DBExec) Delete(sqlBuilder iSQLBuilder) ([]map[string]interface{}, error) {
	return p.Do(Delete, sqlBuilder)
}

//Select Select
func (p *DBExec) Select(sqlBuilder iSQLBuilder) ([]map[string]interface{}, error) {
	return p.Do(Select, sqlBuilder)
}
