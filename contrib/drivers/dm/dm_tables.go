// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package dm

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

const (
	tablesSqlTmp = `SELECT * FROM ALL_TABLES`
)

// Tables 获取并返回当前模式下的表格列表。
//主要用于命令行工具链，用于自动生成模型。
// md5:bce161ba95454bf5
// ff:
// d:
// ctx:
// schema:
// tables:
// err:
func (d *Driver) Tables(ctx context.Context, schema ...string) (tables []string, err error) {
	var result gdb.Result
	// 当schema为空时，返回默认链接. md5:fd8d52adc6efeedb
	link, err := d.SlaveLink(schema...)
	if err != nil {
		return nil, err
	}
	// 链接已经区分，不再需要判断归属. md5:397cb7fafe12c367
	result, err = d.DoSelect(ctx, link, tablesSqlTmp)
	if err != nil {
		return
	}
	for _, m := range result {
		if v, ok := m["IOT_NAME"]; ok {
			tables = append(tables, v.String())
		}
	}
	return
}
