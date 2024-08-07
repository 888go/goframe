// 版权归GoFrame作者所有（https://goframe.org）。保留所有权利。
//
// 本源代码形式受MIT许可证条款的约束。如果gm文件中未附带MIT许可证的副本，
// 您可以从https://github.com/gogf/gf获取。
// md5:1d281c30cdc3423b

package 集合类_test

import (
	"fmt"

	gset "github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/frame/g"
)

func ExampleSet_Intersect() {
	s1 := gset.X创建并按值(g.Slice别名{1, 2, 3})
	s2 := gset.X创建并按值(g.Slice别名{4, 5, 6})
	s3 := gset.X创建并按值(g.Slice别名{1, 2, 3, 4, 5, 6, 7})

	fmt.Println(s3.X取交集(s1).X取集合切片())
	fmt.Println(s3.X取差集(s1).X取集合切片())
	fmt.Println(s1.X取并集(s2).X取集合切片())
	fmt.Println(s1.X取补集(s3).X取集合切片())

	// May Output:
	// [2 3 1]
	// [5 6 7 4]
	// [6 1 2 3 4 5]
	// [4 5 6 7]
}

func ExampleSet_Diff() {
	s1 := gset.X创建并按值(g.Slice别名{1, 2, 3})
	s2 := gset.X创建并按值(g.Slice别名{4, 5, 6})
	s3 := gset.X创建并按值(g.Slice别名{1, 2, 3, 4, 5, 6, 7})

	fmt.Println(s3.X取交集(s1).X取集合切片())
	fmt.Println(s3.X取差集(s1).X取集合切片())
	fmt.Println(s1.X取并集(s2).X取集合切片())
	fmt.Println(s1.X取补集(s3).X取集合切片())

	// May Output:
	// [2 3 1]
	// [5 6 7 4]
	// [6 1 2 3 4 5]
	// [4 5 6 7]
}

func ExampleSet_Union() {
	s1 := gset.X创建并按值(g.Slice别名{1, 2, 3})
	s2 := gset.X创建并按值(g.Slice别名{4, 5, 6})
	s3 := gset.X创建并按值(g.Slice别名{1, 2, 3, 4, 5, 6, 7})

	fmt.Println(s3.X取交集(s1).X取集合切片())
	fmt.Println(s3.X取差集(s1).X取集合切片())
	fmt.Println(s1.X取并集(s2).X取集合切片())
	fmt.Println(s1.X取补集(s3).X取集合切片())

	// May Output:
	// [2 3 1]
	// [5 6 7 4]
	// [6 1 2 3 4 5]
	// [4 5 6 7]
}

func ExampleSet_Complement() {
	s1 := gset.X创建并按值(g.Slice别名{1, 2, 3})
	s2 := gset.X创建并按值(g.Slice别名{4, 5, 6})
	s3 := gset.X创建并按值(g.Slice别名{1, 2, 3, 4, 5, 6, 7})

	fmt.Println(s3.X取交集(s1).X取集合切片())
	fmt.Println(s3.X取差集(s1).X取集合切片())
	fmt.Println(s1.X取并集(s2).X取集合切片())
	fmt.Println(s1.X取补集(s3).X取集合切片())

	// May Output:
	// [2 3 1]
	// [5 6 7 4]
	// [6 1 2 3 4 5]
	// [4 5 6 7]
}

func ExampleSet_IsSubsetOf() {
	var s1, s2 gset.Set
	s1.X加入(g.Slice别名{1, 2, 3}...)
	s2.X加入(g.Slice别名{2, 3}...)
	fmt.Println(s1.X是否为子集(&s2))
	fmt.Println(s2.X是否为子集(&s1))

	// Output:
	// false
	// true
}

func ExampleSet_AddIfNotExist() {
	var set gset.Set
	fmt.Println(set.X加入值并跳过已存在(1))
	fmt.Println(set.X加入值并跳过已存在(1))
	fmt.Println(set.X取集合切片())

	// Output:
	// true
	// false
	// [1]
}

func ExampleSet_Pop() {
	var set gset.Set
	set.X加入(1, 2, 3, 4)
	fmt.Println(set.X出栈())
	fmt.Println(set.X出栈多个(2))
	fmt.Println(set.X取数量())

	// May Output:
	// 1
	// [2 3]
	// 1
}

func ExampleSet_Pops() {
	var set gset.Set
	set.X加入(1, 2, 3, 4)
	fmt.Println(set.X出栈())
	fmt.Println(set.X出栈多个(2))
	fmt.Println(set.X取数量())

	// May Output:
	// 1
	// [2 3]
	// 1
}

func ExampleSet_Join() {
	var set gset.Set
	set.X加入("a", "b", "c", "d")
	fmt.Println(set.X取集合文本(","))

	// May Output:
	// a,b,c,d
}

func ExampleSet_Contains() {
	var set gset.StrSet
	set.X加入("a")
	fmt.Println(set.X是否存在("a"))
	fmt.Println(set.X是否存在("A"))
	fmt.Println(set.X是否存在并忽略大小写("A"))

	// Output:
	// true
	// false
	// true
}

func ExampleSet_ContainsI() {
	var set gset.StrSet
	set.X加入("a")
	fmt.Println(set.X是否存在("a"))
	fmt.Println(set.X是否存在("A"))
	fmt.Println(set.X是否存在并忽略大小写("A"))

	// Output:
	// true
	// false
	// true
}
