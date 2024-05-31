// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gdb

import (
	"database/sql"
)

// DriverDefault is the default driver for mysql database, which does nothing.
type DriverDefault struct {
	*Core
}

func init() {
	if err := Register("default", &DriverDefault{}); err != nil {
		panic(err)
	}
}

// New creates and returns a database object for mysql.
// It implements the interface of gdb.Driver for extra database driver installation.

// ff:
// DB:
// node:
// core:
func (d *DriverDefault) New(core *Core, node *ConfigNode) (DB, error) {
	return &DriverDefault{
		Core: core,
	}, nil
}

// Open creates and returns an underlying sql.DB object for mysql.
// Note that it converts time.Time argument to local timezone in default.

// ff:底层Open
// err:
// db:
// config:配置对象
func (d *DriverDefault) Open(config *ConfigNode) (db *sql.DB, err error) {
	return
}

// PingMaster pings the master node to check authentication or keeps the connection alive.

// ff:向主节点发送心跳
func (d *DriverDefault) PingMaster() error {
	return nil
}

// PingSlave pings the slave node to check authentication or keeps the connection alive.

// ff:向从节点发送心跳
func (d *DriverDefault) PingSlave() error {
	return nil
}
