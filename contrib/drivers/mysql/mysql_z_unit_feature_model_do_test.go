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
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

func Test_Model_Insert_Data_DO(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			g.Meta     `orm:"do:true"`
			Id         interface{}
			Passport   interface{}
			Password   interface{}
			Nickname   interface{}
			CreateTime interface{}
		}
		data := User{
			Id:       1,
			Passport: "user_1",
			Password: "pass_1",
		}
		result, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)
		n, _ := result.LastInsertId()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one[`id`], `1`)
		t.Assert(one[`passport`], `user_1`)
		t.Assert(one[`password`], `pass_1`)
		t.Assert(one[`nickname`], ``)
		t.Assert(one[`create_time`], ``)
	})
}

func Test_Model_Insert_Data_List_DO(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			g.Meta     `orm:"do:true"`
			Id         interface{}
			Passport   interface{}
			Password   interface{}
			Nickname   interface{}
			CreateTime interface{}
		}
		data := g.Slice别名{
			User{
				Id:       1,
				Passport: "user_1",
				Password: "pass_1",
			},
			User{
				Id:       2,
				Passport: "user_2",
				Password: "pass_2",
			},
		}
		result, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)
		n, _ := result.LastInsertId()
		t.Assert(n, 2)

		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one[`id`], `1`)
		t.Assert(one[`passport`], `user_1`)
		t.Assert(one[`password`], `pass_1`)
		t.Assert(one[`nickname`], ``)
		t.Assert(one[`create_time`], ``)

		one, err = db.X创建Model对象(table).X条件并识别主键(2).X查询一条()
		t.AssertNil(err)
		t.Assert(one[`id`], `2`)
		t.Assert(one[`passport`], `user_2`)
		t.Assert(one[`password`], `pass_2`)
		t.Assert(one[`nickname`], ``)
		t.Assert(one[`create_time`], ``)
	})
}

func Test_Model_Update_Data_DO(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			g.Meta     `orm:"do:true"`
			Id         interface{}
			Passport   interface{}
			Password   interface{}
			Nickname   interface{}
			CreateTime interface{}
		}
		data := User{
			Id:       1,
			Passport: "user_100",
			Password: "pass_100",
		}
		_, err := db.X创建Model对象(table).X设置数据(data).X条件并识别主键(1).X更新()
		t.AssertNil(err)

		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one[`id`], `1`)
		t.Assert(one[`passport`], `user_100`)
		t.Assert(one[`password`], `pass_100`)
		t.Assert(one[`nickname`], `name_1`)
	})
}

func Test_Model_Update_Pointer_Data_DO(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type NN string
		type Req struct {
			Id       int
			Passport *string
			Password *string
			Nickname *NN
		}
		type UserDo struct {
			g.Meta     `orm:"do:true"`
			Id         interface{}
			Passport   interface{}
			Password   interface{}
			Nickname   interface{}
			CreateTime interface{}
		}
		var (
			nickname = NN("nickname_111")
			req      = Req{
				Password: gconv.X取文本指针("12345678"),
				Nickname: &nickname,
			}
			data = UserDo{
				Passport: req.Passport,
				Password: req.Password,
				Nickname: req.Nickname,
			}
		)

		_, err := db.X创建Model对象(table).X设置数据(data).X条件并识别主键(1).X更新()
		t.AssertNil(err)

		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one[`id`], `1`)
		t.Assert(one[`password`], `12345678`)
		t.Assert(one[`nickname`], `nickname_111`)
	})
}

func Test_Model_Where_DO(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type User struct {
			g.Meta     `orm:"do:true"`
			Id         interface{}
			Passport   interface{}
			Password   interface{}
			Nickname   interface{}
			CreateTime interface{}
		}
		where := User{
			Id:       1,
			Passport: "user_1",
			Password: "pass_1",
		}
		one, err := db.X创建Model对象(table).X条件(where).X查询一条()
		t.AssertNil(err)
		t.Assert(one[`id`], `1`)
		t.Assert(one[`passport`], `user_1`)
		t.Assert(one[`password`], `pass_1`)
		t.Assert(one[`nickname`], `name_1`)
	})
}

func Test_Model_Insert_Data_ForDao(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type UserForDao struct {
			Id         interface{}
			Passport   interface{}
			Password   interface{}
			Nickname   interface{}
			CreateTime interface{}
		}
		data := UserForDao{
			Id:       1,
			Passport: "user_1",
			Password: "pass_1",
		}
		result, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)
		n, _ := result.LastInsertId()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one[`id`], `1`)
		t.Assert(one[`passport`], `user_1`)
		t.Assert(one[`password`], `pass_1`)
		t.Assert(one[`nickname`], ``)
		t.Assert(one[`create_time`], ``)
	})
}

func Test_Model_Insert_Data_List_ForDao(t *testing.T) {
	table := createTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type UserForDao struct {
			Id         interface{}
			Passport   interface{}
			Password   interface{}
			Nickname   interface{}
			CreateTime interface{}
		}
		data := g.Slice别名{
			UserForDao{
				Id:       1,
				Passport: "user_1",
				Password: "pass_1",
			},
			UserForDao{
				Id:       2,
				Passport: "user_2",
				Password: "pass_2",
			},
		}
		result, err := db.X创建Model对象(table).X设置数据(data).X插入()
		t.AssertNil(err)
		n, _ := result.LastInsertId()
		t.Assert(n, 2)

		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one[`id`], `1`)
		t.Assert(one[`passport`], `user_1`)
		t.Assert(one[`password`], `pass_1`)
		t.Assert(one[`nickname`], ``)
		t.Assert(one[`create_time`], ``)

		one, err = db.X创建Model对象(table).X条件并识别主键(2).X查询一条()
		t.AssertNil(err)
		t.Assert(one[`id`], `2`)
		t.Assert(one[`passport`], `user_2`)
		t.Assert(one[`password`], `pass_2`)
		t.Assert(one[`nickname`], ``)
		t.Assert(one[`create_time`], ``)
	})
}

func Test_Model_Update_Data_ForDao(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type UserForDao struct {
			Id         interface{}
			Passport   interface{}
			Password   interface{}
			Nickname   interface{}
			CreateTime interface{}
		}
		data := UserForDao{
			Id:       1,
			Passport: "user_100",
			Password: "pass_100",
		}
		_, err := db.X创建Model对象(table).X设置数据(data).X条件并识别主键(1).X更新()
		t.AssertNil(err)

		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one[`id`], `1`)
		t.Assert(one[`passport`], `user_100`)
		t.Assert(one[`password`], `pass_100`)
		t.Assert(one[`nickname`], `name_1`)
	})
}

func Test_Model_Where_ForDao(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		type UserForDao struct {
			Id         interface{}
			Passport   interface{}
			Password   interface{}
			Nickname   interface{}
			CreateTime interface{}
		}
		where := UserForDao{
			Id:       1,
			Passport: "user_1",
			Password: "pass_1",
		}
		one, err := db.X创建Model对象(table).X条件(where).X查询一条()
		t.AssertNil(err)
		t.Assert(one[`id`], `1`)
		t.Assert(one[`passport`], `user_1`)
		t.Assert(one[`password`], `pass_1`)
		t.Assert(one[`nickname`], `name_1`)
	})
}

func Test_Model_Where_FieldPrefix(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array := gstr.X分割并忽略空值(gtest.DataContent(`table_with_prefix.sql`), ";")
		for _, v := range array {
			if _, err := db.X原生SQL执行(ctx, v); err != nil {
				gtest.Error(err)
			}
		}
		defer dropTable("instance")

		type Instance struct {
			ID   int `orm:"f_id"`
			Name string
		}

		type InstanceDo struct {
			g.Meta `orm:"table:instance, do:true"`
			ID     interface{} `orm:"f_id"`
		}
		var instance *Instance
		err := db.X创建Model对象("instance").X条件(InstanceDo{
			ID: 1,
		}).X查询到结构体指针(&instance)
		t.AssertNil(err)
		t.AssertNE(instance, nil)
		t.Assert(instance.ID, 1)
		t.Assert(instance.Name, "john")
	})
	// With omitempty.
	gtest.C(t, func(t *gtest.T) {
		array := gstr.X分割并忽略空值(gtest.DataContent(`table_with_prefix.sql`), ";")
		for _, v := range array {
			if _, err := db.X原生SQL执行(ctx, v); err != nil {
				gtest.Error(err)
			}
		}
		defer dropTable("instance")

		type Instance struct {
			ID   int `orm:"f_id,omitempty"`
			Name string
		}

		type InstanceDo struct {
			g.Meta `orm:"table:instance, do:true"`
			ID     interface{} `orm:"f_id,omitempty"`
		}
		var instance *Instance
		err := db.X创建Model对象("instance").X条件(InstanceDo{
			ID: 1,
		}).X查询到结构体指针(&instance)
		t.AssertNil(err)
		t.AssertNE(instance, nil)
		t.Assert(instance.ID, 1)
		t.Assert(instance.Name, "john")
	})
}
