// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

import (
	"fmt"

	"github.com/gogf/gf/v2/text/gstr"
)

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter `table` can be joined table and its joined condition,
// and also with its alias name.
//
// Model("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Model("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
// Model("user", "u").LeftJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid").
// ff:左连接
// m:
// tableOrSubQueryAndJoinConditions:表或子查询和连接条件
func (m *Model) LeftJoin(tableOrSubQueryAndJoinConditions ...string) *Model {
	return m.doJoin(joinOperatorLeft, tableOrSubQueryAndJoinConditions...)
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter `table` can be joined table and its joined condition,
// and also with its alias name.
//
// Model("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Model("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
// Model("user", "u").RightJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid").
// ff:右连接
// m:
// tableOrSubQueryAndJoinConditions:表或子查询和连接条件
func (m *Model) RightJoin(tableOrSubQueryAndJoinConditions ...string) *Model {
	return m.doJoin(joinOperatorRight, tableOrSubQueryAndJoinConditions...)
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter `table` can be joined table and its joined condition,
// and also with its alias name。
//
// Model("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Model("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
// Model("user", "u").InnerJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid").
// ff:内连接
// m:
// tableOrSubQueryAndJoinConditions:表或子查询和连接条件
func (m *Model) InnerJoin(tableOrSubQueryAndJoinConditions ...string) *Model {
	return m.doJoin(joinOperatorInner, tableOrSubQueryAndJoinConditions...)
}

// LeftJoinOnField performs as LeftJoin, but it joins both tables with the `same field name`.
//
// Model("order").LeftJoinOnField("user", "user_id")
// Model("order").LeftJoinOnField("product", "product_id").
// ff:左连接相同字段
// m:
// table:表名
// field:相同字段名
func (m *Model) LeftJoinOnField(table, field string) *Model {
	return m.doJoin(joinOperatorLeft, table, fmt.Sprintf(
		`%s.%s=%s.%s`,
		m.tablesInit,
		m.db.GetCore().QuoteWord(field),
		m.db.GetCore().QuoteWord(table),
		m.db.GetCore().QuoteWord(field),
	))
}

// RightJoinOnField performs as RightJoin, but it joins both tables with the `same field name`.
//
// Model("order").InnerJoinOnField("user", "user_id")
// Model("order").InnerJoinOnField("product", "product_id").
// ff:右连接相同字段
// m:
// table:表名
// field:相同字段名
func (m *Model) RightJoinOnField(table, field string) *Model {
	return m.doJoin(joinOperatorRight, table, fmt.Sprintf(
		`%s.%s=%s.%s`,
		m.tablesInit,
		m.db.GetCore().QuoteWord(field),
		m.db.GetCore().QuoteWord(table),
		m.db.GetCore().QuoteWord(field),
	))
}

// InnerJoinOnField performs as InnerJoin, but it joins both tables with the `same field name`.
//
// Model("order").InnerJoinOnField("user", "user_id")
// Model("order").InnerJoinOnField("product", "product_id").
// ff:内连接相同字段
// m:
// table:表名
// field:相同字段名
func (m *Model) InnerJoinOnField(table, field string) *Model {
	return m.doJoin(joinOperatorInner, table, fmt.Sprintf(
		`%s.%s=%s.%s`,
		m.tablesInit,
		m.db.GetCore().QuoteWord(field),
		m.db.GetCore().QuoteWord(table),
		m.db.GetCore().QuoteWord(field),
	))
}

// LeftJoinOnFields performs as LeftJoin. It specifies different fields and comparison operator.
//
// Model("user").LeftJoinOnFields("order", "id", "=", "user_id")
// Model("user").LeftJoinOnFields("order", "id", ">", "user_id")
// Model("user").LeftJoinOnFields("order", "id", "<", "user_id")
// ff:左连接带比较运算符
// m:
// table:表名
// firstField:第一个字段
// operator:比较运算符
// secondField:第二个字段
func (m *Model) LeftJoinOnFields(table, firstField, operator, secondField string) *Model {
	return m.doJoin(joinOperatorLeft, table, fmt.Sprintf(
		`%s.%s %s %s.%s`,
		m.tablesInit,
		m.db.GetCore().QuoteWord(firstField),
		operator,
		m.db.GetCore().QuoteWord(table),
		m.db.GetCore().QuoteWord(secondField),
	))
}

// RightJoinOnFields performs as RightJoin. It specifies different fields and comparison operator.
//
// Model("user").RightJoinOnFields("order", "id", "=", "user_id")
// Model("user").RightJoinOnFields("order", "id", ">", "user_id")
// Model("user").RightJoinOnFields("order", "id", "<", "user_id")
// ff:右连接带比较运算符
// m:
// table:表名
// firstField:第一个字段
// operator:比较运算符
// secondField:第二个字段
func (m *Model) RightJoinOnFields(table, firstField, operator, secondField string) *Model {
	return m.doJoin(joinOperatorRight, table, fmt.Sprintf(
		`%s.%s %s %s.%s`,
		m.tablesInit,
		m.db.GetCore().QuoteWord(firstField),
		operator,
		m.db.GetCore().QuoteWord(table),
		m.db.GetCore().QuoteWord(secondField),
	))
}

// InnerJoinOnFields performs as InnerJoin. It specifies different fields and comparison operator.
//
// Model("user").InnerJoinOnFields("order", "id", "=", "user_id")
// Model("user").InnerJoinOnFields("order", "id", ">", "user_id")
// Model("user").InnerJoinOnFields("order", "id", "<", "user_id")
// ff:内连接带比较运算符
// m:
// table:表名
// firstField:第一个字段
// operator:比较运算符
// secondField:第二个字段
func (m *Model) InnerJoinOnFields(table, firstField, operator, secondField string) *Model {
	return m.doJoin(joinOperatorInner, table, fmt.Sprintf(
		`%s.%s %s %s.%s`,
		m.tablesInit,
		m.db.GetCore().QuoteWord(firstField),
		operator,
		m.db.GetCore().QuoteWord(table),
		m.db.GetCore().QuoteWord(secondField),
	))
}

// doJoin does "LEFT/RIGHT/INNER JOIN ... ON ..." statement on the model.
// The parameter `tableOrSubQueryAndJoinConditions` can be joined table and its joined condition,
// and also with its alias name.
//
// Model("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Model("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
// Model("user", "u").InnerJoin("user_detail", "ud", "ud.uid>u.uid")
// Model("user", "u").InnerJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid")
// Related issues:
func (m *Model) doJoin(operator joinOperator, tableOrSubQueryAndJoinConditions ...string) *Model {
	var (
		model   = m.getModel()
		joinStr = ""
		table   string
		alias   string
	)
	// 检查第一个参数，是否为表格或子查询。 md5:0493998c4b03304e
	if len(tableOrSubQueryAndJoinConditions) > 0 {
		if isSubQuery(tableOrSubQueryAndJoinConditions[0]) {
			joinStr = gstr.Trim(tableOrSubQueryAndJoinConditions[0])
			if joinStr[0] != '(' {
				joinStr = "(" + joinStr + ")"
			}
		} else {
			table = tableOrSubQueryAndJoinConditions[0]
			joinStr = m.db.GetCore().QuotePrefixTableName(table)
		}
	}
	// 生成连接条件的字符串表达式。 md5:54f67a1d882ecd10
	conditionLength := len(tableOrSubQueryAndJoinConditions)
	switch {
	case conditionLength > 2:
		alias = tableOrSubQueryAndJoinConditions[1]
		model.tables += fmt.Sprintf(
			" %s JOIN %s AS %s ON (%s)",
			operator, joinStr,
			m.db.GetCore().QuoteWord(alias),
			tableOrSubQueryAndJoinConditions[2],
		)
		m.tableAliasMap[alias] = table

	case conditionLength == 2:
		model.tables += fmt.Sprintf(
			" %s JOIN %s ON (%s)",
			operator, joinStr, tableOrSubQueryAndJoinConditions[1],
		)

	case conditionLength == 1:
		model.tables += fmt.Sprintf(
			" %s JOIN %s", operator, joinStr,
		)
	}
	return model
}

// getTableNameByPrefixOrAlias 检查`prefixOrAlias`是否是某个表的别名，如果是，则返回该表的实际名称，否则直接返回`prefixOrAlias`。
// md5:ab423b9e1e0ad0ca
func (m *Model) getTableNameByPrefixOrAlias(prefixOrAlias string) string {
	value, ok := m.tableAliasMap[prefixOrAlias]
	if ok {
		return value
	}
	return prefixOrAlias
}

// isSubQuery 检查并返回给定的字符串是否为子查询SQL语句。 md5:0921761c51f20650
func isSubQuery(s string) bool {
	s = gstr.TrimLeft(s, "()")
	if p := gstr.Pos(s, " "); p != -1 {
		if gstr.Equal(s[:p], "select") {
			return true
		}
	}
	return false
}
