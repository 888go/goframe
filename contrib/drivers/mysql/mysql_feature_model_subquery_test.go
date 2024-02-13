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

func Test_Model_SubQuery_Where(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
		subQuery1 := db.X创建Model对象(table).X条件("id", g.Slice别名{1, 3, 5})
		subQuery2 := db.X创建Model对象(table).X条件("id", g.Slice别名{5, 7, 9})
		r, err := db.X创建Model对象("? AS a, ? AS b", subQuery1, subQuery2).X字段保留过滤("a.id").X条件("a.id=b.id").X排序ASC("id").X查询()
		t.AssertNil(err)

		t.Assert(len(r), 1)
		t.Assert(r[0]["id"], 5)
	})
}
