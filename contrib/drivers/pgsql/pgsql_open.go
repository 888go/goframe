// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package pgsql

import (
	"database/sql"
	"fmt"

	gdb "github.com/888go/goframe/database/gdb"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	gregex "github.com/888go/goframe/text/gregex"
	gstr "github.com/888go/goframe/text/gstr"
)

// X底层Open 创建并返回一个用于pgsql的底层sql.DB对象。
// 参考链接：https://pkg.go.dev/github.com/lib/pq
// md5:9889bcb899248a2b
func (d *Driver) X底层Open(config *gdb.ConfigNode) (db *sql.DB, err error) {
	var (
		source               string
		underlyingDriverName = "postgres"
	)
	if config.Link != "" {
		// ============================================================================
		// 从 v2.2.0 版本开始已废弃。
		// ============================================================================
		// md5:73505fc2089cb531
		source = config.Link
				// 自定义在运行时更改架构。 md5:69ce0e441b271151
		if config.Name != "" {
			source, _ = gregex.X替换文本(`dbname=([\w\.\-]+)+`, "dbname="+config.Name, source)
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
			if extraMap, err = gstr.X参数解析(config.Extra); err != nil {
				return nil, err
			}
			for k, v := range extraMap {
				source += fmt.Sprintf(` %s=%s`, k, v)
			}
		}
	}

	if db, err = sql.Open(underlyingDriverName, source); err != nil {
		err = gerror.X多层错误码并格式化(
			gcode.CodeDbOperationError, err,
			`sql.Open failed for driver "%s" by source "%s"`, underlyingDriverName, source,
		)
		return nil, err
	}
	return
}
