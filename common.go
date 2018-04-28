package easydb

var no_value = "no value"
var db_mysql = "db_mysql"
var db_pgsql = "db_pgsql"

// DBType DBType
type DBType int

const (
	// MYSQL mysql
	MYSQL DBType = iota
	// PGSQL postgresql
	PGSQL
)

// DBOptType DBOptType
type DBOptType string

const (
	// Insert Insert
	Insert DBOptType = "INSERT INTO "
	// Delete Delete
	Delete DBOptType = "DELETE FROM "
	// Update Update
	Update DBOptType = "UPDATE "
	// Select Select
	Select DBOptType = "SELECT "
)

// QueryType QueryType
type QueryType string

const (
	// COUNT select count
	COUNT QueryType = " COUNT"
	// SUM select sum
	SUM QueryType = " SUM"
	// MAX select max
	MAX QueryType = " MAX"
	// MIN select min
	MIN QueryType = " MIN"
	// AVG select avg
	AVG QueryType = " AVG"
)

// LogicalOptType LogicalOptType
type LogicalOptType string

const (
	// EQ =
	EQ LogicalOptType = " = "
	// GT >
	GT LogicalOptType = " > "
	// LT <
	LT LogicalOptType = " < "
	// GE >=
	GE LogicalOptType = " >= "
	// LE <=
	LE LogicalOptType = " <= "
	// NE <>
	NE LogicalOptType = " <> "
	// LIKE like
	LIKE LogicalOptType = " LIKE "
	// IN in
	IN LogicalOptType = " IN "
	// IsNULL is null
	IsNULL LogicalOptType = " IS NULL "
	// NotNULL is not null
	NotNULL LogicalOptType = " IS NOT NULL "
)

// WhereGroupType WhereGroupType
type WhereGroupType string

const (
	// AND AND
	AND WhereGroupType = " AND "
	// OR OR
	OR WhereGroupType = " OR "
	// GroupStart (
	GroupStart WhereGroupType = " ( "
	// GroupEnd )
	GroupEnd WhereGroupType = " ) "
)

// JoinType JoinType
type JoinType string

const (
	// LeftJoin LEFT JOIN
	LeftJoin JoinType = " LEFT JOIN "
	// RightJoin RIGHT JOIN
	RightJoin JoinType = " RIGHT JOIN "
)

// OrderType OrderType
type OrderType string

const (
	// DESC DESC
	DESC OrderType = " DESC "
	// ASC ASC
	ASC OrderType = " ASC "
)
