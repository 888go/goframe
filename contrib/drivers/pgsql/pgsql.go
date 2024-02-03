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
	*gdb.Core
}

const (
	internalPrimaryKeyInCtx gctx.StrKey = "primary_key"
	defaultSchema           string      = "public"
	quoteChar               string      = `"`
)

func init() {
	if err := gdb.Register(`pgsql`, New()); err != nil {
		panic(err)
	}
}

// New 创建并返回一个实现了 gdb.Driver 的驱动程序，该驱动程序支持针对 PostgreSql 的操作。
func New() gdb.Driver {
	return &Driver{}
}

// New 创建并返回一个用于 PostgreSQL 的数据库对象。
// 它实现了 gdb.Driver 接口，以便进行额外的数据库驱动安装。
func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &Driver{
		Core: core,
	}, nil
}

// Open 创建并返回一个用于pgsql的底层sql.DB对象。
// 参考文档：https://pkg.go.dev/github.com/lib/pq
func (d *Driver) Open(config *gdb.ConfigNode) (db *sql.DB, err error) {
	var (
		source               string
		underlyingDriverName = "postgres"
	)
	if config.Link != "" {
// ============================================================================
// 从 v2.2.0 版本开始已弃用。
// ============================================================================
		source = config.Link
		// 在运行时自定义更改架构
		if config.Name != "" {
			source, _ = gregex.ReplaceString(`dbname=([\w\.\-]+)+`, "dbname="+config.Name, source)
		}
	} else {
		if config.Name != "" {
			source = fmt.Sprintf(
				"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
				config.User, config.Pass, config.Host, config.Port, config.Name,
			)
		} else {
			source = fmt.Sprintf(
				"user=%s password=%s host=%s port=%s sslmode=disable",
				config.User, config.Pass, config.Host, config.Port,
			)
		}

		if config.Namespace != "" {
			source = fmt.Sprintf("%s search_path=%s", source, config.Namespace)
		}

		if config.Timezone != "" {
			source = fmt.Sprintf("%s timezone=%s", source, config.Timezone)
		}

		if config.Extra != "" {
			var extraMap map[string]interface{}
			if extraMap, err = gstr.Parse(config.Extra); err != nil {
				return nil, err
			}
			for k, v := range extraMap {
				source += fmt.Sprintf(` %s=%s`, k, v)
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

// CheckLocalTypeForField 检查并返回给定数据库类型对应的本地 Go 语言类型。
func (d *Driver) CheckLocalTypeForField(ctx context.Context, fieldType string, fieldValue interface{}) (gdb.LocalType, error) {
	var typeName string
	match, _ := gregex.MatchString(`(.+?)\((.+)\)`, fieldType)
	if len(match) == 3 {
		typeName = gstr.Trim(match[1])
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
		return gdb.LocalTypeInt, nil

	case
		// 对于pgsql，int8 等同于 bigint
		"int8":
		return gdb.LocalTypeInt64, nil

	case
		"_int2",
		"_int4":
		return gdb.LocalTypeIntSlice, nil

	case
		"_int8":
		return gdb.LocalTypeInt64Slice, nil

	default:
		return d.Core.CheckLocalTypeForField(ctx, fieldType, fieldValue)
	}
}

// ConvertValueForLocal 将值根据数据库中的字段类型名称转换为本地 Go 语言类型的值。
// 参数 `fieldType` 为小写形式，例如：
// `float(5,2)`, `unsigned double(5,2)`, `decimal(10,2)`, `char(45)`, `varchar(100)` 等。
func (d *Driver) ConvertValueForLocal(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error) {
	typeName, _ := gregex.ReplaceString(`\(.+\)`, "", fieldType)
	typeName = strings.ToLower(typeName)
	switch typeName {
	// 对于pgsql，int2代表smallint，而int4代表integer。
	case "int2", "int4":
		return gconv.Int(gconv.String(fieldValue)), nil

	// 对于pgsql，int8 等同于 bigint.
	case "int8":
		return gconv.Int64(gconv.String(fieldValue)), nil

	// Int32 slice.
	case
		"_int2", "_int4":
		return gconv.Ints(
			gstr.ReplaceByMap(gconv.String(fieldValue),
				map[string]string{
					"{": "[",
					"}": "]",
				},
			),
		), nil

	// Int64 slice.
	case
		"_int8":
		return gconv.Int64s(
			gstr.ReplaceByMap(gconv.String(fieldValue),
				map[string]string{
					"{": "[",
					"}": "]",
				},
			),
		), nil

	default:
		return d.Core.ConvertValueForLocal(ctx, fieldType, fieldValue)
	}
}

// DoFilter 在将 SQL 字符串提交给底层 SQL 驱动程序之前，对其进行处理。
func (d *Driver) DoFilter(ctx context.Context, link gdb.Link, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error) {
	var index int
	// 将占位符字符 '?' 转换为字符串 "$x"。
	newSql, _ = gregex.ReplaceStringFunc(`\?`, sql, func(s string) string {
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
	newSql, _ = gregex.ReplaceStringFuncMatch(`(::jsonb([^\w\d]*)\$\d)`, newSql, func(match []string) string {
		return fmt.Sprintf(`::jsonb%s?`, match[2])
	})
	newSql, _ = gregex.ReplaceString(` LIMIT (\d+),\s*(\d+)`, ` LIMIT $2 OFFSET $1`, newSql)
	return d.Core.DoFilter(ctx, link, newSql, args)
}

// Tables 获取并返回当前模式的表格。
// 它主要用于cli工具链中，用于自动生成模型。
func (d *Driver) Tables(ctx context.Context, schema ...string) (tables []string, err error) {
	var (
		result     gdb.Result
		usedSchema = gutil.GetOrDefaultStr(d.GetConfig().Namespace, schema...)
	)
	if usedSchema == "" {
		usedSchema = defaultSchema
	}
	// **请勿**将 `usedSchema` 作为参数传递给函数 `SlaveLink`。
	link, err := d.SlaveLink(schema...)
	if err != nil {
		return nil, err
	}

	useRelpartbound := ""
	if gstr.CompareVersion(d.version(ctx, link), "10") >= 0 {
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

	query, _ = gregex.ReplaceString(`[\n\r\s]+`, " ", gstr.Trim(query))
	result, err = d.DoSelect(ctx, link, query)
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

// version 检查并返回数据库版本。
func (d *Driver) version(ctx context.Context, link gdb.Link) string {
	result, err := d.DoSelect(ctx, link, "SELECT version();")
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
func (d *Driver) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*gdb.TableField, err error) {
	var (
		result     gdb.Result
		link       gdb.Link
		usedSchema = gutil.GetOrDefaultStr(d.GetSchema(), schema...)
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
			table,
		)
	)
	if link, err = d.SlaveLink(usedSchema); err != nil {
		return nil, err
	}
	structureSql, _ = gregex.ReplaceString(`[\n\r\s]+`, " ", gstr.Trim(structureSql))
	result, err = d.DoSelect(ctx, link, structureSql)
	if err != nil {
		return nil, err
	}
	fields = make(map[string]*gdb.TableField)
	var (
		index = 0
		name  string
		ok    bool
	)
	for _, m := range result {
		name = m["field"].String()
		// 过滤重复字段。
		if _, ok = fields[name]; ok {
			continue
		}
		fields[name] = &gdb.TableField{
			Index:   index,
			Name:    name,
			Type:    m["type"].String(),
			Null:    !m["null"].Bool(),
			Key:     m["key"].String(),
			Default: m["default_value"].Val(),
			Comment: m["comment"].String(),
		}
		index++
	}
	return fields, nil
}

// DoInsert 对给定表执行插入或更新数据操作。
func (d *Driver) DoInsert(ctx context.Context, link gdb.Link, table string, list gdb.List, option gdb.DoInsertOption) (result sql.Result, err error) {
	switch option.InsertOption {
	case gdb.InsertOptionSave:
		return nil, gerror.NewCode(
			gcode.CodeNotSupported,
			`Save operation is not supported by pgsql driver`,
		)

	case gdb.InsertOptionReplace:
		return nil, gerror.NewCode(
			gcode.CodeNotSupported,
			`Replace operation is not supported by pgsql driver`,
		)

	case gdb.InsertOptionIgnore:
		return nil, gerror.NewCode(
			gcode.CodeNotSupported,
			`Insert ignore operation is not supported by pgsql driver`,
		)

	case gdb.InsertOptionDefault:
		tableFields, err := d.GetCore().GetDB().TableFields(ctx, table)
		if err == nil {
			for _, field := range tableFields {
				if field.Key == "pri" {
					pkField := *field
					ctx = context.WithValue(ctx, internalPrimaryKeyInCtx, pkField)
					break
				}
			}
		}
	}
	return d.Core.DoInsert(ctx, link, table, list, option)
}

// DoExec 通过给定的link对象，将sql字符串及其参数提交到底层驱动，并返回执行结果。
func (d *Driver) DoExec(ctx context.Context, link gdb.Link, sql string, args ...interface{}) (result sql.Result, err error) {
	var (
		isUseCoreDoExec bool   = false // 检查是否需要使用默认方法
		primaryKey      string = ""
		pkField         gdb.TableField
	)

	// 事务检查。
	if link == nil {
		if tx := gdb.TXFromCtx(ctx, d.GetGroup()); tx != nil {
			// 首先，从上下文检查并检索事务链接。
			link = tx
		} else if link, err = d.MasterLink(); err != nil {
			// 或者从主节点创建一个。
			return nil, err
		}
	} else if !link.IsTransaction() {
		// 如果当前链接不是事务链接，则检查并从上下文中检索事务。
		if tx := gdb.TXFromCtx(ctx, d.GetGroup()); tx != nil {
			link = tx
		}
	}

	// 检查是否为主键插入操作。
	if value := ctx.Value(internalPrimaryKeyInCtx); value != nil {
		var ok bool
		pkField, ok = value.(gdb.TableField)
		if !ok {
			isUseCoreDoExec = true
		}
	} else {
		isUseCoreDoExec = true
	}

	// 检查是否为插入操作。
	if !isUseCoreDoExec && pkField.Name != "" && strings.Contains(sql, "INSERT INTO") {
		primaryKey = pkField.Name
		sql += " RETURNING " + primaryKey
	} else {
		// 使用默认的DoExec
		return d.Core.DoExec(ctx, link, sql, args...)
	}

	// 仅当使用主键执行插入操作时，才能执行以下代码

	if d.GetConfig().ExecTimeout > 0 {
		var cancelFunc context.CancelFunc
		ctx, cancelFunc = context.WithTimeout(ctx, d.GetConfig().ExecTimeout)
		defer cancelFunc()
	}

	// Sql filtering.
	sql, args = d.FormatSqlBeforeExecuting(sql, args)
	sql, args, err = d.DoFilter(ctx, link, sql, args)
	if err != nil {
		return nil, err
	}

	// Link execution.
	var out gdb.DoCommitOutput
	out, err = d.DoCommit(ctx, gdb.DoCommitInput{
		Link:          link,
		Sql:           sql,
		Args:          args,
		Stmt:          nil,
		Type:          gdb.SqlTypeQueryContext,
		IsTransaction: link.IsTransaction(),
	})

	if err != nil {
		return nil, err
	}
	affected := len(out.Records)
	if affected > 0 {
		if !strings.Contains(pkField.Type, "int") {
			return Result{
				affected:     int64(affected),
				lastInsertId: 0,
				lastInsertIdError: gerror.NewCodef(
					gcode.CodeNotSupported,
					"LastInsertId is not supported by primary key type: %s", pkField.Type),
			}, nil
		}

		if out.Records[affected-1][primaryKey] != nil {
			lastInsertId := out.Records[affected-1][primaryKey].Int64()
			return Result{
				affected:     int64(affected),
				lastInsertId: lastInsertId,
			}, nil
		}
	}

	return Result{}, nil
}
