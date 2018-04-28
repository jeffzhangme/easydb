package restrict

import "github.com/jeffzhangme/easydb"

// IBuildQueryReturn IBuildQueryReturn
type IBuildQueryReturn interface {
	QueryFuncs(funccs ...easydb.QueryFunc) IQFuncReturn
	Columns(columns ...easydb.Column) IColumnsReturn
}

// IQFuncReturn IQFuncReturn
type IQFuncReturn interface {
	QueryFuncs(funccs ...easydb.QueryFunc) IQFuncReturn
	Columns(columns ...easydb.Column) IColumnsReturn
	Tables(tables ...easydb.Table) IFromsReturn
}

// IColumnsReturn IColumnsReturn
type IColumnsReturn interface {
	IQFuncReturn
}

// IFromsReturn IFromsReturn
type IFromsReturn interface {
	Tables(tables ...easydb.Table) IFromsReturn
	RightJoin(table easydb.Table) IJoinReturn
	LeftJoin(table easydb.Table) IJoinReturn
	Where(where easydb.Where) IWheresReturn
	GroupBy(group ...string) IGroupByReturn
	OrderBy(orders ...easydb.Order) IOrderByReturn
	Limit(limit int) ILimitReturn
	Gen() (sql string, err error)
	Val() []string
}

// IJoinReturn IJoinReturn
type IJoinReturn interface {
	On(on easydb.On) IOnReturn
}

// IOnReturn IOnReturn
type IOnReturn interface {
	RightJoin(table easydb.Table) IJoinReturn
	LeftJoin(table easydb.Table) IJoinReturn
	Where(where easydb.Where) IWheresReturn
	GroupBy(group ...string) IGroupByReturn
	OrderBy(orders ...easydb.Order) IOrderByReturn
	Limit(limit int) ILimitReturn
	Gen() (sql string, err error)
	Val() []string
	And() IAndReturn
	Or() IOrReturn
}

// IWheresReturn IWheresReturn
type IWheresReturn interface {
	And() IAndReturn
	Or() IOrReturn
	GroupBy(group ...string) IGroupByReturn
	OrderBy(orders ...easydb.Order) IOrderByReturn
	Limit(limit int) ILimitReturn
	Gen() (sql string, err error)
	Val() []string
}

// IGroupByReturn 制约groupby条件 后调用
type IGroupByReturn interface {
	Having(having easydb.Having) IHavingsReturn
	OrderBy(orders ...easydb.Order) IOrderByReturn
	Limit(limit int) ILimitReturn
	Gen() (sql string, err error)
	Val() []string
}

// IHavingsReturn IHavingsReturn
type IHavingsReturn interface {
	OrderBy(orders ...easydb.Order) IOrderByReturn
	Limit(limit int) ILimitReturn
	And() IAndReturn
	Or() IOrReturn
	Gen() (sql string, err error)
	Val() []string
}

// IOrderByReturn IOrderByReturn
type IOrderByReturn interface {
	Limit(limit int) ILimitReturn
	Gen() (sql string, err error)
	Val() []string
}

// ILimitReturn ILimitReturn
type ILimitReturn interface {
	Offset(offset int) IOffsetReturn
	Gen() (sql string, err error)
	Val() []string
}

// IOffsetReturn IOffsetReturn
type IOffsetReturn interface {
	Gen() (sql string, err error)
	Val() []string
}

// IAndReturn IAndReturn
type IAndReturn interface {
	Where(where easydb.Where) IWheresReturn
	Having(having easydb.Having) IHavingsReturn
	StartGroup() IStartGroupReturn
	EndGroup() IEndGroupReturn
	On(on easydb.On) IOnReturn
}

// IOrReturn IOrReturn
type IOrReturn interface {
	IAndReturn
}

// IStartGroupReturn IStartGroupReturn
type IStartGroupReturn interface {
	Where(where easydb.Where) IWheresReturn
	On(on easydb.On) IOnReturn
	Having(having easydb.Having) IHavingsReturn
}

// IEndGroupReturn IEndGroupReturn
type IEndGroupReturn interface {
	RightJoin(table easydb.Table) IJoinReturn
	LeftJoin(table easydb.Table) IJoinReturn
	Where(where easydb.Where) IWheresReturn
	GroupBy(group ...string) IGroupByReturn
	OrderBy(orders ...easydb.Order) IOrderByReturn
	Limit(limit int) ILimitReturn
	And() IAndReturn
	Or() IOrReturn
	Gen() (sql string, err error)
	Val() []string
}

// IBuildInsertReturn IBuildInsertReturn
type IBuildInsertReturn interface {
	Table(table easydb.Table) IITableReturn
}

// IITableReturn IITableReturn
type IITableReturn interface {
	Values(column ...easydb.Column) IIValuesReturn
}

// IIValuesReturn IIValuesReturn
type IIValuesReturn interface {
	Gen() (sql string, err error)
	Val() []string
}

// IBuildDeleteReturn IBuildDeleteReturn
type IBuildDeleteReturn interface {
	Table(table easydb.Table) IDTableReturn
}

// IDTableReturn IDTableReturn
type IDTableReturn interface {
	Where(where easydb.Where) IDWheresReturn
}

// IDWheresReturn IDWheresReturn
type IDWheresReturn interface {
	And() IDAndReturn
	Or() IDOrReturn
	EndGroup() IDEndGroup
	Gen() (sql string, err error)
	Val() []string
}

// IDAndReturn IDAndReturn
type IDAndReturn interface {
	Where(where easydb.Where) IDWheresReturn
	StartGroup() IDStartGroup
}

// IDOrReturn IDOrReturn
type IDOrReturn interface {
	IDAndReturn
}

// IDStartGroup IDStartGroup
type IDStartGroup interface {
	Where(where easydb.Where) IDWheresReturn
}

// IDEndGroup IDEndGroup
type IDEndGroup interface {
	Where(where easydb.Where) IDWheresReturn
	Gen() (sql string, err error)
	Val() []string
}

// IBuildUpdateReturn IBuildUpdateReturn
type IBuildUpdateReturn interface {
	Table(table easydb.Table) IUTableReturn
}

// IUSetReturn IUSetReturn
type IUSetReturn interface {
	Where(where easydb.Where) IDWheresReturn
	Gen() (sql string, err error)
	Val() []string
}

// IUTableReturn IUTableReturn
type IUTableReturn interface {
	Set(columns ...easydb.Column) IUSetReturn
}
