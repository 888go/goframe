// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/Agogf/gf获取一个。

package gtree_test
import (
	"fmt"
	
	"github.com/888go/goframe/container/gtree"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
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
