// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb

import (
	"database/sql"
)

// DriverDefault 是 mysql 数据库的默认驱动，它实际上什么都不做。
type DriverDefault struct {
	*Core
}

func init() {
	if err := Register("default", &DriverDefault{}); err != nil {
		panic(err)
	}
}

// New 创建并返回一个用于 mysql 的数据库对象。
// 它实现了 gdb.Driver 接口，以便进行额外的数据库驱动安装。
func (d *DriverDefault) New(core *Core, node *ConfigNode) (DB, error) {
	return &DriverDefault{
		Core: core,
	}, nil
}

// Open 创建并返回一个用于 mysql 的底层 sql.DB 对象。
// 注意，它默认会将 time.Time 类型参数转换为本地时区。
func (d *DriverDefault) Open(config *ConfigNode) (db *sql.DB, err error) {
	return
}

// PingMaster 用于向主节点发送心跳以检查身份验证或保持连接存活。
func (d *DriverDefault) PingMaster() error {
	return nil
}

// PingSlave 向从节点发送ping请求，用于检查身份验证或保持连接活跃。
func (d *DriverDefault) PingSlave() error {
	return nil
}
