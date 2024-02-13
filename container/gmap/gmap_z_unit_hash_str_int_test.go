// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

package map类_test

import (
	"testing"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_StrIntMap_Var(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var m map类.StrIntMap
		m.X设置值("a", 1)

		t.Assert(m.X取值("a"), 1)
		t.Assert(m.X取数量(), 1)
		t.Assert(m.X是否为空(), false)

		t.Assert(m.X取值或设置值("b", 2), 2)
		t.Assert(m.X设置值并跳过已存在("b", 2), false)

		t.Assert(m.X设置值并跳过已存在("c", 3), true)

		t.Assert(m.X删除("b"), 2)
		t.Assert(m.X是否存在("b"), false)

		t.AssertIN("c", m.X取所有名称())
		t.AssertIN("a", m.X取所有名称())
		t.AssertIN(3, m.X取所有值())
		t.AssertIN(1, m.X取所有值())

		m_f := map类.X创建StrInt()
		m_f.X设置值("1", 2)
		m_f.X名称值交换()
		t.Assert(m_f.X取Map(), map[string]int{"2": 1})

		m.X清空()
		t.Assert(m.X取数量(), 0)
		t.Assert(m.X是否为空(), true)
	})
}

func Test_StrIntMap_Basic(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建StrInt()
		m.X设置值("a", 1)

		t.Assert(m.X取值("a"), 1)
		t.Assert(m.X取数量(), 1)
		t.Assert(m.X是否为空(), false)

		t.Assert(m.X取值或设置值("b", 2), 2)
		t.Assert(m.X设置值并跳过已存在("b", 2), false)

		t.Assert(m.X设置值并跳过已存在("c", 3), true)

		t.Assert(m.X删除("b"), 2)
		t.Assert(m.X是否存在("b"), false)

		t.AssertIN("c", m.X取所有名称())
		t.AssertIN("a", m.X取所有名称())
		t.AssertIN(3, m.X取所有值())
		t.AssertIN(1, m.X取所有值())

		m_f := map类.X创建StrInt()
		m_f.X设置值("1", 2)
		m_f.X名称值交换()
		t.Assert(m_f.X取Map(), map[string]int{"2": 1})

		m.X清空()
		t.Assert(m.X取数量(), 0)
		t.Assert(m.X是否为空(), true)

		m2 := map类.X创建StrInt并从Map(map[string]int{"a": 1, "b": 2})
		t.Assert(m2.X取Map(), map[string]int{"a": 1, "b": 2})
	})
}

func Test_StrIntMap_Set_Fun(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建StrInt()

		m.X取值或设置值_函数("a", getInt)
		m.X取值或设置值_函数带锁("b", getInt)
		t.Assert(m.X取值("a"), 123)
		t.Assert(m.X取值("b"), 123)
		t.Assert(m.X设置值并跳过已存在_函数("a", getInt), false)
		t.Assert(m.X设置值并跳过已存在_函数("c", getInt), true)

		t.Assert(m.X设置值并跳过已存在_函数带锁("b", getInt), false)
		t.Assert(m.X设置值并跳过已存在_函数带锁("d", getInt), true)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建StrInt并从Map(nil)
		t.Assert(m.X取值或设置值_函数带锁("a", getInt), 123)
	})
}

func Test_StrIntMap_Batch(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建StrInt()

		m.X设置值Map(map[string]int{"a": 1, "b": 2, "c": 3})
		t.Assert(m.X取Map(), map[string]int{"a": 1, "b": 2, "c": 3})
		m.X删除多个值([]string{"a", "b"})
		t.Assert(m.X取Map(), map[string]int{"c": 3})
	})
}

func Test_StrIntMap_Iterator(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		expect := map[string]int{"a": 1, "b": 2}
		m := map类.X创建StrInt并从Map(expect)
		m.X遍历(func(k string, v int) bool {
			t.Assert(expect[k], v)
			return true
		})
		// 断言返回值对遍历控制
		i := 0
		j := 0
		m.X遍历(func(k string, v int) bool {
			i++
			return true
		})
		m.X遍历(func(k string, v int) bool {
			j++
			return false
		})
		t.Assert(i, 2)
		t.Assert(j, 1)
	})
}

func Test_StrIntMap_Lock(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		expect := map[string]int{"a": 1, "b": 2}

		m := map类.X创建StrInt并从Map(expect)
		m.X遍历写锁定(func(m map[string]int) {
			t.Assert(m, expect)
		})
		m.X遍历读锁定(func(m map[string]int) {
			t.Assert(m, expect)
		})
	})
}

func Test_StrIntMap_Clone(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		// clone 方法是深克隆
		m := map类.X创建StrInt并从Map(map[string]int{"a": 1, "b": 2, "c": 3})

		m_clone := m.X取副本()
		m.X删除("a")
		// 修改原 map,clone 后的 map 不影响
		t.AssertIN("a", m_clone.X取所有名称())

		m_clone.X删除("b")
		// 修改clone map,原 map 不影响
		t.AssertIN("b", m.X取所有名称())
	})
}

func Test_StrIntMap_Merge(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m1 := map类.X创建StrInt()
		m2 := map类.X创建StrInt()
		m1.X设置值("a", 1)
		m2.X设置值("b", 2)
		m1.X合并(m2)
		t.Assert(m1.X取Map(), map[string]int{"a": 1, "b": 2})
		m3 := map类.X创建StrInt并从Map(nil)
		m3.X合并(m2)
		t.Assert(m3.X取Map(), m2.X取Map())
	})
}

func Test_StrIntMap_Map(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建StrInt()
		m.X设置值("1", 1)
		m.X设置值("2", 2)
		t.Assert(m.X取值("1"), 1)
		t.Assert(m.X取值("2"), 2)
		data := m.X取Map()
		t.Assert(data["1"], 1)
		t.Assert(data["2"], 2)
		data["3"] = 3
		t.Assert(m.X取值("3"), 3)
		m.X设置值("4", 4)
		t.Assert(data["4"], 4)
	})
}

func Test_StrIntMap_MapCopy(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建StrInt()
		m.X设置值("1", 1)
		m.X设置值("2", 2)
		t.Assert(m.X取值("1"), 1)
		t.Assert(m.X取值("2"), 2)
		data := m.X浅拷贝()
		t.Assert(data["1"], 1)
		t.Assert(data["2"], 2)
		data["3"] = 3
		t.Assert(m.X取值("3"), 0)
		m.X设置值("4", 4)
		t.Assert(data["4"], 0)
	})
}

func Test_StrIntMap_FilterEmpty(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建StrInt()
		m.X设置值("1", 0)
		m.X设置值("2", 2)
		t.Assert(m.X取数量(), 2)
		t.Assert(m.X取值("1"), 0)
		t.Assert(m.X取值("2"), 2)
		m.X删除所有空值()
		t.Assert(m.X取数量(), 1)
		t.Assert(m.X取值("2"), 2)
	})
}

func Test_StrIntMap_Json(t *testing.T) {
	// Marshal
	单元测试类.C(t, func(t *单元测试类.T) {
		data := g.MapStrInt{
			"k1": 1,
			"k2": 2,
		}
		m1 := map类.X创建StrInt并从Map(data)
		b1, err1 := json.Marshal(m1)
		b2, err2 := json.Marshal(data)
		t.Assert(err1, err2)
		t.Assert(b1, b2)
	})
	// Unmarshal
	单元测试类.C(t, func(t *单元测试类.T) {
		data := g.MapStrInt{
			"k1": 1,
			"k2": 2,
		}
		b, err := json.Marshal(data)
		t.AssertNil(err)

		m := map类.X创建StrInt()
		err = json.UnmarshalUseNumber(b, m)
		t.AssertNil(err)
		t.Assert(m.X取值("k1"), data["k1"])
		t.Assert(m.X取值("k2"), data["k2"])
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		data := g.MapStrInt{
			"k1": 1,
			"k2": 2,
		}
		b, err := json.Marshal(data)
		t.AssertNil(err)

		var m map类.StrIntMap
		err = json.UnmarshalUseNumber(b, &m)
		t.AssertNil(err)
		t.Assert(m.X取值("k1"), data["k1"])
		t.Assert(m.X取值("k2"), data["k2"])
	})
}

func Test_StrIntMap_Pop(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建StrInt并从Map(g.MapStrInt{
			"k1": 11,
			"k2": 22,
		})
		t.Assert(m.X取数量(), 2)

		k1, v1 := m.X出栈()
		t.AssertIN(k1, g.Slice别名{"k1", "k2"})
		t.AssertIN(v1, g.Slice别名{11, 22})
		t.Assert(m.X取数量(), 1)
		k2, v2 := m.X出栈()
		t.AssertIN(k2, g.Slice别名{"k1", "k2"})
		t.AssertIN(v2, g.Slice别名{11, 22})
		t.Assert(m.X取数量(), 0)

		t.AssertNE(k1, k2)
		t.AssertNE(v1, v2)

		k3, v3 := m.X出栈()
		t.Assert(k3, "")
		t.Assert(v3, 0)
	})
}

func Test_StrIntMap_Pops(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建StrInt并从Map(g.MapStrInt{
			"k1": 11,
			"k2": 22,
			"k3": 33,
		})
		t.Assert(m.X取数量(), 3)

		kArray := 数组类.X创建()
		vArray := 数组类.X创建()
		for k, v := range m.X出栈多个(1) {
			t.AssertIN(k, g.Slice别名{"k1", "k2", "k3"})
			t.AssertIN(v, g.Slice别名{11, 22, 33})
			kArray.Append别名(k)
			vArray.Append别名(v)
		}
		t.Assert(m.X取数量(), 2)
		for k, v := range m.X出栈多个(2) {
			t.AssertIN(k, g.Slice别名{"k1", "k2", "k3"})
			t.AssertIN(v, g.Slice别名{11, 22, 33})
			kArray.Append别名(k)
			vArray.Append别名(v)
		}
		t.Assert(m.X取数量(), 0)

		t.Assert(kArray.X去重().X取长度(), 3)
		t.Assert(vArray.X去重().X取长度(), 3)

		v := m.X出栈多个(1)
		t.AssertNil(v)
		v = m.X出栈多个(-1)
		t.AssertNil(v)
	})
}

func TestStrIntMap_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Map  *map类.StrIntMap
	}
	// JSON
	单元测试类.C(t, func(t *单元测试类.T) {
		var v *V
		err := 转换类.Struct(map[string]interface{}{
			"name": "john",
			"map":  []byte(`{"k1":1,"k2":2}`),
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Map.X取数量(), 2)
		t.Assert(v.Map.X取值("k1"), 1)
		t.Assert(v.Map.X取值("k2"), 2)
	})
	// Map
	单元测试类.C(t, func(t *单元测试类.T) {
		var v *V
		err := 转换类.Struct(map[string]interface{}{
			"name": "john",
			"map": g.Map{
				"k1": 1,
				"k2": 2,
			},
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Map.X取数量(), 2)
		t.Assert(v.Map.X取值("k1"), 1)
		t.Assert(v.Map.X取值("k2"), 2)
	})
}

func Test_StrIntMap_DeepCopy(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建StrInt并从Map(g.MapStrInt{
			"key1": 1,
			"key2": 2,
		})
		t.Assert(m.X取数量(), 2)

		n := m.DeepCopy().(*map类.StrIntMap)
		n.X设置值("key1", 2)
		t.AssertNE(m.X取值("key1"), n.X取值("key1"))
	})
}

func Test_StrIntMap_IsSubOf(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m1 := map类.X创建StrInt并从Map(g.MapStrInt{
			"k1": 1,
			"k2": 2,
		})
		m2 := map类.X创建StrInt并从Map(g.MapStrInt{
			"k2": 2,
		})
		t.Assert(m1.X是否为子集(m2), false)
		t.Assert(m2.X是否为子集(m1), true)
		t.Assert(m2.X是否为子集(m2), true)
	})
}

func Test_StrIntMap_Diff(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m1 := map类.X创建StrInt并从Map(g.MapStrInt{
			"0": 0,
			"1": 1,
			"2": 2,
			"3": 3,
		})
		m2 := map类.X创建StrInt并从Map(g.MapStrInt{
			"0": 0,
			"2": 2,
			"3": 31,
			"4": 4,
		})
		addedKeys, removedKeys, updatedKeys := m1.X比较(m2)
		t.Assert(addedKeys, []string{"4"})
		t.Assert(removedKeys, []string{"1"})
		t.Assert(updatedKeys, []string{"3"})
	})
}
