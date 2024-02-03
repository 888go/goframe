// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

package gmap_test

import (
	"fmt"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/util/gutil"
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

	// 获取相应键的值。
	fmt.Println(m.Get("key3"))

	// 通过键获取值，如果不存在则使用给定的键值对设置它。
	fmt.Println(m.GetOrSet("key4", "val4"))

	// 如果键不存在，则设置键值对并返回true；否则返回false。
	fmt.Println(m.SetIfNotExist("key3", "val3"))

	// Remove key
	m.Remove("key2")
	fmt.Println(m.Keys())

	// 批量删除键。
	m.Removes([]interface{}{"key1", 1})
	fmt.Println(m.Keys())

	// Contains 检查键是否存在。
	fmt.Println(m.Contains("key3"))

	// Flip交换映射中的键值对，它会将键值对改为值键对。
	m.Flip()
	fmt.Println(m.Map())

	// 清空删除映射中的所有数据。
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
