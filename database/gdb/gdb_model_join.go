// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"fmt"

	gstr "github.com/888go/goframe/text/gstr"
)

// X左连接 在模型上执行 "LEFT JOIN ... ON ..." 语句。
// 参数 `table` 可以是连接的表及其连接条件，也可以包含其别名。
//
// 示例：
// Model("user").X左连接("user_detail", "user_detail.uid=user.uid")
// Model("user", "u").X左连接("user_detail", "ud", "ud.uid=u.uid")
// Model("user", "u").X左连接("SELECT xxx FROM xxx", "a", "a.uid=u.uid")
// md5:5f7464280da64004
func (m *Model) X左连接(表或子查询和连接条件 ...string) *Model {
	return m.doJoin(joinOperatorLeft, 表或子查询和连接条件...)
}

// X右连接 执行 "RIGHT JOIN ... ON ..." 语句在模型上。
// 参数 `table` 可以是待连接的表及其连接条件，
// 也可以包含表的别名。
//
// 例如：
// Model("user").X右连接("user_detail", "user_detail.uid=user.uid")
// Model("user", "u").X右连接("user_detail", "ud", "ud.uid=u.uid")
// Model("user", "u").X右连接("SELECT xxx FROM xxx","a", "a.uid=u.uid")
// md5:dbab2528fb37c84e
func (m *Model) X右连接(表或子查询和连接条件 ...string) *Model {
	return m.doJoin(joinOperatorRight, 表或子查询和连接条件...)
}

// X内连接 在模型上执行 "INNER JOIN ... ON ..." 语句。
// 参数 `table` 可以是需要连接的表及其连接条件，同时也可包含别名名称。
//
// 例如：
// Model("user").X内连接("user_detail", "user_detail.uid=user.uid")
// Model("user", "u").X内连接("user_detail", "ud", "ud.uid=u.uid")
// Model("user", "u").X内连接("SELECT xxx FROM xxx","a", "a.uid=u.uid")
// md5:eda419ad685c559d
func (m *Model) X内连接(表或子查询和连接条件 ...string) *Model {
	return m.doJoin(joinOperatorInner, 表或子查询和连接条件...)
}

// X左连接相同字段 执行左连接，但使用两个表中的`相同字段名`进行连接。
//
// 例如：
// Model("order").X左连接相同字段("user", "user_id")
// Model("order").X左连接相同字段("product", "product_id")
// md5:5ace5bdef45c73d6
func (m *Model) X左连接相同字段(表名, 相同字段名 string) *Model {
	return m.doJoin(joinOperatorLeft, 表名, fmt.Sprintf(
		`%s.%s=%s.%s`,
		m.tablesInit,
		m.db.X取Core对象().X底层QuoteWord(相同字段名),
		m.db.X取Core对象().X底层QuoteWord(表名),
		m.db.X取Core对象().X底层QuoteWord(相同字段名),
	))
}

// X右连接相同字段 执行右连接，但使用相同的字段名称连接两个表。
//
// 例如：
// Model("order").InnerJoinOnField("user", "user_id")
// Model("order").InnerJoinOnField("product", "product_id")
// md5:ac8281e2d383e3d6
func (m *Model) X右连接相同字段(表名, 相同字段名 string) *Model {
	return m.doJoin(joinOperatorRight, 表名, fmt.Sprintf(
		`%s.%s=%s.%s`,
		m.tablesInit,
		m.db.X取Core对象().X底层QuoteWord(相同字段名),
		m.db.X取Core对象().X底层QuoteWord(表名),
		m.db.X取Core对象().X底层QuoteWord(相同字段名),
	))
}

// X内连接相同字段 的行为类似于 InnerJoin，但它使用的是具有`相同字段名`的两个表进行连接。
//
// 例如：
// Model("order").X内连接相同字段("user", "user_id")
// Model("order").X内连接相同字段("product", "product_id")
// md5:bdc954b5bcb8a9c5
func (m *Model) X内连接相同字段(表名, 相同字段名 string) *Model {
	return m.doJoin(joinOperatorInner, 表名, fmt.Sprintf(
		`%s.%s=%s.%s`,
		m.tablesInit,
		m.db.X取Core对象().X底层QuoteWord(相同字段名),
		m.db.X取Core对象().X底层QuoteWord(表名),
		m.db.X取Core对象().X底层QuoteWord(相同字段名),
	))
}

// X左连接带比较运算符 执行类似于 LeftJoin 的操作，但允许指定不同的字段和比较运算符。
//
// 例如：
// Model("user").X左连接带比较运算符("order", "id", "=", "user_id")
// Model("user").X左连接带比较运算符("order", "id", ">", "user_id")
// Model("user").X左连接带比较运算符("order", "id", "<", "user_id")
// md5:90ce0e2226eb4b30
func (m *Model) X左连接带比较运算符(表名, 第一个字段, 比较运算符, 第二个字段 string) *Model {
	return m.doJoin(joinOperatorLeft, 表名, fmt.Sprintf(
		`%s.%s %s %s.%s`,
		m.tablesInit,
		m.db.X取Core对象().X底层QuoteWord(第一个字段),
		比较运算符,
		m.db.X取Core对象().X底层QuoteWord(表名),
		m.db.X取Core对象().X底层QuoteWord(第二个字段),
	))
}

// X右连接带比较运算符 执行右连接操作。它指定了不同的字段和比较运算符。
//
// 例如：
// User("user").X右连接带比较运算符("order", "id", "=", "user_id")
// User("user").X右连接带比较运算符("order", "id", ">", "user_id")
// User("user").X右连接带比较运算符("order", "id", "<", "user_id")
// 
// 这里，`X右连接带比较运算符` 是一个 Go 代码中的函数，用于在查询数据库时执行右连接操作，并且允许用户自定义连接的字段和比较操作。第一个参数是模型名（如 "user"），接下来的参数包括要连接的表名、字段名以及连接条件（等于、大于或小于）。
// md5:563f2b0f155fc829
func (m *Model) X右连接带比较运算符(表名, 第一个字段, 比较运算符, 第二个字段 string) *Model {
	return m.doJoin(joinOperatorRight, 表名, fmt.Sprintf(
		`%s.%s %s %s.%s`,
		m.tablesInit,
		m.db.X取Core对象().X底层QuoteWord(第一个字段),
		比较运算符,
		m.db.X取Core对象().X底层QuoteWord(表名),
		m.db.X取Core对象().X底层QuoteWord(第二个字段),
	))
}

// X内连接带比较运算符 执行 InnerJoin 操作。它指定了不同的字段和比较运算符。
// 
// 例如：
// Model("user").X内连接带比较运算符("order", "id", "=", "user_id")
// Model("user").X内连接带比较运算符("order", "id", ">", "user_id")
// Model("user").X内连接带比较运算符("order", "id", "<", "user_id")
// 
// 这段代码是在 Go 语言中定义了一个方法，用于在两个数据表之间执行内连接（InnerJoin），并允许用户指定连接的字段以及比较运算符。例如，`"user".X内连接带比较运算符("order", "id", "=", "user_id")` 表示连接 "user" 表和 "order" 表，通过 "id" 字段进行等号（=）匹配。其他示例展示了使用大于（>）和小于（<）运算符的情况。
// md5:0499f4b5bbbc2016
func (m *Model) X内连接带比较运算符(表名, 第一个字段, 比较运算符, 第二个字段 string) *Model {
	return m.doJoin(joinOperatorInner, 表名, fmt.Sprintf(
		`%s.%s %s %s.%s`,
		m.tablesInit,
		m.db.X取Core对象().X底层QuoteWord(第一个字段),
		比较运算符,
		m.db.X取Core对象().X底层QuoteWord(表名),
		m.db.X取Core对象().X底层QuoteWord(第二个字段),
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
			joinStr = gstr.X过滤首尾符并含空白(tableOrSubQueryAndJoinConditions[0])
			if joinStr[0] != '(' {
				joinStr = "(" + joinStr + ")"
			}
		} else {
			table = tableOrSubQueryAndJoinConditions[0]
			joinStr = m.db.X取Core对象().X底层添加前缀字符和引用字符(table)
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
			m.db.X取Core对象().X底层QuoteWord(alias),
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
	s = gstr.X过滤首字符并含空白(s, "()")
	if p := gstr.X查找(s, " "); p != -1 {
		if gstr.X相等比较并忽略大小写(s[:p], "select") {
			return true
		}
	}
	return false
}
