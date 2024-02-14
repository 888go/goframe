// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package mssql 实现了 gdb.Driver 接口，该接口支持对 MSSql 数据库进行操作。
//
// 注意：
// 1. 不支持 Save/Replace 功能。
// 2. 不支持 LastInsertId 功能。
// 以下是将上述Go语言代码注释翻译成中文：
// ```markdown
// 这个mssql包实现了gdb.Driver接口，主要用于对MSSql数据库的各种操作提供支持。
//
// 注意事项：
// 1. 该实现暂不支持Save/Replace功能。
// 2. 该实现暂不支持LastInsertId方法。
package mssql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	
	_ "github.com/denisenkom/go-mssqldb"
	
	"github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gutil"
)

// Driver 是 SQL 服务器数据库的驱动程序。
type Driver struct {
	*db类.Core
}

const (
	quoteChar = `"`
)

func init() {
	if err := db类.X注册驱动(`mssql`, New()); err != nil {
		panic(err)
	}
}

// New 创建并返回一个实现 gdb.Driver 接口的驱动程序，该驱动支持对 Mssql 的操作。
func New() db类.Driver {
	return &Driver{}
}

// New 创建并返回一个用于SQL服务器的数据库对象。
// 它实现了gdb.Driver接口，以便支持额外的数据库驱动安装。
func (d *Driver) New(core *db类.Core, node *db类.X配置项) (db类.DB, error) {
	return &Driver{
		Core: core,
	}, nil
}

// Open 创建并返回一个用于mssql的底层sql.DB对象。
func (d *Driver) X底层Open(配置对象 *db类.X配置项) (db *sql.DB, 错误 error) {
	var (
		source               string
		underlyingDriverName = "sqlserver"
	)
	if 配置对象.X自定义链接信息 != "" {
// ============================================================================
// 从 v2.2.0 版本开始已弃用。
// ============================================================================
		source = 配置对象.X自定义链接信息
		// 在运行时自定义更改架构
		if 配置对象.X名称 != "" {
			source, _ = 正则类.X替换文本(`database=([\w\.\-]+)+`, "database="+配置对象.X名称, source)
		}
	} else {
		source = fmt.Sprintf(
			"user id=%s;password=%s;server=%s;port=%s;database=%s;encrypt=disable",
			配置对象.X账号, 配置对象.X密码, 配置对象.X地址, 配置对象.X端口, 配置对象.X名称,
		)
		if 配置对象.X额外 != "" {
			var extraMap map[string]interface{}
			if extraMap, 错误 = 文本类.X参数解析(配置对象.X额外); 错误 != nil {
				return nil, 错误
			}
			for k, v := range extraMap {
				source += fmt.Sprintf(`;%s=%s`, k, v)
			}
		}
	}

	if db, 错误 = sql.Open(underlyingDriverName, source); 错误 != nil {
		错误 = 错误类.X多层错误码并格式化(
			错误码类.CodeDbOperationError, 错误,
			`sql.Open failed for driver "%s" by source "%s"`, underlyingDriverName, source,
		)
		return nil, 错误
	}
	return
}

// GetChars 返回该类型数据库的安全字符。
func (d *Driver) X底层取数据库安全字符() (左字符 string, 右字符 string) {
	return quoteChar, quoteChar
}

// DoFilter 在将 SQL 字符串提交给底层 SQL 驱动程序之前，对其进行处理。
func (d *Driver) X底层DoFilter(ctx context.Context, link db类.Link, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error) {
	var index int
	// 将占位符字符 '?' 转换为字符串 "@px"。
	newSql, _ = 正则类.X替换文本_函数("\\?", sql, func(s string) string {
		index++
		return fmt.Sprintf("@p%d", index)
	})
	newSql, _ = 正则类.X替换文本("\"", "", newSql)
	return d.Core.X底层DoFilter(ctx, link, d.parseSql(newSql), args)
}

// parseSql在将SQL提交给底层驱动程序之前执行一些替换操作，
// 以便支持Microsoft SQL Server。
func (d *Driver) parseSql(sql string) string {
	// 从USER表中选取ID为1的记录，限制返回结果数量为1条
	if m, _ := 正则类.X匹配文本(`^SELECT(.+)LIMIT 1$`, sql); len(m) > 1 {
		return fmt.Sprintf(`SELECT TOP 1 %s`, m[1])
	}
	// 从USER表中选取AGE大于18的所有列，并按ID降序排列，然后获取第101至300条记录（LIMIT offset, count语法）
	patten := `^\s*(?i)(SELECT)|(LIMIT\s*(\d+)\s*,\s*(\d+))`
	if 正则类.X是否匹配文本(patten, sql) == false {
		return sql
	}
	res, err := 正则类.X匹配全部文本(patten, sql)
	if err != nil {
		return ""
	}
	var (
		index   = 0
		keyword = strings.TrimSpace(res[index][0])
	)
	index++
	switch strings.ToUpper(keyword) {
	case "SELECT":
		// LIMIT语句检查。
		if len(res) < 2 ||
			(strings.HasPrefix(res[index][0], "LIMIT") == false &&
				strings.HasPrefix(res[index][0], "limit") == false) {
			break
		}
		if 正则类.X是否匹配文本("((?i)SELECT)(.+)((?i)LIMIT)", sql) == false {
			break
		}
		// ORDER BY 语句检查。
		var (
			selectStr = ""
			orderStr  = ""
			haveOrder = 正则类.X是否匹配文本("((?i)SELECT)(.+)((?i)ORDER BY)", sql)
		)
		if haveOrder {
			queryExpr, _ := 正则类.X匹配文本("((?i)SELECT)(.+)((?i)ORDER BY)", sql)
			if len(queryExpr) != 4 ||
				strings.EqualFold(queryExpr[1], "SELECT") == false ||
				strings.EqualFold(queryExpr[3], "ORDER BY") == false {
				break
			}
			selectStr = queryExpr[2]
			orderExpr, _ := 正则类.X匹配文本("((?i)ORDER BY)(.+)((?i)LIMIT)", sql)
			if len(orderExpr) != 4 ||
				strings.EqualFold(orderExpr[1], "ORDER BY") == false ||
				strings.EqualFold(orderExpr[3], "LIMIT") == false {
				break
			}
			orderStr = orderExpr[2]
		} else {
			queryExpr, _ := 正则类.X匹配文本("((?i)SELECT)(.+)((?i)LIMIT)", sql)
			if len(queryExpr) != 4 ||
				strings.EqualFold(queryExpr[1], "SELECT") == false ||
				strings.EqualFold(queryExpr[3], "LIMIT") == false {
				break
			}
			selectStr = queryExpr[2]
		}
		first, limit := 0, 0
		for i := 1; i < len(res[index]); i++ {
			if len(strings.TrimSpace(res[index][i])) == 0 {
				continue
			}

			if strings.HasPrefix(res[index][i], "LIMIT") ||
				strings.HasPrefix(res[index][i], "limit") {
				first, _ = strconv.Atoi(res[index][i+1])
				limit, _ = strconv.Atoi(res[index][i+2])
				break
			}
		}
		if haveOrder {
			sql = fmt.Sprintf(
				"SELECT * FROM "+
					"(SELECT ROW_NUMBER() OVER (ORDER BY %s) as ROWNUMBER_, %s ) as TMP_ "+
					"WHERE TMP_.ROWNUMBER_ > %d AND TMP_.ROWNUMBER_ <= %d",
				orderStr, selectStr, first, first+limit,
			)
		} else {
			if first == 0 {
				first = limit
			}
			sql = fmt.Sprintf(
				"SELECT * FROM (SELECT TOP %d * FROM (SELECT TOP %d %s) as TMP1_ ) as TMP2_ ",
				limit, first+limit, selectStr,
			)
		}
	default:
	}
	return sql
}

// Tables 获取并返回当前模式的表格。
// 它主要用于cli工具链中，用于自动生成模型。
func (d *Driver) X取表名称数组(上下文 context.Context, schema ...string) (表名称数组 []string, 错误 error) {
	var result db类.Result
	link, 错误 := d.X底层SlaveLink(schema...)
	if 错误 != nil {
		return nil, 错误
	}

	result, 错误 = d.X底层查询(
		上下文, link, `SELECT NAME FROM SYSOBJECTS WHERE XTYPE='U' AND STATUS >= 0 ORDER BY NAME`,
	)
	if 错误 != nil {
		return
	}
	for _, m := range result {
		for _, v := range m {
			表名称数组 = append(表名称数组, v.String())
		}
	}
	return
}

// TableFields 获取并返回当前模式下指定表的字段信息。
//
// 另请参阅 DriverMysql.TableFields。
func (d *Driver) X取表字段信息Map(上下文 context.Context, 表名称 string, schema ...string) (字段信息Map map[string]*db类.TableField, 错误 error) {
	var (
		result     db类.Result
		link       db类.Link
		usedSchema = 工具类.X取文本值或取默认值(d.X取默认数据库名称(), schema...)
	)
	if link, 错误 = d.X底层SlaveLink(usedSchema); 错误 != nil {
		return nil, 错误
	}
	structureSql := fmt.Sprintf(`
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
LEFT JOIN systypes b ON a.xtype=b.xtype AND a.xusertype=b.xusertype
INNER JOIN sysobjects d ON a.id=d.id AND d.xtype='U' AND d.name<>'dtproperties'
LEFT JOIN syscomments e ON a.cdefault=e.id
LEFT JOIN sys.extended_properties g ON a.id=g.major_id AND a.colid=g.minor_id
LEFT JOIN sys.extended_properties f ON d.id=f.major_id AND f.minor_id =0
WHERE d.name='%s'
ORDER BY a.id,a.colorder`,
		表名称,
	)
	structureSql, _ = 正则类.X替换文本(`[\n\r\s]+`, " ", 文本类.X过滤首尾符并含空白(structureSql))
	result, 错误 = d.X底层查询(上下文, link, structureSql)
	if 错误 != nil {
		return nil, 错误
	}
	字段信息Map = make(map[string]*db类.TableField)
	for i, m := range result {
		字段信息Map[m["Field"].String()] = &db类.TableField{
			Index:   i,
			X名称:    m["Field"].String(),
			X类型:    m["Type"].String(),
			Null:    m["Null"].X取布尔(),
			Key:     m["Key"].String(),
			Default: m["Default"].X取值(),
			X额外:   m["Extra"].String(),
			Comment: m["Comment"].String(),
		}
	}
	return 字段信息Map, nil
}

// DoInsert 对给定表执行插入或更新数据操作。
func (d *Driver) X底层插入(上下文 context.Context, 链接 db类.Link, 表名称 string, list db类.Map数组, 选项 db类.DoInsertOption) (结果 sql.Result, 错误 error) {
	switch 选项.InsertOption {
	case db类.InsertOptionSave:
		return nil, 错误类.X创建错误码(
			错误码类.CodeNotSupported,
			`Save operation is not supported by mssql driver`,
		)

	case db类.InsertOptionReplace:
		return nil, 错误类.X创建错误码(
			错误码类.CodeNotSupported,
			`Replace operation is not supported by mssql driver`,
		)

	default:
		return d.Core.X底层插入(上下文, 链接, 表名称, list, 选项)
	}
}
