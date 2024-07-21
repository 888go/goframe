// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package mssql

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gutil"
)

var (
	tableFieldsSqlTmp = `
SELECT 
	a.name Field,
	CASE b.name 
		WHEN 'datetime' THEN 'datetime'
		WHEN 'numeric' THEN b.name + '(' + convert(varchar(20), a.xprec) + ',' + convert(varchar(20), a.xscale) + ')' 
		WHEN 'char' THEN b.name + '(' + convert(varchar(20), a.length)+ ')'
		WHEN 'varchar' THEN b.name + '(' + convert(varchar(20), a.length)+ ')'
		ELSE b.name + '(' + convert(varchar(20),a.length)+ ')' END AS Type,
	CASE WHEN a.isnullable=1 THEN 'YES' ELSE 'NO' end AS [Null],
	CASE WHEN exists (
		SELECT 1 FROM sysobjects WHERE xtype='PK' AND name IN (
			SELECT name FROM sysindexes WHERE indid IN (
				SELECT indid FROM sysindexkeys WHERE id = a.id AND colid=a.colid
			)
		)
	) THEN 'PRI' ELSE '' END AS [Key],
	CASE WHEN COLUMNPROPERTY(a.id,a.name,'IsIdentity')=1 THEN 'auto_increment' ELSE '' END Extra,
	isnull(e.text,'') AS [Default],
	isnull(g.[value],'') AS [Comment]
FROM syscolumns a
LEFT  JOIN systypes b ON a.xtype=b.xtype AND a.xusertype=b.xusertype
INNER JOIN sysobjects d ON a.id=d.id AND d.xtype='U' AND d.name<>'dtproperties'
LEFT  JOIN syscomments e ON a.cdefault=e.id
LEFT  JOIN sys.extended_properties g ON a.id=g.major_id AND a.colid=g.minor_id
LEFT  JOIN sys.extended_properties f ON d.id=f.major_id AND f.minor_id =0
WHERE d.name='%s'
ORDER BY a.id,a.colorder
`
)

func init() {
	var err error
	tableFieldsSqlTmp, err = gdb.FormatMultiLineSqlToSingle(tableFieldsSqlTmp)
	if err != nil {
		panic(err)
	}
}

		// TableFields 获取并返回当前模式下指定表的字段信息。
		//
		// 参见 DriverMysql.TableFields。
		// md5:7f7a75c67e38ad22
// ff:
// d:
// ctx:
// table:
// schema:
// fields:
// err:
func (d *Driver) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*gdb.TableField, err error) {
	var (
		result     gdb.Result
		link       gdb.Link
		usedSchema = gutil.GetOrDefaultStr(d.GetSchema(), schema...)
	)
	if link, err = d.SlaveLink(usedSchema); err != nil {
		return nil, err
	}
	structureSql := fmt.Sprintf(tableFieldsSqlTmp, table)
	result, err = d.DoSelect(ctx, link, structureSql)
	if err != nil {
		return nil, err
	}
	fields = make(map[string]*gdb.TableField)
	for i, m := range result {
		fields[m["Field"].String()] = &gdb.TableField{
			Index:   i,
			Name:    m["Field"].String(),
			Type:    m["Type"].String(),
			Null:    m["Null"].Bool(),
			Key:     m["Key"].String(),
			Default: m["Default"].Val(),
			Extra:   m["Extra"].String(),
			Comment: m["Comment"].String(),
		}
	}
	return fields, nil
}
