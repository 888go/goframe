//go:build 屏蔽单元测试

// 版权所有 2019 gf 作者（https://github.com/gogf/gf）。保留所有权利。
//
// 此源代码形式受麻省理工学院（MIT）许可证的条款约束。
// 如果未随此文件一起分发MIT许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:47e609239e0cb2bc

package mssql_test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	gconv "github.com/888go/goframe/util/gconv"

	garray "github.com/888go/goframe/container/garray"
	gmap "github.com/888go/goframe/container/gmap"
	gdb "github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
	gutil "github.com/888go/goframe/util/gutil"
)

func Test_Page(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	// db.SetDebug(true)
	result, err := db.X创建Model对象(table).X设置分页(1, 2).X排序("id").X查询()
	gtest.Assert(err, nil)
	fmt.Println("page:1--------", result)
	gtest.Assert(len(result), 2)
	gtest.Assert(result[0]["ID"], 1)
	gtest.Assert(result[1]["ID"], 2)

	result, err = db.X创建Model对象(table).X设置分页(2, 2).X排序("id").X查询()
	gtest.Assert(err, nil)
	fmt.Println("page: 2--------", result)
	gtest.Assert(len(result), 2)
	gtest.Assert(result[0]["ID"], 3)
	gtest.Assert(result[1]["ID"], 4)

	result, err = db.X创建Model对象(table).X设置分页(3, 2).X排序("id").X查询()
	gtest.Assert(err, nil)
	fmt.Println("page:3 --------", result)
	gtest.Assert(len(result), 2)
	gtest.Assert(result[0]["ID"], 5)

	result, err = db.X创建Model对象(table).X设置分页(2, 3).X查询()
	gtest.Assert(err, nil)
	gtest.Assert(len(result), 3)
}

func Test_Model_Insert(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		user := db.X创建Model对象(table)
		result, err := user.X设置数据(g.Map{
			"id":          1,
			"uid":         1,
			"passport":    "t1",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "name_1",
			"create_time": gtime.X创建并按当前时间().String(),
		}).X插入()
		t.AssertNil(err)

		result, err = db.X创建Model对象(table).X设置数据(g.Map{
			"id":          "2",
			"uid":         "2",
			"passport":    "t2",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "name_2",
			"create_time": gtime.X创建并按当前时间().String(),
		}).X插入()
		t.AssertNil(err)

		type User struct {
			Id         int         `gconv:"id"`
			Uid        int         `gconv:"uid"`
			Passport   string      `json:"passport"`
			Password   string      `gconv:"password"`
			Nickname   string      `gconv:"nickname"`
			CreateTime *gtime.Time `json:"create_time"`
		}
		// Model inserting.
		result, err = db.X创建Model对象(table).X设置数据(User{
			Id:       3,
			Uid:      3,
			Passport: "t3",
			Password: "25d55ad283aa400af464c76d713c07ad",
			Nickname: "name_3",
		}).X插入()
		t.AssertNil(err)

		value, err := db.X创建Model对象(table).X字段保留过滤("passport").X条件("id=3").X查询一条值()
		t.AssertNil(err)
		t.Assert(value.String(), "t3")

		result, err = db.X创建Model对象(table).X设置数据(&User{
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

		result, err = db.X创建Model对象(table).X条件("id>?", 1).X删除()
		t.AssertNil(err)
		_, _ = result.RowsAffected()

	})
}

func Test_Model_Insert_KeyFieldNameMapping(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			Nickname   string
			CreateTime string
		}
		data := User{
			Id:         1,
			Passport:   "user_1",
			Password:   "pass_1",
			Nickname:   "name_1",
			CreateTime: "2020-10-10 12:00:01",
		}
		_, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)

		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["PASSPORT"], data.Passport)
		t.Assert(one["CREATE_TIME"], data.CreateTime)
		t.Assert(one["NICKNAME"], data.Nickname)
	})
}

func Test_Model_Update_KeyFieldNameMapping(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			Nickname   string
			CreateTime string
		}
		data := User{
			Id:         1,
			Passport:   "user_10",
			Password:   "pass_10",
			Nickname:   "name_10",
			CreateTime: "2020-10-10 12:00:01",
		}
		_, err := db.X创建Model对象(table).X设置数据(data).X条件并识别主键(1).X更新()
		t.AssertNil(err)

		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["PASSPORT"], data.Passport)
		t.Assert(one["CREATE_TIME"], data.CreateTime)
		t.Assert(one["NICKNAME"], data.Nickname)
	})
}

func Test_Model_Insert_Time(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		data := g.Map{
			"id":          1,
			"passport":    "t1",
			"password":    "p1",
			"nickname":    "n1",
			"create_time": "2020-10-10 20:09:18.334",
		}
		_, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)

		one, err := db.X创建Model对象(table).X查询一条("id", 1)
		t.AssertNil(err)
		t.Assert(one["PASSPORT"].String(), data["passport"])
		t.Assert(one["CREATE_TIME"].String(), "2020-10-10 20:09:18")
		t.Assert(one["NICKNAME"].String(), data["nickname"])
	})
}

func Test_Model_BatchInsertWithArrayStruct(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		user := db.X创建Model对象(table)
		array := garray.X创建()
		for i := 1; i <= TableSize; i++ {
			array.Append别名(g.Map{
				"id":          i,
				"uid":         i,
				"passport":    fmt.Sprintf("t%d", i),
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    fmt.Sprintf("name_%d", i),
				"create_time": gtime.X创建并按当前时间().String(),
			})
		}

		_, err := user.X设置数据(array).X插入()
		t.AssertNil(err)

	})
}

func Test_Model_Batch(t *testing.T) {
	// batch insert
	gtest.C(t, func(t *gtest.T) {
		table := createTable()
		defer dropTable(table)
		_, err := db.X创建Model对象(table).X设置数据(g.Map切片{
			{
				"id":          2,
				"uid":         2,
				"passport":    "t2",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "name_2",
				"create_time": gtime.X创建并按当前时间().String(),
			},
			{
				"id":          3,
				"uid":         3,
				"passport":    "t3",
				"password":    "25d55ad283aa400af464c76d713c07ad",
				"nickname":    "name_3",
				"create_time": gtime.X创建并按当前时间().String(),
			},
		}).X设置批量操作行数(1).X插入()
		if err != nil {
			gtest.Error(err)
		}
	})

}

func Test_Model_Update(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X设置数据("passport", "user_22").X条件("passport=?", "user_2").X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X设置数据("passport", "user_2").X条件("passport='user_22'").X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})

	// Update + Data(string)
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X设置数据("passport='user_33'").X条件("passport='user_3'").X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
		// 更新 + Fields(字符串). md5:df4e16d13da67d5e
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X字段保留过滤("passport").X设置数据(g.Map{
			"passport": "user_44",
			"none":     "none",
		}).X条件("passport='user_4'").X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
}

func Test_Model_Clone(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		md := db.X创建Model对象(table).X链式安全(true).X条件("id IN(?)", g.Slice别名{1, 3})
		count, err := md.X查询行数()
		t.AssertNil(err)

		record, err := md.X链式安全(true).X排序("id DESC").X查询一条()
		t.AssertNil(err)

		result, err := md.X链式安全(true).X排序("id ASC").X查询()
		t.AssertNil(err)

		t.Assert(count, int64(2))
		t.Assert(record["ID"].X取整数(), 3)
		t.Assert(len(result), 2)
		t.Assert(result[0]["ID"].X取整数(), 1)
		t.Assert(result[1]["ID"].X取整数(), 3)
	})
}

func Test_Model_Safe(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		md := db.X创建Model对象(table).X链式安全(false).X条件("id IN(?)", g.Slice别名{1, 3})
		count, err := md.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))

		md.X条件("id = ?", 1)
		count, err = md.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(1))
	})
	gtest.C(t, func(t *gtest.T) {
		md := db.X创建Model对象(table).X链式安全(true).X条件("id IN(?)", g.Slice别名{1, 3})
		count, err := md.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))

		md.X条件("id = ?", 1)
		count, err = md.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))
	})

	gtest.C(t, func(t *gtest.T) {
		md := db.X创建Model对象(table).X链式安全().X条件("id IN(?)", g.Slice别名{1, 3})
		count, err := md.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))

		md.X条件("id = ?", 1)
		count, err = md.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))
	})
	gtest.C(t, func(t *gtest.T) {
		md1 := db.X创建Model对象(table).X链式安全()
		md2 := md1.X条件("id in (?)", g.Slice别名{1, 3})
		count, err := md2.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))

		all, err := md2.X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)

		all, err = md2.X设置分页(1, 10).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
	})

	gtest.C(t, func(t *gtest.T) {
		table := createInitTable()
		defer dropTable(table)

		md1 := db.X创建Model对象(table).X条件("id>", 0).X链式安全()
		md2 := md1.X条件("id in (?)", g.Slice别名{1, 3})
		md3 := md1.X条件("id in (?)", g.Slice别名{4, 5, 6})

		// 1,3
		count, err := md2.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))

		all, err := md2.X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
		t.Assert(all[0]["ID"].X取整数(), 1)
		t.Assert(all[1]["ID"].X取整数(), 3)

		all, err = md2.X设置分页(1, 10).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)

		// 4,5,6
		count, err = md3.X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(3))

		all, err = md3.X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(all), 3)
		t.Assert(all[0]["ID"].X取整数(), 4)
		t.Assert(all[1]["ID"].X取整数(), 5)
		t.Assert(all[2]["ID"].X取整数(), 6)

		all, err = md3.X设置分页(1, 10).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 3)
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
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id<0").X查询()
		t.Assert(result, nil)
		t.AssertNil(err)
	})
}

func Test_Model_One(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		record, err := db.X创建Model对象(table).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(record["NICKNAME"].String(), "name_1")
	})

	gtest.C(t, func(t *gtest.T) {
		record, err := db.X创建Model对象(table).X条件("id", 0).X查询一条()
		t.AssertNil(err)
		t.Assert(record, nil)
	})
}

func Test_Model_Value(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		value, err := db.X创建Model对象(table).X字段保留过滤("nickname").X条件("id", 1).X查询一条值()
		t.AssertNil(err)
		t.Assert(value.String(), "name_1")
	})

	gtest.C(t, func(t *gtest.T) {
		value, err := db.X创建Model对象(table).X字段保留过滤("nickname").X条件("id", 0).X查询一条值()
		t.AssertNil(err)
		t.Assert(value, nil)
	})
}

func Test_Model_Array(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X条件("id", g.Slice别名{1, 2, 3}).X查询()
		t.AssertNil(err)
		t.Assert(gconv.X取整数切片(all.X取字段切片("ID")), g.Slice别名{1, 2, 3})
		t.Assert(all.X取字段切片("NICKNAME"), g.Slice别名{"name_1", "name_2", "name_3"})
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

func Test_Model_Count(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
		// 使用缓存计数，检查内部上下文数据特性。 md5:fa8263fd899afcec
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 10; i++ {
			count, err := db.X创建Model对象(table).X缓存(gdb.CacheOption{
				X时长: time.Second * 10,
				X名称:     guid.X生成(),
				X强制缓存:    false,
			}).X查询行数()
			t.AssertNil(err)
			t.Assert(count, int64(TableSize))
		}
	})
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X字段排除过滤("id").X条件("id>8").X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))
	})
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X字段保留过滤("distinct id").X条件("id>8").X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))
	})
	// COUNT...LIMIT...
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X设置分页(1, 2).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})

}

func Test_Model_Select(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	type User struct {
		Id         int
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

func Test_Model_Struct(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		user := new(User)
		err := db.X创建Model对象(table).X条件("id=1").X查询到结构体指针(user)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_1")
	})
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		user := new(User)
		err := db.X创建Model对象(table).X条件("id=1").X查询到结构体指针(user)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_1")
	})
		// 自动创建结构体对象。 md5:4b196dfc1321dc30
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		user := (*User)(nil)
		err := db.X创建Model对象(table).X条件("id=1").X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_1")
	})
	// Just using Scan.
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		user := (*User)(nil)
		err := db.X创建Model对象(table).X条件("id=1").X查询到结构体指针(&user)
		if err != nil {
			gtest.Error(err)
		}
		t.Assert(user.NickName, "name_1")
	})
	// sql.ErrNoRows
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		user := new(User)
		err := db.X创建Model对象(table).X条件("id=-1").X查询到结构体指针(user)
		t.Assert(err, sql.ErrNoRows)
	})
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var user *User
		err := db.X创建Model对象(table).X条件("id=-1").X查询到结构体指针(&user)
		t.AssertNil(err)
	})
}

func Test_Model_Struct_CustomType(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	type MyInt int

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         MyInt
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		user := new(User)
		err := db.X创建Model对象(table).X条件("id=1").X查询到结构体指针(user)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_1")
	})
}

func Test_Model_Structs(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		var users []User
		err := db.X创建Model对象(table).X排序("id asc").X查询到结构体指针(&users)
		if err != nil {
			gtest.Error(err)
		}
		t.Assert(len(users), TableSize)
		t.Assert(users[0].Id, 1)
		t.Assert(users[1].Id, 2)
		t.Assert(users[2].Id, 3)
		t.Assert(users[0].NickName, "name_1")
		t.Assert(users[1].NickName, "name_2")
		t.Assert(users[2].NickName, "name_3")
	})
		// 自动创建结构体切片。 md5:78598f0d7f20b815
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var users []*User
		err := db.X创建Model对象(table).X排序("id asc").X查询到结构体指针(&users)
		if err != nil {
			gtest.Error(err)
		}
		t.Assert(len(users), TableSize)
		t.Assert(users[0].Id, 1)
		t.Assert(users[1].Id, 2)
		t.Assert(users[2].Id, 3)
		t.Assert(users[0].NickName, "name_1")
		t.Assert(users[1].NickName, "name_2")
		t.Assert(users[2].NickName, "name_3")
	})
	// Just using Scan.
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var users []*User
		err := db.X创建Model对象(table).X排序("id asc").X查询到结构体指针(&users)
		if err != nil {
			gtest.Error(err)
		}
		t.Assert(len(users), TableSize)
		t.Assert(users[0].Id, 1)
		t.Assert(users[1].Id, 2)
		t.Assert(users[2].Id, 3)
		t.Assert(users[0].NickName, "name_1")
		t.Assert(users[1].NickName, "name_2")
		t.Assert(users[2].NickName, "name_3")
	})
	// sql.ErrNoRows
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var users []*User
		err := db.X创建Model对象(table).X条件("id<0").X查询到结构体指针(&users)
		t.AssertNil(err)
	})
}

func Test_Model_Scan(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		user := new(User)
		err := db.X创建Model对象(table).X条件("id=1").X查询到结构体指针(user)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_1")
	})
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		user := new(User)
		err := db.X创建Model对象(table).X条件("id=1").X查询到结构体指针(user)
		t.AssertNil(err)
		t.Assert(user.NickName, "name_1")
	})
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		var users []User
		err := db.X创建Model对象(table).X排序("id asc").X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), TableSize)
		t.Assert(users[0].Id, 1)
		t.Assert(users[1].Id, 2)
		t.Assert(users[2].Id, 3)
		t.Assert(users[0].NickName, "name_1")
		t.Assert(users[1].NickName, "name_2")
		t.Assert(users[2].NickName, "name_3")
	})
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var users []*User
		err := db.X创建Model对象(table).X排序("id asc").X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), TableSize)
		t.Assert(users[0].Id, 1)
		t.Assert(users[1].Id, 2)
		t.Assert(users[2].Id, 3)
		t.Assert(users[0].NickName, "name_1")
		t.Assert(users[1].NickName, "name_2")
		t.Assert(users[2].NickName, "name_3")
	})
	// sql.ErrNoRows
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var (
			user  = new(User)
			users = new([]*User)
		)
		err1 := db.X创建Model对象(table).X条件("id < 0").X查询到结构体指针(user)
		err2 := db.X创建Model对象(table).X条件("id < 0").X查询到结构体指针(users)
		t.Assert(err1, sql.ErrNoRows)
		t.Assert(err2, nil)
	})
}

func Test_Model_Scan_NilSliceAttrWhenNoRecordsFound(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		type Response struct {
			Users []User `json:"users"`
		}
		var res Response
		err := db.X创建Model对象(table).X查询到结构体指针(&res.Users)
		t.AssertNil(err)
		t.Assert(res.Users, nil)
	})
}

func Test_Model_OrderBy(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X排序("id DESC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(result[0]["NICKNAME"].String(), fmt.Sprintf("name_%d", TableSize))
	})
}

func Test_Model_Data(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		table := createInitTable()
		defer dropTable(table)
		result, err := db.X创建Model对象(table).X设置数据("nickname=?", "test").X条件("id=?", 3).X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
	gtest.C(t, func(t *gtest.T) {
		table := createTable()
		defer dropTable(table)
		users := make([]g.MapStrAny, 0)
		for i := 1; i <= 10; i++ {
			users = append(users, g.MapStrAny{
				"id":       i,
				"passport": fmt.Sprintf(`passport_%d`, i),
				"password": fmt.Sprintf(`password_%d`, i),
				"nickname": fmt.Sprintf(`nickname_%d`, i),
			})
		}
		result, err := db.X创建Model对象(table).X设置数据(users).X设置批量操作行数(2).X插入()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 10)
	})
	gtest.C(t, func(t *gtest.T) {
		table := createTable()
		defer dropTable(table)
		users := garray.X创建()
		for i := 1; i <= 10; i++ {
			users.Append别名(g.MapStrAny{
				"id":       i,
				"passport": fmt.Sprintf(`passport_%d`, i),
				"password": fmt.Sprintf(`password_%d`, i),
				"nickname": fmt.Sprintf(`nickname_%d`, i),
			})
		}
		result, err := db.X创建Model对象(table).X设置数据(users).X设置批量操作行数(2).X插入()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 10)
	})
}

func Test_Model_Where(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	// string
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id=? and nickname=?", 3, "name_3").X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["ID"].X取整数(), 3)
	})

	// slice
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(g.Slice别名{"id", 3}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(g.Slice别名{"id", 3, "nickname", "name_3"}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["ID"].X取整数(), 3)
	})

	// slice parameter
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id=? and nickname=?", g.Slice别名{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["ID"].X取整数(), 3)
	})
	// map like
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(g.Map{
			"passport like": "user_1%",
		}).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0].X取Map类().X取值("ID"), 1)
		t.Assert(result[1].X取Map类().X取值("ID"), 10)
	})
	// map + slice parameter
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(g.Map{
			"id":       g.Slice别名{1, 2, 3},
			"passport": g.Slice别名{"user_2", "user_3"},
		}).X条件("id=? and nickname=?", g.Slice别名{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id=3", g.Slice别名{}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id=?", g.Slice别名{3}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id", 3).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id", 3).X条件("nickname", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id", 3).X条件("nickname", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id", 30).X条件或("nickname", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id", 30).X条件或("nickname", "name_3").X条件("id>?", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id", 30).X条件或("nickname", "name_3").X条件("id>", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	// slice
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id=? AND nickname=?", g.Slice别名{3, "name_3"}...).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id=? AND nickname=?", g.Slice别名{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("passport like ? and nickname like ?", g.Slice别名{"user_3", "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	// map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(g.Map{"id": 3, "nickname": "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	// map key operator
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(g.Map{"id>": 1, "id<": 3}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 2)
	})

	// gmap.Map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(gmap.X创建并从Map(g.MapAnyAny{"id": 3, "nickname": "name_3"})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	// gmap.Map key operator
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(gmap.X创建并从Map(g.MapAnyAny{"id>": 1, "id<": 3})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 2)
	})

	// list map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(gmap.X创建链表Map并从Map(g.MapAnyAny{"id": 3, "nickname": "name_3"})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	// list map key operator
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(gmap.X创建链表Map并从Map(g.MapAnyAny{"id>": 1, "id<": 3})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 2)
	})

	// tree map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(gmap.X创建红黑树Map并从Map(gutil.X比较文本, g.MapAnyAny{"id": 3, "nickname": "name_3"})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	// tree map key operator
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(gmap.X创建红黑树Map并从Map(gutil.X比较文本, g.MapAnyAny{"id>": 1, "id<": 3})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 2)
	})

	// complicated where 1
	gtest.C(t, func(t *gtest.T) {
		// db.SetDebug(true)
		conditions := g.Map{
			"nickname like ?":    "%name%",
			"id between ? and ?": g.Slice别名{1, 3},
			"id > 0":             nil,
			"create_time > 0":    nil,
			"id":                 g.Slice别名{1, 2, 3},
		}
		result, err := db.X创建Model对象(table).X条件(conditions).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["ID"].X取整数(), 1)
	})
	// complicated where 2
	gtest.C(t, func(t *gtest.T) {
		// db.SetDebug(true)
		conditions := g.Map{
			"nickname like ?":    "%name%",
			"id between ? and ?": g.Slice别名{1, 3},
			"id >= ?":            1,
			"create_time > ?":    0,
			"id in(?)":           g.Slice别名{1, 2, 3},
		}
		result, err := db.X创建Model对象(table).X条件(conditions).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["ID"].X取整数(), 1)
	})
		// 结构体，自动映射和过滤。 md5:8edea55227b914af
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id       int
			Nickname string
		}
		result, err := db.X创建Model对象(table).X条件(User{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)

		result, err = db.X创建Model对象(table).X条件(&User{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	// slice single
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("id IN(?)", g.Slice别名{1, 3}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0]["ID"].X取整数(), 1)
		t.Assert(result[1]["ID"].X取整数(), 3)
	})
	// slice + string
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("nickname=? AND id IN(?)", "name_3", g.Slice别名{1, 3}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["ID"].X取整数(), 3)
	})
	// slice + map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(g.Map{
			"id":       g.Slice别名{1, 3},
			"nickname": "name_3",
		}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["ID"].X取整数(), 3)
	})
	// slice + struct
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Ids      []int  `json:"id"`
			Nickname string `gconv:"nickname"`
		}
		result, err := db.X创建Model对象(table).X条件(User{
			Ids:      []int{1, 3},
			Nickname: "name_3",
		}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["ID"].X取整数(), 3)
	})
}

func Test_Model_Where_ISNULL_1(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		// db.SetDebug(true)
		result, err := db.X创建Model对象(table).X设置数据("nickname", nil).X条件("id", 2).X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).X条件("nickname", nil).X查询一条()
		t.AssertNil(err)
		t.Assert(one.X是否为空(), false)
		t.Assert(one["ID"], 2)
	})
}

func Test_Model_Where_ISNULL_2(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	// complicated one.
	gtest.C(t, func(t *gtest.T) {
		// db.SetDebug(true)
		conditions := g.Map{
			"nickname like ?":    "%name%",
			"id between ? and ?": g.Slice别名{1, 3},
			"id > 0":             nil,
			"create_time > 0":    nil,
			"id":                 g.Slice别名{1, 2, 3},
		}
		result, err := db.X创建Model对象(table).X条件(conditions).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["ID"].X取整数(), 1)
	})
}

func Test_Model_Where_OmitEmpty(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		conditions := g.Map{
			"id < 4": "",
		}
		result, err := db.X创建Model对象(table).X条件(conditions).X排序("id desc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		conditions := g.Map{
			"id < 4": "",
		}
		result, err := db.X创建Model对象(table).X条件(conditions).X过滤空值().X排序("id desc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 10)
		t.Assert(result[0]["ID"].X取整数(), 10)
	})
}

func Test_Model_Where_GTime(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("create_time>?", gtime.X创建并从文本("2010-09-01")).X查询()
		t.AssertNil(err)
		t.Assert(len(result), 10)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("create_time>?", *gtime.X创建并从文本("2010-09-01")).X查询()
		t.AssertNil(err)
		t.Assert(len(result), 10)
	})
}

func Test_Model_WherePri(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	// primary key
	gtest.C(t, func(t *gtest.T) {
		one, err := db.X创建Model对象(table).X条件并识别主键(3).X查询一条()
		t.AssertNil(err)
		t.AssertNE(one, nil)
		t.Assert(one["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X条件并识别主键(g.Slice别名{3, 9}).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
		t.Assert(all[0]["ID"].X取整数(), 3)
		t.Assert(all[1]["ID"].X取整数(), 9)
	})

	// string
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id=? and nickname=?", 3, "name_3").X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["ID"].X取整数(), 3)
	})
	// slice parameter
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id=? and nickname=?", g.Slice别名{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["ID"].X取整数(), 3)
	})
	// map like
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(g.Map{
			"passport like": "user_1%",
		}).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0].X取Map类().X取值("ID"), 1)
		t.Assert(result[1].X取Map类().X取值("ID"), 10)
	})
	// map + slice parameter
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(g.Map{
			"id":       g.Slice别名{1, 2, 3},
			"passport": g.Slice别名{"user_2", "user_3"},
		}).X条件("id=? and nickname=?", g.Slice别名{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(g.Map{
			"id":       g.Slice别名{1, 2, 3},
			"passport": g.Slice别名{"user_2", "user_3"},
		}).X条件或("nickname=?", g.Slice别名{"name_4"}).X条件("id", 3).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["ID"].X取整数(), 2)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id=3", g.Slice别名{}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id=?", g.Slice别名{3}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id", 3).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id", 3).X条件并识别主键("nickname", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id", 3).X条件("nickname", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id", 30).X条件或("nickname", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id", 30).X条件或("nickname", "name_3").X条件("id>?", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id", 30).X条件或("nickname", "name_3").X条件("id>", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	// slice
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id=? AND nickname=?", g.Slice别名{3, "name_3"}...).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id=? AND nickname=?", g.Slice别名{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("passport like ? and nickname like ?", g.Slice别名{"user_3", "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	// map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(g.Map{"id": 3, "nickname": "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	// map key operator
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(g.Map{"id>": 1, "id<": 3}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 2)
	})

	// gmap.Map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(gmap.X创建并从Map(g.MapAnyAny{"id": 3, "nickname": "name_3"})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	// gmap.Map key operator
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(gmap.X创建并从Map(g.MapAnyAny{"id>": 1, "id<": 3})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 2)
	})

	// list map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(gmap.X创建链表Map并从Map(g.MapAnyAny{"id": 3, "nickname": "name_3"})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	// list map key operator
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(gmap.X创建链表Map并从Map(g.MapAnyAny{"id>": 1, "id<": 3})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 2)
	})

	// tree map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(gmap.X创建红黑树Map并从Map(gutil.X比较文本, g.MapAnyAny{"id": 3, "nickname": "name_3"})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	// tree map key operator
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(gmap.X创建红黑树Map并从Map(gutil.X比较文本, g.MapAnyAny{"id>": 1, "id<": 3})).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 2)
	})

	// complicated where 1
	gtest.C(t, func(t *gtest.T) {
		// db.SetDebug(true)
		conditions := g.Map{
			"nickname like ?":    "%name%",
			"id between ? and ?": g.Slice别名{1, 3},
			"id > 0":             nil,
			"create_time > 0":    nil,
			"id":                 g.Slice别名{1, 2, 3},
		}
		result, err := db.X创建Model对象(table).X条件并识别主键(conditions).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["ID"].X取整数(), 1)
	})
	// complicated where 2
	gtest.C(t, func(t *gtest.T) {
		// db.SetDebug(true)
		conditions := g.Map{
			"nickname like ?":    "%name%",
			"id between ? and ?": g.Slice别名{1, 3},
			"id >= ?":            1,
			"create_time > ?":    0,
			"id in(?)":           g.Slice别名{1, 2, 3},
		}
		result, err := db.X创建Model对象(table).X条件并识别主键(conditions).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["ID"].X取整数(), 1)
	})
	// struct
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id       int    `json:"id"`
			Nickname string `gconv:"nickname"`
		}
		result, err := db.X创建Model对象(table).X条件并识别主键(User{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)

		result, err = db.X创建Model对象(table).X条件并识别主键(&User{3, "name_3"}).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	// slice single
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("id IN(?)", g.Slice别名{1, 3}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0]["ID"].X取整数(), 1)
		t.Assert(result[1]["ID"].X取整数(), 3)
	})
	// slice + string
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键("nickname=? AND id IN(?)", "name_3", g.Slice别名{1, 3}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["ID"].X取整数(), 3)
	})
	// slice + map
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件并识别主键(g.Map{
			"id":       g.Slice别名{1, 3},
			"nickname": "name_3",
		}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["ID"].X取整数(), 3)
	})
	// slice + struct
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Ids      []int  `json:"id"`
			Nickname string `gconv:"nickname"`
		}
		result, err := db.X创建Model对象(table).X条件并识别主键(User{
			Ids:      []int{1, 3},
			Nickname: "name_3",
		}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["ID"].X取整数(), 3)
	})
}

func Test_Model_Delete(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("1=1").X删除()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, TableSize)
	})
}

func Test_Model_Offset(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X设置条数(5, 2).X排序("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0]["ID"], 6)
		t.Assert(result[1]["ID"], 7)
	})
}

func Test_Model_Option_Map(t *testing.T) {
	// Insert
	gtest.C(t, func(t *gtest.T) {
		table := createTable()
		defer dropTable(table)
		r, err := db.X创建Model对象(table).X字段保留过滤("id, passport").X设置数据(g.Map{
			"id":       1,
			"passport": "1",
			"password": "1",
			"nickname": "1",
		}).X插入()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)
		one, err := db.X创建Model对象(table).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.AssertNE(one["PASSWORD"].String(), "1")
		t.AssertNE(one["NICKNAME"].String(), "1")
		t.Assert(one["PASSPORT"].String(), "1")
	})
	gtest.C(t, func(t *gtest.T) {
		table := createTable()
		defer dropTable(table)
		r, err := db.X创建Model对象(table).X过滤空值数据().X设置数据(g.Map{
			"id":       1,
			"passport": 0,
			"password": 0,
			"nickname": "1",
		}).X插入()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)
		one, err := db.X创建Model对象(table).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.AssertNE(one["PASSPORT"].String(), "0")
		t.AssertNE(one["PASSWORD"].String(), "0")
		t.Assert(one["NICKNAME"].String(), "1")
	})

	// Update
	gtest.C(t, func(t *gtest.T) {
		table := createInitTable()
		defer dropTable(table)

		r, err := db.X创建Model对象(table).X设置数据(g.Map{"nickname": ""}).X条件("id", 1).X更新()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		_, err = db.X创建Model对象(table).X过滤空值数据().X设置数据(g.Map{"nickname": ""}).X条件("id", 2).X更新()
		t.AssertNil(err)

		r, err = db.X创建Model对象(table).X过滤空值().X设置数据(g.Map{"nickname": "", "password": "123"}).X条件("id", 3).X更新()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)

		_, err = db.X创建Model对象(table).X过滤空值().X字段保留过滤("nickname", "password").X设置数据(g.Map{"nickname": "", "password": "123", "passport": "123"}).X条件("id", 4).X更新()
		t.AssertNil(err)

		r, err = db.X创建Model对象(table).X过滤空值().
			X字段保留过滤("password").X设置数据(g.Map{
			"nickname": "",
			"passport": "123",
			"password": "456",
		}).X条件("id", 5).X更新()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).X条件("id", 5).X查询一条()
		t.AssertNil(err)
		t.Assert(one["PASSWORD"], "456")
		t.AssertNE(one["PASSPORT"].String(), "")
		t.AssertNE(one["PASSPORT"].String(), "123")
	})
}

func Test_Model_Option_Where(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		table := createInitTable()
		defer dropTable(table)
		r, err := db.X创建Model对象(table).X过滤空值().X设置数据("nickname", 1).X条件(g.Map{"id": 0, "passport": ""}).X条件("1=1").X更新()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, TableSize)
	})
	return
	gtest.C(t, func(t *gtest.T) {
		table := createInitTable()
		defer dropTable(table)
		r, err := db.X创建Model对象(table).X过滤空值().X设置数据("nickname", 1).X条件(g.Map{"id": 1, "passport": ""}).X更新()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		v, err := db.X创建Model对象(table).X条件("id", 1).X字段保留过滤("nickname").X查询一条值()
		t.AssertNil(err)
		t.Assert(v.String(), "1")
	})
}

func Test_Model_Where_MultiSliceArguments(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table).X条件(g.Map{
			"id":       g.Slice别名{1, 2, 3, 4},
			"passport": g.Slice别名{"user_2", "user_3", "user_4"},
			"nickname": g.Slice别名{"name_2", "name_4"},
			"id >= 4":  nil,
		}).X查询()
		t.AssertNil(err)
		t.Assert(len(r), 1)
		t.Assert(r[0]["ID"], 4)
	})

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(g.Map{
			"id":       g.Slice别名{1, 2, 3},
			"passport": g.Slice别名{"user_2", "user_3"},
		}).X条件或("nickname=?", g.Slice别名{"name_4"}).X条件("id", 3).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["ID"].X取整数(), 2)
	})
}

func Test_Model_FieldsEx(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	// Select.
	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table).X字段排除过滤("create_time, created_at, updated_at, id").X条件("id in (?)", g.Slice别名{1, 2}).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(len(r[0]), 3)
		t.Assert(r[0]["ID"], "")
		t.Assert(r[0]["PASSPORT"], "user_1")
		t.Assert(r[0]["PASSWORD"], "pass_1")
		t.Assert(r[0]["NICKNAME"], "name_1")
		t.Assert(r[0]["CREATE_TIME"], "")
		t.Assert(r[1]["ID"], "")
		t.Assert(r[1]["PASSPORT"], "user_2")
		t.Assert(r[1]["PASSWORD"], "pass_2")
		t.Assert(r[1]["NICKNAME"], "name_2")
		t.Assert(r[1]["CREATE_TIME"], "")
	})
	// Update.
	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table).X字段排除过滤("password").X设置数据(g.Map{"nickname": "123", "password": "456"}).X条件("id", 3).X更新()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).X条件("id", 3).X查询一条()
		t.AssertNil(err)
		t.Assert(one["NICKNAME"], "123")
		t.AssertNE(one["PASSWORD"], "456")
	})
}

func Test_Model_FieldsExStruct(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id       int    `orm:"id"       json:"id"`
			Passport string `orm:"password" json:"pass_port"`
			Password string `orm:"password" json:"password"`
			NickName string `orm:"nickname" json:"nick__name"`
		}
		user := &User{
			Id:       1,
			Passport: "111",
			Password: "222",
			NickName: "333",
		}
		r, err := db.X创建Model对象(table).X字段排除过滤("create_time, password").X过滤空值().X设置数据(user).X插入()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)
	})
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id       int    `orm:"id"       json:"id"`
			Passport string `orm:"password" json:"pass_port"`
			Password string `orm:"password" json:"password"`
			NickName string `orm:"nickname" json:"nick__name"`
		}
		users := make([]*User, 0)
		for i := 100; i < 110; i++ {
			users = append(users, &User{
				Id:       i,
				Passport: fmt.Sprintf(`passport_%d`, i),
				Password: fmt.Sprintf(`password_%d`, i),
				NickName: fmt.Sprintf(`nickname_%d`, i),
			})
		}
		r, err := db.X创建Model对象(table).X字段排除过滤("create_time, password").
			X过滤空值().
			X设置批量操作行数(2).
			X设置数据(users).
			X插入()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 10)
	})
}

func Test_Model_OmitEmpty_Time(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id       int       `orm:"id"       json:"id"`
			Passport string    `orm:"password" json:"pass_port"`
			Password string    `orm:"password" json:"password"`
			Time     time.Time `orm:"create_time" `
		}
		user := &User{
			Id:       1,
			Passport: "111",
			Password: "222",
			Time:     time.Time{},
		}
		r, err := db.X创建Model对象(table).X过滤空值().X设置数据(user).X条件并识别主键(1).X更新()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 1)
	})
}

func Test_Result_Chunk(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table).X排序("id asc").X查询()
		t.AssertNil(err)
		chunks := r.X分割(3)
		t.Assert(len(chunks), 4)
		t.Assert(chunks[0][0]["ID"].X取整数(), 1)
		t.Assert(chunks[1][0]["ID"].X取整数(), 4)
		t.Assert(chunks[2][0]["ID"].X取整数(), 7)
		t.Assert(chunks[3][0]["ID"].X取整数(), 10)
	})
}

func Test_Model_DryRun(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	db.X设置空跑特性(true)
	defer db.X设置空跑特性(false)

	gtest.C(t, func(t *gtest.T) {
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["ID"], 1)
	})
	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table).X设置数据("passport", "port_1").X条件并识别主键(1).X更新()
		t.AssertNil(err)
		n, err := r.RowsAffected()
		t.AssertNil(err)
		t.Assert(n, 0)
	})
}

func Test_Model_Join_SubQuery(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		subQuery := fmt.Sprintf("select * from %s", table)
		r, err := db.X创建Model对象(table, "t1").X字段保留过滤("t2.id").X左连接(subQuery, "t2", "t2.id=t1.id").X查询切片()
		t.AssertNil(err)
		t.Assert(len(r), TableSize)
		t.Assert(r[0], "1")
		t.Assert(r[TableSize-1], TableSize)
	})
}

func Test_Model_Having(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X字段保留过滤("id, count(*)").X条件("id > 1").X排序分组("id").X设置分组条件("id > 8").X查询()
		t.AssertNil(err)
		t.Assert(len(all), 2)
	})

}

func Test_Model_Distinct(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X条件("id > 1").X设置去重().X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(9))
	})
}

func Test_Model_Min_Max(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		value, err := db.X创建Model对象(table, "t").X字段保留过滤("min(t.id)").X条件("id > 1").X查询一条值()
		t.AssertNil(err)
		t.Assert(value.X取整数(), 2)
	})
	gtest.C(t, func(t *gtest.T) {
		value, err := db.X创建Model对象(table, "t").X字段保留过滤("max(t.id)").X条件("id > 1").X查询一条值()
		t.AssertNil(err)
		t.Assert(value.X取整数(), 10)
	})
}

func Test_Model_Fields_AutoMapping(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		value, err := db.X创建Model对象(table).X字段保留过滤("ID").X条件("id", 2).X查询一条值()
		t.AssertNil(err)
		t.Assert(value.X取整数(), 2)
	})

	gtest.C(t, func(t *gtest.T) {
		value, err := db.X创建Model对象(table).X字段保留过滤("NICK_NAME").X条件("id", 2).X查询一条值()
		t.AssertNil(err)
		t.Assert(value.String(), "name_2")
	})
	// Map
	gtest.C(t, func(t *gtest.T) {
		one, err := db.X创建Model对象(table).X字段保留过滤(g.Map{
			"ID":        1,
			"NICK_NAME": 1,
		}).X条件("id", 2).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one), 2)
		t.Assert(one["ID"], 2)
		t.Assert(one["NICKNAME"], "name_2")
	})
	// Struct
	gtest.C(t, func(t *gtest.T) {
		type T struct {
			ID       int
			NICKNAME int
		}
		one, err := db.X创建Model对象(table).X字段保留过滤(&T{
			ID:       0,
			NICKNAME: 0,
		}).X条件("id", 2).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one), 2)
		t.Assert(one["ID"], 2)
		t.Assert(one["NICKNAME"], "name_2")
	})
}

func Test_Model_NullField(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id       int
			Passport *string
		}
		data := g.Map{
			"id":       1,
			"passport": nil,
		}
		result, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)

		var user *User
		err = one.X取结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, data["id"])
		t.Assert(user.Passport, data["passport"])
	})
}

func Test_Model_HasTable(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		t.AssertNil(db.X取Core对象().X删除所有表查询缓存(ctx))
		result, err := db.X取Core对象().X是否存在表名(table)
		t.Assert(result, true)
		t.AssertNil(err)
	})

	gtest.C(t, func(t *gtest.T) {
		t.AssertNil(db.X取Core对象().X删除所有表查询缓存(ctx))
		result, err := db.X取Core对象().X是否存在表名("table12321")
		t.Assert(result, false)
		t.AssertNil(err)
	})
}

func Test_Model_HasField(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X是否存在字段("ID")
		t.Assert(result, true)
		t.AssertNil(err)
	})

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X是否存在字段("id123")
		t.Assert(result, false)
		t.AssertNil(err)
	})
}

func Test_Model_WhereIn(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件包含("id", g.Slice别名{1, 2, 3, 4}).X条件包含("id", g.Slice别名{3, 4, 5}).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0]["ID"], 3)
		t.Assert(result[1]["ID"], 4)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件包含("id", g.Slice别名{}).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 0)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X过滤空值条件().X条件包含("id", g.Slice别名{}).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
	})
}

func Test_Model_WhereNotIn(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件不包含("id", g.Slice别名{1, 2, 3, 4}).X条件不包含("id", g.Slice别名{3, 4, 5}).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 5)
		t.Assert(result[0]["ID"], 6)
		t.Assert(result[1]["ID"], 7)
	})
}

func Test_Model_WhereOrIn(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件或包含("id", g.Slice别名{1, 2, 3, 4}).X条件或包含("id", g.Slice别名{3, 4, 5}).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 5)
		t.Assert(result[0]["ID"], 1)
		t.Assert(result[4]["ID"], 5)
	})
}

func Test_Model_WhereOrNotIn(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件或不包含("id", g.Slice别名{1, 2, 3, 4}).X条件或不包含("id", g.Slice别名{3, 4, 5}).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 8)
		t.Assert(result[0]["ID"], 1)
		t.Assert(result[4]["ID"], 7)
	})
}

func Test_Model_WhereBetween(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件取范围("id", 1, 4).X条件取范围("id", 3, 5).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0]["ID"], 3)
		t.Assert(result[1]["ID"], 4)
	})
}

func Test_Model_WhereNotBetween(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件取范围以外("id", 2, 8).X条件取范围以外("id", 3, 100).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["ID"], 1)
	})
}

func Test_Model_WhereOrBetween(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件或取范围("id", 1, 4).X条件或取范围("id", 3, 5).X排序Desc("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 5)
		t.Assert(result[0]["ID"], 5)
		t.Assert(result[4]["ID"], 1)
	})
}

func Test_Model_WhereOrNotBetween(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	// db.SetDebug(true)
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件或取范围以外("id", 1, 4).X条件或取范围以外("id", 3, 5).X排序Desc("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 8)
		t.Assert(result[0]["ID"], 10)
		t.Assert(result[4]["ID"], 6)
	})
}

func Test_Model_WhereLike(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件模糊匹配("nickname", "name%").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(result[0]["ID"], 1)
		t.Assert(result[TableSize-1]["ID"], TableSize)
	})
}

func Test_Model_WhereNotLike(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件模糊匹配以外("nickname", "name%").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 0)
	})
}

func Test_Model_WhereOrLike(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件或模糊匹配("nickname", "namexxx%").X条件或模糊匹配("nickname", "name%").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(result[0]["ID"], 1)
		t.Assert(result[TableSize-1]["ID"], TableSize)
	})
}

func Test_Model_WhereOrNotLike(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件或模糊匹配以外("nickname", "namexxx%").X条件或模糊匹配以外("nickname", "name%").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(result[0]["ID"], 1)
		t.Assert(result[TableSize-1]["ID"], TableSize)
	})
}

func Test_Model_WhereNull(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件NULL值("nickname").X条件NULL值("passport").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 0)
	})
}

func Test_Model_WhereNotNull(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件非Null("nickname").X条件非Null("passport").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(result[0]["ID"], 1)
		t.Assert(result[TableSize-1]["ID"], TableSize)
	})
}

func Test_Model_WhereOrNull(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件或NULL值("nickname").X条件或NULL值("passport").X排序ASC("id").X排序随机().X查询()
		t.AssertNil(err)
		t.Assert(len(result), 0)
	})
}

func Test_Model_WhereOrNotNull(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件或非Null("nickname").X条件或非Null("passport").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(result[0]["ID"], 1)
		t.Assert(result[TableSize-1]["ID"], TableSize)
	})
}

func Test_Model_WhereLT(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件小于("id", 3).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0]["ID"], 1)
	})
}

func Test_Model_WhereLTE(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件小于等于("id", 3).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["ID"], 1)
	})
}

func Test_Model_WhereGT(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件大于("id", 8).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0]["ID"], 9)
	})
}

func Test_Model_WhereGTE(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件大于等于("id", 8).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["ID"], 8)
	})
}

func Test_Model_WhereOrLT(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件小于("id", 3).X条件或小于("id", 4).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["ID"], 1)
		t.Assert(result[2]["ID"], 3)
	})
}

func Test_Model_WhereOrLTE(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件小于等于("id", 3).X条件或小于等于("id", 4).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 4)
		t.Assert(result[0]["ID"], 1)
		t.Assert(result[3]["ID"], 4)
	})
}

func Test_Model_WhereOrGT(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件大于("id", 8).X条件或大于("id", 7).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 3)
		t.Assert(result[0]["ID"], 8)
	})
}

func Test_Model_WhereOrGTE(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件大于等于("id", 8).X条件或大于等于("id", 7).X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 4)
		t.Assert(result[0]["ID"], 7)
	})
}

func Test_Model_Min_Max_Avg_Sum(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X查询最小值("id")
		t.AssertNil(err)
		t.Assert(result, 1)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X查询最大值("id")
		t.AssertNil(err)
		t.Assert(result, TableSize)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X查询平均值("id")
		t.AssertNil(err)
		t.Assert(result, 5.5)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X查询求和("id")
		t.AssertNil(err)
		t.Assert(result, 55)
	})
}

func Test_Model_CountColumn(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X查询字段行数("id")
		t.AssertNil(err)
		t.Assert(result, TableSize)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件包含("id", g.Slice别名{1, 2, 3}).X查询字段行数("id")
		t.AssertNil(err)
		t.Assert(result, 3)
	})
}

func Test_Model_Raw(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.
			X原生SQL(fmt.Sprintf("select * from %s where id in (?)", table), g.Slice别名{1, 5, 7, 8, 9, 10}).
			X条件小于("id", 8).
			X条件包含("id", g.Slice别名{1, 2, 3, 4, 5, 6, 7}).
			X排序Desc("id").
			X查询()
		t.AssertNil(err)
		t.Assert(len(all), 3)
		t.Assert(all[0]["ID"], 7)
		t.Assert(all[1]["ID"], 5)
	})

	gtest.C(t, func(t *gtest.T) {
		count, err := db.
			X原生SQL(fmt.Sprintf("select * from %s where id in (?)", table), g.Slice别名{1, 5, 7, 8, 9, 10}).
			X条件小于("id", 8).
			X条件包含("id", g.Slice别名{1, 2, 3, 4, 5, 6, 7}).
			X排序Desc("id").
			X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(6))
	})
}

func Test_Model_Handler(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		m := db.X创建Model对象(table).X链式安全().X处理函数(
			func(m *gdb.Model) *gdb.Model {
				return m.X设置分页(0, 3)
			},
			func(m *gdb.Model) *gdb.Model {
				return m.X条件("id", g.Slice别名{1, 2, 3, 4, 5, 6})
			},
			func(m *gdb.Model) *gdb.Model {
				return m.X排序Desc("id")
			},
		)
		all, err := m.X查询()
		t.AssertNil(err)
		t.Assert(len(all), 3)
		t.Assert(all[0]["ID"], 6)
		t.Assert(all[2]["ID"], 4)
	})
}

func Test_Model_FieldCount(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X字段保留过滤("id").X字段追加计数("id", "total").X排序分组("id").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(all), TableSize)
		t.Assert(all[0]["ID"], 1)
		t.Assert(all[0]["total"].X取整数(), 1)
	})
}

func Test_Model_FieldMax(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X字段保留过滤("id").X字段追加最大值("id", "total").X排序分组("id").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(all), TableSize)
		t.Assert(all[0]["ID"], 1)
		t.Assert(all[0]["total"].X取整数(), 1)
	})
}

func Test_Model_FieldMin(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X字段保留过滤("id").X字段追加最小值("id", "total").X排序分组("id").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(all), TableSize)
		t.Assert(all[0]["ID"], 1)
		t.Assert(all[0]["total"].X取整数(), 1)
	})
}

func Test_Model_FieldAvg(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X字段保留过滤("id").X字段追加平均值("id", "total").X排序分组("id").X排序ASC("id").X查询()
		t.AssertNil(err)
		t.Assert(len(all), TableSize)
		t.Assert(all[0]["ID"], 1)
		t.Assert(all[0]["total"].X取整数(), 1)
	})
}

func Test_Model_OmitEmptyWhere(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	// Basic type where.
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X条件("id", 0).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(0))
	})
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X过滤空值条件().X条件("id", 0).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X过滤空值条件().X条件("id", 0).X条件("nickname", "").X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
	// Slice where.
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X条件("id", g.Slice别名{1, 2, 3}).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(3))
	})
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X条件("id", g.Slice别名{}).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(0))
	})
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X过滤空值条件().X条件("id", g.Slice别名{}).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X条件("id", g.Slice别名{}).X过滤空值条件().X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
	// Struct Where.
	gtest.C(t, func(t *gtest.T) {
		type Input struct {
			Id   []int
			Name []string
		}
		count, err := db.X创建Model对象(table).X条件(Input{}).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(0))
	})
	gtest.C(t, func(t *gtest.T) {
		type Input struct {
			Id   []int
			Name []string
		}
		count, err := db.X创建Model对象(table).X条件(Input{}).X过滤空值条件().X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
	// Map Where.
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X条件(g.Map{
			"id":       []int{},
			"nickname": []string{},
		}).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(0))
	})
	gtest.C(t, func(t *gtest.T) {
		type Input struct {
			Id []int
		}
		count, err := db.X创建Model对象(table).X条件(g.Map{
			"id": []int{},
		}).X过滤空值条件().X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
}

func Test_Model_WherePrefix(t *testing.T) {
	table1 := "table1"
	table2 := "table2"
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table1).
			X字段保留过滤并带前缀(table1, "*").
			X左连接相同字段(table2, "id").
			X条件带前缀(table2, g.Map{
				"id": g.Slice别名{1, 2},
			}).
			X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 2)
		t.Assert(r[0]["ID"], "1")
		t.Assert(r[1]["ID"], "2")
	})
}

func Test_Model_WhereOrPrefix(t *testing.T) {
	table1 := "table1"
	table2 := "table2"
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table1).
			X字段保留过滤并带前缀(table1, "*").
			X左连接相同字段(table2, "id").
			X条件或并带前缀(table1, g.Map{
				"id": g.Slice别名{1, 2},
			}).
			X条件或并带前缀(table2, g.Map{
				"id": g.Slice别名{8, 9},
			}).
			X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 4)
		t.Assert(r[0]["ID"], "1")
		t.Assert(r[1]["ID"], "2")
		t.Assert(r[2]["ID"], "8")
		t.Assert(r[3]["ID"], "9")
	})
}

func Test_Model_WherePrefixLike(t *testing.T) {
	table1 := "table1"
	table2 := "table2"
	createInitTable(table1)
	defer dropTable(table1)
	createInitTable(table2)
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table1).
			X字段保留过滤并带前缀(table1, "*").
			X左连接相同字段(table2, "id").
			X条件带前缀(table1, g.Map{
				"id": g.Slice别名{1, 2, 3},
			}).
			X条件带前缀(table2, g.Map{
				"id": g.Slice别名{3, 4, 5},
			}).
			X条件模糊匹配并带前缀(table2, "nickname", "name%").
			X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(len(r), 1)
		t.Assert(r[0]["ID"], "3")
	})
}

func Test_Model_AllAndCount(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, total, err := db.X创建Model对象(table).X排序("id").X设置条数(0, 3).X查询与行数(false)
		t.AssertNil(err)

		t.Assert(len(result), 3)
		t.Assert(total, TableSize)
	})
}

func Test_Model_ScanAndCount(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}

		users := make([]User, 0)
		total := 0

		err := db.X创建Model对象(table).X排序("id").X设置条数(0, 3).X查询与行数到指针(&users, &total, false)
		t.AssertNil(err)

		t.Assert(len(users), 3)
		t.Assert(total, TableSize)
	})
}

func Test_Model_Save(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id        int
			Passport  string
			Password  string
			NickName  string
			CreatedAt *gtime.Time
			UpdatedAt *gtime.Time
		}
		var (
			user   User
			count  int
			result sql.Result
			err    error
		)

		result, err = db.X创建Model对象(table).X设置数据(g.Map{
			"id":       1,
			"passport": "p1",
			"password": "15d55ad283aa400af464c76d713c07ad",
			"nickname": "n1",
		}).OnConflict("id").X插入并更新已存在()

		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)

		err = db.X创建Model对象(table).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 1)
		t.Assert(user.Passport, "p1")
		t.Assert(user.Password, "15d55ad283aa400af464c76d713c07ad")
		t.Assert(user.NickName, "n1")

				// 暂停1秒以确保更新的时间不同。 md5:733cf93f1135fed7
		time.Sleep(1 * time.Second)
		_, err = db.X创建Model对象(table).X设置数据(g.Map{
			"id":       1,
			"passport": "p1",
			"password": "25d55ad283aa400af464c76d713c07ad",
			"nickname": "n2",
		}).OnConflict("id").X插入并更新已存在()
		t.AssertNil(err)

		err = db.X创建Model对象(table).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Passport, "p1")
		t.Assert(user.Password, "25d55ad283aa400af464c76d713c07ad")
		t.Assert(user.NickName, "n2")
						// 检查created_at是否不等于updated_at. md5:1ce415b9de20266f
		t.AssertNE(user.CreatedAt, user.UpdatedAt)

		count, err = db.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(count, 1)
	})
}

func Test_Model_Replace(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		_, err := db.X创建Model对象(table).X设置数据(g.Map{
			"id":          1,
			"passport":    "t11",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "T11",
			"create_time": "2018-10-24 10:00:00",
		}).X插入并替换已存在()
		t.Assert(err, "Replace operation is not supported by mssql driver")
	})
}
