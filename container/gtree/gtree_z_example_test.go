// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/Agogf/gf获取一个。

package 树形类_test

import (
	"fmt"
	
	"github.com/888go/goframe/container/gtree"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
)

func ExampleNewAVLTree() {
	avlTree := 树形类.NewAVLTree(工具类.X比较文本)
	for i := 0; i < 6; i++ {
		avlTree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(avlTree)

	// Output:
	// │       ┌── key5
	// │   ┌── key4
	// └── key3
	//     │   ┌── key2
	//     └── key1
	//         └── key0
}

func ExampleNewAVLTreeFrom() {
	avlTree := 树形类.NewAVLTree(工具类.X比较文本)
	for i := 0; i < 6; i++ {
		avlTree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	otherAvlTree := 树形类.NewAVLTreeFrom(工具类.X比较文本, avlTree.Map())
	fmt.Println(otherAvlTree)

	// May Output:
	// │   ┌── key5
	// │   │   └── key4
	// └── key3
	//     │   ┌── key2
	//     └── key1
	//         └── key0
}

func ExampleNewBTree() {
	bTree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		bTree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}
	fmt.Println(bTree.Map())

	// Output:
	// map[key0:val0 key1:val1 key2:val2 key3:val3 key4:val4 key5:val5]
}

func ExampleNewBTreeFrom() {
	bTree := 树形类.NewBTree(3, 工具类.X比较文本)
	for i := 0; i < 6; i++ {
		bTree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	otherBTree := 树形类.NewBTreeFrom(3, 工具类.X比较文本, bTree.Map())
	fmt.Println(otherBTree.Map())

	// Output:
	// map[key0:val0 key1:val1 key2:val2 key3:val3 key4:val4 key5:val5]
}

func ExampleNewRedBlackTree() {
	rbTree := 树形类.NewRedBlackTree(工具类.X比较文本)
	for i := 0; i < 6; i++ {
		rbTree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	fmt.Println(rbTree)

	// Output:
	// │           ┌── key5
	// │       ┌── key4
	// │   ┌── key3
	// │   │   └── key2
	// └── key1
	//     └── key0
}

func ExampleNewRedBlackTreeFrom() {
	rbTree := 树形类.NewRedBlackTree(工具类.X比较文本)
	for i := 0; i < 6; i++ {
		rbTree.X设置值("key"+转换类.String(i), "val"+转换类.String(i))
	}

	otherRBTree := 树形类.NewRedBlackTreeFrom(工具类.X比较文本, rbTree.Map())
	fmt.Println(otherRBTree)

	// May Output:
	// │           ┌── key5
	// │       ┌── key4
	// │   ┌── key3
	// │   │   └── key2
	// └── key1
	//     └── key0
}
