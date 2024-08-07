// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package mysql_test

import (
	"testing"

	"github.com/888go/goframe/frame/g"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_Model_SubQuery_Where(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table).X条件(
			"id in ?",
			db.X创建Model对象(table).X字段保留过滤("id").X条件("id", g.Slice别名{1, 3, 5}),
		).X排序ASC("id").X查询()
		t.AssertNil(err)

		t.Assert(len(r), 3)
		t.Assert(r[0]["id"], 1)
		t.Assert(r[1]["id"], 3)
		t.Assert(r[2]["id"], 5)
	})
}

func Test_Model_SubQuery_Having(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table).X条件(
			"id in ?",
			db.X创建Model对象(table).X字段保留过滤("id").X条件("id", g.Slice别名{1, 3, 5}),
		).X设置分组条件(
			"id > ?",
			db.X创建Model对象(table).X字段保留过滤("MAX(id)").X条件("id", g.Slice别名{1, 3}),
		).X排序ASC("id").X查询()
		t.AssertNil(err)

		t.Assert(len(r), 1)
		t.Assert(r[0]["id"], 5)
	})
}

func Test_Model_SubQuery_Model(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		subQuery1 := db.X创建Model对象(table).X条件("id", g.Slice别名{1, 3, 5})
		subQuery2 := db.X创建Model对象(table).X条件("id", g.Slice别名{5, 7, 9})
		r, err := db.X创建Model对象("? AS a, ? AS b", subQuery1, subQuery2).X字段保留过滤("a.id").X条件("a.id=b.id").X排序ASC("id").X查询()
		t.AssertNil(err)

		t.Assert(len(r), 1)
		t.Assert(r[0]["id"], 5)
	})
}
