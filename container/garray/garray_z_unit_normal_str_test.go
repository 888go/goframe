// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// go test *.go

package 数组类_test

import (
	"strings"
	"testing"
	"time"
	
	"github.com/888go/goframe/internal/empty"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_StrArray_Basic(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		expect := []string{"0", "1", "2", "3"}
		array := 数组类.X创建文本并从数组(expect)
		array2 := 数组类.X创建文本并从数组(expect, true)
		array3 := 数组类.X创建文本并从数组([]string{})
		t.Assert(array.X取切片(), expect)
		t.Assert(array.X取any数组(), expect)
		array.X设置值(0, "100")

		v, ok := array.X取值2(0)
		t.Assert(v, 100)
		t.Assert(ok, true)

		v, ok = array3.X取值2(0)
		t.Assert(v, "")
		t.Assert(ok, false)

		t.Assert(array.X查找("100"), 0)
		t.Assert(array.X是否存在("100"), true)

		v, ok = array.X删除(0)
		t.Assert(v, 100)
		t.Assert(ok, true)

		v, ok = array.X删除(-1)
		t.Assert(v, "")
		t.Assert(ok, false)

		v, ok = array.X删除(100000)
		t.Assert(v, "")
		t.Assert(ok, false)

		t.Assert(array.X是否存在("100"), false)
		array.Append别名("4")
		t.Assert(array.X取长度(), 4)
		array.X插入前面(0, "100")
		array.X插入后面(0, "200")
		t.Assert(array.X取切片(), []string{"100", "200", "1", "2", "3", "4"})
		array.X插入前面(5, "300")
		array.X插入后面(6, "400")
		t.Assert(array.X取切片(), []string{"100", "200", "1", "2", "3", "300", "4", "400"})
		t.Assert(array.X清空().X取长度(), 0)
		t.Assert(array2.X取切片(), expect)
		t.Assert(array3.X查找("100"), -1)
		err := array.X插入前面(99, "300")
		t.AssertNE(err, nil)
		array.X插入后面(99, "400")
		t.AssertNE(err, nil)

	})

	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建文本并从数组([]string{"0", "1", "2", "3"})

		copyArray := array.DeepCopy().(*数组类.StrArray)
		copyArray.X设置值(0, "1")
		cval, _ := copyArray.X取值2(0)
		val, _ := array.X取值2(0)
		t.AssertNE(cval, val)
	})
}

func TestStrArray_ContainsI(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := 数组类.X创建文本()
		t.Assert(s.X是否存在("A"), false)
		s.Append别名("a", "b", "C")
		t.Assert(s.X是否存在("A"), false)
		t.Assert(s.X是否存在("a"), true)
		t.Assert(s.X是否存在并忽略大小写("A"), true)
	})
}

func TestStrArray_Sort(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		expect1 := []string{"0", "1", "2", "3"}
		expect2 := []string{"3", "2", "1", "0"}
		array := 数组类.X创建文本()
		for i := 3; i >= 0; i-- {
			array.Append别名(转换类.String(i))
		}
		array.X排序递增()
		t.Assert(array.X取切片(), expect1)
		array.X排序递增(true)
		t.Assert(array.X取切片(), expect2)
	})
}

func TestStrArray_Unique(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		expect := []string{"1", "1", "2", "2", "3", "3", "2", "2"}
		array := 数组类.X创建文本并从数组(expect)
		t.Assert(array.X去重().X取切片(), []string{"1", "2", "3"})
		array1 := 数组类.X创建文本并从数组([]string{})
		t.Assert(array1.X去重().X取切片(), []string{})
	})
}

func TestStrArray_PushAndPop(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		expect := []string{"0", "1", "2", "3"}
		array := 数组类.X创建文本并从数组(expect)
		t.Assert(array.X取切片(), expect)

		v, ok := array.X出栈左()
		t.Assert(v, "0")
		t.Assert(ok, true)

		v, ok = array.X出栈右()
		t.Assert(v, "3")
		t.Assert(ok, true)

		v, ok = array.X出栈随机()
		t.AssertIN(v, []string{"1", "2"})
		t.Assert(ok, true)

		v, ok = array.X出栈随机()
		t.AssertIN(v, []string{"1", "2"})
		t.Assert(ok, true)

		v, ok = array.X出栈随机()
		t.Assert(v, "")
		t.Assert(ok, false)

		t.Assert(array.X取长度(), 0)
		array.X入栈左("1").X入栈右("2")
		t.Assert(array.X取切片(), []string{"1", "2"})
	})
}

func TestStrArray_PopLeft(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建文本并从数组(g.SliceStr别名{"1", "2", "3"})
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

func TestStrArray_PopRight(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建文本并从数组(g.SliceStr别名{"1", "2", "3"})

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

func TestStrArray_PopLefts(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建文本并从数组(g.SliceStr别名{"1", "2", "3"})
		t.Assert(array.X出栈左多个(2), g.Slice别名{"1", "2"})
		t.Assert(array.X取长度(), 1)
		t.Assert(array.X出栈左多个(2), g.Slice别名{"3"})
		t.Assert(array.X取长度(), 0)
	})
}

func TestStrArray_PopRights(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建文本并从数组(g.SliceStr别名{"1", "2", "3"})
		t.Assert(array.X出栈右多个(2), g.Slice别名{"2", "3"})
		t.Assert(array.X取长度(), 1)
		t.Assert(array.X出栈左多个(2), g.Slice别名{"1"})
		t.Assert(array.X取长度(), 0)
	})
}

func TestStrArray_PopLeftsAndPopRights(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建文本()
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

	单元测试类.C(t, func(t *单元测试类.T) {
		value1 := []string{"0", "1", "2", "3", "4", "5", "6"}
		value2 := []string{"0", "1", "2", "3", "4", "5", "6"}
		array1 := 数组类.X创建文本并从数组(value1)
		array2 := 数组类.X创建文本并从数组(value2)
		t.Assert(array1.X出栈左多个(2), []interface{}{"0", "1"})
		t.Assert(array1.X取切片(), []interface{}{"2", "3", "4", "5", "6"})
		t.Assert(array1.X出栈右多个(2), []interface{}{"5", "6"})
		t.Assert(array1.X取切片(), []interface{}{"2", "3", "4"})
		t.Assert(array1.X出栈右多个(20), []interface{}{"2", "3", "4"})
		t.Assert(array1.X取切片(), []interface{}{})
		t.Assert(array2.X出栈左多个(20), []interface{}{"0", "1", "2", "3", "4", "5", "6"})
		t.Assert(array2.X取切片(), []interface{}{})
	})
}

func TestString_Range(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		value1 := []string{"0", "1", "2", "3", "4", "5", "6"}
		array1 := 数组类.X创建文本并从数组(value1)
		array2 := 数组类.X创建文本并从数组(value1, true)
		t.Assert(array1.X取切片并按范围(0, 1), []interface{}{"0"})
		t.Assert(array1.X取切片并按范围(1, 2), []interface{}{"1"})
		t.Assert(array1.X取切片并按范围(0, 2), []interface{}{"0", "1"})
		t.Assert(array1.X取切片并按范围(-1, 10), value1)
		t.Assert(array1.X取切片并按范围(10, 1), nil)
		t.Assert(array2.X取切片并按范围(0, 1), []interface{}{"0"})
	})
}

func TestStrArray_Merge(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a11 := []string{"0", "1", "2", "3"}
		a21 := []string{"4", "5", "6", "7"}
		array1 := 数组类.X创建文本并从数组(a11)
		array2 := 数组类.X创建文本并从数组(a21)
		t.Assert(array1.X合并(array2).X取切片(), []string{"0", "1", "2", "3", "4", "5", "6", "7"})

		func1 := func(v1, v2 interface{}) int {
			if 转换类.X取整数(v1) < 转换类.X取整数(v2) {
				return 0
			}
			return 1
		}

		s1 := []string{"a", "b", "c", "d"}
		s2 := []string{"e", "f"}
		i1 := 数组类.X创建整数并从数组([]int{1, 2, 3})
		i2 := 数组类.X创建并从数组([]interface{}{3})
		s3 := 数组类.X创建文本并从数组([]string{"g", "h"})
		s4 := 数组类.X创建排序并从数组([]interface{}{4, 5}, func1)
		s5 := 数组类.X创建文本排序并从数组(s2)
		s6 := 数组类.X创建整数排序并从数组([]int{1, 2, 3})
		a1 := 数组类.X创建文本并从数组(s1)

		t.Assert(a1.X合并(s2).X取长度(), 6)
		t.Assert(a1.X合并(i1).X取长度(), 9)
		t.Assert(a1.X合并(i2).X取长度(), 10)
		t.Assert(a1.X合并(s3).X取长度(), 12)
		t.Assert(a1.X合并(s4).X取长度(), 14)
		t.Assert(a1.X合并(s5).X取长度(), 16)
		t.Assert(a1.X合并(s6).X取长度(), 19)
	})
}

func TestStrArray_Fill(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"0"}
		a2 := []string{"0"}
		array1 := 数组类.X创建文本并从数组(a1)
		array2 := 数组类.X创建文本并从数组(a2)
		t.Assert(array1.X填充(1, 2, "100"), nil)
		t.Assert(array1.X取切片(), []string{"0", "100", "100"})
		t.Assert(array2.X填充(0, 2, "100"), nil)
		t.Assert(array2.X取切片(), []string{"100", "100"})
		t.AssertNE(array2.X填充(-1, 2, "100"), nil)
		t.Assert(array2.X取长度(), 2)
	})
}

func TestStrArray_Chunk(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"1", "2", "3", "4", "5"}
		array1 := 数组类.X创建文本并从数组(a1)
		chunks := array1.X分割(2)
		t.Assert(len(chunks), 3)
		t.Assert(chunks[0], []string{"1", "2"})
		t.Assert(chunks[1], []string{"3", "4"})
		t.Assert(chunks[2], []string{"5"})
		t.Assert(len(array1.X分割(0)), 0)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"1", "2", "3", "4", "5"}
		array1 := 数组类.X创建文本并从数组(a1)
		chunks := array1.X分割(3)
		t.Assert(len(chunks), 2)
		t.Assert(chunks[0], []string{"1", "2", "3"})
		t.Assert(chunks[1], []string{"4", "5"})
		t.Assert(array1.X分割(0), nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"1", "2", "3", "4", "5", "6"}
		array1 := 数组类.X创建文本并从数组(a1)
		chunks := array1.X分割(2)
		t.Assert(len(chunks), 3)
		t.Assert(chunks[0], []string{"1", "2"})
		t.Assert(chunks[1], []string{"3", "4"})
		t.Assert(chunks[2], []string{"5", "6"})
		t.Assert(array1.X分割(0), nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"1", "2", "3", "4", "5", "6"}
		array1 := 数组类.X创建文本并从数组(a1)
		chunks := array1.X分割(3)
		t.Assert(len(chunks), 2)
		t.Assert(chunks[0], []string{"1", "2", "3"})
		t.Assert(chunks[1], []string{"4", "5", "6"})
		t.Assert(array1.X分割(0), nil)
	})
}

func TestStrArray_Pad(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"0"}
		array1 := 数组类.X创建文本并从数组(a1)
		t.Assert(array1.X填满(3, "1").X取切片(), []string{"0", "1", "1"})
		t.Assert(array1.X填满(-4, "1").X取切片(), []string{"1", "0", "1", "1"})
		t.Assert(array1.X填满(3, "1").X取切片(), []string{"1", "0", "1", "1"})
	})
}

func TestStrArray_SubSlice(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"0", "1", "2", "3", "4", "5", "6"}
		array1 := 数组类.X创建文本并从数组(a1)
		array2 := 数组类.X创建文本并从数组(a1, true)
		t.Assert(array1.X取切片并按数量(0, 2), []string{"0", "1"})
		t.Assert(array1.X取切片并按数量(2, 2), []string{"2", "3"})
		t.Assert(array1.X取切片并按数量(5, 8), []string{"5", "6"})
		t.Assert(array1.X取切片并按数量(8, 2), nil)
		t.Assert(array1.X取切片并按数量(1, -2), nil)
		t.Assert(array1.X取切片并按数量(-5, 2), []string{"2", "3"})
		t.Assert(array1.X取切片并按数量(-10, 1), nil)
		t.Assert(array2.X取切片并按数量(0, 2), []string{"0", "1"})
	})
}

func TestStrArray_Rand(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"0", "1", "2", "3", "4", "5", "6"}
		array1 := 数组类.X创建文本并从数组(a1)
		t.Assert(len(array1.X取值随机多个(2)), "2")
		t.Assert(len(array1.X取值随机多个(10)), 10)
		t.AssertIN(array1.X取值随机多个(1)[0], a1)
		v, ok := array1.X取值随机()
		t.Assert(ok, true)
		t.AssertIN(v, a1)

		array2 := 数组类.X创建文本并从数组([]string{})
		v, ok = array2.X取值随机()
		t.Assert(ok, false)
		t.Assert(v, "")
		strArray := array2.X取值随机多个(1)
		t.AssertNil(strArray)
	})
}

func TestStrArray_PopRands(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"a", "b", "c", "d", "e", "f", "g"}
		array1 := 数组类.X创建文本并从数组(a1)
		t.AssertIN(array1.X出栈随机多个(1), []string{"a", "b", "c", "d", "e", "f", "g"})
		t.AssertIN(array1.X出栈随机多个(1), []string{"a", "b", "c", "d", "e", "f", "g"})
		t.AssertNI(array1.X出栈随机多个(1), array1.X取切片())
		t.AssertNI(array1.X出栈随机多个(1), array1.X取切片())
		t.Assert(len(array1.X出栈随机多个(10)), 3)
	})
}

func TestStrArray_Shuffle(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"0", "1", "2", "3", "4", "5", "6"}
		array1 := 数组类.X创建文本并从数组(a1)
		t.Assert(array1.X随机排序().X取长度(), 7)
	})
}

func TestStrArray_Reverse(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"0", "1", "2", "3", "4", "5", "6"}
		array1 := 数组类.X创建文本并从数组(a1)
		t.Assert(array1.X倒排序().X取切片(), []string{"6", "5", "4", "3", "2", "1", "0"})
	})
}

func TestStrArray_Join(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"0", "1", "2", "3", "4", "5", "6"}
		array1 := 数组类.X创建文本并从数组(a1)
		t.Assert(array1.X连接("."), `0.1.2.3.4.5.6`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"0", "1", `"a"`, `\a`}
		array1 := 数组类.X创建文本并从数组(a1)
		t.Assert(array1.X连接("."), `0.1."a".\a`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{}
		array1 := 数组类.X创建文本并从数组(a1)
		t.Assert(array1.X连接("."), "")
	})
}

func TestStrArray_String(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"0", "1", "2", "3", "4", "5", "6"}
		array1 := 数组类.X创建文本并从数组(a1)
		t.Assert(array1.String(), `["0","1","2","3","4","5","6"]`)

		array1 = nil
		t.Assert(array1.String(), "")
	})
}

func TestNewStrArrayFromCopy(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"0", "1", "2", "3", "4", "5", "6"}
		a2 := 数组类.X创建文本并从数组复制(a1)
		a3 := 数组类.X创建文本并从数组复制(a1, true)
		t.Assert(a2.X是否存在("1"), true)
		t.Assert(a2.X取长度(), 7)
		t.Assert(a2, a3)
	})
}

func TestStrArray_SetArray(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"0", "1", "2", "3", "4", "5", "6"}
		a2 := []string{"a", "b", "c", "d"}
		array1 := 数组类.X创建文本并从数组(a1)
		t.Assert(array1.X是否存在("2"), true)
		t.Assert(array1.X取长度(), 7)

		array1 = array1.X设置数组(a2)
		t.Assert(array1.X是否存在("2"), false)
		t.Assert(array1.X是否存在("c"), true)
		t.Assert(array1.X取长度(), 4)
	})
}

func TestStrArray_Replace(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"0", "1", "2", "3", "4", "5", "6"}
		a2 := []string{"a", "b", "c", "d"}
		a3 := []string{"o", "p", "q", "x", "y", "z", "w", "r", "v"}
		array1 := 数组类.X创建文本并从数组(a1)
		t.Assert(array1.X是否存在("2"), true)
		t.Assert(array1.X取长度(), 7)

		array1 = array1.X替换(a2)
		t.Assert(array1.X是否存在("2"), false)
		t.Assert(array1.X是否存在("c"), true)
		t.Assert(array1.X是否存在("5"), true)
		t.Assert(array1.X取长度(), 7)

		array1 = array1.X替换(a3)
		t.Assert(array1.X是否存在("2"), false)
		t.Assert(array1.X是否存在("c"), false)
		t.Assert(array1.X是否存在("5"), false)
		t.Assert(array1.X是否存在("p"), true)
		t.Assert(array1.X是否存在("r"), false)
		t.Assert(array1.X取长度(), 7)

	})
}

func TestStrArray_Sum(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"0", "1", "2", "3", "4", "5", "6"}
		a2 := []string{"0", "a", "3", "4", "5", "6"}
		array1 := 数组类.X创建文本并从数组(a1)
		array2 := 数组类.X创建文本并从数组(a2)
		t.Assert(array1.X求和(), 21)
		t.Assert(array2.X求和(), 18)
	})
}

func TestStrArray_PopRand(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"0", "1", "2", "3", "4", "5", "6"}
		array1 := 数组类.X创建文本并从数组(a1)
		str1, ok := array1.X出栈随机()
		t.Assert(strings.Contains("0,1,2,3,4,5,6", str1), true)
		t.Assert(array1.X取长度(), 6)
		t.Assert(ok, true)
	})
}

func TestStrArray_Clone(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"0", "1", "2", "3", "4", "5", "6"}
		array1 := 数组类.X创建文本并从数组(a1)
		array2 := array1.X取副本()
		t.Assert(array2, array1)
		t.Assert(array2.X取长度(), 7)
	})
}

func TestStrArray_CountValues(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"0", "1", "2", "3", "4", "4", "6"}
		array1 := 数组类.X创建文本并从数组(a1)

		m1 := array1.X统计()
		t.Assert(len(m1), 6)
		t.Assert(m1["2"], 1)
		t.Assert(m1["4"], 2)
	})
}

func TestStrArray_Remove(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []string{"e", "a", "d", "a", "c"}
		array1 := 数组类.X创建文本并从数组(a1)
		s1, ok := array1.X删除(1)
		t.Assert(s1, "a")
		t.Assert(ok, true)
		t.Assert(array1.X取长度(), 4)
		s1, ok = array1.X删除(3)
		t.Assert(s1, "c")
		t.Assert(ok, true)
		t.Assert(array1.X取长度(), 3)
	})
}

func TestStrArray_RLockFunc(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s1 := []string{"a", "b", "c", "d"}
		a1 := 数组类.X创建文本并从数组(s1, true)

		ch1 := make(chan int64, 3)
		ch2 := make(chan int64, 1)
		// go1
		go a1.X遍历读锁定(func(n1 []string) { // 读锁
			time.Sleep(2 * time.Second) // 暂停1秒
			n1[2] = "g"
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
		t.Assert(a1.X是否存在("g"), true)
	})
}

func TestStrArray_SortFunc(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s1 := []string{"a", "d", "c", "b"}
		a1 := 数组类.X创建文本并从数组(s1)
		func1 := func(v1, v2 string) bool {
			return v1 < v2
		}
		a11 := a1.X排序函数(func1)
		t.Assert(a11, []string{"a", "b", "c", "d"})
	})
}

func TestStrArray_LockFunc(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s1 := []string{"a", "b", "c", "d"}
		a1 := 数组类.X创建文本并从数组(s1, true)

		ch1 := make(chan int64, 3)
		ch2 := make(chan int64, 3)
		// go1
		go a1.X遍历写锁定(func(n1 []string) { // 读写锁
			time.Sleep(2 * time.Second) // 暂停2秒
			n1[2] = "g"
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
		t.Assert(a1.X是否存在("g"), true)
	})
}

func TestStrArray_Json(t *testing.T) {
	// array pointer
	单元测试类.C(t, func(t *单元测试类.T) {
		s1 := []string{"a", "b", "d", "c"}
		a1 := 数组类.X创建文本并从数组(s1)
		b1, err1 := json.Marshal(a1)
		b2, err2 := json.Marshal(s1)
		t.Assert(b1, b2)
		t.Assert(err1, err2)

		a2 := 数组类.X创建文本()
		err1 = json.UnmarshalUseNumber(b2, &a2)
		t.AssertNil(err1)
		t.Assert(a2.X取切片(), s1)

		var a3 数组类.StrArray
		err := json.UnmarshalUseNumber(b2, &a3)
		t.AssertNil(err)
		t.Assert(a3.X取切片(), s1)
	})
	// array value
	单元测试类.C(t, func(t *单元测试类.T) {
		s1 := []string{"a", "b", "d", "c"}
		a1 := *数组类.X创建文本并从数组(s1)
		b1, err1 := json.Marshal(a1)
		b2, err2 := json.Marshal(s1)
		t.Assert(b1, b2)
		t.Assert(err1, err2)

		a2 := 数组类.X创建文本()
		err1 = json.UnmarshalUseNumber(b2, &a2)
		t.AssertNil(err1)
		t.Assert(a2.X取切片(), s1)

		var a3 数组类.StrArray
		err := json.UnmarshalUseNumber(b2, &a3)
		t.AssertNil(err)
		t.Assert(a3.X取切片(), s1)
	})
	// array pointer
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Name   string
			Scores *数组类.StrArray
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
		t.Assert(user.Scores, data["Scores"])
	})
	// array value
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Name   string
			Scores 数组类.StrArray
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
		t.Assert(user.Scores, data["Scores"])
	})
}

func TestStrArray_Iterator(t *testing.T) {
	slice := g.SliceStr别名{"a", "b", "d", "c"}
	array := 数组类.X创建文本并从数组(slice)
	单元测试类.C(t, func(t *单元测试类.T) {
		array.X遍历(func(k int, v string) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array.X遍历升序(func(k int, v string) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array.X遍历降序(func(k int, v string) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		index := 0
		array.X遍历(func(k int, v string) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		index := 0
		array.X遍历升序(func(k int, v string) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		index := 0
		array.X遍历降序(func(k int, v string) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
}

func TestStrArray_RemoveValue(t *testing.T) {
	slice := g.SliceStr别名{"a", "b", "d", "c"}
	array := 数组类.X创建文本并从数组(slice)
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(array.X删除值("e"), false)
		t.Assert(array.X删除值("b"), true)
		t.Assert(array.X删除值("a"), true)
		t.Assert(array.X删除值("c"), true)
		t.Assert(array.X删除值("f"), false)
	})
}

func TestStrArray_RemoveValues(t *testing.T) {
	slice := g.SliceStr别名{"a", "b", "d", "c"}
	array := 数组类.X创建文本并从数组(slice)
	单元测试类.C(t, func(t *单元测试类.T) {
		array.X删除多个值("a", "b", "c")
		t.Assert(array.X取切片(), g.SliceStr别名{"d"})
	})
}

func TestStrArray_UnmarshalValue(t *testing.T) {
	type V struct {
		Name  string
		Array *数组类.StrArray
	}
	// JSON
	单元测试类.C(t, func(t *单元测试类.T) {
		var v *V
		err := 转换类.Struct(g.Map{
			"name":  "john",
			"array": []byte(`["1","2","3"]`),
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Array.X取切片(), g.SliceStr别名{"1", "2", "3"})
	})
	// Map
	单元测试类.C(t, func(t *单元测试类.T) {
		var v *V
		err := 转换类.Struct(g.Map{
			"name":  "john",
			"array": g.SliceStr别名{"1", "2", "3"},
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Array.X取切片(), g.SliceStr别名{"1", "2", "3"})
	})
}
func TestStrArray_Filter(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建文本并从数组(g.SliceStr别名{"", "1", "2", "0"})
		t.Assert(array.X遍历删除(func(index int, value string) bool {
			return empty.IsEmpty(value)
		}), g.SliceStr别名{"1", "2", "0"})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建文本并从数组(g.SliceStr别名{"1", "2"})
		t.Assert(array.X遍历删除(func(index int, value string) bool {
			return empty.IsEmpty(value)
		}), g.SliceStr别名{"1", "2"})
	})
}

func TestStrArray_FilterEmpty(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建文本并从数组(g.SliceStr别名{"", "1", "2", "0"})
		t.Assert(array.X删除所有空值(), g.SliceStr别名{"1", "2", "0"})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建文本并从数组(g.SliceStr别名{"1", "2"})
		t.Assert(array.X删除所有空值(), g.SliceStr别名{"1", "2"})
	})
}

func TestStrArray_Walk(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建文本并从数组(g.SliceStr别名{"1", "2"})
		t.Assert(array.X遍历修改(func(value string) string {
			return "key-" + value
		}), g.Slice别名{"key-1", "key-2"})
	})
}
