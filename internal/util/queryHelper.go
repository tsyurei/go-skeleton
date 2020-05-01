package util

import "github.com/go-pg/pg/v9/orm"

import "fmt"

import "strings"

// QueryHelper will create a temporary storage for query before running it
type QueryHelper struct {
	Query *orm.Query

	AliasMap map[QueryAliasHelper]string
}

type QueryAliasHelper struct {
	tableName, aliasName string
}

func InitQueryHelper(query *orm.Query) *QueryHelper {
	return &QueryHelper{
		Query:  query,
		AliasMap: make(map[QueryAliasHelper]string, 4),
	}
}

func (qh *QueryHelper) Join(joinType, tableName, aliasName, joinOn string) *orm.Query {
	ah := QueryAliasHelper{tableName, aliasName}
	if _, ok := qh.AliasMap[ah]; !ok {
		qh.AliasMap[ah] = ah.aliasName
		qh.Query = qh.Query.Join(fmt.Sprintf("%s JOIN %s AS %s ON %s", joinType, tableName, aliasName, joinOn))
	}
	return qh.Query
}

func (qh *QueryHelper) Relation(tableName string, apply ...func(*orm.Query) (*orm.Query, error)) *orm.Query {
	ah := QueryAliasHelper{strings.ToLower(tableName), strings.ToLower(tableName)}
	if _, ok := qh.AliasMap[ah]; !ok {
		qh.AliasMap[ah] = ah.aliasName
		qh.Query = qh.Query.Relation(tableName, apply...)
	}
	return qh.Query
}

