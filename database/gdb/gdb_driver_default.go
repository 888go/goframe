// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package db类

import (
	"database/sql"
)

// DriverDefault是MySQL数据库的默认驱动，它什么都不做。 md5:c3fc81ab467241cd
type DriverDefault struct {
	*Core
}

func init() {
	if err := X注册驱动("default", &DriverDefault{}); err != nil {
		panic(err)
	}
}

// New 创建并返回一个针对 MySQL 的数据库对象。它实现了 gdb.Driver 接口，以便于额外的数据库驱动程序安装。
// md5:e61df629828efeff
func (d *DriverDefault) New(core *Core, node *ConfigNode) (DB, error) {
	return &DriverDefault{
		Core: core,
	}, nil
}

// X底层Open 创建并返回一个底层的 sql.DB 对象，用于 MySQL。
// 注意，它将时间.Time 参数默认转换为本地时区。
// md5:341df118003c304e
func (d *DriverDefault) X底层Open(配置对象 *ConfigNode) (db *sql.DB, err error) {
	return
}

// X向主节点发送心跳 向主节点发送请求以检查身份验证或保持连接活动。 md5:47a7df55cbee8583
func (d *DriverDefault) X向主节点发送心跳() error {
	return nil
}

// X向从节点发送心跳 调用ping命令检查从节点的认证或者维持连接。 md5:62272b38d874eda6
func (d *DriverDefault) X向从节点发送心跳() error {
	return nil
}
