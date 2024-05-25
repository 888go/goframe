// 版权归GoFrame作者所有（https://goframe.org）。保留所有权利。
//
// 本源代码形式受MIT许可证条款的约束。如果gm文件中未附带MIT许可证的副本，
// 您可以从https://github.com/gogf/gf获取。
// md5:1d281c30cdc3423b

package gmap_test

import (
	"fmt"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/util/gutil"
)

func ExampleNew() {
	m := gmap.New()

	// Add data.
	m.Set("key1", "val1")

	// Print size.
	fmt.Println(m.Size())

	addMap := make(map[interface{}]interface{})
	addMap["key2"] = "val2"
	addMap["key3"] = "val3"
	addMap[1] = 1

	fmt.Println(m.Values())

	// Batch add data.
	m.Sets(addMap)

	// 获取对应键的值。. md5:80c18e88718ab20c
	fmt.Println(m.Get("key3"))

	// 如果键不存在，使用给定的键值对获取或设置值。. md5:6110c826554187ae
	fmt.Println(m.GetOrSet("key4", "val4"))

	// 如果键不存在，则设置键值对并返回true；否则，返回false。. md5:cc39f71c7fcd3e86
	fmt.Println(m.SetIfNotExist("key3", "val3"))

	// Remove key
	m.Remove("key2")
	fmt.Println(m.Keys())

	// Batch remove keys.
	m.Removes([]interface{}{"key1", 1})
	fmt.Println(m.Keys())

	// Contains 检查一个键是否存在。. md5:822d7ed9848c03c4
	fmt.Println(m.Contains("key3"))

	// Flip 会交换映射中的键值，它会将键值对改为值键对。. md5:23981f3e2b855f36
	m.Flip()
	fmt.Println(m.Map())

	// Clear 删除map中的所有数据。. md5:7812c72bbfe807df
	m.Clear()

	fmt.Println(m.Size())

	// May Output:
	// 1
	// [val1]
	// val3
	// val4
	// false
	// [key4 key1 key3 1]
	// [key4 key3]
	// true
	// map[val3:key3 val4:key4]
	// 0
}

func ExampleNewFrom() {
	m := gmap.New()

	m.Set("key1", "val1")
	fmt.Println(m)

	n := gmap.NewFrom(m.MapCopy(), true)
	fmt.Println(n)

	// Output:
	// {"key1":"val1"}
	// {"key1":"val1"}
}

func ExampleNewHashMap() {
	m := gmap.NewHashMap()

	m.Set("key1", "val1")
	fmt.Println(m)

	// Output:
	// {"key1":"val1"}
}

func ExampleNewHashMapFrom() {
	m := gmap.New()

	m.Set("key1", "val1")
	fmt.Println(m)

	n := gmap.NewHashMapFrom(m.MapCopy(), true)
	fmt.Println(n)

	// Output:
	// {"key1":"val1"}
	// {"key1":"val1"}
}

func ExampleNewAnyAnyMap() {
	m := gmap.NewAnyAnyMap()

	m.Set("key1", "val1")
	fmt.Println(m)

	// Output:
	// {"key1":"val1"}
}

func ExampleNewAnyAnyMapFrom() {
	m := gmap.NewAnyAnyMap()

	m.Set("key1", "val1")
	fmt.Println(m)

	n := gmap.NewAnyAnyMapFrom(m.MapCopy(), true)
	fmt.Println(n)

	// Output:
	// {"key1":"val1"}
	// {"key1":"val1"}
}

func ExampleNewIntAnyMap() {
	m := gmap.NewIntAnyMap()

	m.Set(1, "val1")
	fmt.Println(m)

	// Output:
	// {"1":"val1"}
}

func ExampleNewIntAnyMapFrom() {
	m := gmap.NewIntAnyMap()

	m.Set(1, "val1")
	fmt.Println(m)

	n := gmap.NewIntAnyMapFrom(m.MapCopy(), true)
	fmt.Println(n)

	// Output:
	// {"1":"val1"}
	// {"1":"val1"}
}

func ExampleNewIntIntMap() {
	m := gmap.NewIntIntMap()

	m.Set(1, 1)
	fmt.Println(m)

	// Output:
	// {"1":1}
}

func ExampleNewIntIntMapFrom() {
	m := gmap.NewIntIntMap()

	m.Set(1, 1)
	fmt.Println(m)

	n := gmap.NewIntIntMapFrom(m.MapCopy(), true)
	fmt.Println(n)

	// Output:
	// {"1":1}
	// {"1":1}
}

func ExampleNewStrAnyMap() {
	m := gmap.NewStrAnyMap()

	m.Set("key1", "var1")
	fmt.Println(m)

	// Output:
	// {"key1":"var1"}
}

func ExampleNewStrAnyMapFrom() {
	m := gmap.NewStrAnyMap()

	m.Set("key1", "var1")
	fmt.Println(m)

	n := gmap.NewStrAnyMapFrom(m.MapCopy(), true)
	fmt.Println(n)

	// Output:
	// {"key1":"var1"}
	// {"key1":"var1"}
}

func ExampleNewStrIntMap() {
	m := gmap.NewStrIntMap()

	m.Set("key1", 1)
	fmt.Println(m)

	// Output:
	// {"key1":1}
}

func ExampleNewStrIntMapFrom() {
	m := gmap.NewStrIntMap()

	m.Set("key1", 1)
	fmt.Println(m)

	n := gmap.NewStrIntMapFrom(m.MapCopy(), true)
	fmt.Println(n)

	// Output:
	// {"key1":1}
	// {"key1":1}
}

func ExampleNewStrStrMap() {
	m := gmap.NewStrStrMap()

	m.Set("key1", "var1")
	fmt.Println(m)

	// Output:
	// {"key1":"var1"}
}

func ExampleNewStrStrMapFrom() {
	m := gmap.NewStrStrMap()

	m.Set("key1", "var1")
	fmt.Println(m)

	n := gmap.NewStrStrMapFrom(m.MapCopy(), true)
	fmt.Println(n)

	// Output:
	// {"key1":"var1"}
	// {"key1":"var1"}
}

func ExampleNewListMap() {
	m := gmap.NewListMap()

	m.Set("key1", "var1")
	m.Set("key2", "var2")
	fmt.Println(m)

	// Output:
	// {"key1":"var1","key2":"var2"}
}

func ExampleNewListMapFrom() {
	m := gmap.NewListMap()

	m.Set("key1", "var1")
	m.Set("key2", "var2")
	fmt.Println(m)

	n := gmap.NewListMapFrom(m.Map(), true)
	fmt.Println(n)

	// May Output:
	// {"key1":"var1","key2":"var2"}
	// {"key1":"var1","key2":"var2"}
}

func ExampleNewTreeMap() {
	m := gmap.NewTreeMap(gutil.ComparatorString)

	m.Set("key2", "var2")
	m.Set("key1", "var1")

	fmt.Println(m.Map())

	// May Output:
	// map[key1:var1 key2:var2]
}

func ExampleNewTreeMapFrom() {
	m := gmap.NewTreeMap(gutil.ComparatorString)

	m.Set("key2", "var2")
	m.Set("key1", "var1")

	fmt.Println(m.Map())

	n := gmap.NewListMapFrom(m.Map(), true)
	fmt.Println(n.Map())

	// May Output:
	// map[key1:var1 key2:var2]
	// map[key1:var1 key2:var2]
}
