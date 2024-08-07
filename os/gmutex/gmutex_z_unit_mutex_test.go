// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 互斥锁类_test

import (
	"testing"
	"time"

	garray "github.com/888go/goframe/container/garray"
	gmutex "github.com/888go/goframe/os/gmutex"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_Mutex_Unlock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		mu := gmutex.Mutex{}
		array := garray.X创建(true)
		go func() {
			mu.X锁定_函数(func() {
				array.Append别名(1)
				time.Sleep(300 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.X锁定_函数(func() {
				array.Append别名(1)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.X锁定_函数(func() {
				array.Append别名(1)
			})
		}()

		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(400 * time.Millisecond)
		t.Assert(array.X取长度(), 3)
	})
}

func Test_Mutex_LockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		mu := gmutex.Mutex{}
		array := garray.X创建(true)
		go func() {
			mu.X锁定_函数(func() {
				array.Append别名(1)
				time.Sleep(300 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.X锁定_函数(func() {
				array.Append别名(1)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}

func Test_Mutex_TryLockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		mu := gmutex.Mutex{}
		array := garray.X创建(true)
		go func() {
			mu.X锁定_函数(func() {
				array.Append别名(1)
				time.Sleep(300 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.X非阻塞锁定_函数(func() {
				array.Append别名(1)
			})
		}()
		go func() {
			time.Sleep(400 * time.Millisecond)
			mu.X非阻塞锁定_函数(func() {
				array.Append别名(1)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}
