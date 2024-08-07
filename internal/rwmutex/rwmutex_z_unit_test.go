// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package rwmutex_test

import (
	"testing"
	"time"

	garray "github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/internal/rwmutex"
	gtest "github.com/888go/goframe/test/gtest"
)

func TestRWMutexIsSafe(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		lock := rwmutex.New()
		t.Assert(lock.IsSafe(), false)

		lock = rwmutex.New(false)
		t.Assert(lock.IsSafe(), false)

		lock = rwmutex.New(false, false)
		t.Assert(lock.IsSafe(), false)

		lock = rwmutex.New(true, false)
		t.Assert(lock.IsSafe(), true)

		lock = rwmutex.New(true, true)
		t.Assert(lock.IsSafe(), true)

		lock = rwmutex.New(true)
		t.Assert(lock.IsSafe(), true)
	})
}

func TestSafeRWMutex(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			localSafeLock = rwmutex.New(true)
			array         = garray.X创建(true)
		)

		go func() {
			localSafeLock.Lock()
			array.Append别名(1)
			time.Sleep(1000 * time.Millisecond)
			array.Append别名(1)
			localSafeLock.Unlock()
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			localSafeLock.Lock()
			array.Append别名(1)
			time.Sleep(2000 * time.Millisecond)
			array.Append别名(1)
			localSafeLock.Unlock()
		}()
		time.Sleep(500 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(800 * time.Millisecond)
		t.Assert(array.X取长度(), 3)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 3)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 4)
	})
}

func TestSafeReaderRWMutex(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			localSafeLock = rwmutex.New(true)
			array         = garray.X创建(true)
		)
		go func() {
			localSafeLock.RLock()
			array.Append别名(1)
			time.Sleep(1000 * time.Millisecond)
			array.Append别名(1)
			localSafeLock.RUnlock()
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			localSafeLock.RLock()
			array.Append别名(1)
			time.Sleep(2000 * time.Millisecond)
			array.Append别名(1)
			time.Sleep(1000 * time.Millisecond)
			array.Append别名(1)
			localSafeLock.RUnlock()
		}()
		go func() {
			time.Sleep(500 * time.Millisecond)
			localSafeLock.Lock()
			array.Append别名(1)
			localSafeLock.Unlock()
		}()
		time.Sleep(500 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 3)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 4)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 6)
	})
}

func TestUnsafeRWMutex(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			localUnsafeLock = rwmutex.New()
			array           = garray.X创建(true)
		)
		go func() {
			localUnsafeLock.Lock()
			array.Append别名(1)
			time.Sleep(2000 * time.Millisecond)
			array.Append别名(1)
			localUnsafeLock.Unlock()
		}()
		go func() {
			time.Sleep(500 * time.Millisecond)
			localUnsafeLock.Lock()
			array.Append别名(1)
			time.Sleep(500 * time.Millisecond)
			array.Append别名(1)
			localUnsafeLock.Unlock()
		}()
		time.Sleep(800 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
		time.Sleep(800 * time.Millisecond)
		t.Assert(array.X取长度(), 3)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 3)
		time.Sleep(500 * time.Millisecond)
		t.Assert(array.X取长度(), 4)
	})
}
