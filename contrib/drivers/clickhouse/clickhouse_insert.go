// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package clickhouse

import (
	"context"
	"database/sql"
)

// X插入并跳过已存在 不支持用于修改数据部分的其他查询：REPLACE、MERGE、UPSERT、INSERT UPDATE。 md5:ac3efdb87c360d83
func (d *Driver) X插入并跳过已存在(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error) {
	return nil, errUnsupportedInsertIgnore
}

// X插入并取ID 不支持其他用于修改数据部分的查询：REPLACE、MERGE、UPSERT、INSERT UPDATE。 md5:9d4693bead6866d9
func (d *Driver) X插入并取ID(ctx context.Context, table string, data interface{}, batch ...int) (int64, error) {
	return 0, errUnsupportedInsertGetId
}

// 不支持用于修改数据部分的其他查询：REPLACE、MERGE、UPSERT、INSERT UPDATE。 md5:d0b1de268614fdfa
func (d *Driver) X插入并替换已存在(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error) {
	return nil, errUnsupportedReplace
}
