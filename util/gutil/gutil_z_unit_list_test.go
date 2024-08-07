// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 工具类_test

import (
	"testing"

	"github.com/888go/goframe/frame/g"
	gtest "github.com/888go/goframe/test/gtest"
	gutil "github.com/888go/goframe/util/gutil"
)

func Test_ListItemValues_Map(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		listMap := g.Map切片{
			g.Map{"id": 1, "score": 100},
			g.Map{"id": 2, "score": 99},
			g.Map{"id": 3, "score": 99},
		}
		t.Assert(gutil.ListItemValues(listMap, "id"), g.Slice别名{1, 2, 3})
		t.Assert(gutil.ListItemValues(&listMap, "id"), g.Slice别名{1, 2, 3})
		t.Assert(gutil.ListItemValues(listMap, "score"), g.Slice别名{100, 99, 99})
	})
	gtest.C(t, func(t *gtest.T) {
		listMap := g.Map切片{
			g.Map{"id": 1, "score": 100},
			g.Map{"id": 2, "score": nil},
			g.Map{"id": 3, "score": 0},
		}
		t.Assert(gutil.ListItemValues(listMap, "id"), g.Slice别名{1, 2, 3})
		t.Assert(gutil.ListItemValues(listMap, "score"), g.Slice别名{100, nil, 0})
	})
	gtest.C(t, func(t *gtest.T) {
		listMap := g.Map切片{}
		t.Assert(len(gutil.ListItemValues(listMap, "id")), 0)
	})
}

func Test_ListItemValues_Map_SubKey(t *testing.T) {
	type Scores struct {
		Math    int
		English int
	}
	gtest.C(t, func(t *gtest.T) {
		listMap := g.Map切片{
			g.Map{"id": 1, "scores": Scores{100, 60}},
			g.Map{"id": 2, "scores": Scores{0, 100}},
			g.Map{"id": 3, "scores": Scores{59, 99}},
		}
		t.Assert(gutil.ListItemValues(listMap, "scores", "Math"), g.Slice别名{100, 0, 59})
		t.Assert(gutil.ListItemValues(listMap, "scores", "English"), g.Slice别名{60, 100, 99})
		t.Assert(gutil.ListItemValues(listMap, "scores", "PE"), g.Slice别名{})
	})
}

func Test_ListItemValues_Map_Array_SubKey(t *testing.T) {
	type Scores struct {
		Math    int
		English int
	}
	gtest.C(t, func(t *gtest.T) {
		listMap := g.Map切片{
			g.Map{"id": 1, "scores": []Scores{{1, 2}, {3, 4}}},
			g.Map{"id": 2, "scores": []Scores{{5, 6}, {7, 8}}},
			g.Map{"id": 3, "scores": []Scores{{9, 10}, {11, 12}}},
		}
		t.Assert(gutil.ListItemValues(listMap, "scores", "Math"), g.Slice别名{1, 3, 5, 7, 9, 11})
		t.Assert(gutil.ListItemValues(listMap, "scores", "English"), g.Slice别名{2, 4, 6, 8, 10, 12})
		t.Assert(gutil.ListItemValues(listMap, "scores", "PE"), g.Slice别名{})
	})
}

func Test_ListItemValues_Struct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type T struct {
			Id    int
			Score float64
		}
		listStruct := g.Slice别名{
			T{1, 100},
			T{2, 99},
			T{3, 0},
		}
		t.Assert(gutil.ListItemValues(listStruct, "Id"), g.Slice别名{1, 2, 3})
		t.Assert(gutil.ListItemValues(listStruct, "Score"), g.Slice别名{100, 99, 0})
	})
	// Pointer items.
	gtest.C(t, func(t *gtest.T) {
		type T struct {
			Id    int
			Score float64
		}
		listStruct := g.Slice别名{
			&T{1, 100},
			&T{2, 99},
			&T{3, 0},
		}
		t.Assert(gutil.ListItemValues(listStruct, "Id"), g.Slice别名{1, 2, 3})
		t.Assert(gutil.ListItemValues(listStruct, "Score"), g.Slice别名{100, 99, 0})
	})
	// Nil element value.
	gtest.C(t, func(t *gtest.T) {
		type T struct {
			Id    int
			Score interface{}
		}
		listStruct := g.Slice别名{
			T{1, 100},
			T{2, nil},
			T{3, 0},
		}
		t.Assert(gutil.ListItemValues(listStruct, "Id"), g.Slice别名{1, 2, 3})
		t.Assert(gutil.ListItemValues(listStruct, "Score"), g.Slice别名{100, nil, 0})
	})
}

func Test_ListItemValues_Struct_SubKey(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type Student struct {
			Id    int
			Score float64
		}
		type Class struct {
			Total    int
			Students []Student
		}
		listStruct := g.Slice别名{
			Class{2, []Student{{1, 1}, {2, 2}}},
			Class{3, []Student{{3, 3}, {4, 4}, {5, 5}}},
			Class{1, []Student{{6, 6}}},
		}
		t.Assert(gutil.ListItemValues(listStruct, "Total"), g.Slice别名{2, 3, 1})
		t.Assert(gutil.ListItemValues(listStruct, "Students"), `[[{"Id":1,"Score":1},{"Id":2,"Score":2}],[{"Id":3,"Score":3},{"Id":4,"Score":4},{"Id":5,"Score":5}],[{"Id":6,"Score":6}]]`)
		t.Assert(gutil.ListItemValues(listStruct, "Students", "Id"), g.Slice别名{1, 2, 3, 4, 5, 6})
	})
	gtest.C(t, func(t *gtest.T) {
		type Student struct {
			Id    int
			Score float64
		}
		type Class struct {
			Total    int
			Students []*Student
		}
		listStruct := g.Slice别名{
			&Class{2, []*Student{{1, 1}, {2, 2}}},
			&Class{3, []*Student{{3, 3}, {4, 4}, {5, 5}}},
			&Class{1, []*Student{{6, 6}}},
		}
		t.Assert(gutil.ListItemValues(listStruct, "Total"), g.Slice别名{2, 3, 1})
		t.Assert(gutil.ListItemValues(listStruct, "Students"), `[[{"Id":1,"Score":1},{"Id":2,"Score":2}],[{"Id":3,"Score":3},{"Id":4,"Score":4},{"Id":5,"Score":5}],[{"Id":6,"Score":6}]]`)
		t.Assert(gutil.ListItemValues(listStruct, "Students", "Id"), g.Slice别名{1, 2, 3, 4, 5, 6})
	})
}

func Test_ListItemValuesUnique(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		listMap := g.Map切片{
			g.Map{"id": 1, "score": 100},
			g.Map{"id": 2, "score": 100},
			g.Map{"id": 3, "score": 100},
			g.Map{"id": 4, "score": 100},
			g.Map{"id": 5, "score": 100},
		}
		t.Assert(gutil.ListItemValuesUnique(listMap, "id"), g.Slice别名{1, 2, 3, 4, 5})
		t.Assert(gutil.ListItemValuesUnique(listMap, "score"), g.Slice别名{100})
	})
	gtest.C(t, func(t *gtest.T) {
		listMap := g.Map切片{
			g.Map{"id": 1, "score": 100},
			g.Map{"id": 2, "score": 100},
			g.Map{"id": 3, "score": 100},
			g.Map{"id": 4, "score": 100},
			g.Map{"id": 5, "score": 99},
		}
		t.Assert(gutil.ListItemValuesUnique(listMap, "id"), g.Slice别名{1, 2, 3, 4, 5})
		t.Assert(gutil.ListItemValuesUnique(listMap, "score"), g.Slice别名{100, 99})
	})
	gtest.C(t, func(t *gtest.T) {
		listMap := g.Map切片{
			g.Map{"id": 1, "score": 100},
			g.Map{"id": 2, "score": 100},
			g.Map{"id": 3, "score": 0},
			g.Map{"id": 4, "score": 100},
			g.Map{"id": 5, "score": 99},
		}
		t.Assert(gutil.ListItemValuesUnique(listMap, "id"), g.Slice别名{1, 2, 3, 4, 5})
		t.Assert(gutil.ListItemValuesUnique(listMap, "score"), g.Slice别名{100, 0, 99})
	})
}

func Test_ListItemValuesUnique_Struct_SubKey(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type Student struct {
			Id    int
			Score float64
		}
		type Class struct {
			Total    int
			Students []Student
		}
		listStruct := g.Slice别名{
			Class{2, []Student{{1, 1}, {1, 2}}},
			Class{3, []Student{{2, 3}, {2, 4}, {5, 5}}},
			Class{1, []Student{{6, 6}}},
		}
		t.Assert(gutil.ListItemValuesUnique(listStruct, "Total"), g.Slice别名{2, 3, 1})
		t.Assert(gutil.ListItemValuesUnique(listStruct, "Students", "Id"), g.Slice别名{1, 2, 5, 6})
	})
	gtest.C(t, func(t *gtest.T) {
		type Student struct {
			Id    int
			Score float64
		}
		type Class struct {
			Total    int
			Students []*Student
		}
		listStruct := g.Slice别名{
			&Class{2, []*Student{{1, 1}, {1, 2}}},
			&Class{3, []*Student{{2, 3}, {2, 4}, {5, 5}}},
			&Class{1, []*Student{{6, 6}}},
		}
		t.Assert(gutil.ListItemValuesUnique(listStruct, "Total"), g.Slice别名{2, 3, 1})
		t.Assert(gutil.ListItemValuesUnique(listStruct, "Students", "Id"), g.Slice别名{1, 2, 5, 6})
	})
}

func Test_ListItemValuesUnique_Map_Array_SubKey(t *testing.T) {
	type Scores struct {
		Math    int
		English int
	}
	gtest.C(t, func(t *gtest.T) {
		listMap := g.Map切片{
			g.Map{"id": 1, "scores": []Scores{{1, 2}, {1, 2}}},
			g.Map{"id": 2, "scores": []Scores{{5, 8}, {5, 8}}},
			g.Map{"id": 3, "scores": []Scores{{9, 10}, {11, 12}}},
		}
		t.Assert(gutil.ListItemValuesUnique(listMap, "scores", "Math"), g.Slice别名{1, 5, 9, 11})
		t.Assert(gutil.ListItemValuesUnique(listMap, "scores", "English"), g.Slice别名{2, 8, 10, 12})
		t.Assert(gutil.ListItemValuesUnique(listMap, "scores", "PE"), g.Slice别名{})
	})
}

func Test_ListItemValuesUnique_Binary_ID(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		listMap := g.Map切片{
			g.Map{"id": []byte{1}, "score": 100},
			g.Map{"id": []byte{2}, "score": 100},
			g.Map{"id": []byte{3}, "score": 100},
			g.Map{"id": []byte{4}, "score": 100},
			g.Map{"id": []byte{4}, "score": 100},
		}
		t.Assert(gutil.ListItemValuesUnique(listMap, "id"), g.Slice别名{[]byte{1}, []byte{2}, []byte{3}, []byte{4}})
		t.Assert(gutil.ListItemValuesUnique(listMap, "score"), g.Slice别名{100})
	})
}
