// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果没有随 gm 文件一起分布 MIT 许可证的副本，
// 您可以访问 https://github.com/Agogf/gf 获取一个。 md5:8fae8e64a457a737

package gtree_test

import (
	"fmt"

	"github.com/gogf/gf/v2/container/gtree"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

func ExampleNewAVLTree() {
	avlTree := gtree.NewAVLTree(gutil.ComparatorString)
	for i := 0; i < 6; i++ {
		avlTree.Set("key"+gconv.String(i), "val"+gconv.String(i))
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
	avlTree := gtree.NewAVLTree(gutil.ComparatorString)
	for i := 0; i < 6; i++ {
		avlTree.Set("key"+gconv.String(i), "val"+gconv.String(i))
	}

	otherAvlTree := gtree.NewAVLTreeFrom(gutil.ComparatorString, avlTree.Map())
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
	bTree := gtree.NewBTree(3, gutil.ComparatorString)
	for i := 0; i < 6; i++ {
		bTree.Set("key"+gconv.String(i), "val"+gconv.String(i))
	}
	fmt.Println(bTree.Map())

	// Output:
	// map[key0:val0 key1:val1 key2:val2 key3:val3 key4:val4 key5:val5]
}

func ExampleNewBTreeFrom() {
	bTree := gtree.NewBTree(3, gutil.ComparatorString)
	for i := 0; i < 6; i++ {
		bTree.Set("key"+gconv.String(i), "val"+gconv.String(i))
	}

	otherBTree := gtree.NewBTreeFrom(3, gutil.ComparatorString, bTree.Map())
	fmt.Println(otherBTree.Map())

	// Output:
	// map[key0:val0 key1:val1 key2:val2 key3:val3 key4:val4 key5:val5]
}

func ExampleNewRedBlackTree() {
	rbTree := gtree.NewRedBlackTree(gutil.ComparatorString)
	for i := 0; i < 6; i++ {
		rbTree.Set("key"+gconv.String(i), "val"+gconv.String(i))
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
	rbTree := gtree.NewRedBlackTree(gutil.ComparatorString)
	for i := 0; i < 6; i++ {
		rbTree.Set("key"+gconv.String(i), "val"+gconv.String(i))
	}

	otherRBTree := gtree.NewRedBlackTreeFrom(gutil.ComparatorString, rbTree.Map())
	fmt.Println(otherRBTree)

	// May Output:
	// │           ┌── key5
	// │       ┌── key4
	// │   ┌── key3
	// │   │   └── key2
	// └── key1
	//     └── key0
}
