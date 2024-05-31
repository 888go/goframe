// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gmlock implements a concurrent-safe memory-based locker.
package gmlock//bm:内存锁类

var (
	// Default locker.
	locker = New()
)

// Lock locks the `key` with writing lock.
// If there's a write/reading lock the `key`,
// it will blocks until the lock is released.

// ff:写锁定
// key:名称
func Lock(key string) {
	locker.Lock(key)
}

// TryLock tries locking the `key` with writing lock,
// it returns true if success, or if there's a write/reading lock the `key`,
// it returns false.

// ff:非阻塞写锁定
// key:名称
func TryLock(key string) bool {
	return locker.TryLock(key)
}

// Unlock unlocks the writing lock of the `key`.

// ff:退出写锁定
// key:名称
func Unlock(key string) {
	locker.Unlock(key)
}

// RLock locks the `key` with reading lock.
// If there's a writing lock on `key`,
// it will blocks until the writing lock is released.

// ff:读锁定
// key:名称
func RLock(key string) {
	locker.RLock(key)
}

// TryRLock tries locking the `key` with reading lock.
// It returns true if success, or if there's a writing lock on `key`, it returns false.

// ff:非阻塞读锁定
// key:名称
func TryRLock(key string) bool {
	return locker.TryRLock(key)
}

// RUnlock unlocks the reading lock of the `key`.

// ff:退出读锁定
// key:名称
func RUnlock(key string) {
	locker.RUnlock(key)
}

// LockFunc locks the `key` with writing lock and callback function `f`.
// If there's a write/reading lock the `key`,
// it will blocks until the lock is released.
//
// It releases the lock after `f` is executed.

// ff:写锁定_函数
// f:回调函数
// key:名称
func LockFunc(key string, f func()) {
	locker.LockFunc(key, f)
}

// RLockFunc locks the `key` with reading lock and callback function `f`.
// If there's a writing lock the `key`,
// it will blocks until the lock is released.
//
// It releases the lock after `f` is executed.

// ff:读锁定_函数
// f:回调函数
// key:名称
func RLockFunc(key string, f func()) {
	locker.RLockFunc(key, f)
}

// TryLockFunc locks the `key` with writing lock and callback function `f`.
// It returns true if success, or else if there's a write/reading lock the `key`, it return false.
//
// It releases the lock after `f` is executed.

// ff:非阻塞写锁定_函数
// f:回调函数
// key:名称
func TryLockFunc(key string, f func()) bool {
	return locker.TryLockFunc(key, f)
}

// TryRLockFunc locks the `key` with reading lock and callback function `f`.
// It returns true if success, or else if there's a writing lock the `key`, it returns false.
//
// It releases the lock after `f` is executed.

// ff:非阻塞读锁定_函数
// f:回调函数
// key:名称
func TryRLockFunc(key string, f func()) bool {
	return locker.TryRLockFunc(key, f)
}

// Remove removes mutex with given `key`.

// ff:删除锁
// key:名称
func Remove(key string) {
	locker.Remove(key)
}
