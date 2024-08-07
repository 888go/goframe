// 版权归GoFrame作者所有（https://goframe.org）。保留所有权利。
//
// 本源代码形式受MIT许可证条款的约束。如果gm文件中未附带MIT许可证的副本，
// 您可以从https://github.com/gogf/gf获取。
// md5:1d281c30cdc3423b

package map类_test

import (
	"fmt"

	gmap "github.com/888go/goframe/container/gmap"
	gutil "github.com/888go/goframe/util/gutil"
)

func ExampleNew() {
	m := gmap.X创建()

	// Add data.
	m.X设置值("key1", "val1")

	// Print size.
	fmt.Println(m.X取数量())

	addMap := make(map[interface{}]interface{})
	addMap["key2"] = "val2"
	addMap["key3"] = "val3"
	addMap[1] = 1

	fmt.Println(m.X取所有值())

	// Batch add data.
	m.X设置值Map(addMap)

		// 获取对应键的值。 md5:80c18e88718ab20c
	fmt.Println(m.X取值("key3"))

		// 如果键不存在，使用给定的键值对获取或设置值。 md5:6110c826554187ae
	fmt.Println(m.X取值或设置值("key4", "val4"))

		// 如果键不存在，则设置键值对并返回true；否则，返回false。 md5:cc39f71c7fcd3e86
	fmt.Println(m.X设置值并跳过已存在("key3", "val3"))

	// Remove key
	m.X删除("key2")
	fmt.Println(m.X取所有名称())

	// Batch remove keys.
	m.X删除多个值([]interface{}{"key1", 1})
	fmt.Println(m.X取所有名称())

		// Contains 检查一个键是否存在。 md5:822d7ed9848c03c4
	fmt.Println(m.X是否存在("key3"))

		// Flip 会交换映射中的键值，它会将键值对改为值键对。 md5:23981f3e2b855f36
	m.X名称值交换()
	fmt.Println(m.X取Map())

		// Clear 删除map中的所有数据。 md5:7812c72bbfe807df
	m.X清空()

	fmt.Println(m.X取数量())

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
	m := gmap.X创建()

	m.X设置值("key1", "val1")
	fmt.Println(m)

	n := gmap.X创建并从Map(m.X浅拷贝(), true)
	fmt.Println(n)

	// Output:
	// {"key1":"val1"}
	// {"key1":"val1"}
}

func ExampleNewHashMap() {
	m := gmap.NewHashMap别名()

	m.X设置值("key1", "val1")
	fmt.Println(m)

	// Output:
	// {"key1":"val1"}
}

func ExampleNewHashMapFrom() {
	m := gmap.X创建()

	m.X设置值("key1", "val1")
	fmt.Println(m)

	n := gmap.NewHashMapFrom别名(m.X浅拷贝(), true)
	fmt.Println(n)

	// Output:
	// {"key1":"val1"}
	// {"key1":"val1"}
}

func ExampleNewAnyAnyMap() {
	m := gmap.X创建AnyAny()

	m.X设置值("key1", "val1")
	fmt.Println(m)

	// Output:
	// {"key1":"val1"}
}

func ExampleNewAnyAnyMapFrom() {
	m := gmap.X创建AnyAny()

	m.X设置值("key1", "val1")
	fmt.Println(m)

	n := gmap.X创建AnyAny并从Map(m.X浅拷贝(), true)
	fmt.Println(n)

	// Output:
	// {"key1":"val1"}
	// {"key1":"val1"}
}

func ExampleNewIntAnyMap() {
	m := gmap.X创建IntAny()

	m.X设置值(1, "val1")
	fmt.Println(m)

	// Output:
	// {"1":"val1"}
}

func ExampleNewIntAnyMapFrom() {
	m := gmap.X创建IntAny()

	m.X设置值(1, "val1")
	fmt.Println(m)

	n := gmap.X创建IntAny并从Map(m.X浅拷贝(), true)
	fmt.Println(n)

	// Output:
	// {"1":"val1"}
	// {"1":"val1"}
}

func ExampleNewIntIntMap() {
	m := gmap.X创建IntInt()

	m.X设置值(1, 1)
	fmt.Println(m)

	// Output:
	// {"1":1}
}

func ExampleNewIntIntMapFrom() {
	m := gmap.X创建IntInt()

	m.X设置值(1, 1)
	fmt.Println(m)

	n := gmap.X创建IntInt并从Map(m.X浅拷贝(), true)
	fmt.Println(n)

	// Output:
	// {"1":1}
	// {"1":1}
}

func ExampleNewStrAnyMap() {
	m := gmap.X创建StrAny()

	m.X设置值("key1", "var1")
	fmt.Println(m)

	// Output:
	// {"key1":"var1"}
}

func ExampleNewStrAnyMapFrom() {
	m := gmap.X创建StrAny()

	m.X设置值("key1", "var1")
	fmt.Println(m)

	n := gmap.X创建AnyStr并从Map(m.X浅拷贝(), true)
	fmt.Println(n)

	// Output:
	// {"key1":"var1"}
	// {"key1":"var1"}
}

func ExampleNewStrIntMap() {
	m := gmap.X创建StrInt()

	m.X设置值("key1", 1)
	fmt.Println(m)

	// Output:
	// {"key1":1}
}

func ExampleNewStrIntMapFrom() {
	m := gmap.X创建StrInt()

	m.X设置值("key1", 1)
	fmt.Println(m)

	n := gmap.X创建StrInt并从Map(m.X浅拷贝(), true)
	fmt.Println(n)

	// Output:
	// {"key1":1}
	// {"key1":1}
}

func ExampleNewStrStrMap() {
	m := gmap.X创建StrStr()

	m.X设置值("key1", "var1")
	fmt.Println(m)

	// Output:
	// {"key1":"var1"}
}

func ExampleNewStrStrMapFrom() {
	m := gmap.X创建StrStr()

	m.X设置值("key1", "var1")
	fmt.Println(m)

	n := gmap.X创建StrStr并从Map(m.X浅拷贝(), true)
	fmt.Println(n)

	// Output:
	// {"key1":"var1"}
	// {"key1":"var1"}
}

func ExampleNewListMap() {
	m := gmap.X创建链表mp()

	m.X设置值("key1", "var1")
	m.X设置值("key2", "var2")
	fmt.Println(m)

	// Output:
	// {"key1":"var1","key2":"var2"}
}

func ExampleNewListMapFrom() {
	m := gmap.X创建链表mp()

	m.X设置值("key1", "var1")
	m.X设置值("key2", "var2")
	fmt.Println(m)

	n := gmap.X创建链表Map并从Map(m.X取Map(), true)
	fmt.Println(n)

	// May Output:
	// {"key1":"var1","key2":"var2"}
	// {"key1":"var1","key2":"var2"}
}

func ExampleNewTreeMap() {
	m := gmap.X创建红黑树Map(gutil.X比较文本)

	m.X设置值("key2", "var2")
	m.X设置值("key1", "var1")

	fmt.Println(m.Map())

	// May Output:
	// map[key1:var1 key2:var2]
}

func ExampleNewTreeMapFrom() {
	m := gmap.X创建红黑树Map(gutil.X比较文本)

	m.X设置值("key2", "var2")
	m.X设置值("key1", "var1")

	fmt.Println(m.Map())

	n := gmap.X创建链表Map并从Map(m.Map(), true)
	fmt.Println(n.X取Map())

	// May Output:
	// map[key1:var1 key2:var2]
	// map[key1:var1 key2:var2]
}
