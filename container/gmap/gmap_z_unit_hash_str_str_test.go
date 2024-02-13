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

func Test_StrStrMap_Var(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var m map类.StrStrMap
		m.X设置值("a", "a")

		t.Assert(m.X取值("a"), "a")
		t.Assert(m.X取数量(), 1)
		t.Assert(m.X是否为空(), false)

		t.Assert(m.X取值或设置值("b", "b"), "b")
		t.Assert(m.X设置值并跳过已存在("b", "b"), false)

		t.Assert(m.X设置值并跳过已存在("c", "c"), true)

		t.Assert(m.X删除("b"), "b")
		t.Assert(m.X是否存在("b"), false)

		t.AssertIN("c", m.X取所有名称())
		t.AssertIN("a", m.X取所有名称())
		t.AssertIN("a", m.X取所有值())
		t.AssertIN("c", m.X取所有值())

		m.X名称值交换()

		t.Assert(m.X取Map(), map[string]string{"a": "a", "c": "c"})

		m.X清空()
		t.Assert(m.X取数量(), 0)
		t.Assert(m.X是否为空(), true)
	})
}

func Test_StrStrMap_Basic(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建StrStr()
		m.X设置值("a", "a")

		t.Assert(m.X取值("a"), "a")
		t.Assert(m.X取数量(), 1)
		t.Assert(m.X是否为空(), false)

		t.Assert(m.X取值或设置值("b", "b"), "b")
		t.Assert(m.X设置值并跳过已存在("b", "b"), false)

		t.Assert(m.X设置值并跳过已存在("c", "c"), true)

		t.Assert(m.X删除("b"), "b")
		t.Assert(m.X是否存在("b"), false)

		t.AssertIN("c", m.X取所有名称())
		t.AssertIN("a", m.X取所有名称())
		t.AssertIN("a", m.X取所有值())
		t.AssertIN("c", m.X取所有值())

		m.X名称值交换()

		t.Assert(m.X取Map(), map[string]string{"a": "a", "c": "c"})

		m.X清空()
		t.Assert(m.X取数量(), 0)
		t.Assert(m.X是否为空(), true)

		m2 := map类.X创建StrStr并从Map(map[string]string{"a": "a", "b": "b"})
		t.Assert(m2.X取Map(), map[string]string{"a": "a", "b": "b"})
	})
}

func Test_StrStrMap_Set_Fun(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建StrStr()

		m.X取值或设置值_函数("a", getStr)
		m.X取值或设置值_函数带锁("b", getStr)
		t.Assert(m.X取值("a"), "z")
		t.Assert(m.X取值("b"), "z")
		t.Assert(m.X设置值并跳过已存在_函数("a", getStr), false)
		t.Assert(m.X设置值并跳过已存在_函数("c", getStr), true)

		t.Assert(m.X设置值并跳过已存在_函数带锁("b", getStr), false)
		t.Assert(m.X设置值并跳过已存在_函数带锁("d", getStr), true)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建StrStr并从Map(nil)

		t.Assert(m.X取值或设置值_函数带锁("b", getStr), "z")
	})
}

func Test_StrStrMap_Batch(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建StrStr()

		m.X设置值Map(map[string]string{"a": "a", "b": "b", "c": "c"})
		t.Assert(m.X取Map(), map[string]string{"a": "a", "b": "b", "c": "c"})
		m.X删除多个值([]string{"a", "b"})
		t.Assert(m.X取Map(), map[string]string{"c": "c"})
	})
}

func Test_StrStrMap_Iterator(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		expect := map[string]string{"a": "a", "b": "b"}
		m := map类.X创建StrStr并从Map(expect)
		m.X遍历(func(k string, v string) bool {
			t.Assert(expect[k], v)
			return true
		})
		// 断言返回值对遍历控制
		i := 0
		j := 0
		m.X遍历(func(k string, v string) bool {
			i++
			return true
		})
		m.X遍历(func(k string, v string) bool {
			j++
			return false
		})
		t.Assert(i, 2)
		t.Assert(j, 1)
	})
}

func Test_StrStrMap_Lock(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		expect := map[string]string{"a": "a", "b": "b"}

		m := map类.X创建StrStr并从Map(expect)
		m.X遍历写锁定(func(m map[string]string) {
			t.Assert(m, expect)
		})
		m.X遍历读锁定(func(m map[string]string) {
			t.Assert(m, expect)
		})
	})
}

func Test_StrStrMap_Clone(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		// clone 方法是深克隆
		m := map类.X创建StrStr并从Map(map[string]string{"a": "a", "b": "b", "c": "c"})

		m_clone := m.X取副本()
		m.X删除("a")
		// 修改原 map,clone 后的 map 不影响
		t.AssertIN("a", m_clone.X取所有名称())

		m_clone.X删除("b")
		// 修改clone map,原 map 不影响
		t.AssertIN("b", m.X取所有名称())
	})
}

func Test_StrStrMap_Merge(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m1 := map类.X创建StrStr()
		m2 := map类.X创建StrStr()
		m1.X设置值("a", "a")
		m2.X设置值("b", "b")
		m1.X合并(m2)
		t.Assert(m1.X取Map(), map[string]string{"a": "a", "b": "b"})
		m3 := map类.X创建StrStr并从Map(nil)
		m3.X合并(m2)
		t.Assert(m3.X取Map(), m2.X取Map())
	})
}

func Test_StrStrMap_Map(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建StrStr()
		m.X设置值("1", "1")
		m.X设置值("2", "2")
		t.Assert(m.X取值("1"), "1")
		t.Assert(m.X取值("2"), "2")
		data := m.X取Map()
		t.Assert(data["1"], "1")
		t.Assert(data["2"], "2")
		data["3"] = "3"
		t.Assert(m.X取值("3"), "3")
		m.X设置值("4", "4")
		t.Assert(data["4"], "4")
	})
}

func Test_StrStrMap_MapCopy(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建StrStr()
		m.X设置值("1", "1")
		m.X设置值("2", "2")
		t.Assert(m.X取值("1"), "1")
		t.Assert(m.X取值("2"), "2")
		data := m.X浅拷贝()
		t.Assert(data["1"], "1")
		t.Assert(data["2"], "2")
		data["3"] = "3"
		t.Assert(m.X取值("3"), "")
		m.X设置值("4", "4")
		t.Assert(data["4"], "")
	})
}

func Test_StrStrMap_FilterEmpty(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建StrStr()
		m.X设置值("1", "")
		m.X设置值("2", "2")
		t.Assert(m.X取数量(), 2)
		t.Assert(m.X取值("1"), "")
		t.Assert(m.X取值("2"), "2")
		m.X删除所有空值()
		t.Assert(m.X取数量(), 1)
		t.Assert(m.X取值("2"), "2")
	})
}

func Test_StrStrMap_Json(t *testing.T) {
	// Marshal
	单元测试类.C(t, func(t *单元测试类.T) {
		data := g.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		}
		m1 := map类.X创建StrStr并从Map(data)
		b1, err1 := json.Marshal(m1)
		b2, err2 := json.Marshal(data)
		t.Assert(err1, err2)
		t.Assert(b1, b2)
	})
	// Unmarshal
	单元测试类.C(t, func(t *单元测试类.T) {
		data := g.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		}
		b, err := json.Marshal(data)
		t.AssertNil(err)

		m := map类.X创建StrStr()
		err = json.UnmarshalUseNumber(b, m)
		t.AssertNil(err)
		t.Assert(m.X取值("k1"), data["k1"])
		t.Assert(m.X取值("k2"), data["k2"])
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		data := g.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		}
		b, err := json.Marshal(data)
		t.AssertNil(err)

		var m map类.StrStrMap
		err = json.UnmarshalUseNumber(b, &m)
		t.AssertNil(err)
		t.Assert(m.X取值("k1"), data["k1"])
		t.Assert(m.X取值("k2"), data["k2"])
	})
}

func Test_StrStrMap_Pop(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建StrStr并从Map(g.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		})
		t.Assert(m.X取数量(), 2)

		k1, v1 := m.X出栈()
		t.AssertIN(k1, g.Slice别名{"k1", "k2"})
		t.AssertIN(v1, g.Slice别名{"v1", "v2"})
		t.Assert(m.X取数量(), 1)
		k2, v2 := m.X出栈()
		t.AssertIN(k2, g.Slice别名{"k1", "k2"})
		t.AssertIN(v2, g.Slice别名{"v1", "v2"})
		t.Assert(m.X取数量(), 0)

		t.AssertNE(k1, k2)
		t.AssertNE(v1, v2)

		k3, v3 := m.X出栈()
		t.Assert(k3, "")
		t.Assert(v3, "")
	})
}

func Test_StrStrMap_Pops(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建StrStr并从Map(g.MapStrStr{
			"k1": "v1",
			"k2": "v2",
			"k3": "v3",
		})
		t.Assert(m.X取数量(), 3)

		kArray := 数组类.X创建()
		vArray := 数组类.X创建()
		for k, v := range m.X出栈多个(1) {
			t.AssertIN(k, g.Slice别名{"k1", "k2", "k3"})
			t.AssertIN(v, g.Slice别名{"v1", "v2", "v3"})
			kArray.Append别名(k)
			vArray.Append别名(v)
		}
		t.Assert(m.X取数量(), 2)
		for k, v := range m.X出栈多个(2) {
			t.AssertIN(k, g.Slice别名{"k1", "k2", "k3"})
			t.AssertIN(v, g.Slice别名{"v1", "v2", "v3"})
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

func TestStrStrMap_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Map  *map类.StrStrMap
	}
	// JSON
	单元测试类.C(t, func(t *单元测试类.T) {
		var v *V
		err := 转换类.Struct(map[string]interface{}{
			"name": "john",
			"map":  []byte(`{"k1":"v1","k2":"v2"}`),
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Map.X取数量(), 2)
		t.Assert(v.Map.X取值("k1"), "v1")
		t.Assert(v.Map.X取值("k2"), "v2")
	})
	// Map
	单元测试类.C(t, func(t *单元测试类.T) {
		var v *V
		err := 转换类.Struct(map[string]interface{}{
			"name": "john",
			"map": g.Map{
				"k1": "v1",
				"k2": "v2",
			},
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Map.X取数量(), 2)
		t.Assert(v.Map.X取值("k1"), "v1")
		t.Assert(v.Map.X取值("k2"), "v2")
	})
}

func Test_StrStrMap_DeepCopy(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建StrStr并从Map(g.MapStrStr{
			"key1": "val1",
			"key2": "val2",
		})
		t.Assert(m.X取数量(), 2)

		n := m.DeepCopy().(*map类.StrStrMap)
		n.X设置值("key1", "v1")
		t.AssertNE(m.X取值("key1"), n.X取值("key1"))
	})
}

func Test_StrStrMap_IsSubOf(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m1 := map类.X创建StrStr并从Map(g.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		})
		m2 := map类.X创建StrStr并从Map(g.MapStrStr{
			"k2": "v2",
		})
		t.Assert(m1.X是否为子集(m2), false)
		t.Assert(m2.X是否为子集(m1), true)
		t.Assert(m2.X是否为子集(m2), true)
	})
}

func Test_StrStrMap_Diff(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m1 := map类.X创建StrStr并从Map(g.MapStrStr{
			"0": "0",
			"1": "1",
			"2": "2",
			"3": "3",
		})
		m2 := map类.X创建StrStr并从Map(g.MapStrStr{
			"0": "0",
			"2": "2",
			"3": "31",
			"4": "4",
		})
		addedKeys, removedKeys, updatedKeys := m1.X比较(m2)
		t.Assert(addedKeys, []string{"4"})
		t.Assert(removedKeys, []string{"1"})
		t.Assert(updatedKeys, []string{"3"})
	})
}
