// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package oracle

import (
	"database/sql"
	"strings"

	gora "github.com/sijms/go-ora/v2"

	gdb "github.com/888go/goframe/database/gdb"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	gregex "github.com/888go/goframe/text/gregex"
	gconv "github.com/888go/goframe/util/gconv"
)

// X底层Open 创建并返回一个底层的 sql.DB 对象，针对 Oracle。 md5:db2b73d9e41929bd
func (d *Driver) X底层Open(config *gdb.ConfigNode) (db *sql.DB, err error) {
	var (
		source               string
		underlyingDriverName = "oracle"
	)

	options := map[string]string{
		"CONNECTION TIMEOUT": "60",
		"PREFETCH_ROWS":      "25",
	}

	if config.Debug {
		options["TRACE FILE"] = "oracle_trace.log"
	}
		// [用户名:[密码]@]主机[:端口][/服务名][?参数1=值1&...&参数N=值N]. md5:94680826666597c6
	if config.Link != "" {
		// ============================================================================
		// 从 v2.2.0 版本开始已废弃。
		// ============================================================================
		// md5:73505fc2089cb531
		source = config.Link
				// 自定义在运行时更改架构。 md5:69ce0e441b271151
		if config.Name != "" {
			source, _ = gregex.X替换文本(`@(.+?)/([\w\.\-]+)+`, "@$1/"+config.Name, source)
		}
	} else {
		if config.Extra != "" {
			// fix #3226
			list := strings.Split(config.Extra, "&")
			for _, v := range list {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					options[kv[0]] = kv[1]
				}
			}
		}
		source = gora.BuildUrl(
			config.Host, gconv.X取整数(config.Port), config.Name, config.User, config.Pass, options,
		)
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
