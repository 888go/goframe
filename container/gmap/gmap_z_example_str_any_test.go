// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

package map类_test

import (
	"fmt"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/util/gconv"
)

func ExampleStrAnyMap_Iterator() {
	m := map类.X创建StrAny()
	for i := 1; i <= 10; i++ {
		m.X设置值(转换类.String(i), i*2)
	}

	var totalValue int
	m.X遍历(func(k string, v interface{}) bool {
		totalValue += v.(int)

		return totalValue < 50
	})

	fmt.Println("totalValue:", totalValue)

	// May Output:
	// totalValue: 52
}

func ExampleStrAnyMap_Clone() {
	m := map类.X创建StrAny()

	m.X设置值("key1", "val1")
	fmt.Println(m)

	n := m.X取副本()
	fmt.Println(n)

	// Output:
	// {"key1":"val1"}
	// {"key1":"val1"}
}

func ExampleStrAnyMap_Map() {
	// 非并发安全，指向底层数据的指针
	m1 := map类.X创建StrAny()
	m1.X设置值("key1", "val1")
	fmt.Println("m1:", m1)

	n1 := m1.X取Map()
	fmt.Println("before n1:", n1)
	m1.X设置值("key1", "val2")
	fmt.Println("after n1:", n1)

	// 并发安全，复制底层数据
	m2 := map类.X创建StrAny(true)
	m2.X设置值("key1", "val1")
	fmt.Println("m2:", m2)

	n2 := m2.X取Map()
	fmt.Println("before n2:", n2)
	m2.X设置值("key1", "val2")
	fmt.Println("after n2:", n2)

	// Output:
	// m1: {"key1":"val1"}
	// before n1: map[key1:val1]
	// after n1: map[key1:val2]
	// m2: {"key1":"val1"}
	// before n2: map[key1:val1]
	// after n2: map[key1:val1]
}

func ExampleStrAnyMap_MapCopy() {
	m := map类.X创建StrAny()

	m.X设置值("key1", "val1")
	m.X设置值("key2", "val2")
	fmt.Println(m)

	n := m.X浅拷贝()
	fmt.Println(n)

	// Output:
	// {"key1":"val1","key2":"val2"}
	// map[key1:val1 key2:val2]
}

func ExampleStrAnyMap_MapStrAny() {
	m := map类.X创建StrAny()
	m.X设置值("key1", "val1")
	m.X设置值("key2", "val2")

	n := m.X取MapStrAny()
	fmt.Printf("%#v", n)

	// Output:
	// map[string]interface {}{"key1":"val1", "key2":"val2"}
}

func ExampleStrAnyMap_FilterEmpty() {
	m := map类.X创建AnyStr并从Map(g.MapStrAny{
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

func ExampleStrAnyMap_FilterNil() {
	m := map类.X创建AnyStr并从Map(g.MapStrAny{
		"k1": "",
		"k2": nil,
		"k3": 0,
		"k4": 1,
	})
	m.X删除所有nil值()
	fmt.Printf("%#v", m.X取Map())

	// Output:
	// map[string]interface {}{"k1":"", "k3":0, "k4":1}
}

func ExampleStrAnyMap_Set() {
	m := map类.X创建StrAny()

	m.X设置值("key1", "val1")
	fmt.Println(m)

	// Output:
	// {"key1":"val1"}
}

func ExampleStrAnyMap_Sets() {
	m := map类.X创建StrAny()

	addMap := make(map[string]interface{})
	addMap["key1"] = "val1"
	addMap["key2"] = "val2"
	addMap["key3"] = "val3"

	m.X设置值Map(addMap)
	fmt.Println(m)

	// Output:
	// {"key1":"val1","key2":"val2","key3":"val3"}
}

func ExampleStrAnyMap_Search() {
	m := map类.X创建StrAny()

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

func ExampleStrAnyMap_Get() {
	m := map类.X创建StrAny()

	m.X设置值("key1", "val1")

	fmt.Println("key1 value:", m.X取值("key1"))
	fmt.Println("key2 value:", m.X取值("key2"))

	// Output:
	// key1 value: val1
	// key2 value: <nil>
}

func ExampleStrAnyMap_Pop() {
	var m map类.StrAnyMap
	m.X设置值Map(g.MapStrAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	fmt.Println(m.X出栈())

	// May Output:
	// k1 v1
}

func ExampleStrAnyMap_Pops() {
	var m map类.StrAnyMap
	m.X设置值Map(g.MapStrAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})
	fmt.Println(m.X出栈多个(-1))
	fmt.Println("size:", m.X取数量())

	m.X设置值Map(g.MapStrAny{
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

func ExampleStrAnyMap_GetOrSet() {
	m := map类.X创建StrAny()
	m.X设置值("key1", "val1")

	fmt.Println(m.X取值或设置值("key1", "NotExistValue"))
	fmt.Println(m.X取值或设置值("key2", "val2"))

	// Output:
	// val1
	// val2
}

func ExampleStrAnyMap_GetOrSetFunc() {
	m := map类.X创建StrAny()
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

func ExampleStrAnyMap_GetOrSetFuncLock() {
	m := map类.X创建StrAny()
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

func ExampleStrAnyMap_GetVar() {
	m := map类.X创建StrAny()
	m.X设置值("key1", "val1")

	fmt.Println(m.X取值泛型类("key1"))
	fmt.Println(m.X取值泛型类("key2").X是否为Nil())

	// Output:
	// val1
	// true
}

func ExampleStrAnyMap_GetVarOrSet() {
	m := map类.X创建StrAny()
	m.X设置值("key1", "val1")

	fmt.Println(m.X取值或设置值泛型类("key1", "NotExistValue"))
	fmt.Println(m.X取值或设置值泛型类("key2", "val2"))

	// Output:
	// val1
	// val2
}

func ExampleStrAnyMap_GetVarOrSetFunc() {
	m := map类.X创建StrAny()
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

func ExampleStrAnyMap_GetVarOrSetFuncLock() {
	m := map类.X创建StrAny()
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

func ExampleStrAnyMap_SetIfNotExist() {
	var m map类.StrAnyMap
	fmt.Println(m.X设置值并跳过已存在("k1", "v1"))
	fmt.Println(m.X设置值并跳过已存在("k1", "v2"))
	fmt.Println(m.X取Map())

	// Output:
	// true
	// false
	// map[k1:v1]
}

func ExampleStrAnyMap_SetIfNotExistFunc() {
	var m map类.StrAnyMap
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

func ExampleStrAnyMap_SetIfNotExistFuncLock() {
	var m map类.StrAnyMap
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

func ExampleStrAnyMap_Remove() {
	var m map类.StrAnyMap
	m.X设置值("k1", "v1")

	fmt.Println(m.X删除("k1"))
	fmt.Println(m.X删除("k2"))
	fmt.Println(m.X取数量())

	// Output:
	// v1
	// <nil>
	// 0
}

func ExampleStrAnyMap_Removes() {
	var m map类.StrAnyMap
	m.X设置值Map(g.MapStrAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	removeList := make([]string, 2)
	removeList = append(removeList, "k1")
	removeList = append(removeList, "k2")

	m.X删除多个值(removeList)

	fmt.Println(m.X取Map())

	// Output:
	// map[k3:v3 k4:v4]
}

func ExampleStrAnyMap_Keys() {
	var m map类.StrAnyMap
	m.X设置值Map(g.MapStrAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})
	fmt.Println(m.X取所有名称())

	// May Output:
	// [k1 k2 k3 k4]
}

func ExampleStrAnyMap_Values() {
	var m map类.StrAnyMap
	m.X设置值Map(g.MapStrAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})
	fmt.Println(m.X取所有值())

	// May Output:
	// [v1 v2 v3 v4]
}

func ExampleStrAnyMap_Contains() {
	var m map类.StrAnyMap
	m.X设置值Map(g.MapStrAny{
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

func ExampleStrAnyMap_Size() {
	var m map类.StrAnyMap
	m.X设置值Map(g.MapStrAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	fmt.Println(m.X取数量())

	// Output:
	// 4
}

func ExampleStrAnyMap_IsEmpty() {
	var m map类.StrAnyMap
	fmt.Println(m.X是否为空())

	m.X设置值("k1", "v1")
	fmt.Println(m.X是否为空())

	// Output:
	// true
	// false
}

func ExampleStrAnyMap_Clear() {
	var m map类.StrAnyMap
	m.X设置值Map(g.MapStrAny{
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

func ExampleStrAnyMap_Replace() {
	var m map类.StrAnyMap
	m.X设置值Map(g.MapStrAny{
		"k1": "v1",
	})

	var n map类.StrAnyMap
	n.X设置值Map(g.MapStrAny{
		"k2": "v2",
	})

	fmt.Println(m.X取Map())

	m.X替换(n.X取Map())
	fmt.Println(m.X取Map())

	n.X设置值("k2", "v1")
	fmt.Println(m.X取Map())

	// Output:
	// map[k1:v1]
	// map[k2:v2]
	// map[k2:v1]
}

func ExampleStrAnyMap_LockFunc() {
	var m map类.StrAnyMap
	m.X设置值Map(g.MapStrAny{
		"k1": 1,
		"k2": 2,
		"k3": 3,
		"k4": 4,
	})

	m.X遍历写锁定(func(m map[string]interface{}) {
		totalValue := 0
		for _, v := range m {
			totalValue += v.(int)
		}
		fmt.Println("totalValue:", totalValue)
	})

	// Output:
	// totalValue: 10
}

func ExampleStrAnyMap_RLockFunc() {
	var m map类.StrAnyMap
	m.X设置值Map(g.MapStrAny{
		"k1": 1,
		"k2": 2,
		"k3": 3,
		"k4": 4,
	})

	m.X遍历读锁定(func(m map[string]interface{}) {
		totalValue := 0
		for _, v := range m {
			totalValue += v.(int)
		}
		fmt.Println("totalValue:", totalValue)
	})

	// Output:
	// totalValue: 10
}

func ExampleStrAnyMap_Flip() {
	var m map类.StrAnyMap
	m.X设置值Map(g.MapStrAny{
		"k1": "v1",
	})
	m.X名称值交换()
	fmt.Println(m.X取Map())

	// Output:
	// map[v1:k1]
}

func ExampleStrAnyMap_Merge() {
	var m1, m2 map类.StrAnyMap
	m1.X设置值("key1", "val1")
	m2.X设置值("key2", "val2")
	m1.X合并(&m2)
	fmt.Println(m1.X取Map())

	// May Output:
	// map[key1:val1 key2:val2]
}

func ExampleStrAnyMap_String() {
	var m map类.StrAnyMap
	m.X设置值Map(g.MapStrAny{
		"k1": "v1",
	})

	fmt.Println(m.String())

	var m1 *map类.StrAnyMap = nil
	fmt.Println(len(m1.String()))

	// Output:
	// {"k1":"v1"}
	// 0
}

func ExampleStrAnyMap_MarshalJSON() {
	var m map类.StrAnyMap
	m.X设置值Map(g.MapStrAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	bytes, err := json.Marshal(&m)
	if err == nil {
		fmt.Println(转换类.String(bytes))
	}

	// Output:
	// {"k1":"v1","k2":"v2","k3":"v3","k4":"v4"}
}

func ExampleStrAnyMap_UnmarshalJSON() {
	var m map类.StrAnyMap
	m.X设置值Map(g.MapStrAny{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	var n map类.StrAnyMap

	err := json.Unmarshal(转换类.X取字节集(m.String()), &n)
	if err == nil {
		fmt.Println(n.X取Map())
	}

	// Output:
	// map[k1:v1 k2:v2 k3:v3 k4:v4]
}

func ExampleStrAnyMap_UnmarshalValue() {
	var m map类.StrAnyMap

	goWeb := map[string]interface{}{
		"goframe": "https://goframe.org",
		"gin":     "https://gin-gonic.com/",
		"echo":    "https://echo.labstack.com/",
	}

	if err := 转换类.Scan(goWeb, &m); err == nil {
		fmt.Printf("%#v", m.X取Map())
	}
	// Output:
	// map[string]interface {}{"echo":"https://echo.labstack.com/", "gin":"https://gin-gonic.com/", "goframe":"https://goframe.org"}
}
