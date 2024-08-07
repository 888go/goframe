//go:build 屏蔽单元测试

// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package pgsql_test

import (
	"database/sql"
	"fmt"
	"testing"

	gdb "github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

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
		n, _ := result.RowsAffected()
		t.Assert(n, 1)

		result, err = db.X创建Model对象(table).X设置数据(g.Map{
			"id":          "2",
			"uid":         "2",
			"passport":    "t2",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "name_2",
			"create_time": gtime.X创建并按当前时间().String(),
		}).X插入()
		t.AssertNil(err)
		n, _ = result.RowsAffected()
		t.Assert(n, 1)

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
			Id:         3,
			Uid:        3,
			Passport:   "t3",
			Password:   "25d55ad283aa400af464c76d713c07ad",
			Nickname:   "name_3",
			CreateTime: gtime.X创建并按当前时间(),
		}).X插入()
		t.AssertNil(err)
		n, _ = result.RowsAffected()
		t.Assert(n, 1)
		value, err := db.X创建Model对象(table).X字段保留过滤("passport").X条件("id=3").X查询一条值() // model value
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
		n, _ = result.RowsAffected()
		t.Assert(n, 1)
		value, err = db.X创建Model对象(table).X字段保留过滤("passport").X条件("id=4").X查询一条值()
		t.AssertNil(err)
		t.Assert(value.String(), "t4")

		result, err = db.X创建Model对象(table).X条件("id>?", 1).X删除() // model delete
		t.AssertNil(err)
		n, _ = result.RowsAffected()
		t.Assert(n, 3)
	})
}

func Test_Model_One(t *testing.T) {
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
		result, err := db.X创建Model对象(table).X条件("id", "2").X删除()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
	})
}

func Test_Model_Update(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

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

func Test_Model_Count(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(TableSize))
	})
	gtest.C(t, func(t *gtest.T) {
		count, err := db.X创建Model对象(table).X字段排除过滤("id").X条件("id>8").X查询行数()
		t.AssertNil(err)
		t.Assert(count, int64(2))
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
			user   User
			count  int
			result sql.Result
			err    error
		)

		result, err = db.X创建Model对象(table).X设置数据(g.Map{
			"id":          1,
			"passport":    "p1",
			"password":    "pw1",
			"nickname":    "n1",
			"create_time": CreateTime,
		}).OnConflict("id").X插入并更新已存在()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)

		err = db.X创建Model对象(table).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 1)
		t.Assert(user.Passport, "p1")
		t.Assert(user.Password, "pw1")
		t.Assert(user.NickName, "n1")
		t.Assert(user.CreateTime.String(), CreateTime)

		_, err = db.X创建Model对象(table).X设置数据(g.Map{
			"id":          1,
			"passport":    "p1",
			"password":    "pw2",
			"nickname":    "n2",
			"create_time": CreateTime,
		}).OnConflict("id").X插入并更新已存在()
		t.AssertNil(err)

		err = db.X创建Model对象(table).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Passport, "p1")
		t.Assert(user.Password, "pw2")
		t.Assert(user.NickName, "n2")
		t.Assert(user.CreateTime.String(), CreateTime)

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
		t.Assert(err, "Replace operation is not supported by pgsql driver")
	})
}

func Test_Model_OnConflict(t *testing.T) {
	var (
		table      = fmt.Sprintf(`%s_%d`, TablePrefix+"test", gtime.X取时间戳纳秒())
		uniqueName = fmt.Sprintf(`%s_%d`, TablePrefix+"test_unique", gtime.X取时间戳纳秒())
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
		CREATE TABLE %s (
		   	id bigserial  NOT NULL,
		   	passport varchar(45) NOT NULL,
		   	password varchar(32) NOT NULL,
		   	nickname varchar(45) NOT NULL,
		   	create_time timestamp NOT NULL,
		   	PRIMARY KEY (id),
			CONSTRAINT %s UNIQUE ("passport", "password")
		) ;`, table, uniqueName,
	)); err != nil {
		gtest.Fatal(err)
	}
	defer dropTable(table)

	// string type 1.
	gtest.C(t, func(t *gtest.T) {
		data := g.Map{
			"id":          1,
			"passport":    "pp1",
			"password":    "pw1",
			"nickname":    "n1",
			"create_time": "2016-06-06",
		}
		_, err := db.X创建Model对象(table).OnConflict("passport,password").X设置数据(data).X插入并更新已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"])
		t.Assert(one["password"], data["password"])
		t.Assert(one["nickname"], "n1")
	})

	// string type 2.
	gtest.C(t, func(t *gtest.T) {
		data := g.Map{
			"id":          1,
			"passport":    "pp1",
			"password":    "pw1",
			"nickname":    "n1",
			"create_time": "2016-06-06",
		}
		_, err := db.X创建Model对象(table).OnConflict("passport", "password").X设置数据(data).X插入并更新已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"])
		t.Assert(one["password"], data["password"])
		t.Assert(one["nickname"], "n1")
	})

	// slice.
	gtest.C(t, func(t *gtest.T) {
		data := g.Map{
			"id":          1,
			"passport":    "pp1",
			"password":    "pw1",
			"nickname":    "n1",
			"create_time": "2016-06-06",
		}
		_, err := db.X创建Model对象(table).OnConflict(g.Slice别名{"passport", "password"}).X设置数据(data).X插入并更新已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件("id", 1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"])
		t.Assert(one["password"], data["password"])
		t.Assert(one["nickname"], "n1")
	})
}

func Test_Model_OnDuplicate(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	// string type 1.
	gtest.C(t, func(t *gtest.T) {
		data := g.Map{
			"id":          1,
			"passport":    "pp1",
			"password":    "pw1",
			"nickname":    "n1",
			"create_time": "2016-06-06",
		}
		_, err := db.X创建Model对象(table).OnConflict("id").X设置插入冲突更新字段("passport,password").X设置数据(data).X插入并更新已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"])
		t.Assert(one["password"], data["password"])
		t.Assert(one["nickname"], "name_1")
	})

	// string type 2.
	gtest.C(t, func(t *gtest.T) {
		data := g.Map{
			"id":          1,
			"passport":    "pp1",
			"password":    "pw1",
			"nickname":    "n1",
			"create_time": "2016-06-06",
		}
		_, err := db.X创建Model对象(table).OnConflict("id").X设置插入冲突更新字段("passport", "password").X设置数据(data).X插入并更新已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"])
		t.Assert(one["password"], data["password"])
		t.Assert(one["nickname"], "name_1")
	})

	// slice.
	gtest.C(t, func(t *gtest.T) {
		data := g.Map{
			"id":          1,
			"passport":    "pp1",
			"password":    "pw1",
			"nickname":    "n1",
			"create_time": "2016-06-06",
		}
		_, err := db.X创建Model对象(table).OnConflict("id").X设置插入冲突更新字段(g.Slice别名{"passport", "password"}).X设置数据(data).X插入并更新已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"])
		t.Assert(one["password"], data["password"])
		t.Assert(one["nickname"], "name_1")
	})

	// map.
	gtest.C(t, func(t *gtest.T) {
		data := g.Map{
			"id":          1,
			"passport":    "pp1",
			"password":    "pw1",
			"nickname":    "n1",
			"create_time": "2016-06-06",
		}
		_, err := db.X创建Model对象(table).OnConflict("id").X设置插入冲突更新字段(g.Map{
			"passport": "nickname",
			"password": "nickname",
		}).X设置数据(data).X插入并更新已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data["nickname"])
		t.Assert(one["password"], data["nickname"])
		t.Assert(one["nickname"], "name_1")
	})

	// map+raw.
	gtest.C(t, func(t *gtest.T) {
		data := g.MapStrStr{
			"id":          "1",
			"passport":    "pp1",
			"password":    "pw1",
			"nickname":    "n1",
			"create_time": "2016-06-06",
		}
		_, err := db.X创建Model对象(table).OnConflict("id").X设置插入冲突更新字段(g.Map{
			"passport": gdb.Raw("CONCAT(EXCLUDED.passport, '1')"),
			"password": gdb.Raw("CONCAT(EXCLUDED.password, '2')"),
		}).X设置数据(data).X插入并更新已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"]+"1")
		t.Assert(one["password"], data["password"]+"2")
		t.Assert(one["nickname"], "name_1")
	})
}

func Test_Model_OnDuplicateEx(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	// string type 1.
	gtest.C(t, func(t *gtest.T) {
		data := g.Map{
			"id":          1,
			"passport":    "pp1",
			"password":    "pw1",
			"nickname":    "n1",
			"create_time": "2016-06-06",
		}
		_, err := db.X创建Model对象(table).OnConflict("id").X设置插入冲突不更新字段("nickname,create_time").X设置数据(data).X插入并更新已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"])
		t.Assert(one["password"], data["password"])
		t.Assert(one["nickname"], "name_1")
	})

	// string type 2.
	gtest.C(t, func(t *gtest.T) {
		data := g.Map{
			"id":          1,
			"passport":    "pp1",
			"password":    "pw1",
			"nickname":    "n1",
			"create_time": "2016-06-06",
		}
		_, err := db.X创建Model对象(table).OnConflict("id").X设置插入冲突不更新字段("nickname", "create_time").X设置数据(data).X插入并更新已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"])
		t.Assert(one["password"], data["password"])
		t.Assert(one["nickname"], "name_1")
	})

	// slice.
	gtest.C(t, func(t *gtest.T) {
		data := g.Map{
			"id":          1,
			"passport":    "pp1",
			"password":    "pw1",
			"nickname":    "n1",
			"create_time": "2016-06-06",
		}
		_, err := db.X创建Model对象(table).OnConflict("id").X设置插入冲突不更新字段(g.Slice别名{"nickname", "create_time"}).X设置数据(data).X插入并更新已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"])
		t.Assert(one["password"], data["password"])
		t.Assert(one["nickname"], "name_1")
	})

	// map.
	gtest.C(t, func(t *gtest.T) {
		data := g.Map{
			"id":          1,
			"passport":    "pp1",
			"password":    "pw1",
			"nickname":    "n1",
			"create_time": "2016-06-06",
		}
		_, err := db.X创建Model对象(table).OnConflict("id").X设置插入冲突不更新字段(g.Map{
			"nickname":    "nickname",
			"create_time": "nickname",
		}).X设置数据(data).X插入并更新已存在()
		t.AssertNil(err)
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["passport"], data["passport"])
		t.Assert(one["password"], data["password"])
		t.Assert(one["nickname"], "name_1")
	})
}
