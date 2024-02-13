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
		table1 = 时间类.X取文本时间戳纳秒() + "_table1"
		table2 = 时间类.X取文本时间戳纳秒() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	单元测试类.C(t, func(t *单元测试类.T) {
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
		table1 = 时间类.X取文本时间戳纳秒() + "_table1"
		table2 = 时间类.X取文本时间戳纳秒() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	单元测试类.C(t, func(t *单元测试类.T) {
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
		table1 = 时间类.X取文本时间戳纳秒() + "_table1"
		table2 = 时间类.X取文本时间戳纳秒() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	单元测试类.C(t, func(t *单元测试类.T) {
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
		table1 = 时间类.X取文本时间戳纳秒() + "_table1"
		table2 = 时间类.X取文本时间戳纳秒() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	单元测试类.C(t, func(t *单元测试类.T) {
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
		table1 = 时间类.X取文本时间戳纳秒() + "_table1"
		table2 = 时间类.X取文本时间戳纳秒() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	单元测试类.C(t, func(t *单元测试类.T) {
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
		table1 = 时间类.X取文本时间戳纳秒() + "_table1"
		table2 = 时间类.X取文本时间戳纳秒() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	单元测试类.C(t, func(t *单元测试类.T) {
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
		table1 = 时间类.X取文本时间戳纳秒() + "_table1"
		table2 = 时间类.X取文本时间戳纳秒() + "_table2"
	)
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	单元测试类.C(t, func(t *单元测试类.T) {
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
