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
	*gdb.Core
}

const (
	quoteChar = `"`
)

func init() {
	if err := gdb.Register(`mssql`, New()); err != nil {
		panic(err)
	}
}

// New 创建并返回一个实现 gdb.Driver 接口的驱动程序，该驱动支持对 Mssql 的操作。
func New() gdb.Driver {
	return &Driver{}
}

// New 创建并返回一个用于SQL服务器的数据库对象。
// 它实现了gdb.Driver接口，以便支持额外的数据库驱动安装。
func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &Driver{
		Core: core,
	}, nil
}

// Open 创建并返回一个用于mssql的底层sql.DB对象。
func (d *Driver) Open(config *gdb.ConfigNode) (db *sql.DB, err error) {
	var (
		source               string
		underlyingDriverName = "sqlserver"
	)
	if config.Link != "" {
// ============================================================================
// 从 v2.2.0 版本开始已弃用。
// ============================================================================
		source = config.Link
		// 在运行时自定义更改架构
		if config.Name != "" {
			source, _ = gregex.ReplaceString(`database=([\w\.\-]+)+`, "database="+config.Name, source)
		}
	} else {
		source = fmt.Sprintf(
			"user id=%s;password=%s;server=%s;port=%s;database=%s;encrypt=disable",
			config.User, config.Pass, config.Host, config.Port, config.Name,
		)
		if config.Extra != "" {
			var extraMap map[string]interface{}
			if extraMap, err = gstr.Parse(config.Extra); err != nil {
				return nil, err
			}
			for k, v := range extraMap {
				source += fmt.Sprintf(`;%s=%s`, k, v)
			}
		}
	}

	if db, err = sql.Open(underlyingDriverName, source); err != nil {
		err = gerror.WrapCodef(
			gcode.CodeDbOperationError, err,
			`sql.Open failed for driver "%s" by source "%s"`, underlyingDriverName, source,
		)
		return nil, err
	}
	return
}

// GetChars 返回该类型数据库的安全字符。
func (d *Driver) GetChars() (charLeft string, charRight string) {
	return quoteChar, quoteChar
}

// DoFilter 在将 SQL 字符串提交给底层 SQL 驱动程序之前，对其进行处理。
func (d *Driver) DoFilter(ctx context.Context, link gdb.Link, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error) {
	var index int
	// 将占位符字符 '?' 转换为字符串 "@px"。
	newSql, _ = gregex.ReplaceStringFunc("\\?", sql, func(s string) string {
		index++
		return fmt.Sprintf("@p%d", index)
	})
	newSql, _ = gregex.ReplaceString("\"", "", newSql)
	return d.Core.DoFilter(ctx, link, d.parseSql(newSql), args)
}

// parseSql在将SQL提交给底层驱动程序之前执行一些替换操作，
// 以便支持Microsoft SQL Server。
func (d *Driver) parseSql(sql string) string {
	// 从USER表中选取ID为1的记录，限制返回结果数量为1条
	if m, _ := gregex.MatchString(`^SELECT(.+)LIMIT 1$`, sql); len(m) > 1 {
		return fmt.Sprintf(`SELECT TOP 1 %s`, m[1])
	}
	// 从USER表中选取AGE大于18的所有列，并按ID降序排列，然后获取第101至300条记录（LIMIT offset, count语法）
	patten := `^\s*(?i)(SELECT)|(LIMIT\s*(\d+)\s*,\s*(\d+))`
	if gregex.IsMatchString(patten, sql) == false {
		return sql
	}
	res, err := gregex.MatchAllString(patten, sql)
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
		if gregex.IsMatchString("((?i)SELECT)(.+)((?i)LIMIT)", sql) == false {
			break
		}
		// ORDER BY 语句检查。
		var (
			selectStr = ""
			orderStr  = ""
			haveOrder = gregex.IsMatchString("((?i)SELECT)(.+)((?i)ORDER BY)", sql)
		)
		if haveOrder {
			queryExpr, _ := gregex.MatchString("((?i)SELECT)(.+)((?i)ORDER BY)", sql)
			if len(queryExpr) != 4 ||
				strings.EqualFold(queryExpr[1], "SELECT") == false ||
				strings.EqualFold(queryExpr[3], "ORDER BY") == false {
				break
			}
			selectStr = queryExpr[2]
			orderExpr, _ := gregex.MatchString("((?i)ORDER BY)(.+)((?i)LIMIT)", sql)
			if len(orderExpr) != 4 ||
				strings.EqualFold(orderExpr[1], "ORDER BY") == false ||
				strings.EqualFold(orderExpr[3], "LIMIT") == false {
				break
			}
			orderStr = orderExpr[2]
		} else {
			queryExpr, _ := gregex.MatchString("((?i)SELECT)(.+)((?i)LIMIT)", sql)
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
func (d *Driver) Tables(ctx context.Context, schema ...string) (tables []string, err error) {
	var result gdb.Result
	link, err := d.SlaveLink(schema...)
	if err != nil {
		return nil, err
	}

	result, err = d.DoSelect(
		ctx, link, `SELECT NAME FROM SYSOBJECTS WHERE XTYPE='U' AND STATUS >= 0 ORDER BY NAME`,
	)
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

// TableFields 获取并返回当前模式下指定表的字段信息。
//
// 另请参阅 DriverMysql.TableFields。
func (d *Driver) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*gdb.TableField, err error) {
	var (
		result     gdb.Result
		link       gdb.Link
		usedSchema = gutil.GetOrDefaultStr(d.GetSchema(), schema...)
	)
	if link, err = d.SlaveLink(usedSchema); err != nil {
		return nil, err
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
		table,
	)
	structureSql, _ = gregex.ReplaceString(`[\n\r\s]+`, " ", gstr.Trim(structureSql))
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

// DoInsert 对给定表执行插入或更新数据操作。
func (d *Driver) DoInsert(ctx context.Context, link gdb.Link, table string, list gdb.List, option gdb.DoInsertOption) (result sql.Result, err error) {
	switch option.InsertOption {
	case gdb.InsertOptionSave:
		return nil, gerror.NewCode(
			gcode.CodeNotSupported,
			`Save operation is not supported by mssql driver`,
		)

	case gdb.InsertOptionReplace:
		return nil, gerror.NewCode(
			gcode.CodeNotSupported,
			`Replace operation is not supported by mssql driver`,
		)

	default:
		return d.Core.DoInsert(ctx, link, table, list, option)
	}
}
