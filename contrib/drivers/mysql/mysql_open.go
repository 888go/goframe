// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package mysql

import (
	"database/sql"
	"fmt"

	"net/url"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gregex"
)

// Open 创建并返回一个底层的 sql.DB 对象，用于 MySQL。
// 注意，它将时间.Time 参数默认转换为本地时区。
// md5:341df118003c304e
func (d *Driver) Open(config *gdb.ConfigNode) (db *sql.DB, err error) {
	var (
		source               string
		underlyingDriverName = "mysql"
	)
	// [用户名[:密码]@][协议[(地址)]]/数据库名[?参数1=值1&...&参数N=值N]. md5:493f32d70b15315e
	if config.Link != "" {
		// ============================================================================
		// 从 v2.2.0 版本开始已废弃。
		// ============================================================================
		// md5:73505fc2089cb531
		source = config.Link
		// 自定义在运行时更改架构。 md5:69ce0e441b271151
		if config.Name != "" {
			source, _ = gregex.ReplaceString(`/([\w\.\-]+)+`, "/"+config.Name, source)
		}
	} else {
		// 待办事项：在未指定字符集时不要设置字符集（在v2.5.0版本中）. md5:2c9a899c402d1e44
		source = fmt.Sprintf(
			"%s:%s@%s(%s:%s)/%s?charset=%s",
			config.User, config.Pass, config.Protocol, config.Host, config.Port, config.Name, config.Charset,
		)
		if config.Timezone != "" {
			if strings.Contains(config.Timezone, "/") {
				config.Timezone = url.QueryEscape(config.Timezone)
			}
			source = fmt.Sprintf("%s&loc=%s", source, config.Timezone)
		}
		if config.Extra != "" {
			source = fmt.Sprintf("%s&%s", source, config.Extra)
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
