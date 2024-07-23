// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package mssql

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

var (
	selectSqlTmp          = `SELECT * FROM (SELECT TOP %d * FROM (SELECT TOP %d %s) as TMP1_ ) as TMP2_ `
	selectWithOrderSqlTmp = `
SELECT * FROM (SELECT ROW_NUMBER() OVER (ORDER BY %s) as ROWNUMBER_, %s ) as TMP_ 
WHERE TMP_.ROWNUMBER_ > %d AND TMP_.ROWNUMBER_ <= %d
`
)

func init() {
	var err error
	selectWithOrderSqlTmp, err = gdb.FormatMultiLineSqlToSingle(selectWithOrderSqlTmp)
	if err != nil {
		panic(err)
	}
}

// DoFilter 在将 SQL 字符串提交给底层 SQL 驱动程序之前处理它。 md5:f9ff7431f1478cfb
func (d *Driver) DoFilter(
	ctx context.Context, link gdb.Link, sql string, args []interface{},
) (newSql string, newArgs []interface{}, err error) {
	var index int
	// 将占位符字符'?'转换为字符串"@px"。 md5:e9602cb8693766e3
	newSql, err = gregex.ReplaceStringFunc("\\?", sql, func(s string) string {
		index++
		return fmt.Sprintf("@p%d", index)
	})
	if err != nil {
		return "", nil, err
	}
	newSql, err = gregex.ReplaceString("\"", "", newSql)
	if err != nil {
		return "", nil, err
	}
	newSql, err = d.parseSql(newSql)
	if err != nil {
		return "", nil, err
	}
	newArgs = args
	return d.Core.DoFilter(ctx, link, newSql, newArgs)
}

// parseSql 在将 SQL 语句提交给底层驱动程序之前，进行一些替换，以支持微软SQL服务器。
// md5:dedc2c424a170a26
func (d *Driver) parseSql(toBeCommittedSql string) (string, error) {
	var (
		err       error
		operation = gstr.StrTillEx(toBeCommittedSql, " ")
		keyword   = strings.ToUpper(gstr.Trim(operation))
	)
	switch keyword {
	case "SELECT":
		toBeCommittedSql, err = d.handleSelectSqlReplacement(toBeCommittedSql)
		if err != nil {
			return "", err
		}
	}
	return toBeCommittedSql, nil
}

func (d *Driver) handleSelectSqlReplacement(toBeCommittedSql string) (newSql string, err error) {
	// 查询USER表中ID为1的所有列，限制返回结果为1条. md5:3f978a0c9e2f99a6
	match, err := gregex.MatchString(`^SELECT(.+)LIMIT 1$`, toBeCommittedSql)
	if err != nil {
		return "", err
	}
	if len(match) > 1 {
		return fmt.Sprintf(`SELECT TOP 1 %s`, match[1]), nil
	}

	// 从USER表中选择所有AGE大于18的记录，按ID降序排序，从第100条开始，取200条数据. md5:b1500e0aa6edbbb7
	patten := `^\s*(?i)(SELECT)|(LIMIT\s*(\d+)\s*,\s*(\d+))`
	if gregex.IsMatchString(patten, toBeCommittedSql) == false {
		return toBeCommittedSql, nil
	}
	allMatch, err := gregex.MatchAllString(patten, toBeCommittedSql)
	if err != nil {
		return "", err
	}
	var index = 1
	// LIMIT语句检查。 md5:b3cd8a3a3e1cd305
	if len(allMatch) < 2 ||
		(strings.HasPrefix(allMatch[index][0], "LIMIT") == false &&
			strings.HasPrefix(allMatch[index][0], "limit") == false) {
		return toBeCommittedSql, nil
	}
	if gregex.IsMatchString("((?i)SELECT)(.+)((?i)LIMIT)", toBeCommittedSql) == false {
		return toBeCommittedSql, nil
	}
	// 检查ORDER BY语句。 md5:5ff030b44f15e56a
	var (
		selectStr = ""
		orderStr  = ""
		haveOrder = gregex.IsMatchString("((?i)SELECT)(.+)((?i)ORDER BY)", toBeCommittedSql)
	)
	if haveOrder {
		queryExpr, _ := gregex.MatchString("((?i)SELECT)(.+)((?i)ORDER BY)", toBeCommittedSql)
		if len(queryExpr) != 4 ||
			strings.EqualFold(queryExpr[1], "SELECT") == false ||
			strings.EqualFold(queryExpr[3], "ORDER BY") == false {
			return toBeCommittedSql, nil
		}
		selectStr = queryExpr[2]
		orderExpr, _ := gregex.MatchString("((?i)ORDER BY)(.+)((?i)LIMIT)", toBeCommittedSql)
		if len(orderExpr) != 4 ||
			strings.EqualFold(orderExpr[1], "ORDER BY") == false ||
			strings.EqualFold(orderExpr[3], "LIMIT") == false {
			return toBeCommittedSql, nil
		}
		orderStr = orderExpr[2]
	} else {
		queryExpr, _ := gregex.MatchString("((?i)SELECT)(.+)((?i)LIMIT)", toBeCommittedSql)
		if len(queryExpr) != 4 ||
			strings.EqualFold(queryExpr[1], "SELECT") == false ||
			strings.EqualFold(queryExpr[3], "LIMIT") == false {
			return toBeCommittedSql, nil
		}
		selectStr = queryExpr[2]
	}
	first, limit := 0, 0
	for i := 1; i < len(allMatch[index]); i++ {
		if len(strings.TrimSpace(allMatch[index][i])) == 0 {
			continue
		}
		if strings.HasPrefix(allMatch[index][i], "LIMIT") ||
			strings.HasPrefix(allMatch[index][i], "limit") {
			first, _ = strconv.Atoi(allMatch[index][i+1])
			limit, _ = strconv.Atoi(allMatch[index][i+2])
			break
		}
	}
	if haveOrder {
		toBeCommittedSql = fmt.Sprintf(
			selectWithOrderSqlTmp,
			orderStr, selectStr, first, first+limit,
		)
		return toBeCommittedSql, nil
	}

	if first == 0 {
		first = limit
	}
	toBeCommittedSql = fmt.Sprintf(
		selectSqlTmp,
		limit, first+limit, selectStr,
	)
	return toBeCommittedSql, nil
}
