// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package mysql 实现了 gdb.Driver 接口，该接口支持对 MySQL 数据库的相关操作。
package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"strings"
	
	_ "github.com/go-sql-driver/mysql"
	
	"github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/util/gutil"
)

// Driver 是 MySQL 数据库的驱动程序。
type Driver struct {
	*db类.Core
}

const (
	quoteChar = "`"
)

func init() {
	var (
		err         error
		driverObj   = New()
		driverNames = g.SliceStr别名{"mysql", "mariadb", "tidb"}
	)
	for _, driverName := range driverNames {
		if err = db类.X注册驱动(driverName, driverObj); err != nil {
			panic(err)
		}
	}
}

// New 创建并返回一个实现 gdb.Driver 接口的驱动程序，该驱动程序支持针对 MySQL 的操作。
func New() db类.Driver {
	return &Driver{}
}

// New 创建并返回一个用于 mysql 的数据库对象。
// 它实现了 gdb.Driver 接口，以便进行额外的数据库驱动安装。
func (d *Driver) New(core *db类.Core, node *db类.X配置项) (db类.DB, error) {
	return &Driver{
		Core: core,
	}, nil
}

// Open 创建并返回一个用于 mysql 的底层 sql.DB 对象。
// 注意，它默认会将 time.Time 类型参数转换为本地时区。
func (d *Driver) X底层Open(配置对象 *db类.X配置项) (db *sql.DB, err error) {
	var (
		source               string
		underlyingDriverName = "mysql"
	)
	// [用户名[:密码]@][协议[(地址)]]/数据库名[?参数1=值1&...&参数N=值N]
// 这段注释是对Go语言中符合MySQL连接格式的字符串进行描述，具体含义如下：
// - `[username[:password]@]`：可选的用户名和密码部分，用于登录数据库。冒号（:）分隔用户名和密码。
// - `[protocol[(address)]]`：指定数据库连接协议以及服务器地址，例如 `tcp(` 或 `unix(` 等，其中括号内的 `address` 为服务器地址或socket路径。
// - `/dbname`：必填项，表示要连接的数据库名称。
// - `[?param1=value1&...&paramN=valueN]`：可选的查询参数部分，通常用于设置额外的连接选项，如 `charset=utf8mb4`、`parseTime=true` 等，多个参数之间用 `&` 符号分隔。
	if 配置对象.X自定义链接信息 != "" {
// ============================================================================
// 从 v2.2.0 版本开始已弃用。
// ============================================================================
		source = 配置对象.X自定义链接信息
		// 自定义在运行时更改架构
		if 配置对象.X名称 != "" {
			source, _ = 正则类.X替换文本(`/([\w\.\-]+)+`, "/"+配置对象.X名称, source)
		}
	} else {
		// TODO: 当未指定字符集时（在v2.5.0版本中），不要设置字符集
		source = fmt.Sprintf(
			"%s:%s@%s(%s:%s)/%s?charset=%s",
			配置对象.X账号, 配置对象.X密码, 配置对象.X协议, 配置对象.X地址, 配置对象.X端口, 配置对象.X名称, 配置对象.X字符集,
		)
		if 配置对象.X时区 != "" {
			if strings.Contains(配置对象.X时区, "/") {
				配置对象.X时区 = url.QueryEscape(配置对象.X时区)
			}
			source = fmt.Sprintf("%s&loc=%s", source, 配置对象.X时区)
		}
		if 配置对象.X额外 != "" {
			source = fmt.Sprintf("%s&%s", source, 配置对象.X额外)
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

// DoFilter 在将 SQL 发送给数据库之前处理 SQL。
func (d *Driver) X底层DoFilter(ctx context.Context, link db类.Link, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error) {
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
	result, 错误 = d.X底层查询(上下文, link, `SHOW TABLES`)
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
// 参数`link`是可选的，如果给定为nil，它会自动获取一个原始sql连接作为链接执行必要的sql查询。
//
// 注意，它返回一个包含字段名及其对应字段信息的map。由于map是无序的，TableField结构体中有一个"Index"字段标记其在所有字段中的顺序。
//
// 为了提高性能，该方法使用了缓存功能，缓存有效期直到进程重启才会失效。
func (d *Driver) X取表字段信息Map(上下文 context.Context, 表名称 string, schema ...string) (字段信息Map map[string]*db类.TableField, err error) {
	var (
		result     db类.Result
		link       db类.Link
		usedSchema = 工具类.X取文本值或取默认值(d.X取默认数据库名称(), schema...)
	)
	if link, err = d.X底层SlaveLink(usedSchema); err != nil {
		return nil, err
	}
	result, err = d.X底层查询(
		上下文, link,
		fmt.Sprintf(`SHOW FULL COLUMNS FROM %s`, d.X底层QuoteWord(表名称)),
	)
	if err != nil {
		return nil, err
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
