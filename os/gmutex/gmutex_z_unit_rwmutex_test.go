// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 互斥锁类_test

import (
	"context"
	"testing"
	"time"

	garray "github.com/888go/goframe/container/garray"
	glog "github.com/888go/goframe/os/glog"
	gmutex "github.com/888go/goframe/os/gmutex"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_RWMutex_RUnlock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		mu := gmutex.RWMutex{}
		mu.X读锁定_函数(func() {
			time.Sleep(200 * time.Millisecond)
		})
	})

	// RLock before Lock
	gtest.C(t, func(t *gtest.T) {
		mu := gmutex.RWMutex{}
		mu.RLock()
		go func() {
			mu.Lock()
			time.Sleep(300 * time.Millisecond)
			mu.Unlock()
		}()
		time.Sleep(100 * time.Millisecond)
		mu.RUnlock()
	})
}

func Test_RWMutex_IsLocked(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		mu := gmutex.RWMutex{}
		go func() {
			mu.X写锁定_函数(func() {
				time.Sleep(200 * time.Millisecond)
			})
		}()
		time.Sleep(100 * time.Millisecond)

		go func() {
			mu.X读锁定_函数(func() {
				time.Sleep(200 * time.Millisecond)
			})
		}()
	})
}

func Test_RWMutex_Unlock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		mu := gmutex.RWMutex{}
		array := garray.X创建(true)
		go func() {
			mu.X写锁定_函数(func() {
				array.Append别名(1)
				time.Sleep(300 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.X写锁定_函数(func() {
				array.Append别名(1)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.X写锁定_函数(func() {
				array.Append别名(1)
			})
		}()

		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(400 * time.Millisecond)
		t.Assert(array.X取长度(), 3)
	})
}

func Test_RWMutex_LockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		mu := gmutex.RWMutex{}
		array := garray.X创建(true)
		go func() {
			mu.X写锁定_函数(func() {
				array.Append别名(1)
				time.Sleep(300 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.X写锁定_函数(func() {
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

func Test_RWMutex_TryLockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		mu := gmutex.RWMutex{}
		array := garray.X创建(true)
		go func() {
			mu.X写锁定_函数(func() {
				array.Append别名(1)
				time.Sleep(300 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.X非阻塞写锁定_函数(func() {
				array.Append别名(1)
			})
		}()
		go func() {
			time.Sleep(400 * time.Millisecond)
			mu.X非阻塞写锁定_函数(func() {
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

func Test_RWMutex_RLockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		mu := gmutex.RWMutex{}
		array := garray.X创建(true)
		go func() {
			mu.X写锁定_函数(func() {
				array.Append别名(1)
				time.Sleep(300 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.X读锁定_函数(func() {
				array.Append别名(1)
				time.Sleep(100 * time.Millisecond)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})

	gtest.C(t, func(t *gtest.T) {
		mu := gmutex.RWMutex{}
		array := garray.X创建(true)
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.X读锁定_函数(func() {
				array.Append别名(1)
				time.Sleep(100 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.X读锁定_函数(func() {
				array.Append别名(1)
				time.Sleep(100 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.X读锁定_函数(func() {
				array.Append别名(1)
				time.Sleep(100 * time.Millisecond)
			})
		}()
		t.Assert(array.X取长度(), 0)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 3)
	})
}

func Test_RWMutex_TryRLockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			mu    = gmutex.RWMutex{}
			array = garray.X创建(true)
		)
		// First writing lock
		go func() {
			mu.X写锁定_函数(func() {
				array.Append别名(1)
				glog.X输出(context.TODO(), "lock1 done")
				time.Sleep(2000 * time.Millisecond)
			})
		}()
				// 这个goroutine永远无法获取锁。 md5:47331ca99d298935
		go func() {
			time.Sleep(1000 * time.Millisecond)
			mu.X非阻塞读锁定_函数(func() {
				array.Append别名(1)
			})
		}()
		for index := 0; index < 1000; index++ {
			go func() {
				time.Sleep(4000 * time.Millisecond)
				mu.X非阻塞读锁定_函数(func() {
					array.Append别名(1)
				})
			}()
		}
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(2000 * time.Millisecond)
		t.Assert(array.X取长度(), 1001)
	})
}
