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
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_Model_LeftJoinOnField(t *testing.T) {
	var (
		table1 = gtime.X取文本时间戳纳秒() + "_table1"
		table2 = gtime.X取文本时间戳纳秒() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table1).
			X字段保留过滤并带前缀(table1, "*").
			X左连接相同字段(table2, "id").
			X条件包含("id", g.Slice别名{1, 2}).
			X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
}

func Test_Model_RightJoinOnField(t *testing.T) {
	var (
		table1 = gtime.X取文本时间戳纳秒() + "_table1"
		table2 = gtime.X取文本时间戳纳秒() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table1).
			X字段保留过滤并带前缀(table1, "*").
			X右连接相同字段(table2, "id").
			X条件包含("id", g.Slice别名{1, 2}).
			X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
}

func Test_Model_InnerJoinOnField(t *testing.T) {
	var (
		table1 = gtime.X取文本时间戳纳秒() + "_table1"
		table2 = gtime.X取文本时间戳纳秒() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table1).
			X字段保留过滤并带前缀(table1, "*").
			X内连接相同字段(table2, "id").
			X条件包含("id", g.Slice别名{1, 2}).
			X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
}

func Test_Model_LeftJoinOnFields(t *testing.T) {
	var (
		table1 = gtime.X取文本时间戳纳秒() + "_table1"
		table2 = gtime.X取文本时间戳纳秒() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table1).
			X字段保留过滤并带前缀(table1, "*").
			X左连接带比较运算符(table2, "id", "=", "id").
			X条件包含("id", g.Slice别名{1, 2}).
			X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
}

func Test_Model_RightJoinOnFields(t *testing.T) {
	var (
		table1 = gtime.X取文本时间戳纳秒() + "_table1"
		table2 = gtime.X取文本时间戳纳秒() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table1).
			X字段保留过滤并带前缀(table1, "*").
			X右连接带比较运算符(table2, "id", "=", "id").
			X条件包含("id", g.Slice别名{1, 2}).
			X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
}

func Test_Model_InnerJoinOnFields(t *testing.T) {
	var (
		table1 = gtime.X取文本时间戳纳秒() + "_table1"
		table2 = gtime.X取文本时间戳纳秒() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table1).
			X字段保留过滤并带前缀(table1, "*").
			X内连接带比较运算符(table2, "id", "=", "id").
			X条件包含("id", g.Slice别名{1, 2}).
			X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[1]["id"], "2")
	})
}

func Test_Model_FieldsPrefix(t *testing.T) {
	var (
		table1 = gtime.X取文本时间戳纳秒() + "_table1"
		table2 = gtime.X取文本时间戳纳秒() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table1).
			X字段保留过滤并带前缀(table1, "id").
			X字段保留过滤并带前缀(table2, "nickname").
			X左连接相同字段(table2, "id").
			X条件包含("id", g.Slice别名{1, 2}).
			X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["id"], "1")
		t.Assert(r[0]["nickname"], "name_1")
	})
}
