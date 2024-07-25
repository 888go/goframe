// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package dm

import (
	"context"
	"fmt"

	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gutil"
)

const (
	tableFieldsSqlTmp = `SELECT * FROM ALL_TAB_COLUMNS WHERE Table_Name= '%s' AND OWNER = '%s'`
)

// TableFields 获取并返回当前模式中指定表的字段信息。 md5:920febaff284f5e7
func (d *Driver) TableFields(
	ctx context.Context, table string, schema ...string,
) (fields map[string]*gdb.TableField, err error) {
	var (
		result gdb.Result
		link   gdb.Link
						// 当没有指定模式时，默认返回配置项. md5:7872fad26e099386
		usedSchema = gutil.GetOrDefaultStr(d.GetSchema(), schema...)
	)
			// 当usedSchema为空时，返回默认链接. md5:8e4a43a3b2726ef6
	if link, err = d.SlaveLink(usedSchema); err != nil {
		return nil, err
	}
		// 链接已经区分，不再需要判断归属. md5:397cb7fafe12c367
	result, err = d.DoSelect(
		ctx, link,
		fmt.Sprintf(
			tableFieldsSqlTmp,
			strings.ToUpper(table),
			strings.ToUpper(d.GetSchema()),
		),
	)
	if err != nil {
		return nil, err
	}
	fields = make(map[string]*gdb.TableField)
	for i, m := range result {
		// m[nullable] 返回"N"或"Y"
		// "N" 表示非空
		// "Y" 表示可能为空
		// md5:63d3103e22469aea
		var nullable bool
		if m["NULLABLE"].String() != "N" {
			nullable = true
		}
		fields[m["COLUMN_NAME"].String()] = &gdb.TableField{
			Index:   i,
			Name:    m["COLUMN_NAME"].String(),
			Type:    m["DATA_TYPE"].String(),
			Null:    nullable,
			Default: m["DATA_DEFAULT"].Val(),
			// Key:     m["Key"].String()， 			// 关键字：将m中"Key"对应的值转换为字符串
			// Extra:   m["Extra"].String()， 			// 副本：将m中"Extra"对应的值转换为字符串
			// Comment: m["Comment"].String()， 			// 注释：将m中"Comment"对应的值转换为字符串
			// md5:ef32e9151c11fe98
		}
	}
	return fields, nil
}
