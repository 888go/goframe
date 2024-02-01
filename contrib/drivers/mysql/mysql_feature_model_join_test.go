// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package mysql_test
import (
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	)

func Test_Model_LeftJoinOnField(t *testing.T) {
	var (
		table1 = gtime.TimestampNanoStr() + "_table1"
		table2 = gtime.TimestampNanoStr() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.Model(table1).
			FieldsPrefix(table1, "*").
			LeftJoinOnField(table2, "id").
			WhereIn("id", g.Slice{1, 2}).
			Order("id asc").All()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
}

func Test_Model_RightJoinOnField(t *testing.T) {
	var (
		table1 = gtime.TimestampNanoStr() + "_table1"
		table2 = gtime.TimestampNanoStr() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.Model(table1).
			FieldsPrefix(table1, "*").
			RightJoinOnField(table2, "id").
			WhereIn("id", g.Slice{1, 2}).
			Order("id asc").All()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
}

func Test_Model_InnerJoinOnField(t *testing.T) {
	var (
		table1 = gtime.TimestampNanoStr() + "_table1"
		table2 = gtime.TimestampNanoStr() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.Model(table1).
			FieldsPrefix(table1, "*").
			InnerJoinOnField(table2, "id").
			WhereIn("id", g.Slice{1, 2}).
			Order("id asc").All()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
}

func Test_Model_LeftJoinOnFields(t *testing.T) {
	var (
		table1 = gtime.TimestampNanoStr() + "_table1"
		table2 = gtime.TimestampNanoStr() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.Model(table1).
			FieldsPrefix(table1, "*").
			LeftJoinOnFields(table2, "id", "=", "id").
			WhereIn("id", g.Slice{1, 2}).
			Order("id asc").All()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
}

func Test_Model_RightJoinOnFields(t *testing.T) {
	var (
		table1 = gtime.TimestampNanoStr() + "_table1"
		table2 = gtime.TimestampNanoStr() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.Model(table1).
			FieldsPrefix(table1, "*").
			RightJoinOnFields(table2, "id", "=", "id").
			WhereIn("id", g.Slice{1, 2}).
			Order("id asc").All()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
}

func Test_Model_InnerJoinOnFields(t *testing.T) {
	var (
		table1 = gtime.TimestampNanoStr() + "_table1"
		table2 = gtime.TimestampNanoStr() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.Model(table1).
			FieldsPrefix(table1, "*").
			InnerJoinOnFields(table2, "id", "=", "id").
			WhereIn("id", g.Slice{1, 2}).
			Order("id asc").All()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
}

func Test_Model_FieldsPrefix(t *testing.T) {
	var (
		table1 = gtime.TimestampNanoStr() + "_table1"
		table2 = gtime.TimestampNanoStr() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.Model(table1).
			FieldsPrefix(table1, "id").
			FieldsPrefix(table2, "nickname").
			LeftJoinOnField(table2, "id").
			WhereIn("id", g.Slice{1, 2}).
			Order("id asc").All()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[0]["nickname"], "name_1")
	})
}
