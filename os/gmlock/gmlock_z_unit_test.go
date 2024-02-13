// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 内存锁类_test

import (
	"sync"
	"testing"
	"time"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/os/gmlock"
	"github.com/888go/goframe/test/gtest"
)

func Test_Locker_Lock(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		key := "testLock"
		array := 数组类.X创建(true)
		go func() {
			内存锁类.X写锁定(key)
			array.Append别名(1)
			time.Sleep(300 * time.Millisecond)
			内存锁类.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X写锁定(key)
			array.Append别名(1)
			内存锁类.X退出写锁定(key)
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
		内存锁类.X删除锁(key)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		key := "testLock"
		array := 数组类.X创建(true)
		lock := 内存锁类.X创建()
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
	单元测试类.C(t, func(t *单元测试类.T) {
		key := "testTryLock"
		array := 数组类.X创建(true)
		go func() {
			内存锁类.X写锁定(key)
			array.Append别名(1)
			time.Sleep(300 * time.Millisecond)
			内存锁类.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(150 * time.Millisecond)
			if 内存锁类.X非阻塞写锁定(key) {
				array.Append别名(1)
				内存锁类.X退出写锁定(key)
			}
		}()
		go func() {
			time.Sleep(400 * time.Millisecond)
			if 内存锁类.X非阻塞写锁定(key) {
				array.Append别名(1)
				内存锁类.X退出写锁定(key)
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
	单元测试类.C(t, func(t *单元测试类.T) {
		key := "testLockFunc"
		array := 数组类.X创建(true)
		go func() {
			内存锁类.X写锁定_函数(key, func() {
				array.Append别名(1)
				time.Sleep(300 * time.Millisecond)
			}) //
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X写锁定_函数(key, func() {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		key := "testTryLockFunc"
		array := 数组类.X创建(true)
		go func() {
			内存锁类.X非阻塞写锁定_函数(key, func() {
				array.Append别名(1)
				time.Sleep(200 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X非阻塞写锁定_函数(key, func() {
				array.Append别名(1)
			})
		}()
		go func() {
			time.Sleep(300 * time.Millisecond)
			内存锁类.X非阻塞写锁定_函数(key, func() {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		key := "testRLockBeforeLock"
		array := 数组类.X创建(true)
		go func() {
			内存锁类.X读锁定(key)
			array.Append别名(1)
			time.Sleep(200 * time.Millisecond)
			内存锁类.X退出读锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X写锁定(key)
			array.Append别名(1)
			内存锁类.X退出写锁定(key)
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})

	// Lock before RLock
	单元测试类.C(t, func(t *单元测试类.T) {
		key := "testLockBeforeRLock"
		array := 数组类.X创建(true)
		go func() {
			内存锁类.X写锁定(key)
			array.Append别名(1)
			time.Sleep(200 * time.Millisecond)
			内存锁类.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X读锁定(key)
			array.Append别名(1)
			内存锁类.X退出读锁定(key)
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})

	// 在RLocks之前加锁
	单元测试类.C(t, func(t *单元测试类.T) {
		key := "testLockBeforeRLocks"
		array := 数组类.X创建(true)
		go func() {
			内存锁类.X写锁定(key)
			array.Append别名(1)
			time.Sleep(300 * time.Millisecond)
			内存锁类.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X读锁定(key)
			array.Append别名(1)
			time.Sleep(200 * time.Millisecond)
			内存锁类.X退出读锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X读锁定(key)
			array.Append别名(1)
			time.Sleep(200 * time.Millisecond)
			内存锁类.X退出读锁定(key)
		}()
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 3)
	})
}

func Test_Locker_TryRLock(t *testing.T) {
	// 在尝试RLock之前加锁
	单元测试类.C(t, func(t *单元测试类.T) {
		key := "testLockBeforeTryRLock"
		array := 数组类.X创建(true)
		go func() {
			内存锁类.X写锁定(key)
			array.Append别名(1)
			time.Sleep(200 * time.Millisecond)
			内存锁类.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			if 内存锁类.X非阻塞读锁定(key) {
				array.Append别名(1)
				内存锁类.X退出读锁定(key)
			}
		}()
		time.Sleep(150 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})

	// 在尝试RLock之前加锁s
	单元测试类.C(t, func(t *单元测试类.T) {
		key := "testLockBeforeTryRLocks"
		array := 数组类.X创建(true)
		go func() {
			内存锁类.X写锁定(key)
			array.Append别名(1)
			time.Sleep(200 * time.Millisecond)
			内存锁类.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			if 内存锁类.X非阻塞读锁定(key) {
				array.Append别名(1)
				内存锁类.X退出读锁定(key)
			}
		}()
		go func() {
			time.Sleep(300 * time.Millisecond)
			if 内存锁类.X非阻塞读锁定(key) {
				array.Append别名(1)
				内存锁类.X退出读锁定(key)
			}
		}()
		time.Sleep(150 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}

func Test_Locker_RLockFunc(t *testing.T) {
	// RLockFunc 在 Lock 之前执行
	单元测试类.C(t, func(t *单元测试类.T) {
		key := "testRLockFuncBeforeLock"
		array := 数组类.X创建(true)
		go func() {
			内存锁类.X读锁定_函数(key, func() {
				array.Append别名(1)
				time.Sleep(200 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X写锁定(key)
			array.Append别名(1)
			内存锁类.X退出写锁定(key)
		}()
		time.Sleep(150 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})

	// 在RLockFunc之前加锁
	单元测试类.C(t, func(t *单元测试类.T) {
		key := "testLockBeforeRLockFunc"
		array := 数组类.X创建(true)
		go func() {
			内存锁类.X写锁定(key)
			array.Append别名(1)
			time.Sleep(200 * time.Millisecond)
			内存锁类.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X读锁定_函数(key, func() {
				array.Append别名(1)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})

	// 在RLockFunc之前加锁s
	单元测试类.C(t, func(t *单元测试类.T) {
		key := "testLockBeforeRLockFuncs"
		array := 数组类.X创建(true)
		go func() {
			内存锁类.X写锁定(key)
			array.Append别名(1)
			time.Sleep(200 * time.Millisecond)
			内存锁类.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X读锁定_函数(key, func() {
				array.Append别名(1)
				time.Sleep(200 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X读锁定_函数(key, func() {
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
	// 在尝试RLock之前加锁Func
	单元测试类.C(t, func(t *单元测试类.T) {
		key := "testLockBeforeTryRLockFunc"
		array := 数组类.X创建(true)
		go func() {
			内存锁类.X写锁定(key)
			array.Append别名(1)
			time.Sleep(200 * time.Millisecond)
			内存锁类.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X非阻塞读锁定_函数(key, func() {
				array.Append别名(1)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
	})

	// 在尝试RLock之前加锁Funcs
	单元测试类.C(t, func(t *单元测试类.T) {
		key := "testLockBeforeTryRLockFuncs"
		array := 数组类.X创建(true)
		go func() {
			内存锁类.X写锁定(key)
			array.Append别名(1)
			time.Sleep(200 * time.Millisecond)
			内存锁类.X退出写锁定(key)
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			内存锁类.X非阻塞读锁定_函数(key, func() {
				array.Append别名(1)
			})
		}()
		go func() {
			time.Sleep(300 * time.Millisecond)
			内存锁类.X非阻塞读锁定_函数(key, func() {
				array.Append别名(1)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.X取长度(), 1)
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.X取长度(), 2)
	})
}
