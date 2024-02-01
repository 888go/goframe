// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可协议条款。如果随gm文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf获取一个。

package gset_test
import (
	"fmt"
	
	"github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/frame/g"
	)

func ExampleSet_Intersect() {
	s1 := gset.NewFrom(g.Slice{1, 2, 3})
	s2 := gset.NewFrom(g.Slice{4, 5, 6})
	s3 := gset.NewFrom(g.Slice{1, 2, 3, 4, 5, 6, 7})

	fmt.Println(s3.Intersect(s1).Slice())
	fmt.Println(s3.Diff(s1).Slice())
	fmt.Println(s1.Union(s2).Slice())
	fmt.Println(s1.Complement(s3).Slice())

	// May Output:
	// [2 3 1]
	// [5 6 7 4]
	// [6 1 2 3 4 5]
	// [4 5 6 7]
}

func ExampleSet_Diff() {
	s1 := gset.NewFrom(g.Slice{1, 2, 3})
	s2 := gset.NewFrom(g.Slice{4, 5, 6})
	s3 := gset.NewFrom(g.Slice{1, 2, 3, 4, 5, 6, 7})

	fmt.Println(s3.Intersect(s1).Slice())
	fmt.Println(s3.Diff(s1).Slice())
	fmt.Println(s1.Union(s2).Slice())
	fmt.Println(s1.Complement(s3).Slice())

	// May Output:
	// [2 3 1]
	// [5 6 7 4]
	// [6 1 2 3 4 5]
	// [4 5 6 7]
}

func ExampleSet_Union() {
	s1 := gset.NewFrom(g.Slice{1, 2, 3})
	s2 := gset.NewFrom(g.Slice{4, 5, 6})
	s3 := gset.NewFrom(g.Slice{1, 2, 3, 4, 5, 6, 7})

	fmt.Println(s3.Intersect(s1).Slice())
	fmt.Println(s3.Diff(s1).Slice())
	fmt.Println(s1.Union(s2).Slice())
	fmt.Println(s1.Complement(s3).Slice())

	// May Output:
	// [2 3 1]
	// [5 6 7 4]
	// [6 1 2 3 4 5]
	// [4 5 6 7]
}

func ExampleSet_Complement() {
	s1 := gset.NewFrom(g.Slice{1, 2, 3})
	s2 := gset.NewFrom(g.Slice{4, 5, 6})
	s3 := gset.NewFrom(g.Slice{1, 2, 3, 4, 5, 6, 7})

	fmt.Println(s3.Intersect(s1).Slice())
	fmt.Println(s3.Diff(s1).Slice())
	fmt.Println(s1.Union(s2).Slice())
	fmt.Println(s1.Complement(s3).Slice())

	// May Output:
	// [2 3 1]
	// [5 6 7 4]
	// [6 1 2 3 4 5]
	// [4 5 6 7]
}

func ExampleSet_IsSubsetOf() {
	var s1, s2 gset.Set
	s1.Add(g.Slice{1, 2, 3}...)
	s2.Add(g.Slice{2, 3}...)
	fmt.Println(s1.IsSubsetOf(&s2))
	fmt.Println(s2.IsSubsetOf(&s1))

	// Output:
	// false
	// true
}

func ExampleSet_AddIfNotExist() {
	var set gset.Set
	fmt.Println(set.AddIfNotExist(1))
	fmt.Println(set.AddIfNotExist(1))
	fmt.Println(set.Slice())

	// Output:
	// true
	// false
	// [1]
}

func ExampleSet_Pop() {
	var set gset.Set
	set.Add(1, 2, 3, 4)
	fmt.Println(set.Pop())
	fmt.Println(set.Pops(2))
	fmt.Println(set.Size())

	// May Output:
	// 1
	// [2 3]
	// 1
}

func ExampleSet_Pops() {
	var set gset.Set
	set.Add(1, 2, 3, 4)
	fmt.Println(set.Pop())
	fmt.Println(set.Pops(2))
	fmt.Println(set.Size())

	// May Output:
	// 1
	// [2 3]
	// 1
}

func ExampleSet_Join() {
	var set gset.Set
	set.Add("a", "b", "c", "d")
	fmt.Println(set.Join(","))

	// May Output:
	// a,b,c,d
}

func ExampleSet_Contains() {
	var set gset.StrSet
	set.Add("a")
	fmt.Println(set.Contains("a"))
	fmt.Println(set.Contains("A"))
	fmt.Println(set.ContainsI("A"))

	// Output:
	// true
	// false
	// true
}

func ExampleSet_ContainsI() {
	var set gset.StrSet
	set.Add("a")
	fmt.Println(set.Contains("a"))
	fmt.Println(set.Contains("A"))
	fmt.Println(set.ContainsI("A"))

	// Output:
	// true
	// false
	// true
}
