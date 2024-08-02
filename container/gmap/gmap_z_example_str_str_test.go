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

func ExampleStrStrMap_Iterator() {
	m := gmap.NewStrStrMap()
	for i := 0; i < 10; i++ {
		m.Set("key"+gconv.String(i), "var"+gconv.String(i))
	}

	var str string
	m.Iterator(func(k string, v string) bool {

		str += v + "|"

		return len(str) < 20
	})

	fmt.Println("str:", str)

	// May Output:
	// var0|var1|var2|var3|
}

func ExampleStrStrMap_Clone() {
	m := gmap.NewStrStrMap()

	m.Set("key1", "val1")
	fmt.Println(m)

	n := m.Clone()
	fmt.Println(n)

	// Output:
	// {"key1":"val1"}
	// {"key1":"val1"}
}

func ExampleStrStrMap_Map() {
		// 非并发安全，指向底层数据的指针. md5:0c201eaf65f11ed8
	m1 := gmap.NewStrStrMap()
	m1.Set("key1", "val1")
	fmt.Println("m1:", m1)

	n1 := m1.Map()
	fmt.Println("before n1:", n1)
	m1.Set("key1", "val2")
	fmt.Println("after n1:", n1)

			// 并发安全，底层数据的副本. md5:114a4273430037c7
	m2 := gmap.NewStrStrMap(true)
	m2.Set("key1", "val1")
	fmt.Println("m2:", m2)

	n2 := m2.Map()
	fmt.Println("before n2:", n2)
	m2.Set("key1", "val2")
	fmt.Println("after n2:", n2)

	// Output:
	// m1: {"key1":"val1"}
	// before n1: map[key1:val1]
	// after n1: map[key1:val2]
	// m2: {"key1":"val1"}
	// before n2: map[key1:val1]
	// after n2: map[key1:val1]
}

func ExampleStrStrMap_MapCopy() {
	m := gmap.NewStrStrMap()

	m.Set("key1", "val1")
	m.Set("key2", "val2")
	fmt.Println(m)

	n := m.MapCopy()
	fmt.Println(n)

	// Output:
	// {"key1":"val1","key2":"val2"}
	// map[key1:val1 key2:val2]
}

func ExampleStrStrMap_MapStrAny() {
	m := gmap.NewStrStrMap()
	m.Set("key1", "val1")
	m.Set("key2", "val2")

	n := m.MapStrAny()
	fmt.Printf("%#v", n)

	// Output:
	// map[string]interface {}{"key1":"val1", "key2":"val2"}
}

func ExampleStrStrMap_FilterEmpty() {
	m := gmap.NewStrStrMapFrom(g.MapStrStr{
		"k1": "",
		"k2": "v2",
	})
	m.FilterEmpty()
	fmt.Println(m.Map())

	// Output:
	// map[k2:v2]
}

func ExampleStrStrMap_Set() {
	m := gmap.NewStrStrMap()

	m.Set("key1", "val1")
	fmt.Println(m)

	// Output:
	// {"key1":"val1"}
}

func ExampleStrStrMap_Sets() {
	m := gmap.NewStrStrMap()

	addMap := make(map[string]string)
	addMap["key1"] = "val1"
	addMap["key2"] = "val2"
	addMap["key3"] = "val3"

	m.Sets(addMap)
	fmt.Println(m)

	// Output:
	// {"key1":"val1","key2":"val2","key3":"val3"}
}

func ExampleStrStrMap_Search() {
	m := gmap.NewStrStrMap()

	m.Set("key1", "val1")

	value, found := m.Search("key1")
	if found {
		fmt.Println("find key1 value:", value)
	}

	value, found = m.Search("key2")
	if !found {
		fmt.Println("key2 not find")
	}

	// Output:
	// find key1 value: val1
	// key2 not find
}

func ExampleStrStrMap_Get() {
	m := gmap.NewStrStrMap()

	m.Set("key1", "val1")

	fmt.Println("key1 value:", m.Get("key1"))
	fmt.Println("key2 value:", m.Get("key2"))

	// Output:
	// key1 value: val1
	// key2 value:
}

func ExampleStrStrMap_Pop() {
	var m gmap.StrStrMap
	m.Sets(g.MapStrStr{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	fmt.Println(m.Pop())

	// May Output:
	// k1 v1
}

func ExampleStrStrMap_Pops() {
	var m gmap.StrStrMap
	m.Sets(g.MapStrStr{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})
	fmt.Println(m.Pops(-1))
	fmt.Println("size:", m.Size())

	m.Sets(g.MapStrStr{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})
	fmt.Println(m.Pops(2))
	fmt.Println("size:", m.Size())

	// May Output:
	// map[k1:v1 k2:v2 k3:v3 k4:v4]
	// size: 0
	// map[k1:v1 k2:v2]
	// size: 2
}

func ExampleStrStrMap_GetOrSet() {
	m := gmap.NewStrStrMap()
	m.Set("key1", "val1")

	fmt.Println(m.GetOrSet("key1", "NotExistValue"))
	fmt.Println(m.GetOrSet("key2", "val2"))

	// Output:
	// val1
	// val2
}

func ExampleStrStrMap_GetOrSetFunc() {
	m := gmap.NewStrStrMap()
	m.Set("key1", "val1")

	fmt.Println(m.GetOrSetFunc("key1", func() string {
		return "NotExistValue"
	}))
	fmt.Println(m.GetOrSetFunc("key2", func() string {
		return "NotExistValue"
	}))

	// Output:
	// val1
	// NotExistValue
}

func ExampleStrStrMap_GetOrSetFuncLock() {
	m := gmap.NewStrStrMap()
	m.Set("key1", "val1")

	fmt.Println(m.GetOrSetFuncLock("key1", func() string {
		return "NotExistValue"
	}))
	fmt.Println(m.GetOrSetFuncLock("key2", func() string {
		return "NotExistValue"
	}))

	// Output:
	// val1
	// NotExistValue
}

func ExampleStrStrMap_SetIfNotExist() {
	var m gmap.StrStrMap
	fmt.Println(m.SetIfNotExist("k1", "v1"))
	fmt.Println(m.SetIfNotExist("k1", "v2"))
	fmt.Println(m.Map())

	// Output:
	// true
	// false
	// map[k1:v1]
}

func ExampleStrStrMap_SetIfNotExistFunc() {
	var m gmap.StrStrMap
	fmt.Println(m.SetIfNotExistFunc("k1", func() string {
		return "v1"
	}))
	fmt.Println(m.SetIfNotExistFunc("k1", func() string {
		return "v2"
	}))
	fmt.Println(m.Map())

	// Output:
	// true
	// false
	// map[k1:v1]
}

func ExampleStrStrMap_SetIfNotExistFuncLock() {
	var m gmap.StrStrMap
	fmt.Println(m.SetIfNotExistFuncLock("k1", func() string {
		return "v1"
	}))
	fmt.Println(m.SetIfNotExistFuncLock("k1", func() string {
		return "v2"
	}))
	fmt.Println(m.Map())

	// Output:
	// true
	// false
	// map[k1:v1]
}

func ExampleStrStrMap_Remove() {
	var m gmap.StrStrMap
	m.Set("k1", "v1")

	fmt.Println(m.Remove("k1"))
	fmt.Println(len(m.Remove("k2")))
	fmt.Println(m.Size())

	// Output:
	// v1
	// 0
	// 0
}

func ExampleStrStrMap_Removes() {
	var m gmap.StrStrMap
	m.Sets(g.MapStrStr{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	removeList := make([]string, 2)
	removeList = append(removeList, "k1")
	removeList = append(removeList, "k2")

	m.Removes(removeList)

	fmt.Println(m.Map())

	// Output:
	// map[k3:v3 k4:v4]
}

func ExampleStrStrMap_Keys() {
	var m gmap.StrStrMap
	m.Sets(g.MapStrStr{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})
	fmt.Println(m.Keys())

	// May Output:
	// [k1 k2 k3 k4]
}

func ExampleStrStrMap_Values() {
	var m gmap.StrStrMap
	m.Sets(g.MapStrStr{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})
	fmt.Println(m.Values())

	// May Output:
	// [v1 v2 v3 v4]
}

func ExampleStrStrMap_Contains() {
	var m gmap.StrStrMap
	m.Sets(g.MapStrStr{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	fmt.Println(m.Contains("k1"))
	fmt.Println(m.Contains("k5"))

	// Output:
	// true
	// false
}

func ExampleStrStrMap_Size() {
	var m gmap.StrStrMap
	m.Sets(g.MapStrStr{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	fmt.Println(m.Size())

	// Output:
	// 4
}

func ExampleStrStrMap_IsEmpty() {
	var m gmap.StrStrMap
	fmt.Println(m.IsEmpty())

	m.Set("k1", "v1")
	fmt.Println(m.IsEmpty())

	// Output:
	// true
	// false
}

func ExampleStrStrMap_Clear() {
	var m gmap.StrStrMap
	m.Sets(g.MapStrStr{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	m.Clear()

	fmt.Println(m.Map())

	// Output:
	// map[]
}

func ExampleStrStrMap_Replace() {
	var m gmap.StrStrMap
	m.Sets(g.MapStrStr{
		"k1": "v1",
	})

	var n gmap.StrStrMap
	n.Sets(g.MapStrStr{
		"k2": "v2",
	})

	fmt.Println(m.Map())

	m.Replace(n.Map())
	fmt.Println(m.Map())

	n.Set("k2", "v1")
	fmt.Println(m.Map())

	// Output:
	// map[k1:v1]
	// map[k2:v2]
	// map[k2:v1]
}

func ExampleStrStrMap_LockFunc() {
	var m gmap.StrStrMap
	m.Sets(g.MapStrStr{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	m.LockFunc(func(m map[string]string) {
		for k, v := range m {
			fmt.Println("key:", k, " value:", v)
		}
	})

	// May Output:
	// key: k1  value: v1
	// key: k2  value: v2
	// key: k3  value: v3
	// key: k4  value: v4
}

func ExampleStrStrMap_RLockFunc() {
	var m gmap.StrStrMap
	m.Sets(g.MapStrStr{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	m.RLockFunc(func(m map[string]string) {
		for k, v := range m {
			fmt.Println("key:", k, " value:", v)
		}
	})

	// May Output:
	// key: k1  value: v1
	// key: k2  value: v2
	// key: k3  value: v3
	// key: k4  value: v4
}

func ExampleStrStrMap_Flip() {
	var m gmap.StrStrMap
	m.Sets(g.MapStrStr{
		"k1": "v1",
	})
	m.Flip()
	fmt.Println(m.Map())

	// Output:
	// map[v1:k1]
}

func ExampleStrStrMap_Merge() {
	var m1, m2 gmap.StrStrMap
	m1.Set("key1", "val1")
	m2.Set("key2", "val2")
	m1.Merge(&m2)
	fmt.Println(m1.Map())

	// May Output:
	// map[key1:val1 key2:val2]
}

func ExampleStrStrMap_String() {
	var m gmap.StrStrMap
	m.Sets(g.MapStrStr{
		"k1": "v1",
	})

	fmt.Println(m.String())

	var m1 *gmap.StrStrMap = nil
	fmt.Println(len(m1.String()))

	// Output:
	// {"k1":"v1"}
	// 0
}

func ExampleStrStrMap_MarshalJSON() {
	var m gmap.StrStrMap
	m.Sets(g.MapStrStr{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	bytes, err := json.Marshal(&m)
	if err == nil {
		fmt.Println(gconv.String(bytes))
	}

	// Output:
	// {"k1":"v1","k2":"v2","k3":"v3","k4":"v4"}
}

func ExampleStrStrMap_UnmarshalJSON() {
	var m gmap.StrStrMap
	m.Sets(g.MapStrStr{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	var n gmap.StrStrMap

	err := json.Unmarshal(gconv.Bytes(m.String()), &n)
	if err == nil {
		fmt.Println(n.Map())
	}

	// Output:
	// map[k1:v1 k2:v2 k3:v3 k4:v4]
}

func ExampleStrStrMap_UnmarshalValue() {
	var m gmap.StrStrMap

	goWeb := map[string]string{
		"goframe": "https://goframe.org",
		"gin":     "https://gin-gonic.com/",
		"echo":    "https://echo.labstack.com/",
	}

	if err := gconv.Scan(goWeb, &m); err == nil {
		fmt.Printf("%#v", m.Map())
	}
	// Output:
	// map[string]string{"echo":"https://echo.labstack.com/", "gin":"https://gin-gonic.com/", "goframe":"https://goframe.org"}
}
