// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

package map类_test

import (
	"fmt"
	
	"github.com/888go/goframe/gmap"
	"github.com/gogf/gf/v2/util/gutil"
)

func ExampleNew() {
	m := map类.X创建()

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

	// 获取相应键的值。
	fmt.Println(m.X取值("key3"))

	// 通过键获取值，如果不存在则使用给定的键值对设置它。
	fmt.Println(m.X取值或设置值("key4", "val4"))

	// 如果键不存在，则设置键值对并返回true；否则返回false。
	fmt.Println(m.X设置值并跳过已存在("key3", "val3"))

	// Remove key
	m.X删除("key2")
	fmt.Println(m.X取所有名称())

	// 批量删除键。
	m.X删除多个值([]interface{}{"key1", 1})
	fmt.Println(m.X取所有名称())

	// Contains 检查键是否存在。
	fmt.Println(m.X是否存在("key3"))

	// Flip交换映射中的键值对，它会将键值对改为值键对。
	m.X名称值交换()
	fmt.Println(m.X取Map())

	// 清空删除映射中的所有数据。
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
	m := map类.X创建()

	m.X设置值("key1", "val1")
	fmt.Println(m)

	n := map类.X创建并从Map(m.X浅拷贝(), true)
	fmt.Println(n)

	// Output:
	// {"key1":"val1"}
	// {"key1":"val1"}
}

func ExampleNewHashMap() {
	m := map类.NewHashMap别名()

	m.X设置值("key1", "val1")
	fmt.Println(m)

	// Output:
	// {"key1":"val1"}
}

func ExampleNewHashMapFrom() {
	m := map类.X创建()

	m.X设置值("key1", "val1")
	fmt.Println(m)

	n := map类.NewHashMapFrom别名(m.X浅拷贝(), true)
	fmt.Println(n)

	// Output:
	// {"key1":"val1"}
	// {"key1":"val1"}
}

func ExampleNewAnyAnyMap() {
	m := map类.X创建AnyAny()

	m.X设置值("key1", "val1")
	fmt.Println(m)

	// Output:
	// {"key1":"val1"}
}

func ExampleNewAnyAnyMapFrom() {
	m := map类.X创建AnyAny()

	m.X设置值("key1", "val1")
	fmt.Println(m)

	n := map类.X创建AnyAny并从Map(m.X浅拷贝(), true)
	fmt.Println(n)

	// Output:
	// {"key1":"val1"}
	// {"key1":"val1"}
}

func ExampleNewIntAnyMap() {
	m := map类.X创建IntAny()

	m.X设置值(1, "val1")
	fmt.Println(m)

	// Output:
	// {"1":"val1"}
}

func ExampleNewIntAnyMapFrom() {
	m := map类.X创建IntAny()

	m.X设置值(1, "val1")
	fmt.Println(m)

	n := map类.X创建IntAny并从Map(m.X浅拷贝(), true)
	fmt.Println(n)

	// Output:
	// {"1":"val1"}
	// {"1":"val1"}
}

func ExampleNewIntIntMap() {
	m := map类.X创建IntInt()

	m.X设置值(1, 1)
	fmt.Println(m)

	// Output:
	// {"1":1}
}

func ExampleNewIntIntMapFrom() {
	m := map类.X创建IntInt()

	m.X设置值(1, 1)
	fmt.Println(m)

	n := map类.X创建IntInt并从Map(m.X浅拷贝(), true)
	fmt.Println(n)

	// Output:
	// {"1":1}
	// {"1":1}
}

func ExampleNewStrAnyMap() {
	m := map类.X创建StrAny()

	m.X设置值("key1", "var1")
	fmt.Println(m)

	// Output:
	// {"key1":"var1"}
}

func ExampleNewStrAnyMapFrom() {
	m := map类.X创建StrAny()

	m.X设置值("key1", "var1")
	fmt.Println(m)

	n := map类.X创建AnyStr并从Map(m.X浅拷贝(), true)
	fmt.Println(n)

	// Output:
	// {"key1":"var1"}
	// {"key1":"var1"}
}

func ExampleNewStrIntMap() {
	m := map类.X创建StrInt()

	m.X设置值("key1", 1)
	fmt.Println(m)

	// Output:
	// {"key1":1}
}

func ExampleNewStrIntMapFrom() {
	m := map类.X创建StrInt()

	m.X设置值("key1", 1)
	fmt.Println(m)

	n := map类.X创建StrInt并从Map(m.X浅拷贝(), true)
	fmt.Println(n)

	// Output:
	// {"key1":1}
	// {"key1":1}
}

func ExampleNewStrStrMap() {
	m := map类.X创建StrStr()

	m.X设置值("key1", "var1")
	fmt.Println(m)

	// Output:
	// {"key1":"var1"}
}

func ExampleNewStrStrMapFrom() {
	m := map类.X创建StrStr()

	m.X设置值("key1", "var1")
	fmt.Println(m)

	n := map类.X创建StrStr并从Map(m.X浅拷贝(), true)
	fmt.Println(n)

	// Output:
	// {"key1":"var1"}
	// {"key1":"var1"}
}

func ExampleNewListMap() {
	m := map类.X创建链表mp()

	m.X设置值("key1", "var1")
	m.X设置值("key2", "var2")
	fmt.Println(m)

	// Output:
	// {"key1":"var1","key2":"var2"}
}

func ExampleNewListMapFrom() {
	m := map类.X创建链表mp()

	m.X设置值("key1", "var1")
	m.X设置值("key2", "var2")
	fmt.Println(m)

	n := map类.X创建链表Map并从Map(m.X取Map(), true)
	fmt.Println(n)

	// May Output:
	// {"key1":"var1","key2":"var2"}
	// {"key1":"var1","key2":"var2"}
}

func ExampleNewTreeMap() {
	m := map类.X创建红黑树Map(gutil.ComparatorString)

	m.Set("key2", "var2")
	m.Set("key1", "var1")

	fmt.Println(m.Map())

	// May Output:
	// map[key1:var1 key2:var2]
}

func ExampleNewTreeMapFrom() {
	m := map类.X创建红黑树Map(gutil.ComparatorString)

	m.Set("key2", "var2")
	m.Set("key1", "var1")

	fmt.Println(m.Map())

	n := map类.X创建链表Map并从Map(m.Map(), true)
	fmt.Println(n.X取Map())

	// May Output:
	// map[key1:var1 key2:var2]
	// map[key1:var1 key2:var2]
}
