// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// go test *.go

package 数组类_test

import (
	"testing"
	"time"
	
	"github.com/888go/goframe/internal/empty"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func TestNewSortedIntArrayComparator(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{0, 3, 2, 1, 4, 5, 6}
		array1 := 数组类.X创建整数排序并带排序函数(func(a, b int) int {
			return a - b
		}, true)
		array1.Append别名(a1...)
		t.Assert(array1.X取长度(), 7)
		t.Assert(array1.X取any数组(), []int{0, 1, 2, 3, 4, 5, 6})
	})
}

func TestNewSortedIntArrayRange(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array1 := 数组类.X创建整数排序并按范围(1, 5, 1)
		t.Assert(array1.X取长度(), 5)
		t.Assert(array1.X取any数组(), []int{1, 2, 3, 4, 5})
	})
}

func TestNewSortedIntArrayFrom(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{0, 3, 2, 1, 4, 5, 6}
		array1 := 数组类.X创建整数排序并从数组(a1, true)
		t.Assert(array1.X连接("."), "0.1.2.3.4.5.6")
		t.Assert(array1.X取切片(), a1)
		t.Assert(array1.X取any数组(), a1)
	})
}

func TestNewSortedIntArrayFromCopy(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{0, 5, 2, 1, 4, 3, 6}
		array1 := 数组类.X创建整数排序并从数组复制(a1, false)
		t.Assert(array1.X连接("."), "0.1.2.3.4.5.6")
	})
}

func TestSortedIntArray_At(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{0, 3, 2, 1}

		array1 := 数组类.X创建整数排序并从数组(a1)
		v := array1.X取值(1)

		t.Assert(v, 1)
	})
}

func TestSortedIntArray_SetArray(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{0, 1, 2, 3}
		a2 := []int{4, 5, 6}
		array1 := 数组类.X创建整数排序并从数组(a1)
		array2 := array1.X设置数组(a2)

		t.Assert(array2.X取长度(), 3)
		t.Assert(array2.X查找(3), -1)
		t.Assert(array2.X查找(5), 1)
		t.Assert(array2.X查找(6), 2)
	})
}

func TestSortedIntArray_Sort(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{0, 3, 2, 1}

		array1 := 数组类.X创建整数排序并从数组(a1)
		array2 := array1.X排序递增()

		t.Assert(array2.X取长度(), 4)
		t.Assert(array2, []int{0, 1, 2, 3})
	})
}

func TestSortedIntArray_Get(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 3, 5, 0}
		array1 := 数组类.X创建整数排序并从数组(a1)
		v, ok := array1.X取值2(0)
		t.Assert(v, 0)
		t.Assert(ok, true)

		v, ok = array1.X取值2(1)
		t.Assert(v, 1)
		t.Assert(ok, true)

		v, ok = array1.X取值2(3)
		t.Assert(v, 5)
		t.Assert(ok, true)

		v, ok = array1.X取值2(99)
		t.Assert(v, 0)
		t.Assert(ok, false)
	})
}

func TestSortedIntArray_Remove(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 3, 5, 0}
		array1 := 数组类.X创建整数排序并从数组(a1)

		v, ok := array1.X删除(-1)
		t.Assert(v, 0)
		t.Assert(ok, false)

		v, ok = array1.X删除(-100000)
		t.Assert(v, 0)
		t.Assert(ok, false)

		v, ok = array1.X删除(2)
		t.Assert(v, 3)
		t.Assert(ok, true)

		t.Assert(array1.X查找(5), 2)

		v, ok = array1.X删除(0)
		t.Assert(v, 0)
		t.Assert(ok, true)

		t.Assert(array1.X查找(5), 1)

		a2 := []int{1, 3, 4}
		array2 := 数组类.X创建整数排序并从数组(a2)

		v, ok = array2.X删除(1)
		t.Assert(v, 3)
		t.Assert(ok, true)
		t.Assert(array2.X查找(1), 0)

		v, ok = array2.X删除(1)
		t.Assert(v, 4)
		t.Assert(ok, true)

		t.Assert(array2.X查找(4), -1)
	})
}

func TestSortedIntArray_PopLeft(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 3, 5, 2}
		array1 := 数组类.X创建整数排序并从数组(a1)
		v, ok := array1.X出栈左()
		t.Assert(v, 1)
		t.Assert(ok, true)
		t.Assert(array1.X取长度(), 3)
		t.Assert(array1.X查找(1), -1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建整数排序并从数组(g.SliceInt别名{1, 2, 3})
		v, ok := array.X出栈左()
		t.Assert(v, 1)
		t.Assert(ok, true)
		t.Assert(array.X取长度(), 2)
		v, ok = array.X出栈左()
		t.Assert(v, 2)
		t.Assert(ok, true)
		t.Assert(array.X取长度(), 1)
		v, ok = array.X出栈左()
		t.Assert(v, 3)
		t.Assert(ok, true)
		t.Assert(array.X取长度(), 0)
	})
}

func TestSortedIntArray_PopRight(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 3, 5, 2}
		array1 := 数组类.X创建整数排序并从数组(a1)
		v, ok := array1.X出栈右()
		t.Assert(v, 5)
		t.Assert(ok, true)
		t.Assert(array1.X取长度(), 3)
		t.Assert(array1.X查找(5), -1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建整数排序并从数组(g.SliceInt别名{1, 2, 3})
		v, ok := array.X出栈右()
		t.Assert(v, 3)
		t.Assert(ok, true)
		t.Assert(array.X取长度(), 2)

		v, ok = array.X出栈右()
		t.Assert(v, 2)
		t.Assert(ok, true)
		t.Assert(array.X取长度(), 1)

		v, ok = array.X出栈右()
		t.Assert(v, 1)
		t.Assert(ok, true)
		t.Assert(array.X取长度(), 0)
	})
}

func TestSortedIntArray_PopRand(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 3, 5, 2}
		array1 := 数组类.X创建整数排序并从数组(a1)
		i1, ok := array1.X出栈随机()
		t.Assert(ok, true)
		t.Assert(array1.X取长度(), 3)
		t.Assert(array1.X查找(i1), -1)
		t.AssertIN(i1, []int{1, 3, 5, 2})
	})
}

func TestSortedIntArray_PopRands(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 3, 5, 2}
		array1 := 数组类.X创建整数排序并从数组(a1)
		ns1 := array1.X出栈随机多个(2)
		t.Assert(array1.X取长度(), 2)
		t.AssertIN(ns1, []int{1, 3, 5, 2})

		a2 := []int{1, 3, 5, 2}
		array2 := 数组类.X创建整数排序并从数组(a2)
		ns2 := array2.X出栈随机多个(5)
		t.Assert(array2.X取长度(), 0)
		t.Assert(len(ns2), 4)
		t.AssertIN(ns2, []int{1, 3, 5, 2})
	})
}

func TestSortedIntArray_Empty(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建整数排序()
		v, ok := array.X出栈左()
		t.Assert(v, 0)
		t.Assert(ok, false)
		t.Assert(array.X出栈左多个(10), nil)

		v, ok = array.X出栈右()
		t.Assert(v, 0)
		t.Assert(ok, false)
		t.Assert(array.X出栈右多个(10), nil)

		v, ok = array.X出栈随机()
		t.Assert(v, 0)
		t.Assert(ok, false)
		t.Assert(array.X出栈随机多个(10), nil)
	})
}

func TestSortedIntArray_PopLefts(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 3, 5, 2}
		array1 := 数组类.X创建整数排序并从数组(a1)
		ns1 := array1.X出栈左多个(2)
		t.Assert(array1.X取长度(), 2)
		t.Assert(ns1, []int{1, 2})

		a2 := []int{1, 3, 5, 2}
		array2 := 数组类.X创建整数排序并从数组(a2)
		ns2 := array2.X出栈左多个(5)
		t.Assert(array2.X取长度(), 0)
		t.AssertIN(ns2, []int{1, 3, 5, 2})
	})
}

func TestSortedIntArray_PopRights(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 3, 5, 2}
		array1 := 数组类.X创建整数排序并从数组(a1)
		ns1 := array1.X出栈右多个(2)
		t.Assert(array1.X取长度(), 2)
		t.Assert(ns1, []int{3, 5})

		a2 := []int{1, 3, 5, 2}
		array2 := 数组类.X创建整数排序并从数组(a2)
		ns2 := array2.X出栈右多个(5)
		t.Assert(array2.X取长度(), 0)
		t.AssertIN(ns2, []int{1, 3, 5, 2})
	})
}

func TestSortedIntArray_Range(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 3, 5, 2, 6, 7}
		array1 := 数组类.X创建整数排序并从数组(a1)
		array2 := 数组类.X创建整数排序并从数组(a1, true)
		ns1 := array1.X取切片并按范围(1, 4)
		t.Assert(len(ns1), 3)
		t.Assert(ns1, []int{2, 3, 5})

		ns2 := array1.X取切片并按范围(5, 4)
		t.Assert(len(ns2), 0)

		ns3 := array1.X取切片并按范围(-1, 4)
		t.Assert(len(ns3), 4)

		nsl := array1.X取切片并按范围(5, 8)
		t.Assert(len(nsl), 1)
		t.Assert(array2.X取切片并按范围(1, 2), []int{2})
	})
}

func TestSortedIntArray_Sum(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 3, 5}
		array1 := 数组类.X创建整数排序并从数组(a1)
		n1 := array1.X求和()
		t.Assert(n1, 9)
	})
}

func TestSortedIntArray_Join(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 3, 5}
		array1 := 数组类.X创建整数排序并从数组(a1)
		t.Assert(array1.X连接("."), `1.3.5`)

		array2 := 数组类.X创建整数排序并从数组([]int{})
		t.Assert(array2.X连接("."), "")
	})
}

func TestSortedIntArray_String(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 3, 5}
		array1 := 数组类.X创建整数排序并从数组(a1)
		t.Assert(array1.String(), `[1,3,5]`)

		array1 = nil
		t.Assert(array1.String(), "")
	})
}

func TestSortedIntArray_Contains(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 3, 5}
		array1 := 数组类.X创建整数排序并从数组(a1)
		t.Assert(array1.X是否存在(4), false)
	})
}

func TestSortedIntArray_Clone(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 3, 5}
		array1 := 数组类.X创建整数排序并从数组(a1)
		array2 := array1.X取副本()
		t.Assert(array2.X取长度(), 3)
		t.Assert(array2, array1)
	})
}

func TestSortedIntArray_Clear(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 3, 5}
		array1 := 数组类.X创建整数排序并从数组(a1)
		array1.X清空()
		t.Assert(array1.X取长度(), 0)
	})
}

func TestSortedIntArray_Chunk(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 2, 3, 4, 5}
		array1 := 数组类.X创建整数排序并从数组(a1)
		ns1 := array1.X分割(2) // 按每几个元素切成一个数组
		ns2 := array1.X分割(-1)
		t.Assert(len(ns1), 3)
		t.Assert(ns1[0], []int{1, 2})
		t.Assert(ns1[2], []int{5})
		t.Assert(len(ns2), 0)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 2, 3, 4, 5}
		array1 := 数组类.X创建整数排序并从数组(a1)
		chunks := array1.X分割(3)
		t.Assert(len(chunks), 2)
		t.Assert(chunks[0], []int{1, 2, 3})
		t.Assert(chunks[1], []int{4, 5})
		t.Assert(array1.X分割(0), nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 2, 3, 4, 5, 6}
		array1 := 数组类.X创建整数排序并从数组(a1)
		chunks := array1.X分割(2)
		t.Assert(len(chunks), 3)
		t.Assert(chunks[0], []int{1, 2})
		t.Assert(chunks[1], []int{3, 4})
		t.Assert(chunks[2], []int{5, 6})
		t.Assert(array1.X分割(0), nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 2, 3, 4, 5, 6}
		array1 := 数组类.X创建整数排序并从数组(a1)
		chunks := array1.X分割(3)
		t.Assert(len(chunks), 2)
		t.Assert(chunks[0], []int{1, 2, 3})
		t.Assert(chunks[1], []int{4, 5, 6})
		t.Assert(array1.X分割(0), nil)
	})
}

func TestSortedIntArray_SubSlice(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 2, 3, 4, 5}
		array1 := 数组类.X创建整数排序并从数组(a1)
		array2 := 数组类.X创建整数排序并从数组(a1, true)
		ns1 := array1.X取切片并按数量(1, 2)
		t.Assert(len(ns1), 2)
		t.Assert(ns1, []int{2, 3})

		ns2 := array1.X取切片并按数量(7, 2)
		t.Assert(len(ns2), 0)

		ns3 := array1.X取切片并按数量(3, 5)
		t.Assert(len(ns3), 2)
		t.Assert(ns3, []int{4, 5})

		ns4 := array1.X取切片并按数量(3, 1)
		t.Assert(len(ns4), 1)
		t.Assert(ns4, []int{4})
		t.Assert(array1.X取切片并按数量(-1, 1), []int{5})
		t.Assert(array1.X取切片并按数量(-9, 1), nil)
		t.Assert(array1.X取切片并按数量(1, -9), nil)
		t.Assert(array2.X取切片并按数量(1, 2), []int{2, 3})
	})
}

func TestSortedIntArray_Rand(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 2, 3, 4, 5}
		array1 := 数组类.X创建整数排序并从数组(a1)
		ns1, ok := array1.X取值随机()
		t.AssertIN(ns1, a1)
		t.Assert(ok, true)

		array2 := 数组类.X创建整数排序并从数组([]int{})
		ns2, ok := array2.X取值随机()
		t.Assert(ns2, 0)
		t.Assert(ok, false)
	})
}

func TestSortedIntArray_Rands(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 2, 3, 4, 5}
		array1 := 数组类.X创建整数排序并从数组(a1)
		ns1 := array1.X取值随机多个(2)
		t.AssertIN(ns1, a1)
		t.Assert(len(ns1), 2)

		ns2 := array1.X取值随机多个(6)
		t.Assert(len(ns2), 6)

		array2 := 数组类.X创建整数排序并从数组([]int{})
		val := array2.X取值随机多个(1)
		t.Assert(val, nil)
	})
}

func TestSortedIntArray_CountValues(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 2, 3, 4, 5, 3}
		array1 := 数组类.X创建整数排序并从数组(a1)
		ns1 := array1.X统计() // 按每几个元素切成一个数组
		t.Assert(len(ns1), 5)
		t.Assert(ns1[2], 1)
		t.Assert(ns1[3], 2)
	})
}

func TestSortedIntArray_SetUnique(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 2, 3, 4, 5, 3, 2, 2, 3, 5, 5}
		array1 := 数组类.X创建整数排序并从数组(a1)
		array1.X设置去重(true)
		t.Assert(array1.X取长度(), 5)
		t.Assert(array1, []int{1, 2, 3, 4, 5})
	})
}

func TestSortedIntArray_Unique(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []int{1, 2, 3, 4, 5, 3, 2, 2, 3, 5, 5}
		array1 := 数组类.X创建整数排序并从数组(a1)
		array1.X去重()
		t.Assert(array1.X取长度(), 5)
		t.Assert(array1, []int{1, 2, 3, 4, 5})

		array2 := 数组类.X创建整数排序并从数组([]int{})
		array2.X去重()
		t.Assert(array2.X取长度(), 0)
	})
}

func TestSortedIntArray_LockFunc(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s1 := []int{1, 2, 3, 4}
		a1 := 数组类.X创建整数排序并从数组(s1, true)
		ch1 := make(chan int64, 3)
		ch2 := make(chan int64, 3)
		// go1
		go a1.X遍历写锁定(func(n1 []int) { // 读写锁
			time.Sleep(2 * time.Second) // 暂停2秒
			n1[2] = 6
			ch2 <- 转换类.X取整数64位(time.Now().UnixNano() / 1000 / 1000)
		})

		// go2
		go func() {
			time.Sleep(100 * time.Millisecond) // 故意暂停0.01秒,等go1执行锁后，再开始执行.
			ch1 <- 转换类.X取整数64位(time.Now().UnixNano() / 1000 / 1000)
			a1.X取长度()
			ch1 <- 转换类.X取整数64位(time.Now().UnixNano() / 1000 / 1000)
		}()

		t1 := <-ch1
		t2 := <-ch1
		<-ch2 // 等待go1完成

		// 防止ci抖动,以豪秒为单位
		t.AssertGT(t2-t1, 20) // go1加的读写互斥锁，所go2读的时候被阻塞。
		t.Assert(a1.X是否存在(6), true)
	})
}

func TestSortedIntArray_RLockFunc(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s1 := []int{1, 2, 3, 4}
		a1 := 数组类.X创建整数排序并从数组(s1, true)

		ch1 := make(chan int64, 3)
		ch2 := make(chan int64, 1)
		// go1
		go a1.X遍历读锁定(func(n1 []int) { // 读锁
			time.Sleep(2 * time.Second) // 暂停1秒
			n1[2] = 6
			ch2 <- 转换类.X取整数64位(time.Now().UnixNano() / 1000 / 1000)
		})

		// go2
		go func() {
			time.Sleep(100 * time.Millisecond) // 故意暂停0.01秒,等go1执行锁后，再开始执行.
			ch1 <- 转换类.X取整数64位(time.Now().UnixNano() / 1000 / 1000)
			a1.X取长度()
			ch1 <- 转换类.X取整数64位(time.Now().UnixNano() / 1000 / 1000)
		}()

		t1 := <-ch1
		t2 := <-ch1
		<-ch2 // 等待go1完成

		// 防止ci抖动,以豪秒为单位
		t.AssertLT(t2-t1, 20) // go1加的读锁，所go2读的时候，并没有阻塞。
		t.Assert(a1.X是否存在(6), true)
	})
}

func TestSortedIntArray_Merge(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		func1 := func(v1, v2 interface{}) int {
			if 转换类.X取整数(v1) < 转换类.X取整数(v2) {
				return 0
			}
			return 1
		}
		i0 := []int{1, 2, 3, 4}
		s2 := []string{"e", "f"}
		i1 := 数组类.X创建整数并从数组([]int{1, 2, 3})
		i2 := 数组类.X创建并从数组([]interface{}{3})
		s3 := 数组类.X创建文本并从数组([]string{"g", "h"})
		s4 := 数组类.X创建排序并从数组([]interface{}{4, 5}, func1)
		s5 := 数组类.X创建文本排序并从数组(s2)
		s6 := 数组类.X创建整数排序并从数组([]int{1, 2, 3})
		a1 := 数组类.X创建整数排序并从数组(i0)

		t.Assert(a1.X合并(s2).X取长度(), 6)
		t.Assert(a1.X合并(i1).X取长度(), 9)
		t.Assert(a1.X合并(i2).X取长度(), 10)
		t.Assert(a1.X合并(s3).X取长度(), 12)
		t.Assert(a1.X合并(s4).X取长度(), 14)
		t.Assert(a1.X合并(s5).X取长度(), 16)
		t.Assert(a1.X合并(s6).X取长度(), 19)
	})
}

func TestSortedIntArray_Json(t *testing.T) {
	// array pointer
	单元测试类.C(t, func(t *单元测试类.T) {
		s1 := []int{1, 4, 3, 2}
		s2 := []int{1, 2, 3, 4}
		a1 := 数组类.X创建整数排序并从数组(s1)
		b1, err1 := json.Marshal(a1)
		b2, err2 := json.Marshal(s1)
		t.Assert(b1, b2)
		t.Assert(err1, err2)

		a2 := 数组类.X创建整数排序()
		err1 = json.UnmarshalUseNumber(b2, &a2)
		t.AssertNil(err1)
		t.Assert(a2.X取切片(), s2)

		var a3 数组类.SortedIntArray
		err := json.UnmarshalUseNumber(b2, &a3)
		t.AssertNil(err)
		t.Assert(a3.X取切片(), s1)
	})
	// array value
	单元测试类.C(t, func(t *单元测试类.T) {
		s1 := []int{1, 4, 3, 2}
		s2 := []int{1, 2, 3, 4}
		a1 := *数组类.X创建整数排序并从数组(s1)
		b1, err1 := json.Marshal(a1)
		b2, err2 := json.Marshal(s1)
		t.Assert(b1, b2)
		t.Assert(err1, err2)

		a2 := 数组类.X创建整数排序()
		err1 = json.UnmarshalUseNumber(b2, &a2)
		t.AssertNil(err1)
		t.Assert(a2.X取切片(), s2)

		var a3 数组类.SortedIntArray
		err := json.UnmarshalUseNumber(b2, &a3)
		t.AssertNil(err)
		t.Assert(a3.X取切片(), s1)
	})
	// array pointer
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Name   string
			Scores *数组类.SortedIntArray
		}
		data := g.Map{
			"Name":   "john",
			"Scores": []int{99, 100, 98},
		}
		b, err := json.Marshal(data)
		t.AssertNil(err)

		user := new(User)
		err = json.UnmarshalUseNumber(b, user)
		t.AssertNil(err)
		t.Assert(user.Name, data["Name"])
		t.Assert(user.Scores, []int{98, 99, 100})
	})
	// array value
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Name   string
			Scores 数组类.SortedIntArray
		}
		data := g.Map{
			"Name":   "john",
			"Scores": []int{99, 100, 98},
		}
		b, err := json.Marshal(data)
		t.AssertNil(err)

		user := new(User)
		err = json.UnmarshalUseNumber(b, user)
		t.AssertNil(err)
		t.Assert(user.Name, data["Name"])
		t.Assert(user.Scores, []int{98, 99, 100})
	})
}

func TestSortedIntArray_Iterator(t *testing.T) {
	slice := g.SliceInt别名{10, 20, 30, 40}
	array := 数组类.X创建整数排序并从数组(slice)
	单元测试类.C(t, func(t *单元测试类.T) {
		array.X遍历(func(k int, v int) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array.X遍历升序(func(k int, v int) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array.X遍历降序(func(k int, v int) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		index := 0
		array.X遍历(func(k int, v int) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		index := 0
		array.X遍历升序(func(k int, v int) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		index := 0
		array.X遍历降序(func(k int, v int) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
}

func TestSortedIntArray_RemoveValue(t *testing.T) {
	slice := g.SliceInt别名{10, 20, 30, 40}
	array := 数组类.X创建整数排序并从数组(slice)
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(array.X删除值(99), false)
		t.Assert(array.X删除值(20), true)
		t.Assert(array.X删除值(10), true)
		t.Assert(array.X删除值(20), false)
		t.Assert(array.X删除值(88), false)
		t.Assert(array.X取长度(), 2)
	})
}

func TestSortedIntArray_RemoveValues(t *testing.T) {
	slice := g.SliceInt别名{10, 20, 30, 40}
	array := 数组类.X创建整数排序并从数组(slice)
	单元测试类.C(t, func(t *单元测试类.T) {
		array.X删除多个值(10, 40, 20)
		t.Assert(array.X取切片(), g.SliceInt别名{30})
	})
}

func TestSortedIntArray_UnmarshalValue(t *testing.T) {
	type V struct {
		Name  string
		Array *数组类.SortedIntArray
	}
	// JSON
	单元测试类.C(t, func(t *单元测试类.T) {
		var v *V
		err := 转换类.Struct(g.Map{
			"name":  "john",
			"array": []byte(`[2,3,1]`),
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Array.X取切片(), g.Slice别名{1, 2, 3})
	})
	// Map
	单元测试类.C(t, func(t *单元测试类.T) {
		var v *V
		err := 转换类.Struct(g.Map{
			"name":  "john",
			"array": g.Slice别名{2, 3, 1},
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Array.X取切片(), g.Slice别名{1, 2, 3})
	})
}
func TestSortedIntArray_Filter(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建整数排序并从数组(g.SliceInt别名{0, 1, 2, 3, 4, 0})
		t.Assert(array.X遍历删除(func(index int, value int) bool {
			return empty.IsEmpty(value)
		}), g.SliceInt别名{1, 2, 3, 4})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建整数排序并从数组(g.SliceInt别名{1, 2, 3, 4})
		t.Assert(array.X遍历删除(func(index int, value int) bool {
			return empty.IsEmpty(value)
		}), g.SliceInt别名{1, 2, 3, 4})
	})
}

func TestSortedIntArray_FilterEmpty(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建整数排序并从数组(g.SliceInt别名{0, 1, 2, 3, 4, 0})
		t.Assert(array.X删除所有空值(), g.SliceInt别名{1, 2, 3, 4})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建整数排序并从数组(g.SliceInt别名{1, 2, 3, 4})
		t.Assert(array.X删除所有空值(), g.SliceInt别名{1, 2, 3, 4})
	})
}

func TestSortedIntArray_Walk(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建整数排序并从数组(g.SliceInt别名{1, 2})
		t.Assert(array.X遍历修改(func(value int) int {
			return 10 + value
		}), g.Slice别名{11, 12})
	})
}

func TestSortedIntArray_IsEmpty(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建整数排序并从数组([]int{})
		t.Assert(array.X是否为空(), true)
	})
}

func TestSortedIntArray_DeepCopy(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建整数排序并从数组([]int{1, 2, 3, 4, 5})
		copyArray := array.DeepCopy().(*数组类.SortedIntArray)
		array.X入栈右(6)
		copyArray.X入栈右(7)
		cval, _ := copyArray.X取值2(5)
		val, _ := array.X取值2(5)
		t.AssertNE(cval, val)
	})
}
