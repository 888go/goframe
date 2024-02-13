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

	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := db.X多表去重查询(
			db.X创建Model对象(table).X条件("id", 1),
			db.X创建Model对象(table).X条件("id", 2),
			db.X创建Model对象(table).X条件包含("id", g.Slice别名{1, 2, 3}).X排序Desc("id"),
		).X排序Desc("id").X查询()

		t.AssertNil(err)

		t.Assert(len(r), 3)
		t.Assert(r[0]["id"], 3)
		t.Assert(r[1]["id"], 2)
		t.Assert(r[2]["id"], 1)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := db.X多表去重查询(
			db.X创建Model对象(table).X条件("id", 1),
			db.X创建Model对象(table).X条件("id", 2),
			db.X创建Model对象(table).X条件包含("id", g.Slice别名{1, 2, 3}).X排序Desc("id"),
		).X排序Desc("id").X查询一条()

		t.AssertNil(err)

		t.Assert(r["id"], 3)
	})
}

func Test_UnionAll(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := db.X多表查询(
			db.X创建Model对象(table).X条件("id", 1),
			db.X创建Model对象(table).X条件("id", 2),
			db.X创建Model对象(table).X条件包含("id", g.Slice别名{1, 2, 3}).X排序Desc("id"),
		).X排序Desc("id").X查询()

		t.AssertNil(err)

		t.Assert(len(r), 5)
		t.Assert(r[0]["id"], 3)
		t.Assert(r[1]["id"], 2)
		t.Assert(r[2]["id"], 2)
		t.Assert(r[3]["id"], 1)
		t.Assert(r[4]["id"], 1)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := db.X多表查询(
			db.X创建Model对象(table).X条件("id", 1),
			db.X创建Model对象(table).X条件("id", 2),
			db.X创建Model对象(table).X条件包含("id", g.Slice别名{1, 2, 3}).X排序Desc("id"),
		).X排序Desc("id").X查询一条()

		t.AssertNil(err)

		t.Assert(r["id"], 3)
	})
}

func Test_Model_Union(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := db.X创建Model对象(table).X多表去重查询(
			db.X创建Model对象(table).X条件("id", 1),
			db.X创建Model对象(table).X条件("id", 2),
			db.X创建Model对象(table).X条件包含("id", g.Slice别名{1, 2, 3}).X排序Desc("id"),
		).X排序Desc("id").X查询()

		t.AssertNil(err)

		t.Assert(len(r), 3)
		t.Assert(r[0]["id"], 3)
		t.Assert(r[1]["id"], 2)
		t.Assert(r[2]["id"], 1)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := db.X创建Model对象(table).X多表去重查询(
			db.X创建Model对象(table).X条件("id", 1),
			db.X创建Model对象(table).X条件("id", 2),
			db.X创建Model对象(table).X条件包含("id", g.Slice别名{1, 2, 3}).X排序Desc("id"),
		).X排序Desc("id").X查询一条()

		t.AssertNil(err)

		t.Assert(r["id"], 3)
	})
}

func Test_Model_UnionAll(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := db.X创建Model对象(table).X多表查询(
			db.X创建Model对象(table).X条件("id", 1),
			db.X创建Model对象(table).X条件("id", 2),
			db.X创建Model对象(table).X条件包含("id", g.Slice别名{1, 2, 3}).X排序Desc("id"),
		).X排序Desc("id").X查询()

		t.AssertNil(err)

		t.Assert(len(r), 5)
		t.Assert(r[0]["id"], 3)
		t.Assert(r[1]["id"], 2)
		t.Assert(r[2]["id"], 2)
		t.Assert(r[3]["id"], 1)
		t.Assert(r[4]["id"], 1)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := db.X创建Model对象(table).X多表查询(
			db.X创建Model对象(table).X条件("id", 1),
			db.X创建Model对象(table).X条件("id", 2),
			db.X创建Model对象(table).X条件包含("id", g.Slice别名{1, 2, 3}).X排序Desc("id"),
		).X排序Desc("id").X查询一条()

		t.AssertNil(err)

		t.Assert(r["id"], 3)
	})
}
