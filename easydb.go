package easydb

// GetInst GetInst
func GetInst(dbType DBType) iAdapter {
	var adapter iAdapter
	switch dbType {
	case MYSQL:
		adapter = &dbAdapter{initMysql()}
		break
	case PGSQL:
		adapter = &dbAdapter{initPgsql()}
		break
	}
	return adapter
}

// Close Close
func Close() {
	if nil != mysqlInst {
		mysqlInst.Close()
	}
	if nil != pgsqlInst {
		pgsqlInst.Close()
	}
}
