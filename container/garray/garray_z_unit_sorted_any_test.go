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
	"github.com/888go/goframe/util/gutil"
)

func TestSortedArray_NewSortedArrayFrom(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "f", "c"}
		a2 := []interface{}{"h", "j", "i", "k"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		func2 := func(v1, v2 interface{}) int {
			return -1
		}
		array1 := 数组类.X创建排序并从数组(a1, func1)
		array2 := 数组类.X创建排序并从数组(a2, func2)

		t.Assert(array1.X取长度(), 3)
		t.Assert(array1, []interface{}{"a", "c", "f"})

		t.Assert(array2.X取长度(), 4)
		t.Assert(array2, []interface{}{"k", "i", "j", "h"})
	})
}

func TestNewSortedArrayFromCopy(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "f", "c"}

		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		func2 := func(v1, v2 interface{}) int {
			return -1
		}
		array1 := 数组类.X创建排序并从数组复制(a1, func1)
		array2 := 数组类.X创建排序并从数组复制(a1, func2)
		t.Assert(array1.X取长度(), 3)
		t.Assert(array1, []interface{}{"a", "c", "f"})
		t.Assert(array1.X取长度(), 3)
		t.Assert(array2, []interface{}{"c", "f", "a"})
	})
}

func TestNewSortedArrayRange(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		func1 := func(v1, v2 interface{}) int {
			return 转换类.X取整数(v1) - 转换类.X取整数(v2)
		}

		array1 := 数组类.X创建排序并按范围(1, 5, 1, func1)
		t.Assert(array1.X取长度(), 5)
		t.Assert(array1, []interface{}{1, 2, 3, 4, 5})
	})
}

func TestSortedArray_SetArray(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "f", "c"}
		a2 := []interface{}{"e", "h", "g", "k"}

		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}

		array1 := 数组类.X创建排序并从数组(a1, func1)
		array1.X设置数组(a2)
		t.Assert(array1.X取长度(), 4)
		t.Assert(array1, []interface{}{"e", "g", "h", "k"})
	})

}

func TestSortedArray_Sort(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "f", "c"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		array1 := 数组类.X创建排序并从数组(a1, func1)
		array1.X排序递增()
		t.Assert(array1.X取长度(), 3)
		t.Assert(array1, []interface{}{"a", "c", "f"})
	})

}

func TestSortedArray_Get(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "f", "c"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		array1 := 数组类.X创建排序并从数组(a1, func1)
		v, ok := array1.X取值2(2)
		t.Assert(v, "f")
		t.Assert(ok, true)

		v, ok = array1.X取值2(1)
		t.Assert(v, "c")
		t.Assert(ok, true)

		v, ok = array1.X取值2(99)
		t.Assert(v, nil)
		t.Assert(ok, false)
	})

}

func TestSortedArray_At(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "f", "c"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		array1 := 数组类.X创建排序并从数组(a1, func1)
		v := array1.X取值(2)
		t.Assert(v, "f")
	})
}

func TestSortedArray_Remove(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "d", "c", "b"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		array1 := 数组类.X创建排序并从数组(a1, func1)
		i1, ok := array1.X删除(1)
		t.Assert(ok, true)
		t.Assert(转换类.String(i1), "b")
		t.Assert(array1.X取长度(), 3)
		t.Assert(array1.X是否存在("b"), false)

		v, ok := array1.X删除(-1)
		t.Assert(v, nil)
		t.Assert(ok, false)

		v, ok = array1.X删除(100000)
		t.Assert(v, nil)
		t.Assert(ok, false)

		i2, ok := array1.X删除(0)
		t.Assert(ok, true)
		t.Assert(转换类.String(i2), "a")
		t.Assert(array1.X取长度(), 2)
		t.Assert(array1.X是否存在("a"), false)

		i3, ok := array1.X删除(1)
		t.Assert(ok, true)
		t.Assert(转换类.String(i3), "d")
		t.Assert(array1.X取长度(), 1)
		t.Assert(array1.X是否存在("d"), false)
	})

}

func TestSortedArray_PopLeft(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array1 := 数组类.X创建排序并从数组(
			[]interface{}{"a", "d", "c", "b"},
			工具类.X比较文本,
		)
		i1, ok := array1.X出栈左()
		t.Assert(ok, true)
		t.Assert(转换类.String(i1), "a")
		t.Assert(array1.X取长度(), 3)
		t.Assert(array1, []interface{}{"b", "c", "d"})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建排序并从数组(g.Slice别名{1, 2, 3}, 工具类.X比较整数)
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

func TestSortedArray_PopRight(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array1 := 数组类.X创建排序并从数组(
			[]interface{}{"a", "d", "c", "b"},
			工具类.X比较文本,
		)
		i1, ok := array1.X出栈右()
		t.Assert(ok, true)
		t.Assert(转换类.String(i1), "d")
		t.Assert(array1.X取长度(), 3)
		t.Assert(array1, []interface{}{"a", "b", "c"})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建排序并从数组(g.Slice别名{1, 2, 3}, 工具类.X比较整数)
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

func TestSortedArray_PopRand(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "d", "c", "b"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		array1 := 数组类.X创建排序并从数组(a1, func1)
		i1, ok := array1.X出栈随机()
		t.Assert(ok, true)
		t.AssertIN(i1, []interface{}{"a", "d", "c", "b"})
		t.Assert(array1.X取长度(), 3)

	})
}

func TestSortedArray_PopRands(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "d", "c", "b"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		array1 := 数组类.X创建排序并从数组(a1, func1)
		i1 := array1.X出栈随机多个(2)
		t.Assert(len(i1), 2)
		t.AssertIN(i1, []interface{}{"a", "d", "c", "b"})
		t.Assert(array1.X取长度(), 2)

		i2 := array1.X出栈随机多个(3)
		t.Assert(len(i1), 2)
		t.AssertIN(i2, []interface{}{"a", "d", "c", "b"})
		t.Assert(array1.X取长度(), 0)

	})
}

func TestSortedArray_Empty(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建排序(工具类.X比较整数)
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
}

func TestSortedArray_PopLefts(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "d", "c", "b", "e", "f"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		array1 := 数组类.X创建排序并从数组(a1, func1)
		i1 := array1.X出栈左多个(2)
		t.Assert(len(i1), 2)
		t.AssertIN(i1, []interface{}{"a", "d", "c", "b", "e", "f"})
		t.Assert(array1.X取长度(), 4)

		i2 := array1.X出栈左多个(5)
		t.Assert(len(i2), 4)
		t.AssertIN(i1, []interface{}{"a", "d", "c", "b", "e", "f"})
		t.Assert(array1.X取长度(), 0)
	})
}

func TestSortedArray_PopRights(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "d", "c", "b", "e", "f"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		array1 := 数组类.X创建排序并从数组(a1, func1)
		i1 := array1.X出栈右多个(2)
		t.Assert(len(i1), 2)
		t.Assert(i1, []interface{}{"e", "f"})
		t.Assert(array1.X取长度(), 4)

		i2 := array1.X出栈右多个(10)
		t.Assert(len(i2), 4)
		t.Assert(array1.X取长度(), 0)
	})
}

func TestSortedArray_Range(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "d", "c", "b", "e", "f"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		array1 := 数组类.X创建排序并从数组(a1, func1)
		array2 := 数组类.X创建排序并从数组(a1, func1, true)
		i1 := array1.X取切片并按范围(2, 5)
		t.Assert(i1, []interface{}{"c", "d", "e"})
		t.Assert(array1.X取长度(), 6)

		i2 := array1.X取切片并按范围(7, 5)
		t.Assert(len(i2), 0)
		i2 = array1.X取切片并按范围(-1, 2)
		t.Assert(i2, []interface{}{"a", "b"})

		i2 = array1.X取切片并按范围(4, 10)
		t.Assert(len(i2), 2)
		t.Assert(i2, []interface{}{"e", "f"})

		t.Assert(array2.X取切片并按范围(1, 3), []interface{}{"b", "c"})

	})
}

func TestSortedArray_Sum(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "d", "c", "b", "e", "f"}
		a2 := []interface{}{"1", "2", "3", "b", "e", "f"}
		a3 := []interface{}{"4", "5", "6"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		array1 := 数组类.X创建排序并从数组(a1, func1)
		array2 := 数组类.X创建排序并从数组(a2, func1)
		array3 := 数组类.X创建排序并从数组(a3, func1)
		t.Assert(array1.X求和(), 0)
		t.Assert(array2.X求和(), 6)
		t.Assert(array3.X求和(), 15)

	})
}

func TestSortedArray_Clone(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "d", "c", "b", "e", "f"}

		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		array1 := 数组类.X创建排序并从数组(a1, func1)
		array2 := array1.X取副本()
		t.Assert(array1, array2)
		array1.X删除(1)
		t.AssertNE(array1, array2)

	})
}

func TestSortedArray_Clear(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "d", "c", "b", "e", "f"}

		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		array1 := 数组类.X创建排序并从数组(a1, func1)
		t.Assert(array1.X取长度(), 6)
		array1.X清空()
		t.Assert(array1.X取长度(), 0)

	})
}

func TestSortedArray_Chunk(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "d", "c", "b", "e"}

		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		array1 := 数组类.X创建排序并从数组(a1, func1)
		i1 := array1.X分割(2)
		t.Assert(len(i1), 3)
		t.Assert(i1[0], []interface{}{"a", "b"})
		t.Assert(i1[2], []interface{}{"e"})

		i1 = array1.X分割(0)
		t.Assert(len(i1), 0)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{1, 2, 3, 4, 5}
		array1 := 数组类.X创建排序并从数组(a1, 工具类.X比较整数)
		chunks := array1.X分割(3)
		t.Assert(len(chunks), 2)
		t.Assert(chunks[0], []interface{}{1, 2, 3})
		t.Assert(chunks[1], []interface{}{4, 5})
		t.Assert(array1.X分割(0), nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{1, 2, 3, 4, 5, 6}
		array1 := 数组类.X创建排序并从数组(a1, 工具类.X比较整数)
		chunks := array1.X分割(2)
		t.Assert(len(chunks), 3)
		t.Assert(chunks[0], []interface{}{1, 2})
		t.Assert(chunks[1], []interface{}{3, 4})
		t.Assert(chunks[2], []interface{}{5, 6})
		t.Assert(array1.X分割(0), nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{1, 2, 3, 4, 5, 6}
		array1 := 数组类.X创建排序并从数组(a1, 工具类.X比较整数)
		chunks := array1.X分割(3)
		t.Assert(len(chunks), 2)
		t.Assert(chunks[0], []interface{}{1, 2, 3})
		t.Assert(chunks[1], []interface{}{4, 5, 6})
		t.Assert(array1.X分割(0), nil)
	})
}

func TestSortedArray_SubSlice(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "d", "c", "b", "e"}

		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		array1 := 数组类.X创建排序并从数组(a1, func1)
		array2 := 数组类.X创建排序并从数组(a1, func1, true)
		i1 := array1.X取切片并按数量(2, 3)
		t.Assert(len(i1), 3)
		t.Assert(i1, []interface{}{"c", "d", "e"})

		i1 = array1.X取切片并按数量(2, 6)
		t.Assert(len(i1), 3)
		t.Assert(i1, []interface{}{"c", "d", "e"})

		i1 = array1.X取切片并按数量(7, 2)
		t.Assert(len(i1), 0)

		s1 := array1.X取切片并按数量(1, -2)
		t.Assert(s1, nil)

		s1 = array1.X取切片并按数量(-9, 2)
		t.Assert(s1, nil)
		t.Assert(array2.X取切片并按数量(1, 3), []interface{}{"b", "c", "d"})

	})
}

func TestSortedArray_Rand(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "d", "c"}

		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		array1 := 数组类.X创建排序并从数组(a1, func1)
		i1, ok := array1.X取值随机()
		t.Assert(ok, true)
		t.AssertIN(i1, []interface{}{"a", "d", "c"})
		t.Assert(array1.X取长度(), 3)

		array2 := 数组类.X创建排序并从数组([]interface{}{}, func1)
		v, ok := array2.X取值随机()
		t.Assert(ok, false)
		t.Assert(v, nil)
	})
}

func TestSortedArray_Rands(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "d", "c"}

		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		array1 := 数组类.X创建排序并从数组(a1, func1)
		i1 := array1.X取值随机多个(2)
		t.AssertIN(i1, []interface{}{"a", "d", "c"})
		t.Assert(len(i1), 2)
		t.Assert(array1.X取长度(), 3)

		i1 = array1.X取值随机多个(4)
		t.Assert(len(i1), 4)

		array2 := 数组类.X创建排序并从数组([]interface{}{}, func1)
		v := array2.X取值随机多个(1)
		t.Assert(v, nil)
	})
}

func TestSortedArray_Join(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "d", "c"}
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		array1 := 数组类.X创建排序并从数组(a1, func1)
		t.Assert(array1.X连接(","), `a,c,d`)
		t.Assert(array1.X连接("."), `a.c.d`)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{0, 1, `"a"`, `\a`}
		array1 := 数组类.X创建排序并从数组(a1, 工具类.X比较文本)
		t.Assert(array1.X连接("."), `"a".0.1.\a`)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{}
		array1 := 数组类.X创建排序并从数组(a1, 工具类.X比较文本)
		t.Assert(array1.X连接("."), "")
	})
}

func TestSortedArray_String(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{0, 1, "a", "b"}
		array1 := 数组类.X创建排序并从数组(a1, 工具类.X比较文本)
		t.Assert(array1.String(), `[0,1,"a","b"]`)

		array1 = nil
		t.Assert(array1.String(), "")
	})
}

func TestSortedArray_CountValues(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{"a", "d", "c", "c"}

		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		array1 := 数组类.X创建排序并从数组(a1, func1)
		m1 := array1.X统计()
		t.Assert(len(m1), 3)
		t.Assert(m1["c"], 2)
		t.Assert(m1["a"], 1)

	})
}

func TestSortedArray_SetUnique(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{1, 2, 3, 4, 5, 3, 2, 2, 3, 5, 5}
		array1 := 数组类.X创建排序并从数组(a1, 工具类.X比较整数)
		array1.X设置去重(true)
		t.Assert(array1.X取长度(), 5)
		t.Assert(array1, []interface{}{1, 2, 3, 4, 5})
	})
}

func TestSortedArray_Unique(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a1 := []interface{}{1, 2, 3, 4, 5, 3, 2, 2, 3, 5, 5}
		array1 := 数组类.X创建排序并从数组(a1, 工具类.X比较整数)
		array1.X去重()
		t.Assert(array1.X取长度(), 5)
		t.Assert(array1, []interface{}{1, 2, 3, 4, 5})

		array2 := 数组类.X创建排序并从数组([]interface{}{}, 工具类.X比较整数)
		array2.X去重()
		t.Assert(array2.X取长度(), 0)
		t.Assert(array2, []interface{}{})
	})
}

func TestSortedArray_LockFunc(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		s1 := []interface{}{"a", "b", "c", "d"}
		a1 := 数组类.X创建排序并从数组(s1, func1, true)

		ch1 := make(chan int64, 3)
		ch2 := make(chan int64, 3)
		// go1
		go a1.X遍历写锁定(func(n1 []interface{}) { // 读写锁
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

func TestSortedArray_RLockFunc(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		func1 := func(v1, v2 interface{}) int {
			return strings.Compare(转换类.String(v1), 转换类.String(v2))
		}
		s1 := []interface{}{"a", "b", "c", "d"}
		a1 := 数组类.X创建排序并从数组(s1, func1, true)

		ch1 := make(chan int64, 3)
		ch2 := make(chan int64, 3)
		// go1
		go a1.X遍历读锁定(func(n1 []interface{}) { // 读写锁
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
		t.AssertLT(t2-t1, 20) // go1加的读锁，所go2读的时候不会被阻塞。
		t.Assert(a1.X是否存在("g"), true)
	})
}

func TestSortedArray_Merge(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		func1 := func(v1, v2 interface{}) int {
			if 转换类.X取整数(v1) < 转换类.X取整数(v2) {
				return 0
			}
			return 1
		}

		s1 := []interface{}{"a", "b", "c", "d"}
		s2 := []string{"e", "f"}
		i1 := 数组类.X创建整数并从数组([]int{1, 2, 3})
		i2 := 数组类.X创建并从数组([]interface{}{3})
		s3 := 数组类.X创建文本并从数组([]string{"g", "h"})
		s4 := 数组类.X创建排序并从数组([]interface{}{4, 5}, func1)
		s5 := 数组类.X创建文本排序并从数组(s2)
		s6 := 数组类.X创建整数排序并从数组([]int{1, 2, 3})

		a1 := 数组类.X创建排序并从数组(s1, func1)

		t.Assert(a1.X合并(s2).X取长度(), 6)
		t.Assert(a1.X合并(i1).X取长度(), 9)
		t.Assert(a1.X合并(i2).X取长度(), 10)
		t.Assert(a1.X合并(s3).X取长度(), 12)
		t.Assert(a1.X合并(s4).X取长度(), 14)
		t.Assert(a1.X合并(s5).X取长度(), 16)
		t.Assert(a1.X合并(s6).X取长度(), 19)
	})
}

func TestSortedArray_Json(t *testing.T) {
	// array pointer
	单元测试类.C(t, func(t *单元测试类.T) {
		s1 := []interface{}{"a", "b", "d", "c"}
		s2 := []interface{}{"a", "b", "c", "d"}
		a1 := 数组类.X创建排序并从数组(s1, 工具类.X比较文本)
		b1, err1 := json.Marshal(a1)
		b2, err2 := json.Marshal(s1)
		t.Assert(b1, b2)
		t.Assert(err1, err2)

		a2 := 数组类.X创建排序(工具类.X比较文本)
		err1 = json.UnmarshalUseNumber(b2, &a2)
		t.AssertNil(err1)
		t.Assert(a2.X取切片(), s2)

		var a3 数组类.SortedArray
		err := json.UnmarshalUseNumber(b2, &a3)
		t.AssertNil(err)
		t.Assert(a3.X取切片(), s1)
		t.Assert(a3.X取any数组(), s1)
	})
	// array value
	单元测试类.C(t, func(t *单元测试类.T) {
		s1 := []interface{}{"a", "b", "d", "c"}
		s2 := []interface{}{"a", "b", "c", "d"}
		a1 := *数组类.X创建排序并从数组(s1, 工具类.X比较文本)
		b1, err1 := json.Marshal(a1)
		b2, err2 := json.Marshal(s1)
		t.Assert(b1, b2)
		t.Assert(err1, err2)

		a2 := 数组类.X创建排序(工具类.X比较文本)
		err1 = json.UnmarshalUseNumber(b2, &a2)
		t.AssertNil(err1)
		t.Assert(a2.X取切片(), s2)

		var a3 数组类.SortedArray
		err := json.UnmarshalUseNumber(b2, &a3)
		t.AssertNil(err)
		t.Assert(a3.X取切片(), s1)
		t.Assert(a3.X取any数组(), s1)
	})
	// array pointer
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Name   string
			Scores *数组类.SortedArray
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
		t.AssertNE(user.Scores, nil)
		t.Assert(user.Scores.X取长度(), 3)

		v, ok := user.Scores.X出栈左()
		t.AssertIN(v, data["Scores"])
		t.Assert(ok, true)

		v, ok = user.Scores.X出栈左()
		t.AssertIN(v, data["Scores"])
		t.Assert(ok, true)

		v, ok = user.Scores.X出栈左()
		t.AssertIN(v, data["Scores"])
		t.Assert(ok, true)

		v, ok = user.Scores.X出栈左()
		t.Assert(v, nil)
		t.Assert(ok, false)
	})
	// array value
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Name   string
			Scores 数组类.SortedArray
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
		t.AssertNE(user.Scores, nil)
		t.Assert(user.Scores.X取长度(), 3)

		v, ok := user.Scores.X出栈左()
		t.AssertIN(v, data["Scores"])
		t.Assert(ok, true)

		v, ok = user.Scores.X出栈左()
		t.AssertIN(v, data["Scores"])
		t.Assert(ok, true)

		v, ok = user.Scores.X出栈左()
		t.AssertIN(v, data["Scores"])
		t.Assert(ok, true)

		v, ok = user.Scores.X出栈左()
		t.Assert(v, nil)
		t.Assert(ok, false)
	})
}

func TestSortedArray_Iterator(t *testing.T) {
	slice := g.Slice别名{"a", "b", "d", "c"}
	array := 数组类.X创建排序并从数组(slice, 工具类.X比较文本)
	单元测试类.C(t, func(t *单元测试类.T) {
		array.X遍历(func(k int, v interface{}) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array.X遍历升序(func(k int, v interface{}) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array.X遍历降序(func(k int, v interface{}) bool {
			t.Assert(v, slice[k])
			return true
		})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		index := 0
		array.X遍历(func(k int, v interface{}) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		index := 0
		array.X遍历升序(func(k int, v interface{}) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		index := 0
		array.X遍历降序(func(k int, v interface{}) bool {
			index++
			return false
		})
		t.Assert(index, 1)
	})
}

func TestSortedArray_RemoveValue(t *testing.T) {
	slice := g.Slice别名{"a", "b", "d", "c"}
	array := 数组类.X创建排序并从数组(slice, 工具类.X比较文本)
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(array.X删除值("e"), false)
		t.Assert(array.X删除值("b"), true)
		t.Assert(array.X删除值("a"), true)
		t.Assert(array.X删除值("c"), true)
		t.Assert(array.X删除值("f"), false)
	})
}

func TestSortedArray_RemoveValues(t *testing.T) {
	slice := g.Slice别名{"a", "b", "d", "c"}
	array := 数组类.X创建排序并从数组(slice, 工具类.X比较文本)
	单元测试类.C(t, func(t *单元测试类.T) {
		array.X删除多个值("a", "b", "c")
		t.Assert(array.X取切片(), g.SliceStr别名{"d"})
	})
}

func TestSortedArray_UnmarshalValue(t *testing.T) {
	type V struct {
		Name  string
		Array *数组类.SortedArray
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
func TestSortedArray_Filter(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		values := g.Slice别名{0, 1, 2, 3, 4, "", g.Slice别名{}}
		array := 数组类.X创建排序并从数组复制(values, 工具类.X比较整数)
		t.Assert(array.X遍历删除(func(index int, value interface{}) bool {
			return empty.X是否为Nil(value)
		}).X取切片(), g.Slice别名{0, "", g.Slice别名{}, 1, 2, 3, 4})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建排序并从数组复制(g.Slice别名{nil, 1, 2, 3, 4, nil}, 工具类.X比较整数)
		t.Assert(array.X遍历删除(func(index int, value interface{}) bool {
			return empty.X是否为Nil(value)
		}), g.Slice别名{1, 2, 3, 4})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建排序并从数组(g.Slice别名{0, 1, 2, 3, 4, "", g.Slice别名{}}, 工具类.X比较整数)
		t.Assert(array.X遍历删除(func(index int, value interface{}) bool {
			return empty.IsEmpty(value)
		}), g.Slice别名{1, 2, 3, 4})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建排序并从数组(g.Slice别名{1, 2, 3, 4}, 工具类.X比较整数)
		t.Assert(array.X遍历删除(func(index int, value interface{}) bool {
			return empty.IsEmpty(value)
		}), g.Slice别名{1, 2, 3, 4})
	})
}

func TestSortedArray_FilterNil(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		values := g.Slice别名{0, 1, 2, 3, 4, "", g.Slice别名{}}
		array := 数组类.X创建排序并从数组复制(values, 工具类.X比较整数)
		t.Assert(array.X删除所有nil().X取切片(), g.Slice别名{0, "", g.Slice别名{}, 1, 2, 3, 4})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建排序并从数组复制(g.Slice别名{nil, 1, 2, 3, 4, nil}, 工具类.X比较整数)
		t.Assert(array.X删除所有nil(), g.Slice别名{1, 2, 3, 4})
	})
}

func TestSortedArray_FilterEmpty(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建排序并从数组(g.Slice别名{0, 1, 2, 3, 4, "", g.Slice别名{}}, 工具类.X比较整数)
		t.Assert(array.X删除所有空值(), g.Slice别名{1, 2, 3, 4})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建排序并从数组(g.Slice别名{1, 2, 3, 4}, 工具类.X比较整数)
		t.Assert(array.X删除所有空值(), g.Slice别名{1, 2, 3, 4})
	})
}

func TestSortedArray_Walk(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建排序并从数组(g.Slice别名{"1", "2"}, 工具类.X比较文本)
		t.Assert(array.X遍历修改(func(value interface{}) interface{} {
			return "key-" + 转换类.String(value)
		}), g.Slice别名{"key-1", "key-2"})
	})
}

func TestSortedArray_IsEmpty(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建排序并从数组([]interface{}{}, 工具类.X比较文本)
		t.Assert(array.X是否为空(), true)
	})
}

func TestSortedArray_DeepCopy(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := 数组类.X创建排序并从数组([]interface{}{1, 2, 3, 4, 5}, 工具类.X比较文本)
		copyArray := array.DeepCopy().(*数组类.SortedArray)
		array.X入栈右(6)
		copyArray.X入栈右(7)
		cval, _ := copyArray.X取值2(5)
		val, _ := array.X取值2(5)
		t.AssertNE(cval, val)
	})
}
