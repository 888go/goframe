// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

package map类_test

import (
	"testing"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
)

func Test_TreeMap_Var(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var m map类.TreeMap
		m.SetComparator(工具类.X比较文本)
		m.X设置值("key1", "val1")
		t.Assert(m.Keys(), []interface{}{"key1"})

		t.Assert(m.Get("key1"), "val1")
		t.Assert(m.Size(), 1)
		t.Assert(m.IsEmpty(), false)

		t.Assert(m.GetOrSet("key2", "val2"), "val2")
		t.Assert(m.SetIfNotExist("key2", "val2"), false)

		t.Assert(m.SetIfNotExist("key3", "val3"), true)

		t.Assert(m.Remove("key2"), "val2")
		t.Assert(m.Contains("key2"), false)

		t.AssertIN("key3", m.Keys())
		t.AssertIN("key1", m.Keys())
		t.AssertIN("val3", m.Values())
		t.AssertIN("val1", m.Values())

		m.Flip()
		t.Assert(m.Map(), map[interface{}]interface{}{"val3": "key3", "val1": "key1"})

		m.Clear()
		t.Assert(m.Size(), 0)
		t.Assert(m.IsEmpty(), true)
	})
}

func Test_TreeMap_Basic(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建红黑树Map(工具类.X比较文本)
		m.X设置值("key1", "val1")
		t.Assert(m.Keys(), []interface{}{"key1"})

		t.Assert(m.Get("key1"), "val1")
		t.Assert(m.Size(), 1)
		t.Assert(m.IsEmpty(), false)

		t.Assert(m.GetOrSet("key2", "val2"), "val2")
		t.Assert(m.SetIfNotExist("key2", "val2"), false)

		t.Assert(m.SetIfNotExist("key3", "val3"), true)

		t.Assert(m.Remove("key2"), "val2")
		t.Assert(m.Contains("key2"), false)

		t.AssertIN("key3", m.Keys())
		t.AssertIN("key1", m.Keys())
		t.AssertIN("val3", m.Values())
		t.AssertIN("val1", m.Values())

		m.Flip()
		t.Assert(m.Map(), map[interface{}]interface{}{"val3": "key3", "val1": "key1"})

		m.Clear()
		t.Assert(m.Size(), 0)
		t.Assert(m.IsEmpty(), true)

		m2 := map类.X创建红黑树Map并从Map(工具类.X比较文本, map[interface{}]interface{}{1: 1, "key1": "val1"})
		t.Assert(m2.Map(), map[interface{}]interface{}{1: 1, "key1": "val1"})
	})
}

func Test_TreeMap_Set_Fun(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建红黑树Map(工具类.X比较文本)
		m.GetOrSetFunc("fun", getValue)
		m.GetOrSetFuncLock("funlock", getValue)
		t.Assert(m.Get("funlock"), 3)
		t.Assert(m.Get("fun"), 3)
		m.GetOrSetFunc("fun", getValue)
		t.Assert(m.SetIfNotExistFunc("fun", getValue), false)
		t.Assert(m.SetIfNotExistFuncLock("funlock", getValue), false)
	})
}

func Test_TreeMap_Batch(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := map类.X创建红黑树Map(工具类.X比较文本)
		m.Sets(map[interface{}]interface{}{1: 1, "key1": "val1", "key2": "val2", "key3": "val3"})
		t.Assert(m.Map(), map[interface{}]interface{}{1: 1, "key1": "val1", "key2": "val2", "key3": "val3"})
		m.Removes([]interface{}{"key1", 1})
		t.Assert(m.Map(), map[interface{}]interface{}{"key2": "val2", "key3": "val3"})
	})
}

func Test_TreeMap_Iterator(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		expect := map[interface{}]interface{}{1: 1, "key1": "val1"}
		m := map类.X创建红黑树Map并从Map(工具类.X比较文本, expect)
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

	单元测试类.C(t, func(t *单元测试类.T) {
		expect := map[interface{}]interface{}{1: 1, "key1": "val1"}
		m := map类.X创建红黑树Map并从Map(工具类.X比较文本, expect)
		for i := 0; i < 10; i++ {
			m.IteratorAsc(func(k interface{}, v interface{}) bool {
				t.Assert(expect[k], v)
				return true
			})
		}
		j := 0
		for i := 0; i < 10; i++ {
			m.IteratorAsc(func(k interface{}, v interface{}) bool {
				j++
				return false
			})
		}
		t.Assert(j, 10)
	})
}

func Test_TreeMap_Clone(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		// clone 方法是深克隆
		m := map类.X创建红黑树Map并从Map(工具类.X比较文本, map[interface{}]interface{}{1: 1, "key1": "val1"})
		m_clone := m.Clone()
		m.Remove(1)
		// 修改原 map,clone 后的 map 不影响
		t.AssertIN(1, m_clone.Keys())

		m_clone.Remove("key1")
		// 修改clone map,原 map 不影响
		t.AssertIN("key1", m.Keys())
	})
}

func Test_TreeMap_Json(t *testing.T) {
	// Marshal
	单元测试类.C(t, func(t *单元测试类.T) {
		data := g.MapAnyAny{
			"k1": "v1",
			"k2": "v2",
		}
		m1 := map类.X创建红黑树Map并从Map(工具类.X比较文本, data)
		b1, err1 := json.Marshal(m1)
		b2, err2 := json.Marshal(转换类.X取Map(data))
		t.Assert(err1, err2)
		t.Assert(b1, b2)
	})
	// Unmarshal
	单元测试类.C(t, func(t *单元测试类.T) {
		data := g.MapAnyAny{
			"k1": "v1",
			"k2": "v2",
		}
		b, err := json.Marshal(转换类.X取Map(data))
		t.AssertNil(err)

		m := map类.X创建红黑树Map(工具类.X比较文本)
		err = json.UnmarshalUseNumber(b, m)
		t.AssertNil(err)
		t.Assert(m.Get("k1"), data["k1"])
		t.Assert(m.Get("k2"), data["k2"])
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		data := g.MapAnyAny{
			"k1": "v1",
			"k2": "v2",
		}
		b, err := json.Marshal(转换类.X取Map(data))
		t.AssertNil(err)

		var m map类.TreeMap
		err = json.UnmarshalUseNumber(b, &m)
		t.AssertNil(err)
		t.Assert(m.Get("k1"), data["k1"])
		t.Assert(m.Get("k2"), data["k2"])
	})
}

func TestTreeMap_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Map  *map类.TreeMap
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
		t.Assert(v.Map.Size(), 2)
		t.Assert(v.Map.Get("k1"), "v1")
		t.Assert(v.Map.Get("k2"), "v2")
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
		t.Assert(v.Map.Size(), 2)
		t.Assert(v.Map.Get("k1"), "v1")
		t.Assert(v.Map.Get("k2"), "v2")
	})
}
