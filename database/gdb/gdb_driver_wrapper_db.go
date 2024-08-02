// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"context"
	"database/sql"
	"fmt"

	gjson "github.com/888go/goframe/encoding/gjson"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	gstr "github.com/888go/goframe/text/gstr"
	gutil "github.com/888go/goframe/util/gutil"
)

// DriverWrapperDB是一个DB包装器，用于通过嵌入式DB扩展功能。 md5:a926644143c69c76
type DriverWrapperDB struct {
	DB
}

// Open 创建并返回一个用于pgsql的底层sql.DB对象。
// 参考链接：https://pkg.go.dev/github.com/lib/pq
// md5:9889bcb899248a2b
func (d *DriverWrapperDB) Open(node *ConfigNode) (db *sql.DB, err error) {
	var ctx = d.GetCtx()
	intlog.PrintFunc(ctx, func() string {
		return fmt.Sprintf(`open new connection:%s`, gjson.MustEncode(node))
	})
	return d.DB.Open(node)
}

// Tables 获取并返回当前模式下的表格列表。
//主要用于命令行工具链，用于自动生成模型。
// md5:bce161ba95454bf5
func (d *DriverWrapperDB) Tables(ctx context.Context, schema ...string) (tables []string, err error) {
	ctx = context.WithValue(ctx, ctxKeyInternalProducedSQL, struct{}{})
	return d.DB.Tables(ctx, schema...)
}

// TableFields 获取并返回当前模式指定表的字段信息。
// 
// 参数 `link` 是可选的，如果为 nil，则自动获取一个原始 SQL 连接，用于执行必要的 SQL 查询。
// 
// 它返回一个包含字段名及其对应字段的映射。由于映射是无序的，TableField 结构体有一个 "Index" 字段，标记其在字段中的顺序。
// 
// 该方法使用缓存功能来提高性能，直到进程重启，缓存永不过期。
// md5:c844572d5210b35e
func (d *DriverWrapperDB) TableFields(
	ctx context.Context, table string, schema ...string,
) (fields map[string]*TableField, err error) {
	if table == "" {
		return nil, nil
	}
	charL, charR := d.GetChars()
	table = gstr.Trim(table, charL+charR)
	if gstr.Contains(table, " ") {
		return nil, gerror.NewCode(
			gcode.CodeInvalidParameter,
			"function TableFields supports only single table operations",
		)
	}
	var (
				// 前缀:组@模式#表. md5:b22e67d9da02a91b
		cacheKey = fmt.Sprintf(
			`%s%s@%s#%s`,
			cachePrefixTableFields,
			d.GetGroup(),
			gutil.GetOrDefaultStr(d.GetSchema(), schema...),
			table,
		)
		value = tableFieldsMap.GetOrSetFuncLock(cacheKey, func() interface{} {
			ctx = context.WithValue(ctx, ctxKeyInternalProducedSQL, struct{}{})
			fields, err = d.DB.TableFields(ctx, table, schema...)
			if err != nil {
				return nil
			}
			return fields
		})
	)
	if value != nil {
		fields = value.(map[string]*TableField)
	}
	return
}

// DoInsert 用于插入或更新给定表的数据。
// 此函数通常用于自定义接口定义，您无需手动调用。
// 参数 `data` 可以为 map/gmap/struct/*struct/[]map/[]struct 等类型。
// 例如：
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
//
// 参数 `option` 的值如下：
// InsertOptionDefault：仅插入，如果数据中包含唯一键或主键，则返回错误；
// InsertOptionReplace：如果数据中包含唯一键或主键，先从表中删除原有记录，再插入新记录；
// InsertOptionSave：如果数据中包含唯一键或主键，进行更新，否则插入新记录；
// InsertOptionIgnore：如果数据中包含唯一键或主键，忽略插入操作。
// md5:9fab32fdc41df179
func (d *DriverWrapperDB) DoInsert(ctx context.Context, link Link, table string, list List, option DoInsertOption) (result sql.Result, err error) {
		// 在将数据类型提交给底层数据库驱动程序之前进行转换。 md5:58b56ae1ed22196f
	for i, item := range list {
		list[i], err = d.GetCore().ConvertDataForRecord(ctx, item, table)
		if err != nil {
			return nil, err
		}
	}
	return d.DB.DoInsert(ctx, link, table, list, option)
}
