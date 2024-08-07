//go:build 屏蔽单元测试

// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package clickhouse_test

import (
	"fmt"
	"strings"
	"testing"

	gdb "github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_New(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		node := gdb.ConfigNode{
			Host:  "127.0.0.1",
			Port:  "9000",
			User:  "default",
			Name:  "default",
			Type:  "clickhouse",
			Debug: false,
		}
		newDb, err := gdb.X创建DB对象(node)
		t.AssertNil(err)
		value, err := newDb.X原生SQL查询字段值(ctx, `select 1`)
		t.AssertNil(err)
		t.Assert(value, `1`)
		t.AssertNil(newDb.X关闭数据库(ctx))
	})
}

func Test_Model_Raw(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X原生SQL(fmt.Sprintf("select id from %s ", table)).X查询行数()
		t.Assert(count, 10)
		t.AssertNil(err)
	})

	gtest.C(t, func(t *gtest.T) {
		model := db.X创建Model对象(table)
		result, err := model.X设置数据(g.Map{
			"id":       gdb.Raw("1+5"),
			"passport": "port_1",
			"password": "pass_1",
			"nickname": "name_1",
		}).X插入()
		t.Assert(strings.Contains(err.Error(), "converting gdb.Raw to UInt64 is unsupported"), true)

		t.AssertNil(result)
	})
}

func Test_Model_Insert(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		user := db.X创建Model对象(table)
		_, err := user.X设置数据(g.Map{
			"id":          uint64(1),
			"uid":         1,
			"passport":    "t1",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "name_1",
			"create_time": gtime.X创建并按当前时间(),
		}).X插入()
		t.AssertNil(err)

		_, err = db.X创建Model对象(table).X设置数据(g.Map{
			"id":          uint64(2),
			"uid":         "2",
			"passport":    "t2",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "name_2",
			"create_time": gtime.X创建并按当前时间(),
		}).X插入()
		t.AssertNil(err)

		type User struct {
			Id         uint64      `gconv:"id"`
			Uid        int         `gconv:"uid"`
			Passport   string      `json:"passport"`
			Password   string      `gconv:"password"`
			Nickname   string      `gconv:"nickname"`
			CreateTime *gtime.Time `json:"create_time"`
		}
		// Model inserting.
		_, err = db.X创建Model对象(table).X设置数据(User{
			Id:         3,
			Uid:        3,
			Passport:   "t3",
			Password:   "25d55ad283aa400af464c76d713c07ad",
			Nickname:   "name_3",
			CreateTime: gtime.X创建并按当前时间(),
		}).X插入()
		t.AssertNil(err)

		value, err := db.X创建Model对象(table).X字段保留过滤("passport").X条件("id=3").X查询一条值() // model value
		t.AssertNil(err)
		t.Assert(value.String(), "t3")

		_, err = db.X创建Model对象(table).X设置数据(&User{
			Id:         4,
			Uid:        4,
			Passport:   "t4",
			Password:   "25d55ad283aa400af464c76d713c07ad",
			Nickname:   "T4",
			CreateTime: gtime.X创建并按当前时间(),
		}).X插入()
		t.AssertNil(err)

		value, err = db.X创建Model对象(table).X字段保留过滤("passport").X条件("id=4").X查询一条值()
		t.AssertNil(err)
		t.Assert(value.String(), "t4")

		_, err = db.X创建Model对象(table).X条件("id>?", 1).X删除() // model delete
		t.AssertNil(err)

	})
}

func Test_Model_One(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         uint64
			Passport   string
			Password   string
			Nickname   string
			CreateTime *gtime.Time
		}
		data := User{
			Id:         1,
			Passport:   "user_1",
			Password:   "pass_1",
			Nickname:   "name_1",
			CreateTime: gtime.X创建并按当前时间(),
		}
		_, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)

		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条() // model one
		t.AssertNil(err)
		t.Assert(one["passport"], data.Passport)
		t.Assert(one["create_time"], data.CreateTime)
		t.Assert(one["nickname"], data.Nickname)
	})
}

func Test_Model_All(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
	})
}

func Test_Model_Delete(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		_, err := db.X创建Model对象(table).X条件("id", "2").X删除()
		t.AssertNil(err)

	})
}

func Test_Model_Update(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		_, err := db.X创建Model对象(table).X条件("passport='user_3'").X更新()
		t.AssertEQ(err.Error(), "updating table with empty data")
	})

		// 更新 + Fields(字符串). md5:df4e16d13da67d5e
	gtest.C(t, func(t *gtest.T) {
		_, err := db.X创建Model对象(table).X字段保留过滤("passport").X设置数据(g.Map{
			"passport": "user_44",
			"none":     "none",
		}).X条件("passport='user_4'").X更新()
		t.AssertNil(err)

		_, err = db.X创建Model对象(table).X条件("passport='user_44'").X查询一条()
		t.AssertNil(err)
	})
}

func Test_Model_Array(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X条件("id", g.Slice别名{1, 2, 3}).X查询()
		t.AssertNil(err)
		t.Assert(all.X取字段切片("id"), g.Slice别名{1, 2, 3})
		t.Assert(all.X取字段切片("nickname"), g.Slice别名{"name_1", "name_2", "name_3"})
	})
	gtest.C(t, func(t *gtest.T) {
		array, err := db.X创建Model对象(table).X字段保留过滤("nickname").X条件("id", g.Slice别名{1, 2, 3}).X查询切片()
		t.AssertNil(err)
		t.Assert(array, g.Slice别名{"name_1", "name_2", "name_3"})
	})
	gtest.C(t, func(t *gtest.T) {
		array, err := db.X创建Model对象(table).X查询切片("nickname", "id", g.Slice别名{1, 2, 3})
		t.AssertNil(err)
		t.Assert(array, g.Slice别名{"name_1", "name_2", "name_3"})
	})
}

func Test_Model_Scan(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	type User struct {
		Id         uint64
		Passport   string
		Password   string
		NickName   string
		CreateTime gtime.Time
	}
	gtest.C(t, func(t *gtest.T) {
		var users []User
		err := db.X创建Model对象(table).X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), TableSize)
	})
}

func Test_Model_Count(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(count, TableSize)
	})
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X字段排除过滤("id").X条件("id>8").X查询行数()
		t.AssertNil(err)
		t.Assert(count, 2)
	})
}

func Test_Model_Where(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	// map + slice parameter
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(g.Map{
			"id":       g.Slice别名{1, 2, 3},
			"passport": g.Slice别名{"user_2", "user_3"},
		}).X条件("id=? and nickname=?", g.Slice别名{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["id"].X取整数(), 3)
	})

		// 结构体，自动映射和过滤。 md5:8edea55227b914af
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id       int
			Nickname string
		}
		result, err := db.X创建Model对象(table).X条件(User{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)

		result, err = db.X创建Model对象(table).X条件(&User{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["id"].X取整数(), 3)
	})
}

func Test_Model_Sav(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		_, err := db.X创建Model对象(table).X设置数据(g.Map{
			"id":          uint64(1),
			"passport":    "t111",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "T111",
			"create_time": gtime.X创建并按当前时间(),
		}).X插入并更新已存在()
		t.AssertNil(err)
	})
}

func Test_Model_Replace(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		_, err := db.X创建Model对象(table).X设置数据(g.Map{
			"id":          uint64(1),
			"passport":    "t11",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "T11",
			"create_time": gtime.X创建并按当前时间(),
		}).X插入并替换已存在()
		t.AssertNil(err)
	})
}
