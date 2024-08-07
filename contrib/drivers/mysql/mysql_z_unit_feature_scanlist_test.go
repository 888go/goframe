// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package mysql_test

import (
	"context"
	"fmt"
	"testing"

	gdb "github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gconv "github.com/888go/goframe/util/gconv"
)

func Test_Table_Relation_One(t *testing.T) {
	var (
		tableUser       = "user_" + gtime.X取文本时间戳微秒()
		tableUserDetail = "user_detail_" + gtime.X取文本时间戳微秒()
		tableUserScores = "user_scores_" + gtime.X取文本时间戳微秒()
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		gtest.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		gtest.Error(err)
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
		gtest.Error(err)
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

	// Initialize the data.
	var err error
	gtest.C(t, func(t *gtest.T) {
		err = db.X事务(context.TODO(), func(ctx context.Context, tx gdb.TX) error {
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
	gtest.C(t, func(t *gtest.T) {
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
	gtest.C(t, func(t *gtest.T) {
		var user Entity
				// 从"user"表中选择所有列，其中"name"为'john'. md5:032af229cd8affac
		err := db.X创建Model对象(tableUser).X查询到结构体指针(&user.User, "name", "john")
		t.AssertNil(err)

				// 从"user_detail"表中选择所有列，WHERE子句的条件是`uid`等于1. md5:d5e73807445a5607
		err = db.X创建Model对象(tableUserDetail).X查询到结构体指针(&user.UserDetail, "uid", user.User.Uid)
		t.AssertNil(err)

						// 从`user_scores`表中SELECT * WHERE `uid`=1. md5:d5e5d47d2cdd7d33
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
		tableUser       = "user_" + gtime.X取文本时间戳微秒()
		tableUserDetail = "user_detail_" + gtime.X取文本时间戳微秒()
		tableUserScores = "user_scores_" + gtime.X取文本时间戳微秒()
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		gtest.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		gtest.Error(err)
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
		gtest.Error(err)
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

	// Initialize the data.
	gtest.C(t, func(t *gtest.T) {
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
	gtest.C(t, func(t *gtest.T) {
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
		t.Assert(gconv.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[0])["uid"], 3)
		t.Assert(gconv.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[0])["score"], 1)
		t.Assert(gconv.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[4])["uid"], 3)
		t.Assert(gconv.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[4])["score"], 5)
	})
		// Result 使用具有结构体元素和指针属性的ScanList。 md5:b23d106d13859ad5
	gtest.C(t, func(t *gtest.T) {
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
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("id asc").X查询()
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

		// 使用指针元素和指针属性的ScanList结果。 md5:137ae715e99be611
	gtest.C(t, func(t *gtest.T) {
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
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("id asc").X查询()
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

		// 使用结构体元素和属性扫描Result。 md5:3af2572786856fc5
	gtest.C(t, func(t *gtest.T) {
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
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("id asc").X查询()
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

		// 结果 ScanList 包含指针元素和结构体属性。 md5:57e6be396681268f
	gtest.C(t, func(t *gtest.T) {
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
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("id asc").X查询()
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

		// Model 是一个 ScanList 类型，其中的元素和属性为指针。 md5:815b904cdb4dea16
	gtest.C(t, func(t *gtest.T) {
		var users []*Entity
		// User
		err := db.X创建Model对象(tableUser).
			X条件("uid", g.Slice别名{3, 4}).
			X排序("uid asc").
			X查询到指针列表(&users, "User")
		t.AssertNil(err)
		// Detail
		err = db.X创建Model对象(tableUserDetail).
			X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).
			X排序("uid asc").
			X查询到指针列表(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)
		// Scores
		err = db.X创建Model对象(tableUserScores).
			X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).
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
		tableUser       = "user_" + gtime.X取文本时间戳微秒()
		tableUserDetail = "user_detail_" + gtime.X取文本时间戳微秒()
		tableUserScores = "user_scores_" + gtime.X取文本时间戳微秒()
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		gtest.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		gtest.Error(err)
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
		gtest.Error(err)
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

	// Initialize the data.
	gtest.C(t, func(t *gtest.T) {
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
		// Result 使用具有结构体元素和指针属性的ScanList。 md5:b23d106d13859ad5
	gtest.C(t, func(t *gtest.T) {
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
			X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).
			X排序("uid asc").
			X查询到指针列表(&users, "UserDetail", "User", "uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})

		// Scores
		err = db.X创建Model对象(tableUserScores).
			X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).
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
		tableUser       = "user_" + gtime.X取文本时间戳微秒()
		tableUserDetail = "user_detail_" + gtime.X取文本时间戳微秒()
		tableUserScores = "user_scores_" + gtime.X取文本时间戳微秒()
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		gtest.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		gtest.Error(err)
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
		gtest.Error(err)
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

	// Initialize the data.
	gtest.C(t, func(t *gtest.T) {
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
	gtest.C(t, func(t *gtest.T) {
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
		t.Assert(gconv.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[0])["uid"], 3)
		t.Assert(gconv.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[0])["score"], 1)
		t.Assert(gconv.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[4])["uid"], 3)
		t.Assert(gconv.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[4])["score"], 5)
	})
		// Result 使用具有结构体元素和指针属性的ScanList。 md5:b23d106d13859ad5
	gtest.C(t, func(t *gtest.T) {
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
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid:uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("id asc").X查询()
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

		// 使用指针元素和指针属性的ScanList结果。 md5:137ae715e99be611
	gtest.C(t, func(t *gtest.T) {
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
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "Uid:UID")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("id asc").X查询()
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

		// 使用结构体元素和属性扫描Result。 md5:3af2572786856fc5
	gtest.C(t, func(t *gtest.T) {
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
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid:UId")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("id asc").X查询()
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

		// 结果 ScanList 包含指针元素和结构体属性。 md5:57e6be396681268f
	gtest.C(t, func(t *gtest.T) {
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
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("id asc").X查询()
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

		// Model 是一个 ScanList 类型，其中的元素和属性为指针。 md5:815b904cdb4dea16
	gtest.C(t, func(t *gtest.T) {
		var users []*Entity
		// User
		err := db.X创建Model对象(tableUser).
			X条件("uid", g.Slice别名{3, 4}).
			X排序("uid asc").
			X查询到指针列表(&users, "User")
		t.AssertNil(err)
		// Detail
		err = db.X创建Model对象(tableUserDetail).
			X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).
			X排序("uid asc").
			X查询到指针列表(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)
		// Scores
		err = db.X创建Model对象(tableUserScores).
			X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).
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
		tableUser       = "user_" + gtime.X取文本时间戳微秒()
		tableUserDetail = "user_detail_" + gtime.X取文本时间戳微秒()
		tableUserScores = "user_scores_" + gtime.X取文本时间戳微秒()
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		gtest.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		gtest.Error(err)
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
		gtest.Error(err)
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

	// Initialize the data.
	gtest.C(t, func(t *gtest.T) {
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

		// Result 使用具有结构体元素和指针属性的ScanList。 md5:b23d106d13859ad5
	gtest.C(t, func(t *gtest.T) {
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
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("id asc").X查询()
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

		// 使用指针元素和指针属性的ScanList结果。 md5:137ae715e99be611
	gtest.C(t, func(t *gtest.T) {
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
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "Uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("id asc").X查询()
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

		// 使用结构体元素和属性扫描Result。 md5:3af2572786856fc5
	gtest.C(t, func(t *gtest.T) {
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
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "UId")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("id asc").X查询()
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

		// 结果 ScanList 包含指针元素和结构体属性。 md5:57e6be396681268f
	gtest.C(t, func(t *gtest.T) {
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
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("id asc").X查询()
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

		// Model 是一个 ScanList 类型，其中的元素和属性为指针。 md5:815b904cdb4dea16
	gtest.C(t, func(t *gtest.T) {
		var users []*Entity
		// User
		err := db.X创建Model对象(tableUser).
			X条件("uid", g.Slice别名{3, 4}).
			X排序("uid asc").
			X查询到指针列表(&users, "User")
		t.AssertNil(err)
		// Detail
		err = db.X创建Model对象(tableUserDetail).
			X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).
			X排序("uid asc").
			X查询到指针列表(&users, "UserDetail", "User", "uid")
		t.AssertNil(err)
		// Scores
		err = db.X创建Model对象(tableUserScores).
			X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).
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
		tableUser       = "user_" + gtime.X取文本时间戳微秒()
		tableUserDetail = "user_detail_" + gtime.X取文本时间戳微秒()
		tableUserScores = "user_scores_" + gtime.X取文本时间戳微秒()
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		gtest.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		gtest.Error(err)
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
		gtest.Error(err)
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

		// Result 使用具有结构体元素和指针属性的ScanList。 md5:b23d106d13859ad5
	gtest.C(t, func(t *gtest.T) {
		var users []Entity
		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 0)
		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid:uid")
		t.AssertNil(err)

		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "uid:uid")
		t.AssertNil(err)
	})
	return
		// 使用指针元素和指针属性的ScanList结果。 md5:137ae715e99be611
	gtest.C(t, func(t *gtest.T) {
		var users []*Entity
		// User
		all, err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "User")
		t.AssertNil(err)
		t.Assert(len(users), 0)

		// Detail
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "Uid:UID")
		t.AssertNil(err)

		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "Uid:UID")
		t.AssertNil(err)
	})

		// 使用结构体元素和属性扫描Result。 md5:3af2572786856fc5
	gtest.C(t, func(t *gtest.T) {
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
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid:UId")
		t.AssertNil(err)

		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "UId:Uid")
		t.AssertNil(err)
	})

		// 结果 ScanList 包含指针元素和结构体属性。 md5:57e6be396681268f
	gtest.C(t, func(t *gtest.T) {
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
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)

		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "UID:Uid")
		t.AssertNil(err)
	})

		// Model 是一个 ScanList 类型，其中的元素和属性为指针。 md5:815b904cdb4dea16
	gtest.C(t, func(t *gtest.T) {
		var users []*Entity
		// User
		err := db.X创建Model对象(tableUser).
			X条件("uid", g.Slice别名{3, 4}).
			X排序("uid asc").
			X查询到指针列表(&users, "User")
		t.AssertNil(err)
		// Detail
		err = db.X创建Model对象(tableUserDetail).
			X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).
			X排序("uid asc").
			X查询到指针列表(&users, "UserDetail", "User", "uid:Uid")
		t.AssertNil(err)
		// Scores
		err = db.X创建Model对象(tableUserScores).
			X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).
			X排序("id asc").
			X查询到指针列表(&users, "UserScores", "User", "uid:Uid")
		t.AssertNil(err)

		t.Assert(len(users), 0)
	})
}

func Test_Table_Relation_NoneEqualDataSize(t *testing.T) {
	var (
		tableUser       = "user_" + gtime.X取文本时间戳微秒()
		tableUserDetail = "user_detail_" + gtime.X取文本时间戳微秒()
		tableUserScores = "user_scores_" + gtime.X取文本时间戳微秒()
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		gtest.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		gtest.Error(err)
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
		gtest.Error(err)
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

	// Initialize the data.
	gtest.C(t, func(t *gtest.T) {
		var err error
		for i := 1; i <= 5; i++ {
			// User.
			_, err = db.X插入(ctx, tableUser, g.Map{
				"uid":  i,
				"name": fmt.Sprintf(`name_%d`, i),
			})
			t.AssertNil(err)
			// 细节。
			// 插入用户详细信息到数据库，其中uid为i，address为格式化的`address_i`
			// t.AssertNil(err) 验证错误是否为nil
			// 分数。
			// 循环遍历1到5，插入用户分数到数据库，其中uid为i，score为j
			// t.AssertNil(err) 验证错误是否为nil
			// md5:1b3376e15e2dc36d
		}
	})

		// Result 使用具有结构体元素和指针属性的ScanList。 md5:b23d106d13859ad5
	gtest.C(t, func(t *gtest.T) {
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
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, nil)
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "uid")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 0)
	})

		// 使用指针元素和指针属性的ScanList结果。 md5:137ae715e99be611
	gtest.C(t, func(t *gtest.T) {
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
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "Uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, nil)
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "UID")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 0)
	})

		// 使用结构体元素和属性扫描Result。 md5:3af2572786856fc5
	gtest.C(t, func(t *gtest.T) {
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
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "UId")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, EntityUserDetail{})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "Uid")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 0)
	})

		// 结果 ScanList 包含指针元素和结构体属性。 md5:57e6be396681268f
	gtest.C(t, func(t *gtest.T) {
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
		all, err = db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "User", "uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, EntityUserDetail{})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).X排序("id asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserScores", "User", "UID")
		t.AssertNil(err)
		t.Assert(len(users[0].UserScores), 0)
	})

		// Model 是一个 ScanList 类型，其中的元素和属性为指针。 md5:815b904cdb4dea16
	gtest.C(t, func(t *gtest.T) {
		var users []*Entity
		// User
		err := db.X创建Model对象(tableUser).
			X条件("uid", g.Slice别名{3, 4}).
			X排序("uid asc").
			X查询到指针列表(&users, "User")
		t.AssertNil(err)
		// Detail
		err = db.X创建Model对象(tableUserDetail).
			X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).
			X排序("uid asc").
			X查询到指针列表(&users, "UserDetail", "User", "uid")
		t.AssertNil(err)
		// Scores
		err = db.X创建Model对象(tableUserScores).
			X条件("uid", gdb.X取结构体切片或Map切片值(users, "User", "Uid")).
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
		tableUser       = "user_" + gtime.X取文本时间戳微秒()
		tableUserDetail = "user_detail_" + gtime.X取文本时间戳微秒()
		tableUserScores = "user_scores_" + gtime.X取文本时间戳微秒()
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		gtest.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		gtest.Error(err)
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
		gtest.Error(err)
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

	// Initialize the data.
	gtest.C(t, func(t *gtest.T) {
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

	gtest.C(t, func(t *gtest.T) {
		var (
			err    error
			scores []*EntityUserScores
		)
		// SELECT * FROM `user_scores`
		err = db.X创建Model对象(tableUserScores).X查询到结构体指针(&scores)
		t.AssertNil(err)

				// 从 `user_scores` 表中 SELECT * WHERE `uid` 在 (1,2,3,4,5) 中. md5:b0a4359d4663bf31
		err = db.X创建Model对象(tableUser).
			X条件("uid", gdb.X取结构体切片或Map切片值并去重(&scores, "Uid")).
			X查询到指针列表(&scores, "EntityUser", "uid:Uid")
		t.AssertNil(err)

				// 从'user_detail'表中选择所有列，其中`uid`在(1,2,3,4,5)范围内. md5:fc3208d19b9f10f6
		err = db.X创建Model对象(tableUserDetail).
			X条件("uid", gdb.X取结构体切片或Map切片值并去重(&scores, "Uid")).
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
		tableUser       = "user_" + gtime.X取文本时间戳微秒()
		tableUserDetail = "user_detail_" + gtime.X取文本时间戳微秒()
		tableUserScores = "user_scores_" + gtime.X取文本时间戳微秒()
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUser)); err != nil {
		gtest.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  uid int(10) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(45) NOT NULL,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, tableUserDetail)); err != nil {
		gtest.Error(err)
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
		gtest.Error(err)
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

	// Initialize the data.
	gtest.C(t, func(t *gtest.T) {
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
	gtest.C(t, func(t *gtest.T) {
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
		t.Assert(gconv.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[0])["uid"], 3)
		t.Assert(gconv.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[0])["score"], 1)
		t.Assert(gconv.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[4])["uid"], 3)
		t.Assert(gconv.X取Map(all.X取字段Map泛型类("uid")["3"].Slice别名()[4])["score"], 5)
	})

		// Result 使用具有结构体元素和指针属性的ScanList。 md5:b23d106d13859ad5
	gtest.C(t, func(t *gtest.T) {
		var users []Entity
		// User
		err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].EntityUser, &EntityUser{3, "name_3"})
		t.Assert(users[1].EntityUser, &EntityUser{4, "name_4"})
		// Detail
		all, err := db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "Uid")).X排序("id asc").X查询()
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

		// 使用指针元素和指针属性的ScanList结果。 md5:137ae715e99be611
	gtest.C(t, func(t *gtest.T) {
		var users []*Entity
		// User
		err := db.X创建Model对象(tableUser).X条件("uid", g.Slice别名{3, 4}).X排序("uid asc").X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].EntityUser, &EntityUser{3, "name_3"})
		t.Assert(users[1].EntityUser, &EntityUser{4, "name_4"})
		// Detail
		all, err := db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "Uid")).X排序("id asc").X查询()
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

		// 使用结构体元素和属性扫描Result。 md5:3af2572786856fc5
	gtest.C(t, func(t *gtest.T) {
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
		all, err := db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "Uid")).X排序("id asc").X查询()
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

		// 结果 ScanList 包含指针元素和结构体属性。 md5:57e6be396681268f
	gtest.C(t, func(t *gtest.T) {
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
		all, err := db.X创建Model对象(tableUserDetail).X条件("uid", gdb.X取结构体切片或Map切片值(users, "Uid")).X排序("uid asc").X查询()
		t.AssertNil(err)
		err = all.X取指针列表(&users, "UserDetail", "uid")
		t.AssertNil(err)
		t.Assert(users[0].UserDetail, &EntityUserDetail{3, "address_3"})
		t.Assert(users[1].UserDetail, &EntityUserDetail{4, "address_4"})
		// Scores
		all, err = db.X创建Model对象(tableUserScores).X条件("uid", gdb.X取结构体切片或Map切片值(users, "Uid")).X排序("id asc").X查询()
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

		// Model 是一个 ScanList 类型，其中的元素和属性为指针。 md5:815b904cdb4dea16
	gtest.C(t, func(t *gtest.T) {
		var users []*Entity
		// User
		err := db.X创建Model对象(tableUser).
			X条件("uid", g.Slice别名{3, 4}).
			X排序("uid asc").
			X查询到结构体指针(&users)
		t.AssertNil(err)
		// Detail
		err = db.X创建Model对象(tableUserDetail).
			X条件("uid", gdb.X取结构体切片或Map切片值(users, "Uid")).
			X排序("uid asc").
			X查询到指针列表(&users, "UserDetail", "uid:Uid")
		t.AssertNil(err)
		// Scores
		err = db.X创建Model对象(tableUserScores).
			X条件("uid", gdb.X取结构体切片或Map切片值(users, "Uid")).
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
