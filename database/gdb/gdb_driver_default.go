// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gdb

import (
	"database/sql"
)

// DriverDefault是MySQL数据库的默认驱动，它什么都不做。 md5:c3fc81ab467241cd
type DriverDefault struct {
	*Core
}

func init() {
	if err := Register("default", &DriverDefault{}); err != nil {
		panic(err)
	}
}

// New 创建并返回一个针对 MySQL 的数据库对象。它实现了 gdb.Driver 接口，以便于额外的数据库驱动程序安装。
// md5:e61df629828efeff
// ff:
// d:
// core:
// node:
// DB:
func (d *DriverDefault) New(core *Core, node *ConfigNode) (DB, error) {
	return &DriverDefault{
		Core: core,
	}, nil
}

// Open 创建并返回一个底层的 sql.DB 对象，用于 MySQL。
// 注意，它将时间.Time 参数默认转换为本地时区。
// md5:341df118003c304e
// ff:底层Open
// d:
// config:配置对象
// db:
// err:
func (d *DriverDefault) Open(config *ConfigNode) (db *sql.DB, err error) {
	return
}

// PingMaster 向主节点发送请求以检查身份验证或保持连接活动。 md5:47a7df55cbee8583
// ff:向主节点发送心跳
// d:
func (d *DriverDefault) PingMaster() error {
	return nil
}

// PingSlave 调用ping命令检查从节点的认证或者维持连接。 md5:62272b38d874eda6
// ff:向从节点发送心跳
// d:
func (d *DriverDefault) PingSlave() error {
	return nil
}
