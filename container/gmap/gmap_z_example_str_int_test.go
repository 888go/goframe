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

func ExampleStrIntMap_Iterator() {
	m := gmap.X创建StrInt()
	for i := 0; i < 10; i++ {
		m.X设置值(gconv.String(i), i*2)
	}

	var totalValue int
	m.X遍历(func(k string, v int) bool {
		totalValue += v

		return totalValue < 50
	})

	fmt.Println("totalValue:", totalValue)

	// May Output:
	// totalValue: 52
}

func ExampleStrIntMap_Clone() {
	m := gmap.X创建StrInt()

	m.X设置值("key1", 1)
	fmt.Println(m)

	n := m.X取副本()
	fmt.Println(n)

	// Output:
	// {"key1":1}
	// {"key1":1}
}

func ExampleStrIntMap_Map() {
		// 非并发安全，指向底层数据的指针. md5:0c201eaf65f11ed8
	m1 := gmap.X创建StrInt()
	m1.X设置值("key1", 1)
	fmt.Println("m1:", m1)

	n1 := m1.X取Map()
	fmt.Println("before n1:", n1)
	m1.X设置值("key1", 2)
	fmt.Println("after n1:", n1)

			// 并发安全，底层数据的副本. md5:114a4273430037c7
	m2 := gmap.X创建StrInt(true)
	m2.X设置值("key1", 1)
	fmt.Println("m2:", m2)

	n2 := m2.X取Map()
	fmt.Println("before n2:", n2)
	m2.X设置值("key1", 2)
	fmt.Println("after n2:", n2)

	// Output:
	// m1: {"key1":1}
	// before n1: map[key1:1]
	// after n1: map[key1:2]
	// m2: {"key1":1}
	// before n2: map[key1:1]
	// after n2: map[key1:1]
}

func ExampleStrIntMap_MapCopy() {
	m := gmap.X创建StrInt()

	m.X设置值("key1", 1)
	m.X设置值("key2", 2)
	fmt.Println(m)

	n := m.X浅拷贝()
	fmt.Println(n)

	// Output:
	// {"key1":1,"key2":2}
	// map[key1:1 key2:2]
}

func ExampleStrIntMap_MapStrAny() {
	m := gmap.X创建StrInt()
	m.X设置值("key1", 1)
	m.X设置值("key2", 2)

	n := m.X取MapStrAny()
	fmt.Printf("%#v", n)

	// Output:
	// map[string]interface {}{"key1":1, "key2":2}
}

func ExampleStrIntMap_FilterEmpty() {
	m := gmap.X创建StrInt并从Map(g.MapStrInt{
		"k1": 0,
		"k2": 1,
	})
	m.X删除所有空值()
	fmt.Println(m.X取Map())

	// Output:
	// map[k2:1]
}

func ExampleStrIntMap_Set() {
	m := gmap.X创建StrInt()

	m.X设置值("key1", 1)
	fmt.Println(m)

	// Output:
	// {"key1":1}
}

func ExampleStrIntMap_Sets() {
	m := gmap.X创建StrInt()

	addMap := make(map[string]int)
	addMap["key1"] = 1
	addMap["key2"] = 2
	addMap["key3"] = 3

	m.X设置值Map(addMap)
	fmt.Println(m)

	// Output:
	// {"key1":1,"key2":2,"key3":3}
}

func ExampleStrIntMap_Search() {
	m := gmap.X创建StrInt()

	m.X设置值("key1", 1)

	value, found := m.X查找("key1")
	if found {
		fmt.Println("find key1 value:", value)
	}

	value, found = m.X查找("key2")
	if !found {
		fmt.Println("key2 not find")
	}

	// Output:
	// find key1 value: 1
	// key2 not find
}

func ExampleStrIntMap_Get() {
	m := gmap.X创建StrInt()

	m.X设置值("key1", 1)

	fmt.Println("key1 value:", m.X取值("key1"))
	fmt.Println("key2 value:", m.X取值("key2"))

	// Output:
	// key1 value: 1
	// key2 value: 0
}

func ExampleStrIntMap_Pop() {
	var m gmap.StrIntMap
	m.X设置值Map(g.MapStrInt{
		"k1": 1,
		"k2": 2,
		"k3": 3,
		"k4": 4,
	})

	fmt.Println(m.X出栈())

	// May Output:
	// k1 1
}

func ExampleStrIntMap_Pops() {
	var m gmap.StrIntMap
	m.X设置值Map(g.MapStrInt{
		"k1": 1,
		"k2": 2,
		"k3": 3,
		"k4": 4,
	})
	fmt.Println(m.X出栈多个(-1))
	fmt.Println("size:", m.X取数量())

	m.X设置值Map(g.MapStrInt{
		"k1": 1,
		"k2": 2,
		"k3": 3,
		"k4": 4,
	})
	fmt.Println(m.X出栈多个(2))
	fmt.Println("size:", m.X取数量())

	// May Output:
	// map[k1:1 k2:2 k3:3 k4:4]
	// size: 0
	// map[k1:1 k2:2]
	// size: 2
}

func ExampleStrIntMap_GetOrSet() {
	m := gmap.X创建StrInt()
	m.X设置值("key1", 1)

	fmt.Println(m.X取值或设置值("key1", 0))
	fmt.Println(m.X取值或设置值("key2", 2))

	// Output:
	// 1
	// 2
}

func ExampleStrIntMap_GetOrSetFunc() {
	m := gmap.X创建StrInt()
	m.X设置值("key1", 1)

	fmt.Println(m.X取值或设置值_函数("key1", func() int {
		return 0
	}))
	fmt.Println(m.X取值或设置值_函数("key2", func() int {
		return 0
	}))

	// Output:
	// 1
	// 0
}

func ExampleStrIntMap_GetOrSetFuncLock() {
	m := gmap.X创建StrInt()
	m.X设置值("key1", 1)

	fmt.Println(m.X取值或设置值_函数带锁("key1", func() int {
		return 0
	}))
	fmt.Println(m.X取值或设置值_函数带锁("key2", func() int {
		return 0
	}))

	// Output:
	// 1
	// 0
}

func ExampleStrIntMap_SetIfNotExist() {
	var m gmap.StrIntMap
	fmt.Println(m.X设置值并跳过已存在("k1", 1))
	fmt.Println(m.X设置值并跳过已存在("k1", 2))
	fmt.Println(m.X取Map())

	// Output:
	// true
	// false
	// map[k1:1]
}

func ExampleStrIntMap_SetIfNotExistFunc() {
	var m gmap.StrIntMap
	fmt.Println(m.X设置值并跳过已存在_函数("k1", func() int {
		return 1
	}))
	fmt.Println(m.X设置值并跳过已存在_函数("k1", func() int {
		return 2
	}))
	fmt.Println(m.X取Map())

	// Output:
	// true
	// false
	// map[k1:1]
}

func ExampleStrIntMap_SetIfNotExistFuncLock() {
	var m gmap.StrIntMap
	fmt.Println(m.X设置值并跳过已存在_函数带锁("k1", func() int {
		return 1
	}))
	fmt.Println(m.X设置值并跳过已存在_函数带锁("k1", func() int {
		return 2
	}))
	fmt.Println(m.X取Map())

	// Output:
	// true
	// false
	// map[k1:1]
}

func ExampleStrIntMap_Remove() {
	var m gmap.StrIntMap
	m.X设置值("k1", 1)

	fmt.Println(m.X删除("k1"))
	fmt.Println(m.X删除("k2"))
	fmt.Println(m.X取数量())

	// Output:
	// 1
	// 0
	// 0
}

func ExampleStrIntMap_Removes() {
	var m gmap.StrIntMap
	m.X设置值Map(g.MapStrInt{
		"k1": 1,
		"k2": 2,
		"k3": 3,
		"k4": 4,
	})

	removeList := make([]string, 2)
	removeList = append(removeList, "k1")
	removeList = append(removeList, "k2")

	m.X删除多个值(removeList)

	fmt.Println(m.X取Map())

	// Output:
	// map[k3:3 k4:4]
}

func ExampleStrIntMap_Keys() {
	var m gmap.StrIntMap
	m.X设置值Map(g.MapStrInt{
		"k1": 1,
		"k2": 2,
		"k3": 3,
		"k4": 4,
	})
	fmt.Println(m.X取所有名称())

	// May Output:
	// [k1 k2 k3 k4]
}

func ExampleStrIntMap_Values() {
	var m gmap.StrIntMap
	m.X设置值Map(g.MapStrInt{
		"k1": 1,
		"k2": 2,
		"k3": 3,
		"k4": 4,
	})
	fmt.Println(m.X取所有值())

	// May Output:
	// [1 2 3 4]
}

func ExampleStrIntMap_Contains() {
	var m gmap.StrIntMap
	m.X设置值Map(g.MapStrInt{
		"k1": 1,
		"k2": 2,
		"k3": 3,
		"k4": 4,
	})

	fmt.Println(m.X是否存在("k1"))
	fmt.Println(m.X是否存在("k5"))

	// Output:
	// true
	// false
}

func ExampleStrIntMap_Size() {
	var m gmap.StrIntMap
	m.X设置值Map(g.MapStrInt{
		"k1": 1,
		"k2": 2,
		"k3": 3,
		"k4": 4,
	})

	fmt.Println(m.X取数量())

	// Output:
	// 4
}

func ExampleStrIntMap_IsEmpty() {
	var m gmap.StrIntMap
	fmt.Println(m.X是否为空())

	m.X设置值("k1", 1)
	fmt.Println(m.X是否为空())

	// Output:
	// true
	// false
}

func ExampleStrIntMap_Clear() {
	var m gmap.StrIntMap
	m.X设置值Map(g.MapStrInt{
		"k1": 1,
		"k2": 2,
		"k3": 3,
		"k4": 4,
	})

	m.X清空()

	fmt.Println(m.X取Map())

	// Output:
	// map[]
}

func ExampleStrIntMap_Replace() {
	var m gmap.StrIntMap
	m.X设置值Map(g.MapStrInt{
		"k1": 1,
	})

	var n gmap.StrIntMap
	n.X设置值Map(g.MapStrInt{
		"k2": 2,
	})

	fmt.Println(m.X取Map())

	m.X替换(n.X取Map())
	fmt.Println(m.X取Map())

	n.X设置值("k2", 1)
	fmt.Println(m.X取Map())

	// Output:
	// map[k1:1]
	// map[k2:2]
	// map[k2:1]
}

func ExampleStrIntMap_LockFunc() {
	var m gmap.StrIntMap
	m.X设置值Map(g.MapStrInt{
		"k1": 1,
		"k2": 2,
		"k3": 3,
		"k4": 4,
	})

	m.X遍历写锁定(func(m map[string]int) {
		totalValue := 0
		for _, v := range m {
			totalValue += v
		}
		fmt.Println("totalValue:", totalValue)
	})

	// Output:
	// totalValue: 10
}

func ExampleStrIntMap_RLockFunc() {
	var m gmap.StrIntMap
	m.X设置值Map(g.MapStrInt{
		"k1": 1,
		"k2": 2,
		"k3": 3,
		"k4": 4,
	})

	m.X遍历读锁定(func(m map[string]int) {
		totalValue := 0
		for _, v := range m {
			totalValue += v
		}
		fmt.Println("totalValue:", totalValue)
	})

	// Output:
	// totalValue: 10
}

func ExampleStrIntMap_Flip() {
	var m gmap.StrIntMap
	m.X设置值Map(g.MapStrInt{
		"k1": 1,
	})
	m.X名称值交换()
	fmt.Println(m.X取Map())

	var n gmap.StrIntMap
	n.X设置值Map(g.MapStrInt{
		"11": 1,
	})
	n.X名称值交换()
	fmt.Println(n.X取Map())

	// Output:
	// map[1:0]
	// map[1:11]
}

func ExampleStrIntMap_Merge() {
	var m1, m2 gmap.StrIntMap
	m1.X设置值("key1", 1)
	m2.X设置值("key2", 2)
	m1.X合并(&m2)
	fmt.Println(m1.X取Map())

	// May Output:
	// map[key1:1 key2:2]
}

func ExampleStrIntMap_String() {
	var m gmap.StrIntMap
	m.X设置值Map(g.MapStrInt{
		"k1": 1,
	})

	fmt.Println(m.String())

	var m1 *gmap.StrIntMap = nil
	fmt.Println(len(m1.String()))

	// Output:
	// {"k1":1}
	// 0
}

func ExampleStrIntMap_MarshalJSON() {
	var m gmap.StrIntMap
	m.X设置值Map(g.MapStrInt{
		"k1": 1,
		"k2": 2,
		"k3": 3,
		"k4": 4,
	})

	bytes, err := json.Marshal(&m)
	if err == nil {
		fmt.Println(gconv.String(bytes))
	}

	// Output:
	// {"k1":1,"k2":2,"k3":3,"k4":4}
}

func ExampleStrIntMap_UnmarshalJSON() {
	var m gmap.StrIntMap
	m.X设置值Map(g.MapStrInt{
		"k1": 1,
		"k2": 2,
		"k3": 3,
		"k4": 4,
	})

	var n gmap.StrIntMap

	err := json.Unmarshal(gconv.X取字节集(m.String()), &n)
	if err == nil {
		fmt.Println(n.X取Map())
	}

	// Output:
	// map[k1:1 k2:2 k3:3 k4:4]
}

func ExampleStrIntMap_UnmarshalValue() {
	var m gmap.StrIntMap

	goWeb := map[string]int{
		"goframe": 1,
		"gin":     2,
		"echo":    3,
	}

	if err := gconv.Scan(goWeb, &m); err == nil {
		fmt.Printf("%#v", m.X取Map())
	}
	// Output:
	// map[string]int{"echo":3, "gin":2, "goframe":1}
}
