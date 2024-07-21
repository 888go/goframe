// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包 gmutex 继承并扩展了 sync.Mutex 和 sync.RWMutex，提供了更多的功能。
// 
// 注意，从 GoFrame 版本 v2.5.2 开始，它采用了标准库同步包 sync 的 mutex。
// md5:63811cedd95f3f75
package gmutex//bm:互斥锁类

// New creates and returns a new mutex.
// ff:创建
func New() *RWMutex {
	return &RWMutex{}
}
