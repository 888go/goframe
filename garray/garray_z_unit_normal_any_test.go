// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// go test *.go

package 数组类_test

import (
	"testing"
	"time"
	
	"github.com/888go/goframe/garray/internal/empty"
	
	"github.com/888go/goframe/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/888go/goframe/garray/internal/json"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

func Test_Array_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []interface{}{0, 1, 2, 3}
		array := 数组类.X创建并从数组(expect)
		array2 := 数组类.X创建并从数组(expect)
		array3 := 数组类.X创建并从数组([]interface{}{})
		array4 := 数组类.X创建并按范围(1, 5, 1)

		t.Assert(array.X取切片(), expect)
		t.Assert(array.Interfaces(), expect)
		err := array.X设置值(0, 100)
		t.AssertNil(err)

		err = array.X设置值(100, 100)
		t.AssertNE(err, nil)

		t.Assert(array.X是否为空(), false)

		copyArray := array.DeepCopy()
		ca := copyArray.(*数组类.Array)
		ca.X设置值(0, 1)
		cval, _ := ca.X取值2(0)
		val, _ := array.X取值2(0)
		t.AssertNE(cval, val)

		v, ok := array.X取值2(0)
		t.Assert(v, 100)
		t.Assert(ok, true)

		v, ok = array.X取值2(1)
		t.Assert(v, 1)
		t.Assert(ok, true)

		v, ok = array.X取值2(4)
		t.Assert(v, nil)
		t.Assert(ok, false)

		t.Assert(array.X查找(100), 0)
		t.Assert(array3.X查找(100), -1)
		t.Assert(array.X是否存在(100), true)

		v, ok = array.X删除(0)
		t.Assert(v, 100)
		t.Assert(ok, true)

		v, ok = array.X删除(-1)
		t.Assert(v, nil)
		t.Assert(ok, false)

		v, ok = array.X删除(100000)
		t.Assert(v, nil)
		t.Assert(ok, false)

		v, ok = array2.X删除(3)
		t.Assert(v, 3)
		t.Assert(ok, true)

		v, ok = array2.X删除(1)
		t.Assert(v, 1)
		t.Assert(ok, true)

		t.Assert(array.X是否存在(100), false)
		array.Append别名(4)
		t.Assert(array.X取长度(), 4)
		array.X插入前面(0, 100)
		array.X插入后面(0, 200)
		t.Assert(array.X取切片(), []interface{}{100, 200, 2, 2, 3, 4})
		array.X插入前面(5, 300)
		array.X插入后面(6, 400)
		t.Assert(array.X取切片(), []interface{}{100, 200, 2, 2, 3, 300, 4, 400})
		t.Assert(array.X清空().X取长度(), 0)
		err = array.X插入前面(99, 9900)
		t.AssertNE(err, nil)
		err = array.X插入后面(99, 9900)
		t.AssertNE(err, nil)

		t.Assert(array4.String(), "[1,2,3,4,5]")
	})
}

func TestArray_Sort(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect1 := []interface{}{0, 1, 2, 3}
		expect2 := []interface{}{3, 2, 1, 0}
		array := 数组类.NewArray别名()
		for i := 3; i >= 0; i-- {
			array.Append别名(i)
		}
		array.X排序并带函数(func(v1, v2 interface{}) bool {
			return v1.(int) < v2.(int)
		})
		t.Assert(array.X取切片(), expect1)
		array.X排序并带函数(func(v1, v2 interface{}) bool {
			return v1.(int) > v2.(int)
		})
		t.Assert(array.X取切片(), expect2)
	})
}

func TestArray_Unique(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []interface{}{1, 2, 3, 4, 5, 3, 2, 2, 3, 5, 5}
		array := 数组类.X创建并从数组(expect)
		t.Assert(array.X去重().X取切片(), []interface{}{1, 2, 3, 4, 5})
	})
	gtest.C(t, func(t *gtest.T) {
		expect := []interface{}{}
		array := 数组类.X创建并从数组(expect)
		t.Assert(array.X去重().X取切片(), []interface{}{})
	})
}

func TestArray_PushAndPop(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := []interface{}{0, 1, 2, 3}
		array := 数组类.X创建并从数组(expect)
		t.Assert(array.X取切片(), expect)

		v, ok := array.X出栈左()
		t.Assert(v, 0)
		t.Assert(ok, true)

		v, ok = array.X出栈右()
		t.Assert(v, 3)
		t.Assert(ok, true)

		v, ok = array.X出栈随机()
		t.AssertIN(v, []interface{}{1, 2})
		t.Assert(ok, true)

		v, ok = array.X出栈随机()
		t.AssertIN(v, []interface{}{1, 2})
		t.Assert(ok, true)

		t.Assert(array.X取长度(), 0)
		array.X入栈左(1).X入栈右(2)
		t.Assert(array.X取切片(), []interface{}{1, 2})
	})
}

func TestArray_PopRands(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{100, 200, 300, 400, 500, 600}
		array := 数组类.NewFromCopy别名(a1)
		t.AssertIN(array.X出栈随机多个(2), []interface{}{100, 200, 300, 400, 500, 600})
	})
}

func TestArray_PopLeft(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 数组类.NewFrom别名(g.Slice{1, 2, 3})
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

func TestArray_PopRight(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 数组类.NewFrom别名(g.Slice{1, 2, 3})

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

func TestArray_PopLefts(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 数组类.NewFrom别名(g.Slice{1, 2, 3})
		t.Assert(array.X出栈左多个(2), g.Slice{1, 2})
		t.Assert(array.X取长度(), 1)
		t.Assert(array.X出栈左多个(2), g.Slice{3})
		t.Assert(array.X取长度(), 0)
	})
}

func TestArray_PopRights(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 数组类.NewFrom别名(g.Slice{1, 2, 3})
		t.Assert(array.X出栈右多个(2), g.Slice{2, 3})
		t.Assert(array.X取长度(), 1)
		t.Assert(array.X出栈左多个(2), g.Slice{1})
		t.Assert(array.X取长度(), 0)
	})
}

func TestArray_PopLeftsAndPopRights(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 数组类.X创建()
		v, ok := array.X出栈左()
		t.Assert(v, nil)
		t.Assert(ok, false)
		t.Assert(array.X出栈左多个(10), nil)

		v, ok = array.X出栈右()
		t.Assert(v, nil)
		t.Assert(ok, false)
		t.Assert(array.X出栈右多个(10), nil)

		v, ok = array.X出栈随机()
		t.Assert(v, nil)
		t.Assert(ok, false)
		t.Assert(array.X出栈随机多个(10), nil)
	})

	gtest.C(t, func(t *gtest.T) {
		value1 := []interface{}{0, 1, 2, 3, 4, 5, 6}
		value2 := []interface{}{0, 1, 2, 3, 4, 5, 6}
		array1 := 数组类.X创建并从数组(value1)
		array2 := 数组类.X创建并从数组(value2)
		t.Assert(array1.X出栈左多个(2), []interface{}{0, 1})
		t.Assert(array1.X取切片(), []interface{}{2, 3, 4, 5, 6})
		t.Assert(array1.X出栈右多个(2), []interface{}{5, 6})
		t.Assert(array1.X取切片(), []interface{}{2, 3, 4})
		t.Assert(array1.X出栈右多个(20), []interface{}{2, 3, 4})
		t.Assert(array1.X取切片(), []interface{}{})
		t.Assert(array2.X出栈左多个(20), []interface{}{0, 1, 2, 3, 4, 5, 6})
		t.Assert(array2.X取切片(), []interface{}{})
	})
}

func TestArray_Range(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		value1 := []interface{}{0, 1, 2, 3, 4, 5, 6}
		array1 := 数组类.X创建并从数组(value1)
		array2 := 数组类.X创建并从数组(value1, true)
		t.Assert(array1.X取切片并按范围(0, 1), []interface{}{0})
		t.Assert(array1.X取切片并按范围(1, 2), []interface{}{1})
		t.Assert(array1.X取切片并按范围(0, 2), []interface{}{0, 1})
		t.Assert(array1.X取切片并按范围(-1, 10), value1)
		t.Assert(array1.X取切片并按范围(10, 2), nil)
		t.Assert(array2.X取切片并按范围(1, 3), []interface{}{1, 2})
	})
}

func TestArray_Merge(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		func1 := func(v1, v2 interface{}) int {
			if gconv.Int(v1) < gconv.Int(v2) {
				return 0
			}
			return 1
		}

		i1 := []interface{}{0, 1, 2, 3}
		i2 := []interface{}{4, 5, 6, 7}
		array1 := 数组类.X创建并从数组(i1)
		array2 := 数组类.X创建并从数组(i2)
		t.Assert(array1.X合并(array2).X取切片(), []interface{}{0, 1, 2, 3, 4, 5, 6, 7})

		// 声明并初始化一个字符串切片s1，其中包含了四个元素："a", "b", "c", "d"
// s1 := []string{"a", "b", "c", "d"}
		s2 := []string{"e", "f"}
		i3 := 数组类.X创建整数并从数组([]int{1, 2, 3})
		i4 := 数组类.X创建并从数组([]interface{}{3})
		s3 := 数组类.X创建文本并从数组([]string{"g", "h"})
		s4 := 数组类.X创建排序并从数组([]interface{}{4, 5}, func1)
		s5 := 数组类.X创建文本排序并从数组(s2)
		s6 := 数组类.X创建整数排序并从数组([]int{1, 2, 3})
		a1 := 数组类.X创建并从数组(i1)

		t.Assert(a1.X合并(s2).X取长度(), 6)
		t.Assert(a1.X合并(i3).X取长度(), 9)
		t.Assert(a1.X合并(i4).X取长度(), 10)
		t.Assert(a1.X合并(s3).X取长度(), 12)
		t.Assert(a1.X合并(s4).X取长度(), 14)
		t.Assert(a1.X合并(s5).X取长度(), 16)
		t.Assert(a1.X合并(s6).X取长度(), 19)
	})
}

func TestArray_Fill(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{0}
		a2 := []interface{}{0}
		array1 := 数组类.X创建并从数组(a1)
		array2 := 数组类.X创建并从数组(a2, true)

		t.Assert(array1.X填充(1, 2, 100), nil)
		t.Assert(array1.X取切片(), []interface{}{0, 100, 100})

		t.Assert(array2.X填充(0, 2, 100), nil)
		t.Assert(array2.X取切片(), []interface{}{100, 100})

		t.AssertNE(array2.X填充(-1, 2, 100), nil)
		t.Assert(array2.X取切片(), []interface{}{100, 100})
	})
}

func TestArray_Chunk(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{1, 2, 3, 4, 5}
		array1 := 数组类.X创建并从数组(a1)
		chunks := array1.X分割(2)
		t.Assert(len(chunks), 3)
		t.Assert(chunks[0], []interface{}{1, 2})
		t.Assert(chunks[1], []interface{}{3, 4})
		t.Assert(chunks[2], []interface{}{5})
		t.Assert(array1.X分割(0), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{1, 2, 3, 4, 5}
		array1 := 数组类.X创建并从数组(a1)
		chunks := array1.X分割(3)
		t.Assert(len(chunks), 2)
		t.Assert(chunks[0], []interface{}{1, 2, 3})
		t.Assert(chunks[1], []interface{}{4, 5})
		t.Assert(array1.X分割(0), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{1, 2, 3, 4, 5, 6}
		array1 := 数组类.X创建并从数组(a1)
		chunks := array1.X分割(2)
		t.Assert(len(chunks), 3)
		t.Assert(chunks[0], []interface{}{1, 2})
		t.Assert(chunks[1], []interface{}{3, 4})
		t.Assert(chunks[2], []interface{}{5, 6})
		t.Assert(array1.X分割(0), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{1, 2, 3, 4, 5, 6}
		array1 := 数组类.X创建并从数组(a1)
		chunks := array1.X分割(3)
		t.Assert(len(chunks), 2)
		t.Assert(chunks[0], []interface{}{1, 2, 3})
		t.Assert(chunks[1], []interface{}{4, 5, 6})
		t.Assert(array1.X分割(0), nil)
	})
}

func TestArray_Pad(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{0}
		array1 := 数组类.X创建并从数组(a1)
		t.Assert(array1.X填满(3, 1).X取切片(), []interface{}{0, 1, 1})
		t.Assert(array1.X填满(-4, 1).X取切片(), []interface{}{1, 0, 1, 1})
		t.Assert(array1.X填满(3, 1).X取切片(), []interface{}{1, 0, 1, 1})
	})
}

func TestArray_SubSlice(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{0, 1, 2, 3, 4, 5, 6}
		array1 := 数组类.X创建并从数组(a1)
		array2 := 数组类.X创建并从数组(a1, true)
		t.Assert(array1.X取切片并按数量(0, 2), []interface{}{0, 1})
		t.Assert(array1.X取切片并按数量(2, 2), []interface{}{2, 3})
		t.Assert(array1.X取切片并按数量(5, 8), []interface{}{5, 6})
		t.Assert(array1.X取切片并按数量(9, 1), nil)
		t.Assert(array1.X取切片并按数量(-2, 2), []interface{}{5, 6})
		t.Assert(array1.X取切片并按数量(-9, 2), nil)
		t.Assert(array1.X取切片并按数量(1, -2), nil)
		t.Assert(array2.X取切片并按数量(0, 2), []interface{}{0, 1})
	})
}

func TestArray_Rand(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{0, 1, 2, 3, 4, 5, 6}
		array1 := 数组类.X创建并从数组(a1)
		t.Assert(len(array1.X取值随机多个(2)), 2)
		t.Assert(len(array1.X取值随机多个(10)), 10)
		t.AssertIN(array1.X取值随机多个(1)[0], a1)
	})

	gtest.C(t, func(t *gtest.T) {
		s1 := []interface{}{"a", "b", "c", "d"}
		a1 := 数组类.X创建并从数组(s1)
		i1, ok := a1.X取值随机()
		t.Assert(ok, true)
		t.Assert(a1.X是否存在(i1), true)
		t.Assert(a1.X取长度(), 4)
	})

	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{}
		array1 := 数组类.X创建并从数组(a1)
		rand, found := array1.X取值随机()
		t.AssertNil(rand)
		t.Assert(found, false)
	})

	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{}
		array1 := 数组类.X创建并从数组(a1)
		rand := array1.X取值随机多个(1)
		t.AssertNil(rand)
	})
}

func TestArray_Shuffle(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{0, 1, 2, 3, 4, 5, 6}
		array1 := 数组类.X创建并从数组(a1)
		t.Assert(array1.X随机排序().X取长度(), 7)
	})
}

func TestArray_Reverse(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{0, 1, 2, 3, 4, 5, 6}
		array1 := 数组类.X创建并从数组(a1)
		t.Assert(array1.X倒排序().X取切片(), []interface{}{6, 5, 4, 3, 2, 1, 0})
	})
}

func TestArray_Join(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{0, 1, 2, 3, 4, 5, 6}
		array1 := 数组类.X创建并从数组(a1)
		t.Assert(array1.X连接("."), `0.1.2.3.4.5.6`)
	})

	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{0, 1, `"a"`, `\a`}
		array1 := 数组类.X创建并从数组(a1)
		t.Assert(array1.X连接("."), `0.1."a".\a`)
	})

	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{}
		array1 := 数组类.X创建并从数组(a1)
		t.Assert(len(array1.X连接(".")), 0)
	})
}

func TestArray_String(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{0, 1, 2, 3, 4, 5, 6}
		array1 := 数组类.X创建并从数组(a1)
		t.Assert(array1.String(), `[0,1,2,3,4,5,6]`)
		array1 = nil
		t.Assert(array1.String(), "")
	})
}

func TestArray_Replace(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{0, 1, 2, 3, 4, 5, 6}
		a2 := []interface{}{"a", "b", "c"}
		a3 := []interface{}{"m", "n", "p", "z", "x", "y", "d", "u"}
		array1 := 数组类.X创建并从数组(a1)
		array2 := array1.X替换(a2)
		t.Assert(array2.X取长度(), 7)
		t.Assert(array2.X是否存在("b"), true)
		t.Assert(array2.X是否存在(4), true)
		t.Assert(array2.X是否存在("v"), false)
		array3 := array1.X替换(a3)
		t.Assert(array3.X取长度(), 7)
		t.Assert(array3.X是否存在(4), false)
		t.Assert(array3.X是否存在("p"), true)
		t.Assert(array3.X是否存在("u"), false)
	})
}

func TestArray_SetArray(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{0, 1, 2, 3, 4, 5, 6}
		a2 := []interface{}{"a", "b", "c"}

		array1 := 数组类.X创建并从数组(a1)
		array1 = array1.X设置数组(a2)
		t.Assert(array1.X取长度(), 3)
		t.Assert(array1.X是否存在("b"), true)
		t.Assert(array1.X是否存在("5"), false)
	})
}

func TestArray_Sum(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{0, 1, 2, 3}
		a2 := []interface{}{"a", "b", "c"}
		a3 := []interface{}{"a", "1", "2"}

		array1 := 数组类.X创建并从数组(a1)
		array2 := 数组类.X创建并从数组(a2)
		array3 := 数组类.X创建并从数组(a3)

		t.Assert(array1.X求和(), 6)
		t.Assert(array2.X求和(), 0)
		t.Assert(array3.X求和(), 3)

	})
}

func TestArray_Clone(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{0, 1, 2, 3}
		array1 := 数组类.X创建并从数组(a1)
		array2 := array1.X取副本()

		t.Assert(array1.X取长度(), 4)
		t.Assert(array2.X求和(), 6)
		t.AssertEQ(array1, array2)

	})
}

func TestArray_CountValues(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a1 := []interface{}{"a", "b", "c", "d", "e", "d"}
		array1 := 数组类.X创建并从数组(a1)
		array2 := array1.X统计()
		t.Assert(len(array2), 5)
		t.Assert(array2["b"], 1)
		t.Assert(array2["d"], 2)
	})
}

func TestArray_LockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := []interface{}{"a", "b", "c", "d"}
		a1 := 数组类.X创建并从数组(s1, true)

		ch1 := make(chan int64, 3)
		ch2 := make(chan int64, 3)
		// go1
		go a1.X遍历并写锁定(func(n1 []interface{}) { // 读写锁
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

func TestArray_RLockFunc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s1 := []interface{}{"a", "b", "c", "d"}
		a1 := 数组类.X创建并从数组(s1, true)

		ch1 := make(chan int64, 3)
		ch2 := make(chan int64, 1)
		// go1
		go a1.X遍历并读锁定(func(n1 []interface{}) { // 读锁
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

func TestArray_Json(t *testing.T) {
	// pointer
	gtest.C(t, func(t *gtest.T) {
		s1 := []interface{}{"a", "b", "d", "c"}
		a1 := 数组类.X创建并从数组(s1)
		b1, err1 := json.Marshal(a1)
		b2, err2 := json.Marshal(s1)
		t.Assert(b1, b2)
		t.Assert(err1, err2)

		a2 := 数组类.X创建()
		err2 = json.UnmarshalUseNumber(b2, &a2)
		t.Assert(err2, nil)
		t.Assert(a2.X取切片(), s1)

		var a3 数组类.Array
		err := json.UnmarshalUseNumber(b2, &a3)
		t.AssertNil(err)
		t.Assert(a3.X取切片(), s1)
	})
	// value.
	gtest.C(t, func(t *gtest.T) {
		s1 := []interface{}{"a", "b", "d", "c"}
		a1 := *数组类.X创建并从数组(s1)
		b1, err1 := json.Marshal(a1)
		b2, err2 := json.Marshal(s1)
		t.Assert(b1, b2)
		t.Assert(err1, err2)

		a2 := 数组类.X创建()
		err2 = json.UnmarshalUseNumber(b2, &a2)
		t.Assert(err2, nil)
		t.Assert(a2.X取切片(), s1)

		var a3 数组类.Array
		err := json.UnmarshalUseNumber(b2, &a3)
		t.AssertNil(err)
		t.Assert(a3.X取切片(), s1)
	})
	// pointer
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Name   string
			Scores *数组类.Array
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
	// value
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Name   string
			Scores 数组类.Array
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

func TestArray_Iterator(t *testing.T) {
	slice := g.Slice{"a", "b", "d", "c"}
	array := 数组类.X创建并从数组(slice)
	gtest.C(t, func(t *gtest.T) {
		array.X遍历(func(k int, v interface{}) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	gtest.C(t, func(t *gtest.T) {
		array.X遍历升序(func(k int, v interface{}) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	gtest.C(t, func(t *gtest.T) {
		array.X遍历降序(func(k int, v interface{}) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	gtest.C(t, func(t *gtest.T) {
		index := 0
		array.X遍历(func(k int, v interface{}) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
	gtest.C(t, func(t *gtest.T) {
		index := 0
		array.X遍历升序(func(k int, v interface{}) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
	gtest.C(t, func(t *gtest.T) {
		index := 0
		array.X遍历降序(func(k int, v interface{}) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
}

func TestArray_RemoveValue(t *testing.T) {
	slice := g.Slice{"a", "b", "d", "c"}
	array := 数组类.X创建并从数组(slice)
	gtest.C(t, func(t *gtest.T) {
		t.Assert(array.X删除值("e"), false)
		t.Assert(array.X删除值("b"), true)
		t.Assert(array.X删除值("a"), true)
		t.Assert(array.X删除值("c"), true)
		t.Assert(array.X删除值("f"), false)
	})
}

func TestArray_RemoveValues(t *testing.T) {
	slice := g.Slice{"a", "b", "d", "c"}
	array := 数组类.X创建并从数组(slice)
	gtest.C(t, func(t *gtest.T) {
		array.X删除多个值("a", "b", "c")
		t.Assert(array.X取切片(), g.Slice{"d"})
	})
}

func TestArray_UnmarshalValue(t *testing.T) {
	type V struct {
		Name  string
		Array *数组类.Array
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
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(g.Map{
			"name":  "john",
			"array": g.Slice{1, 2, 3},
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Array.X取切片(), g.Slice{1, 2, 3})
	})
}

func TestArray_FilterNil(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		values := g.Slice{0, 1, 2, 3, 4, "", g.Slice{}}
		array := 数组类.X创建并从数组复制(values)
		t.Assert(array.X删除所有nil().X取切片(), values)
	})
	gtest.C(t, func(t *gtest.T) {
		array := 数组类.X创建并从数组复制(g.Slice{nil, 1, 2, 3, 4, nil})
		t.Assert(array.X删除所有nil(), g.Slice{1, 2, 3, 4})
	})
}

func TestArray_Filter(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		values := g.Slice{0, 1, 2, 3, 4, "", g.Slice{}}
		array := 数组类.X创建并从数组复制(values)
		t.Assert(array.X遍历删除(func(index int, value interface{}) bool {
			return empty.IsNil(value)
		}).X取切片(), values)
	})
	gtest.C(t, func(t *gtest.T) {
		array := 数组类.X创建并从数组复制(g.Slice{nil, 1, 2, 3, 4, nil})
		t.Assert(array.X遍历删除(func(index int, value interface{}) bool {
			return empty.IsNil(value)
		}), g.Slice{1, 2, 3, 4})
	})
	gtest.C(t, func(t *gtest.T) {
		array := 数组类.X创建并从数组(g.Slice{0, 1, 2, 3, 4, "", g.Slice{}})

		t.Assert(array.X遍历删除(func(index int, value interface{}) bool {
			return empty.IsEmpty(value)
		}), g.Slice{1, 2, 3, 4})
	})
	gtest.C(t, func(t *gtest.T) {
		array := 数组类.X创建并从数组(g.Slice{1, 2, 3, 4})

		t.Assert(array.X遍历删除(func(index int, value interface{}) bool {
			return empty.IsEmpty(value)
		}), g.Slice{1, 2, 3, 4})
	})
}

func TestArray_FilterEmpty(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 数组类.X创建并从数组(g.Slice{0, 1, 2, 3, 4, "", g.Slice{}})
		t.Assert(array.X删除所有空值(), g.Slice{1, 2, 3, 4})
	})
	gtest.C(t, func(t *gtest.T) {
		array := 数组类.X创建并从数组(g.Slice{1, 2, 3, 4})
		t.Assert(array.X删除所有空值(), g.Slice{1, 2, 3, 4})
	})
}

func TestArray_Walk(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := 数组类.X创建并从数组(g.Slice{"1", "2"})
		t.Assert(array.X遍历修改(func(value interface{}) interface{} {
			return "key-" + gconv.String(value)
		}), g.Slice{"key-1", "key-2"})
	})
}
