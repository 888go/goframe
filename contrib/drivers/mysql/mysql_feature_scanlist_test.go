// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package mysql_test

import (
	"context"
	"fmt"
	"testing"
	
	"github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_Table_Relation_One(t *testing.T) {
	var (
		tableUser       = "user_" + 时间类.X取文本时间戳微秒()
		tableUserDetail = "user_detail_" + 时间类.X取文本时间戳微秒()
		tableUserScores = "user_scores_" + 时间类.X取文本时间戳微秒()
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  uid int(10) unsigned NOT NULL,
  score int(10) unsigned NOT NULL,
  course varchar(45) NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type EntityUser struct {
		Uid  int    `orm:"uid"`
		Name string `orm:"name"`
	}

	type EntityUserDetail struct {
		Uid     int    `orm:"uid"`
		Address string `orm:"address"`
	}

	type EntityUserScores struct {
		Id     int    `orm:"id"`
		Uid    int    `orm:"uid"`
		Score  int    `orm:"score"`
		Course string `orm:"course"`
	}

	type Entity struct {
		User       *EntityUser
		UserDetail *EntityUserDetail
		UserScores []*EntityUserScores
	}

	// 初始化数据
	var err error
	单元测试类.C(t, func(t *单元测试类.T) {
		err = db.X事务(context.TODO(), func(ctx context.Context, tx db类.TX) error {
			r, err := tx.X创建Model对象(tableUser).X插入并更新已存在(EntityUser{
				Name: "john",
			})
			if err != nil {
				return err
			}
			uid, err := r.LastInsertId()
			if err != nil {
				return err
			}
			_, err = tx.X创建Model对象(tableUserDetail).X插入并更新已存在(EntityUserDetail{
				Uid:     int(uid),
				Address: "Beijing DongZhiMen #66",
			})
			if err != nil {
				return err
			}
			_, err = tx.X创建Model对象(tableUserScores).X插入并更新已存在(g.Slice别名{
				EntityUserScores{Uid: int(uid), Score: 100, Course: "math"},
				EntityUserScores{Uid: int(uid), Score: 99, Course: "physics"},
			})
			return err
		})
		t.AssertNil(err)
	})
	// Data check.
	单元测试类.C(t, func(t *单元测试类.T) {
		r, err := db.X创建Model对象(tableUser).X查询()
		t.AssertNil(err)
		t.Assert(r.X取数量(), 1)
		t.Assert(r[0]["uid"].X取整数(), 1)
		t.Assert(r[0]["name"].String(), "john")

		r, err = db.X创建Model对象(tableUserDetail).X条件("uid", r[0]["uid"].X取整数()).X查询()
		t.AssertNil(err)
		t.Assert(r.X取数量(), 1)
		t.Assert(r[0]["uid"].X取整数(), 1)
		t.Assert(r[0]["address"].String(), `Beijing DongZhiMen #66`)

		r, err = db.X创建Model对象(tableUserScores).X条件("uid", r[0]["uid"].X取整数()).X查询()
		t.AssertNil(err)
		t.Assert(r.X取数量(), 2)
		t.Assert(r[0]["uid"].X取整数(), 1)
		t.Assert(r[1]["uid"].X取整数(), 1)
		t.Assert(r[0]["course"].String(), `math`)
		t.Assert(r[1]["course"].String(), `physics`)
	})
	// Entity query.
	单元测试类.C(t, func(t *单元测试类.T) {
		var user Entity
		// 从`user`表中选择所有列，条件为`name`字段等于'john'
// 即：查询用户表中名为'john'的所有记录
		err := db.X创建Model对象(tableUser).X查询到结构体指针(&user.User, "name", "john")
		t.AssertNil(err)

		// 从`user_detail`表中选取所有列，其条件为`uid`等于1
		err = db.X创建Model对象(tableUserDetail).X查询到结构体指针(&user.UserDetail, "uid", user.User.Uid)
		t.AssertNil(err)

		// 从`user_scores`表中选择所有列，其条件是`uid`等于1
		err = db.X创建Model对象(tableUserScores).X查询到结构体指针(&user.UserScores, "uid", user.User.Uid)
		t.AssertNil(err)

		t.Assert(user.User, EntityUser{
			Uid:  1,
			Name: "john",
		})
		t.Assert(user.UserDetail, EntityUserDetail{
			Uid:     1,
			Address: "Beijing DongZhiMen #66",
		})
		t.Assert(user.UserScores, []EntityUserScores{
			{Id: 1, Uid: 1, Course: "math", Score: 100},
			{Id: 2, Uid: 1, Course: "physics", Score: 99},
		})
	})
}

func Test_Table_Relation_Many(t *testing.T) {
	var (
		tableUser       = "user_" + 时间类.X取文本时间戳微秒()
		tableUserDetail = "user_detail_" + 时间类.X取文本时间戳微秒()
		tableUserScores = "user_scores_" + 时间类.X取文本时间戳微秒()
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  uid int(10) unsigned NOT NULL,
  score int(10) unsigned NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type EntityUser struct {
		Uid  int    `json:"uid"`
		Name string `json:"name"`
	}
	type EntityUserDetail struct {
		Uid     int    `json:"uid"`
		Address string `json:"address"`
	}
	type EntityUserScores struct {
		Id    int `json:"id"`
		Uid   int `json:"uid"`
		Score int `json:"score"`
	}
	type Entity struct {
		User       *EntityUser
		UserDetail *EntityUserDetail
		UserScores []*EntityUserScores
	}

	// 初始化数据
	单元测试类.C(t, func(t *单元测试类.T) {
		var err error
		for i := 1; i <= 5; i++ {
			// User.
			_, err = db.X插入(ctx, tableUser, g.Map{
				"uid":  i,
				"name": fmt.Sprintf(`name_%d`, i),
			})
			t.AssertNil(err)
			// Detail.
			_, err = db.X插入(ctx, tableUserDetail, g.Map{
				"uid":     i,
				"address": fmt.Sprintf(`address_%d`, i),
			})
			t.AssertNil(err)
			// Scores.
			for j := 1; j <= 5; j++ {
				_, err = db.X插入(ctx, tableUserScores, g.Map{
					"uid":   i,
					"score": j,
				})
				t.AssertNil(err)
			}
		}
	})

	// MapKeyValue.
	单元测试类.C(t, func(t *单元测试类.T) {
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		t.Assert(all.X取数量(), 2)
		t.Assert(len(all.X取字段Map泛型类("uid")), 2)
		t.Assert(all.X取字段Map泛型类("uid")["3"].X取Map()["uid"], 3)
		t.Assert(all.X取字段Map泛型类("uid")["4"].X取Map()["uid"], 4)
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", g.Slice别名{3, 4}).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(all.X取数量(), 10)
		t.Assert(len(all.X取字段Map泛型类("uid")), 2)
		t.Assert(len(all.X取字段Map泛型类("uid")["3"].Slice别名()), 5)
		t.Assert(len(all.X取字段Map泛型类("uid")["4"].Slice别名()), 5)
		t.Assert(转换类.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[0])["uid"], 3)
		t.Assert(转换类.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[0])["score"], 1)
		t.Assert(转换类.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[4])["uid"], 3)
		t.Assert(转换类.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[4])["score"], 5)
	})
	// Result ScanList，用于包含结构体元素和指针属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []Entity
		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})
		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "uid:Uid")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})

	// Result ScanList，具有指针元素和指针属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []*Entity
		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})
		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "uid:Uid")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})

	// Result ScanList 用于包含结构体元素及结构体属性的扫描列表。
	单元测试类.C(t, func(t *单元测试类.T) {
		type EntityUser struct {
			Uid  int    `json:"uid"`
			Name string `json:"name"`
		}
		type EntityUserDetail struct {
			Uid     int    `json:"uid"`
			Address string `json:"address"`
		}
		type EntityUserScores struct {
			Id    int `json:"id"`
			Uid   int `json:"uid"`
			Score int `json:"score"`
		}
		type Entity struct {
			User       EntityUser
			UserDetail EntityUserDetail
			UserScores []EntityUserScores
		}
		var users []Entity
		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})
		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "uid:Uid")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})

	// Result 扫描并生成一个具有指针元素和结构体属性的列表。
	单元测试类.C(t, func(t *单元测试类.T) {
		type EntityUser struct {
			Uid  int    `json:"uid"`
			Name string `json:"name"`
		}
		type EntityUserDetail struct {
			Uid     int    `json:"uid"`
			Address string `json:"address"`
		}
		type EntityUserScores struct {
			Id    int `json:"id"`
			Uid   int `json:"uid"`
			Score int `json:"score"`
		}
		type Entity struct {
			User       EntityUser
			UserDetail EntityUserDetail
			UserScores []EntityUserScores
		}
		var users []*Entity

		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})
		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "uid:Uid")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})

	// Model ScanList，其中包含指针元素和指针属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []*Entity
		// User
		err := db.X创建Model对象(tableUser).
			X条件("uid", g.Slice别名{3, 4}).
			X排序("uid asc").
			X查询到指针列表(&users, "User")
		t.AssertNil(err)
		// Detail
		err = db.X创建Model对象(tableUserDetail).
			X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).
			X排序("uid asc").
			X查询到指针列表(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)
		// Scores
		err = db.X创建Model对象(tableUserScores).
			X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).
			X排序("id asc").
			X查询到指针列表(&users, "UserScores", "User", "uid:Uid")
		t.AssertNil(err)

		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})

		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})

		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})
}

func Test_Table_Relation_Many_ModelScanList(t *testing.T) {
	var (
		tableUser       = "user_" + 时间类.X取文本时间戳微秒()
		tableUserDetail = "user_detail_" + 时间类.X取文本时间戳微秒()
		tableUserScores = "user_scores_" + 时间类.X取文本时间戳微秒()
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  uid int(10) unsigned NOT NULL,
  score int(10) unsigned NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type EntityUser struct {
		Uid  int    `json:"uid"`
		Name string `json:"name"`
	}
	type EntityUserDetail struct {
		Uid     int    `json:"uid"`
		Address string `json:"address"`
	}
	type EntityUserScores struct {
		Id    int `json:"id"`
		Uid   int `json:"uid"`
		Score int `json:"score"`
	}
	type Entity struct {
		User       *EntityUser
		UserDetail *EntityUserDetail
		UserScores []*EntityUserScores
	}

	// 初始化数据
	单元测试类.C(t, func(t *单元测试类.T) {
		var err error
		for i := 1; i <= 5; i++ {
			// User.
			_, err = db.X插入(ctx, tableUser, g.Map{
				"uid":  i,
				"name": fmt.Sprintf(`name_%d`, i),
			})
			t.AssertNil(err)
			// Detail.
			_, err = db.X插入(ctx, tableUserDetail, g.Map{
				"uid":     i,
				"address": fmt.Sprintf(`address_%d`, i),
			})
			t.AssertNil(err)
			// Scores.
			for j := 1; j <= 5; j++ {
				_, err = db.X插入(ctx, tableUserScores, g.Map{
					"uid":   i,
					"score": j,
				})
				t.AssertNil(err)
			}
		}
	})

	//db.SetDebug(true)
	// Result ScanList，用于包含结构体元素和指针属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []Entity
		// User
		err := db.X创建Model对象(tableUser).
			X条件("uid", g.Slice别名{3, 4}).
			X排序("uid asc").
			X查询到指针列表(&users, "User")
		t.AssertNil(err)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})

		// Detail
		err = db.X创建Model对象(tableUserDetail).
			X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).
			X排序("uid asc").
			X查询到指针列表(&users, "UserDetail", "User", "uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})

		// Scores
		err = db.X创建Model对象(tableUserScores).
			X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).
			X排序("id asc").
			X查询到指针列表(&users, "UserScores", "User", "uid:Uid")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})
}

func Test_Table_Relation_Many_RelationKeyCaseInsensitive(t *testing.T) {
	var (
		tableUser       = "user_" + 时间类.X取文本时间戳微秒()
		tableUserDetail = "user_detail_" + 时间类.X取文本时间戳微秒()
		tableUserScores = "user_scores_" + 时间类.X取文本时间戳微秒()
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  uid int(10) unsigned NOT NULL,
  score int(10) unsigned NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type EntityUser struct {
		Uid  int    `json:"uid"`
		Name string `json:"name"`
	}
	type EntityUserDetail struct {
		Uid     int    `json:"uid"`
		Address string `json:"address"`
	}
	type EntityUserScores struct {
		Id    int `json:"id"`
		Uid   int `json:"uid"`
		Score int `json:"score"`
	}
	type Entity struct {
		User       *EntityUser
		UserDetail *EntityUserDetail
		UserScores []*EntityUserScores
	}

	// 初始化数据
	单元测试类.C(t, func(t *单元测试类.T) {
		var err error
		for i := 1; i <= 5; i++ {
			// User.
			_, err = db.X插入(ctx, tableUser, g.Map{
				"uid":  i,
				"name": fmt.Sprintf(`name_%d`, i),
			})
			t.AssertNil(err)
			// Detail.
			_, err = db.X插入(ctx, tableUserDetail, g.Map{
				"uid":     i,
				"address": fmt.Sprintf(`address_%d`, i),
			})
			t.AssertNil(err)
			// Scores.
			for j := 1; j <= 5; j++ {
				_, err = db.X插入(ctx, tableUserScores, g.Map{
					"uid":   i,
					"score": j,
				})
				t.AssertNil(err)
			}
		}
	})

	// MapKeyValue.
	单元测试类.C(t, func(t *单元测试类.T) {
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		t.Assert(all.X取数量(), 2)
		t.Assert(len(all.X取字段Map泛型类("uid")), 2)
		t.Assert(all.X取字段Map泛型类("uid")["3"].X取Map()["uid"], 3)
		t.Assert(all.X取字段Map泛型类("uid")["4"].X取Map()["uid"], 4)
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", g.Slice别名{3, 4}).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(all.X取数量(), 10)
		t.Assert(len(all.X取字段Map泛型类("uid")), 2)
		t.Assert(len(all.X取字段Map泛型类("uid")["3"].Slice别名()), 5)
		t.Assert(len(all.X取字段Map泛型类("uid")["4"].Slice别名()), 5)
		t.Assert(转换类.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[0])["uid"], 3)
		t.Assert(转换类.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[0])["score"], 1)
		t.Assert(转换类.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[4])["uid"], 3)
		t.Assert(转换类.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[4])["score"], 5)
	})
	// Result ScanList，用于包含结构体元素和指针属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []Entity
		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})
		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid:uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "uid:uid")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})

	// Result ScanList，具有指针元素和指针属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []*Entity
		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})
		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "Uid:UID")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "Uid:UID")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})

	// Result ScanList 用于包含结构体元素及结构体属性的扫描列表。
	单元测试类.C(t, func(t *单元测试类.T) {
		type EntityUser struct {
			Uid  int    `json:"uid"`
			Name string `json:"name"`
		}
		type EntityUserDetail struct {
			Uid     int    `json:"uid"`
			Address string `json:"address"`
		}
		type EntityUserScores struct {
			Id    int `json:"id"`
			Uid   int `json:"uid"`
			Score int `json:"score"`
		}
		type Entity struct {
			User       EntityUser
			UserDetail EntityUserDetail
			UserScores []EntityUserScores
		}
		var users []Entity
		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})
		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid:UId")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "UId:Uid")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})

	// Result 扫描并生成一个具有指针元素和结构体属性的列表。
	单元测试类.C(t, func(t *单元测试类.T) {
		type EntityUser struct {
			Uid  int    `json:"uid"`
			Name string `json:"name"`
		}
		type EntityUserDetail struct {
			Uid     int    `json:"uid"`
			Address string `json:"address"`
		}
		type EntityUserScores struct {
			Id    int `json:"id"`
			Uid   int `json:"uid"`
			Score int `json:"score"`
		}
		type Entity struct {
			User       EntityUser
			UserDetail EntityUserDetail
			UserScores []EntityUserScores
		}
		var users []*Entity

		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})
		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "UID:Uid")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})

	// Model ScanList，其中包含指针元素和指针属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []*Entity
		// User
		err := db.X创建Model对象(tableUser).
			X条件("uid", g.Slice别名{3, 4}).
			X排序("uid asc").
			X查询到指针列表(&users, "User")
		t.AssertNil(err)
		// Detail
		err = db.X创建Model对象(tableUserDetail).
			X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).
			X排序("uid asc").
			X查询到指针列表(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)
		// Scores
		err = db.X创建Model对象(tableUserScores).
			X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).
			X排序("id asc").
			X查询到指针列表(&users, "UserScores", "User", "uid:Uid")
		t.AssertNil(err)

		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})

		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})

		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})
}

func Test_Table_Relation_Many_TheSameRelationNames(t *testing.T) {
	var (
		tableUser       = "user_" + 时间类.X取文本时间戳微秒()
		tableUserDetail = "user_detail_" + 时间类.X取文本时间戳微秒()
		tableUserScores = "user_scores_" + 时间类.X取文本时间戳微秒()
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  uid int(10) unsigned NOT NULL,
  score int(10) unsigned NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type EntityUser struct {
		Uid  int    `json:"uid"`
		Name string `json:"name"`
	}
	type EntityUserDetail struct {
		Uid     int    `json:"uid"`
		Address string `json:"address"`
	}
	type EntityUserScores struct {
		Id    int `json:"id"`
		Uid   int `json:"uid"`
		Score int `json:"score"`
	}
	type Entity struct {
		User       *EntityUser
		UserDetail *EntityUserDetail
		UserScores []*EntityUserScores
	}

	// 初始化数据
	单元测试类.C(t, func(t *单元测试类.T) {
		var err error
		for i := 1; i <= 5; i++ {
			// User.
			_, err = db.X插入(ctx, tableUser, g.Map{
				"uid":  i,
				"name": fmt.Sprintf(`name_%d`, i),
			})
			t.AssertNil(err)
			// Detail.
			_, err = db.X插入(ctx, tableUserDetail, g.Map{
				"uid":     i,
				"address": fmt.Sprintf(`address_%d`, i),
			})
			t.AssertNil(err)
			// Scores.
			for j := 1; j <= 5; j++ {
				_, err = db.X插入(ctx, tableUserScores, g.Map{
					"uid":   i,
					"score": j,
				})
				t.AssertNil(err)
			}
		}
	})

	// Result ScanList，用于包含结构体元素和指针属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []Entity
		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})
		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "uid")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})

	// Result ScanList，具有指针元素和指针属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []*Entity
		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})
		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "Uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "UID")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})

	// Result ScanList 用于包含结构体元素及结构体属性的扫描列表。
	单元测试类.C(t, func(t *单元测试类.T) {
		type EntityUser struct {
			Uid  int    `json:"uid"`
			Name string `json:"name"`
		}
		type EntityUserDetail struct {
			Uid     int    `json:"uid"`
			Address string `json:"address"`
		}
		type EntityUserScores struct {
			Id    int `json:"id"`
			Uid   int `json:"uid"`
			Score int `json:"score"`
		}
		type Entity struct {
			User       EntityUser
			UserDetail EntityUserDetail
			UserScores []EntityUserScores
		}
		var users []Entity
		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})
		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "UId")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "Uid")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})

	// Result 扫描并生成一个具有指针元素和结构体属性的列表。
	单元测试类.C(t, func(t *单元测试类.T) {
		type EntityUser struct {
			Uid  int    `json:"uid"`
			Name string `json:"name"`
		}
		type EntityUserDetail struct {
			Uid     int    `json:"uid"`
			Address string `json:"address"`
		}
		type EntityUserScores struct {
			Id    int `json:"id"`
			Uid   int `json:"uid"`
			Score int `json:"score"`
		}
		type Entity struct {
			User       EntityUser
			UserDetail EntityUserDetail
			UserScores []EntityUserScores
		}
		var users []*Entity

		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})
		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "UID")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})

	// Model ScanList，其中包含指针元素和指针属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []*Entity
		// User
		err := db.X创建Model对象(tableUser).
			X条件("uid", g.Slice别名{3, 4}).
			X排序("uid asc").
			X查询到指针列表(&users, "User")
		t.AssertNil(err)
		// Detail
		err = db.X创建Model对象(tableUserDetail).
			X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).
			X排序("uid asc").
			X查询到指针列表(&users, "UserDetail", "User", "uid")
		t.AssertNil(err)
		// Scores
		err = db.X创建Model对象(tableUserScores).
			X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).
			X排序("id asc").
			X查询到指针列表(&users, "UserScores", "User", "uid")
		t.AssertNil(err)

		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})

		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})

		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})
}

func Test_Table_Relation_EmptyData(t *testing.T) {
	var (
		tableUser       = "user_" + 时间类.X取文本时间戳微秒()
		tableUserDetail = "user_detail_" + 时间类.X取文本时间戳微秒()
		tableUserScores = "user_scores_" + 时间类.X取文本时间戳微秒()
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  uid int(10) unsigned NOT NULL,
  score int(10) unsigned NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type EntityUser struct {
		Uid  int    `json:"uid"`
		Name string `json:"name"`
	}
	type EntityUserDetail struct {
		Uid     int    `json:"uid"`
		Address string `json:"address"`
	}
	type EntityUserScores struct {
		Id    int `json:"id"`
		Uid   int `json:"uid"`
		Score int `json:"score"`
	}
	type Entity struct {
		User       *EntityUser
		UserDetail *EntityUserDetail
		UserScores []*EntityUserScores
	}

	// Result ScanList，用于包含结构体元素和指针属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []Entity
		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 0)
		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid:uid")
		t.AssertNil(err)

		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "uid:uid")
		t.AssertNil(err)
	})
	return
	// Result ScanList，具有指针元素和指针属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []*Entity
		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 0)

		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "Uid:UID")
		t.AssertNil(err)

		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "Uid:UID")
		t.AssertNil(err)
	})

	// Result ScanList 用于包含结构体元素及结构体属性的扫描列表。
	单元测试类.C(t, func(t *单元测试类.T) {
		type EntityUser struct {
			Uid  int    `json:"uid"`
			Name string `json:"name"`
		}
		type EntityUserDetail struct {
			Uid     int    `json:"uid"`
			Address string `json:"address"`
		}
		type EntityUserScores struct {
			Id    int `json:"id"`
			Uid   int `json:"uid"`
			Score int `json:"score"`
		}
		type Entity struct {
			User       EntityUser
			UserDetail EntityUserDetail
			UserScores []EntityUserScores
		}
		var users []Entity
		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)

		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid:UId")
		t.AssertNil(err)

		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "UId:Uid")
		t.AssertNil(err)
	})

	// Result 扫描并生成一个具有指针元素和结构体属性的列表。
	单元测试类.C(t, func(t *单元测试类.T) {
		type EntityUser struct {
			Uid  int    `json:"uid"`
			Name string `json:"name"`
		}
		type EntityUserDetail struct {
			Uid     int    `json:"uid"`
			Address string `json:"address"`
		}
		type EntityUserScores struct {
			Id    int `json:"id"`
			Uid   int `json:"uid"`
			Score int `json:"score"`
		}
		type Entity struct {
			User       EntityUser
			UserDetail EntityUserDetail
			UserScores []EntityUserScores
		}
		var users []*Entity

		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 0)
		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)

		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "UID:Uid")
		t.AssertNil(err)
	})

	// Model ScanList，其中包含指针元素和指针属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []*Entity
		// User
		err := db.X创建Model对象(tableUser).
			X条件("uid", g.Slice别名{3, 4}).
			X排序("uid asc").
			X查询到指针列表(&users, "User")
		t.AssertNil(err)
		// Detail
		err = db.X创建Model对象(tableUserDetail).
			X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).
			X排序("uid asc").
			X查询到指针列表(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)
		// Scores
		err = db.X创建Model对象(tableUserScores).
			X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).
			X排序("id asc").
			X查询到指针列表(&users, "UserScores", "User", "uid:Uid")
		t.AssertNil(err)

		t.Assert(len(users), 0)
	})
}

func Test_Table_Relation_NoneEqualDataSize(t *testing.T) {
	var (
		tableUser       = "user_" + 时间类.X取文本时间戳微秒()
		tableUserDetail = "user_detail_" + 时间类.X取文本时间戳微秒()
		tableUserScores = "user_scores_" + 时间类.X取文本时间戳微秒()
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  uid int(10) unsigned NOT NULL,
  score int(10) unsigned NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type EntityUser struct {
		Uid  int    `json:"uid"`
		Name string `json:"name"`
	}
	type EntityUserDetail struct {
		Uid     int    `json:"uid"`
		Address string `json:"address"`
	}
	type EntityUserScores struct {
		Id    int `json:"id"`
		Uid   int `json:"uid"`
		Score int `json:"score"`
	}
	type Entity struct {
		User       *EntityUser
		UserDetail *EntityUserDetail
		UserScores []*EntityUserScores
	}

	// 初始化数据
	单元测试类.C(t, func(t *单元测试类.T) {
		var err error
		for i := 1; i <= 5; i++ {
			// User.
			_, err = db.X插入(ctx, tableUser, g.Map{
				"uid":  i,
				"name": fmt.Sprintf(`name_%d`, i),
			})
			t.AssertNil(err)
// 用户详情.
// _, err = db.Insert(ctx, tableUserDetail, g.Map{
//	"uid":     i, // 用户ID
//	"address": fmt.Sprintf(`address_%d`, i), // 格式化后的用户地址信息，索引为i
// })
// t.AssertNil(err) // 断言错误是否为nil，即检查插入用户详情操作是否成功
// 用户得分.
// for j := 1; j <= 5; j++ { // 遍历1到5的得分值
//	_, err = db.Insert(ctx, tableUserScores, g.Map{
//		"uid":   i, // 用户ID
//		"score": j, // 用户对应得分值
//	})
//	t.AssertNil(err) // 断言错误是否为nil，即检查插入用户得分记录操作是否成功
// }
		}
	})

	// Result ScanList，用于包含结构体元素和指针属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []Entity
		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})
		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, nil)
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "uid")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 0)
	})

	// Result ScanList，具有指针元素和指针属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []*Entity
		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})
		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "Uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, nil)
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "UID")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 0)
	})

	// Result ScanList 用于包含结构体元素及结构体属性的扫描列表。
	单元测试类.C(t, func(t *单元测试类.T) {
		type EntityUser struct {
			Uid  int    `json:"uid"`
			Name string `json:"name"`
		}
		type EntityUserDetail struct {
			Uid     int    `json:"uid"`
			Address string `json:"address"`
		}
		type EntityUserScores struct {
			Id    int `json:"id"`
			Uid   int `json:"uid"`
			Score int `json:"score"`
		}
		type Entity struct {
			User       EntityUser
			UserDetail EntityUserDetail
			UserScores []EntityUserScores
		}
		var users []Entity
		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})
		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "UId")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, EntityUserDetail{})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "Uid")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 0)
	})

	// Result 扫描并生成一个具有指针元素和结构体属性的列表。
	单元测试类.C(t, func(t *单元测试类.T) {
		type EntityUser struct {
			Uid  int    `json:"uid"`
			Name string `json:"name"`
		}
		type EntityUserDetail struct {
			Uid     int    `json:"uid"`
			Address string `json:"address"`
		}
		type EntityUserScores struct {
			Id    int `json:"id"`
			Uid   int `json:"uid"`
			Score int `json:"score"`
		}
		type Entity struct {
			User       EntityUser
			UserDetail EntityUserDetail
			UserScores []EntityUserScores
		}
		var users []*Entity

		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})
		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, EntityUserDetail{})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "UID")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 0)
	})

	// Model ScanList，其中包含指针元素和指针属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []*Entity
		// User
		err := db.X创建Model对象(tableUser).
			X条件("uid", g.Slice别名{3, 4}).
			X排序("uid asc").
			X查询到指针列表(&users, "User")
		t.AssertNil(err)
		// Detail
		err = db.X创建Model对象(tableUserDetail).
			X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).
			X排序("uid asc").
			X查询到指针列表(&users, "UserDetail", "User", "uid")
		t.AssertNil(err)
		// Scores
		err = db.X创建Model对象(tableUserScores).
			X条件("uid", db类.X取结构体数组或Map数组值(users, "User", "Uid")).
			X排序("id asc").
			X查询到指针列表(&users, "UserScores", "User", "uid")
		t.AssertNil(err)

		t.Assert(len(users), 2)
		t.Assert(users[0].User, &EntityUser{3, "name_3"})
		t.Assert(users[1].User, &EntityUser{4, "name_4"})

		t.Assert(users[0].UserDetail, nil)

		t.Assert(len(users[0].UserScores), 0)
	})
}

func Test_Table_Relation_EmbeddedStruct1(t *testing.T) {
	var (
		tableUser       = "user_" + 时间类.X取文本时间戳微秒()
		tableUserDetail = "user_detail_" + 时间类.X取文本时间戳微秒()
		tableUserScores = "user_scores_" + 时间类.X取文本时间戳微秒()
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  uid int(10) unsigned NOT NULL,
  score int(10) unsigned NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type EntityUser struct {
		Uid  int    `json:"uid"`
		Name string `json:"name"`
	}
	type EntityUserDetail struct {
		*EntityUser
		Uid     int    `json:"uid"`
		Address string `json:"address"`
	}
	type EntityUserScores struct {
		*EntityUser
		*EntityUserDetail
		Id    int `json:"id"`
		Uid   int `json:"uid"`
		Score int `json:"score"`
	}

	// 初始化数据
	单元测试类.C(t, func(t *单元测试类.T) {
		var err error
		for i := 1; i <= 5; i++ {
			// User.
			_, err = db.X插入(ctx, tableUser, g.Map{
				"uid":  i,
				"name": fmt.Sprintf(`name_%d`, i),
			})
			t.AssertNil(err)
			// Detail.
			_, err = db.X插入(ctx, tableUserDetail, g.Map{
				"uid":     i,
				"address": fmt.Sprintf(`address_%d`, i),
			})
			t.AssertNil(err)
			// Scores.
			for j := 1; j <= 5; j++ {
				_, err = db.X插入(ctx, tableUserScores, g.Map{
					"uid":   i,
					"score": j,
				})
				t.AssertNil(err)
			}
		}
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			err    error
			scores []*EntityUserScores
		)
		// SELECT * FROM `user_scores`
		err = db.X创建Model对象(tableUserScores).X查询到结构体指针(&scores)
		t.AssertNil(err)

		// 从`user_scores`表中选取所有列，其条件是`uid`字段的值存在于列表(1,2,3,4,5)中
// 即：查询`user_scores`表中uid为1、2、3、4或5的所有记录
		err = db.X创建Model对象(tableUser).
			X条件("uid", db类.X取结构体数组或Map数组值并去重(&scores, "Uid")).
			X查询到指针列表(&scores, "EntityUser", "uid:Uid")
		t.AssertNil(err)

		// 从`user_detail`表中选择所有列，其条件是`uid`在(1,2,3,4,5)这个列表内
		err = db.X创建Model对象(tableUserDetail).
			X条件("uid", db类.X取结构体数组或Map数组值并去重(&scores, "Uid")).
			X查询到指针列表(&scores, "EntityUserDetail", "uid:Uid")
		t.AssertNil(err)

		// Assertions.
		t.Assert(len(scores), 25)
		t.Assert(scores[0].Id, 1)
		t.Assert(scores[0].Uid, 1)
		t.Assert(scores[0].Name, "name_1")
		t.Assert(scores[0].Address, "address_1")
		t.Assert(scores[24].Id, 25)
		t.Assert(scores[24].Uid, 5)
		t.Assert(scores[24].Name, "name_5")
		t.Assert(scores[24].Address, "address_5")
	})
}

func Test_Table_Relation_EmbeddedStruct2(t *testing.T) {
	var (
		tableUser       = "user_" + 时间类.X取文本时间戳微秒()
		tableUserDetail = "user_detail_" + 时间类.X取文本时间戳微秒()
		tableUserScores = "user_scores_" + 时间类.X取文本时间戳微秒()
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  uid int(10) unsigned NOT NULL,
  score int(10) unsigned NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type EntityUser struct {
		Uid  int    `json:"uid"`
		Name string `json:"name"`
	}
	type EntityUserDetail struct {
		Uid     int    `json:"uid"`
		Address string `json:"address"`
	}
	type EntityUserScores struct {
		Id    int `json:"id"`
		Uid   int `json:"uid"`
		Score int `json:"score"`
	}
	type Entity struct {
		*EntityUser
		UserDetail *EntityUserDetail
		UserScores []*EntityUserScores
	}

	// 初始化数据
	单元测试类.C(t, func(t *单元测试类.T) {
		var err error
		for i := 1; i <= 5; i++ {
			// User.
			_, err = db.X插入(ctx, tableUser, g.Map{
				"uid":  i,
				"name": fmt.Sprintf(`name_%d`, i),
			})
			t.AssertNil(err)
			// Detail.
			_, err = db.X插入(ctx, tableUserDetail, g.Map{
				"uid":     i,
				"address": fmt.Sprintf(`address_%d`, i),
			})
			t.AssertNil(err)
			// Scores.
			for j := 1; j <= 5; j++ {
				_, err = db.X插入(ctx, tableUserScores, g.Map{
					"uid":   i,
					"score": j,
				})
				t.AssertNil(err)
			}
		}
	})

	// MapKeyValue.
	单元测试类.C(t, func(t *单元测试类.T) {
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		t.Assert(all.X取数量(), 2)
		t.Assert(len(all.X取字段Map泛型类("uid")), 2)
		t.Assert(all.X取字段Map泛型类("uid")["3"].X取Map()["uid"], 3)
		t.Assert(all.X取字段Map泛型类("uid")["4"].X取Map()["uid"], 4)
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", g.Slice别名{3, 4}).X排序("id asc").X查询()
		t.AssertNil(err)
		t.Assert(all.X取数量(), 10)
		t.Assert(len(all.X取字段Map泛型类("uid")), 2)
		t.Assert(len(all.X取字段Map泛型类("uid")["3"].Slice别名()), 5)
		t.Assert(len(all.X取字段Map泛型类("uid")["4"].Slice别名()), 5)
		t.Assert(转换类.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[0])["uid"], 3)
		t.Assert(转换类.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[0])["score"], 1)
		t.Assert(转换类.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[4])["uid"], 3)
		t.Assert(转换类.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[4])["score"], 5)
	})

	// Result ScanList，用于包含结构体元素和指针属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []Entity
		// User
		err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].EntityUser, &EntityUser{3, "name_3"})
		t.Assert(users[1].EntityUser, &EntityUser{4, "name_4"})
		// Detail
		all, err := db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "uid")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})

	// Result ScanList，具有指针元素和指针属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []*Entity
		// User
		err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].EntityUser, &EntityUser{3, "name_3"})
		t.Assert(users[1].EntityUser, &EntityUser{4, "name_4"})
		// Detail
		all, err := db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "uid")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})

	// Result ScanList 用于包含结构体元素及结构体属性的扫描列表。
	单元测试类.C(t, func(t *单元测试类.T) {
		type EntityUser struct {
			Uid  int    `json:"uid"`
			Name string `json:"name"`
		}
		type EntityUserDetail struct {
			Uid     int    `json:"uid"`
			Address string `json:"address"`
		}
		type EntityUserScores struct {
			Id    int `json:"id"`
			Uid   int `json:"uid"`
			Score int `json:"score"`
		}
		type Entity struct {
			EntityUser
			UserDetail EntityUserDetail
			UserScores []EntityUserScores
		}
		var users []Entity
		// User
		err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].EntityUser, &EntityUser{3, "name_3"})
		t.Assert(users[1].EntityUser, &EntityUser{4, "name_4"})
		// Detail
		all, err := db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "uid")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})

	// Result 扫描并生成一个具有指针元素和结构体属性的列表。
	单元测试类.C(t, func(t *单元测试类.T) {
		type EntityUser struct {
			Uid  int    `json:"uid"`
			Name string `json:"name"`
		}
		type EntityUserDetail struct {
			Uid     int    `json:"uid"`
			Address string `json:"address"`
		}
		type EntityUserScores struct {
			Id    int `json:"id"`
			Uid   int `json:"uid"`
			Score int `json:"score"`
		}
		type Entity struct {
			EntityUser
			UserDetail EntityUserDetail
			UserScores []EntityUserScores
		}
		var users []*Entity

		// User
		err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].EntityUser, &EntityUser{3, "name_3"})
		t.Assert(users[1].EntityUser, &EntityUser{4, "name_4"})
		// Detail
		all, err := db.X创建Model对象(tableUserDetail).X条件("uid", db类.X取结构体数组或Map数组值(users, "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", db类.X取结构体数组或Map数组值(users, "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "uid")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})

	// Model ScanList，其中包含指针元素和指针属性。
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []*Entity
		// User
		err := db.X创建Model对象(tableUser).
			X条件("uid", g.Slice别名{3, 4}).
			X排序("uid asc").
			X查询到结构体指针(&users)
		t.AssertNil(err)
		// Detail
		err = db.X创建Model对象(tableUserDetail).
			X条件("uid", db类.X取结构体数组或Map数组值(users, "Uid")).
			X排序("uid asc").
			X查询到指针列表(&users, "UserDetail", "uid:Uid")
		t.AssertNil(err)
		// Scores
		err = db.X创建Model对象(tableUserScores).
			X条件("uid", db类.X取结构体数组或Map数组值(users, "Uid")).
			X排序("id asc").
			X查询到指针列表(&users, "UserScores", "uid:Uid")
		t.AssertNil(err)

		t.Assert(len(users), 2)
		t.Assert(users[0].EntityUser, &EntityUser{3, "name_3"})
		t.Assert(users[1].EntityUser, &EntityUser{4, "name_4"})

		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})

		t.Assert(len(users[0].UserScores), 5)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Score, 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Score, 5)
	})
}
