// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package mysql_test

import (
	"database/sql"
	"reflect"
	"testing"

	gdb "github.com/888go/goframe/database/gdb"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

func Test_Model_Embedded_Insert(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type Base struct {
			Id         int    `json:"id"`
			Uid        int    `json:"uid"`
			CreateTime string `json:"create_time"`
		}
		type User struct {
			Base
			Passport string `json:"passport"`
			Password string `json:"password"`
			Nickname string `json:"nickname"`
		}
		result, err := db.X创建Model对象(table).X设置数据(User{
			Passport: "john-test",
			Password: "123456",
			Nickname: "John",
			Base: Base{
				Id:         100,
				Uid:        100,
				CreateTime: gtime.X创建并按当前时间().String(),
			},
		}).X插入()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)
		value, err := db.X创建Model对象(table).X字段保留过滤("passport").X条件("id=100").X查询一条值()
		t.AssertNil(err)
		t.Assert(value.String(), "john-test")
	})
}

func Test_Model_Embedded_MapToStruct(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type Ids struct {
			Id  int `json:"id"`
			Uid int `json:"uid"`
		}
		type Base struct {
			Ids
			CreateTime string `json:"create_time"`
		}
		type User struct {
			Base
			Passport string `json:"passport"`
			Password string `json:"password"`
			Nickname string `json:"nickname"`
		}
		data := g.Map{
			"id":          100,
			"uid":         101,
			"passport":    "t1",
			"password":    "123456",
			"nickname":    "T1",
			"create_time": gtime.X创建并按当前时间().String(),
		}
		result, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)
		n, _ := result.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).X条件("id=100").X查询一条()
		t.AssertNil(err)

		user := new(User)

		t.Assert(one.X取结构体指针(user), nil)
		t.Assert(user.Id, data["id"])
		t.Assert(user.Passport, data["passport"])
		t.Assert(user.Password, data["password"])
		t.Assert(user.Nickname, data["nickname"])
		t.Assert(user.CreateTime, data["create_time"])
	})
}

func Test_Struct_Pointer_Attribute(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	type User struct {
		Id       *int
		Passport *string
		Password *string
		Nickname string
	}

	gtest.C(t, func(t *gtest.T) {
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		user := new(User)
		err = one.X取结构体指针(user)
		t.AssertNil(err)
		t.Assert(*user.Id, 1)
		t.Assert(*user.Passport, "user_1")
		t.Assert(*user.Password, "pass_1")
		t.Assert(user.Nickname, "name_1")
	})
	gtest.C(t, func(t *gtest.T) {
		user := new(User)
		err := db.X创建Model对象(table).X查询到结构体指针(user, "id=1")
		t.AssertNil(err)
		t.Assert(*user.Id, 1)
		t.Assert(*user.Passport, "user_1")
		t.Assert(*user.Password, "pass_1")
		t.Assert(user.Nickname, "name_1")
	})
	gtest.C(t, func(t *gtest.T) {
		var user *User
		err := db.X创建Model对象(table).X查询到结构体指针(&user, "id=1")
		t.AssertNil(err)
		t.Assert(*user.Id, 1)
		t.Assert(*user.Passport, "user_1")
		t.Assert(*user.Password, "pass_1")
		t.Assert(user.Nickname, "name_1")
	})
}

func Test_Structs_Pointer_Attribute(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	type User struct {
		Id       *int
		Passport *string
		Password *string
		Nickname string
	}
	// All
	gtest.C(t, func(t *gtest.T) {
		one, err := db.X创建Model对象(table).X查询("id < 3")
		t.AssertNil(err)
		users := make([]User, 0)
		err = one.X取切片结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(*users[0].Id, 1)
		t.Assert(*users[0].Passport, "user_1")
		t.Assert(*users[0].Password, "pass_1")
		t.Assert(users[0].Nickname, "name_1")
	})
	gtest.C(t, func(t *gtest.T) {
		one, err := db.X创建Model对象(table).X查询("id < 3")
		t.AssertNil(err)
		users := make([]*User, 0)
		err = one.X取切片结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(*users[0].Id, 1)
		t.Assert(*users[0].Passport, "user_1")
		t.Assert(*users[0].Password, "pass_1")
		t.Assert(users[0].Nickname, "name_1")
	})
	gtest.C(t, func(t *gtest.T) {
		var users []User
		one, err := db.X创建Model对象(table).X查询("id < 3")
		t.AssertNil(err)
		err = one.X取切片结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(*users[0].Id, 1)
		t.Assert(*users[0].Passport, "user_1")
		t.Assert(*users[0].Password, "pass_1")
		t.Assert(users[0].Nickname, "name_1")
	})
	gtest.C(t, func(t *gtest.T) {
		var users []*User
		one, err := db.X创建Model对象(table).X查询("id < 3")
		t.AssertNil(err)
		err = one.X取切片结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(*users[0].Id, 1)
		t.Assert(*users[0].Passport, "user_1")
		t.Assert(*users[0].Password, "pass_1")
		t.Assert(users[0].Nickname, "name_1")
	})
	// Structs
	gtest.C(t, func(t *gtest.T) {
		users := make([]User, 0)
		err := db.X创建Model对象(table).X查询到结构体指针(&users, "id < 3")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(*users[0].Id, 1)
		t.Assert(*users[0].Passport, "user_1")
		t.Assert(*users[0].Password, "pass_1")
		t.Assert(users[0].Nickname, "name_1")
	})
	gtest.C(t, func(t *gtest.T) {
		users := make([]*User, 0)
		err := db.X创建Model对象(table).X查询到结构体指针(&users, "id < 3")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(*users[0].Id, 1)
		t.Assert(*users[0].Passport, "user_1")
		t.Assert(*users[0].Password, "pass_1")
		t.Assert(users[0].Nickname, "name_1")
	})
	gtest.C(t, func(t *gtest.T) {
		var users []User
		err := db.X创建Model对象(table).X查询到结构体指针(&users, "id < 3")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(*users[0].Id, 1)
		t.Assert(*users[0].Passport, "user_1")
		t.Assert(*users[0].Password, "pass_1")
		t.Assert(users[0].Nickname, "name_1")
	})
	gtest.C(t, func(t *gtest.T) {
		var users []*User
		err := db.X创建Model对象(table).X查询到结构体指针(&users, "id < 3")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(*users[0].Id, 1)
		t.Assert(*users[0].Passport, "user_1")
		t.Assert(*users[0].Password, "pass_1")
		t.Assert(users[0].Nickname, "name_1")
	})
}

func Test_Struct_Empty(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	type User struct {
		Id       int
		Passport string
		Password string
		Nickname string
	}

	gtest.C(t, func(t *gtest.T) {
		user := new(User)
		err := db.X创建Model对象(table).X条件("id=100").X查询到结构体指针(user)
		t.Assert(err, sql.ErrNoRows)
		t.AssertNE(user, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		one, err := db.X创建Model对象(table).X条件("id=100").X查询一条()
		t.AssertNil(err)
		var user *User
		t.Assert(one.X取结构体指针(&user), nil)
		t.Assert(user, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		var user *User
		err := db.X创建Model对象(table).X条件("id=100").X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user, nil)
	})
}

func Test_Structs_Empty(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	type User struct {
		Id       int
		Passport string
		Password string
		Nickname string
	}

	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X条件("id>100").X查询()
		t.AssertNil(err)
		users := make([]User, 0)
		t.Assert(all.X取切片结构体指针(&users), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X条件("id>100").X查询()
		t.AssertNil(err)
		users := make([]User, 10)
		t.Assert(all.X取切片结构体指针(&users), sql.ErrNoRows)
	})
	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X条件("id>100").X查询()
		t.AssertNil(err)
		var users []User
		t.Assert(all.X取切片结构体指针(&users), nil)
	})

	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X条件("id>100").X查询()
		t.AssertNil(err)
		users := make([]*User, 0)
		t.Assert(all.X取切片结构体指针(&users), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X条件("id>100").X查询()
		t.AssertNil(err)
		users := make([]*User, 10)
		t.Assert(all.X取切片结构体指针(&users), sql.ErrNoRows)
	})
	gtest.C(t, func(t *gtest.T) {
		all, err := db.X创建Model对象(table).X条件("id>100").X查询()
		t.AssertNil(err)
		var users []*User
		t.Assert(all.X取切片结构体指针(&users), nil)
	})
}

type MyTime struct {
	gtime.Time
}

type MyTimeSt struct {
	CreateTime MyTime
}

func (st *MyTimeSt) UnmarshalValue(v interface{}) error {
	m := gconv.X取Map(v)
	t, err := gtime.X转换文本(gconv.String(m["create_time"]))
	if err != nil {
		return err
	}
	st.CreateTime = MyTime{*t}
	return nil
}

func Test_Model_Scan_CustomType_Time(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		st := new(MyTimeSt)
		err := db.X创建Model对象(table).X字段保留过滤("create_time").X查询到结构体指针(st)
		t.AssertNil(err)
		t.Assert(st.CreateTime.String(), "2018-10-24 10:00:00")
	})
	gtest.C(t, func(t *gtest.T) {
		var stSlice []*MyTimeSt
		err := db.X创建Model对象(table).X字段保留过滤("create_time").X查询到结构体指针(&stSlice)
		t.AssertNil(err)
		t.Assert(len(stSlice), TableSize)
		t.Assert(stSlice[0].CreateTime.String(), "2018-10-24 10:00:00")
		t.Assert(stSlice[9].CreateTime.String(), "2018-10-24 10:00:00")
	})
}

func Test_Model_Scan_CustomType_String(t *testing.T) {
	type MyString string

	type MyStringSt struct {
		Passport MyString
	}

	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		st := new(MyStringSt)
		err := db.X创建Model对象(table).X字段保留过滤("Passport").X条件并识别主键(1).X查询到结构体指针(st)
		t.AssertNil(err)
		t.Assert(st.Passport, "user_1")
	})
	gtest.C(t, func(t *gtest.T) {
		var sts []MyStringSt
		err := db.X创建Model对象(table).X字段保留过滤("Passport").X排序("id asc").X查询到结构体指针(&sts)
		t.AssertNil(err)
		t.Assert(len(sts), TableSize)
		t.Assert(sts[0].Passport, "user_1")
	})
}

type User struct {
	Id         int
	Passport   string
	Password   string
	Nickname   string
	CreateTime *gtime.Time
}

func (user *User) UnmarshalValue(value interface{}) error {
	if record, ok := value.(gdb.Record); ok {
		*user = User{
			Id:         record["id"].X取整数(),
			Passport:   record["passport"].String(),
			Password:   "",
			Nickname:   record["nickname"].String(),
			CreateTime: record["create_time"].X取gtime时间类(),
		}
		return nil
	}
	return gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, `unsupported value type for UnmarshalValue: %v`, reflect.TypeOf(value))
}

func Test_Model_Scan_UnmarshalValue(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)
	gtest.C(t, func(t *gtest.T) {
		var users []*User
		err := db.X创建Model对象(table).X排序("id asc").X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), TableSize)
		t.Assert(users[0].Id, 1)
		t.Assert(users[0].Passport, "user_1")
		t.Assert(users[0].Password, "")
		t.Assert(users[0].Nickname, "name_1")
		t.Assert(users[0].CreateTime.String(), CreateTime)

		t.Assert(users[9].Id, 10)
		t.Assert(users[9].Passport, "user_10")
		t.Assert(users[9].Password, "")
		t.Assert(users[9].Nickname, "name_10")
		t.Assert(users[9].CreateTime.String(), CreateTime)
	})
}

func Test_Model_Scan_Map(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		var users []*User
		err := db.X创建Model对象(table).X排序("id asc").X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), TableSize)
		t.Assert(users[0].Id, 1)
		t.Assert(users[0].Passport, "user_1")
		t.Assert(users[0].Password, "")
		t.Assert(users[0].Nickname, "name_1")
		t.Assert(users[0].CreateTime.String(), CreateTime)

		t.Assert(users[9].Id, 10)
		t.Assert(users[9].Passport, "user_10")
		t.Assert(users[9].Password, "")
		t.Assert(users[9].Nickname, "name_10")
		t.Assert(users[9].CreateTime.String(), CreateTime)
	})
}

func Test_Scan_AutoFilteringByStructAttributes(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	type User struct {
		Id       int
		Passport string
	}
	// db.SetDebug(true)
	gtest.C(t, func(t *gtest.T) {
		var user *User
		err := db.X创建Model对象(table).X排序ASC("id").X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 1)
	})
	gtest.C(t, func(t *gtest.T) {
		var users []User
		err := db.X创建Model对象(table).X排序ASC("id").X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), TableSize)
		t.Assert(users[0].Id, 1)
	})
}

func Test_Scan_JsonAttributes(t *testing.T) {
	type GiftImage struct {
		Uid    string `json:"uid"`
		Url    string `json:"url"`
		Status string `json:"status"`
		Name   string `json:"name"`
	}

	type GiftComment struct {
		Name     string `json:"name"`
		Field    string `json:"field"`
		Required bool   `json:"required"`
	}

	type Prop struct {
		Name   string   `json:"name"`
		Values []string `json:"values"`
	}

	type Sku struct {
		GiftId      int64  `json:"gift_id"`
		Name        string `json:"name"`
		ScorePrice  int    `json:"score_price"`
		MarketPrice int    `json:"market_price"`
		CostPrice   int    `json:"cost_price"`
		Stock       int    `json:"stock"`
	}

	type Covers struct {
		List []GiftImage `json:"list"`
	}

	type GiftEntity struct {
		Id                   int64         `json:"id"`
		StoreId              int64         `json:"store_id"`
		GiftType             int           `json:"gift_type"`
		GiftName             string        `json:"gift_name"`
		Description          string        `json:"description"`
		Covers               Covers        `json:"covers"`
		Cover                string        `json:"cover"`
		GiftCategoryId       []int64       `json:"gift_category_id"`
		HasProps             bool          `json:"has_props"`
		OutSn                string        `json:"out_sn"`
		IsLimitSell          bool          `json:"is_limit_sell"`
		LimitSellType        int           `json:"limit_sell_type"`
		LimitSellCycle       string        `json:"limit_sell_cycle"`
		LimitSellCycleCount  int           `json:"limit_sell_cycle_count"`
		LimitSellCustom      bool          `json:"limit_sell_custom"`   // 只允许特定会员兑换
		LimitCustomerTags    []int64       `json:"limit_customer_tags"` // 允许兑换的成员
		ScorePrice           int           `json:"score_price"`
		MarketPrice          float64       `json:"market_price"`
		CostPrice            int           `json:"cost_price"`
		Stock                int           `json:"stock"`
		Props                []Prop        `json:"props"`
		Skus                 []Sku         `json:"skus"`
		ExpressType          []string      `json:"express_type"`
		Comments             []GiftComment `json:"comments"`
		Content              string        `json:"content"`
		AtLeastRechargeCount int           `json:"at_least_recharge_count"`
		Status               int           `json:"status"`
	}

	type User struct {
		Id       int
		Passport string
	}

	table := "jfy_gift"
	array := gstr.X分割并忽略空值(gtest.DataContent(`issue1380.sql`), ";")
	for _, v := range array {
		if _, err := db.X原生SQL执行(ctx, v); err != nil {
			gtest.Error(err)
		}
	}
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		var (
			entity = new(GiftEntity)
			err    = db.X创建Model对象(table).X条件("id", 17).X查询到结构体指针(entity)
		)
		t.AssertNil(err)
		t.Assert(len(entity.Skus), 2)

		t.Assert(entity.Skus[0].Name, "red")
		t.Assert(entity.Skus[0].Stock, 10)
		t.Assert(entity.Skus[0].GiftId, 1)
		t.Assert(entity.Skus[0].CostPrice, 80)
		t.Assert(entity.Skus[0].ScorePrice, 188)
		t.Assert(entity.Skus[0].MarketPrice, 388)

		t.Assert(entity.Skus[1].Name, "blue")
		t.Assert(entity.Skus[1].Stock, 100)
		t.Assert(entity.Skus[1].GiftId, 2)
		t.Assert(entity.Skus[1].CostPrice, 81)
		t.Assert(entity.Skus[1].ScorePrice, 200)
		t.Assert(entity.Skus[1].MarketPrice, 288)

		t.Assert(entity.Id, 17)
		t.Assert(entity.StoreId, 100004)
		t.Assert(entity.GiftType, 1)
		t.Assert(entity.GiftName, "GIFT")
		t.Assert(entity.Description, "支持个性定制的父亲节老师长辈的专属礼物")
		t.Assert(len(entity.Covers.List), 3)
		t.Assert(entity.OutSn, "259402")
		t.Assert(entity.LimitCustomerTags, "[]")
		t.Assert(entity.ScorePrice, 10)
		t.Assert(len(entity.Props), 1)
		t.Assert(len(entity.Comments), 2)
		t.Assert(entity.Status, 99)
		t.Assert(entity.Content, `<p>礼品详情</p>`)
	})
}
