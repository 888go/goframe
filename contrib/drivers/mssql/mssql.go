// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package mssql implements gdb.Driver, which supports operations for database MSSql.
//
// 1. It does not support Replace features.
// 2. It does not support LastInsertId.
package mssql

import (
	_ "github.com/denisenkom/go-mssqldb"

	"github.com/gogf/gf/v2/database/gdb"
)

// Driver is the driver for SQL server database.
type Driver struct {
	*gdb.Core
}

const (
	quoteChar = `"`
)

func init() {
	if err := gdb.Register(`mssql`, New()); err != nil {
		panic(err)
	}
}

// New create and returns a driver that implements gdb.Driver, which supports operations for Mssql.
// ff:
func New() gdb.Driver {
	return &Driver{}
}

// New creates and returns a database object for SQL server.
// It implements the interface of gdb.Driver for extra database driver installation.
// ff:
// d:
// core:
// node:
func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &Driver{
		Core: core,
	}, nil
}

// GetChars returns the security char for this type of database.
// ff:底层取数据库安全字符
// d:
// charLeft:左字符
// charRight:右字符
func (d *Driver) GetChars() (charLeft string, charRight string) {
	return quoteChar, quoteChar
}
