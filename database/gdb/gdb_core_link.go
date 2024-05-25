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

// dbLink用于实现DB接口的Link。. md5:70a34111ff15fefa
type dbLink struct {
	*sql.DB         // Underlying DB object.
	isOnMaster bool // isOnMaster 标记当前链接是否在主节点上运行。. md5:451060dcca211e1a
}

// txLink 用于为 TX 实现接口 Link。. md5:1c367f806a5592fd
type txLink struct {
	*sql.Tx
}

// IsTransaction 返回当前Link是否为交易。. md5:83b57bf1eb9cb599
func (l *dbLink) IsTransaction() bool {
	return false
}

// IsOnMaster 检查并返回当前链接是否在主节点上运行。. md5:ec0eef810e78c3f2
func (l *dbLink) IsOnMaster() bool {
	return l.isOnMaster
}

// IsTransaction 返回当前Link是否为交易。. md5:83b57bf1eb9cb599
func (l *txLink) IsTransaction() bool {
	return true
}

// IsOnMaster 检查并返回当前链接是否在主节点上运行。. md5:ec0eef810e78c3f2
// Note that, transaction operation is always operated on master node.
func (l *txLink) IsOnMaster() bool {
	return true
}
