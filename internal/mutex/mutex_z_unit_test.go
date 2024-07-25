// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package mutex_test

import (
	"testing"
	"time"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/internal/mutex"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestMutexIsSafe(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		lock := mutex.New()
		t.Assert(lock.IsSafe(), false)

		lock = mutex.New(false)
		t.Assert(lock.IsSafe(), false)

		lock = mutex.New(false, false)
		t.Assert(lock.IsSafe(), false)

		lock = mutex.New(true, false)
		t.Assert(lock.IsSafe(), true)

		lock = mutex.New(true, true)
		t.Assert(lock.IsSafe(), true)

		lock = mutex.New(true)
		t.Assert(lock.IsSafe(), true)
	})
}

func TestSafeMutex(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		safeLock := mutex.New(true)
		array := garray.New(true)

		go func() {
			safeLock.Lock()
			array.Append(1)
			time.Sleep(1000 * time.Millisecond)
			array.Append(1)
			safeLock.Unlock()
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			safeLock.Lock()
			array.Append(1)
			time.Sleep(2000 * time.Millisecond)
			array.Append(1)
			safeLock.Unlock()
		}()
		time.Sleep(500 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(800 * time.Millisecond)
		t.Assert(array.Len(), 3)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 3)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 4)
	})
}

func TestUnsafeMutex(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			unsafeLock = mutex.New()
			array      = garray.New(true)
		)

		go func() {
			unsafeLock.Lock()
			array.Append(1)
			time.Sleep(1000 * time.Millisecond)
			array.Append(1)
			unsafeLock.Unlock()
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			unsafeLock.Lock()
			array.Append(1)
			time.Sleep(2000 * time.Millisecond)
			array.Append(1)
			unsafeLock.Unlock()
		}()
		time.Sleep(500 * time.Millisecond)
		t.Assert(array.Len(), 2)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 3)
		time.Sleep(500 * time.Millisecond)
		t.Assert(array.Len(), 3)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 4)
	})
}
