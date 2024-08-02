// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package oracle

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	gdb "github.com/888go/goframe/database/gdb"
	gregex "github.com/888go/goframe/text/gregex"
	gstr "github.com/888go/goframe/text/gstr"
)

var (
	newSqlReplacementTmp = `
SELECT * FROM (
	SELECT GFORM.*, ROWNUM ROWNUM_ FROM (%s %s) GFORM WHERE ROWNUM <= %d
) 
	WHERE ROWNUM_ > %d
`
)

func init() {
	var err error
	newSqlReplacementTmp, err = gdb.FormatMultiLineSqlToSingle(newSqlReplacementTmp)
	if err != nil {
		panic(err)
	}
}

// DoFilter 在将 SQL 字符串提交给底层 SQL 驱动程序之前处理它。 md5:f9ff7431f1478cfb
func (d *Driver) DoFilter(ctx context.Context, link gdb.Link, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error) {
	var index int
	newArgs = args
		// 将占位符字符 '?' 转换为字符串 ":vx"。 md5:14aed71041f34fec
	newSql, err = gregex.ReplaceStringFunc("\\?", sql, func(s string) string {
		index++
		return fmt.Sprintf(":v%d", index)
	})
	if err != nil {
		return
	}
	newSql, err = gregex.ReplaceString("\"", "", newSql)
	if err != nil {
		return
	}
	newSql, err = d.parseSql(newSql)
	if err != nil {
		return
	}
	return d.Core.DoFilter(ctx, link, newSql, newArgs)
}

// parseSql 在将 SQL 语句提交给底层驱动程序之前，进行一些替换，以支持 Oracle 服务器。
// md5:be49fd2231417d4d
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
	var (
		match  [][]string
		patten = `^\s*(?i)(SELECT)|(LIMIT\s*(\d+)\s*,{0,1}\s*(\d*))`
	)
	match, err = gregex.MatchAllString(patten, toBeCommittedSql)
	if err != nil {
		return "", err
	}
	if len(match) == 0 {
		return toBeCommittedSql, nil
	}
	var index = 1
	if len(match) < 2 || strings.HasPrefix(match[index][0], "LIMIT") == false {
		return toBeCommittedSql, nil
	}
		// 只处理`SELECT ... LIMIT ...`语句。 md5:94e3efd3997b47e2
	queryExpr, err := gregex.MatchString("((?i)SELECT)(.+)((?i)LIMIT)", toBeCommittedSql)
	if err != nil {
		return "", err
	}
	if len(queryExpr) == 0 {
		return toBeCommittedSql, nil
	}
	if len(queryExpr) != 4 ||
		strings.EqualFold(queryExpr[1], "SELECT") == false ||
		strings.EqualFold(queryExpr[3], "LIMIT") == false {
		return toBeCommittedSql, nil
	}
	page, limit := 0, 0
	for i := 1; i < len(match[index]); i++ {
		if len(strings.TrimSpace(match[index][i])) == 0 {
			continue
		}
		if strings.HasPrefix(match[index][i], "LIMIT") {
			if match[index][i+2] != "" {
				page, err = strconv.Atoi(match[index][i+1])
				if err != nil {
					return "", err
				}
				limit, err = strconv.Atoi(match[index][i+2])
				if err != nil {
					return "", err
				}
				if page <= 0 {
					page = 1
				}
				limit = (page/limit + 1) * limit
				page, err = strconv.Atoi(match[index][i+1])
				if err != nil {
					return "", err
				}
			} else {
				limit, err = strconv.Atoi(match[index][i+1])
				if err != nil {
					return "", err
				}
			}
			break
		}
	}
	var newReplacedSql = fmt.Sprintf(
		newSqlReplacementTmp,
		queryExpr[1], queryExpr[2], limit, page,
	)
	return newReplacedSql, nil
}
