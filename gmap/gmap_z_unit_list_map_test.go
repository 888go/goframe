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

func Test_ListMap_Var(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var m map类.ListMap
		m.X设置值("key1", "val1")
		t.Assert(m.X取所有名称(), []interface{}{"key1"})

		t.Assert(m.X取值("key1"), "val1")
		t.Assert(m.X取数量(), 1)
		t.Assert(m.X是否为空(), false)

		t.Assert(m.X取值或设置值("key2", "val2"), "val2")
		t.Assert(m.X设置值并跳过已存在("key2", "val2"), false)

		t.Assert(m.X设置值并跳过已存在("key3", "val3"), true)
		t.Assert(m.X删除("key2"), "val2")
		t.Assert(m.X是否存在("key2"), false)

		t.AssertIN("key3", m.X取所有名称())
		t.AssertIN("key1", m.X取所有名称())
		t.AssertIN("val3", m.X取所有值())
		t.AssertIN("val1", m.X取所有值())

		m.X名称值交换()

		t.Assert(m.X取Map(), map[interface{}]interface{}{"val3": "key3", "val1": "key1"})

		m.X清空()
		t.Assert(m.X取数量(), 0)
		t.Assert(m.X是否为空(), true)
	})
}

func Test_ListMap_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建链表mp()
		m.X设置值("key1", "val1")
		t.Assert(m.X取所有名称(), []interface{}{"key1"})

		t.Assert(m.X取值("key1"), "val1")
		t.Assert(m.X取数量(), 1)
		t.Assert(m.X是否为空(), false)

		t.Assert(m.X取值或设置值("key2", "val2"), "val2")
		t.Assert(m.X设置值并跳过已存在("key2", "val2"), false)

		t.Assert(m.X设置值并跳过已存在("key3", "val3"), true)
		t.Assert(m.X删除("key2"), "val2")
		t.Assert(m.X是否存在("key2"), false)

		t.AssertIN("key3", m.X取所有名称())
		t.AssertIN("key1", m.X取所有名称())
		t.AssertIN("val3", m.X取所有值())
		t.AssertIN("val1", m.X取所有值())

		m.X名称值交换()

		t.Assert(m.X取Map(), map[interface{}]interface{}{"val3": "key3", "val1": "key1"})

		m.X清空()
		t.Assert(m.X取数量(), 0)
		t.Assert(m.X是否为空(), true)

		m2 := map类.X创建链表Map并从Map(map[interface{}]interface{}{1: 1, "key1": "val1"})
		t.Assert(m2.X取Map(), map[interface{}]interface{}{1: 1, "key1": "val1"})
	})
}

func Test_ListMap_Set_Fun(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建链表mp()
		m.X取值或设置值_函数("fun", getValue)
		m.X取值或设置值_函数带锁("funlock", getValue)
		t.Assert(m.X取值("funlock"), 3)
		t.Assert(m.X取值("fun"), 3)
		m.X取值或设置值_函数("fun", getValue)
		t.Assert(m.X设置值并跳过已存在_函数("fun", getValue), false)
		t.Assert(m.X设置值并跳过已存在_函数带锁("funlock", getValue), false)
	})
}

func Test_ListMap_Batch(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建链表mp()
		m.X设置值Map(map[interface{}]interface{}{1: 1, "key1": "val1", "key2": "val2", "key3": "val3"})
		t.Assert(m.X取Map(), map[interface{}]interface{}{1: 1, "key1": "val1", "key2": "val2", "key3": "val3"})
		m.X删除多个值([]interface{}{"key1", 1})
		t.Assert(m.X取Map(), map[interface{}]interface{}{"key2": "val2", "key3": "val3"})
	})
}

func Test_ListMap_Iterator(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := map[interface{}]interface{}{1: 1, "key1": "val1"}

		m := map类.X创建链表Map并从Map(expect)
		m.X遍历(func(k interface{}, v interface{}) bool {
			t.Assert(expect[k], v)
			return true
		})
		// 断言返回值对遍历控制
		i := 0
		j := 0
		m.X遍历(func(k interface{}, v interface{}) bool {
			i++
			return true
		})
		m.X遍历(func(k interface{}, v interface{}) bool {
			j++
			return false
		})
		t.Assert(i, 2)
		t.Assert(j, 1)
	})
}

func Test_ListMap_Clone(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// clone 方法是深克隆
		m := map类.X创建链表Map并从Map(map[interface{}]interface{}{1: 1, "key1": "val1"})
		m_clone := m.X取副本()
		m.X删除(1)
		// 修改原 map,clone 后的 map 不影响
		t.AssertIN(1, m_clone.X取所有名称())

		m_clone.X删除("key1")
		// 修改clone map,原 map 不影响
		t.AssertIN("key1", m.X取所有名称())
	})
}

func Test_ListMap_Basic_Merge(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := map类.X创建链表mp()
		m2 := map类.X创建链表mp()
		m1.X设置值("key1", "val1")
		m2.X设置值("key2", "val2")
		m1.X合并(m2)
		t.Assert(m1.X取Map(), map[interface{}]interface{}{"key1": "val1", "key2": "val2"})
		m3 := map类.X创建链表Map并从Map(nil)
		m3.X合并(m2)
		t.Assert(m3.X取Map(), m2.X取Map())
	})
}

func Test_ListMap_Order(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建链表mp()
		m.X设置值("k1", "v1")
		m.X设置值("k2", "v2")
		m.X设置值("k3", "v3")
		t.Assert(m.X取所有名称(), g.Slice{"k1", "k2", "k3"})
		t.Assert(m.X取所有值(), g.Slice{"v1", "v2", "v3"})
	})
}

func Test_ListMap_FilterEmpty(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建链表mp()
		m.X设置值(1, "")
		m.X设置值(2, "2")
		t.Assert(m.X取数量(), 2)
		t.Assert(m.X取值(2), "2")
		m.X删除所有空值()
		t.Assert(m.X取数量(), 1)
		t.Assert(m.X取值(2), "2")
	})
}

func Test_ListMap_Json(t *testing.T) {
	// Marshal
	gtest.C(t, func(t *gtest.T) {
		data := g.MapAnyAny{
			"k1": "v1",
		}
		m1 := map类.X创建链表Map并从Map(data)
		b1, err1 := json.Marshal(m1)
		t.AssertNil(err1)
		b2, err2 := json.Marshal(gconv.Map(data))
		t.AssertNil(err2)
		t.Assert(b1, b2)
	})
	// Unmarshal
	gtest.C(t, func(t *gtest.T) {
		data := g.MapAnyAny{
			"k1": "v1",
			"k2": "v2",
		}
		b, err := json.Marshal(gconv.Map(data))
		t.AssertNil(err)

		m := map类.X创建链表mp()
		err = json.UnmarshalUseNumber(b, m)
		t.AssertNil(err)
		t.Assert(m.X取值("k1"), data["k1"])
		t.Assert(m.X取值("k2"), data["k2"])
	})

	gtest.C(t, func(t *gtest.T) {
		data := g.MapAnyAny{
			"k1": "v1",
			"k2": "v2",
		}
		b, err := json.Marshal(gconv.Map(data))
		t.AssertNil(err)

		var m map类.ListMap
		err = json.UnmarshalUseNumber(b, &m)
		t.AssertNil(err)
		t.Assert(m.X取值("k1"), data["k1"])
		t.Assert(m.X取值("k2"), data["k2"])
	})
}

func Test_ListMap_Json_Sequence(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建链表mp()
		for i := 'z'; i >= 'a'; i-- {
			m.X设置值(string(i), i)
		}
		b, err := json.Marshal(m)
		t.AssertNil(err)
		t.Assert(b, `{"z":122,"y":121,"x":120,"w":119,"v":118,"u":117,"t":116,"s":115,"r":114,"q":113,"p":112,"o":111,"n":110,"m":109,"l":108,"k":107,"j":106,"i":105,"h":104,"g":103,"f":102,"e":101,"d":100,"c":99,"b":98,"a":97}`)
	})
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建链表mp()
		for i := 'a'; i <= 'z'; i++ {
			m.X设置值(string(i), i)
		}
		b, err := json.Marshal(m)
		t.AssertNil(err)
		t.Assert(b, `{"a":97,"b":98,"c":99,"d":100,"e":101,"f":102,"g":103,"h":104,"i":105,"j":106,"k":107,"l":108,"m":109,"n":110,"o":111,"p":112,"q":113,"r":114,"s":115,"t":116,"u":117,"v":118,"w":119,"x":120,"y":121,"z":122}`)
	})
}

func Test_ListMap_Pop(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建链表Map并从Map(g.MapAnyAny{
			"k1": "v1",
			"k2": "v2",
		})
		t.Assert(m.X取数量(), 2)

		k1, v1 := m.X出栈()
		t.AssertIN(k1, g.Slice{"k1", "k2"})
		t.AssertIN(v1, g.Slice{"v1", "v2"})
		t.Assert(m.X取数量(), 1)
		k2, v2 := m.X出栈()
		t.AssertIN(k2, g.Slice{"k1", "k2"})
		t.AssertIN(v2, g.Slice{"v1", "v2"})
		t.Assert(m.X取数量(), 0)

		t.AssertNE(k1, k2)
		t.AssertNE(v1, v2)

		k3, v3 := m.X出栈()
		t.AssertNil(k3)
		t.AssertNil(v3)
	})
}

func Test_ListMap_Pops(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建链表Map并从Map(g.MapAnyAny{
			"k1": "v1",
			"k2": "v2",
			"k3": "v3",
		})
		t.Assert(m.X取数量(), 3)

		kArray := garray.New()
		vArray := garray.New()
		for k, v := range m.X出栈多个(1) {
			t.AssertIN(k, g.Slice{"k1", "k2", "k3"})
			t.AssertIN(v, g.Slice{"v1", "v2", "v3"})
			kArray.Append(k)
			vArray.Append(v)
		}
		t.Assert(m.X取数量(), 2)
		for k, v := range m.X出栈多个(2) {
			t.AssertIN(k, g.Slice{"k1", "k2", "k3"})
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

func TestListMap_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Map  *map类.ListMap
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
		t.Assert(v.Map.X取值("1"), "v1")
		t.Assert(v.Map.X取值("2"), "v2")
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
		t.Assert(v.Map.X取值("1"), "v1")
		t.Assert(v.Map.X取值("2"), "v2")
	})
}

func TestListMap_String(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建链表mp()
		m.X设置值(1, "")
		m.X设置值(2, "2")
		t.Assert(m.String(), "{\"1\":\"\",\"2\":\"2\"}")

		m1 := map类.X创建链表Map并从Map(nil)
		t.Assert(m1.String(), "{}")
	})
}

func TestListMap_MarshalJSON(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建链表mp()
		m.X设置值(1, "")
		m.X设置值(2, "2")
		res, err := m.MarshalJSON()
		t.Assert(res, []byte("{\"1\":\"\",\"2\":\"2\"}"))
		t.AssertNil(err)

		m1 := map类.X创建链表Map并从Map(nil)
		res, err = m1.MarshalJSON()
		t.Assert(res, []byte("{}"))
		t.AssertNil(err)
	})
}

func TestListMap_DeepCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建链表mp()
		m.X设置值(1, "1")
		m.X设置值(2, "2")
		t.Assert(m.X取数量(), 2)

		n := m.DeepCopy().(*map类.ListMap)
		n.X设置值(1, "val1")
		t.AssertNE(m.X取值(1), n.X取值(1))
	})
}
