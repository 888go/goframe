// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gmutex_test
import (
	"context"
	"testing"
	"time"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/os/gmutex"
	"github.com/888go/goframe/test/gtest"
	)

func Test_RWMutex_RUnlock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		mu := gmutex.RWMutex{}
		mu.RLockFunc(func() {
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
			mu.LockFunc(func() {
				time.Sleep(200 * time.Millisecond)
			})
		}()
		time.Sleep(100 * time.Millisecond)

		go func() {
			mu.RLockFunc(func() {
				time.Sleep(200 * time.Millisecond)
			})
		}()
	})
}

func Test_RWMutex_Unlock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		mu := gmutex.RWMutex{}
		array := garray.New(true)
		go func() {
			mu.LockFunc(func() {
				array.Append(1)
				time.Sleep(300 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.LockFunc(func() {
				array.Append(1)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.LockFunc(func() {
				array.Append(1)
			})
		}()

		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(400 * time.Millisecond)
		t.Assert(array.Len(), 3)
	})
}

func Test_RWMutex_LockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		mu := gmutex.RWMutex{}
		array := garray.New(true)
		go func() {
			mu.LockFunc(func() {
				array.Append(1)
				time.Sleep(300 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.LockFunc(func() {
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

func Test_RWMutex_TryLockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		mu := gmutex.RWMutex{}
		array := garray.New(true)
		go func() {
			mu.LockFunc(func() {
				array.Append(1)
				time.Sleep(300 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.TryLockFunc(func() {
				array.Append(1)
			})
		}()
		go func() {
			time.Sleep(400 * time.Millisecond)
			mu.TryLockFunc(func() {
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

func Test_RWMutex_RLockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		mu := gmutex.RWMutex{}
		array := garray.New(true)
		go func() {
			mu.LockFunc(func() {
				array.Append(1)
				time.Sleep(300 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.RLockFunc(func() {
				array.Append(1)
				time.Sleep(100 * time.Millisecond)
			})
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})

	gtest.C(t, func(t *gtest.T) {
		mu := gmutex.RWMutex{}
		array := garray.New(true)
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.RLockFunc(func() {
				array.Append(1)
				time.Sleep(100 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.RLockFunc(func() {
				array.Append(1)
				time.Sleep(100 * time.Millisecond)
			})
		}()
		go func() {
			time.Sleep(100 * time.Millisecond)
			mu.RLockFunc(func() {
				array.Append(1)
				time.Sleep(100 * time.Millisecond)
			})
		}()
		t.Assert(array.Len(), 0)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 3)
	})
}

func Test_RWMutex_TryRLockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			mu    = gmutex.RWMutex{}
			array = garray.New(true)
		)
		// 首次写入锁
		go func() {
			mu.LockFunc(func() {
				array.Append(1)
				glog.Print(context.TODO(), "lock1 done")
				time.Sleep(2000 * time.Millisecond)
			})
		}()
		// 这个goroutine永远不会获取到锁。
		go func() {
			time.Sleep(1000 * time.Millisecond)
			mu.TryRLockFunc(func() {
				array.Append(1)
			})
		}()
		for index := 0; index < 1000; index++ {
			go func() {
				time.Sleep(4000 * time.Millisecond)
				mu.TryRLockFunc(func() {
					array.Append(1)
				})
			}()
		}
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(1000 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(2000 * time.Millisecond)
		t.Assert(array.Len(), 1001)
	})
}
