// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package dm

import (
	"context"

	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

// DoFilter 在将 SQL 字符串提交给底层 SQL 驱动程序之前处理它。 md5:f9ff7431f1478cfb
func (d *Driver) DoFilter(
	ctx context.Context, link gdb.Link, sql string, args []interface{},
) (newSql string, newArgs []interface{}, err error) {
	// 因为之前已经在字段处理中完成了大写，所以不需要再大写. md5:32e3319fc42c6edf
	newSql, _ = gregex.ReplaceString(`["\n\t]`, "", sql)
	newSql = gstr.ReplaceI(gstr.ReplaceI(newSql, "GROUP_CONCAT", "LISTAGG"), "SEPARATOR", ",")

	// 待办事项：当前的方法太过粗略。我们应该处理GROUP_CONCAT函数，以及从匹配的select语句中解析索引字段的问题。
	// （GROUP_CONCAT功能DM不支持；索引不能作为查询列名使用，并且需要添加安全字符，例如将"index"转义）
	// md5:125ee1107dd70034
	l, r := d.GetChars()
	if strings.Contains(newSql, "INDEX") || strings.Contains(newSql, "index") {
		if !(strings.Contains(newSql, "_INDEX") || strings.Contains(newSql, "_index")) {
			newSql = gstr.ReplaceI(newSql, "INDEX", l+"INDEX"+r)
		}
	}

	// 待办事项：我尝试过但从未成功：
	// 通过正则表达式匹配SQL中的"INDEX"：
	// array, err := gregex.MatchAllString(`SELECT (.*INDEX.*) FROM .*`, newSql)
	// 打印错误信息：
	// g.Dump("err:", err)
	// 打印匹配结果：
	// g.Dump("array:", array)
	// 打印第一个匹配项的第二部分：
	// g.Dump("array:", array[0][1])
	// md5:46650cd1fe9bb3a8

	// 使用正则表达式 `SELECT (.*INDEX.*) FROM .*` 替换原SQL（将 `l` 后面跟着 "INDEX"，再接 `r`），并将结果赋值给新的SQL字符串 `newSql`
	// 打印 "err:" 后面的错误信息
	// 打印 "newSql:" 后面的新SQL字符串
	// md5:5e9ef3312146be4d

	// 使用正则表达式编译模式：`.*SELECT (.*INDEX.*) FROM .*`
	// 将新的SQL字符串中的所有匹配到的子串用自定义函数替换
	// 自定义函数接受一个字符串参数data，打印"data: "和数据本身，然后返回数据本身
	// re, err := regexp.Compile(`.*SELECT (.*INDEX.*) FROM .*`)
	// newSql = re.ReplaceAllStringFunc(newSql, func(data string) string {
	// 	fmt.Println("data:", data)
	// 	return data
	// })
	// md5:e2b3231602b36621

	return d.Core.DoFilter(
		ctx,
		link,
		newSql,
		args,
	)
}
