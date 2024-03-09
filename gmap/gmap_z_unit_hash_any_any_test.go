// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

package map类_test

import (
	"testing"
	"time"
	
	"github.com/gogf/gf/v2/container/garray"
	"github.com/888go/goframe/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/888go/goframe/gmap/internal/json"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

func Test_AnyAnyMap_Var(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var m map类.AnyAnyMap
		m.X设置值(1, 1)

		t.Assert(m.X取值(1), 1)
		t.Assert(m.X取数量(), 1)
		t.Assert(m.X是否为空(), false)

		t.Assert(m.X取值或设置值(2, "2"), "2")
		t.Assert(m.X设置值并跳过已存在(2, "2"), false)

		t.Assert(m.X设置值并跳过已存在(3, 3), true)

		t.Assert(m.X删除(2), "2")
		t.Assert(m.X是否存在(2), false)

		t.AssertIN(3, m.X取所有名称())
		t.AssertIN(1, m.X取所有名称())
		t.AssertIN(3, m.X取所有值())
		t.AssertIN(1, m.X取所有值())
		m.X名称值交换()
		t.Assert(m.X取Map(), map[interface{}]int{1: 1, 3: 3})

		m.X清空()
		t.Assert(m.X取数量(), 0)
		t.Assert(m.X是否为空(), true)
	})
}

func Test_AnyAnyMap_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建AnyAny()
		m.X设置值(1, 1)

		t.Assert(m.X取值(1), 1)
		t.Assert(m.X取数量(), 1)
		t.Assert(m.X是否为空(), false)

		t.Assert(m.X取值或设置值(2, "2"), "2")
		t.Assert(m.X设置值并跳过已存在(2, "2"), false)

		t.Assert(m.X设置值并跳过已存在(3, 3), true)

		t.Assert(m.X删除(2), "2")
		t.Assert(m.X是否存在(2), false)

		t.AssertIN(3, m.X取所有名称())
		t.AssertIN(1, m.X取所有名称())
		t.AssertIN(3, m.X取所有值())
		t.AssertIN(1, m.X取所有值())
		m.X名称值交换()
		t.Assert(m.X取Map(), map[interface{}]int{1: 1, 3: 3})

		m.X清空()
		t.Assert(m.X取数量(), 0)
		t.Assert(m.X是否为空(), true)

		m2 := map类.X创建AnyAny并从Map(map[interface{}]interface{}{1: 1, 2: "2"})
		t.Assert(m2.X取Map(), map[interface{}]interface{}{1: 1, 2: "2"})
	})
}

func Test_AnyAnyMap_Set_Fun(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建AnyAny()

		m.X取值或设置值_函数(1, getAny)
		m.X取值或设置值_函数带锁(2, getAny)
		t.Assert(m.X取值(1), 123)
		t.Assert(m.X取值(2), 123)

		t.Assert(m.X设置值并跳过已存在_函数(1, getAny), false)
		t.Assert(m.X设置值并跳过已存在_函数(3, getAny), true)

		t.Assert(m.X设置值并跳过已存在_函数带锁(2, getAny), false)
		t.Assert(m.X设置值并跳过已存在_函数带锁(4, getAny), true)
	})

}

func Test_AnyAnyMap_Batch(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建AnyAny()

		m.X设置值Map(map[interface{}]interface{}{1: 1, 2: "2", 3: 3})
		t.Assert(m.X取Map(), map[interface{}]interface{}{1: 1, 2: "2", 3: 3})
		m.X删除多个值([]interface{}{1, 2})
		t.Assert(m.X取Map(), map[interface{}]interface{}{3: 3})
	})
}

func Test_AnyAnyMap_Iterator(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := map[interface{}]interface{}{1: 1, 2: "2"}
		m := map类.X创建AnyAny并从Map(expect)
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
		t.Assert(i, "2")
		t.Assert(j, 1)
	})
}

func Test_AnyAnyMap_Lock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := map[interface{}]interface{}{1: 1, 2: "2"}
		m := map类.X创建AnyAny并从Map(expect)
		m.X遍历写锁定(func(m map[interface{}]interface{}) {
			t.Assert(m, expect)
		})
		m.X遍历读锁定(func(m map[interface{}]interface{}) {
			t.Assert(m, expect)
		})
	})
}

func Test_AnyAnyMap_Clone(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// clone 方法是深克隆
		m := map类.X创建AnyAny并从Map(map[interface{}]interface{}{1: 1, 2: "2"})

		m_clone := m.X取副本()
		m.X删除(1)
		// 修改原 map,clone 后的 map 不影响
		t.AssertIN(1, m_clone.X取所有名称())

		m_clone.X删除(2)
		// 修改clone map,原 map 不影响
		t.AssertIN(2, m.X取所有名称())
	})
}

func Test_AnyAnyMap_Merge(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := map类.X创建AnyAny()
		m2 := map类.X创建AnyAny()
		m1.X设置值(1, 1)
		m2.X设置值(2, "2")
		m1.X合并(m2)
		t.Assert(m1.X取Map(), map[interface{}]interface{}{1: 1, 2: "2"})
		m3 := map类.X创建AnyAny并从Map(nil)
		m3.X合并(m2)
		t.Assert(m3.X取Map(), m2.X取Map())
	})
}

func Test_AnyAnyMap_Map(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建AnyAny()
		m.X设置值(1, 0)
		m.X设置值(2, 2)
		t.Assert(m.X取值(1), 0)
		t.Assert(m.X取值(2), 2)
		data := m.X取Map()
		t.Assert(data[1], 0)
		t.Assert(data[2], 2)
		data[3] = 3
		t.Assert(m.X取值(3), 3)
		m.X设置值(4, 4)
		t.Assert(data[4], 4)
	})
}

func Test_AnyAnyMap_MapCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建AnyAny()
		m.X设置值(1, 0)
		m.X设置值(2, 2)
		t.Assert(m.X取值(1), 0)
		t.Assert(m.X取值(2), 2)
		data := m.X浅拷贝()
		t.Assert(data[1], 0)
		t.Assert(data[2], 2)
		data[3] = 3
		t.Assert(m.X取值(3), nil)
		m.X设置值(4, 4)
		t.Assert(data[4], nil)
	})
}

func Test_AnyAnyMap_FilterEmpty(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建AnyAny()
		m.X设置值(1, 0)
		m.X设置值(2, 2)
		t.Assert(m.X取值(1), 0)
		t.Assert(m.X取值(2), 2)
		m.X删除所有空值()
		t.Assert(m.X取值(1), nil)
		t.Assert(m.X取值(2), 2)
	})
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建AnyAny()
		m.X设置值(1, 0)
		m.X设置值("time1", time.Time{})
		m.X设置值("time2", time.Now())
		t.Assert(m.X取值(1), 0)
		t.Assert(m.X取值("time1"), time.Time{})
		m.X删除所有空值()
		t.Assert(m.X取值(1), nil)
		t.Assert(m.X取值("time1"), nil)
		t.AssertNE(m.X取值("time2"), nil)
	})
}

func Test_AnyAnyMap_Json(t *testing.T) {
	// Marshal
	gtest.C(t, func(t *gtest.T) {
		data := g.MapAnyAny{
			"k1": "v1",
			"k2": "v2",
		}
		m1 := map类.X创建AnyAny并从Map(data)
		b1, err1 := json.Marshal(m1)
		b2, err2 := json.Marshal(gconv.Map(data))
		t.Assert(err1, err2)
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

		m := map类.X创建()
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

		var m map类.Map
		err = json.UnmarshalUseNumber(b, &m)
		t.AssertNil(err)
		t.Assert(m.X取值("k1"), data["k1"])
		t.Assert(m.X取值("k2"), data["k2"])
	})
}

func Test_AnyAnyMap_Pop(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建AnyAny并从Map(g.MapAnyAny{
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

func Test_AnyAnyMap_Pops(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建AnyAny并从Map(g.MapAnyAny{
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

func TestAnyAnyMap_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Map  *map类.Map
	}
	// JSON
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(map[string]interface{}{
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
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(map[string]interface{}{
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

func Test_AnyAnyMap_DeepCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建AnyAny并从Map(g.MapAnyAny{
			"k1": "v1",
			"k2": "v2",
		})
		t.Assert(m.X取数量(), 2)

		n := m.DeepCopy().(*map类.AnyAnyMap)
		n.X设置值("k1", "val1")
		t.AssertNE(m.X取值("k1"), n.X取值("k1"))
	})
}

func Test_AnyAnyMap_IsSubOf(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := map类.X创建AnyAny并从Map(g.MapAnyAny{
			"k1": "v1",
			"k2": "v2",
		})
		m2 := map类.X创建AnyAny并从Map(g.MapAnyAny{
			"k2": "v2",
		})
		t.Assert(m1.X是否为子集(m2), false)
		t.Assert(m2.X是否为子集(m1), true)
		t.Assert(m2.X是否为子集(m2), true)
	})
}

func Test_AnyAnyMap_Diff(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := map类.X创建AnyAny并从Map(g.MapAnyAny{
			"0": "v0",
			"1": "v1",
			2:   "v2",
			3:   3,
		})
		m2 := map类.X创建AnyAny并从Map(g.MapAnyAny{
			"0": "v0",
			2:   "v2",
			3:   "v3",
			4:   "v4",
		})
		addedKeys, removedKeys, updatedKeys := m1.X比较(m2)
		t.Assert(addedKeys, []interface{}{4})
		t.Assert(removedKeys, []interface{}{"1"})
		t.Assert(updatedKeys, []interface{}{3})
	})
}
