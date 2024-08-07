// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package sqlite

import (
	"database/sql"
	"fmt"

	gdb "github.com/888go/goframe/database/gdb"
	gurl "github.com/888go/goframe/encoding/gurl"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	gfile "github.com/888go/goframe/os/gfile"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

// X底层Open 创建并返回一个底层的 sql.DB 对象，用于 SQLite。
// https://github.com/glebarez/go-sqlite
// md5:f3c66f09b7236ffa
func (d *Driver) X底层Open(config *gdb.ConfigNode) (db *sql.DB, err error) {
	var (
		source               string
		underlyingDriverName = "sqlite"
	)
	if config.Link != "" {
		// ============================================================================
		// 从 v2.2.0 版本开始已废弃。
		// ============================================================================
		// md5:73505fc2089cb531
		source = config.Link
	} else {
		source = config.Name
	}
		// 它会搜索源文件以找到其绝对路径。 md5:c47257e12a6e34cf
	if absolutePath, _ := gfile.X查找(source); absolutePath != "" {
		source = absolutePath
	}

	// 可以指定多个PRAGMA，例如：
	// path/to/some.db?_pragma=忙闲超时(5000)&_pragma=日志模式(WAL)
	// md5:407d435292848935
	if config.Extra != "" {
		var (
			options  string
			extraMap map[string]interface{}
		)
		if extraMap, err = gstr.X参数解析(config.Extra); err != nil {
			return nil, err
		}
		for k, v := range extraMap {
			if options != "" {
				options += "&"
			}
			options += fmt.Sprintf(`_pragma=%s(%s)`, k, gurl.X编码(gconv.String(v)))
		}
		if len(options) > 1 {
			source += "?" + options
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
