// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

package map类_test

import (
	"testing"
	
	"github.com/888go/goframe/gmap"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gutil"
)

func getValue() interface{} {
	return 3
}

func Test_Map_Var(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var m map类.Map
		m.X设置值(1, 11)
		t.Assert(m.X取值(1), 11)
	})
	gtest.C(t, func(t *gtest.T) {
		var m map类.IntAnyMap
		m.X设置值(1, 11)
		t.Assert(m.X取值(1), 11)
	})
	gtest.C(t, func(t *gtest.T) {
		var m map类.IntIntMap
		m.X设置值(1, 11)
		t.Assert(m.X取值(1), 11)
	})
	gtest.C(t, func(t *gtest.T) {
		var m map类.IntStrMap
		m.X设置值(1, "11")
		t.Assert(m.X取值(1), "11")
	})
	gtest.C(t, func(t *gtest.T) {
		var m map类.StrAnyMap
		m.X设置值("1", "11")
		t.Assert(m.X取值("1"), "11")
	})
	gtest.C(t, func(t *gtest.T) {
		var m map类.StrStrMap
		m.X设置值("1", "11")
		t.Assert(m.X取值("1"), "11")
	})
	gtest.C(t, func(t *gtest.T) {
		var m map类.StrIntMap
		m.X设置值("1", 11)
		t.Assert(m.X取值("1"), 11)
	})
	gtest.C(t, func(t *gtest.T) {
		var m map类.ListMap
		m.X设置值("1", 11)
		t.Assert(m.X取值("1"), 11)
	})
	gtest.C(t, func(t *gtest.T) {
		var m map类.TreeMap
		m.SetComparator(gutil.ComparatorString)
		m.Set("1", 11)
		t.Assert(m.Get("1"), 11)
	})
}

func Test_Map_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建()
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

		m2 := map类.X创建并从Map(map[interface{}]interface{}{1: 1, "key1": "val1"})
		t.Assert(m2.X取Map(), map[interface{}]interface{}{1: 1, "key1": "val1"})
	})
}

func Test_Map_Set_Fun(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建()
		m.X取值或设置值_函数("fun", getValue)
		m.X取值或设置值_函数带锁("funlock", getValue)
		t.Assert(m.X取值("funlock"), 3)
		t.Assert(m.X取值("fun"), 3)
		m.X取值或设置值_函数("fun", getValue)
		t.Assert(m.X设置值并跳过已存在_函数("fun", getValue), false)
		t.Assert(m.X设置值并跳过已存在_函数带锁("funlock", getValue), false)
	})
}

func Test_Map_Batch(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := map类.X创建()
		m.X设置值Map(map[interface{}]interface{}{1: 1, "key1": "val1", "key2": "val2", "key3": "val3"})
		t.Assert(m.X取Map(), map[interface{}]interface{}{1: 1, "key1": "val1", "key2": "val2", "key3": "val3"})
		m.X删除多个值([]interface{}{"key1", 1})
		t.Assert(m.X取Map(), map[interface{}]interface{}{"key2": "val2", "key3": "val3"})
	})
}

func Test_Map_Iterator(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := map[interface{}]interface{}{1: 1, "key1": "val1"}

		m := map类.X创建并从Map(expect)
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

func Test_Map_Lock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		expect := map[interface{}]interface{}{1: 1, "key1": "val1"}
		m := map类.X创建并从Map(expect)
		m.X遍历写锁定(func(m map[interface{}]interface{}) {
			t.Assert(m, expect)
		})
		m.X遍历读锁定(func(m map[interface{}]interface{}) {
			t.Assert(m, expect)
		})
	})
}

func Test_Map_Clone(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// clone 方法是深克隆
		m := map类.X创建并从Map(map[interface{}]interface{}{1: 1, "key1": "val1"})
		m_clone := m.X取副本()
		m.X删除(1)
		// 修改原 map,clone 后的 map 不影响
		t.AssertIN(1, m_clone.X取所有名称())

		m_clone.X删除("key1")
		// 修改clone map,原 map 不影响
		t.AssertIN("key1", m.X取所有名称())
	})
}

func Test_Map_Basic_Merge(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := map类.X创建()
		m2 := map类.X创建()
		m1.X设置值("key1", "val1")
		m2.X设置值("key2", "val2")
		m1.X合并(m2)
		t.Assert(m1.X取Map(), map[interface{}]interface{}{"key1": "val1", "key2": "val2"})
	})
}
