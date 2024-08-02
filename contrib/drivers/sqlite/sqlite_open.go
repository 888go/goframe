// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

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

// Open creates and returns an underlying sql.DB object for sqlite.
// https://github.com/glebarez/go-sqlite
func (d *Driver) Open(config *gdb.ConfigNode) (db *sql.DB, err error) {
	var (
		source               string
		underlyingDriverName = "sqlite"
	)
	if config.Link != "" {
		// ============================================================================
		// Deprecated from v2.2.0.
		// ============================================================================
		source = config.Link
	} else {
		source = config.Name
	}
	// It searches the source file to locate its absolute path..
	if absolutePath, _ := gfile.Search(source); absolutePath != "" {
		source = absolutePath
	}

	// Multiple PRAGMAs can be specified, e.g.:
	// path/to/some.db?_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)
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
