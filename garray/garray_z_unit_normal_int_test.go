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
	"github.com/gogf/gf/v2/util/gconv"
)

func Test_IntArray_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []int{0, 1, 2, 3}
		expect2 := []int{}
		array := 切片类.X创建整数并从切片(expect)
		array2 := 切片类.X创建整数并从切片(expect2)
		t.Assert(array.X取切片(), expect)
		t.Assert(array.Interfaces(), expect)
		array.X设置值(0, 100)

		v, ok := array.X取值2(0)
		t.Assert(v, 100)
		t.Assert(ok, true)

		v, ok = array.X取值2(1)
		t.Assert(v, 1)
		t.Assert(ok, true)

		t.Assert(array.X查找(100), 0)
		t.Assert(array2.X查找(100), -1)
		t.Assert(array.X是否存在(100), true)

		v, ok = array.X删除(0)
		t.Assert(v, 100)
		t.Assert(ok, true)

		v, ok = array.X删除(-1)
		t.Assert(v, 0)
		t.Assert(ok, false)

		v, ok = array.X删除(100000)
		t.Assert(v, 0)
		t.Assert(ok, false)

		t.Assert(array.X是否存在(100), false)
		array.Append别名(4)
		t.Assert(array.X取长度(), 4)
		array.X插入前面(0, 100)
		array.X插入后面(0, 200)
		t.Assert(array.X取切片(), []int{100, 200, 1, 2, 3, 4})
		array.X插入前面(5, 300)
		array.X插入后面(6, 400)
		t.Assert(array.X取切片(), []int{100, 200, 1, 2, 3, 300, 4, 400})
		t.Assert(array.X清空().X取长度(), 0)
		err := array.X插入前面(99, 300)
		t.AssertNE(err, nil)
		err = array.X插入后面(99, 400)
		t.AssertNE(err, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建整数并从切片([]int{0, 1, 2, 3})
		copyArray := array.DeepCopy().(*切片类.IntArray)
		copyArray.X设置值(0, 1)
		cval, _ := copyArray.X取值2(0)
		val, _ := array.X取值2(0)
		t.AssertNE(cval, val)
	})
}

func TestIntArray_Sort(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect1 := []int{0, 1, 2, 3}
		expect2 := []int{3, 2, 1, 0}
		array := 切片类.X创建整数()
		array2 := 切片类.X创建整数(true)
		for i := 3; i >= 0; i-- {
			array.Append别名(i)
			array2.Append别名(i)
		}
		array.X排序递增()
		t.Assert(array.X取切片(), expect1)
		array.X排序递增(true)
		t.Assert(array.X取切片(), expect2)
		t.Assert(array2.X取切片(), expect2)
	})
}

func TestIntArray_Unique(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []int{1, 2, 3, 4, 5, 3, 2, 2, 3, 5, 5}
		array := 切片类.X创建整数并从切片(expect)
		t.Assert(array.X去重().X取切片(), []int{1, 2, 3, 4, 5})
		array2 := 切片类.X创建整数并从切片([]int{})
		t.Assert(array2.X去重().X取切片(), []int{})
	})
}

func TestIntArray_PushAndPop(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []int{0, 1, 2, 3}
		array := 切片类.X创建整数并从切片(expect)
		t.Assert(array.X取切片(), expect)

		v, ok := array.X出栈左()
		t.Assert(v, 0)
		t.Assert(ok, true)

		v, ok = array.X出栈右()
		t.Assert(v, 3)
		t.Assert(ok, true)

		v, ok = array.X出栈随机()
		t.AssertIN(v, []int{1, 2})
		t.Assert(ok, true)

		v, ok = array.X出栈随机()
		t.AssertIN(v, []int{1, 2})
		t.Assert(ok, true)

		v, ok = array.X出栈随机()
		t.Assert(v, 0)
		t.Assert(ok, false)

		t.Assert(array.X取长度(), 0)
		array.X入栈左(1).X入栈右(2)
		t.Assert(array.X取切片(), []int{1, 2})
	})
}

func TestIntArray_PopLeftsAndPopRights(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建整数()

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

	gtest.C(t, func(t *gtest.T) {
		value1 := []int{0, 1, 2, 3, 4, 5, 6}
		value2 := []int{0, 1, 2, 3, 4, 5, 6}
		array1 := 切片类.X创建整数并从切片(value1)
		array2 := 切片类.X创建整数并从切片(value2)
		t.Assert(array1.X出栈左多个(2), []int{0, 1})
		t.Assert(array1.X取切片(), []int{2, 3, 4, 5, 6})
		t.Assert(array1.X出栈右多个(2), []int{5, 6})
		t.Assert(array1.X取切片(), []int{2, 3, 4})
		t.Assert(array1.X出栈右多个(20), []int{2, 3, 4})
		t.Assert(array1.X取切片(), []int{})
		t.Assert(array2.X出栈左多个(20), []int{0, 1, 2, 3, 4, 5, 6})
		t.Assert(array2.X取切片(), []int{})
	})
}

func TestIntArray_Range(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		value1 := []int{0, 1, 2, 3, 4, 5, 6}
		array1 := 切片类.X创建整数并从切片(value1)
		array2 := 切片类.X创建整数并从切片(value1, true)
		t.Assert(array1.X取切片并按范围(0, 1), []int{0})
		t.Assert(array1.X取切片并按范围(1, 2), []int{1})
		t.Assert(array1.X取切片并按范围(0, 2), []int{0, 1})
		t.Assert(array1.X取切片并按范围(10, 2), nil)
		t.Assert(array1.X取切片并按范围(-1, 10), value1)
		t.Assert(array2.X取切片并按范围(1, 2), []int{1})
	})
}

func TestIntArray_Merge(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		func1 := func(v1, v2 interface{}) int {
			if gconv.Int(v1) < gconv.Int(v2) {
				return 0
			}
			return 1
		}

		n1 := []int{0, 1, 2, 3}
		n2 := []int{4, 5, 6, 7}
		i1 := []interface{}{"1", "2"}
		s1 := []string{"a", "b", "c"}
		s2 := []string{"e", "f"}
		a1 := 切片类.X创建整数并从切片(n1)
		a2 := 切片类.X创建整数并从切片(n2)
		a3 := 切片类.X创建并从切片(i1)
		a4 := 切片类.X创建文本并从切片(s1)

		a5 := 切片类.X创建文本排序并从切片(s2)
		a6 := 切片类.X创建整数排序并从切片([]int{1, 2, 3})

		a7 := 切片类.X创建文本排序并从切片(s1)
		a8 := 切片类.X创建排序并从切片([]interface{}{4, 5}, func1)

		t.Assert(a1.X合并(a2).X取切片(), []int{0, 1, 2, 3, 4, 5, 6, 7})
		t.Assert(a1.X合并(a3).X取长度(), 10)
		t.Assert(a1.X合并(a4).X取长度(), 13)
		t.Assert(a1.X合并(a5).X取长度(), 15)
		t.Assert(a1.X合并(a6).X取长度(), 18)
		t.Assert(a1.X合并(a7).X取长度(), 21)
		t.Assert(a1.X合并(a8).X取长度(), 23)
	})
}

func TestIntArray_Fill(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{0}
		a2 := []int{0}
		array1 := 切片类.X创建整数并从切片(a1)
		array2 := 切片类.X创建整数并从切片(a2)
		t.Assert(array1.X填充(1, 2, 100), nil)
		t.Assert(array1.X取切片(), []int{0, 100, 100})

		t.Assert(array2.X填充(0, 2, 100), nil)
		t.Assert(array2.X取切片(), []int{100, 100})

		t.AssertNE(array2.X填充(-1, 2, 100), nil)
		t.Assert(array2.X取切片(), []int{100, 100})
	})
}

func TestIntArray_PopLeft(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建整数并从切片(g.SliceInt{1, 2, 3})
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

func TestIntArray_PopRight(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建整数并从切片(g.SliceInt{1, 2, 3})

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

func TestIntArray_PopLefts(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建整数并从切片(g.SliceInt{1, 2, 3})
		t.Assert(array.X出栈左多个(2), g.Slice{1, 2})
		t.Assert(array.X取长度(), 1)
		t.Assert(array.X出栈左多个(2), g.Slice{3})
		t.Assert(array.X取长度(), 0)
	})
}

func TestIntArray_PopRights(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建整数并从切片(g.SliceInt{1, 2, 3})
		t.Assert(array.X出栈右多个(2), g.Slice{2, 3})
		t.Assert(array.X取长度(), 1)
		t.Assert(array.X出栈左多个(2), g.Slice{1})
		t.Assert(array.X取长度(), 0)
	})
}

func TestIntArray_Chunk(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{1, 2, 3, 4, 5}
		array1 := 切片类.X创建整数并从切片(a1)
		chunks := array1.X分割(2)
		t.Assert(len(chunks), 3)
		t.Assert(chunks[0], []int{1, 2})
		t.Assert(chunks[1], []int{3, 4})
		t.Assert(chunks[2], []int{5})
		t.Assert(array1.X分割(0), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{1, 2, 3, 4, 5}
		array1 := 切片类.X创建整数并从切片(a1)
		chunks := array1.X分割(3)
		t.Assert(len(chunks), 2)
		t.Assert(chunks[0], []int{1, 2, 3})
		t.Assert(chunks[1], []int{4, 5})
		t.Assert(array1.X分割(0), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{1, 2, 3, 4, 5, 6}
		array1 := 切片类.X创建整数并从切片(a1)
		chunks := array1.X分割(2)
		t.Assert(len(chunks), 3)
		t.Assert(chunks[0], []int{1, 2})
		t.Assert(chunks[1], []int{3, 4})
		t.Assert(chunks[2], []int{5, 6})
		t.Assert(array1.X分割(0), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{1, 2, 3, 4, 5, 6}
		array1 := 切片类.X创建整数并从切片(a1)
		chunks := array1.X分割(3)
		t.Assert(len(chunks), 2)
		t.Assert(chunks[0], []int{1, 2, 3})
		t.Assert(chunks[1], []int{4, 5, 6})
		t.Assert(array1.X分割(0), nil)
	})
}

func TestIntArray_Pad(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{0}
		array1 := 切片类.X创建整数并从切片(a1)
		t.Assert(array1.X填满(3, 1).X取切片(), []int{0, 1, 1})
		t.Assert(array1.X填满(-4, 1).X取切片(), []int{1, 0, 1, 1})
		t.Assert(array1.X填满(3, 1).X取切片(), []int{1, 0, 1, 1})
	})
}

func TestIntArray_SubSlice(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{0, 1, 2, 3, 4, 5, 6}
		array1 := 切片类.X创建整数并从切片(a1)
		array2 := 切片类.X创建整数并从切片(a1, true)
		t.Assert(array1.X取切片并按数量(6), []int{6})
		t.Assert(array1.X取切片并按数量(5), []int{5, 6})
		t.Assert(array1.X取切片并按数量(8), nil)
		t.Assert(array1.X取切片并按数量(0, 2), []int{0, 1})
		t.Assert(array1.X取切片并按数量(2, 2), []int{2, 3})
		t.Assert(array1.X取切片并按数量(5, 8), []int{5, 6})
		t.Assert(array1.X取切片并按数量(-1, 1), []int{6})
		t.Assert(array1.X取切片并按数量(-1, 9), []int{6})
		t.Assert(array1.X取切片并按数量(-2, 3), []int{5, 6})
		t.Assert(array1.X取切片并按数量(-7, 3), []int{0, 1, 2})
		t.Assert(array1.X取切片并按数量(-8, 3), nil)
		t.Assert(array1.X取切片并按数量(-1, -3), []int{3, 4, 5})
		t.Assert(array1.X取切片并按数量(-9, 3), nil)
		t.Assert(array1.X取切片并按数量(1, -1), []int{0})
		t.Assert(array1.X取切片并按数量(1, -3), nil)
		t.Assert(array2.X取切片并按数量(0, 2), []int{0, 1})
	})
}

func TestIntArray_Rand(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{0, 1, 2, 3, 4, 5, 6}
		array1 := 切片类.X创建整数并从切片(a1)
		t.Assert(len(array1.X取值随机多个(2)), 2)
		t.Assert(len(array1.X取值随机多个(10)), 10)
		t.AssertIN(array1.X取值随机多个(1)[0], a1)

		v, ok := array1.X取值随机()
		t.AssertIN(v, a1)
		t.Assert(ok, true)

		array2 := 切片类.X创建整数并从切片([]int{})
		v, ok = array2.X取值随机()
		t.Assert(v, 0)
		t.Assert(ok, false)

		intSlices := array2.X取值随机多个(1)
		t.Assert(intSlices, nil)
	})
}

func TestIntArray_PopRands(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{100, 200, 300, 400, 500, 600}
		array := 切片类.X创建整数并从切片(a1)
		ns1 := array.X出栈随机多个(2)
		t.AssertIN(ns1, []int{100, 200, 300, 400, 500, 600})
		t.Assert(len(ns1), 2)

		ns2 := array.X出栈随机多个(7)
		t.Assert(len(ns2), 4)
		t.AssertIN(ns2, []int{100, 200, 300, 400, 500, 600})
	})
}

func TestIntArray_Shuffle(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{0, 1, 2, 3, 4, 5, 6}
		array1 := 切片类.X创建整数并从切片(a1)
		t.Assert(array1.X随机排序().X取长度(), 7)
	})
}

func TestIntArray_Reverse(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{0, 1, 2, 3, 4, 5, 6}
		array1 := 切片类.X创建整数并从切片(a1)
		t.Assert(array1.X倒排序().X取切片(), []int{6, 5, 4, 3, 2, 1, 0})
	})
}

func TestIntArray_Join(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{0, 1, 2, 3, 4, 5, 6}
		array1 := 切片类.X创建整数并从切片(a1)
		t.Assert(array1.X连接("."), "0.1.2.3.4.5.6")
	})
}

func TestIntArray_String(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{0, 1, 2, 3, 4, 5, 6}
		array1 := 切片类.X创建整数并从切片(a1)
		t.Assert(array1.String(), "[0,1,2,3,4,5,6]")
		array1 = nil
		t.Assert(array1.String(), "")
	})
}

func TestIntArray_SetArray(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{1, 2, 3, 5}
		a2 := []int{6, 7}
		array1 := 切片类.X创建整数并从切片(a1)
		array1.X设置切片(a2)
		t.Assert(array1.X取长度(), 2)
		t.Assert(array1, []int{6, 7})
	})
}

func TestIntArray_Replace(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{1, 2, 3, 5}
		a2 := []int{6, 7}
		a3 := []int{9, 10, 11, 12, 13}
		array1 := 切片类.X创建整数并从切片(a1)
		array1.X替换(a2)
		t.Assert(array1, []int{6, 7, 3, 5})

		array1.X替换(a3)
		t.Assert(array1, []int{9, 10, 11, 12})
	})
}

func TestIntArray_Clear(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{1, 2, 3, 5}
		array1 := 切片类.X创建整数并从切片(a1)
		array1.X清空()
		t.Assert(array1.X取长度(), 0)
	})
}

func TestIntArray_Clone(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{1, 2, 3, 5}
		array1 := 切片类.X创建整数并从切片(a1)
		array2 := array1.X取副本()
		t.Assert(array1, array2)
	})
}

func TestArray_Get(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{1, 2, 3, 5}
		array1 := 切片类.X创建整数并从切片(a1)
		v, ok := array1.X取值2(2)
		t.Assert(v, 3)
		t.Assert(ok, true)
		t.Assert(array1.X取长度(), 4)
	})
}

func TestIntArray_Sum(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{1, 2, 3, 5}
		array1 := 切片类.X创建整数并从切片(a1)
		t.Assert(array1.X求和(), 11)
	})
}

func TestIntArray_CountValues(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{1, 2, 3, 5, 3}
		array1 := 切片类.X创建整数并从切片(a1)
		m1 := array1.X统计()
		t.Assert(len(m1), 4)
		t.Assert(m1[1], 1)
		t.Assert(m1[3], 2)
	})
}

func TestNewIntArrayFromCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{1, 2, 3, 5, 3}
		array1 := 切片类.X创建整数并从切片复制(a1)
		t.Assert(array1.X取长度(), 5)
		t.Assert(array1, a1)
	})
}

func TestIntArray_Remove(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []int{1, 2, 3, 5, 4}
		array1 := 切片类.X创建整数并从切片(a1)
		v, ok := array1.X删除(1)
		t.Assert(v, 2)
		t.Assert(ok, true)
		t.Assert(array1.X取长度(), 4)

		v, ok = array1.X删除(0)
		t.Assert(v, 1)
		t.Assert(ok, true)
		t.Assert(array1.X取长度(), 3)

		v, ok = array1.X删除(2)
		t.Assert(v, 4)
		t.Assert(ok, true)
		t.Assert(array1.X取长度(), 2)
	})
}

func TestIntArray_LockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := []int{1, 2, 3, 4}
		a1 := 切片类.X创建整数并从切片(s1, true)

		ch1 := make(chan int64, 3)
		ch2 := make(chan int64, 3)
		// go1
		go a1.X遍历写锁定(func(n1 []int) { // 读写锁
			time.Sleep(2 * time.Second) // 暂停2秒
			n1[2] = 6
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
		t.Assert(a1.X是否存在(6), true)
	})
}

func TestIntArray_SortFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := []int{1, 4, 3, 2}
		a1 := 切片类.X创建整数并从切片(s1)
		func1 := func(v1, v2 int) bool {
			return v1 < v2
		}
		a11 := a1.X排序函数(func1)
		t.Assert(a11, []int{1, 2, 3, 4})

	})
}

func TestIntArray_RLockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := []int{1, 2, 3, 4}
		a1 := 切片类.X创建整数并从切片(s1, true)

		ch1 := make(chan int64, 3)
		ch2 := make(chan int64, 1)
		// go1
		go a1.X遍历读锁定(func(n1 []int) { // 读锁
			time.Sleep(2 * time.Second) // 暂停1秒
			n1[2] = 6
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
		t.Assert(a1.X是否存在(6), true)
	})
}

func TestIntArray_Json(t *testing.T) {
	// array pointer
	gtest.C(t, func(t *gtest.T) {
		s1 := []int{1, 4, 3, 2}
		a1 := 切片类.X创建整数并从切片(s1)
		b1, err1 := json.Marshal(a1)
		b2, err2 := json.Marshal(s1)
		t.Assert(b1, b2)
		t.Assert(err1, err2)

		a2 := 切片类.X创建整数()
		err1 = json.UnmarshalUseNumber(b2, &a2)
		t.AssertNil(err1)
		t.Assert(a2.X取切片(), s1)

		var a3 切片类.IntArray
		err := json.UnmarshalUseNumber(b2, &a3)
		t.AssertNil(err)
		t.Assert(a3.X取切片(), s1)
	})
	// array value
	gtest.C(t, func(t *gtest.T) {
		s1 := []int{1, 4, 3, 2}
		a1 := *切片类.X创建整数并从切片(s1)
		b1, err1 := json.Marshal(a1)
		b2, err2 := json.Marshal(s1)
		t.Assert(b1, b2)
		t.Assert(err1, err2)

		a2 := 切片类.X创建整数()
		err1 = json.UnmarshalUseNumber(b2, &a2)
		t.AssertNil(err1)
		t.Assert(a2.X取切片(), s1)

		var a3 切片类.IntArray
		err := json.UnmarshalUseNumber(b2, &a3)
		t.AssertNil(err)
		t.Assert(a3.X取切片(), s1)
	})
	// array pointer
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Name   string
			Scores *切片类.IntArray
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
		t.Assert(user.Scores, data["Scores"])
	})
	// array value
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Name   string
			Scores 切片类.IntArray
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
		t.Assert(user.Scores, data["Scores"])
	})
}

func TestIntArray_Iterator(t *testing.T) {
	slice := g.SliceInt{10, 20, 30, 40}
	array := 切片类.X创建整数并从切片(slice)
	gtest.C(t, func(t *gtest.T) {
		array.X遍历(func(k int, v int) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	gtest.C(t, func(t *gtest.T) {
		array.X遍历升序(func(k int, v int) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	gtest.C(t, func(t *gtest.T) {
		array.X遍历降序(func(k int, v int) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	gtest.C(t, func(t *gtest.T) {
		index := 0
		array.X遍历(func(k int, v int) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
	gtest.C(t, func(t *gtest.T) {
		index := 0
		array.X遍历升序(func(k int, v int) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
	gtest.C(t, func(t *gtest.T) {
		index := 0
		array.X遍历降序(func(k int, v int) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
}

func TestIntArray_RemoveValue(t *testing.T) {
	slice := g.SliceInt{10, 20, 30, 40}
	array := 切片类.X创建整数并从切片(slice)
	gtest.C(t, func(t *gtest.T) {
		t.Assert(array.X删除值(99), false)
		t.Assert(array.X删除值(20), true)
		t.Assert(array.X删除值(10), true)
		t.Assert(array.X删除值(20), false)
		t.Assert(array.X删除值(88), false)
		t.Assert(array.X取长度(), 2)
	})
}

func TestIntArray_RemoveValues(t *testing.T) {
	slice := g.SliceInt{10, 20, 30, 40}
	array := 切片类.X创建整数并从切片(slice)
	gtest.C(t, func(t *gtest.T) {
		array.X删除多个值(10, 20, 40)
		t.Assert(array.X取切片(), g.SliceInt{30})
	})
}

func TestIntArray_UnmarshalValue(t *testing.T) {
	type V struct {
		Name  string
		Array *切片类.IntArray
	}
	// JSON
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(g.Map{
			"name":  "john",
			"array": []byte(`[1,2,3]`),
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Array.X取切片(), g.Slice{1, 2, 3})
	})
	// Map
	// 使用gtest.C进行测试，参数t为gtest.T类型指针
	// 定义一个*v类型的指针变量v
	// 将g.Map类型的键值对（"name": "john", "array": g.Slice{1, 2, 3}）通过gconv.Struct转换到结构体变量v中
	// 断言转换过程中的错误为nil
	// 断言转换后v的Name字段值为"john"
	// 断言转换后v的Array字段转换为切片后的值等于g.Slice{1, 2, 3}
}

func TestIntArray_Filter(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建整数并从切片(g.SliceInt{0, 1, 2, 3, 4, 0})
		t.Assert(array.X遍历删除(func(index int, value int) bool {
			return empty.IsEmpty(value)
		}), g.SliceInt{1, 2, 3, 4})
	})
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建整数并从切片(g.SliceInt{1, 2, 3, 4})
		t.Assert(array.X遍历删除(func(index int, value int) bool {
			return empty.IsEmpty(value)

		}), g.SliceInt{1, 2, 3, 4})
	})
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建整数并从切片(g.SliceInt{1, 2, 3, 4})
		t.Assert(array.X遍历删除(func(index int, value int) bool {
			return value%2 == 0
		}), g.SliceInt{1, 3})
	})
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建整数并从切片(g.SliceInt{1, 2, 3, 4})
		t.Assert(array.X遍历删除(func(index int, value int) bool {
			return value%2 == 1
		}), g.SliceInt{2, 4})
	})
}

func TestIntArray_FilterEmpty(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建整数并从切片(g.SliceInt{0, 1, 2, 3, 4, 0})
		t.Assert(array.X删除所有零值(), g.SliceInt{1, 2, 3, 4})
	})
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建整数并从切片(g.SliceInt{1, 2, 3, 4})
		t.Assert(array.X删除所有零值(), g.SliceInt{1, 2, 3, 4})
	})
}

func TestIntArray_Walk(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建整数并从切片(g.SliceInt{1, 2})
		t.Assert(array.X遍历修改(func(value int) int {
			return 10 + value
		}), g.Slice{11, 12})
	})
}

func TestIntArray_NewIntArrayRange(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建整数并按范围(0, 128, 4)
		t.Assert(array.String(), `[0,4,8,12,16,20,24,28,32,36,40,44,48,52,56,60,64,68,72,76,80,84,88,92,96,100,104,108,112,116,120,124,128]`)
	})
	gtest.C(t, func(t *gtest.T) {
		array := 切片类.X创建整数并按范围(1, 128, 4)
		t.Assert(array.String(), `[1,5,9,13,17,21,25,29,33,37,41,45,49,53,57,61,65,69,73,77,81,85,89,93,97,101,105,109,113,117,121,125]`)
	})
}
