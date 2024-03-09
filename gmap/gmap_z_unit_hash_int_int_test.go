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

func getInt() int {
	return 123
}

func intIntCallBack(int, int) bool {
	return true
}

func Test_IntIntMap_Var(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var m map类.IntIntMap
		m.X设置值(1, 1)

		t.Assert(m.X取值(1), 1)
		t.Assert(m.X取数量(), 1)
		t.Assert(m.X是否为空(), false)

		t.Assert(m.X取值或设置值(2, 2), 2)
		t.Assert(m.X设置值并跳过已存在(2, 2), false)

		t.Assert(m.X设置值并跳过已存在(3, 3), true)

		t.Assert(m.X删除(2), 2)
		t.Assert(m.X是否存在(2), false)

		t.AssertIN(3, m.X取所有名称())
		t.AssertIN(1, m.X取所有名称())
		t.AssertIN(3, m.X取所有值())
		t.AssertIN(1, m.X取所有值())
		m.X名称值交换()
		t.Assert(m.X取Map(), map[int]int{1: 1, 3: 3})

		m.X清空()
		t.Assert(m.X取数量(), 0)
		t.Assert(m.X是否为空(), true)
	})
}

func Test_IntIntMap_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntInt()
		m.X设置值(1, 1)

		t.Assert(m.X取值(1), 1)
		t.Assert(m.X取数量(), 1)
		t.Assert(m.X是否为空(), false)

		t.Assert(m.X取值或设置值(2, 2), 2)
		t.Assert(m.X设置值并跳过已存在(2, 2), false)

		t.Assert(m.X设置值并跳过已存在(3, 3), true)

		t.Assert(m.X删除(2), 2)
		t.Assert(m.X是否存在(2), false)

		t.AssertIN(3, m.X取所有名称())
		t.AssertIN(1, m.X取所有名称())
		t.AssertIN(3, m.X取所有值())
		t.AssertIN(1, m.X取所有值())
		m.X名称值交换()
		t.Assert(m.X取Map(), map[int]int{1: 1, 3: 3})

		m.X清空()
		t.Assert(m.X取数量(), 0)
		t.Assert(m.X是否为空(), true)

		m2 := map类.X创建IntInt并从Map(map[int]int{1: 1, 2: 2})
		t.Assert(m2.X取Map(), map[int]int{1: 1, 2: 2})
	})

	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntInt(true)
		m.X设置值(1, 1)
		t.Assert(m.X取Map(), map[int]int{1: 1})
	})
}

func Test_IntIntMap_Set_Fun(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntInt()

		m.X取值或设置值_函数(1, getInt)
		m.X取值或设置值_函数带锁(2, getInt)
		t.Assert(m.X取值(1), 123)
		t.Assert(m.X取值(2), 123)
		t.Assert(m.X设置值并跳过已存在_函数(1, getInt), false)
		t.Assert(m.X设置值并跳过已存在_函数(3, getInt), true)

		t.Assert(m.X设置值并跳过已存在_函数带锁(2, getInt), false)
		t.Assert(m.X设置值并跳过已存在_函数带锁(4, getInt), true)
	})

	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntInt并从Map(nil)
		t.Assert(m.X取值或设置值_函数带锁(1, getInt), getInt())
	})
}

func Test_IntIntMap_Batch(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntInt()

		m.X设置值Map(map[int]int{1: 1, 2: 2, 3: 3})
		m.X遍历(intIntCallBack)
		t.Assert(m.X取Map(), map[int]int{1: 1, 2: 2, 3: 3})
		m.X删除多个值([]int{1, 2})
		t.Assert(m.X取Map(), map[int]int{3: 3})
	})
}

func Test_IntIntMap_Iterator(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := map[int]int{1: 1, 2: 2}
		m := map类.X创建IntInt并从Map(expect)
		m.X遍历(func(k int, v int) bool {
			t.Assert(expect[k], v)
			return true
		})
		// 断言返回值对遍历控制
		i := 0
		j := 0
		m.X遍历(func(k int, v int) bool {
			i++
			return true
		})
		m.X遍历(func(k int, v int) bool {
			j++
			return false
		})
		t.Assert(i, 2)
		t.Assert(j, 1)
	})
}

func Test_IntIntMap_Lock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := map[int]int{1: 1, 2: 2}
		m := map类.X创建IntInt并从Map(expect)
		m.X遍历写锁定(func(m map[int]int) {
			t.Assert(m, expect)
		})
		m.X遍历读锁定(func(m map[int]int) {
			t.Assert(m, expect)
		})
	})
}

func Test_IntIntMap_Clone(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// clone 方法是深克隆
		m := map类.X创建IntInt并从Map(map[int]int{1: 1, 2: 2})

		m_clone := m.X取副本()
		m.X删除(1)
		// 修改原 map,clone 后的 map 不影响
		t.AssertIN(1, m_clone.X取所有名称())

		m_clone.X删除(2)
		// 修改clone map,原 map 不影响
		t.AssertIN(2, m.X取所有名称())
	})
}

func Test_IntIntMap_Merge(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := map类.X创建IntInt()
		m2 := map类.X创建IntInt()
		m1.X设置值(1, 1)
		m2.X设置值(2, 2)
		m1.X合并(m2)
		t.Assert(m1.X取Map(), map[int]int{1: 1, 2: 2})
		m3 := map类.X创建IntInt并从Map(nil)
		m3.X合并(m2)
		t.Assert(m3.X取Map(), m2.X取Map())
	})
}

func Test_IntIntMap_Map(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntInt()
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

func Test_IntIntMap_MapCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntInt()
		m.X设置值(1, 0)
		m.X设置值(2, 2)
		t.Assert(m.X取值(1), 0)
		t.Assert(m.X取值(2), 2)
		data := m.X浅拷贝()
		t.Assert(data[1], 0)
		t.Assert(data[2], 2)
		data[3] = 3
		t.Assert(m.X取值(3), 0)
		m.X设置值(4, 4)
		t.Assert(data[4], 0)
	})
}

func Test_IntIntMap_FilterEmpty(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntInt()
		m.X设置值(1, 0)
		m.X设置值(2, 2)
		t.Assert(m.X取数量(), 2)
		t.Assert(m.X取值(1), 0)
		t.Assert(m.X取值(2), 2)
		m.X删除所有空值()
		t.Assert(m.X取数量(), 1)
		t.Assert(m.X取值(2), 2)
	})
}

func Test_IntIntMap_Json(t *testing.T) {
	// Marshal
	gtest.C(t, func(t *gtest.T) {
		data := g.MapIntInt{
			1: 10,
			2: 20,
		}
		m1 := map类.X创建IntInt并从Map(data)
		b1, err1 := json.Marshal(m1)
		b2, err2 := json.Marshal(data)
		t.Assert(err1, err2)
		t.Assert(b1, b2)
	})
	// Unmarshal
	gtest.C(t, func(t *gtest.T) {
		data := g.MapIntInt{
			1: 10,
			2: 20,
		}
		b, err := json.Marshal(data)
		t.AssertNil(err)

		m := map类.X创建IntInt()
		err = json.UnmarshalUseNumber(b, m)
		t.AssertNil(err)
		t.Assert(m.X取值(1), data[1])
		t.Assert(m.X取值(2), data[2])
	})
}

func Test_IntIntMap_Pop(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntInt并从Map(g.MapIntInt{
			1: 11,
			2: 22,
		})
		t.Assert(m.X取数量(), 2)

		k1, v1 := m.X出栈()
		t.AssertIN(k1, g.Slice{1, 2})
		t.AssertIN(v1, g.Slice{11, 22})
		t.Assert(m.X取数量(), 1)
		k2, v2 := m.X出栈()
		t.AssertIN(k2, g.Slice{1, 2})
		t.AssertIN(v2, g.Slice{11, 22})
		t.Assert(m.X取数量(), 0)

		t.AssertNE(k1, k2)
		t.AssertNE(v1, v2)

		k3, v3 := m.X出栈()
		t.Assert(k3, 0)
		t.Assert(v3, 0)
	})
}

func Test_IntIntMap_Pops(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntInt并从Map(g.MapIntInt{
			1: 11,
			2: 22,
			3: 33,
		})
		t.Assert(m.X取数量(), 3)

		kArray := garray.New()
		vArray := garray.New()
		for k, v := range m.X出栈多个(1) {
			t.AssertIN(k, g.Slice{1, 2, 3})
			t.AssertIN(v, g.Slice{11, 22, 33})
			kArray.Append(k)
			vArray.Append(v)
		}
		t.Assert(m.X取数量(), 2)
		for k, v := range m.X出栈多个(2) {
			t.AssertIN(k, g.Slice{1, 2, 3})
			t.AssertIN(v, g.Slice{11, 22, 33})
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

func TestIntIntMap_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Map  *map类.IntIntMap
	}
	// JSON
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(map[string]interface{}{
			"name": "john",
			"map":  []byte(`{"1":1,"2":2}`),
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Map.X取数量(), 2)
		t.Assert(v.Map.X取值(1), "1")
		t.Assert(v.Map.X取值(2), "2")
	})
	// Map
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(map[string]interface{}{
			"name": "john",
			"map": g.MapIntAny{
				1: 1,
				2: 2,
			},
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Map.X取数量(), 2)
		t.Assert(v.Map.X取值(1), "1")
		t.Assert(v.Map.X取值(2), "2")
	})
}

func Test_IntIntMap_DeepCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建IntInt并从Map(g.MapIntInt{
			1: 1,
			2: 2,
		})
		t.Assert(m.X取数量(), 2)

		n := m.DeepCopy().(*map类.IntIntMap)
		n.X设置值(1, 2)
		t.AssertNE(m.X取值(1), n.X取值(1))
	})
}

func Test_IntIntMap_IsSubOf(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := map类.X创建IntAny并从Map(g.MapIntAny{
			1: 1,
			2: 2,
		})
		m2 := map类.X创建IntAny并从Map(g.MapIntAny{
			2: 2,
		})
		t.Assert(m1.X是否为子集(m2), false)
		t.Assert(m2.X是否为子集(m1), true)
		t.Assert(m2.X是否为子集(m2), true)
	})
}

func Test_IntIntMap_Diff(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := map类.X创建IntInt并从Map(g.MapIntInt{
			0: 0,
			1: 1,
			2: 2,
			3: 3,
		})
		m2 := map类.X创建IntInt并从Map(g.MapIntInt{
			0: 0,
			2: 2,
			3: 31,
			4: 4,
		})
		addedKeys, removedKeys, updatedKeys := m1.X比较(m2)
		t.Assert(addedKeys, []int{4})
		t.Assert(removedKeys, []int{1})
		t.Assert(updatedKeys, []int{3})
	})
}
