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

// LeftJoin 在模型上执行 "LEFT JOIN ... ON ..." 语句。
// 参数 `table` 可以是连接的表及其连接条件，也可以包含其别名。
//
// 示例：
// Model("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Model("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
// Model("user", "u").LeftJoin("SELECT xxx FROM xxx", "a", "a.uid=u.uid")
// md5:5f7464280da64004
func (m *Model) LeftJoin(tableOrSubQueryAndJoinConditions ...string) *Model {
	return m.doJoin(joinOperatorLeft, tableOrSubQueryAndJoinConditions...)
}

// RightJoin 执行 "RIGHT JOIN ... ON ..." 语句在模型上。
// 参数 `table` 可以是待连接的表及其连接条件，
// 也可以包含表的别名。
//
// 例如：
// Model("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Model("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
// Model("user", "u").RightJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid")
// md5:dbab2528fb37c84e
func (m *Model) RightJoin(tableOrSubQueryAndJoinConditions ...string) *Model {
	return m.doJoin(joinOperatorRight, tableOrSubQueryAndJoinConditions...)
}

// InnerJoin 在模型上执行 "INNER JOIN ... ON ..." 语句。
// 参数 `table` 可以是需要连接的表及其连接条件，同时也可包含别名名称。
//
// 例如：
// Model("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Model("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
// Model("user", "u").InnerJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid")
// md5:eda419ad685c559d
func (m *Model) InnerJoin(tableOrSubQueryAndJoinConditions ...string) *Model {
	return m.doJoin(joinOperatorInner, tableOrSubQueryAndJoinConditions...)
}

// LeftJoinOnField 执行左连接，但使用两个表中的`相同字段名`进行连接。
//
// 例如：
// Model("order").LeftJoinOnField("user", "user_id")
// Model("order").LeftJoinOnField("product", "product_id")
// md5:5ace5bdef45c73d6
func (m *Model) LeftJoinOnField(table, field string) *Model {
	return m.doJoin(joinOperatorLeft, table, fmt.Sprintf(
		`%s.%s=%s.%s`,
		m.tablesInit,
		m.db.GetCore().QuoteWord(field),
		m.db.GetCore().QuoteWord(table),
		m.db.GetCore().QuoteWord(field),
	))
}

// RightJoinOnField 执行右连接，但使用相同的字段名称连接两个表。
//
// 例如：
// Model("order").InnerJoinOnField("user", "user_id")
// Model("order").InnerJoinOnField("product", "product_id")
// md5:ac8281e2d383e3d6
func (m *Model) RightJoinOnField(table, field string) *Model {
	return m.doJoin(joinOperatorRight, table, fmt.Sprintf(
		`%s.%s=%s.%s`,
		m.tablesInit,
		m.db.GetCore().QuoteWord(field),
		m.db.GetCore().QuoteWord(table),
		m.db.GetCore().QuoteWord(field),
	))
}

// InnerJoinOnField 的行为类似于 InnerJoin，但它使用的是具有`相同字段名`的两个表进行连接。
//
// 例如：
// Model("order").InnerJoinOnField("user", "user_id")
// Model("order").InnerJoinOnField("product", "product_id")
// md5:bdc954b5bcb8a9c5
func (m *Model) InnerJoinOnField(table, field string) *Model {
	return m.doJoin(joinOperatorInner, table, fmt.Sprintf(
		`%s.%s=%s.%s`,
		m.tablesInit,
		m.db.GetCore().QuoteWord(field),
		m.db.GetCore().QuoteWord(table),
		m.db.GetCore().QuoteWord(field),
	))
}

// LeftJoinOnFields 执行类似于 LeftJoin 的操作，但允许指定不同的字段和比较运算符。
//
// 例如：
// Model("user").LeftJoinOnFields("order", "id", "=", "user_id")
// Model("user").LeftJoinOnFields("order", "id", ">", "user_id")
// Model("user").LeftJoinOnFields("order", "id", "<", "user_id")
// md5:90ce0e2226eb4b30
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

// RightJoinOnFields 执行右连接操作。它指定了不同的字段和比较运算符。
//
// 例如：
// User("user").RightJoinOnFields("order", "id", "=", "user_id")
// User("user").RightJoinOnFields("order", "id", ">", "user_id")
// User("user").RightJoinOnFields("order", "id", "<", "user_id") 
// 
// 这里，`RightJoinOnFields` 是一个 Go 代码中的函数，用于在查询数据库时执行右连接操作，并且允许用户自定义连接的字段和比较操作。第一个参数是模型名（如 "user"），接下来的参数包括要连接的表名、字段名以及连接条件（等于、大于或小于）。
// md5:563f2b0f155fc829
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

// InnerJoinOnFields 执行 InnerJoin 操作。它指定了不同的字段和比较运算符。
// 
// 例如：
// Model("user").InnerJoinOnFields("order", "id", "=", "user_id")
// Model("user").InnerJoinOnFields("order", "id", ">", "user_id")
// Model("user").InnerJoinOnFields("order", "id", "<", "user_id") 
// 
// 这段代码是在 Go 语言中定义了一个方法，用于在两个数据表之间执行内连接（InnerJoin），并允许用户指定连接的字段以及比较运算符。例如，`"user".InnerJoinOnFields("order", "id", "=", "user_id")` 表示连接 "user" 表和 "order" 表，通过 "id" 字段进行等号（=）匹配。其他示例展示了使用大于（>）和小于（<）运算符的情况。
// md5:0499f4b5bbbc2016
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

// doJoin 在模型上执行 "LEFT/RIGHT/INNER JOIN ... ON ..." 语句。
// 参数 `tableOrSubQueryAndJoinConditions` 可以是待连接的表及其连接条件，
// 同时也可以包含表的别名。
//
// 例如：
// Model("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Model("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
// Model("user", "u").InnerJoin("user_detail", "ud", "ud.uid>u.uid")
// Model("user", "u").InnerJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid")
// 相关问题讨论：
// https://github.com/gogf/gf/issues/1024
// md5:7b792ce187933a04
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
