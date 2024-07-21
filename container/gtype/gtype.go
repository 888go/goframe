// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gtype提供高性能和线程安全的基础变量类型。 md5:94b883ebf0b43fd8
package gtype

// New 是 NewAny 的别名。
// 请参阅 NewAny，NewInterface。
// md5:a0f9c97b9c253975
func New(value ...interface{}) *Any {
	return NewAny(value...)
}
