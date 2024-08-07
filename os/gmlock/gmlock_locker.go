// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 内存锁类

import (
	"sync"

	gmap "github.com/888go/goframe/container/gmap"
)

// Locker 是一个基于内存的锁。
// 注意，Lockers 中的 mutex 没有缓存过期机制。
// 当你不再需要使用某个锁时，需要手动删除它。
// md5:330f85347bba3cc8
type Locker struct {
	m *gmap.StrAnyMap
}

// X创建 创建并返回一个新的内存锁。
// 这个内存锁能够使用动态字符串键进行锁定和解锁。
// md5:37aaba9921b3e711
func X创建() *Locker {
	return &Locker{
		m: gmap.X创建StrAny(true),
	}
}

// X写锁定 以写锁方式锁定`key`。
// 如果`key`已有写锁或读锁，它会阻塞直到锁被释放。
// md5:7b2d56ac41ec0a40
func (l *Locker) X写锁定(名称 string) {
	l.getOrNewMutex(名称).Lock()
}

// X非阻塞写锁定尝试使用写入锁锁定`key`，如果成功返回true，如果`key`已经有写入或读取锁则返回false。
// md5:1e86a7888ed1621a
func (l *Locker) X非阻塞写锁定(名称 string) bool {
	return l.getOrNewMutex(名称).TryLock()
}

// X退出写锁定 解锁对`key`的写入锁。 md5:b54a3ae386cfa500
func (l *Locker) X退出写锁定(名称 string) {
	if v := l.m.X取值(名称); v != nil {
		v.(*sync.RWMutex).Unlock()
	}
}

// X读锁定 使用读锁锁定`key`。如果`key`上有一个写锁，它将阻塞直到写锁被释放。
// md5:f45f660f368bbb78
func (l *Locker) X读锁定(名称 string) {
	l.getOrNewMutex(名称).RLock()
}

// X非阻塞读锁定 尝试使用读取锁对 `key` 进行加锁。
// 如果成功，则返回true；如果 `key` 上存在写入锁，则返回false。
// md5:8733aa161c104b87
func (l *Locker) X非阻塞读锁定(名称 string) bool {
	return l.getOrNewMutex(名称).TryRLock()
}

// X退出读锁定 释放对 `key` 的读取锁。 md5:d4f823abaa858783
func (l *Locker) X退出读锁定(名称 string) {
	if v := l.m.X取值(名称); v != nil {
		v.(*sync.RWMutex).RUnlock()
	}
}

// X写锁定_函数 使用写入锁锁定`key`，并使用回调函数`f`。
// 如果`key`已有写入或读取锁，它将阻塞直到锁被释放。
//
// 在执行完`f`后，它会释放锁。
// md5:fc66c542fa813208
func (l *Locker) X写锁定_函数(名称 string, 回调函数 func()) {
	l.X写锁定(名称)
	defer l.X退出写锁定(名称)
	回调函数()
}

// X读锁定_函数 使用读取锁对`key`进行加锁，并执行回调函数`f`。
// 如果`key`已被写入锁锁定，
// 则会阻塞直到该锁被释放。
//
// 在`f`执行完毕后，它将释放锁。
// md5:3f30fb5d911cd5e7
func (l *Locker) X读锁定_函数(名称 string, 回调函数 func()) {
	l.X读锁定(名称)
	defer l.X退出读锁定(名称)
	回调函数()
}

// X非阻塞写锁定_函数 使用写锁锁定`key`并执行回调函数`f`。
// 如果操作成功，返回true；如果`key`已被写锁或读锁占用，则返回false。
//
// 在回调函数`f`执行完毕后，它会释放锁。
// md5:a016db0c6b2bc67e
func (l *Locker) X非阻塞写锁定_函数(名称 string, 回调函数 func()) bool {
	if l.X非阻塞写锁定(名称) {
		defer l.X退出写锁定(名称)
		回调函数()
		return true
	}
	return false
}

// X非阻塞读锁定_函数 使用读取锁尝试锁定`key`，并调用回调函数`f`。
// 如果成功，它返回true；如果`key`已有写入锁，则返回false。
//
// 在`f`执行完毕后释放锁。
// md5:527ef8bb470bd8fd
func (l *Locker) X非阻塞读锁定_函数(名称 string, 回调函数 func()) bool {
	if l.X非阻塞读锁定(名称) {
		defer l.X退出读锁定(名称)
		回调函数()
		return true
	}
	return false
}

// X删除锁 根据给定的 `key` 从 locker 中移除互斥锁。 md5:e557087320b6a672
func (l *Locker) X删除锁(名称 string) {
	l.m.X删除(名称)
}

// X移除所有锁 从 locker 中移除所有互斥锁。 md5:6e7b8ead4ad69f9d
func (l *Locker) X移除所有锁() {
	l.m.X清空()
}

// getOrNewMutex 如果给定的`key`存在，则返回对应的互斥锁，否则创建一个新的并返回。
// md5:08c35eff58386554
func (l *Locker) getOrNewMutex(key string) *sync.RWMutex {
	return l.m.X取值或设置值_函数带锁(key, func() interface{} {
		return &sync.RWMutex{}
	}).(*sync.RWMutex)
}
