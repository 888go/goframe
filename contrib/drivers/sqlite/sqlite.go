// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package sqlite 实现了 gdb.Driver 接口，该接口支持对 SQLite 数据库的操作。
//
// 注意：
// 1. 此包不支持 Save 功能。
package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	
	_ "github.com/glebarez/go-sqlite"
	
	"github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/encoding/gurl"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
)

// Driver 是用于 SQLite 数据库的驱动程序。
type Driver struct {
	*db类.Core
}

const (
	quoteChar = "`"
)

func init() {
	if err := db类.X注册驱动(`sqlite`, New()); err != nil {
		panic(err)
	}
}

// New 创建并返回一个实现了 gdb.Driver 接口的驱动器，该驱动器支持对 SQLite 的操作。
func New() db类.Driver {
	return &Driver{}
}

// New 创建并返回一个用于 sqlite 的数据库对象。
// 它实现了 gdb.Driver 接口，以便进行额外的数据库驱动安装。
func (d *Driver) New(core *db类.Core, node *db类.ConfigNode) (db类.DB, error) {
	return &Driver{
		Core: core,
	}, nil
}

// Open创建并返回一个用于sqlite的底层sql.DB对象。
// 参考链接：https://github.com/glebarez/go-sqlite
func (d *Driver) X底层Open(配置对象 *db类.ConfigNode) (db *sql.DB, err error) {
	var (
		source               string
		underlyingDriverName = "sqlite"
	)
	if 配置对象.Link != "" {
// ============================================================================
// 从 v2.2.0 版本开始已弃用。
// ============================================================================
		source = 配置对象.Link
	} else {
		source = 配置对象.Name
	}
	// 它在源文件中搜索以定位其绝对路径。
	if absolutePath, _ := 文件类.X查找(source); absolutePath != "" {
		source = absolutePath
	}

// 多个PRAGMA指令可以通过如下方式指定：
// path/to/some.db?_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)
// （翻译成中文）
// 可以通过以下方式同时指定多个PRAGMA参数：
// path/to/some.db?_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)
// 其中，"busy_timeout"设置为5000毫秒，"journal_mode"设置为WAL模式。
	if 配置对象.Extra != "" {
		var (
			options  string
			extraMap map[string]interface{}
		)
		if extraMap, err = 文本类.X参数解析(配置对象.Extra); err != nil {
			return nil, err
		}
		for k, v := range extraMap {
			if options != "" {
				options += "&"
			}
			options += fmt.Sprintf(`_pragma=%s(%s)`, k, url类.X编码(转换类.String(v)))
		}
		if len(options) > 1 {
			source += "?" + options
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

// GetChars 返回此类型数据库的安全字符。
func (d *Driver) X底层取数据库安全字符() (左字符 string, 右字符 string) {
	return quoteChar, quoteChar
}

// DoFilter 在将SQL字符串提交给底层SQL驱动程序之前对其进行处理。
func (d *Driver) X底层DoFilter(ctx context.Context, link db类.Link, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error) {
	// 特殊的插入/忽略操作，用于SQLite.
	switch {
	case 文本类.X开头判断(sql, db类.InsertOperationIgnore):
		sql = "INSERT OR IGNORE" + sql[len(db类.InsertOperationIgnore):]

	case 文本类.X开头判断(sql, db类.InsertOperationReplace):
		sql = "INSERT OR REPLACE" + sql[len(db类.InsertOperationReplace):]

	default:
		if 文本类.X是否包含(sql, db类.InsertOnDuplicateKeyUpdate) {
			return sql, args, 错误类.X创建错误码(
				错误码类.CodeNotSupported,
				`Save operation is not supported by sqlite driver`,
			)
		}
	}
	return d.Core.X底层DoFilter(ctx, link, sql, args)
}

// Tables 获取并返回当前模式的表。
// 它主要用于cli工具链中，用于自动生成模型。
func (d *Driver) X取表名称数组(上下文 context.Context, schema ...string) (表名称数组 []string, 错误 error) {
	var result db类.Result
	link, 错误 := d.X底层SlaveLink(schema...)
	if 错误 != nil {
		return nil, 错误
	}

	result, 错误 = d.X底层查询(
		上下文,
		link,
		`SELECT NAME FROM SQLITE_MASTER WHERE TYPE='table' ORDER BY NAME`,
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
func (d *Driver) X取表字段信息Map(上下文 context.Context, 表名称 string, schema ...string) (字段信息Map map[string]*db类.TableField, err error) {
	var (
		result     db类.Result
		link       db类.Link
		usedSchema = 工具类.X取文本值或取默认值(d.X取默认数据库名称(), schema...)
	)
	if link, err = d.X底层SlaveLink(usedSchema); err != nil {
		return nil, err
	}
	result, err = d.X底层查询(上下文, link, fmt.Sprintf(`PRAGMA TABLE_INFO(%s)`, d.X底层QuoteWord(表名称)))
	if err != nil {
		return nil, err
	}
	字段信息Map = make(map[string]*db类.TableField)
	for i, m := range result {
		mKey := ""
		if m["pk"].X取布尔() {
			mKey = "pri"
		}
		字段信息Map[m["name"].String()] = &db类.TableField{
			Index:   i,
			Name:    m["name"].String(),
			Type:    m["type"].String(),
			Key:     mKey,
			Default: m["dflt_value"].X取值(),
			Null:    !m["notnull"].X取布尔(),
		}
	}
	return 字段信息Map, nil
}
