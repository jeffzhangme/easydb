package easydb

// iBuildQueryReturn iBuildQueryReturn
type iBuildQueryReturn interface {
	QueryFuncs(funccs ...QueryFunc) iQFuncReturn
	Columns(columns ...Column) iColumnsReturn
}

// iQFuncReturn iQFuncReturn
type iQFuncReturn interface {
	QueryFuncs(funccs ...QueryFunc) iQFuncReturn
	Columns(columns ...Column) iColumnsReturn
	Tables(tables ...Table) iFromsReturn
}

// iColumnsReturn iColumnsReturn
type iColumnsReturn interface {
	iQFuncReturn
}

// iFromsReturn iFromsReturn
type iFromsReturn interface {
	Tables(tables ...Table) iFromsReturn
	RightJoin(table Table) iJoinReturn
	LeftJoin(table Table) iJoinReturn
	Where(where Where) iWheresReturn
	GroupBy(group ...string) iGroupByReturn
	OrderBy(orders ...Order) iOrderByReturn
	Limit(limit int) iLimitReturn
	Gen() (sql string, err error)
	Val() []interface{}
}

// iJoinReturn iJoinReturn
type iJoinReturn interface {
	On(on On) iOnReturn
}

// iOnReturn iOnReturn
type iOnReturn interface {
	RightJoin(table Table) iJoinReturn
	LeftJoin(table Table) iJoinReturn
	Where(where Where) iWheresReturn
	GroupBy(group ...string) iGroupByReturn
	OrderBy(orders ...Order) iOrderByReturn
	Limit(limit int) iLimitReturn
	Gen() (sql string, err error)
	Val() []interface{}
	And() iAndReturn
	Or() iOrReturn
}

// iWheresReturn iWheresReturn
type iWheresReturn interface {
	And() iAndReturn
	Or() iOrReturn
	GroupBy(group ...string) iGroupByReturn
	OrderBy(orders ...Order) iOrderByReturn
	Limit(limit int) iLimitReturn
	EndGroup() iEndGroupReturn
	Gen() (sql string, err error)
	Val() []interface{}
}

// iGroupByReturn 制约groupby条件 后调用
type iGroupByReturn interface {
	Having(having Having) iHavingsReturn
	OrderBy(orders ...Order) iOrderByReturn
	Limit(limit int) iLimitReturn
	Gen() (sql string, err error)
	Val() []interface{}
}

// iHavingsReturn iHavingsReturn
type iHavingsReturn interface {
	OrderBy(orders ...Order) iOrderByReturn
	Limit(limit int) iLimitReturn
	And() iAndReturn
	Or() iOrReturn
	EndGroup() iEndGroupReturn
	Gen() (sql string, err error)
	Val() []interface{}
}

// iOrderByReturn iOrderByReturn
type iOrderByReturn interface {
	Limit(limit int) iLimitReturn
	Gen() (sql string, err error)
	Val() []interface{}
}

// iLimitReturn iLimitReturn
type iLimitReturn interface {
	Offset(offset int) iOffsetReturn
	Gen() (sql string, err error)
	Val() []interface{}
}

// iOffsetReturn iOffsetReturn
type iOffsetReturn interface {
	Gen() (sql string, err error)
	Val() []interface{}
}

// iAndReturn iAndReturn
type iAndReturn interface {
	Where(where Where) iWheresReturn
	Having(having Having) iHavingsReturn
	StartGroup() iStartGroupReturn
	On(on On) iOnReturn
}

// iOrReturn iOrReturn
type iOrReturn interface {
	iAndReturn
}

// iStartGroupReturn iStartGroupReturn
type iStartGroupReturn interface {
	Where(where Where) iWheresReturn
	On(on On) iOnReturn
	Having(having Having) iHavingsReturn
}

// iEndGroupReturn iEndGroupReturn
type iEndGroupReturn interface {
	RightJoin(table Table) iJoinReturn
	LeftJoin(table Table) iJoinReturn
	Where(where Where) iWheresReturn
	GroupBy(group ...string) iGroupByReturn
	OrderBy(orders ...Order) iOrderByReturn
	Limit(limit int) iLimitReturn
	And() iAndReturn
	Or() iOrReturn
	Gen() (sql string, err error)
	Val() []interface{}
}

// iBuildInsertReturn iBuildInsertReturn
type iBuildInsertReturn interface {
	Table(table Table) iITableReturn
}

// iITableReturn iITableReturn
type iITableReturn interface {
	Values(column ...Column) iIValuesReturn
}

// iIValuesReturn iIValuesReturn
type iIValuesReturn interface {
	Gen() (sql string, err error)
	Val() []interface{}
}

// iBuildDeleteReturn iBuildDeleteReturn
type iBuildDeleteReturn interface {
	Table(table Table) iDTableReturn
}

// iDTableReturn iDTableReturn
type iDTableReturn interface {
	Where(where Where) iDWheresReturn
}

// iDWheresReturn iDWheresReturn
type iDWheresReturn interface {
	And() iDAndReturn
	Or() iDOrReturn
	EndGroup() iDEndGroup
	Gen() (sql string, err error)
	Val() []interface{}
}

// iDAndReturn iDAndReturn
type iDAndReturn interface {
	Where(where Where) iDWheresReturn
	StartGroup() iDStartGroup
}

// iDOrReturn iDOrReturn
type iDOrReturn interface {
	iDAndReturn
}

// iDStartGroup iDStartGroup
type iDStartGroup interface {
	Where(where Where) iDWheresReturn
}

// iDEndGroup iDEndGroup
type iDEndGroup interface {
	Where(where Where) iDWheresReturn
	Gen() (sql string, err error)
	Val() []interface{}
}

// iBuildUpdateReturn iBuildUpdateReturn
type iBuildUpdateReturn interface {
	Table(table Table) iUTableReturn
}

// iUSetReturn iUSetReturn
type iUSetReturn interface {
	Where(where Where) iDWheresReturn
	Gen() (sql string, err error)
	Val() []interface{}
}

// iUTableReturn iUTableReturn
type iUTableReturn interface {
	Set(columns ...Column) iUSetReturn
}
