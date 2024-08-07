//go:build 屏蔽单元测试

// 版权所有 2019 gf 作者（https://github.com/gogf/gf）。保留所有权利。
//
// 此源代码形式受麻省理工学院（MIT）许可证的条款约束。
// 如果未随此文件一起分发MIT许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:47e609239e0cb2bc

package oracle_test

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
	"time"

	gmap "github.com/888go/goframe/container/gmap"
	gdb "github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
	gutil "github.com/888go/goframe/util/gutil"
)

func Test_Model_InnerJoin(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		table1 := createInitTable("user1")
		table2 := createInitTable("user2")

		defer dropTable(table1)
		defer dropTable(table2)

		res, err := db.X创建Model对象(table1).X条件("id > ?", 5).X删除()
		if err != nil {
			t.Fatal(err)
		}

		n, err := res.RowsAffected()
		if err != nil {
			t.Fatal(err)
		}

		t.Assert(n, 5)

		result, err := db.X创建Model对象(table1+" u1").X内连接(table2+" u2", "u1.id = u2.id").X排序("u1.id").X查询()
		if err != nil {
			t.Fatal(err)
		}

		t.Assert(len(result), 5)

		result, err = db.X创建Model对象(table1+" u1").X内连接(table2+" u2", "u1.id = u2.id").X条件("u1.id > ?", 1).X排序("u1.id").X查询()
		if err != nil {
			t.Fatal(err)
		}

		t.Assert(len(result), 4)
	})
}

func Test_Model_LeftJoin(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		table1 := createInitTable("user1")
		table2 := createInitTable("user2")

		defer dropTable(table1)
		defer dropTable(table2)

		res, err := db.X创建Model对象(table2).X条件("id > ?", 3).X删除()
		if err != nil {
			t.Fatal(err)
		}

		n, err := res.RowsAffected()
		if err != nil {
			t.Fatal(err)
		} else {
			t.Assert(n, 7)
		}

		result, err := db.X创建Model对象(table1+" u1").X左连接(table2+" u2", "u1.id = u2.id").X查询()
		if err != nil {
			t.Fatal(err)
		}

		t.Assert(len(result), 10)

		result, err = db.X创建Model对象(table1+" u1").X左连接(table2+" u2", "u1.id = u2.id").X条件("u1.id > ? ", 2).X查询()
		if err != nil {
			t.Fatal(err)
		}

		t.Assert(len(result), 8)
	})
}

func Test_Model_RightJoin(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		table1 := createInitTable("user1")
		table2 := createInitTable("user2")

		defer dropTable(table1)
		defer dropTable(table2)

		res, err := db.X创建Model对象(table1).X条件("id > ?", 3).X删除()
		if err != nil {
			t.Fatal(err)
		}

		n, err := res.RowsAffected()
		if err != nil {
			t.Fatal(err)
		}

		t.Assert(n, 7)

		result, err := db.X创建Model对象(table1+" u1").X右连接(table2+" u2", "u1.id = u2.id").X查询()
		if err != nil {
			t.Fatal(err)
		}
		t.Assert(len(result), 10)

		result, err = db.X创建Model对象(table1+" u1").X右连接(table2+" u2", "u1.id = u2.id").X条件("u1.id > 2").X查询()
		if err != nil {
			t.Fatal(err)
		}
		t.Assert(len(result), 1)
	})
}

func Test_Page(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	result, err := db.X创建Model对象(table).X设置分页(1, 2).X排序("ID").X查询()
	gtest.Assert(err, nil)
	fmt.Println("page:1--------", result)
	gtest.Assert(len(result), 2)
	gtest.Assert(result[0]["ID"], 1)
	gtest.Assert(result[1]["ID"], 2)

	result, err = db.X创建Model对象(table).X设置分页(2, 2).X排序("ID").X查询()
	gtest.Assert(err, nil)
	fmt.Println("page: 2--------", result)
	gtest.Assert(len(result), 2)
	gtest.Assert(result[0]["ID"], 3)
	gtest.Assert(result[1]["ID"], 4)

	result, err = db.X创建Model对象(table).X设置分页(3, 2).X排序("ID").X查询()
	gtest.Assert(err, nil)
	fmt.Println("page:3 --------", result)
	gtest.Assert(len(result), 2)
	gtest.Assert(result[0]["ID"], 5)

	result, err = db.X创建Model对象(table).X设置分页(2, 3).X查询()
	gtest.Assert(err, nil)
	gtest.Assert(len(result), 3)
	gtest.Assert(result[0]["ID"], 4)
	gtest.Assert(result[1]["ID"], 5)
	gtest.Assert(result[2]["ID"], 6)
}

func Test_Model_Insert(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		user := db.X创建Model对象(table)
		result, err := user.X设置数据(g.Map{
			"ID":          1,
			"UID":         1,
			"PASSPORT":    "t1",
			"PASSWORD":    "25d55ad283aa400af464c76d713c07ad",
			"NICKNAME":    "name_1",
			"SALARY":      2675.11,
			"CREATE_TIME": gtime.X创建并按当前时间().String(),
		}).X插入()
		t.AssertNil(err)

		result, err = db.X创建Model对象(table).X设置数据(g.Map{
			"ID":          "2",
			"UID":         "2",
			"PASSPORT":    "t2",
			"PASSWORD":    "25d55ad283aa400af464c76d713c07ad",
			"NICKNAME":    "name_2",
			"SALARY":      2675.12,
			"CREATE_TIME": gtime.X创建并按当前时间().String(),
		}).X插入()
		t.AssertNil(err)

		type User struct {
			Id         int         `gconv:"ID"`
			Uid        int         `gconv:"uid"`
			Passport   string      `json:"PASSPORT"`
			Password   string      `gconv:"PASSWORD"`
			Nickname   string      `gconv:"NICKNAME"`
			Salary     float64     `gconv:"SALARY"`
			CreateTime *gtime.Time `json:"CREATE_TIME"`
		}
		// Model inserting.
		result, err = db.X创建Model对象(table).X设置数据(User{
			Id:         3,
			Uid:        3,
			Passport:   "t3",
			Password:   "25d55ad283aa400af464c76d713c07ad",
			Nickname:   "name_3",
			Salary:     2675.13,
			CreateTime: gtime.X创建并按当前时间(),
		}).X插入()
		t.AssertNil(err)

		value, err := db.X创建Model对象(table).X字段保留过滤("PASSPORT").X条件("id=3").X查询一条值()
		t.AssertNil(err)
		t.Assert(value.String(), "t3")

		result, err = db.X创建Model对象(table).X设置数据(&User{
			Id:         4,
			Uid:        4,
			Passport:   "t4",
			Password:   "25d55ad283aa400af464c76d713c07ad",
			Nickname:   "T4",
			Salary:     2675.14,
			CreateTime: gtime.X创建并按当前时间(),
		}).X插入()
		t.AssertNil(err)

		value, err = db.X创建Model对象(table).X字段保留过滤("PASSPORT").X条件("id=4").X查询一条值()
		t.AssertNil(err)
		t.Assert(value.String(), "t4")

		result, err = db.X创建Model对象(table).X条件("id>?", 1).X删除()
		t.AssertNil(err)
		_, _ = result.RowsAffected()

	})
}

// 关于gf（Go Foundation）框架的问题#3286. md5:99d564c3a7918995
func Test_Model_Insert_Raw(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		_, err := db.X创建Model对象(table).X设置数据(g.Map{
			"ID":          1,
			"UID":         1,
			"PASSPORT":    "t1",
			"PASSWORD":    "25d55ad283aa400af464c76d713c07ad",
			"NICKNAME":    gdb.Raw("name_1"),
			"SALARY":      2675.11,
			"CREATE_TIME": gtime.X创建并按当前时间().String(),
		}).X插入()
		t.AssertNil(err)

		value, err := db.X创建Model对象(table).X字段保留过滤("PASSPORT").X条件("id=1").X查询一条值()
		t.AssertNil(err)
		t.Assert(value.String(), "t1")
	})
}

func Test_Model_Insert_Time(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		data := g.Map{
			"ID":          1,
			"PASSPORT":    "t1",
			"PASSWORD":    "p1",
			"NICKNAME":    "n1",
			"SALARY":      2675.11,
			"CREATE_TIME": "2020-10-10 20:09:18",
		}
		_, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)

		one, err := db.X创建Model对象(table).X查询一条("ID", 1)
		t.AssertNil(err)
		t.Assert(one["PASSPORT"].String(), data["PASSPORT"])
		t.Assert(one["CREATE_TIME"].String(), "2020-10-10 20:09:18")
		t.Assert(one["NICKNAME"].String(), data["NICKNAME"])
		t.Assert(one["SALARY"].X取小数64位(), data["SALARY"])
	})
}

func Test_Model_Batch(t *testing.T) {
	// batch insert
	gtest.C(t, func(t *gtest.T) {
		table := createTable()
		defer dropTable(table)
		_, err := db.X创建Model对象(table).X设置数据(g.Map切片{
			{
				"ID":          2,
				"uid":         2,
				"PASSPORT":    "t2",
				"PASSWORD":    "25d55ad283aa400af464c76d713c07ad",
				"NICKNAME":    "name_2",
				"SALARY":      2675.12,
				"CREATE_TIME": gtime.X创建并按当前时间().String(),
			},
			{
				"ID":          3,
				"uid":         3,
				"PASSPORT":    "t3",
				"PASSWORD":    "25d55ad283aa400af464c76d713c07ad",
				"NICKNAME":    "name_3",
				"SALARY":      2675.13,
				"CREATE_TIME": gtime.X创建并按当前时间().String(),
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
		result, err := db.X创建Model对象(table).X设置数据("PASSPORT", "user_22").X条件("passport=?", "user_2").X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X设置数据("PASSPORT", "user_2").X条件("passport='user_22'").X更新()
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
		result, err := db.X创建Model对象(table).X字段保留过滤("PASSPORT").X设置数据(g.Map{
			"PASSPORT": "user_44",
			"none":     "none",
		}).X条件("passport='user_4'").X更新()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
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
		record, err := db.X创建Model对象(table).X条件("ID", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(record["NICKNAME"].String(), "name_1")
	})

	gtest.C(t, func(t *gtest.T) {
		record, err := db.X创建Model对象(table).X条件("ID", 0).X查询一条()
		t.AssertNil(err)
		t.Assert(record, nil)
	})
}

func Test_Model_Value(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		value, err := db.X创建Model对象(table).X字段保留过滤("NICKNAME").X条件("ID", 1).X查询一条值()
		t.AssertNil(err)
		t.Assert(value.String(), "name_1")
	})

	gtest.C(t, func(t *gtest.T) {
		value, err := db.X创建Model对象(table).X字段保留过滤("NICKNAME").X条件("ID", 0).X查询一条值()
		t.AssertNil(err)
		t.Assert(value, nil)
	})
}

func Test_Model_Array(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X条件("ID", g.Slice别名{1, 2, 3}).X查询()
		t.AssertNil(err)
		for k, v := range all.X取字段切片("ID") {
			t.Assert(v.X取整数(), k+1)
		}
		t.Assert(all.X取字段切片("NICKNAME"), g.Slice别名{"name_1", "name_2", "name_3"})
	})
	gtest.C(t, func(t *gtest.T) {
		array, err := db.X创建Model对象(table).X字段保留过滤("NICKNAME").X条件("ID", g.Slice别名{1, 2, 3}).X查询切片()
		t.AssertNil(err)
		t.Assert(array, g.Slice别名{"name_1", "name_2", "name_3"})
	})
	gtest.C(t, func(t *gtest.T) {
		array, err := db.X创建Model对象(table).X查询切片("NICKNAME", "ID", g.Slice别名{1, 2, 3})
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
		count, err := db.X创建Model对象(table).X字段排除过滤("ID").X条件("id>8").X查询行数()
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
		Salary     float64
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
			Salary     float64
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
			Salary     float64
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
			Salary     float64
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
			Salary     float64
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
			Salary     float64
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
			Salary     float64
			CreateTime *gtime.Time
		}
		var user *User
		err := db.X创建Model对象(table).X条件("id=-1").X查询到结构体指针(&user)
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
			Salary     float64
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
			Salary     float64
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
			Salary     float64
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
			Salary     float64
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
			Salary     float64
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
		result, err := db.X创建Model对象(table).X条件(g.Slice别名{"ID", 3}).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件(g.Slice别名{"ID", 3, "NICKNAME", "name_3"}).X查询一条()
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
			"ID":       g.Slice别名{1, 2, 3},
			"PASSPORT": g.Slice别名{"user_2", "user_3"},
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
		result, err := db.X创建Model对象(table).X条件("ID", 3).X查询一条()
		t.AssertNil(err)
		t.AssertGT(len(result), 0)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("ID", 3).X条件("NICKNAME", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("ID", 3).X条件("NICKNAME", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("ID", 30).X条件或("NICKNAME", "name_3").X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("ID", 30).X条件或("NICKNAME", "name_3").X条件("id>?", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(result["ID"].X取整数(), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件("ID", 30).X条件或("NICKNAME", "name_3").X条件("id>", 1).X查询一条()
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
		result, err := db.X创建Model对象(table).X条件(g.Map{"ID": 3, "NICKNAME": "name_3"}).X查询一条()
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
		result, err := db.X创建Model对象(table).X条件(gmap.X创建并从Map(g.MapAnyAny{"ID": 3, "NICKNAME": "name_3"})).X查询一条()
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
		result, err := db.X创建Model对象(table).X条件(gmap.X创建链表Map并从Map(g.MapAnyAny{"ID": 3, "NICKNAME": "name_3"})).X查询一条()
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
		result, err := db.X创建Model对象(table).X条件(gmap.X创建红黑树Map并从Map(gutil.X比较文本, g.MapAnyAny{"ID": 3, "NICKNAME": "name_3"})).X查询一条()
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
			"ID":                 g.Slice别名{1, 2, 3},
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
			"ID":       g.Slice别名{1, 3},
			"NICKNAME": "name_3",
		}).X排序("id ASC").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 1)
		t.Assert(result[0]["ID"].X取整数(), 3)
	})
	// slice + struct
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Ids      []int  `json:"ID"`
			Nickname string `gconv:"NICKNAME"`
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

func Test_Model_Having(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X字段保留过滤("id, count(*)").X条件("id > 1").X排序分组("ID").X设置分组条件("id > 8").X查询()
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

// not support
/*
func Test_Model_Min_Max(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		value, err := db.Model(table, "t").Fields("min(t.id)").Where("id > 1").Value()
		t.AssertNil(err)
		t.Assert(value.Int(), 2)
	})
	gtest.C(t, func(t *gtest.T) {
		value, err := db.Model(table, "t").Fields("max(t.id)").Where("id > 1").Value()
		t.AssertNil(err)
		t.Assert(value.Int(), 10)
	})
}
*/
func Test_Model_HasTable(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	// db.SetDebug(true)
	gtest.C(t, func(t *gtest.T) {
		t.AssertNil(db.X取Core对象().X删除所有表查询缓存(ctx))
		result, err := db.X取Core对象().X是否存在表名(strings.ToUpper(table))
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
		result, err := db.X创建Model对象(table).X条件包含("ID", g.Slice别名{1, 2, 3, 4}).X条件包含("ID", g.Slice别名{3, 4, 5}).X排序ASC("ID").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 2)
		t.Assert(result[0]["ID"], 3)
		t.Assert(result[1]["ID"], 4)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件包含("ID", g.Slice别名{}).X排序ASC("ID").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 0)
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X过滤空值条件().X条件包含("ID", g.Slice别名{}).X排序ASC("ID").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
	})
}

func Test_Model_WhereNotIn(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件不包含("ID", g.Slice别名{1, 2, 3, 4}).X条件不包含("ID", g.Slice别名{3, 4, 5}).X排序ASC("ID").X查询()
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
		result, err := db.X创建Model对象(table).X条件或包含("ID", g.Slice别名{1, 2, 3, 4}).X条件或包含("ID", g.Slice别名{3, 4, 5}).X排序ASC("ID").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 5)
		t.Assert(result[0]["ID"], 1)
		t.Assert(result[4]["ID"], 5)
	})
}

func Test_Model_WhereLike(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件模糊匹配("NICKNAME", "name%").X排序ASC("ID").X查询()
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
		result, err := db.X创建Model对象(table).X条件模糊匹配以外("NICKNAME", "name%").X排序ASC("ID").X查询()
		t.AssertNil(err)
		t.Assert(len(result), 0)
	})
}

func Test_Model_WhereOrLike(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		result, err := db.X创建Model对象(table).X条件或模糊匹配("NICKNAME", "namexxx%").X条件或模糊匹配("NICKNAME", "name%").X排序ASC("ID").X查询()
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
		result, err := db.X创建Model对象(table).X条件或模糊匹配以外("NICKNAME", "namexxx%").X条件或模糊匹配以外("NICKNAME", "name%").X排序ASC("ID").X查询()
		t.AssertNil(err)
		t.Assert(len(result), TableSize)
		t.Assert(result[0]["ID"], 1)
		t.Assert(result[TableSize-1]["ID"], TableSize)
	})
}

func Test_Model_Save(t *testing.T) {
	table := createTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var (
			user       User
			count      int
			result     sql.Result
			createTime = gtime.X创建并按当前时间().X取格式文本("Y-m-d")
			err        error
		)

		result, err = db.X创建Model对象(table).X设置数据(g.Map{
			"id":          1,
			"passport":    "p1",
			"password":    "15d55ad283aa400af464c76d713c07ad",
			"nickname":    "n1",
			"create_time": createTime,
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
		t.Assert(user.CreateTime.X取格式文本("Y-m-d"), createTime)

		_, err = db.X创建Model对象(table).X设置数据(g.Map{
			"id":          1,
			"passport":    "p1",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "n2",
			"create_time": createTime,
		}).OnConflict("id").X插入并更新已存在()
		t.AssertNil(err)

		err = db.X创建Model对象(table).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Passport, "p1")
		t.Assert(user.Password, "25d55ad283aa400af464c76d713c07ad")
		t.Assert(user.NickName, "n2")
		t.Assert(user.CreateTime.X取格式文本("Y-m-d"), createTime)

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
		t.Assert(err, "Replace operation is not supported by oracle driver")
	})
}

/* not support the "AS"
func Test_Model_Raw(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.
			Raw(fmt.Sprintf("select * from %s where id in (?)", table), g.Slice{1, 5, 7, 8, 9, 10}).
			WhereLT("ID", 8).
			WhereIn("ID", g.Slice{1, 2, 3, 4, 5, 6, 7}).
			OrderDesc("ID").
			All()
		t.AssertNil(err)
		t.Assert(len(all), 3)
		t.Assert(all[0]["ID"], 7)
		t.Assert(all[1]["ID"], 5)
	})

	gtest.C(t, func(t *gtest.T) {
		count, err := db.
			Raw(fmt.Sprintf("select * from %s where id in (?)", table), g.Slice{1, 5, 7, 8, 9, 10}).
			WhereLT("ID", 8).
			WhereIn("ID", g.Slice{1, 2, 3, 4, 5, 6, 7}).
			OrderDesc("ID").
			Count()
		t.AssertNil(err)
		t.Assert(count, 6)
	})
}

func Test_Model_FieldCount(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.Model(table).Fields("ID").FieldCount("ID", "total").Group("ID").OrderAsc("ID").All()
		t.AssertNil(err)
		t.Assert(len(all), TableSize)
		t.Assert(all[0]["ID"], 1)
		t.Assert(all[0]["total"], 1)
	})
}

func Test_Model_FieldMax(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		all, err := db.Model(table).Fields("ID").FieldMax("ID", "total").Group("ID").OrderAsc("ID").All()
		t.AssertNil(err)
		t.Assert(len(all), TableSize)
		t.Assert(all[0]["ID"], 1)
		t.Assert(all[0]["total"], 1)
	})
}*/
