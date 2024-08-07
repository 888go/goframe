// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package json类

// X设置参数分隔符 设置用于层次数据访问的分隔符字符。 md5:99655266409bea6a
func (j *Json) X设置参数分隔符(分隔符 byte) {
	j.mu.Lock()
	j.c = 分隔符
	j.mu.Unlock()
}

// X设置分层冲突检查 用于启用或禁用层次数据访问时的暴力检查。 md5:b2fa0bb88e62957d
func (j *Json) X设置分层冲突检查(启用 bool) {
	j.mu.Lock()
	j.vc = 启用
	j.mu.Unlock()
}
