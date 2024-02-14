// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package pgsql 实现了 gdb.Driver 接口，该接口支持对 PostgreSQL 数据库进行操作。
// 注意：
// 1. 该包不支持 Save/Replace 功能。
// 2. 该包不支持 Insert Ignore 特性。
package pgsql

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strings"
	
	_ "github.com/lib/pq"
	
	"github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/os/gctx"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
)

// Driver 是用于 PostgreSQL 数据库的驱动。
type Driver struct {
	*db类.Core
}

const (
	internalPrimaryKeyInCtx 上下文类.StrKey = "primary_key"
	defaultSchema           string      = "public"
	quoteChar               string      = `"`
)

func init() {
	if err := db类.X注册驱动(`pgsql`, New()); err != nil {
		panic(err)
	}
}

// New 创建并返回一个实现了 gdb.Driver 的驱动程序，该驱动程序支持针对 PostgreSql 的操作。
func New() db类.Driver {
	return &Driver{}
}

// New 创建并返回一个用于 PostgreSQL 的数据库对象。
// 它实现了 gdb.Driver 接口，以便进行额外的数据库驱动安装。
func (d *Driver) New(core *db类.Core, node *db类.X配置项) (db类.DB, error) {
	return &Driver{
		Core: core,
	}, nil
}

// Open 创建并返回一个用于pgsql的底层sql.DB对象。
// 参考文档：https://pkg.go.dev/github.com/lib/pq
func (d *Driver) X底层Open(配置对象 *db类.X配置项) (db *sql.DB, err error) {
	var (
		source               string
		underlyingDriverName = "postgres"
	)
	if 配置对象.X自定义链接信息 != "" {
// ============================================================================
// 从 v2.2.0 版本开始已弃用。
// ============================================================================
		source = 配置对象.X自定义链接信息
		// 在运行时自定义更改架构
		if 配置对象.X名称 != "" {
			source, _ = 正则类.X替换文本(`dbname=([\w\.\-]+)+`, "dbname="+配置对象.X名称, source)
		}
	} else {
		if 配置对象.X名称 != "" {
			source = fmt.Sprintf(
				"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
				配置对象.X账号, 配置对象.X密码, 配置对象.X地址, 配置对象.X端口, 配置对象.X名称,
			)
		} else {
			source = fmt.Sprintf(
				"user=%s password=%s host=%s port=%s sslmode=disable",
				配置对象.X账号, 配置对象.X密码, 配置对象.X地址, 配置对象.X端口,
			)
		}

		if 配置对象.X命名空间 != "" {
			source = fmt.Sprintf("%s search_path=%s", source, 配置对象.X命名空间)
		}

		if 配置对象.X时区 != "" {
			source = fmt.Sprintf("%s timezone=%s", source, 配置对象.X时区)
		}

		if 配置对象.X额外 != "" {
			var extraMap map[string]interface{}
			if extraMap, err = 文本类.X参数解析(配置对象.X额外); err != nil {
				return nil, err
			}
			for k, v := range extraMap {
				source += fmt.Sprintf(` %s=%s`, k, v)
			}
		}
	}

	if db, err = sql.Open(underlyingDriverName, source); err != nil {
		err = 错误类.X多层错误码并格式化(
			错误码类.CodeDbOperationError, err,
			`sql.Open failed for driver "%s" by source "%s"`, underlyingDriverName, source,
		)
		return nil, err
	}
	return
}

// GetChars 返回该类型数据库的安全字符。
func (d *Driver) X底层取数据库安全字符() (左字符 string, 右字符 string) {
	return quoteChar, quoteChar
}

// CheckLocalTypeForField 检查并返回给定数据库类型对应的本地 Go 语言类型。
func (d *Driver) X底层CheckLocalTypeForField(ctx context.Context, fieldType string, fieldValue interface{}) (db类.LocalType, error) {
	var typeName string
	match, _ := 正则类.X匹配文本(`(.+?)\((.+)\)`, fieldType)
	if len(match) == 3 {
		typeName = 文本类.X过滤首尾符并含空白(match[1])
	} else {
		typeName = fieldType
	}
	typeName = strings.ToLower(typeName)
	switch typeName {
	case
		// 对于pgsql，int2等于smallint。
		"int2",
		// 对于pgsql，int4等于integer（整数类型）
		"int4":
		return db类.LocalTypeInt, nil

	case
		// 对于pgsql，int8 等同于 bigint
		"int8":
		return db类.LocalTypeInt64, nil

	case
		"_int2",
		"_int4":
		return db类.LocalTypeIntSlice, nil

	case
		"_int8":
		return db类.LocalTypeInt64Slice, nil

	default:
		return d.Core.X底层CheckLocalTypeForField(ctx, fieldType, fieldValue)
	}
}

// ConvertValueForLocal 将值根据数据库中的字段类型名称转换为本地 Go 语言类型的值。
// 参数 `fieldType` 为小写形式，例如：
// `float(5,2)`, `unsigned double(5,2)`, `decimal(10,2)`, `char(45)`, `varchar(100)` 等。
func (d *Driver) X底层ConvertValueForLocal(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error) {
	typeName, _ := 正则类.X替换文本(`\(.+\)`, "", fieldType)
	typeName = strings.ToLower(typeName)
	switch typeName {
	// 对于pgsql，int2代表smallint，而int4代表integer。
	case "int2", "int4":
		return 转换类.X取整数(转换类.String(fieldValue)), nil

	// 对于pgsql，int8 等同于 bigint.
	case "int8":
		return 转换类.X取整数64位(转换类.String(fieldValue)), nil

	// Int32 slice.
	case
		"_int2", "_int4":
		return 转换类.X取整数数组(
			文本类.Map替换(转换类.String(fieldValue),
				map[string]string{
					"{": "[",
					"}": "]",
				},
			),
		), nil

	// Int64 slice.
	case
		"_int8":
		return 转换类.X取整数64位数组(
			文本类.Map替换(转换类.String(fieldValue),
				map[string]string{
					"{": "[",
					"}": "]",
				},
			),
		), nil

	default:
		return d.Core.X底层ConvertValueForLocal(ctx, fieldType, fieldValue)
	}
}

// DoFilter 在将 SQL 字符串提交给底层 SQL 驱动程序之前，对其进行处理。
func (d *Driver) X底层DoFilter(ctx context.Context, link db类.Link, sql string, 参数 []interface{}) (newSql string, newArgs []interface{}, err error) {
	var index int
	// 将占位符字符 '?' 转换为字符串 "$x"。
	newSql, _ = 正则类.X替换文本_函数(`\?`, sql, func(s string) string {
		index++
		return fmt.Sprintf(`$%d`, index)
	})
// 处理pgsql对jsonb特性的支持，其中包含占位符字符 '?'。
// 参考：
// https://github.com/gogf/gf/issues/1537
// https://www.postgresql.org/docs/12/functions-json.html
// 这段Go语言代码的注释翻译成中文后，其含义为：
// 该处用于处理PostgreSQL中对jsonb类型功能的支持，这些功能可能包含问号（'?'）作为占位符字符。
// 参考文档：
// GitHub上gf框架的issue #1537
// PostgreSQL官方文档中关于12版本的JSON函数介绍
	newSql, _ = 正则类.ReplaceStringFuncMatch(`(::jsonb([^\w\d]*)\$\d)`, newSql, func(match []string) string {
		return fmt.Sprintf(`::jsonb%s?`, match[2])
	})
	newSql, _ = 正则类.X替换文本(` LIMIT (\d+),\s*(\d+)`, ` LIMIT $2 OFFSET $1`, newSql)
	return d.Core.X底层DoFilter(ctx, link, newSql, 参数)
}

// Tables 获取并返回当前模式的表格。
// 它主要用于cli工具链中，用于自动生成模型。
func (d *Driver) X取表名称数组(上下文 context.Context, schema ...string) (表名称数组 []string, 错误 error) {
	var (
		result     db类.Result
		usedSchema = 工具类.X取文本值或取默认值(d.X取当前节点配置().X命名空间, schema...)
	)
	if usedSchema == "" {
		usedSchema = defaultSchema
	}
	// **请勿**将 `usedSchema` 作为参数传递给函数 `SlaveLink`。
	link, 错误 := d.X底层SlaveLink(schema...)
	if 错误 != nil {
		return nil, 错误
	}

	useRelpartbound := ""
	if 文本类.X版本号比较GNU格式(d.version(上下文, link), "10") >= 0 {
		useRelpartbound = "AND c.relpartbound IS NULL"
	}

	var query = fmt.Sprintf(`
SELECT
	c.relname
FROM
	pg_class c
INNER JOIN pg_namespace n ON
	c.relnamespace = n.oid
WHERE
	n.nspname = '%s'
	AND c.relkind IN ('r', 'p')
	%s
ORDER BY
	c.relname`,
		usedSchema,
		useRelpartbound,
	)

	query, _ = 正则类.X替换文本(`[\n\r\s]+`, " ", 文本类.X过滤首尾符并含空白(query))
	result, 错误 = d.X底层查询(上下文, link, query)
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

// version 检查并返回数据库版本。
func (d *Driver) version(ctx context.Context, link db类.Link) string {
	result, err := d.X底层查询(ctx, link, "SELECT version();")
	if err != nil {
		return ""
	}
	if len(result) > 0 {
		if v, ok := result[0]["version"]; ok {
			matches := regexp.MustCompile(`PostgreSQL (\d+\.\d+)`).FindStringSubmatch(v.String())
			if len(matches) >= 2 {
				return matches[1]
			}
		}
	}
	return ""
}

// TableFields 获取并返回当前模式下指定表的字段信息。
func (d *Driver) X取表字段信息Map(上下文 context.Context, 表名称 string, schema ...string) (字段信息Map map[string]*db类.TableField, err error) {
	var (
		result     db类.Result
		link       db类.Link
		usedSchema = 工具类.X取文本值或取默认值(d.X取默认数据库名称(), schema...)
		// TODO 是否存在重复的`id`结果？
		structureSql = fmt.Sprintf(`
SELECT a.attname AS field, t.typname AS type,a.attnotnull as null,
    (case when d.contype is not null then 'pri' else '' end)  as key
      ,ic.column_default as default_value,b.description as comment
      ,coalesce(character_maximum_length, numeric_precision, -1) as length
      ,numeric_scale as scale
FROM pg_attribute a
         left join pg_class c on a.attrelid = c.oid
         left join pg_constraint d on d.conrelid = c.oid and a.attnum = d.conkey[1]
         left join pg_description b ON a.attrelid=b.objoid AND a.attnum = b.objsubid
         left join pg_type t ON a.atttypid = t.oid
         left join information_schema.columns ic on ic.column_name = a.attname and ic.table_name = c.relname
WHERE c.relname = '%s' and a.attisdropped is false and a.attnum > 0
ORDER BY a.attnum`,
			表名称,
		)
	)
	if link, err = d.X底层SlaveLink(usedSchema); err != nil {
		return nil, err
	}
	structureSql, _ = 正则类.X替换文本(`[\n\r\s]+`, " ", 文本类.X过滤首尾符并含空白(structureSql))
	result, err = d.X底层查询(上下文, link, structureSql)
	if err != nil {
		return nil, err
	}
	字段信息Map = make(map[string]*db类.TableField)
	var (
		index = 0
		name  string
		ok    bool
	)
	for _, m := range result {
		name = m["field"].String()
		// 过滤重复字段。
		if _, ok = 字段信息Map[name]; ok {
			continue
		}
		字段信息Map[name] = &db类.TableField{
			Index:   index,
			X名称:    name,
			X类型:    m["type"].String(),
			Null:    !m["null"].X取布尔(),
			Key:     m["key"].String(),
			Default: m["default_value"].X取值(),
			Comment: m["comment"].String(),
		}
		index++
	}
	return 字段信息Map, nil
}

// DoInsert 对给定表执行插入或更新数据操作。
func (d *Driver) X底层插入(上下文 context.Context, 链接 db类.Link, 表名称 string, list db类.Map数组, 选项 db类.DoInsertOption) (结果 sql.Result, 错误 error) {
	switch 选项.InsertOption {
	case db类.InsertOptionSave:
		return nil, 错误类.X创建错误码(
			错误码类.CodeNotSupported,
			`Save operation is not supported by pgsql driver`,
		)

	case db类.InsertOptionReplace:
		return nil, 错误类.X创建错误码(
			错误码类.CodeNotSupported,
			`Replace operation is not supported by pgsql driver`,
		)

	case db类.InsertOptionIgnore:
		return nil, 错误类.X创建错误码(
			错误码类.CodeNotSupported,
			`Insert ignore operation is not supported by pgsql driver`,
		)

	case db类.InsertOptionDefault:
		tableFields, err := d.X取Core对象().X取DB对象().X取表字段信息Map(上下文, 表名称)
		if err == nil {
			for _, field := range tableFields {
				if field.Key == "pri" {
					pkField := *field
					上下文 = context.WithValue(上下文, internalPrimaryKeyInCtx, pkField)
					break
				}
			}
		}
	}
	return d.Core.X底层插入(上下文, 链接, 表名称, list, 选项)
}

// DoExec 通过给定的link对象，将sql字符串及其参数提交到底层驱动，并返回执行结果。
func (d *Driver) X底层原生SQL执行(上下文 context.Context, 链接 db类.Link, sql string, 参数 ...interface{}) (结果 sql.Result, 错误 error) {
	var (
		isUseCoreDoExec bool   = false // 检查是否需要使用默认方法
		primaryKey      string = ""
		pkField         db类.TableField
	)

	// 事务检查。
	if 链接 == nil {
		if tx := db类.X事务从上下文取对象(上下文, d.X取配置组名称()); tx != nil {
			// 首先，从上下文检查并检索事务链接。
			链接 = tx
		} else if 链接, 错误 = d.X底层MasterLink(); 错误 != nil {
			// 或者从主节点创建一个。
			return nil, 错误
		}
	} else if !链接.IsTransaction() {
		// 如果当前链接不是事务链接，则检查并从上下文中检索事务。
		if tx := db类.X事务从上下文取对象(上下文, d.X取配置组名称()); tx != nil {
			链接 = tx
		}
	}

	// 检查是否为主键插入操作。
	if value := 上下文.Value(internalPrimaryKeyInCtx); value != nil {
		var ok bool
		pkField, ok = value.(db类.TableField)
		if !ok {
			isUseCoreDoExec = true
		}
	} else {
		isUseCoreDoExec = true
	}

	// 检查是否为插入操作。
	if !isUseCoreDoExec && pkField.X名称 != "" && strings.Contains(sql, "INSERT INTO") {
		primaryKey = pkField.X名称
		sql += " RETURNING " + primaryKey
	} else {
		// 使用默认的DoExec
		return d.Core.X底层原生SQL执行(上下文, 链接, sql, 参数...)
	}

	// 仅当使用主键执行插入操作时，才能执行以下代码

	if d.X取当前节点配置().X执行超时时长 > 0 {
		var cancelFunc context.CancelFunc
		上下文, cancelFunc = context.WithTimeout(上下文, d.X取当前节点配置().X执行超时时长)
		defer cancelFunc()
	}

	// Sql filtering.
	sql, 参数 = d.X格式化Sql(sql, 参数)
	sql, 参数, 错误 = d.X底层DoFilter(上下文, 链接, sql, 参数)
	if 错误 != nil {
		return nil, 错误
	}

	// Link execution.
	var out db类.X输出
	out, 错误 = d.X底层DoCommit(上下文, db类.DoCommitInput{
		Link:          链接,
		Sql:           sql,
		Args:          参数,
		Stmt:          nil,
		X类型:          db类.SqlTypeQueryContext,
		IsTransaction: 链接.IsTransaction(),
	})

	if 错误 != nil {
		return nil, 错误
	}
	affected := len(out.X行记录数组)
	if affected > 0 {
		if !strings.Contains(pkField.X类型, "int") {
			return Result{
				affected:     int64(affected),
				lastInsertId: 0,
				lastInsertIdError: 错误类.X创建错误码并格式化(
					错误码类.CodeNotSupported,
					"LastInsertId is not supported by primary key type: %s", pkField.X类型),
			}, nil
		}

		if out.X行记录数组[affected-1][primaryKey] != nil {
			lastInsertId := out.X行记录数组[affected-1][primaryKey].X取整数64位()
			return Result{
				affected:     int64(affected),
				lastInsertId: lastInsertId,
			}, nil
		}
	}

	return Result{}, nil
}
