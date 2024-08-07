// 版权归GoFrame作者所有（https://goframe.org）。保留所有权利。
//
// 本源代码形式受MIT许可证条款的约束。如果gm文件中未附带MIT许可证的副本，
// 您可以从https://github.com/gogf/gf获取。
// md5:1d281c30cdc3423b

package map类_test

import (
	"fmt"

	gmap "github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/json"
	gconv "github.com/888go/goframe/util/gconv"
)

func ExampleIntIntMap_Iterator() {
	m := gmap.X创建IntInt()
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
	m := gmap.X创建IntInt()

	m.X设置值(1, 1)
	fmt.Println(m)

	n := m.X取副本()
	fmt.Println(n)

	// Output:
	// {"1":1}
	// {"1":1}
}

func ExampleIntIntMap_Map() {
		// 非并发安全，指向底层数据的指针. md5:0c201eaf65f11ed8
	m1 := gmap.X创建IntInt()
	m1.X设置值(1, 1)
	fmt.Println("m1:", m1)

	n1 := m1.X取Map()
	fmt.Println("before n1:", n1)
	m1.X设置值(1, 2)
	fmt.Println("after n1:", n1)

			// 并发安全，底层数据的副本. md5:114a4273430037c7
	m2 := gmap.X创建(true)
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
	m := gmap.X创建IntInt()

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
	m := gmap.X创建IntInt()
	m.X设置值(1001, 1)
	m.X设置值(1002, 2)

	n := m.X取MapStrAny()
	fmt.Printf("%#v", n)

	// Output:
	// map[string]interface {}{"1001":1, "1002":2}
}

func ExampleIntIntMap_FilterEmpty() {
	m := gmap.X创建IntInt并从Map(g.MapIntInt{
		1: 0,
		2: 1,
	})
	m.X删除所有空值()
	fmt.Println(m.X取Map())

	// Output:
	// map[2:1]
}

func ExampleIntIntMap_Set() {
	m := gmap.X创建IntInt()

	m.X设置值(1, 1)
	fmt.Println(m)

	// Output:
	// {"1":1}
}

func ExampleIntIntMap_Sets() {
	m := gmap.X创建IntInt()

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
	m := gmap.X创建IntInt()

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
	m := gmap.X创建IntInt()

	m.X设置值(1, 1)

	fmt.Println("key1 value:", m.X取值(1))
	fmt.Println("key2 value:", m.X取值(2))

	// Output:
	// key1 value: 1
	// key2 value: 0
}

func ExampleIntIntMap_Pop() {
	var m gmap.IntIntMap
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
	var m gmap.IntIntMap
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
	m := gmap.X创建IntInt()
	m.X设置值(1, 1)

	fmt.Println(m.X取值或设置值(1, 0))
	fmt.Println(m.X取值或设置值(2, 2))

	// Output:
	// 1
	// 2
}

func ExampleIntIntMap_GetOrSetFunc() {
	m := gmap.X创建IntInt()
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
	m := gmap.X创建IntInt()
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
	var m gmap.IntIntMap
	fmt.Println(m.X设置值并跳过已存在(1, 1))
	fmt.Println(m.X设置值并跳过已存在(1, 2))
	fmt.Println(m.X取Map())

	// Output:
	// true
	// false
	// map[1:1]
}

func ExampleIntIntMap_SetIfNotExistFunc() {
	var m gmap.IntIntMap
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
	var m gmap.IntIntMap
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
	var m gmap.IntIntMap
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
	var m gmap.IntIntMap
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
	var m gmap.IntIntMap
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
	var m gmap.IntIntMap
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
	var m gmap.IntIntMap
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
	var m gmap.IntIntMap
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
	var m gmap.IntIntMap
	fmt.Println(m.X是否为空())

	m.X设置值(1, 1)
	fmt.Println(m.X是否为空())

	// Output:
	// true
	// false
}

func ExampleIntIntMap_Clear() {
	var m gmap.IntIntMap
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
	var m gmap.IntIntMap
	m.X设置值Map(g.MapIntInt{
		1: 1,
	})

	var n gmap.IntIntMap
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
	var m gmap.IntIntMap
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
	var m gmap.IntIntMap
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
	var m gmap.IntIntMap
	m.X设置值Map(g.MapIntInt{
		1: 10,
	})
	m.X名称值交换()
	fmt.Println(m.X取Map())

	// Output:
	// map[10:1]
}

func ExampleIntIntMap_Merge() {
	var m1, m2 gmap.Map
	m1.X设置值(1, "1")
	m2.X设置值(2, "2")
	m1.X合并(&m2)
	fmt.Println(m1.X取Map())

	// May Output:
	// map[key1:1 key2:2]
}

func ExampleIntIntMap_String() {
	var m gmap.IntIntMap
	m.X设置值Map(g.MapIntInt{
		1: 1,
	})

	fmt.Println(m.String())

	var m1 *gmap.IntIntMap = nil
	fmt.Println(len(m1.String()))

	// Output:
	// {"1":1}
	// 0
}

func ExampleIntIntMap_MarshalJSON() {
	var m gmap.IntIntMap
	m.X设置值Map(g.MapIntInt{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	})

	bytes, err := json.Marshal(&m)
	if err == nil {
		fmt.Println(gconv.String(bytes))
	}

	// Output:
	// {"1":1,"2":2,"3":3,"4":4}
}

func ExampleIntIntMap_UnmarshalJSON() {
	var m gmap.IntIntMap
	m.X设置值Map(g.MapIntInt{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	})

	var n gmap.Map

	err := json.Unmarshal(gconv.X取字节集(m.String()), &n)
	if err == nil {
		fmt.Println(n.X取Map())
	}

	// Output:
	// map[1:1 2:2 3:3 4:4]
}

func ExampleIntIntMap_UnmarshalValue() {
	var m gmap.IntIntMap

	n := map[int]int{
		1: 1001,
		2: 1002,
		3: 1003,
	}

	if err := gconv.Scan(n, &m); err == nil {
		fmt.Printf("%#v", m.X取Map())
	}
	// Output:
	// map[int]int{1:1001, 2:1002, 3:1003}
}
