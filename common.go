package easydb

var no_value = "no value"
var db_mysql = "db_mysql"
var db_pgsql = "db_pgsql"

// dbType dbType
type dbType int

const (
	// MYSQL mysql
	MYSQL dbType = iota
	// PGSQL postgresql
	PGSQL
)

// dbOptType dbOptType
type dbOptType string

const (
	// Insert Insert
	Insert dbOptType = "INSERT INTO "
	// Delete Delete
	Delete dbOptType = "DELETE FROM "
	// Update Update
	Update dbOptType = "UPDATE "
	// Select Select
	Select dbOptType = "SELECT "
)

// queryType queryType
type queryType string

const (
	// COUNT select count
	COUNT queryType = " COUNT"
	// SUM select sum
	SUM queryType = " SUM"
	// MAX select max
	MAX queryType = " MAX"
	// MIN select min
	MIN queryType = " MIN"
	// AVG select avg
	AVG queryType = " AVG"
)

// logicalOptType logicalOptType
type logicalOptType string

const (
	// EQ =
	EQ logicalOptType = " = "
	// GT >
	GT logicalOptType = " > "
	// LT <
	LT logicalOptType = " < "
	// GE >=
	GE logicalOptType = " >= "
	// LE <=
	LE logicalOptType = " <= "
	// NE <>
	NE logicalOptType = " <> "
	// LIKE like
	LIKE logicalOptType = " LIKE "
	// IN in
	IN logicalOptType = " IN "
	// IsNULL is null
	IsNULL logicalOptType = " IS NULL "
	// NotNULL is not null
	NotNULL logicalOptType = " IS NOT NULL "
)

// whereGroupType whereGroupType
type whereGroupType string

const (
	// AND AND
	AND whereGroupType = " AND "
	// OR OR
	OR whereGroupType = " OR "
	// GroupStart (
	GroupStart whereGroupType = " ( "
	// GroupEnd )
	GroupEnd whereGroupType = " ) "
)

// JoinType JoinType
type JoinType string

const (
	// LeftJoin LEFT JOIN
	LeftJoin JoinType = " LEFT JOIN "
	// RightJoin RIGHT JOIN
	RightJoin JoinType = " RIGHT JOIN "
)

// orderType orderType
type orderType string

const (
	// DESC DESC
	DESC orderType = " DESC "
	// ASC ASC
	ASC orderType = " ASC "
)
