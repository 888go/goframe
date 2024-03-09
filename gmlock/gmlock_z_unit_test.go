// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 内存锁类_test

import (
	"sync"
	"testing"
	"time"
	
	"github.com/gogf/gf/v2/container/garray"
	"github.com/888go/goframe/gmlock"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Locker_Lock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		key := "testLock"
		array := garray.New(true)
		go func() {
			内存锁类.X写锁定(key)
			array.Append(1)
			time.Sleep(300 * time.Millisecond)
			内存锁类.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X写锁定(key)
			array.Append(1)
			内存锁类.X退出写锁定(key)
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 2)
		内存锁类.X删除锁(key)
	})

	gtest.C(t, func(t *gtest.T) {
		key := "testLock"
		array := garray.New(true)
		lock := 内存锁类.X创建()
		go func() {
			lock.X写锁定(key)
			array.Append(1)
			time.Sleep(300 * time.Millisecond)
			lock.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			lock.X写锁定(key)
			array.Append(1)
			lock.X退出写锁定(key)
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 2)
		lock.X移除所有锁()
	})

}

func Test_Locker_TryLock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		key := "testTryLock"
		array := garray.New(true)
		go func() {
			内存锁类.X写锁定(key)
			array.Append(1)
			time.Sleep(300 * time.Millisecond)
			内存锁类.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(150 * time.Millisecond)
			if 内存锁类.X非阻塞写锁定(key) {
				array.Append(1)
				内存锁类.X退出写锁定(key)
			}
		}()
		go func() {
			time.Sleep(400 * time.Millisecond)
			if 内存锁类.X非阻塞写锁定(key) {
				array.Append(1)
				内存锁类.X退出写锁定(key)
			}
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})

}

func Test_Locker_LockFunc(t *testing.T) {
	//no expire
	gtest.C(t, func(t *gtest.T) {
		key := "testLockFunc"
		array := garray.New(true)
		go func() {
			内存锁类.X写锁定_函数(key, func() {
				array.Append(1)
				time.Sleep(300 * time.Millisecond)
			}) //
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X写锁定_函数(key, func() {
				array.Append(1)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1) //
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func Test_Locker_TryLockFunc(t *testing.T) {
	//no expire
	gtest.C(t, func(t *gtest.T) {
		key := "testTryLockFunc"
		array := garray.New(true)
		go func() {
			内存锁类.X非阻塞写锁定_函数(key, func() {
				array.Append(1)
				time.Sleep(200 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X非阻塞写锁定_函数(key, func() {
				array.Append(1)
			})
		}()
		go func() {
			time.Sleep(300 * time.Millisecond)
			内存锁类.X非阻塞写锁定_函数(key, func() {
				array.Append(1)
			})
		}()
		time.Sleep(150 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(400 * time.Millisecond)
		t.Assert(array.Len(), 2)
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
				内存锁类.X写锁定("test")
				defer 内存锁类.X退出写锁定("test")
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
				内存锁类.X写锁定("test")
				defer 内存锁类.X退出写锁定("test")
				time.Sleep(time.Millisecond)
			}()
		}
		for i := 0; i < num; i++ {
			go func() {
				defer wait.Done()
				<-ch
				内存锁类.X读锁定("test")
				defer 内存锁类.X退出读锁定("test")
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
		array := garray.New(true)
		go func() {
			内存锁类.X读锁定(key)
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			内存锁类.X退出读锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X写锁定(key)
			array.Append(1)
			内存锁类.X退出写锁定(key)
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})

	// Lock before RLock
	gtest.C(t, func(t *gtest.T) {
		key := "testLockBeforeRLock"
		array := garray.New(true)
		go func() {
			内存锁类.X写锁定(key)
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			内存锁类.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X读锁定(key)
			array.Append(1)
			内存锁类.X退出读锁定(key)
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})

	// 在RLocks之前加锁
	gtest.C(t, func(t *gtest.T) {
		key := "testLockBeforeRLocks"
		array := garray.New(true)
		go func() {
			内存锁类.X写锁定(key)
			array.Append(1)
			time.Sleep(300 * time.Millisecond)
			内存锁类.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X读锁定(key)
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			内存锁类.X退出读锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X读锁定(key)
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			内存锁类.X退出读锁定(key)
		}()
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 3)
	})
}

func Test_Locker_TryRLock(t *testing.T) {
	// 在尝试RLock之前加锁
	gtest.C(t, func(t *gtest.T) {
		key := "testLockBeforeTryRLock"
		array := garray.New(true)
		go func() {
			内存锁类.X写锁定(key)
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			内存锁类.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			if 内存锁类.X非阻塞读锁定(key) {
				array.Append(1)
				内存锁类.X退出读锁定(key)
			}
		}()
		time.Sleep(150 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})

	// 在尝试RLock之前加锁s
	gtest.C(t, func(t *gtest.T) {
		key := "testLockBeforeTryRLocks"
		array := garray.New(true)
		go func() {
			内存锁类.X写锁定(key)
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			内存锁类.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			if 内存锁类.X非阻塞读锁定(key) {
				array.Append(1)
				内存锁类.X退出读锁定(key)
			}
		}()
		go func() {
			time.Sleep(300 * time.Millisecond)
			if 内存锁类.X非阻塞读锁定(key) {
				array.Append(1)
				内存锁类.X退出读锁定(key)
			}
		}()
		time.Sleep(150 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func Test_Locker_RLockFunc(t *testing.T) {
	// RLockFunc 在 Lock 之前执行
	gtest.C(t, func(t *gtest.T) {
		key := "testRLockFuncBeforeLock"
		array := garray.New(true)
		go func() {
			内存锁类.X读锁定_函数(key, func() {
				array.Append(1)
				time.Sleep(200 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X写锁定(key)
			array.Append(1)
			内存锁类.X退出写锁定(key)
		}()
		time.Sleep(150 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})

	// 在RLockFunc之前加锁
	gtest.C(t, func(t *gtest.T) {
		key := "testLockBeforeRLockFunc"
		array := garray.New(true)
		go func() {
			内存锁类.X写锁定(key)
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			内存锁类.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X读锁定_函数(key, func() {
				array.Append(1)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})

	// 在RLockFunc之前加锁s
	gtest.C(t, func(t *gtest.T) {
		key := "testLockBeforeRLockFuncs"
		array := garray.New(true)
		go func() {
			内存锁类.X写锁定(key)
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			内存锁类.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X读锁定_函数(key, func() {
				array.Append(1)
				time.Sleep(200 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X读锁定_函数(key, func() {
				array.Append(1)
				time.Sleep(200 * time.Millisecond)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 3)
	})
}

func Test_Locker_TryRLockFunc(t *testing.T) {
	// 在尝试RLock之前加锁Func
	gtest.C(t, func(t *gtest.T) {
		key := "testLockBeforeTryRLockFunc"
		array := garray.New(true)
		go func() {
			内存锁类.X写锁定(key)
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			内存锁类.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X非阻塞读锁定_函数(key, func() {
				array.Append(1)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 1)
	})

	// 在尝试RLock之前加锁Funcs
	gtest.C(t, func(t *gtest.T) {
		key := "testLockBeforeTryRLockFuncs"
		array := garray.New(true)
		go func() {
			内存锁类.X写锁定(key)
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			内存锁类.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X非阻塞读锁定_函数(key, func() {
				array.Append(1)
			})
		}()
		go func() {
			time.Sleep(300 * time.Millisecond)
			内存锁类.X非阻塞读锁定_函数(key, func() {
				array.Append(1)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}
