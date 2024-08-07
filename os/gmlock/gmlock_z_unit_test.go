// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 内存锁类_test

import (
	"sync"
	"testing"
	"time"

	garray "github.com/888go/goframe/container/garray"
	gmlock "github.com/888go/goframe/os/gmlock"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_Locker_Lock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		key := "testLock"
		array := garray.X创建(true)
		go func() {
			gmlock.X写锁定(key)
			array.Append别名(1)
			time.Sleep(300 * time.Millisecond)
			gmlock.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.X写锁定(key)
			array.Append别名(1)
			gmlock.X退出写锁定(key)
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
		gmlock.X删除锁(key)
	})

	gtest.C(t, func(t *gtest.T) {
		key := "testLock"
		array := garray.X创建(true)
		lock := gmlock.X创建()
		go func() {
			lock.X写锁定(key)
			array.Append别名(1)
			time.Sleep(300 * time.Millisecond)
			lock.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			lock.X写锁定(key)
			array.Append别名(1)
			lock.X退出写锁定(key)
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
		lock.X移除所有锁()
	})

}

func Test_Locker_TryLock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		key := "testTryLock"
		array := garray.X创建(true)
		go func() {
			gmlock.X写锁定(key)
			array.Append别名(1)
			time.Sleep(300 * time.Millisecond)
			gmlock.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(150 * time.Millisecond)
			if gmlock.X非阻塞写锁定(key) {
				array.Append别名(1)
				gmlock.X退出写锁定(key)
			}
		}()
		go func() {
			time.Sleep(400 * time.Millisecond)
			if gmlock.X非阻塞写锁定(key) {
				array.Append别名(1)
				gmlock.X退出写锁定(key)
			}
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})

}

func Test_Locker_LockFunc(t *testing.T) {
	//no expire
	gtest.C(t, func(t *gtest.T) {
		key := "testLockFunc"
		array := garray.X创建(true)
		go func() {
			gmlock.X写锁定_函数(key, func() {
				array.Append别名(1)
				time.Sleep(300 * time.Millisecond)
			}) //
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.X写锁定_函数(key, func() {
				array.Append别名(1)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1) //
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}

func Test_Locker_TryLockFunc(t *testing.T) {
	//no expire
	gtest.C(t, func(t *gtest.T) {
		key := "testTryLockFunc"
		array := garray.X创建(true)
		go func() {
			gmlock.X非阻塞写锁定_函数(key, func() {
				array.Append别名(1)
				time.Sleep(200 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.X非阻塞写锁定_函数(key, func() {
				array.Append别名(1)
			})
		}()
		go func() {
			time.Sleep(300 * time.Millisecond)
			gmlock.X非阻塞写锁定_函数(key, func() {
				array.Append别名(1)
			})
		}()
		time.Sleep(150 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(400 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}

func Test_Multiple_Goroutine(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ch := make(chan struct{})
		num := 1000
		wait := sync.WaitGroup{}
		wait.Add(num)
		for i := 0; i < num; i++ {
			go func() {
				defer wait.Done()
				<-ch
				gmlock.X写锁定("test")
				defer gmlock.X退出写锁定("test")
				time.Sleep(time.Millisecond)
			}()
		}
		close(ch)
		wait.Wait()
	})

	gtest.C(t, func(t *gtest.T) {
		ch := make(chan struct{})
		num := 100
		wait := sync.WaitGroup{}
		wait.Add(num * 2)
		for i := 0; i < num; i++ {
			go func() {
				defer wait.Done()
				<-ch
				gmlock.X写锁定("test")
				defer gmlock.X退出写锁定("test")
				time.Sleep(time.Millisecond)
			}()
		}
		for i := 0; i < num; i++ {
			go func() {
				defer wait.Done()
				<-ch
				gmlock.X读锁定("test")
				defer gmlock.X退出读锁定("test")
				time.Sleep(time.Millisecond)
			}()
		}
		close(ch)
		wait.Wait()
	})
}

func Test_Locker_RLock(t *testing.T) {
	// RLock before Lock
	gtest.C(t, func(t *gtest.T) {
		key := "testRLockBeforeLock"
		array := garray.X创建(true)
		go func() {
			gmlock.X读锁定(key)
			array.Append别名(1)
			time.Sleep(200 * time.Millisecond)
			gmlock.X退出读锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.X写锁定(key)
			array.Append别名(1)
			gmlock.X退出写锁定(key)
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})

	// Lock before RLock
	gtest.C(t, func(t *gtest.T) {
		key := "testLockBeforeRLock"
		array := garray.X创建(true)
		go func() {
			gmlock.X写锁定(key)
			array.Append别名(1)
			time.Sleep(200 * time.Millisecond)
			gmlock.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.X读锁定(key)
			array.Append别名(1)
			gmlock.X退出读锁定(key)
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})

	// Lock before RLocks
	gtest.C(t, func(t *gtest.T) {
		key := "testLockBeforeRLocks"
		array := garray.X创建(true)
		go func() {
			gmlock.X写锁定(key)
			array.Append别名(1)
			time.Sleep(300 * time.Millisecond)
			gmlock.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.X读锁定(key)
			array.Append别名(1)
			time.Sleep(200 * time.Millisecond)
			gmlock.X退出读锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.X读锁定(key)
			array.Append别名(1)
			time.Sleep(200 * time.Millisecond)
			gmlock.X退出读锁定(key)
		}()
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 3)
	})
}

func Test_Locker_TryRLock(t *testing.T) {
	// Lock before TryRLock
	gtest.C(t, func(t *gtest.T) {
		key := "testLockBeforeTryRLock"
		array := garray.X创建(true)
		go func() {
			gmlock.X写锁定(key)
			array.Append别名(1)
			time.Sleep(200 * time.Millisecond)
			gmlock.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			if gmlock.X非阻塞读锁定(key) {
				array.Append别名(1)
				gmlock.X退出读锁定(key)
			}
		}()
		time.Sleep(150 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})

	// Lock before TryRLocks
	gtest.C(t, func(t *gtest.T) {
		key := "testLockBeforeTryRLocks"
		array := garray.X创建(true)
		go func() {
			gmlock.X写锁定(key)
			array.Append别名(1)
			time.Sleep(200 * time.Millisecond)
			gmlock.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			if gmlock.X非阻塞读锁定(key) {
				array.Append别名(1)
				gmlock.X退出读锁定(key)
			}
		}()
		go func() {
			time.Sleep(300 * time.Millisecond)
			if gmlock.X非阻塞读锁定(key) {
				array.Append别名(1)
				gmlock.X退出读锁定(key)
			}
		}()
		time.Sleep(150 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}

func Test_Locker_RLockFunc(t *testing.T) {
	// RLockFunc before Lock
	gtest.C(t, func(t *gtest.T) {
		key := "testRLockFuncBeforeLock"
		array := garray.X创建(true)
		go func() {
			gmlock.X读锁定_函数(key, func() {
				array.Append别名(1)
				time.Sleep(200 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.X写锁定(key)
			array.Append别名(1)
			gmlock.X退出写锁定(key)
		}()
		time.Sleep(150 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})

	// Lock before RLockFunc
	gtest.C(t, func(t *gtest.T) {
		key := "testLockBeforeRLockFunc"
		array := garray.X创建(true)
		go func() {
			gmlock.X写锁定(key)
			array.Append别名(1)
			time.Sleep(200 * time.Millisecond)
			gmlock.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.X读锁定_函数(key, func() {
				array.Append别名(1)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})

	// Lock before RLockFuncs
	gtest.C(t, func(t *gtest.T) {
		key := "testLockBeforeRLockFuncs"
		array := garray.X创建(true)
		go func() {
			gmlock.X写锁定(key)
			array.Append别名(1)
			time.Sleep(200 * time.Millisecond)
			gmlock.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.X读锁定_函数(key, func() {
				array.Append别名(1)
				time.Sleep(200 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.X读锁定_函数(key, func() {
				array.Append别名(1)
				time.Sleep(200 * time.Millisecond)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 3)
	})
}

func Test_Locker_TryRLockFunc(t *testing.T) {
		// 在尝试读取锁函数之前加锁. md5:4b4b42302a1b68f7
	gtest.C(t, func(t *gtest.T) {
		key := "testLockBeforeTryRLockFunc"
		array := garray.X创建(true)
		go func() {
			gmlock.X写锁定(key)
			array.Append别名(1)
			time.Sleep(200 * time.Millisecond)
			gmlock.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.X非阻塞读锁定_函数(key, func() {
				array.Append别名(1)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})

			// 在尝试读取锁函数之前加锁. md5:4b4b42302a1b68f7
	gtest.C(t, func(t *gtest.T) {
		key := "testLockBeforeTryRLockFuncs"
		array := garray.X创建(true)
		go func() {
			gmlock.X写锁定(key)
			array.Append别名(1)
			time.Sleep(200 * time.Millisecond)
			gmlock.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			gmlock.X非阻塞读锁定_函数(key, func() {
				array.Append别名(1)
			})
		}()
		go func() {
			time.Sleep(300 * time.Millisecond)
			gmlock.X非阻塞读锁定_函数(key, func() {
				array.Append别名(1)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}
