// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

package map类_test

import (
	"fmt"
	
	"github.com/888go/goframe/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/888go/goframe/gmap/internal/json"
	"github.com/gogf/gf/v2/util/gconv"
)

func ExampleListMap_Iterator() {
	m := map类.X创建链表mp()
	for i := 0; i < 10; i++ {
		m.X设置值(i, i*2)
	}

	var totalKey, totalValue int
	m.X遍历(func(k interface{}, v interface{}) bool {
		totalKey += k.(int)
		totalValue += v.(int)

		return totalKey < 10
	})

	fmt.Println("totalKey:", totalKey)
	fmt.Println("totalValue:", totalValue)

	// Output:
	// totalKey: 10
	// totalValue: 20
}

func ExampleListMap_IteratorAsc() {
	m := map类.X创建链表mp()
	for i := 0; i < 10; i++ {
		m.X设置值(i, i*2)
	}

	var totalKey, totalValue int
	m.X遍历升序(func(k interface{}, v interface{}) bool {
		totalKey += k.(int)
		totalValue += v.(int)

		return totalKey < 10
	})

	fmt.Println("totalKey:", totalKey)
	fmt.Println("totalValue:", totalValue)

	// Output:
	// totalKey: 10
	// totalValue: 20
}

func ExampleListMap_IteratorDesc() {
	m := map类.X创建链表mp()
	for i := 0; i < 10; i++ {
		m.X设置值(i, i*2)
	}

	var totalKey, totalValue int
	m.X遍历降序(func(k interface{}, v interface{}) bool {
		totalKey += k.(int)
		totalValue += v.(int)

		return totalKey < 10
	})

	fmt.Println("totalKey:", totalKey)
	fmt.Println("totalValue:", totalValue)

	// Output:
	// totalKey: 17
	// totalValue: 34
}

func ExampleListMap_Clone() {
	m := map类.X创建链表mp()

	m.X设置值("key1", "val1")
	fmt.Println(m)

	n := m.X取副本()
	fmt.Println(n)

	// Output:
	// {"key1":"val1"}
	// {"key1":"val1"}
}

func ExampleListMap_Clear() {
	var m map类.ListMap
	m.X设置值Map(g.MapAnyAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	m.X清空()

	fmt.Println(m.X取Map())

	// Output:
	// map[]
}

func ExampleListMap_Replace() {
	var m map类.ListMap
	m.X设置值Map(g.MapAnyAny{
		"k1": "v1",
	})

	var n map类.ListMap
	n.X设置值Map(g.MapAnyAny{
		"k2": "v2",
	})

	fmt.Println(m.X取Map())

	m.X替换(n.X取Map())
	fmt.Println(m.X取Map())

	// Output:
	// map[k1:v1]
	// map[k2:v2]
}

func ExampleListMap_Map() {
	m1 := map类.X创建链表mp()
	m1.X设置值("key1", "val1")
	fmt.Println("m1:", m1)

	n1 := m1.X取Map()
	fmt.Println("before n1:", n1)
	m1.X设置值("key1", "val2")
	fmt.Println("after n1:", n1)

	// Output:
	// m1: {"key1":"val1"}
	// before n1: map[key1:val1]
	// after n1: map[key1:val1]
}

func ExampleListMap_MapStrAny() {
	m := map类.X创建链表mp()
	m.X设置值("key1", "val1")
	m.X设置值("key2", "val2")

	n := m.X取MapStrAny()
	fmt.Printf("%#v", n)

	// Output:
	// map[string]interface {}{"key1":"val1", "key2":"val2"}
}

func ExampleListMap_FilterEmpty() {
	m := map类.X创建链表Map并从Map(g.MapAnyAny{
		"k1": "",
		"k2": nil,
		"k3": 0,
		"k4": 1,
	})
	m.X删除所有空值()
	fmt.Println(m.X取Map())

	// Output:
	// map[k4:1]
}

func ExampleListMap_Set() {
	m := map类.X创建链表mp()

	m.X设置值("key1", "val1")
	fmt.Println(m)

	// Output:
	// {"key1":"val1"}
}

func ExampleListMap_Sets() {
	m := map类.X创建链表mp()

	addMap := make(map[interface{}]interface{})
	addMap["key1"] = "val1"
	addMap["key2"] = "val2"
	addMap["key3"] = "val3"

	m.X设置值Map(addMap)
	fmt.Println(m)

	// May Output:
	// {"key1":"val1","key2":"val2","key3":"val3"}
}

func ExampleListMap_Search() {
	m := map类.X创建链表mp()

	m.X设置值("key1", "val1")

	value, found := m.X查找("key1")
	if found {
		fmt.Println("find key1 value:", value)
	}

	value, found = m.X查找("key2")
	if !found {
		fmt.Println("key2 not find")
	}

	// Output:
	// find key1 value: val1
	// key2 not find
}

func ExampleListMap_Get() {
	m := map类.X创建链表mp()

	m.X设置值("key1", "val1")

	fmt.Println("key1 value:", m.X取值("key1"))
	fmt.Println("key2 value:", m.X取值("key2"))

	// Output:
	// key1 value: val1
	// key2 value: <nil>
}

func ExampleListMap_Pop() {
	var m map类.ListMap
	m.X设置值Map(g.MapAnyAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	fmt.Println(m.X出栈())

	// May Output:
	// k1 v1
}

func ExampleListMap_Pops() {
	var m map类.ListMap
	m.X设置值Map(g.MapAnyAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})
	fmt.Println(m.X出栈多个(-1))
	fmt.Println("size:", m.X取数量())

	m.X设置值Map(g.MapAnyAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})
	fmt.Println(m.X出栈多个(2))
	fmt.Println("size:", m.X取数量())

	// May Output:
	// map[k1:v1 k2:v2 k3:v3 k4:v4]
	// size: 0
	// map[k1:v1 k2:v2]
	// size: 2
}

func ExampleListMap_GetOrSet() {
	m := map类.X创建链表mp()
	m.X设置值("key1", "val1")

	fmt.Println(m.X取值或设置值("key1", "NotExistValue"))
	fmt.Println(m.X取值或设置值("key2", "val2"))

	// Output:
	// val1
	// val2
}

func ExampleListMap_GetOrSetFunc() {
	m := map类.X创建链表mp()
	m.X设置值("key1", "val1")

	fmt.Println(m.X取值或设置值_函数("key1", func() interface{} {
		return "NotExistValue"
	}))
	fmt.Println(m.X取值或设置值_函数("key2", func() interface{} {
		return "NotExistValue"
	}))

	// Output:
	// val1
	// NotExistValue
}

func ExampleListMap_GetOrSetFuncLock() {
	m := map类.X创建链表mp()
	m.X设置值("key1", "val1")

	fmt.Println(m.X取值或设置值_函数带锁("key1", func() interface{} {
		return "NotExistValue"
	}))
	fmt.Println(m.X取值或设置值_函数带锁("key2", func() interface{} {
		return "NotExistValue"
	}))

	// Output:
	// val1
	// NotExistValue
}

func ExampleListMap_GetVar() {
	m := map类.X创建链表mp()
	m.X设置值("key1", "val1")

	fmt.Println(m.X取值泛型类("key1"))
	fmt.Println(m.X取值泛型类("key2").IsNil())

	// Output:
	// val1
	// true
}

func ExampleListMap_GetVarOrSet() {
	m := map类.X创建链表mp()
	m.X设置值("key1", "val1")

	fmt.Println(m.X取值或设置值泛型类("key1", "NotExistValue"))
	fmt.Println(m.X取值或设置值泛型类("key2", "val2"))

	// Output:
	// val1
	// val2
}

func ExampleListMap_GetVarOrSetFunc() {
	m := map类.X创建链表mp()
	m.X设置值("key1", "val1")

	fmt.Println(m.X取值或设置值泛型类_函数("key1", func() interface{} {
		return "NotExistValue"
	}))
	fmt.Println(m.X取值或设置值泛型类_函数("key2", func() interface{} {
		return "NotExistValue"
	}))

	// Output:
	// val1
	// NotExistValue
}

func ExampleListMap_GetVarOrSetFuncLock() {
	m := map类.X创建链表mp()
	m.X设置值("key1", "val1")

	fmt.Println(m.X取值或设置值泛型类_函数带锁("key1", func() interface{} {
		return "NotExistValue"
	}))
	fmt.Println(m.X取值或设置值泛型类_函数带锁("key2", func() interface{} {
		return "NotExistValue"
	}))

	// Output:
	// val1
	// NotExistValue
}

func ExampleListMap_SetIfNotExist() {
	var m map类.ListMap
	fmt.Println(m.X设置值并跳过已存在("k1", "v1"))
	fmt.Println(m.X设置值并跳过已存在("k1", "v2"))
	fmt.Println(m.X取Map())

	// Output:
	// true
	// false
	// map[k1:v1]
}

func ExampleListMap_SetIfNotExistFunc() {
	var m map类.ListMap
	fmt.Println(m.X设置值并跳过已存在_函数("k1", func() interface{} {
		return "v1"
	}))
	fmt.Println(m.X设置值并跳过已存在_函数("k1", func() interface{} {
		return "v2"
	}))
	fmt.Println(m.X取Map())

	// Output:
	// true
	// false
	// map[k1:v1]
}

func ExampleListMap_SetIfNotExistFuncLock() {
	var m map类.ListMap
	fmt.Println(m.X设置值并跳过已存在_函数带锁("k1", func() interface{} {
		return "v1"
	}))
	fmt.Println(m.X设置值并跳过已存在_函数带锁("k1", func() interface{} {
		return "v2"
	}))
	fmt.Println(m.X取Map())

	// Output:
	// true
	// false
	// map[k1:v1]
}

func ExampleListMap_Remove() {
	var m map类.ListMap
	m.X设置值("k1", "v1")

	fmt.Println(m.X删除("k1"))
	fmt.Println(m.X删除("k2"))
	fmt.Println(m.X取数量())

	// Output:
	// v1
	// <nil>
	// 0
}

func ExampleListMap_Removes() {
	var m map类.ListMap
	m.X设置值Map(g.MapAnyAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	removeList := make([]interface{}, 2)
	removeList = append(removeList, "k1")
	removeList = append(removeList, "k2")

	m.X删除多个值(removeList)

	fmt.Println(m.X取Map())

	// Output:
	// map[k3:v3 k4:v4]
}

func ExampleListMap_Keys() {
	var m map类.ListMap
	m.X设置值Map(g.MapAnyAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})
	fmt.Println(m.X取所有名称())

	// May Output:
	// [k1 k2 k3 k4]
}

func ExampleListMap_Values() {
	var m map类.ListMap
	m.X设置值Map(g.MapAnyAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})
	fmt.Println(m.X取所有值())

	// May Output:
	// [v1 v2 v3 v4]
}

func ExampleListMap_Contains() {
	var m map类.ListMap
	m.X设置值Map(g.MapAnyAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	fmt.Println(m.X是否存在("k1"))
	fmt.Println(m.X是否存在("k5"))

	// Output:
	// true
	// false
}

func ExampleListMap_Size() {
	var m map类.ListMap
	m.X设置值Map(g.MapAnyAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	fmt.Println(m.X取数量())

	// Output:
	// 4
}

func ExampleListMap_IsEmpty() {
	var m map类.ListMap
	fmt.Println(m.X是否为空())

	m.X设置值("k1", "v1")
	fmt.Println(m.X是否为空())

	// Output:
	// true
	// false
}

func ExampleListMap_Flip() {
	var m map类.ListMap
	m.X设置值Map(g.MapAnyAny{
		"k1": "v1",
	})
	m.X名称值交换()
	fmt.Println(m.X取Map())

	// Output:
	// map[v1:k1]
}

func ExampleListMap_Merge() {
	var m1, m2 map类.ListMap
	m1.X设置值("key1", "val1")
	m2.X设置值("key2", "val2")
	m1.X合并(&m2)
	fmt.Println(m1.X取Map())

	// May Output:
	// map[key1:val1 key2:val2]
}

func ExampleListMap_String() {
	var m map类.ListMap
	m.X设置值Map(g.MapAnyAny{
		"k1": "v1",
	})

	fmt.Println(m.String())

	// Output:
	// {"k1":"v1"}
}

func ExampleListMap_MarshalJSON() {
	var m map类.ListMap
	m.X设置值("k1", "v1")
	m.X设置值("k2", "v2")
	m.X设置值("k3", "v3")
	m.X设置值("k4", "v4")

	bytes, err := json.Marshal(&m)
	if err == nil {
		fmt.Println(gconv.String(bytes))
	}

	// Output:
	// {"k1":"v1","k2":"v2","k3":"v3","k4":"v4"}
}

func ExampleListMap_UnmarshalJSON() {
	var m map类.ListMap
	m.X设置值Map(g.MapAnyAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	var n map类.ListMap

	err := json.Unmarshal(gconv.Bytes(m.String()), &n)
	if err == nil {
		fmt.Println(n.X取Map())
	}

	// Output:
	// map[k1:v1 k2:v2 k3:v3 k4:v4]
}

func ExampleListMap_UnmarshalValue() {
	type User struct {
		Uid   int
		Name  string
		Pass1 string `gconv:"password1"`
		Pass2 string `gconv:"password2"`
	}

	var (
		m    map类.AnyAnyMap
		user = User{
			Uid:   1,
			Name:  "john",
			Pass1: "123",
			Pass2: "456",
		}
	)
	if err := gconv.Scan(user, &m); err == nil {
		fmt.Printf("%#v", m.X取Map())
	}

	// Output:
	// map[interface {}]interface {}{"Name":"john", "Uid":1, "password1":"123", "password2":"456"}
}
