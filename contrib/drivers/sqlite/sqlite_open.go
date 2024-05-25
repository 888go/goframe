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

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// Open 创建并返回一个底层的 sql.DB 对象，用于 SQLite。
// https://github.com/glebarez/go-sqlite
// md5:f3c66f09b7236ffa
func (d *Driver) Open(config *gdb.ConfigNode) (db *sql.DB, err error) {
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
	// 它会搜索源文件以找到其绝对路径。. md5:c47257e12a6e34cf
	if absolutePath, _ := gfile.Search(source); absolutePath != "" {
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
