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

func ExampleIntAnyMap_Iterator() {
	m := map类.X创建IntAny()
	for i := 0; i < 10; i++ {
		m.X设置值(i, i*2)
	}

	var totalKey, totalValue int
	m.X遍历(func(k int, v interface{}) bool {
		totalKey += k
		totalValue += v.(int)

		return totalKey < 10
	})

	fmt.Println("totalKey:", totalKey)
	fmt.Println("totalValue:", totalValue)

	// May Output:
	// totalKey: 11
	// totalValue: 22
}

func ExampleIntAnyMap_Clone() {
	m := map类.X创建IntAny()

	m.X设置值(1, "val1")
	fmt.Println(m)

	n := m.X取副本()
	fmt.Println(n)

	// Output:
	// {"1":"val1"}
	// {"1":"val1"}
}

func ExampleIntAnyMap_Map() {
	// 非并发安全，指向底层数据的指针
	m1 := map类.X创建IntAny()
	m1.X设置值(1, "val1")
	fmt.Println("m1:", m1)

	n1 := m1.X取Map()
	fmt.Println("before n1:", n1)
	m1.X设置值(1, "val2")
	fmt.Println("after n1:", n1)

	// 并发安全，复制底层数据
	m2 := map类.X创建(true)
	m2.X设置值(1, "val1")
	fmt.Println("m2:", m2)

	n2 := m2.X取Map()
	fmt.Println("before n2:", n2)
	m2.X设置值(1, "val2")
	fmt.Println("after n2:", n2)

	// Output:
	// m1: {"1":"val1"}
	// before n1: map[1:val1]
	// after n1: map[1:val2]
	// m2: {"1":"val1"}
	// before n2: map[1:val1]
	// after n2: map[1:val1]
}

func ExampleIntAnyMap_MapCopy() {
	m := map类.X创建IntAny()

	m.X设置值(1, "val1")
	m.X设置值(2, "val2")
	fmt.Println(m)

	n := m.X浅拷贝()
	fmt.Println(n)

	// Output:
	// {"1":"val1","2":"val2"}
	// map[1:val1 2:val2]
}

func ExampleIntAnyMap_MapStrAny() {
	m := map类.X创建IntAny()
	m.X设置值(1001, "val1")
	m.X设置值(1002, "val2")

	n := m.X取MapStrAny()
	fmt.Printf("%#v", n)

	// Output:
	// map[string]interface {}{"1001":"val1", "1002":"val2"}
}

func ExampleIntAnyMap_FilterEmpty() {
	m := map类.X创建IntAny并从Map(g.MapIntAny{
		1: "",
		2: nil,
		3: 0,
		4: 1,
	})
	m.X删除所有空值()
	fmt.Println(m.X取Map())

	// Output:
	// map[4:1]
}

func ExampleIntAnyMap_FilterNil() {
	m := map类.X创建IntAny并从Map(g.MapIntAny{
		1: "",
		2: nil,
		3: 0,
		4: 1,
	})
	m.X删除所有nil值()
	fmt.Printf("%#v", m.X取Map())

	// Output:
	// map[int]interface {}{1:"", 3:0, 4:1}
}

func ExampleIntAnyMap_Set() {
	m := map类.X创建IntAny()

	m.X设置值(1, "val1")
	fmt.Println(m)

	// Output:
	// {"1":"val1"}
}

func ExampleIntAnyMap_Sets() {
	m := map类.X创建IntAny()

	addMap := make(map[int]interface{})
	addMap[1] = "val1"
	addMap[2] = "val2"
	addMap[3] = "val3"

	m.X设置值Map(addMap)
	fmt.Println(m)

	// Output:
	// {"1":"val1","2":"val2","3":"val3"}
}

func ExampleIntAnyMap_Search() {
	m := map类.X创建IntAny()

	m.X设置值(1, "val1")

	value, found := m.X查找(1)
	if found {
		fmt.Println("find key1 value:", value)
	}

	value, found = m.X查找(2)
	if !found {
		fmt.Println("key2 not find")
	}

	// Output:
	// find key1 value: val1
	// key2 not find
}

func ExampleIntAnyMap_Get() {
	m := map类.X创建IntAny()

	m.X设置值(1, "val1")

	fmt.Println("key1 value:", m.X取值(1))
	fmt.Println("key2 value:", m.X取值(2))

	// Output:
	// key1 value: val1
	// key2 value: <nil>
}

func ExampleIntAnyMap_Pop() {
	var m map类.IntAnyMap
	m.X设置值Map(g.MapIntAny{
		1: "v1",
		2: "v2",
		3: "v3",
		4: "v4",
	})

	fmt.Println(m.X出栈())

	// May Output:
	// 1 v1
}

func ExampleIntAnyMap_Pops() {
	var m map类.IntAnyMap
	m.X设置值Map(g.MapIntAny{
		1: "v1",
		2: "v2",
		3: "v3",
		4: "v4",
	})
	fmt.Println(m.X出栈多个(-1))
	fmt.Println("size:", m.X取数量())

	m.X设置值Map(g.MapIntAny{
		1: "v1",
		2: "v2",
		3: "v3",
		4: "v4",
	})
	fmt.Println(m.X出栈多个(2))
	fmt.Println("size:", m.X取数量())

	// May Output:
	// map[1:v1 2:v2 3:v3 4:v4]
	// size: 0
	// map[1:v1 2:v2]
	// size: 2
}

func ExampleIntAnyMap_GetOrSet() {
	m := map类.X创建IntAny()
	m.X设置值(1, "val1")

	fmt.Println(m.X取值或设置值(1, "NotExistValue"))
	fmt.Println(m.X取值或设置值(2, "val2"))

	// Output:
	// val1
	// val2
}

func ExampleIntAnyMap_GetOrSetFunc() {
	m := map类.X创建IntAny()
	m.X设置值(1, "val1")

	fmt.Println(m.X取值或设置值_函数(1, func() interface{} {
		return "NotExistValue"
	}))
	fmt.Println(m.X取值或设置值_函数(2, func() interface{} {
		return "NotExistValue"
	}))

	// Output:
	// val1
	// NotExistValue
}

func ExampleIntAnyMap_GetOrSetFuncLock() {
	m := map类.X创建IntAny()
	m.X设置值(1, "val1")

	fmt.Println(m.X取值或设置值_函数带锁(1, func() interface{} {
		return "NotExistValue"
	}))
	fmt.Println(m.X取值或设置值_函数带锁(2, func() interface{} {
		return "NotExistValue"
	}))

	// Output:
	// val1
	// NotExistValue
}

func ExampleIntAnyMap_GetVar() {
	m := map类.X创建IntAny()
	m.X设置值(1, "val1")

	fmt.Println(m.X取值泛型类(1))
	fmt.Println(m.X取值泛型类(2).X是否为Nil())

	// Output:
	// val1
	// true
}

func ExampleIntAnyMap_GetVarOrSet() {
	m := map类.X创建IntAny()
	m.X设置值(1, "val1")

	fmt.Println(m.X取值或设置值泛型类(1, "NotExistValue"))
	fmt.Println(m.X取值或设置值泛型类(2, "val2"))

	// Output:
	// val1
	// val2
}

func ExampleIntAnyMap_GetVarOrSetFunc() {
	m := map类.X创建IntAny()
	m.X设置值(1, "val1")

	fmt.Println(m.X取值或设置值泛型类_函数(1, func() interface{} {
		return "NotExistValue"
	}))
	fmt.Println(m.X取值或设置值泛型类_函数(2, func() interface{} {
		return "NotExistValue"
	}))

	// Output:
	// val1
	// NotExistValue
}

func ExampleIntAnyMap_GetVarOrSetFuncLock() {
	m := map类.X创建IntAny()
	m.X设置值(1, "val1")

	fmt.Println(m.X取值或设置值泛型类_函数带锁(1, func() interface{} {
		return "NotExistValue"
	}))
	fmt.Println(m.X取值或设置值泛型类_函数带锁(2, func() interface{} {
		return "NotExistValue"
	}))

	// Output:
	// val1
	// NotExistValue
}

func ExampleIntAnyMap_SetIfNotExist() {
	var m map类.IntAnyMap
	fmt.Println(m.X设置值并跳过已存在(1, "v1"))
	fmt.Println(m.X设置值并跳过已存在(1, "v2"))
	fmt.Println(m.X取Map())

	// Output:
	// true
	// false
	// map[1:v1]
}

func ExampleIntAnyMap_SetIfNotExistFunc() {
	var m map类.IntAnyMap
	fmt.Println(m.X设置值并跳过已存在_函数(1, func() interface{} {
		return "v1"
	}))
	fmt.Println(m.X设置值并跳过已存在_函数(1, func() interface{} {
		return "v2"
	}))
	fmt.Println(m.X取Map())

	// Output:
	// true
	// false
	// map[1:v1]
}

func ExampleIntAnyMap_SetIfNotExistFuncLock() {
	var m map类.IntAnyMap
	fmt.Println(m.X设置值并跳过已存在_函数带锁(1, func() interface{} {
		return "v1"
	}))
	fmt.Println(m.X设置值并跳过已存在_函数带锁(1, func() interface{} {
		return "v2"
	}))
	fmt.Println(m.X取Map())

	// Output:
	// true
	// false
	// map[1:v1]
}

func ExampleIntAnyMap_Remove() {
	var m map类.IntAnyMap
	m.X设置值(1, "v1")

	fmt.Println(m.X删除(1))
	fmt.Println(m.X删除(2))
	fmt.Println(m.X取数量())

	// Output:
	// v1
	// <nil>
	// 0
}

func ExampleIntAnyMap_Removes() {
	var m map类.IntAnyMap
	m.X设置值Map(g.MapIntAny{
		1: "v1",
		2: "v2",
		3: "v3",
		4: "v4",
	})

	removeList := make([]int, 2)
	removeList = append(removeList, 1)
	removeList = append(removeList, 2)

	m.X删除多个值(removeList)

	fmt.Println(m.X取Map())

	// Output:
	// map[3:v3 4:v4]
}

func ExampleIntAnyMap_Keys() {
	var m map类.IntAnyMap
	m.X设置值Map(g.MapIntAny{
		1: "v1",
		2: "v2",
		3: "v3",
		4: "v4",
	})
	fmt.Println(m.X取所有名称())

	// May Output:
	// [1 2 3 4]
}

func ExampleIntAnyMap_Values() {
	var m map类.IntAnyMap
	m.X设置值Map(g.MapIntAny{
		1: "v1",
		2: "v2",
		3: "v3",
		4: "v4",
	})
	fmt.Println(m.X取所有值())

	// May Output:
	// [v1 v2 v3 v4]
}

func ExampleIntAnyMap_Contains() {
	var m map类.IntAnyMap
	m.X设置值Map(g.MapIntAny{
		1: "v1",
		2: "v2",
		3: "v3",
		4: "v4",
	})

	fmt.Println(m.X是否存在(1))
	fmt.Println(m.X是否存在(5))

	// Output:
	// true
	// false
}

func ExampleIntAnyMap_Size() {
	var m map类.IntAnyMap
	m.X设置值Map(g.MapIntAny{
		1: "v1",
		2: "v2",
		3: "v3",
		4: "v4",
	})

	fmt.Println(m.X取数量())

	// Output:
	// 4
}

func ExampleIntAnyMap_IsEmpty() {
	var m map类.IntAnyMap
	fmt.Println(m.X是否为空())

	m.X设置值(1, "v1")
	fmt.Println(m.X是否为空())

	// Output:
	// true
	// false
}

func ExampleIntAnyMap_Clear() {
	var m map类.IntAnyMap
	m.X设置值Map(g.MapIntAny{
		1: "v1",
		2: "v2",
		3: "v3",
		4: "v4",
	})

	m.X清空()

	fmt.Println(m.X取Map())

	// Output:
	// map[]
}

func ExampleIntAnyMap_Replace() {
	var m map类.IntAnyMap
	m.X设置值Map(g.MapIntAny{
		1: "v1",
	})

	var n map类.IntAnyMap
	n.X设置值Map(g.MapIntAny{
		2: "v2",
	})

	fmt.Println(m.X取Map())

	m.X替换(n.X取Map())
	fmt.Println(m.X取Map())

	n.X设置值(2, "v1")
	fmt.Println(m.X取Map())

	// Output:
	// map[1:v1]
	// map[2:v2]
	// map[2:v1]
}

func ExampleIntAnyMap_LockFunc() {
	var m map类.IntAnyMap
	m.X设置值Map(g.MapIntAny{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	})

	m.X遍历写锁定(func(m map[int]interface{}) {
		totalValue := 0
		for _, v := range m {
			totalValue += v.(int)
		}
		fmt.Println("totalValue:", totalValue)
	})

	// Output:
	// totalValue: 10
}

func ExampleIntAnyMap_RLockFunc() {
	var m map类.IntAnyMap
	m.X设置值Map(g.MapIntAny{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	})

	m.X遍历读锁定(func(m map[int]interface{}) {
		totalValue := 0
		for _, v := range m {
			totalValue += v.(int)
		}
		fmt.Println("totalValue:", totalValue)
	})

	// Output:
	// totalValue: 10
}

func ExampleIntAnyMap_Flip() {
	var m map类.IntAnyMap
	m.X设置值Map(g.MapIntAny{
		1: 10,
	})
	m.X名称值交换()
	fmt.Println(m.X取Map())

	// Output:
	// map[10:1]
}

func ExampleIntAnyMap_Merge() {
	var m1, m2 map类.Map
	m1.X设置值(1, "val1")
	m2.X设置值(2, "val2")
	m1.X合并(&m2)
	fmt.Println(m1.X取Map())

	// May Output:
	// map[key1:val1 key2:val2]
}

func ExampleIntAnyMap_String() {
	var m map类.IntAnyMap
	m.X设置值Map(g.MapIntAny{
		1: "v1",
	})

	fmt.Println(m.String())

	var m1 *map类.IntAnyMap = nil
	fmt.Println(len(m1.String()))

	// Output:
	// {"1":"v1"}
	// 0
}

func ExampleIntAnyMap_MarshalJSON() {
	var m map类.IntAnyMap
	m.X设置值Map(g.MapIntAny{
		1: "v1",
		2: "v2",
		3: "v3",
		4: "v4",
	})

	bytes, err := json.Marshal(&m)
	if err == nil {
		fmt.Println(转换类.String(bytes))
	}

	// Output:
	// {"1":"v1","2":"v2","3":"v3","4":"v4"}
}

func ExampleIntAnyMap_UnmarshalJSON() {
	var m map类.IntAnyMap
	m.X设置值Map(g.MapIntAny{
		1: "v1",
		2: "v2",
		3: "v3",
		4: "v4",
	})

	var n map类.Map

	err := json.Unmarshal(转换类.X取字节集(m.String()), &n)
	if err == nil {
		fmt.Println(n.X取Map())
	}

	// Output:
	// map[1:v1 2:v2 3:v3 4:v4]
}

func ExampleIntAnyMap_UnmarshalValue() {
	var m map类.IntAnyMap

	goWeb := map[int]interface{}{
		1: "goframe",
		2: "gin",
		3: "echo",
	}

	if err := 转换类.Scan(goWeb, &m); err == nil {
		fmt.Printf("%#v", m.X取Map())
	}

	// Output:
	// map[int]interface {}{1:"goframe", 2:"gin", 3:"echo"}
}
