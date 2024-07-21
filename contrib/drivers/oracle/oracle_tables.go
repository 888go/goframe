// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package oracle

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

const (
	tablesSqlTmp = `SELECT TABLE_NAME FROM USER_TABLES ORDER BY TABLE_NAME`
)

// Tables 获取并返回当前模式下的表。
// 主要用于CLI工具链，用于自动生成模型。
// 注意，在Oracle数据库中，它会忽略`schema`参数，因为该参数是不必要的。
// md5:75a87bb44fddc91a
func (d *Driver) Tables(ctx context.Context, schema ...string) (tables []string, err error) {
	var result gdb.Result
	// 不要将`usedSchema`作为`SlaveLink`函数的参数。 md5:283541defa4ac558
	link, err := d.SlaveLink(schema...)
	if err != nil {
		return nil, err
	}
	result, err = d.DoSelect(ctx, link, tablesSqlTmp)
	if err != nil {
		return
	}
	for _, m := range result {
		for _, v := range m {
			tables = append(tables, v.String())
		}
	}
	return
}
