// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package clickhouse

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gregex"
)

// Open 创建并返回clickhouse的底层sql.DB对象。 md5:af49366510276559
func (d *Driver) Open(config *gdb.ConfigNode) (db *sql.DB, err error) {
	source := config.Link
	// 使用clickhouse协议，连接到主机host1和host2的9000端口，数据库为'database'。设置连接超时时间为200毫秒，最大执行时间为60秒。其中，'username'和'password'是用户名和密码。 md5:18f0aa9304f1e37e
	if config.Link != "" {
		// ============================================================================
		// 从v2.2.0版本开始已废弃。
		// ============================================================================
		// 自定义在运行时更改架构。
		// md5:636e7e3d9951c8fa
		if config.Name != "" {
			source, _ = gregex.ReplaceString(replaceSchemaPattern, "@$1/"+config.Name, config.Link)
		} else {
			// 如果没有模式，该链接将被用于替换. md5:20ef7460fd455598
			dbName, _ := gregex.MatchString(replaceSchemaPattern, config.Link)
			if len(dbName) > 0 {
				config.Name = dbName[len(dbName)-1]
			}
		}
	} else {
		if config.Pass != "" {
			source = fmt.Sprintf(
				"clickhouse://%s:%s@%s:%s/%s?debug=%t",
				config.User, url.PathEscape(config.Pass),
				config.Host, config.Port, config.Name, config.Debug,
			)
		} else {
			source = fmt.Sprintf(
				"clickhouse://%s@%s:%s/%s?debug=%t",
				config.User, config.Host, config.Port, config.Name, config.Debug,
			)
		}
		if config.Extra != "" {
			source = fmt.Sprintf("%s&%s", source, config.Extra)
		}
	}
	if db, err = sql.Open(driverName, source); err != nil {
		err = gerror.WrapCodef(
			gcode.CodeDbOperationError, err,
			`sql.Open failed for driver "%s" by source "%s"`, driverName, source,
		)
		return nil, err
	}
	return
}
