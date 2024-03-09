// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 互斥锁类_test

import (
	"testing"
	"time"
	
	"github.com/gogf/gf/v2/container/garray"
	"github.com/888go/goframe/gmutex"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Mutex_Unlock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		mu := 互斥锁类.X互斥锁{}
		array := garray.New(true)
		go func() {
			mu.X锁定_函数(func() {
				array.Append(1)
				time.Sleep(300 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.X锁定_函数(func() {
				array.Append(1)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.X锁定_函数(func() {
				array.Append(1)
			})
		}()

		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(400 * time.Millisecond)
		t.Assert(array.Len(), 3)
	})
}

func Test_Mutex_LockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		mu := 互斥锁类.X互斥锁{}
		array := garray.New(true)
		go func() {
			mu.X锁定_函数(func() {
				array.Append(1)
				time.Sleep(300 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.X锁定_函数(func() {
				array.Append(1)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func Test_Mutex_TryLockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		mu := 互斥锁类.X互斥锁{}
		array := garray.New(true)
		go func() {
			mu.X锁定_函数(func() {
				array.Append(1)
				time.Sleep(300 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.X非阻塞锁定_函数(func() {
				array.Append(1)
			})
		}()
		go func() {
			time.Sleep(400 * time.Millisecond)
			mu.X非阻塞锁定_函数(func() {
				array.Append(1)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}
