// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package mysql

import (
	"context"

	gdb "github.com/888go/goframe/database/gdb"
)

// X取表名称切片 获取并返回当前模式下的表格列表。
//主要用于命令行工具链，用于自动生成模型。
// md5:bce161ba95454bf5
func (d *Driver) X取表名称切片(ctx context.Context, schema ...string) (tables []string, err error) {
	var result gdb.Result
	link, err := d.X底层SlaveLink(schema...)
	if err != nil {
		return nil, err
	}
	result, err = d.X底层查询(ctx, link, `SHOW TABLES`)
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
