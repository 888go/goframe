// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package mysql_test
import (
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
	)

func Test_Union(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.Union(
			db.Model(table).Where("id", 1),
			db.Model(table).Where("id", 2),
			db.Model(table).WhereIn("id", g.Slice{1, 2, 3}).OrderDesc("id"),
		).OrderDesc("id").All()

		t.AssertNil(err)

		t.Assert(len(r), 3)
		t.Assert(r[0]["id"], 3)
		t.Assert(r[1]["id"], 2)
		t.Assert(r[2]["id"], 1)
	})

	gtest.C(t, func(t *gtest.T) {
		r, err := db.Union(
			db.Model(table).Where("id", 1),
			db.Model(table).Where("id", 2),
			db.Model(table).WhereIn("id", g.Slice{1, 2, 3}).OrderDesc("id"),
		).OrderDesc("id").One()

		t.AssertNil(err)

		t.Assert(r["id"], 3)
	})
}

func Test_UnionAll(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.UnionAll(
			db.Model(table).Where("id", 1),
			db.Model(table).Where("id", 2),
			db.Model(table).WhereIn("id", g.Slice{1, 2, 3}).OrderDesc("id"),
		).OrderDesc("id").All()

		t.AssertNil(err)

		t.Assert(len(r), 5)
		t.Assert(r[0]["id"], 3)
		t.Assert(r[1]["id"], 2)
		t.Assert(r[2]["id"], 2)
		t.Assert(r[3]["id"], 1)
		t.Assert(r[4]["id"], 1)
	})

	gtest.C(t, func(t *gtest.T) {
		r, err := db.UnionAll(
			db.Model(table).Where("id", 1),
			db.Model(table).Where("id", 2),
			db.Model(table).WhereIn("id", g.Slice{1, 2, 3}).OrderDesc("id"),
		).OrderDesc("id").One()

		t.AssertNil(err)

		t.Assert(r["id"], 3)
	})
}

func Test_Model_Union(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.Model(table).Union(
			db.Model(table).Where("id", 1),
			db.Model(table).Where("id", 2),
			db.Model(table).WhereIn("id", g.Slice{1, 2, 3}).OrderDesc("id"),
		).OrderDesc("id").All()

		t.AssertNil(err)

		t.Assert(len(r), 3)
		t.Assert(r[0]["id"], 3)
		t.Assert(r[1]["id"], 2)
		t.Assert(r[2]["id"], 1)
	})

	gtest.C(t, func(t *gtest.T) {
		r, err := db.Model(table).Union(
			db.Model(table).Where("id", 1),
			db.Model(table).Where("id", 2),
			db.Model(table).WhereIn("id", g.Slice{1, 2, 3}).OrderDesc("id"),
		).OrderDesc("id").One()

		t.AssertNil(err)

		t.Assert(r["id"], 3)
	})
}

func Test_Model_UnionAll(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.Model(table).UnionAll(
			db.Model(table).Where("id", 1),
			db.Model(table).Where("id", 2),
			db.Model(table).WhereIn("id", g.Slice{1, 2, 3}).OrderDesc("id"),
		).OrderDesc("id").All()

		t.AssertNil(err)

		t.Assert(len(r), 5)
		t.Assert(r[0]["id"], 3)
		t.Assert(r[1]["id"], 2)
		t.Assert(r[2]["id"], 2)
		t.Assert(r[3]["id"], 1)
		t.Assert(r[4]["id"], 1)
	})

	gtest.C(t, func(t *gtest.T) {
		r, err := db.Model(table).UnionAll(
			db.Model(table).Where("id", 1),
			db.Model(table).Where("id", 2),
			db.Model(table).WhereIn("id", g.Slice{1, 2, 3}).OrderDesc("id"),
		).OrderDesc("id").One()

		t.AssertNil(err)

		t.Assert(r["id"], 3)
	})
}
