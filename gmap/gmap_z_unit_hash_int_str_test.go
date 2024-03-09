// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

package map类_test

import (
	"testing"
	
	"github.com/gogf/gf/v2/container/garray"
	"github.com/888go/goframe/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/888go/goframe/gmap/internal/json"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

func getStr() string {
	return "z"
}

func Test_IntStrMap_Var(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var m map类.IntStrMap
		m.X设置值(1, "a")

		t.Assert(m.X取值(1), "a")
		t.Assert(m.X取数量(), 1)
		t.Assert(m.X是否为空(), false)

		t.Assert(m.X取值或设置值(2, "b"), "b")
		t.Assert(m.X设置值并跳过已存在(2, "b"), false)

		t.Assert(m.X设置值并跳过已存在(3, "c"), true)

		t.Assert(m.X删除(2), "b")
		t.Assert(m.X是否存在(2), false)

		t.AssertIN(3, m.X取所有名称())
		t.AssertIN(1, m.X取所有名称())
		t.AssertIN("a", m.X取所有值())
		t.AssertIN("c", m.X取所有值())

		m_f := map类.X创建IntStr()
		m_f.X设置值(1, "2")
		m_f.X名称值交换()
		t.Assert(m_f.X取Map(), map[int]string{2: "1"})

		m.X清空()
		t.Assert(m.X取数量(), 0)
		t.Assert(m.X是否为空(), true)
	})
}

func Test_IntStrMap_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntStr()
		m.X设置值(1, "a")

		t.Assert(m.X取值(1), "a")
		t.Assert(m.X取数量(), 1)
		t.Assert(m.X是否为空(), false)

		t.Assert(m.X取值或设置值(1, "a"), "a")
		t.Assert(m.X取值或设置值(2, "b"), "b")
		t.Assert(m.X设置值并跳过已存在(2, "b"), false)

		t.Assert(m.X设置值并跳过已存在(3, "c"), true)

		t.Assert(m.X删除(2), "b")
		t.Assert(m.X是否存在(2), false)

		t.AssertIN(3, m.X取所有名称())
		t.AssertIN(1, m.X取所有名称())
		t.AssertIN("a", m.X取所有值())
		t.AssertIN("c", m.X取所有值())

		// 反转之后不成为以下 map,flip 操作只是翻转原 map
		// t.Assert(m.Map(), map[string]int{"a": 1, "c": 3})
		m_f := map类.X创建IntStr()
		m_f.X设置值(1, "2")
		m_f.X名称值交换()
		t.Assert(m_f.X取Map(), map[int]string{2: "1"})

		m.X清空()
		t.Assert(m.X取数量(), 0)
		t.Assert(m.X是否为空(), true)

		m2 := map类.X创建IntStr并从Map(map[int]string{1: "a", 2: "b"})
		t.Assert(m2.X取Map(), map[int]string{1: "a", 2: "b"})
	})

	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntStr(true)
		m.X设置值(1, "val1")
		t.Assert(m.X取Map(), map[int]string{1: "val1"})
	})
}

func TestIntStrMap_MapStrAny(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntStr()
		m.X取值或设置值_函数(1, getStr)
		m.X取值或设置值_函数带锁(2, getStr)
		t.Assert(m.X取MapStrAny(), g.MapStrAny{"1": "z", "2": "z"})
	})
}

func TestIntStrMap_Sets(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntStr并从Map(nil)
		m.X设置值Map(g.MapIntStr{1: "z", 2: "z"})
		t.Assert(len(m.X取Map()), 2)
	})
}

func Test_IntStrMap_Set_Fun(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntStr()
		m.X取值或设置值_函数(1, getStr)
		m.X取值或设置值_函数带锁(2, getStr)
		t.Assert(m.X取值或设置值_函数(1, getStr), "z")
		t.Assert(m.X取值或设置值_函数带锁(2, getStr), "z")
		t.Assert(m.X取值(1), "z")
		t.Assert(m.X取值(2), "z")
		t.Assert(m.X设置值并跳过已存在_函数(1, getStr), false)
		t.Assert(m.X设置值并跳过已存在_函数(3, getStr), true)

		t.Assert(m.X设置值并跳过已存在_函数带锁(2, getStr), false)
		t.Assert(m.X设置值并跳过已存在_函数带锁(4, getStr), true)
	})

	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntStr并从Map(nil)
		t.Assert(m.X取值或设置值_函数带锁(1, getStr), "z")
	})

	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntStr并从Map(nil)
		t.Assert(m.X设置值并跳过已存在_函数带锁(1, getStr), true)
	})
}

func Test_IntStrMap_Batch(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntStr()
		m.X设置值Map(map[int]string{1: "a", 2: "b", 3: "c"})
		t.Assert(m.X取Map(), map[int]string{1: "a", 2: "b", 3: "c"})
		m.X删除多个值([]int{1, 2})
		t.Assert(m.X取Map(), map[int]interface{}{3: "c"})
	})
}

func Test_IntStrMap_Iterator(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := map[int]string{1: "a", 2: "b"}
		m := map类.X创建IntStr并从Map(expect)
		m.X遍历(func(k int, v string) bool {
			t.Assert(expect[k], v)
			return true
		})
		// 断言返回值对遍历控制
		i := 0
		j := 0
		m.X遍历(func(k int, v string) bool {
			i++
			return true
		})
		m.X遍历(func(k int, v string) bool {
			j++
			return false
		})
		t.Assert(i, 2)
		t.Assert(j, 1)
	})
}

func Test_IntStrMap_Lock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := map[int]string{1: "a", 2: "b", 3: "c"}
		m := map类.X创建IntStr并从Map(expect)
		m.X遍历写锁定(func(m map[int]string) {
			t.Assert(m, expect)
		})
		m.X遍历读锁定(func(m map[int]string) {
			t.Assert(m, expect)
		})
	})
}

func Test_IntStrMap_Clone(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// clone 方法是深克隆
		m := map类.X创建IntStr并从Map(map[int]string{1: "a", 2: "b", 3: "c"})

		m_clone := m.X取副本()
		m.X删除(1)
		// 修改原 map,clone 后的 map 不影响
		t.AssertIN(1, m_clone.X取所有名称())

		m_clone.X删除(2)
		// 修改clone map,原 map 不影响
		t.AssertIN(2, m.X取所有名称())
	})
}

func Test_IntStrMap_Merge(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := map类.X创建IntStr()
		m2 := map类.X创建IntStr()
		m1.X设置值(1, "a")
		m2.X设置值(2, "b")
		m1.X合并(m2)
		t.Assert(m1.X取Map(), map[int]string{1: "a", 2: "b"})

		m3 := map类.X创建IntStr并从Map(nil)
		m3.X合并(m2)
		t.Assert(m3.X取Map(), m2.X取Map())
	})
}

func Test_IntStrMap_Map(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntStr()
		m.X设置值(1, "0")
		m.X设置值(2, "2")
		t.Assert(m.X取值(1), "0")
		t.Assert(m.X取值(2), "2")
		data := m.X取Map()
		t.Assert(data[1], "0")
		t.Assert(data[2], "2")
		data[3] = "3"
		t.Assert(m.X取值(3), "3")
		m.X设置值(4, "4")
		t.Assert(data[4], "4")
	})
}

func Test_IntStrMap_MapCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntStr()
		m.X设置值(1, "0")
		m.X设置值(2, "2")
		t.Assert(m.X取值(1), "0")
		t.Assert(m.X取值(2), "2")
		data := m.X浅拷贝()
		t.Assert(data[1], "0")
		t.Assert(data[2], "2")
		data[3] = "3"
		t.Assert(m.X取值(3), "")
		m.X设置值(4, "4")
		t.Assert(data[4], "")
	})
}

func Test_IntStrMap_FilterEmpty(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntStr()
		m.X设置值(1, "")
		m.X设置值(2, "2")
		t.Assert(m.X取数量(), 2)
		t.Assert(m.X取值(2), "2")
		m.X删除所有空值()
		t.Assert(m.X取数量(), 1)
		t.Assert(m.X取值(2), "2")
	})
}

func Test_IntStrMap_Json(t *testing.T) {
	// Marshal
	gtest.C(t, func(t *gtest.T) {
		data := g.MapIntStr{
			1: "v1",
			2: "v2",
		}
		m1 := map类.X创建IntStr并从Map(data)
		b1, err1 := json.Marshal(m1)
		b2, err2 := json.Marshal(data)
		t.Assert(err1, err2)
		t.Assert(b1, b2)
	})
	// Unmarshal
	gtest.C(t, func(t *gtest.T) {
		data := g.MapIntStr{
			1: "v1",
			2: "v2",
		}
		b, err := json.Marshal(data)
		t.AssertNil(err)

		m := map类.X创建IntStr()
		err = json.UnmarshalUseNumber(b, m)
		t.AssertNil(err)
		t.Assert(m.X取值(1), data[1])
		t.Assert(m.X取值(2), data[2])
	})
}

func Test_IntStrMap_Pop(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntStr并从Map(g.MapIntStr{
			1: "v1",
			2: "v2",
		})
		t.Assert(m.X取数量(), 2)

		k1, v1 := m.X出栈()
		t.AssertIN(k1, g.Slice{1, 2})
		t.AssertIN(v1, g.Slice{"v1", "v2"})
		t.Assert(m.X取数量(), 1)
		k2, v2 := m.X出栈()
		t.AssertIN(k2, g.Slice{1, 2})
		t.AssertIN(v2, g.Slice{"v1", "v2"})
		t.Assert(m.X取数量(), 0)

		t.AssertNE(k1, k2)
		t.AssertNE(v1, v2)

		k3, v3 := m.X出栈()
		t.Assert(k3, 0)
		t.Assert(v3, "")
	})
}

func Test_IntStrMap_Pops(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntStr并从Map(g.MapIntStr{
			1: "v1",
			2: "v2",
			3: "v3",
		})
		t.Assert(m.X取数量(), 3)

		kArray := garray.New()
		vArray := garray.New()
		for k, v := range m.X出栈多个(1) {
			t.AssertIN(k, g.Slice{1, 2, 3})
			t.AssertIN(v, g.Slice{"v1", "v2", "v3"})
			kArray.Append(k)
			vArray.Append(v)
		}
		t.Assert(m.X取数量(), 2)
		for k, v := range m.X出栈多个(2) {
			t.AssertIN(k, g.Slice{1, 2, 3})
			t.AssertIN(v, g.Slice{"v1", "v2", "v3"})
			kArray.Append(k)
			vArray.Append(v)
		}
		t.Assert(m.X取数量(), 0)

		t.Assert(kArray.Unique().Len(), 3)
		t.Assert(vArray.Unique().Len(), 3)

		v := m.X出栈多个(1)
		t.AssertNil(v)
		v = m.X出栈多个(-1)
		t.AssertNil(v)
	})
}

func TestIntStrMap_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Map  *map类.IntStrMap
	}
	// JSON
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(map[string]interface{}{
			"name": "john",
			"map":  []byte(`{"1":"v1","2":"v2"}`),
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Map.X取数量(), 2)
		t.Assert(v.Map.X取值(1), "v1")
		t.Assert(v.Map.X取值(2), "v2")
	})
	// Map
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(map[string]interface{}{
			"name": "john",
			"map": g.MapIntAny{
				1: "v1",
				2: "v2",
			},
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Map.X取数量(), 2)
		t.Assert(v.Map.X取值(1), "v1")
		t.Assert(v.Map.X取值(2), "v2")
	})
}

func TestIntStrMap_Replace(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntStr并从Map(g.MapIntStr{
			1: "v1",
			2: "v2",
			3: "v3",
		})

		t.Assert(m.X取值(1), "v1")
		t.Assert(m.X取值(2), "v2")
		t.Assert(m.X取值(3), "v3")

		m.X替换(g.MapIntStr{
			1: "v2",
			2: "v3",
			3: "v1",
		})

		t.Assert(m.X取值(1), "v2")
		t.Assert(m.X取值(2), "v3")
		t.Assert(m.X取值(3), "v1")
	})
}

func TestIntStrMap_String(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntStr并从Map(g.MapIntStr{
			1: "v1",
			2: "v2",
			3: "v3",
		})
		t.Assert(m.String(), "{\"1\":\"v1\",\"2\":\"v2\",\"3\":\"v3\"}")

		m = nil
		t.Assert(len(m.String()), 0)
	})
}

func Test_IntStrMap_DeepCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntStr并从Map(g.MapIntStr{
			1: "val1",
			2: "val2",
		})
		t.Assert(m.X取数量(), 2)

		n := m.DeepCopy().(*map类.IntStrMap)
		n.X设置值(1, "v1")
		t.AssertNE(m.X取值(1), n.X取值(1))
	})
}

func Test_IntStrMap_IsSubOf(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := map类.X创建IntStr并从Map(g.MapIntStr{
			1: "v1",
			2: "v2",
		})
		m2 := map类.X创建IntStr并从Map(g.MapIntStr{
			2: "v2",
		})
		t.Assert(m1.X是否为子集(m2), false)
		t.Assert(m2.X是否为子集(m1), true)
		t.Assert(m2.X是否为子集(m2), true)
	})
}

func Test_IntStrMap_Diff(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := map类.X创建IntStr并从Map(g.MapIntStr{
			0: "0",
			1: "1",
			2: "2",
			3: "3",
		})
		m2 := map类.X创建IntStr并从Map(g.MapIntStr{
			0: "0",
			2: "2",
			3: "31",
			4: "4",
		})
		addedKeys, removedKeys, updatedKeys := m1.X比较(m2)
		t.Assert(addedKeys, []int{4})
		t.Assert(removedKeys, []int{1})
		t.Assert(updatedKeys, []int{3})
	})
}
