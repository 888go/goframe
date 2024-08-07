// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gmlock实现了一个基于内存的并发安全锁。 md5:46d1d349c220f670
package 内存锁类

var (
	// Default locker.
	locker = X创建()
)

// X写锁定 以写锁方式锁定 `key`。
// 如果有写锁或读锁正在锁定 `key`，
// 它将阻塞直到锁被释放。
// md5:8dcc0b1e059e3831
func X写锁定(名称 string) {
	locker.X写锁定(名称)
}

// X非阻塞写锁定 尝试对`key`加写锁，如果成功则返回true。如果`key`已有写锁或读锁，则返回false。
// md5:e4f172d7beca094d
func X非阻塞写锁定(名称 string) bool {
	return locker.X非阻塞写锁定(名称)
}

// X退出写锁定 解锁对`key`的写入锁。 md5:b54a3ae386cfa500
func X退出写锁定(名称 string) {
	locker.X退出写锁定(名称)
}

// X读锁定 使用读锁锁定`key`。如果`key`上有一个写锁，它将阻塞直到写锁被释放。
// md5:f45f660f368bbb78
func X读锁定(名称 string) {
	locker.X读锁定(名称)
}

// X非阻塞读锁定 尝试使用读取锁对 `key` 进行加锁。
// 如果成功，则返回true；如果 `key` 上存在写入锁，则返回false。
// md5:8733aa161c104b87
func X非阻塞读锁定(名称 string) bool {
	return locker.X非阻塞读锁定(名称)
}

// X退出读锁定 释放对 `key` 的读取锁。 md5:d4f823abaa858783
func X退出读锁定(名称 string) {
	locker.X退出读锁定(名称)
}

// X写锁定_函数 使用写入锁锁定`key`，并使用回调函数`f`。如果`key`已有写入或读取锁，它将阻塞直到锁被释放。
// 
// 在`f`执行后，它会释放锁。
// md5:3e35c1977b58dac3
func X写锁定_函数(名称 string, 回调函数 func()) {
	locker.X写锁定_函数(名称, 回调函数)
}

// X读锁定_函数 使用读取锁对 `key` 进行锁定，并调用回调函数 `f`。
// 如果 `key` 上有写入锁，它将阻塞直到锁被释放。
//
// 在 `f` 执行完毕后，它会释放锁。
// md5:e4a03ce9029d1911
func X读锁定_函数(名称 string, 回调函数 func()) {
	locker.X读锁定_函数(名称, 回调函数)
}

// X非阻塞写锁定_函数 使用写锁锁定`key`并执行回调函数`f`。
// 如果操作成功，返回true；如果`key`已被写锁或读锁占用，则返回false。
//
// 在回调函数`f`执行完毕后，它会释放锁。
// md5:a016db0c6b2bc67e
func X非阻塞写锁定_函数(名称 string, 回调函数 func()) bool {
	return locker.X非阻塞写锁定_函数(名称, 回调函数)
}

// X非阻塞读锁定_函数 使用读取锁尝试锁定`key`，并调用回调函数`f`。
// 如果成功，它返回true；如果`key`已有写入锁，则返回false。
//
// 在`f`执行完毕后释放锁。
// md5:527ef8bb470bd8fd
func X非阻塞读锁定_函数(名称 string, 回调函数 func()) bool {
	return locker.X非阻塞读锁定_函数(名称, 回调函数)
}

// X删除锁 删除具有给定`key`的互斥锁。 md5:a8a4db10705ecf7a
func X删除锁(名称 string) {
	locker.X删除锁(名称)
}
