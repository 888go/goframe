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

func ExampleIntIntMap_Iterator() {
	m := map类.X创建IntInt()
	for i := 0; i < 10; i++ {
		m.X设置值(i, i*2)
	}

	var totalKey, totalValue int
	m.X遍历(func(k int, v int) bool {
		totalKey += k
		totalValue += v

		return totalKey < 10
	})

	fmt.Println("totalKey:", totalKey)
	fmt.Println("totalValue:", totalValue)

	// May Output:
	// totalKey: 11
	// totalValue: 22
}

func ExampleIntIntMap_Clone() {
	m := map类.X创建IntInt()

	m.X设置值(1, 1)
	fmt.Println(m)

	n := m.X取副本()
	fmt.Println(n)

	// Output:
	// {"1":1}
	// {"1":1}
}

func ExampleIntIntMap_Map() {
	// 非并发安全，指向底层数据的指针
	m1 := map类.X创建IntInt()
	m1.X设置值(1, 1)
	fmt.Println("m1:", m1)

	n1 := m1.X取Map()
	fmt.Println("before n1:", n1)
	m1.X设置值(1, 2)
	fmt.Println("after n1:", n1)

	// 并发安全，复制底层数据
	m2 := map类.X创建(true)
	m2.X设置值(1, "1")
	fmt.Println("m2:", m2)

	n2 := m2.X取Map()
	fmt.Println("before n2:", n2)
	m2.X设置值(1, "2")
	fmt.Println("after n2:", n2)

	// Output:
	// m1: {"1":1}
	// before n1: map[1:1]
	// after n1: map[1:2]
	// m2: {"1":"1"}
	// before n2: map[1:1]
	// after n2: map[1:1]
}

func ExampleIntIntMap_MapCopy() {
	m := map类.X创建IntInt()

	m.X设置值(1, 1)
	m.X设置值(2, 2)
	fmt.Println(m)

	n := m.X浅拷贝()
	fmt.Println(n)

	// Output:
	// {"1":1,"2":2}
	// map[1:1 2:2]
}

func ExampleIntIntMap_MapStrAny() {
	m := map类.X创建IntInt()
	m.X设置值(1001, 1)
	m.X设置值(1002, 2)

	n := m.X取MapStrAny()
	fmt.Printf("%#v", n)

	// Output:
	// map[string]interface {}{"1001":1, "1002":2}
}

func ExampleIntIntMap_FilterEmpty() {
	m := map类.X创建IntInt并从Map(g.MapIntInt{
		1: 0,
		2: 1,
	})
	m.X删除所有空值()
	fmt.Println(m.X取Map())

	// Output:
	// map[2:1]
}

func ExampleIntIntMap_Set() {
	m := map类.X创建IntInt()

	m.X设置值(1, 1)
	fmt.Println(m)

	// Output:
	// {"1":1}
}

func ExampleIntIntMap_Sets() {
	m := map类.X创建IntInt()

	addMap := make(map[int]int)
	addMap[1] = 1
	addMap[2] = 12
	addMap[3] = 123

	m.X设置值Map(addMap)
	fmt.Println(m)

	// Output:
	// {"1":1,"2":12,"3":123}
}

func ExampleIntIntMap_Search() {
	m := map类.X创建IntInt()

	m.X设置值(1, 1)

	value, found := m.X查找(1)
	if found {
		fmt.Println("find key1 value:", value)
	}

	value, found = m.X查找(2)
	if !found {
		fmt.Println("key2 not find")
	}

	// Output:
	// find key1 value: 1
	// key2 not find
}

func ExampleIntIntMap_Get() {
	m := map类.X创建IntInt()

	m.X设置值(1, 1)

	fmt.Println("key1 value:", m.X取值(1))
	fmt.Println("key2 value:", m.X取值(2))

	// Output:
	// key1 value: 1
	// key2 value: 0
}

func ExampleIntIntMap_Pop() {
	var m map类.IntIntMap
	m.X设置值Map(g.MapIntInt{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	})

	fmt.Println(m.X出栈())

	// May Output:
	// 1 1
}

func ExampleIntIntMap_Pops() {
	var m map类.IntIntMap
	m.X设置值Map(g.MapIntInt{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	})
	fmt.Println(m.X出栈多个(-1))
	fmt.Println("size:", m.X取数量())

	m.X设置值Map(g.MapIntInt{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	})
	fmt.Println(m.X出栈多个(2))
	fmt.Println("size:", m.X取数量())

	// May Output:
	// map[1:1 2:2 3:3 4:4]
	// size: 0
	// map[1:1 2:2]
	// size: 2
}

func ExampleIntIntMap_GetOrSet() {
	m := map类.X创建IntInt()
	m.X设置值(1, 1)

	fmt.Println(m.X取值或设置值(1, 0))
	fmt.Println(m.X取值或设置值(2, 2))

	// Output:
	// 1
	// 2
}

func ExampleIntIntMap_GetOrSetFunc() {
	m := map类.X创建IntInt()
	m.X设置值(1, 1)

	fmt.Println(m.X取值或设置值_函数(1, func() int {
		return 0
	}))
	fmt.Println(m.X取值或设置值_函数(2, func() int {
		return 0
	}))

	// Output:
	// 1
	// 0
}

func ExampleIntIntMap_GetOrSetFuncLock() {
	m := map类.X创建IntInt()
	m.X设置值(1, 1)

	fmt.Println(m.X取值或设置值_函数带锁(1, func() int {
		return 0
	}))
	fmt.Println(m.X取值或设置值_函数带锁(2, func() int {
		return 0
	}))

	// Output:
	// 1
	// 0
}

func ExampleIntIntMap_SetIfNotExist() {
	var m map类.IntIntMap
	fmt.Println(m.X设置值并跳过已存在(1, 1))
	fmt.Println(m.X设置值并跳过已存在(1, 2))
	fmt.Println(m.X取Map())

	// Output:
	// true
	// false
	// map[1:1]
}

func ExampleIntIntMap_SetIfNotExistFunc() {
	var m map类.IntIntMap
	fmt.Println(m.X设置值并跳过已存在_函数(1, func() int {
		return 1
	}))
	fmt.Println(m.X设置值并跳过已存在_函数(1, func() int {
		return 2
	}))
	fmt.Println(m.X取Map())

	// Output:
	// true
	// false
	// map[1:1]
}

func ExampleIntIntMap_SetIfNotExistFuncLock() {
	var m map类.IntIntMap
	fmt.Println(m.X设置值并跳过已存在_函数带锁(1, func() int {
		return 1
	}))
	fmt.Println(m.X设置值并跳过已存在_函数带锁(1, func() int {
		return 2
	}))
	fmt.Println(m.X取Map())

	// Output:
	// true
	// false
	// map[1:1]
}

func ExampleIntIntMap_Remove() {
	var m map类.IntIntMap
	m.X设置值(1, 1)

	fmt.Println(m.X删除(1))
	fmt.Println(m.X删除(2))
	fmt.Println(m.X取数量())

	// Output:
	// 1
	// 0
	// 0
}

func ExampleIntIntMap_Removes() {
	var m map类.IntIntMap
	m.X设置值Map(g.MapIntInt{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	})

	removeList := make([]int, 2)
	removeList = append(removeList, 1)
	removeList = append(removeList, 2)

	m.X删除多个值(removeList)

	fmt.Println(m.X取Map())

	// Output:
	// map[3:3 4:4]
}

func ExampleIntIntMap_Keys() {
	var m map类.IntIntMap
	m.X设置值Map(g.MapIntInt{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	})
	fmt.Println(m.X取所有名称())

	// May Output:
	// [1 2 3 4]
}

func ExampleIntIntMap_Values() {
	var m map类.IntIntMap
	m.X设置值Map(g.MapIntInt{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	})
	fmt.Println(m.X取所有值())

	// May Output:
	// [1 v2 v3 4]
}

func ExampleIntIntMap_Contains() {
	var m map类.IntIntMap
	m.X设置值Map(g.MapIntInt{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	})

	fmt.Println(m.X是否存在(1))
	fmt.Println(m.X是否存在(5))

	// Output:
	// true
	// false
}

func ExampleIntIntMap_Size() {
	var m map类.IntIntMap
	m.X设置值Map(g.MapIntInt{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	})

	fmt.Println(m.X取数量())

	// Output:
	// 4
}

func ExampleIntIntMap_IsEmpty() {
	var m map类.IntIntMap
	fmt.Println(m.X是否为空())

	m.X设置值(1, 1)
	fmt.Println(m.X是否为空())

	// Output:
	// true
	// false
}

func ExampleIntIntMap_Clear() {
	var m map类.IntIntMap
	m.X设置值Map(g.MapIntInt{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	})

	m.X清空()

	fmt.Println(m.X取Map())

	// Output:
	// map[]
}

func ExampleIntIntMap_Replace() {
	var m map类.IntIntMap
	m.X设置值Map(g.MapIntInt{
		1: 1,
	})

	var n map类.IntIntMap
	n.X设置值Map(g.MapIntInt{
		2: 2,
	})

	fmt.Println(m.X取Map())

	m.X替换(n.X取Map())
	fmt.Println(m.X取Map())

	n.X设置值(2, 1)
	fmt.Println(m.X取Map())

	// Output:
	// map[1:1]
	// map[2:2]
	// map[2:1]
}

func ExampleIntIntMap_LockFunc() {
	var m map类.IntIntMap
	m.X设置值Map(g.MapIntInt{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	})

	m.X遍历写锁定(func(m map[int]int) {
		totalValue := 0
		for _, v := range m {
			totalValue += v
		}
		fmt.Println("totalValue:", totalValue)
	})

	// Output:
	// totalValue: 10
}

func ExampleIntIntMap_RLockFunc() {
	var m map类.IntIntMap
	m.X设置值Map(g.MapIntInt{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	})

	m.X遍历读锁定(func(m map[int]int) {
		totalValue := 0
		for _, v := range m {
			totalValue += v
		}
		fmt.Println("totalValue:", totalValue)
	})

	// Output:
	// totalValue: 10
}

func ExampleIntIntMap_Flip() {
	var m map类.IntIntMap
	m.X设置值Map(g.MapIntInt{
		1: 10,
	})
	m.X名称值交换()
	fmt.Println(m.X取Map())

	// Output:
	// map[10:1]
}

func ExampleIntIntMap_Merge() {
	var m1, m2 map类.Map
	m1.X设置值(1, "1")
	m2.X设置值(2, "2")
	m1.X合并(&m2)
	fmt.Println(m1.X取Map())

	// May Output:
	// map[key1:1 key2:2]
}

func ExampleIntIntMap_String() {
	var m map类.IntIntMap
	m.X设置值Map(g.MapIntInt{
		1: 1,
	})

	fmt.Println(m.String())

	var m1 *map类.IntIntMap = nil
	fmt.Println(len(m1.String()))

	// Output:
	// {"1":1}
	// 0
}

func ExampleIntIntMap_MarshalJSON() {
	var m map类.IntIntMap
	m.X设置值Map(g.MapIntInt{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	})

	bytes, err := json.Marshal(&m)
	if err == nil {
		fmt.Println(转换类.String(bytes))
	}

	// Output:
	// {"1":1,"2":2,"3":3,"4":4}
}

func ExampleIntIntMap_UnmarshalJSON() {
	var m map类.IntIntMap
	m.X设置值Map(g.MapIntInt{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	})

	var n map类.Map

	err := json.Unmarshal(转换类.X取字节集(m.String()), &n)
	if err == nil {
		fmt.Println(n.X取Map())
	}

	// Output:
	// map[1:1 2:2 3:3 4:4]
}

func ExampleIntIntMap_UnmarshalValue() {
	var m map类.IntIntMap

	n := map[int]int{
		1: 1001,
		2: 1002,
		3: 1003,
	}

	if err := 转换类.Scan(n, &m); err == nil {
		fmt.Printf("%#v", m.X取Map())
	}
	// Output:
	// map[int]int{1:1001, 2:1002, 3:1003}
}
