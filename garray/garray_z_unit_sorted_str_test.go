// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// go test *.go

package 切片类_test

import (
	"testing"
	"time"

	"github.com/888go/goframe/garray/internal/empty"

	"github.com/888go/goframe/garray"
	"github.com/888go/goframe/garray/internal/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

func TestNewSortedStrArrayComparator(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"a", "d", "c", "b"}
		s1 := 切片类.X创建文本排序并带排序函数(func(a, b string) int {
			return gstr.Compare(a, b)
		})
		s1.X入栈右(a1...)
		t.Assert(s1.X取长度(), 4)
		t.Assert(s1, []string{"a", "b", "c", "d"})
	})
}

func TestNewSortedStrArrayFrom(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"a", "d", "c", "b"}
		s1 := 切片类.X创建文本排序并从切片(a1, true)
		t.Assert(s1, []string{"a", "b", "c", "d"})
		s2 := 切片类.X创建文本排序并从切片(a1, false)
		t.Assert(s2, []string{"a", "b", "c", "d"})
	})
}

func TestNewSortedStrArrayFromCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"a", "d", "c", "b"}
		s1 := 切片类.X创建文本排序并从切片复制(a1, true)
		t.Assert(s1.X取长度(), 4)
		t.Assert(s1, []string{"a", "b", "c", "d"})
	})
}

func TestSortedStrArray_SetArray(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"a", "d", "c", "b"}
		a2 := []string{"f", "g", "h"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		array1.X设置切片(a2)
		t.Assert(array1.X取长度(), 3)
		t.Assert(array1.X是否存在("d"), false)
		t.Assert(array1.X是否存在("b"), false)
		t.Assert(array1.X是否存在("g"), true)
	})
}

func TestSortedStrArray_ContainsI(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := 切片类.X创建文本排序()
		s.Append别名("a", "b", "C")
		t.Assert(s.X是否存在("A"), false)
		t.Assert(s.X是否存在("a"), true)
		t.Assert(s.X是否存在并忽略大小写("A"), true)

		s = 切片类.X创建文本排序()
		t.Assert(s.X是否存在("A"), false)
	})
}

func TestSortedStrArray_Sort(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"a", "d", "c", "b"}
		array1 := 切片类.X创建文本排序并从切片(a1)

		t.Assert(array1, []string{"a", "b", "c", "d"})
		array1.X排序递增()
		t.Assert(array1.X取长度(), 4)
		t.Assert(array1.X是否存在("c"), true)
		t.Assert(array1, []string{"a", "b", "c", "d"})
	})
}

func TestSortedStrArray_Get(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"a", "d", "c", "b"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		v, ok := array1.X取值2(2)
		t.Assert(v, "c")
		t.Assert(ok, true)

		v, ok = array1.X取值2(0)
		t.Assert(v, "a")
		t.Assert(ok, true)

		v, ok = array1.X取值2(99)
		t.Assert(v, "")
		t.Assert(ok, false)
	})
}

func TestSortedStrArray_Remove(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"a", "d", "c", "b"}
		array1 := 切片类.X创建文本排序并从切片(a1)

		v, ok := array1.X删除(-1)
		t.Assert(v, "")
		t.Assert(ok, false)

		v, ok = array1.X删除(100000)
		t.Assert(v, "")
		t.Assert(ok, false)

		v, ok = array1.X删除(2)
		t.Assert(v, "c")
		t.Assert(ok, true)

		v, ok = array1.X取值2(2)
		t.Assert(v, "d")
		t.Assert(ok, true)

		t.Assert(array1.X取长度(), 3)
		t.Assert(array1.X是否存在("c"), false)

		v, ok = array1.X删除(0)
		t.Assert(v, "a")
		t.Assert(ok, true)

		t.Assert(array1.X取长度(), 2)
		t.Assert(array1.X是否存在("a"), false)

		v, ok = array1.X删除(1)
		t.Assert(v, "d")
		t.Assert(ok, true)

		t.Assert(array1.X取长度(), 1)
	})
}

func TestSortedStrArray_PopLeft(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"e", "a", "d", "c", "b"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		v, ok := array1.X出栈左()
		t.Assert(v, "a")
		t.Assert(ok, true)
		t.Assert(array1.X取长度(), 4)
		t.Assert(array1.X是否存在("a"), false)
	})
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建文本排序并从切片(g.SliceStr{"1", "2", "3"})
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

func TestSortedStrArray_PopRight(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"e", "a", "d", "c", "b"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		v, ok := array1.X出栈右()
		t.Assert(v, "e")
		t.Assert(ok, ok)
		t.Assert(array1.X取长度(), 4)
		t.Assert(array1.X是否存在("e"), false)
	})
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建文本排序并从切片(g.SliceStr{"1", "2", "3"})
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

func TestSortedStrArray_PopRand(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"e", "a", "d", "c", "b"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		s1, ok := array1.X出栈随机()
		t.Assert(ok, true)
		t.AssertIN(s1, []string{"e", "a", "d", "c", "b"})
		t.Assert(array1.X取长度(), 4)
		t.Assert(array1.X是否存在(s1), false)
	})
}

func TestSortedStrArray_PopRands(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"e", "a", "d", "c", "b"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		s1 := array1.X出栈随机多个(2)
		t.AssertIN(s1, []string{"e", "a", "d", "c", "b"})
		t.Assert(array1.X取长度(), 3)
		t.Assert(len(s1), 2)

		s1 = array1.X出栈随机多个(4)
		t.Assert(len(s1), 3)
		t.AssertIN(s1, []string{"e", "a", "d", "c", "b"})
	})
}

func TestSortedStrArray_Empty(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建文本排序()
		v, ok := array.X出栈左()
		t.Assert(v, "")
		t.Assert(ok, false)
		t.Assert(array.X出栈左多个(10), nil)

		v, ok = array.X出栈右()
		t.Assert(v, "")
		t.Assert(ok, false)
		t.Assert(array.X出栈右多个(10), nil)

		v, ok = array.X出栈随机()
		t.Assert(v, "")
		t.Assert(ok, false)
		t.Assert(array.X出栈随机多个(10), nil)
	})
}

func TestSortedStrArray_PopLefts(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"e", "a", "d", "c", "b"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		s1 := array1.X出栈左多个(2)
		t.Assert(s1, []string{"a", "b"})
		t.Assert(array1.X取长度(), 3)
		t.Assert(len(s1), 2)

		s1 = array1.X出栈左多个(4)
		t.Assert(len(s1), 3)
		t.Assert(s1, []string{"c", "d", "e"})
		t.Assert(array1.X取长度(), 0)
	})
}

func TestSortedStrArray_PopRights(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"e", "a", "d", "c", "b", "f", "g"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		s1 := array1.X出栈右多个(2)
		t.Assert(s1, []string{"f", "g"})
		t.Assert(array1.X取长度(), 5)
		t.Assert(len(s1), 2)
		s1 = array1.X出栈右多个(6)
		t.Assert(len(s1), 5)
		t.Assert(s1, []string{"a", "b", "c", "d", "e"})
		t.Assert(array1.X取长度(), 0)
	})
}

func TestSortedStrArray_Range(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"e", "a", "d", "c", "b", "f", "g"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		array2 := 切片类.X创建文本排序并从切片(a1, true)
		s1 := array1.X取切片并按范围(2, 4)
		t.Assert(len(s1), 2)
		t.Assert(s1, []string{"c", "d"})

		s1 = array1.X取切片并按范围(-1, 2)
		t.Assert(len(s1), 2)
		t.Assert(s1, []string{"a", "b"})

		s1 = array1.X取切片并按范围(4, 8)
		t.Assert(len(s1), 3)
		t.Assert(s1, []string{"e", "f", "g"})
		t.Assert(array1.X取切片并按范围(10, 2), nil)

		s2 := array2.X取切片并按范围(2, 4)
		t.Assert(s2, []string{"c", "d"})

	})
}

func TestSortedStrArray_Sum(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"e", "a", "d", "c", "b", "f", "g"}
		a2 := []string{"1", "2", "3", "4", "a"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		array2 := 切片类.X创建文本排序并从切片(a2)
		t.Assert(array1.X求和(), 0)
		t.Assert(array2.X求和(), 10)
	})
}

func TestSortedStrArray_Clone(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"e", "a", "d", "c", "b", "f", "g"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		array2 := array1.X取副本()
		t.Assert(array1, array2)
		array1.X删除(1)
		t.Assert(array2.X取长度(), 7)
	})
}

func TestSortedStrArray_Clear(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"e", "a", "d", "c", "b", "f", "g"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		array1.X清空()
		t.Assert(array1.X取长度(), 0)
	})
}

func TestSortedStrArray_SubSlice(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"e", "a", "d", "c", "b", "f", "g"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		array2 := 切片类.X创建文本排序并从切片(a1, true)
		s1 := array1.X取切片并按数量(1, 3)
		t.Assert(len(s1), 3)
		t.Assert(s1, []string{"b", "c", "d"})
		t.Assert(array1.X取长度(), 7)

		s2 := array1.X取切片并按数量(1, 10)
		t.Assert(len(s2), 6)

		s3 := array1.X取切片并按数量(10, 2)
		t.Assert(len(s3), 0)

		s3 = array1.X取切片并按数量(-5, 2)
		t.Assert(s3, []string{"c", "d"})

		s3 = array1.X取切片并按数量(-10, 2)
		t.Assert(s3, nil)

		s3 = array1.X取切片并按数量(1, -2)
		t.Assert(s3, nil)

		t.Assert(array2.X取切片并按数量(1, 3), []string{"b", "c", "d"})
	})
}

func TestSortedStrArray_Len(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"e", "a", "d", "c", "b", "f", "g"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		t.Assert(array1.X取长度(), 7)

	})
}

func TestSortedStrArray_Rand(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"e", "a", "d"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		v, ok := array1.X取值随机()
		t.AssertIN(v, []string{"e", "a", "d"})
		t.Assert(ok, true)

		array2 := 切片类.X创建文本排序并从切片([]string{})
		v, ok = array2.X取值随机()
		t.Assert(v, "")
		t.Assert(ok, false)
	})
}

func TestSortedStrArray_Rands(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"e", "a", "d"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		s1 := array1.X取值随机多个(2)

		t.AssertIN(s1, []string{"e", "a", "d"})
		t.Assert(len(s1), 2)

		s1 = array1.X取值随机多个(4)
		t.Assert(len(s1), 4)

		array2 := 切片类.X创建文本排序并从切片([]string{})
		val := array2.X取值随机多个(1)
		t.Assert(val, nil)
	})
}

func TestSortedStrArray_Join(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"e", "a", "d"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		t.Assert(array1.X连接(","), `a,d,e`)
		t.Assert(array1.X连接("."), `a.d.e`)
	})

	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"a", `"b"`, `\c`}
		array1 := 切片类.X创建文本排序并从切片(a1)
		t.Assert(array1.X连接("."), `"b".\c.a`)
	})

	gtest.C(t, func(t *gtest.T) {
		array1 := 切片类.X创建文本排序并从切片([]string{})
		t.Assert(array1.X连接("."), "")
	})
}

func TestSortedStrArray_String(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"e", "a", "d"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		t.Assert(array1.String(), `["a","d","e"]`)

		array1 = nil
		t.Assert(array1.String(), "")
	})
}

func TestSortedStrArray_CountValues(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"e", "a", "d", "a", "c"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		m1 := array1.X统计()
		t.Assert(m1["a"], 2)
		t.Assert(m1["d"], 1)

	})
}

func TestSortedStrArray_Chunk(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"e", "a", "d", "a", "c"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		array2 := array1.X分割(2)
		t.Assert(len(array2), 3)
		t.Assert(len(array2[0]), 2)
		t.Assert(array2[1], []string{"c", "d"})
		t.Assert(array1.X分割(0), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"1", "2", "3", "4", "5"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		chunks := array1.X分割(3)
		t.Assert(len(chunks), 2)
		t.Assert(chunks[0], []string{"1", "2", "3"})
		t.Assert(chunks[1], []string{"4", "5"})
		t.Assert(array1.X分割(0), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"1", "2", "3", "4", "5", "6"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		chunks := array1.X分割(2)
		t.Assert(len(chunks), 3)
		t.Assert(chunks[0], []string{"1", "2"})
		t.Assert(chunks[1], []string{"3", "4"})
		t.Assert(chunks[2], []string{"5", "6"})
		t.Assert(array1.X分割(0), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"1", "2", "3", "4", "5", "6"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		chunks := array1.X分割(3)
		t.Assert(len(chunks), 2)
		t.Assert(chunks[0], []string{"1", "2", "3"})
		t.Assert(chunks[1], []string{"4", "5", "6"})
		t.Assert(array1.X分割(0), nil)
	})
}

func TestSortedStrArray_SetUnique(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"1", "1", "2", "2", "3", "3", "2", "2"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		array2 := array1.X设置去重(true)
		t.Assert(array2.X取长度(), 3)
		t.Assert(array2, []string{"1", "2", "3"})
	})
}

func TestSortedStrArray_Unique(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []string{"1", "1", "2", "2", "3", "3", "2", "2"}
		array1 := 切片类.X创建文本排序并从切片(a1)
		array1.X去重()
		t.Assert(array1.X取长度(), 3)
		t.Assert(array1, []string{"1", "2", "3"})

		array2 := 切片类.X创建文本排序并从切片([]string{})
		array2.X去重()
		t.Assert(array2.X取长度(), 0)
		t.Assert(array2, []string{})
	})
}

func TestSortedStrArray_LockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := []string{"a", "b", "c", "d"}
		a1 := 切片类.X创建文本排序并从切片(s1, true)

		ch1 := make(chan int64, 3)
		ch2 := make(chan int64, 3)
		// go1
		go a1.X遍历写锁定(func(n1 []string) { // 读写锁
			time.Sleep(2 * time.Second) // 暂停2秒
			n1[2] = "g"
			ch2 <- gconv.Int64(time.Now().UnixNano() / 1000 / 1000)
		})

		// go2
		go func() {
			time.Sleep(100 * time.Millisecond) // 故意暂停0.01秒,等go1执行锁后，再开始执行.
			ch1 <- gconv.Int64(time.Now().UnixNano() / 1000 / 1000)
			a1.X取长度()
			ch1 <- gconv.Int64(time.Now().UnixNano() / 1000 / 1000)
		}()

		t1 := <-ch1
		t2 := <-ch1
		<-ch2 // 等待go1完成

		// 防止ci抖动,以豪秒为单位
		t.AssertGT(t2-t1, 20) // go1加的读写互斥锁，所go2读的时候被阻塞。
		t.Assert(a1.X是否存在("g"), true)
	})
}

func TestSortedStrArray_RLockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := []string{"a", "b", "c", "d"}
		a1 := 切片类.X创建文本排序并从切片(s1, true)

		ch1 := make(chan int64, 3)
		ch2 := make(chan int64, 1)
		// go1
		go a1.X遍历读锁定(func(n1 []string) { // 读锁
			time.Sleep(2 * time.Second) // 暂停1秒
			n1[2] = "g"
			ch2 <- gconv.Int64(time.Now().UnixNano() / 1000 / 1000)
		})

		// go2
		go func() {
			time.Sleep(100 * time.Millisecond) // 故意暂停0.01秒,等go1执行锁后，再开始执行.
			ch1 <- gconv.Int64(time.Now().UnixNano() / 1000 / 1000)
			a1.X取长度()
			ch1 <- gconv.Int64(time.Now().UnixNano() / 1000 / 1000)
		}()

		t1 := <-ch1
		t2 := <-ch1
		<-ch2 // 等待go1完成

		// 防止ci抖动,以豪秒为单位
		t.AssertLT(t2-t1, 20) // go1加的读锁，所go2读的时候，并没有阻塞。
		t.Assert(a1.X是否存在("g"), true)
	})
}

func TestSortedStrArray_Merge(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		func1 := func(v1, v2 interface{}) int {
			if gconv.Int(v1) < gconv.Int(v2) {
				return 0
			}
			return 1
		}

		s1 := []string{"a", "b", "c", "d"}
		s2 := []string{"e", "f"}
		i1 := 切片类.X创建整数并从切片([]int{1, 2, 3})
		i2 := 切片类.X创建并从切片([]interface{}{3})
		s3 := 切片类.X创建文本并从切片([]string{"g", "h"})
		s4 := 切片类.X创建排序并从切片([]interface{}{4, 5}, func1)
		s5 := 切片类.X创建文本排序并从切片(s2)
		s6 := 切片类.X创建整数排序并从切片([]int{1, 2, 3})
		a1 := 切片类.X创建文本排序并从切片(s1)

		t.Assert(a1.X合并(s2).X取长度(), 6)
		t.Assert(a1.X合并(i1).X取长度(), 9)
		t.Assert(a1.X合并(i2).X取长度(), 10)
		t.Assert(a1.X合并(s3).X取长度(), 12)
		t.Assert(a1.X合并(s4).X取长度(), 14)
		t.Assert(a1.X合并(s5).X取长度(), 16)
		t.Assert(a1.X合并(s6).X取长度(), 19)
	})
}

func TestSortedStrArray_Json(t *testing.T) {
	// array pointer
	gtest.C(t, func(t *gtest.T) {
		s1 := []string{"a", "b", "d", "c"}
		s2 := []string{"a", "b", "c", "d"}
		a1 := 切片类.X创建文本排序并从切片(s1)
		b1, err1 := json.Marshal(a1)
		b2, err2 := json.Marshal(s1)
		t.Assert(b1, b2)
		t.Assert(err1, err2)

		a2 := 切片类.X创建文本排序()
		err1 = json.UnmarshalUseNumber(b2, &a2)
		t.AssertNil(err1)
		t.Assert(a2.X取切片(), s2)
		t.Assert(a2.Interfaces(), s2)

		var a3 切片类.SortedStrArray
		err := json.UnmarshalUseNumber(b2, &a3)
		t.AssertNil(err)
		t.Assert(a3.X取切片(), s1)
		t.Assert(a3.Interfaces(), s1)
	})
	// array value
	gtest.C(t, func(t *gtest.T) {
		s1 := []string{"a", "b", "d", "c"}
		s2 := []string{"a", "b", "c", "d"}
		a1 := *切片类.X创建文本排序并从切片(s1)
		b1, err1 := json.Marshal(a1)
		b2, err2 := json.Marshal(s1)
		t.Assert(b1, b2)
		t.Assert(err1, err2)

		a2 := 切片类.X创建文本排序()
		err1 = json.UnmarshalUseNumber(b2, &a2)
		t.AssertNil(err1)
		t.Assert(a2.X取切片(), s2)
		t.Assert(a2.Interfaces(), s2)

		var a3 切片类.SortedStrArray
		err := json.UnmarshalUseNumber(b2, &a3)
		t.AssertNil(err)
		t.Assert(a3.X取切片(), s1)
		t.Assert(a3.Interfaces(), s1)
	})
	// array pointer
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Name   string
			Scores *切片类.SortedStrArray
		}
		data := g.Map{
			"Name":   "john",
			"Scores": []string{"A+", "A", "A"},
		}
		b, err := json.Marshal(data)
		t.AssertNil(err)

		user := new(User)
		err = json.UnmarshalUseNumber(b, user)
		t.AssertNil(err)
		t.Assert(user.Name, data["Name"])
		t.Assert(user.Scores, []string{"A", "A", "A+"})
	})
	// array value
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Name   string
			Scores 切片类.SortedStrArray
		}
		data := g.Map{
			"Name":   "john",
			"Scores": []string{"A+", "A", "A"},
		}
		b, err := json.Marshal(data)
		t.AssertNil(err)

		user := new(User)
		err = json.UnmarshalUseNumber(b, user)
		t.AssertNil(err)
		t.Assert(user.Name, data["Name"])
		t.Assert(user.Scores, []string{"A", "A", "A+"})
	})
}

func TestSortedStrArray_Iterator(t *testing.T) {
	slice := g.SliceStr{"a", "b", "d", "c"}
	array := 切片类.X创建文本排序并从切片(slice)
	gtest.C(t, func(t *gtest.T) {
		array.X遍历(func(k int, v string) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	gtest.C(t, func(t *gtest.T) {
		array.X遍历升序(func(k int, v string) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	gtest.C(t, func(t *gtest.T) {
		array.X遍历降序(func(k int, v string) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	gtest.C(t, func(t *gtest.T) {
		index := 0
		array.X遍历(func(k int, v string) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
	gtest.C(t, func(t *gtest.T) {
		index := 0
		array.X遍历升序(func(k int, v string) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
	gtest.C(t, func(t *gtest.T) {
		index := 0
		array.X遍历降序(func(k int, v string) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
}

func TestSortedStrArray_RemoveValue(t *testing.T) {
	slice := g.SliceStr{"a", "b", "d", "c"}
	array := 切片类.X创建文本排序并从切片(slice)
	gtest.C(t, func(t *gtest.T) {
		t.Assert(array.X删除值("e"), false)
		t.Assert(array.X删除值("b"), true)
		t.Assert(array.X删除值("a"), true)
		t.Assert(array.X删除值("c"), true)
		t.Assert(array.X删除值("f"), false)
	})
}

func TestSortedStrArray_RemoveValues(t *testing.T) {
	slice := g.SliceStr{"a", "b", "d", "c"}
	array := 切片类.X创建文本排序并从切片(slice)
	gtest.C(t, func(t *gtest.T) {
		array.X删除多个值("a", "b", "c")
		t.Assert(array.X取切片(), g.SliceStr{"d"})
	})
}

func TestSortedStrArray_UnmarshalValue(t *testing.T) {
	type V struct {
		Name  string
		Array *切片类.SortedStrArray
	}
	// JSON
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(g.Map{
			"name":  "john",
			"array": []byte(`["1","3","2"]`),
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Array.X取切片(), g.SliceStr{"1", "2", "3"})
	})
	// Map
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(g.Map{
			"name":  "john",
			"array": g.SliceStr{"1", "3", "2"},
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Array.X取切片(), g.SliceStr{"1", "2", "3"})
	})
}
func TestSortedStrArray_Filter(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建文本排序并从切片(g.SliceStr{"", "1", "2", "0"})
		t.Assert(array.X遍历删除(func(index int, value string) bool {
			return empty.IsEmpty(value)
		}), g.SliceStr{"0", "1", "2"})
	})
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建文本排序并从切片(g.SliceStr{"1", "2"})
		t.Assert(array.X遍历删除(func(index int, value string) bool {
			return empty.IsEmpty(value)
		}), g.SliceStr{"1", "2"})
	})
}

func TestSortedStrArray_FilterEmpty(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建文本排序并从切片(g.SliceStr{"", "1", "2", "0"})
		t.Assert(array.X删除所有空值(), g.SliceStr{"0", "1", "2"})
	})
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建文本排序并从切片(g.SliceStr{"1", "2"})
		t.Assert(array.X删除所有空值(), g.SliceStr{"1", "2"})
	})
}

func TestSortedStrArray_Walk(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建文本排序并从切片(g.SliceStr{"1", "2"})
		t.Assert(array.X遍历修改(func(value string) string {
			return "key-" + value
		}), g.Slice{"key-1", "key-2"})
	})
}

func TestSortedStrArray_DeepCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建文本排序并从切片([]string{"a", "b", "c", "d"})
		copyArray := array.DeepCopy().(*切片类.SortedStrArray)
		array.X入栈右("e")
		copyArray.X入栈右("f")
		cval, _ := copyArray.X取值2(4)
		val, _ := array.X取值2(4)
		t.AssertNE(cval, val)
	})
}
