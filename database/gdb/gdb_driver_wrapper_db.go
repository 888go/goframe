// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"context"
	"database/sql"
	"fmt"
	
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gutil"
)

// DriverWrapperDB 是一个数据库（DB）包装器，用于通过嵌入式数据库扩展功能。
type DriverWrapperDB struct {
	DB
}

// Open创建并返回一个用于pgsql的底层sql.DB对象。
// 参考文档：https://pkg.go.dev/github.com/lib/pq
func (d *DriverWrapperDB) Open(node *X配置项) (db *sql.DB, err error) {
	var ctx = d.X取上下文对象()
	intlog.PrintFunc(ctx, func() string {
		return fmt.Sprintf(`open new connection:%s`, json类.X变量到json字节集PANI(node))
	})
	return d.DB.X底层Open(node)
}

// Tables 获取并返回当前模式的表。
// 它主要用于cli工具链中，用于自动生成模型。
func (d *DriverWrapperDB) X取表名称数组(上下文 context.Context, schema ...string) (表名称数组 []string, 错误 error) {
	上下文 = context.WithValue(上下文, ctxKeyInternalProducedSQL, struct{}{})
	return d.DB.X取表名称数组(上下文, schema...)
}

// TableFields 获取并返回当前模式下指定表的字段信息。
//
// 参数 `link` 是可选的，如果给出 nil，则会自动获取一个原始的 SQL 连接作为其链接以执行必要的 SQL 查询。
//
// 注意，它返回一个包含字段名及其对应字段信息的映射。由于映射是无序的，TableField 结构体有一个 "Index" 字段用于标记其在所有字段中的顺序。
//
// 为了提高性能，该函数使用了缓存特性，缓存有效期直到进程重启才会过期。
func (d *DriverWrapperDB) X取表字段信息Map(
	ctx context.Context, table string, schema ...string,
) (fields map[string]*X字段信息, err error) {
	if table == "" {
		return nil, nil
	}
	charL, charR := d.X底层取数据库安全字符()
	table = 文本类.X过滤首尾符并含空白(table, charL+charR)
	if 文本类.X是否包含(table, " ") {
		return nil, 错误类.X创建错误码(
			错误码类.CodeInvalidParameter,
			"function TableFields supports only single table operations",
		)
	}
	var (
		cacheKey = fmt.Sprintf(
			`%s%s@%s#%s`,
			cachePrefixTableFields,
			d.X取配置组名称(),
			工具类.X取文本值或取默认值(d.X取默认数据库名称(), schema...),
			table,
		)
		value = tableFieldsMap.X取值或设置值_函数带锁(cacheKey, func() interface{} {
			ctx = context.WithValue(ctx, ctxKeyInternalProducedSQL, struct{}{})
			fields, err = d.DB.X取表字段信息Map(ctx, table, schema...)
			if err != nil {
				return nil
			}
			return fields
		})
	)
	if value != nil {
		fields = value.(map[string]*X字段信息)
	}
	return
}

// DoInsert 对给定表插入或更新数据。
// 该函数通常用于自定义接口定义，您无需手动调用它。
// 参数`data`的类型可以是 map/gmap/struct/*struct/[]map/[]struct 等。
// 示例：
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
// 参数 `option` 的取值如下：
// InsertOptionDefault：仅插入，如果数据中存在唯一/主键，则返回错误；
// InsertOptionReplace：如果数据中存在唯一/主键，先从表中删除并插入新的记录；
// InsertOptionSave：如果数据中存在唯一/主键，则更新记录，否则插入新记录；
// InsertOptionIgnore：如果数据中存在唯一/主键，则忽略插入操作。
func (d *DriverWrapperDB) X底层插入(上下文 context.Context, 链接 X底层链接, 表名称 string, list Map数组, 选项 X底层输入) (结果 sql.Result, 错误 error) {
	// 在提交给底层数据库驱动之前，转换数据类型。
	for i, item := range list {
		list[i], 错误 = d.X取Core对象().X底层ConvertDataForRecord(上下文, item, 表名称)
		if 错误 != nil {
			return nil, 错误
		}
	}
	return d.DB.X底层插入(上下文, 链接, 表名称, list, 选项)
}
