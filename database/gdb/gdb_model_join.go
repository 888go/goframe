// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb

import (
	"fmt"
	
	"github.com/888go/goframe/text/gstr"
)

// LeftJoin 对模型执行 "LEFT JOIN ... ON ..." 语句。
// 参数 `table` 可以是联接表及其联接条件，并且可以带有别名名称。
//
// 示例：
// Model("user").LeftJoin("user_detail", "user_detail.uid=user.uid") // 用户表与用户详情表进行左连接，关联条件为 user_detail.uid 等于 user.uid
// Model("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid") // 使用别名，用户表（别名u）与用户详情表（别名ud）进行左连接，关联条件为 ud.uid 等于 u.uid
// Model("user", "u").LeftJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid") // 将子查询结果作为联接表，并使用别名，用户表（别名u）与子查询结果（别名a）进行左连接，关联条件为 a.uid 等于 u.uid
func (m *Model) LeftJoin(tableOrSubQueryAndJoinConditions ...string) *Model {
	return m.doJoin(joinOperatorLeft, tableOrSubQueryAndJoinConditions...)
}

// RightJoin 对模型执行 "RIGHT JOIN ... ON ..." 语句。
// 参数 `table` 可以是待连接的表及其连接条件，并且可以包含别名名称。
//
// 示例：
// Model("user").RightJoin("user_detail", "user_detail.uid=user.uid") // 连接 user 表和 user_detail 表，连接条件为 user_detail.uid 等于 user.uid
// Model("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")   // 使用别名 u 和 ud 连接 user 表和 user_detail 表，连接条件为 ud.uid 等于 u.uid
// Model("user", "u").RightJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid") // 将查询结果作为连接表，使用别名 u 和 a 连接 user 表和查询结果表，连接条件为 a.uid 等于 u.uid
func (m *Model) RightJoin(tableOrSubQueryAndJoinConditions ...string) *Model {
	return m.doJoin(joinOperatorRight, tableOrSubQueryAndJoinConditions...)
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter `table` can be joined table and its joined condition,
// and also with its alias name。
//
// Eg:
// Model("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Model("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
// Model("user", "u").InnerJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid").
func (m *Model) InnerJoin(tableOrSubQueryAndJoinConditions ...string) *Model {
	return m.doJoin(joinOperatorInner, tableOrSubQueryAndJoinConditions...)
}

// LeftJoinOnField 执行类似 LeftJoin 的操作，但是它通过“相同字段名”将两个表连接起来。
//
// 例如：
// Model("order").LeftJoinOnField("user", "user_id") // 使用 user 表与 order 表的 user_id 字段进行左连接
// Model("order").LeftJoinOnField("product", "product_id") // 使用 product 表与 order 表的 product_id 字段进行左连接
func (m *Model) LeftJoinOnField(table, field string) *Model {
	return m.doJoin(joinOperatorLeft, table, fmt.Sprintf(
		`%s.%s=%s.%s`,
		m.tablesInit,
		m.db.GetCore().QuoteWord(field),
		m.db.GetCore().QuoteWord(table),
		m.db.GetCore().QuoteWord(field),
	))
}

// RightJoinOnField 执行 RightJoin 操作，但是它通过“相同字段名”将两个表连接起来。
//
// 例如：
// Model("order").InnerJoinOnField("user", "user_id") // 根据 user_id 字段将 order 表与 user 表进行内连接
// Model("order").InnerJoinOnField("product", "product_id") // 根据 product_id 字段将 order 表与 product 表进行内连接
// （注意：代码中的 InnerJoinOnField 应该是 RightJoinOnField，因为注释中提到的是 RightJoin。若确实为 RightJoin，请将上述翻译中的“内连接”替换为“右连接”。）
// 修正后的翻译：
// RightJoinOnField 执行 RightJoin 操作，但它是基于“相同字段名”将两个表进行连接。
//
// 例如：
// Model("order").RightJoinOnField("user", "user_id") // 根据 user_id 字段将 order 表与 user 表进行右连接
// Model("order").RightJoinOnField("product", "product_id") // 根据 product_id 字段将 order 表与 product 表进行右连接
func (m *Model) RightJoinOnField(table, field string) *Model {
	return m.doJoin(joinOperatorRight, table, fmt.Sprintf(
		`%s.%s=%s.%s`,
		m.tablesInit,
		m.db.GetCore().QuoteWord(field),
		m.db.GetCore().QuoteWord(table),
		m.db.GetCore().QuoteWord(field),
	))
}

// InnerJoinOnField 执行内连接操作，但它通过`相同的字段名`将两个表连接起来。
//
// 例如：
// Model("order").InnerJoinOnField("user", "user_id") // 使用user表与order表中user_id字段进行内连接
// Model("order").InnerJoinOnField("product", "product_id") // 使用product表与order表中product_id字段进行内连接
func (m *Model) InnerJoinOnField(table, field string) *Model {
	return m.doJoin(joinOperatorInner, table, fmt.Sprintf(
		`%s.%s=%s.%s`,
		m.tablesInit,
		m.db.GetCore().QuoteWord(field),
		m.db.GetCore().QuoteWord(table),
		m.db.GetCore().QuoteWord(field),
	))
}

// LeftJoinOnFields 执行左连接操作。它用于指定不同的字段以及比较运算符。
//
// 例如：
// Model("user").LeftJoinOnFields("order", "id", "=", "user_id") // 根据 id 等于 user_id 进行左连接
// Model("user").LeftJoinOnFields("order", "id", ">", "user_id") // 根据 id 大于 user_id 进行左连接
// Model("user").LeftJoinOnFields("order", "id", "<", "user_id") // 根据 id 小于 user_id 进行左连接
// 这段 Go 代码的注释翻译成中文注释如上，该函数主要是对两个数据表进行左连接操作，并允许用户自定义连接条件中的字段和比较运算符。
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

// RightJoinOnFields 执行右连接操作。它用于指定不同的字段以及比较运算符。
//
// 示例：
// Model("user").RightJoinOnFields("order", "id", "=", "user_id") // 用户表通过id字段等于order表中的user_id字段进行右连接
// Model("user").RightJoinOnFields("order", "id", ">", "user_id") // 用户表通过id字段大于order表中的user_id字段进行右连接
// Model("user").RightJoinOnFields("order", "id", "<", "user_id") // 用户表通过id字段小于order表中的user_id字段进行右连接
// 这段Go语言代码的注释翻译为中文注释，其功能是定义一个RightJoinOnFields方法，该方法用于执行SQL查询中的右连接操作，并允许用户自定义连接的字段及比较条件。
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

// InnerJoinOnFields 执行内连接操作。它用于指定不同的字段及比较运算符。
//
// 示例：
// Model("user").InnerJoinOnFields("order", "id", "=", "user_id")  // 根据 id 等于 user_id 进行内连接
// Model("user").InnerJoinOnFields("order", "id", ">", "user_id")  // 根据 id 大于 user_id 进行内连接
// Model("user").InnerJoinOnFields("order", "id", "<", "user_id")  // 根据 id 小于 user_id 进行内连接
// 该函数主要用于在Go语言中进行数据库查询时，通过自定义字段与运算符实现两个表之间的内连接操作。
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

// doJoin 对模型执行 "LEFT/RIGHT/INNER JOIN ... ON ..." 语句。
// 参数 `tableOrSubQueryAndJoinConditions` 可以是联接表及其联接条件，同时也可以包含别名名称。
//
// 示例：
// Model("user").InnerJoin("user_detail", "user_detail.uid=user.uid") // 对用户表进行内连接，条件为 user_detail.uid 等于 user.uid
// Model("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid") // 使用别名对用户表（u）进行内连接，连接到详情表（ud），条件为 ud.uid 等于 u.uid
// Model("user", "u").InnerJoin("user_detail", "ud", "ud.uid>u.uid") // 同上，但条件为 ud.uid 大于 u.uid
// Model("user", "u").InnerJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid") // 连接到子查询结果（a），条件为 a.uid 等于 u.uid
// 相关问题：
// https://github.com/gogf/gf/issues/1024
func (m *Model) doJoin(operator joinOperator, tableOrSubQueryAndJoinConditions ...string) *Model {
	var (
		model   = m.getModel()
		joinStr = ""
		table   string
		alias   string
	)
	// 检查第一个参数，表格或子查询。
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
	// 生成连接条件语句字符串。
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

// getTableNameByPrefixOrAlias 检查并返回表名，如果 `prefixOrAlias` 是某个表的别名，则返回该表名；
// 否则直接返回 `prefixOrAlias`。
func (m *Model) getTableNameByPrefixOrAlias(prefixOrAlias string) string {
	value, ok := m.tableAliasMap[prefixOrAlias]
	if ok {
		return value
	}
	return prefixOrAlias
}

// isSubQuery 检查并返回给定的字符串是否为子查询SQL字符串。
func isSubQuery(s string) bool {
	s = gstr.TrimLeft(s, "()")
	if p := gstr.Pos(s, " "); p != -1 {
		if gstr.Equal(s[:p], "select") {
			return true
		}
	}
	return false
}
