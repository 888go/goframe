// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/Agogf/gf获取一个。

package 树形类_test

import (
	"fmt"
	
	"github.com/888go/goframe/container/gtree"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
)

func ExampleBTree_Clone() {
	b := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		b.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	tree := b.Clone()

	fmt.Println(tree.Map())
	fmt.Println(tree.Size())

	// Output:
	// map[key0:val0 key1:val1 key2:val2 key3:val3 key4:val4 key5:val5]
	// 6
}

func ExampleBTree_Set() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.Map())
	fmt.Println(tree.Size())

	// Output:
	// map[key0:val0 key1:val1 key2:val2 key3:val3 key4:val4 key5:val5]
	// 6
}

func ExampleBTree_Sets() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)

	tree.Sets(map[interface{}]interface{}{
		"key1": "val1",
		"key2": "val2",
	})

	fmt.Println(tree.Map())
	fmt.Println(tree.Size())

	// Output:
	// map[key1:val1 key2:val2]
	// 2
}

func ExampleBTree_Get() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.Get("key1"))
	fmt.Println(tree.Get("key10"))

	// Output:
	// val1
	// <nil>
}

func ExampleBTree_GetOrSet() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.GetOrSet("key1", "newVal1"))
	fmt.Println(tree.GetOrSet("key6", "val6"))

	// Output:
	// val1
	// val6
}

func ExampleBTree_GetOrSetFunc() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.GetOrSetFunc("key1", func() interface{} {
		return "newVal1"
	}))
	fmt.Println(tree.GetOrSetFunc("key6", func() interface{} {
		return "val6"
	}))

	// Output:
	// val1
	// val6
}

func ExampleBTree_GetOrSetFuncLock() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.GetOrSetFuncLock("key1", func() interface{} {
		return "newVal1"
	}))
	fmt.Println(tree.GetOrSetFuncLock("key6", func() interface{} {
		return "val6"
	}))

	// Output:
	// val1
	// val6
}

func ExampleBTree_GetVar() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.GetVar("key1").String())

	// Output:
	// val1
}

func ExampleBTree_GetVarOrSet() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.GetVarOrSet("key1", "newVal1"))
	fmt.Println(tree.GetVarOrSet("key6", "val6"))

	// Output:
	// val1
	// val6
}

func ExampleBTree_GetVarOrSetFunc() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.GetVarOrSetFunc("key1", func() interface{} {
		return "newVal1"
	}))
	fmt.Println(tree.GetVarOrSetFunc("key6", func() interface{} {
		return "val6"
	}))

	// Output:
	// val1
	// val6
}

func ExampleBTree_GetVarOrSetFuncLock() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.GetVarOrSetFuncLock("key1", func() interface{} {
		return "newVal1"
	}))
	fmt.Println(tree.GetVarOrSetFuncLock("key6", func() interface{} {
		return "val6"
	}))

	// Output:
	// val1
	// val6
}

func ExampleBTree_SetIfNotExist() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.SetIfNotExist("key1", "newVal1"))
	fmt.Println(tree.SetIfNotExist("key6", "val6"))

	// Output:
	// false
	// true
}

func ExampleBTree_SetIfNotExistFunc() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.SetIfNotExistFunc("key1", func() interface{} {
		return "newVal1"
	}))
	fmt.Println(tree.SetIfNotExistFunc("key6", func() interface{} {
		return "val6"
	}))

	// Output:
	// false
	// true
}

func ExampleBTree_SetIfNotExistFuncLock() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.SetIfNotExistFuncLock("key1", func() interface{} {
		return "newVal1"
	}))
	fmt.Println(tree.SetIfNotExistFuncLock("key6", func() interface{} {
		return "val6"
	}))

	// Output:
	// false
	// true
}

func ExampleBTree_Contains() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.Contains("key1"))
	fmt.Println(tree.Contains("key6"))

	// Output:
	// true
	// false
}

func ExampleBTree_Remove() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.Remove("key1"))
	fmt.Println(tree.Remove("key6"))
	fmt.Println(tree.Map())

	// Output:
	// val1
	// <nil>
	// map[key0:val0 key2:val2 key3:val3 key4:val4 key5:val5]
}

func ExampleBTree_Removes() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	removeKeys := make([]interface{}, 2)
	removeKeys = append(removeKeys, "key1")
	removeKeys = append(removeKeys, "key6")

	tree.Removes(removeKeys)

	fmt.Println(tree.Map())

	// Output:
	// map[key0:val0 key2:val2 key3:val3 key4:val4 key5:val5]
}

func ExampleBTree_IsEmpty() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)

	fmt.Println(tree.IsEmpty())

	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.IsEmpty())

	// Output:
	// true
	// false
}

func ExampleBTree_Size() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)

	fmt.Println(tree.Size())

	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.Size())

	// Output:
	// 0
	// 6
}

func ExampleBTree_Keys() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 6; i > 0; i-- {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.Keys())

	// Output:
	// [key1 key2 key3 key4 key5 key6]
}

func ExampleBTree_Values() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 6; i > 0; i-- {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.Values())

	// Output:
	// [val1 val2 val3 val4 val5 val6]
}

func ExampleBTree_Map() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.Map())

	// Output:
	// map[key0:val0 key1:val1 key2:val2 key3:val3 key4:val4 key5:val5]
}

func ExampleBTree_MapStrAny() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值(1000+i, "val"+转换类.String(i))
	}

	fmt.Println(tree.X取MapStrAny())

	// Output:
	// map[1000:val0 1001:val1 1002:val2 1003:val3 1004:val4 1005:val5]
}

func ExampleBTree_Clear() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值(1000+i, "val"+转换类.String(i))
	}
	fmt.Println(tree.Size())

	tree.Clear()
	fmt.Println(tree.Size())

	// Output:
	// 6
	// 0
}

func ExampleBTree_Replace() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.Map())

	data := map[interface{}]interface{}{
		"newKey0": "newVal0",
		"newKey1": "newVal1",
		"newKey2": "newVal2",
	}

	tree.Replace(data)

	fmt.Println(tree.Map())

	// Output:
	// map[key0:val0 key1:val1 key2:val2 key3:val3 key4:val4 key5:val5]
	// map[newKey0:newVal0 newKey1:newVal1 newKey2:newVal2]
}

func ExampleBTree_Height() {
	tree := 树形类.NewBTree(3, 工具类.X比较整数)
	for i := 0; i < 100; i++ {
		tree.X设置值(i, i)
	}
	fmt.Println(tree.Height())

	// Output:
	// 6
}

func ExampleBTree_Left() {
	tree := 树形类.NewBTree(3, 工具类.X比较整数)
	for i := 1; i < 100; i++ {
		tree.X设置值(i, i)
	}
	fmt.Println(tree.Left().Key, tree.Left().Value)

	emptyTree := 树形类.NewBTree(3, 工具类.X比较整数)
	fmt.Println(emptyTree.Left())

	// Output:
	// 1 1
	// <nil>
}

func ExampleBTree_Right() {
	tree := 树形类.NewBTree(3, 工具类.X比较整数)
	for i := 1; i < 100; i++ {
		tree.X设置值(i, i)
	}
	fmt.Println(tree.Right().Key, tree.Right().Value)

	emptyTree := 树形类.NewBTree(3, 工具类.X比较整数)
	fmt.Println(emptyTree.Left())

	// Output:
	// 99 99
	// <nil>
}

func ExampleBTree_String() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.String())

	// Output:
	// key0
	// key1
	//     key2
	// key3
	//     key4
	//     key5
}

func ExampleBTree_Search() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(tree.Search("key0"))
	fmt.Println(tree.Search("key6"))

	// Output:
	// val0 true
	// <nil> false
}

func ExampleBTree_Print() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	tree.Print()

	// Output:
	// key0
	// key1
	//     key2
	// key3
	//     key4
	//     key5
}

func ExampleBTree_Iterator() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 10; i++ {
		tree.X设置值(i, 10-i)
	}

	var totalKey, totalValue int
	tree.X遍历(func(key, value interface{}) bool {
		totalKey += key.(int)
		totalValue += value.(int)

		return totalValue < 20
	})

	fmt.Println("totalKey:", totalKey)
	fmt.Println("totalValue:", totalValue)

	// Output:
	// totalKey: 3
	// totalValue: 27
}

func ExampleBTree_IteratorFrom() {
	m := make(map[interface{}]interface{})
	for i := 1; i <= 5; i++ {
		m[i] = i * 10
	}
	tree := 树形类.NewBTreeFrom(3, 工具类.X比较整数, m)

	tree.IteratorFrom(1, true, func(key, value interface{}) bool {
		fmt.Println("key:", key, ", value:", value)
		return true
	})

	// Output:
	// key: 1 , value: 10
	// key: 2 , value: 20
	// key: 3 , value: 30
	// key: 4 , value: 40
	// key: 5 , value: 50
}

func ExampleBTree_IteratorAsc() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 10; i++ {
		tree.X设置值(i, 10-i)
	}

	tree.IteratorAsc(func(key, value interface{}) bool {
		fmt.Println("key:", key, ", value:", value)
		return true
	})

	// Output:
	// key: 0 , value: 10
	// key: 1 , value: 9
	// key: 2 , value: 8
	// key: 3 , value: 7
	// key: 4 , value: 6
	// key: 5 , value: 5
	// key: 6 , value: 4
	// key: 7 , value: 3
	// key: 8 , value: 2
	// key: 9 , value: 1
}

func ExampleBTree_IteratorAscFrom_Normal() {
	m := make(map[interface{}]interface{})
	for i := 1; i <= 5; i++ {
		m[i] = i * 10
	}
	tree := 树形类.NewBTreeFrom(3, 工具类.X比较整数, m)

	tree.IteratorAscFrom(1, true, func(key, value interface{}) bool {
		fmt.Println("key:", key, ", value:", value)
		return true
	})

	// Output:
	// key: 1 , value: 10
	// key: 2 , value: 20
	// key: 3 , value: 30
	// key: 4 , value: 40
	// key: 5 , value: 50
}

func ExampleBTree_IteratorAscFrom_NoExistKey() {
	m := make(map[interface{}]interface{})
	for i := 1; i <= 5; i++ {
		m[i] = i * 10
	}
	tree := 树形类.NewBTreeFrom(3, 工具类.X比较整数, m)

	tree.IteratorAscFrom(0, true, func(key, value interface{}) bool {
		fmt.Println("key:", key, ", value:", value)
		return true
	})

	// Output:
}

func ExampleBTree_IteratorAscFrom_NoExistKeyAndMatchFalse() {
	m := make(map[interface{}]interface{})
	for i := 1; i <= 5; i++ {
		m[i] = i * 10
	}
	tree := 树形类.NewBTreeFrom(3, 工具类.X比较整数, m)

	tree.IteratorAscFrom(0, false, func(key, value interface{}) bool {
		fmt.Println("key:", key, ", value:", value)
		return true
	})

	// Output:
	// key: 1 , value: 10
	// key: 2 , value: 20
	// key: 3 , value: 30
	// key: 4 , value: 40
	// key: 5 , value: 50
}

func ExampleBTree_IteratorDesc() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 10; i++ {
		tree.X设置值(i, 10-i)
	}

	tree.IteratorDesc(func(key, value interface{}) bool {
		fmt.Println("key:", key, ", value:", value)
		return true
	})

	// Output:
	// key: 9 , value: 1
	// key: 8 , value: 2
	// key: 7 , value: 3
	// key: 6 , value: 4
	// key: 5 , value: 5
	// key: 4 , value: 6
	// key: 3 , value: 7
	// key: 2 , value: 8
	// key: 1 , value: 9
	// key: 0 , value: 10
}

func ExampleBTree_IteratorDescFrom() {
	m := make(map[interface{}]interface{})
	for i := 1; i <= 5; i++ {
		m[i] = i * 10
	}
	tree := 树形类.NewBTreeFrom(3, 工具类.X比较整数, m)

	tree.IteratorDescFrom(5, true, func(key, value interface{}) bool {
		fmt.Println("key:", key, ", value:", value)
		return true
	})

	// Output:
	// key: 5 , value: 50
	// key: 4 , value: 40
	// key: 3 , value: 30
	// key: 2 , value: 20
	// key: 1 , value: 10
}

func ExampleBTree_MarshalJSON() {
	tree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		tree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	bytes, err := json.Marshal(tree)
	if err == nil {
		fmt.Println(转换类.String(bytes))
	}

	// Output:
	// {"key0":"val0","key1":"val1","key2":"val2","key3":"val3","key4":"val4","key5":"val5"}
}
