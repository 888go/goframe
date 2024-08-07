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
	"strings"

	gdb "github.com/888go/goframe/database/gdb"
	gutil "github.com/888go/goframe/util/gutil"
)

var (
	tableFieldsSqlTmp = `
SELECT 
    COLUMN_NAME AS FIELD, 
    CASE   
    WHEN (DATA_TYPE='NUMBER' AND NVL(DATA_SCALE,0)=0) THEN 'INT'||'('||DATA_PRECISION||','||DATA_SCALE||')'
    WHEN (DATA_TYPE='NUMBER' AND NVL(DATA_SCALE,0)>0) THEN 'FLOAT'||'('||DATA_PRECISION||','||DATA_SCALE||')'
    WHEN DATA_TYPE='FLOAT' THEN DATA_TYPE||'('||DATA_PRECISION||','||DATA_SCALE||')' 
    ELSE DATA_TYPE||'('||DATA_LENGTH||')' END AS TYPE,NULLABLE
FROM USER_TAB_COLUMNS WHERE TABLE_NAME = '%s' ORDER BY COLUMN_ID
`
)

func init() {
	var err error
	tableFieldsSqlTmp, err = gdb.FormatMultiLineSqlToSingle(tableFieldsSqlTmp)
	if err != nil {
		panic(err)
	}
}

// X取表字段信息Map 获取并返回当前模式下指定表的字段信息。
//
// 参见 DriverMysql.X取表字段信息Map。
// md5:7f7a75c67e38ad22
func (d *Driver) X取表字段信息Map(ctx context.Context, table string, schema ...string) (fields map[string]*gdb.TableField, err error) {
	var (
		result       gdb.Result
		link         gdb.Link
		usedSchema   = gutil.X取文本值或取默认值(d.X取默认数据库名称(), schema...)
		structureSql = fmt.Sprintf(tableFieldsSqlTmp, strings.ToUpper(table))
	)
	if link, err = d.X底层SlaveLink(usedSchema); err != nil {
		return nil, err
	}
	result, err = d.X底层查询(ctx, link, structureSql)
	if err != nil {
		return nil, err
	}
	fields = make(map[string]*gdb.TableField)
	for i, m := range result {
		isNull := false
		if m["NULLABLE"].String() == "Y" {
			isNull = true
		}

		fields[m["FIELD"].String()] = &gdb.TableField{
			Index: i,
			X名称:  m["FIELD"].String(),
			X类型:  m["TYPE"].String(),
			Null:  isNull,
		}
	}
	return fields, nil
}
