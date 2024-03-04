// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gmlock 实现了一个基于内存的并发安全锁。
package gmlock

var (
	// Default locker.
	locker = New()
)

// Lock 对`key`使用写入锁进行加锁。
// 如果有对`key`的写入/读取锁存在，
// 它将阻塞直到锁被释放。
func Lock(key string) {
	locker.Lock(key)
}

// TryLock 尝试对`key`进行写锁锁定，
// 如果成功则返回 true，如果`key`存在写锁或读锁，则返回 false。
func TryLock(key string) bool {
	return locker.TryLock(key)
}

// Unlock 解锁指定`key`的写入锁。
func Unlock(key string) {
	locker.Unlock(key)
}

// RLock 对“key”加读锁。
// 如果“key”上存在写锁，
// 它将阻塞直到写锁被释放。
func RLock(key string) {
	locker.RLock(key)
}

// TryRLock 尝试对`key`使用读锁进行加锁。
// 如果加锁成功，返回true；如果`key`上存在写锁，则返回false。
func TryRLock(key string) bool {
	return locker.TryRLock(key)
}

// RUnlock 解除对`key`的读取锁。
func RUnlock(key string) {
	locker.RUnlock(key)
}

// LockFunc 对`key`使用写锁并调用回调函数`f`进行锁定。
// 如果有对`key`的写锁/读锁存在，
// 它将阻塞直到锁被释放。
//
// 在`f`执行完毕后，它会自动释放锁。
func LockFunc(key string, f func()) {
	locker.LockFunc(key, f)
}

// RLockFunc 对 `key` 加上读锁并执行回调函数 `f`。
// 如果存在对 `key` 的写锁，则会阻塞直到该锁被释放。
//
// 在 `f` 执行完毕后，它将自动释放锁。
func RLockFunc(key string, f func()) {
	locker.RLockFunc(key, f)
}

// TryLockFunc 尝试对`key`使用写锁并执行回调函数`f`。
// 如果成功获取锁，则返回 true，否则如果`key`已有写锁或读锁存在，则返回 false。
//
// 在`f`执行完毕后会自动释放锁。
func TryLockFunc(key string, f func()) bool {
	return locker.TryLockFunc(key, f)
}

// TryRLockFunc 尝试对`key`加读锁并执行回调函数`f`。
// 如果成功获取读锁则返回true，否则如果`key`已有写锁存在，则返回false。
//
// 在`f`执行完毕后会自动释放锁。
func TryRLockFunc(key string, f func()) bool {
	return locker.TryRLockFunc(key, f)
}

// Remove 通过给定的`key`移除互斥锁。
func Remove(key string) {
	locker.Remove(key)
}
