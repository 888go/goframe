// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb

import (
	"database/sql"
)

// dbLink 用于实现接口 Link 对于数据库的操作。
type dbLink struct {
	*sql.DB         // 基础数据库对象。
	isOnMaster bool // isOnMaster 标记当前链接是否在主节点上操作。
}

// txLink 用于实现接口 Link 对于 TX 的需求。
type txLink struct {
	*sql.Tx
}

// IsTransaction 返回当前 Link 是否为一个事务。
func (l *dbLink) IsTransaction() bool {
	return false
}

// IsOnMaster 检查并返回当前链接是否在主节点上操作。
func (l *dbLink) IsOnMaster() bool {
	return l.isOnMaster
}

// IsTransaction 返回当前 Link 是否为一个事务。
func (l *txLink) IsTransaction() bool {
	return true
}

// IsOnMaster 检查并返回当前链接是否在主节点上操作。
// Note that, transaction operation is always operated on master node.
func (l *txLink) IsOnMaster() bool {
	return true
}
