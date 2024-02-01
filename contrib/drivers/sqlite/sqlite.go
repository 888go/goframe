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
	*gdb.Core
}

const (
	quoteChar = "`"
)

func init() {
	if err := gdb.Register(`sqlite`, New()); err != nil {
		panic(err)
	}
}

// New 创建并返回一个实现了 gdb.Driver 接口的驱动器，该驱动器支持对 SQLite 的操作。
func New() gdb.Driver {
	return &Driver{}
}

// New 创建并返回一个用于 sqlite 的数据库对象。
// 它实现了 gdb.Driver 接口，以便进行额外的数据库驱动安装。
func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &Driver{
		Core: core,
	}, nil
}

// Open创建并返回一个用于sqlite的底层sql.DB对象。
// 参考链接：https://github.com/glebarez/go-sqlite
func (d *Driver) Open(config *gdb.ConfigNode) (db *sql.DB, err error) {
	var (
		source               string
		underlyingDriverName = "sqlite"
	)
	if config.Link != "" {
// ============================================================================
// 从 v2.2.0 版本开始已弃用。
// ============================================================================
		source = config.Link
	} else {
		source = config.Name
	}
	// 它在源文件中搜索以定位其绝对路径。
	if absolutePath, _ := gfile.Search(source); absolutePath != "" {
		source = absolutePath
	}

// 多个PRAGMA指令可以通过如下方式指定：
// path/to/some.db?_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)
// （翻译成中文）
// 可以通过以下方式同时指定多个PRAGMA参数：
// path/to/some.db?_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)
// 其中，"busy_timeout"设置为5000毫秒，"journal_mode"设置为WAL模式。
	if config.Extra != "" {
		var (
			options  string
			extraMap map[string]interface{}
		)
		if extraMap, err = gstr.Parse(config.Extra); err != nil {
			return nil, err
		}
		for k, v := range extraMap {
			if options != "" {
				options += "&"
			}
			options += fmt.Sprintf(`_pragma=%s(%s)`, k, gurl.Encode(gconv.String(v)))
		}
		if len(options) > 1 {
			source += "?" + options
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

// GetChars 返回此类型数据库的安全字符。
func (d *Driver) GetChars() (charLeft string, charRight string) {
	return quoteChar, quoteChar
}

// DoFilter 在将SQL字符串提交给底层SQL驱动程序之前对其进行处理。
func (d *Driver) DoFilter(ctx context.Context, link gdb.Link, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error) {
	// 特殊的插入/忽略操作，用于SQLite.
	switch {
	case gstr.HasPrefix(sql, gdb.InsertOperationIgnore):
		sql = "INSERT OR IGNORE" + sql[len(gdb.InsertOperationIgnore):]

	case gstr.HasPrefix(sql, gdb.InsertOperationReplace):
		sql = "INSERT OR REPLACE" + sql[len(gdb.InsertOperationReplace):]

	default:
		if gstr.Contains(sql, gdb.InsertOnDuplicateKeyUpdate) {
			return sql, args, gerror.NewCode(
				gcode.CodeNotSupported,
				`Save operation is not supported by sqlite driver`,
			)
		}
	}
	return d.Core.DoFilter(ctx, link, sql, args)
}

// Tables 获取并返回当前模式的表。
// 它主要用于cli工具链中，用于自动生成模型。
func (d *Driver) Tables(ctx context.Context, schema ...string) (tables []string, err error) {
	var result gdb.Result
	link, err := d.SlaveLink(schema...)
	if err != nil {
		return nil, err
	}

	result, err = d.DoSelect(
		ctx,
		link,
		`SELECT NAME FROM SQLITE_MASTER WHERE TYPE='table' ORDER BY NAME`,
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
	result, err = d.DoSelect(ctx, link, fmt.Sprintf(`PRAGMA TABLE_INFO(%s)`, d.QuoteWord(table)))
	if err != nil {
		return nil, err
	}
	fields = make(map[string]*gdb.TableField)
	for i, m := range result {
		mKey := ""
		if m["pk"].Bool() {
			mKey = "pri"
		}
		fields[m["name"].String()] = &gdb.TableField{
			Index:   i,
			Name:    m["name"].String(),
			Type:    m["type"].String(),
			Key:     mKey,
			Default: m["dflt_value"].Val(),
			Null:    !m["notnull"].Bool(),
		}
	}
	return fields, nil
}
