// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package mysql_test

import (
	"fmt"
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gmeta"
)

/*
mysql> show tables;
+----------------+
| Tables_in_test |
+----------------+
| user           |
| user_detail    |
| user_score     |
+----------------+
3 rows in set (0.01 sec)

mysql> select * from `user`;
+----+--------+
| id | name   |
+----+--------+
|  1 | name_1 |
|  2 | name_2 |
|  3 | name_3 |
|  4 | name_4 |
|  5 | name_5 |
+----+--------+
5 rows in set (0.01 sec)

mysql> select * from `user_detail`;
+-----+-----------+
| uid | address   |
+-----+-----------+
|   1 | address_1 |
|   2 | address_2 |
|   3 | address_3 |
|   4 | address_4 |
|   5 | address_5 |
+-----+-----------+
5 rows in set (0.00 sec)

mysql> select * from `user_score`;
+----+-----+-------+
| id | uid | score |
+----+-----+-------+
|  1 |   1 |     1 |
|  2 |   1 |     2 |
|  3 |   1 |     3 |
|  4 |   1 |     4 |
|  5 |   1 |     5 |
|  6 |   2 |     1 |
|  7 |   2 |     2 |
|  8 |   2 |     3 |
|  9 |   2 |     4 |
| 10 |   2 |     5 |
| 11 |   3 |     1 |
| 12 |   3 |     2 |
| 13 |   3 |     3 |
| 14 |   3 |     4 |
| 15 |   3 |     5 |
| 16 |   4 |     1 |
| 17 |   4 |     2 |
| 18 |   4 |     3 |
| 19 |   4 |     4 |
| 20 |   4 |     5 |
| 21 |   5 |     1 |
| 22 |   5 |     2 |
| 23 |   5 |     3 |
| 24 |   5 |     4 |
| 25 |   5 |     5 |
+----+-----+-------+
25 rows in set (0.00 sec)
*/

func Test_Table_Relation_With_Scan(t *testing.T) {
	var (
		tableUser       = "user"
		tableUserDetail = "user_detail"
		tableUserScores = "user_score"
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
name varchar(45) NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
uid int(10) unsigned NOT NULL AUTO_INCREMENT,
address varchar(45) NOT NULL,
PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
uid int(10) unsigned NOT NULL,
score int(10) unsigned NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type UserDetail struct {
		元数据类.Meta `orm:"table:user_detail"`
		Uid        int    `json:"uid"`
		Address    string `json:"address"`
	}

	type UserScore struct {
		元数据类.Meta `orm:"table:user_score"`
		Id         int `json:"id"`
		Uid        int `json:"uid"`
		Score      int `json:"score"`
	}

	type User struct {
		元数据类.Meta `orm:"table:user"`
		Id         int          `json:"id"`
		Name       string       `json:"name"`
		UserDetail *UserDetail  `orm:"with:uid=id"`
		UserScores []*UserScore `orm:"with:uid=id"`
	}

	// 初始化数据
	单元测试类.C(t, func(t *单元测试类.T) {
		for i := 1; i <= 5; i++ {
			// User.
			user := User{
				Name: fmt.Sprintf(`name_%d`, i),
			}
			lastInsertId, err := db.X创建Model对象(user).X设置数据(user).X过滤空值().X插入并取ID()
			t.AssertNil(err)
			// Detail.
			userDetail := UserDetail{
				Uid:     int(lastInsertId),
				Address: fmt.Sprintf(`address_%d`, lastInsertId),
			}
			_, err = db.X创建Model对象(userDetail).X设置数据(userDetail).X过滤空值().X插入()
			t.AssertNil(err)
			// Scores.
			for j := 1; j <= 5; j++ {
				userScore := UserScore{
					Uid:   int(lastInsertId),
					Score: j,
				}
				_, err = db.X创建Model对象(userScore).X设置数据(userScore).X过滤空值().X插入()
				t.AssertNil(err)
			}
		}
	})
	for i := 1; i <= 5; i++ {
		// User.
		user := User{
			Name: fmt.Sprintf(`name_%d`, i),
		}
		lastInsertId, err := db.X创建Model对象(user).X设置数据(user).X过滤空值().X插入并取ID()
		单元测试类.AssertNil(err)
		// Detail.
		userDetail := UserDetail{
			Uid:     int(lastInsertId),
			Address: fmt.Sprintf(`address_%d`, lastInsertId),
		}
		_, err = db.X创建Model对象(userDetail).X设置数据(userDetail).X插入()
		单元测试类.AssertNil(err)
		// Scores.
		for j := 1; j <= 5; j++ {
			userScore := UserScore{
				Uid:   int(lastInsertId),
				Score: j,
			}
			_, err = db.X创建Model对象(userScore).X设置数据(userScore).X插入()
			单元测试类.AssertNil(err)
		}
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		var user *User
		err := db.X关联对象(User{}).
			X关联对象(User{}.UserDetail).
			X关联对象(User{}.UserScores).
			X条件("id", 3).
			X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 3)
		t.AssertNE(user.UserDetail, nil)
		t.Assert(user.UserDetail.Uid, 3)
		t.Assert(user.UserDetail.Address, `address_3`)
		t.Assert(len(user.UserScores), 5)
		t.Assert(user.UserScores[0].Uid, 3)
		t.Assert(user.UserScores[0].Score, 1)
		t.Assert(user.UserScores[4].Uid, 3)
		t.Assert(user.UserScores[4].Score, 5)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var user User
		err := db.X关联对象(user).
			X关联对象(user.UserDetail).
			X关联对象(user.UserScores).
			X条件("id", 4).
			X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 4)
		t.AssertNE(user.UserDetail, nil)
		t.Assert(user.UserDetail.Uid, 4)
		t.Assert(user.UserDetail.Address, `address_4`)
		t.Assert(len(user.UserScores), 5)
		t.Assert(user.UserScores[0].Uid, 4)
		t.Assert(user.UserScores[0].Score, 1)
		t.Assert(user.UserScores[4].Uid, 4)
		t.Assert(user.UserScores[4].Score, 5)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var user *User
		err := db.X关联对象(User{}).
			X关联对象(UserDetail{}).
			X关联对象(UserScore{}).
			X条件("id", 4).
			X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 4)
		t.AssertNE(user.UserDetail, nil)
		t.Assert(user.UserDetail.Uid, 4)
		t.Assert(user.UserDetail.Address, `address_4`)
		t.Assert(len(user.UserScores), 5)
		t.Assert(user.UserScores[0].Uid, 4)
		t.Assert(user.UserScores[0].Score, 1)
		t.Assert(user.UserScores[4].Uid, 4)
		t.Assert(user.UserScores[4].Score, 5)
	})
	// 带有部分属性: UserDetail.
	单元测试类.C(t, func(t *单元测试类.T) {
		var user User
		err := db.X关联对象(user).
			X关联对象(user.UserDetail).
			X条件("id", 4).
			X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 4)
		t.AssertNE(user.UserDetail, nil)
		t.Assert(user.UserDetail.Uid, 4)
		t.Assert(user.UserDetail.Address, `address_4`)
		t.Assert(len(user.UserScores), 0)
	})
	// 带有部分属性: UserScores.
	单元测试类.C(t, func(t *单元测试类.T) {
		var user User
		err := db.X关联对象(user).
			X关联对象(user.UserScores).
			X条件("id", 4).
			X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 4)
		t.Assert(user.UserDetail, nil)
		t.Assert(len(user.UserScores), 5)
		t.Assert(user.UserScores[0].Uid, 4)
		t.Assert(user.UserScores[0].Score, 1)
		t.Assert(user.UserScores[4].Uid, 4)
		t.Assert(user.UserScores[4].Score, 5)
	})
}

func Test_Table_Relation_With(t *testing.T) {
	var (
		tableUser       = "user"
		tableUserDetail = "user_detail"
		tableUserScores = "user_scores"
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
name varchar(45) NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
uid int(10) unsigned NOT NULL AUTO_INCREMENT,
address varchar(45) NOT NULL,
PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
uid int(10) unsigned NOT NULL,
score int(10) unsigned NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type UserDetail struct {
		元数据类.Meta `orm:"table:user_detail"`
		Uid        int    `json:"uid"`
		Address    string `json:"address"`
	}

	type UserScores struct {
		元数据类.Meta `orm:"table:user_scores"`
		Id         int `json:"id"`
		Uid        int `json:"uid"`
		Score      int `json:"score"`
	}

	type User struct {
		元数据类.Meta `orm:"table:user"`
		Id         int           `json:"id"`
		Name       string        `json:"name"`
		UserDetail *UserDetail   `orm:"with:uid=id"`
		UserScores []*UserScores `orm:"with:uid=id"`
	}

	// 初始化数据
	var err error
	for i := 1; i <= 5; i++ {
		// User.
		_, err = db.X插入(ctx, tableUser, g.Map{
			"id":   i,
			"name": fmt.Sprintf(`name_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Detail.
		_, err = db.X插入(ctx, tableUserDetail, g.Map{
			"uid":     i,
			"address": fmt.Sprintf(`address_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Scores.
		for j := 1; j <= 5; j++ {
			_, err = db.X插入(ctx, tableUserScores, g.Map{
				"uid":   i,
				"score": j,
			})
			单元测试类.AssertNil(err)
		}
	}

	单元测试类.C(t, func(t *单元测试类.T) {
		var users []*User
		err := db.X关联对象(User{}).
			X关联对象(User{}.UserDetail).
			X关联对象(User{}.UserScores).
			X条件("id", []int{3, 4}).
			X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Id, 3)
		t.Assert(users[0].Name, "name_3")
		t.AssertNE(users[0].UserDetail, nil)
		t.Assert(users[0].UserDetail.Uid, 3)
		t.Assert(users[0].UserDetail.Address, "address_3")
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Uid, 3)
		t.Assert(users[0].UserScores[4].Score, 5)

		t.Assert(users[1].Id, 4)
		t.Assert(users[1].Name, "name_4")
		t.AssertNE(users[1].UserDetail, nil)
		t.Assert(users[1].UserDetail.Uid, 4)
		t.Assert(users[1].UserDetail.Address, "address_4")
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Uid, 4)
		t.Assert(users[1].UserScores[4].Score, 5)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []User
		err := db.X关联对象(User{}).
			X关联对象(User{}.UserDetail).
			X关联对象(User{}.UserScores).
			X条件("id", []int{3, 4}).
			X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Id, 3)
		t.Assert(users[0].Name, "name_3")
		t.AssertNE(users[0].UserDetail, nil)
		t.Assert(users[0].UserDetail.Uid, 3)
		t.Assert(users[0].UserDetail.Address, "address_3")
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Uid, 3)
		t.Assert(users[0].UserScores[4].Score, 5)

		t.Assert(users[1].Id, 4)
		t.Assert(users[1].Name, "name_4")
		t.AssertNE(users[1].UserDetail, nil)
		t.Assert(users[1].UserDetail.Uid, 4)
		t.Assert(users[1].UserDetail.Address, "address_4")
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Uid, 4)
		t.Assert(users[1].UserScores[4].Score, 5)
	})
	// 带有部分属性: UserDetail.
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []*User
		err := db.X关联对象(User{}).
			X关联对象(User{}.UserDetail).
			X条件("id", []int{3, 4}).
			X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Id, 3)
		t.Assert(users[0].Name, "name_3")
		t.AssertNE(users[0].UserDetail, nil)
		t.Assert(users[0].UserDetail.Uid, 3)
		t.Assert(users[0].UserDetail.Address, "address_3")
		t.Assert(len(users[0].UserScores), 0)

		t.Assert(users[1].Id, 4)
		t.Assert(users[1].Name, "name_4")
		t.AssertNE(users[1].UserDetail, nil)
		t.Assert(users[1].UserDetail.Uid, 4)
		t.Assert(users[1].UserDetail.Address, "address_4")
		t.Assert(len(users[1].UserScores), 0)
	})
	// 带有部分属性: UserScores.
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []*User
		err := db.X关联对象(User{}).
			X关联对象(User{}.UserScores).
			X条件("id", []int{3, 4}).
			X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Id, 3)
		t.Assert(users[0].Name, "name_3")
		t.Assert(users[0].UserDetail, nil)
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Uid, 3)
		t.Assert(users[0].UserScores[4].Score, 5)

		t.Assert(users[1].Id, 4)
		t.Assert(users[1].Name, "name_4")
		t.Assert(users[1].UserDetail, nil)
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Uid, 4)
		t.Assert(users[1].UserScores[4].Score, 5)
	})
}

func Test_Table_Relation_WithAll(t *testing.T) {
	var (
		tableUser       = "user"
		tableUserDetail = "user_detail"
		tableUserScores = "user_scores"
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
name varchar(45) NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
uid int(10) unsigned NOT NULL AUTO_INCREMENT,
address varchar(45) NOT NULL,
PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
uid int(10) unsigned NOT NULL,
score int(10) unsigned NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type UserDetail struct {
		元数据类.Meta `orm:"table:user_detail"`
		Uid        int    `json:"uid"`
		Address    string `json:"address"`
	}

	type UserScores struct {
		元数据类.Meta `orm:"table:user_scores"`
		Id         int `json:"id"`
		Uid        int `json:"uid"`
		Score      int `json:"score"`
	}

	type User struct {
		元数据类.Meta `orm:"table:user"`
		Id         int           `json:"id"`
		Name       string        `json:"name"`
		UserDetail *UserDetail   `orm:"with:uid=id"`
		UserScores []*UserScores `orm:"with:uid=id"`
	}

	// 初始化数据
	var err error
	for i := 1; i <= 5; i++ {
		// User.
		_, err = db.X插入(ctx, tableUser, g.Map{
			"id":   i,
			"name": fmt.Sprintf(`name_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Detail.
		_, err = db.X插入(ctx, tableUserDetail, g.Map{
			"uid":     i,
			"address": fmt.Sprintf(`address_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Scores.
		for j := 1; j <= 5; j++ {
			_, err = db.X插入(ctx, tableUserScores, g.Map{
				"uid":   i,
				"score": j,
			})
			单元测试类.AssertNil(err)
		}
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		var user *User
		err := db.X创建Model对象(tableUser).X关联全部对象().X条件("id", 3).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 3)
		t.AssertNE(user.UserDetail, nil)
		t.Assert(user.UserDetail.Uid, 3)
		t.Assert(user.UserDetail.Address, `address_3`)
		t.Assert(len(user.UserScores), 5)
		t.Assert(user.UserScores[0].Uid, 3)
		t.Assert(user.UserScores[0].Score, 1)
		t.Assert(user.UserScores[4].Uid, 3)
		t.Assert(user.UserScores[4].Score, 5)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var user User
		err := db.X创建Model对象(tableUser).X关联全部对象().X条件("id", 4).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 4)
		t.AssertNE(user.UserDetail, nil)
		t.Assert(user.UserDetail.Uid, 4)
		t.Assert(user.UserDetail.Address, `address_4`)
		t.Assert(len(user.UserScores), 5)
		t.Assert(user.UserScores[0].Uid, 4)
		t.Assert(user.UserScores[0].Score, 1)
		t.Assert(user.UserScores[4].Uid, 4)
		t.Assert(user.UserScores[4].Score, 5)
	})
}

func Test_Table_Relation_WithAll_List(t *testing.T) {
	var (
		tableUser       = "user"
		tableUserDetail = "user_detail"
		tableUserScores = "user_scores"
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
name varchar(45) NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
uid int(10) unsigned NOT NULL AUTO_INCREMENT,
address varchar(45) NOT NULL,
PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
uid int(10) unsigned NOT NULL,
score int(10) unsigned NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type UserDetail struct {
		元数据类.Meta `orm:"table:user_detail"`
		Uid        int    `json:"uid"`
		Address    string `json:"address"`
	}

	type UserScores struct {
		元数据类.Meta `orm:"table:user_scores"`
		Id         int `json:"id"`
		Uid        int `json:"uid"`
		Score      int `json:"score"`
	}

	type User struct {
		元数据类.Meta `orm:"table:user"`
		Id         int           `json:"id"`
		Name       string        `json:"name"`
		UserDetail *UserDetail   `orm:"with:uid=id"`
		UserScores []*UserScores `orm:"with:uid=id"`
	}

	// 初始化数据
	var err error
	for i := 1; i <= 5; i++ {
		// User.
		_, err = db.X插入(ctx, tableUser, g.Map{
			"id":   i,
			"name": fmt.Sprintf(`name_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Detail.
		_, err = db.X插入(ctx, tableUserDetail, g.Map{
			"uid":     i,
			"address": fmt.Sprintf(`address_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Scores.
		for j := 1; j <= 5; j++ {
			_, err = db.X插入(ctx, tableUserScores, g.Map{
				"uid":   i,
				"score": j,
			})
			单元测试类.AssertNil(err)
		}
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []*User
		err := db.X创建Model对象(tableUser).X关联全部对象().X条件("id", []int{3, 4}).X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Id, 3)
		t.Assert(users[0].Name, "name_3")
		t.AssertNE(users[0].UserDetail, nil)
		t.Assert(users[0].UserDetail.Uid, 3)
		t.Assert(users[0].UserDetail.Address, "address_3")
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Uid, 3)
		t.Assert(users[0].UserScores[4].Score, 5)

		t.Assert(users[1].Id, 4)
		t.Assert(users[1].Name, "name_4")
		t.AssertNE(users[1].UserDetail, nil)
		t.Assert(users[1].UserDetail.Uid, 4)
		t.Assert(users[1].UserDetail.Address, "address_4")
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Uid, 4)
		t.Assert(users[1].UserScores[4].Score, 5)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []User
		err := db.X创建Model对象(tableUser).X关联全部对象().X条件("id", []int{3, 4}).X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Id, 3)
		t.Assert(users[0].Name, "name_3")
		t.AssertNE(users[0].UserDetail, nil)
		t.Assert(users[0].UserDetail.Uid, 3)
		t.Assert(users[0].UserDetail.Address, "address_3")
		t.Assert(len(users[0].UserScores), 5)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 1)
		t.Assert(users[0].UserScores[4].Uid, 3)
		t.Assert(users[0].UserScores[4].Score, 5)

		t.Assert(users[1].Id, 4)
		t.Assert(users[1].Name, "name_4")
		t.AssertNE(users[1].UserDetail, nil)
		t.Assert(users[1].UserDetail.Uid, 4)
		t.Assert(users[1].UserDetail.Address, "address_4")
		t.Assert(len(users[1].UserScores), 5)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 1)
		t.Assert(users[1].UserScores[4].Uid, 4)
		t.Assert(users[1].UserScores[4].Score, 5)
	})
}

func Test_Table_Relation_WithAllCondition_List(t *testing.T) {
	var (
		tableUser       = "user"
		tableUserDetail = "user_detail"
		tableUserScores = "user_scores"
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
name varchar(45) NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
`, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
uid int(10) unsigned NOT NULL AUTO_INCREMENT,
address varchar(45) NOT NULL,
PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
`, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
uid int(10) unsigned NOT NULL,
score int(10) unsigned NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
`, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type UserDetail struct {
		元数据类.Meta `orm:"table:user_detail"`
		Uid        int    `json:"uid"`
		Address    string `json:"address"`
	}

	type UserScores struct {
		元数据类.Meta `orm:"table:user_scores"`
		Id         int `json:"id"`
		Uid        int `json:"uid"`
		Score      int `json:"score"`
	}

	type User struct {
		元数据类.Meta `orm:"table:user"`
		Id         int           `json:"id"`
		Name       string        `json:"name"`
		UserDetail *UserDetail   `orm:"with:uid=id, where:uid > 3"`
		UserScores []*UserScores `orm:"with:uid=id, where:score>1 and score<5, order:score desc"`
	}

	// 初始化数据
	var err error
	for i := 1; i <= 5; i++ {
		// User.
		_, err = db.X插入(ctx, tableUser, g.Map{
			"id":   i,
			"name": fmt.Sprintf(`name_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Detail.
		_, err = db.X插入(ctx, tableUserDetail, g.Map{
			"uid":     i,
			"address": fmt.Sprintf(`address_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Scores.
		for j := 1; j <= 5; j++ {
			_, err = db.X插入(ctx, tableUserScores, g.Map{
				"uid":   i,
				"score": j,
			})
			单元测试类.AssertNil(err)
		}
	}

	db.X设置调试模式(true)
	defer db.X设置调试模式(false)

	单元测试类.C(t, func(t *单元测试类.T) {
		var users []*User
		err := db.X创建Model对象(tableUser).X关联全部对象().X条件("id", []int{3, 4}).X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Id, 3)
		t.Assert(users[0].Name, "name_3")
		t.Assert(users[0].UserDetail, nil)
		t.Assert(users[1].Id, 4)
		t.Assert(users[1].Name, "name_4")
		t.AssertNE(users[1].UserDetail, nil)
		t.Assert(users[1].UserDetail.Uid, 4)
		t.Assert(users[1].UserDetail.Address, "address_4")
		t.Assert(len(users[1].UserScores), 3)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 4)
		t.Assert(users[1].UserScores[2].Uid, 4)
		t.Assert(users[1].UserScores[2].Score, 2)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []User
		err := db.X创建Model对象(tableUser).X关联全部对象().X条件("id", []int{3, 4}).X查询到结构体指针(&users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Id, 3)
		t.Assert(users[0].Name, "name_3")
		t.Assert(users[0].UserDetail, nil)

		t.Assert(len(users[0].UserScores), 3)
		t.Assert(users[0].UserScores[0].Uid, 3)
		t.Assert(users[0].UserScores[0].Score, 4)
		t.Assert(users[0].UserScores[2].Uid, 3)
		t.Assert(users[0].UserScores[2].Score, 2)

		t.Assert(users[1].Id, 4)
		t.Assert(users[1].Name, "name_4")
		t.AssertNE(users[1].UserDetail, nil)
		t.Assert(users[1].UserDetail.Uid, 4)
		t.Assert(users[1].UserDetail.Address, "address_4")
		t.Assert(len(users[1].UserScores), 3)
		t.Assert(users[1].UserScores[0].Uid, 4)
		t.Assert(users[1].UserScores[0].Score, 4)
		t.Assert(users[1].UserScores[2].Uid, 4)
		t.Assert(users[1].UserScores[2].Score, 2)
	})
}

func Test_Table_Relation_WithAll_Embedded_With_SelfMaintained_Attributes(t *testing.T) {
	var (
		tableUser       = "user"
		tableUserDetail = "user_detail"
		tableUserScores = "user_scores"
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
name varchar(45) NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
uid int(10) unsigned NOT NULL AUTO_INCREMENT,
address varchar(45) NOT NULL,
PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
uid int(10) unsigned NOT NULL,
score int(10) unsigned NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type UserDetail struct {
		元数据类.Meta `orm:"table:user_detail"`
		Uid        int    `json:"uid"`
		Address    string `json:"address"`
	}

	type UserScores struct {
		元数据类.Meta `orm:"table:user_scores"`
		Id         int `json:"id"`
		Uid        int `json:"uid"`
		Score      int `json:"score"`
	}

	type User struct {
		元数据类.Meta  `orm:"table:user"`
		*UserDetail `orm:"with:uid=id"`
		Id          int           `json:"id"`
		Name        string        `json:"name"`
		UserScores  []*UserScores `orm:"with:uid=id"`
	}

	// 初始化数据
	var err error
	for i := 1; i <= 5; i++ {
		// User.
		_, err = db.X插入(ctx, tableUser, g.Map{
			"id":   i,
			"name": fmt.Sprintf(`name_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Detail.
		_, err = db.X插入(ctx, tableUserDetail, g.Map{
			"uid":     i,
			"address": fmt.Sprintf(`address_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Scores.
		for j := 1; j <= 5; j++ {
			_, err = db.X插入(ctx, tableUserScores, g.Map{
				"uid":   i,
				"score": j,
			})
			单元测试类.AssertNil(err)
		}
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		var user *User
		err := db.X创建Model对象(tableUser).X关联全部对象().X条件("id", 3).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 3)
		t.AssertNE(user.UserDetail, nil)
		t.Assert(user.UserDetail.Uid, 3)
		t.Assert(user.UserDetail.Address, `address_3`)
		t.Assert(len(user.UserScores), 5)
		t.Assert(user.UserScores[0].Uid, 3)
		t.Assert(user.UserScores[0].Score, 1)
		t.Assert(user.UserScores[4].Uid, 3)
		t.Assert(user.UserScores[4].Score, 5)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var user User
		err := db.X创建Model对象(tableUser).X关联全部对象().X条件("id", 4).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 4)
		t.AssertNE(user.UserDetail, nil)
		t.Assert(user.UserDetail.Uid, 4)
		t.Assert(user.UserDetail.Address, `address_4`)
		t.Assert(len(user.UserScores), 5)
		t.Assert(user.UserScores[0].Uid, 4)
		t.Assert(user.UserScores[0].Score, 1)
		t.Assert(user.UserScores[4].Uid, 4)
		t.Assert(user.UserScores[4].Score, 5)
	})
}

func Test_Table_Relation_WithAll_Embedded_Without_SelfMaintained_Attributes(t *testing.T) {
	var (
		tableUser       = "user"
		tableUserDetail = "user_detail"
		tableUserScores = "user_scores"
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
name varchar(45) NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
uid int(10) unsigned NOT NULL AUTO_INCREMENT,
address varchar(45) NOT NULL,
PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
uid int(10) unsigned NOT NULL,
score int(10) unsigned NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type UserDetail struct {
		元数据类.Meta `orm:"table:user_detail"`
		Uid        int    `json:"uid"`
		Address    string `json:"address"`
	}

	type UserScores struct {
		元数据类.Meta `orm:"table:user_scores"`
		Id         int `json:"id"`
		Uid        int `json:"uid"`
		Score      int `json:"score"`
	}

	// For Test Only
	type UserEmbedded struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	type User struct {
		元数据类.Meta  `orm:"table:user"`
		*UserDetail `orm:"with:uid=id"`
		UserEmbedded
		UserScores []*UserScores `orm:"with:uid=id"`
	}

	// 初始化数据
	var err error
	for i := 1; i <= 5; i++ {
		// User.
		_, err = db.X插入(ctx, tableUser, g.Map{
			"id":   i,
			"name": fmt.Sprintf(`name_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Detail.
		_, err = db.X插入(ctx, tableUserDetail, g.Map{
			"uid":     i,
			"address": fmt.Sprintf(`address_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Scores.
		for j := 1; j <= 5; j++ {
			_, err = db.X插入(ctx, tableUserScores, g.Map{
				"uid":   i,
				"score": j,
			})
			单元测试类.AssertNil(err)
		}
	}
	db.X设置调试模式(true)
	defer db.X设置调试模式(false)

	单元测试类.C(t, func(t *单元测试类.T) {
		var user *User
		err := db.X创建Model对象(tableUser).X关联全部对象().X条件("id", 3).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 3)
		t.AssertNE(user.UserDetail, nil)
		t.Assert(user.UserDetail.Uid, 3)
		t.Assert(user.UserDetail.Address, `address_3`)
		t.Assert(len(user.UserScores), 5)
		t.Assert(user.UserScores[0].Uid, 3)
		t.Assert(user.UserScores[0].Score, 1)
		t.Assert(user.UserScores[4].Uid, 3)
		t.Assert(user.UserScores[4].Score, 5)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var user User
		err := db.X创建Model对象(tableUser).X关联全部对象().X条件("id", 4).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 4)
		t.AssertNE(user.UserDetail, nil)
		t.Assert(user.UserDetail.Uid, 4)
		t.Assert(user.UserDetail.Address, `address_4`)
		t.Assert(len(user.UserScores), 5)
		t.Assert(user.UserScores[0].Uid, 4)
		t.Assert(user.UserScores[0].Score, 1)
		t.Assert(user.UserScores[4].Uid, 4)
		t.Assert(user.UserScores[4].Score, 5)
	})
}

func Test_Table_Relation_WithAll_Embedded_WithoutMeta(t *testing.T) {
	var (
		tableUser       = "user"
		tableUserDetail = "user_detail"
		tableUserScores = "user_scores"
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
name varchar(45) NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
uid int(10) unsigned NOT NULL AUTO_INCREMENT,
address varchar(45) NOT NULL,
PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
uid int(10) unsigned NOT NULL,
score int(10) unsigned NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type UserDetailBase struct {
		Uid     int    `json:"uid"`
		Address string `json:"address"`
	}

	type UserDetail struct {
		UserDetailBase
	}

	type UserScores struct {
		Id    int `json:"id"`
		Uid   int `json:"uid"`
		Score int `json:"score"`
	}

	type User struct {
		*UserDetail `orm:"with:uid=id"`
		Id          int           `json:"id"`
		Name        string        `json:"name"`
		UserScores  []*UserScores `orm:"with:uid=id"`
	}

	// 初始化数据
	var err error
	for i := 1; i <= 5; i++ {
		// User.
		_, err = db.X插入(ctx, tableUser, g.Map{
			"id":   i,
			"name": fmt.Sprintf(`name_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Detail.
		_, err = db.X插入(ctx, tableUserDetail, g.Map{
			"uid":     i,
			"address": fmt.Sprintf(`address_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Scores.
		for j := 1; j <= 5; j++ {
			_, err = db.X插入(ctx, tableUserScores, g.Map{
				"uid":   i,
				"score": j,
			})
			单元测试类.AssertNil(err)
		}
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		var user *User
		err := db.X创建Model对象(tableUser).X关联全部对象().X条件("id", 3).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 3)
		t.AssertNE(user.UserDetail, nil)
		t.Assert(user.UserDetail.Uid, 3)
		t.Assert(user.UserDetail.Address, `address_3`)
		t.Assert(len(user.UserScores), 5)
		t.Assert(user.UserScores[0].Uid, 3)
		t.Assert(user.UserScores[0].Score, 1)
		t.Assert(user.UserScores[4].Uid, 3)
		t.Assert(user.UserScores[4].Score, 5)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var user User
		err := db.X创建Model对象(tableUser).X关联全部对象().X条件("id", 4).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 4)
		t.AssertNE(user.UserDetail, nil)
		t.Assert(user.UserDetail.Uid, 4)
		t.Assert(user.UserDetail.Address, `address_4`)
		t.Assert(len(user.UserScores), 5)
		t.Assert(user.UserScores[0].Uid, 4)
		t.Assert(user.UserScores[0].Score, 1)
		t.Assert(user.UserScores[4].Uid, 4)
		t.Assert(user.UserScores[4].Score, 5)
	})
}

func Test_Table_Relation_WithAll_AttributeStructAlsoHasWithTag(t *testing.T) {
	var (
		tableUser       = "user"
		tableUserDetail = "user_detail"
		tableUserScores = "user_scores"
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
name varchar(45) NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
`, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
uid int(10) unsigned NOT NULL AUTO_INCREMENT,
address varchar(45) NOT NULL,
PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
`, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
uid int(10) unsigned NOT NULL,
score int(10) unsigned NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
`, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type UserScores struct {
		元数据类.Meta `orm:"table:user_scores"`
		Id         int `json:"id"`
		Uid        int `json:"uid"`
		Score      int `json:"score"`
	}

	type UserDetail struct {
		元数据类.Meta `orm:"table:user_detail"`
		Uid        int           `json:"uid"`
		Address    string        `json:"address"`
		UserScores []*UserScores `orm:"with:uid"`
	}

	type User struct {
		元数据类.Meta  `orm:"table:user"`
		*UserDetail `orm:"with:uid=id"`
		Id          int    `json:"id"`
		Name        string `json:"name"`
	}

	// 初始化数据
	var err error
	for i := 1; i <= 5; i++ {
		// User.
		_, err = db.X插入(ctx, tableUser, g.Map{
			"id":   i,
			"name": fmt.Sprintf(`name_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Detail.
		_, err = db.X插入(ctx, tableUserDetail, g.Map{
			"uid":     i,
			"address": fmt.Sprintf(`address_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Scores.
		for j := 1; j <= 5; j++ {
			_, err = db.X插入(ctx, tableUserScores, g.Map{
				"uid":   i,
				"score": j,
			})
			单元测试类.AssertNil(err)
		}
	}

	单元测试类.C(t, func(t *单元测试类.T) {
		var user *User
		err := db.X创建Model对象(tableUser).X关联全部对象().X条件("id", 3).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 3)
		t.AssertNE(user.UserDetail, nil)
		t.Assert(user.UserDetail.Uid, 3)
		t.Assert(user.UserDetail.Address, `address_3`)
		t.Assert(len(user.UserDetail.UserScores), 5)
		t.Assert(user.UserDetail.UserScores[0].Uid, 3)
		t.Assert(user.UserDetail.UserScores[0].Score, 1)
		t.Assert(user.UserDetail.UserScores[4].Uid, 3)
		t.Assert(user.UserDetail.UserScores[4].Score, 5)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var user User
		err := db.X创建Model对象(tableUser).X关联全部对象().X条件("id", 4).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 4)
		t.AssertNE(user.UserDetail, nil)
		t.Assert(user.UserDetail.Uid, 4)
		t.Assert(user.UserDetail.Address, `address_4`)
		t.Assert(len(user.UserDetail.UserScores), 5)
		t.Assert(user.UserDetail.UserScores[0].Uid, 4)
		t.Assert(user.UserDetail.UserScores[0].Score, 1)
		t.Assert(user.UserDetail.UserScores[4].Uid, 4)
		t.Assert(user.UserDetail.UserScores[4].Score, 5)
	})
}

func Test_Table_Relation_WithAll_AttributeStructAlsoHasWithTag_MoreDeep(t *testing.T) {
	var (
		tableUser       = "user"
		tableUserDetail = "user_detail"
		tableUserScores = "user_scores"
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
name varchar(45) NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
`, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
uid int(10) unsigned NOT NULL AUTO_INCREMENT,
address varchar(45) NOT NULL,
PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
`, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
uid int(10) unsigned NOT NULL,
score int(10) unsigned NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
`, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type UserScores struct {
		元数据类.Meta `orm:"table:user_scores"`
		Id         int `json:"id"`
		Uid        int `json:"uid"`
		Score      int `json:"score"`
	}

	type UserDetail1 struct {
		元数据类.Meta `orm:"table:user_detail"`
		Uid        int           `json:"uid"`
		Address    string        `json:"address"`
		UserScores []*UserScores `orm:"with:uid"`
	}

	type UserDetail2 struct {
		元数据类.Meta  `orm:"table:user_detail"`
		Uid         int           `json:"uid"`
		Address     string        `json:"address"`
		UserDetail1 *UserDetail1  `orm:"with:uid"`
		UserScores  []*UserScores `orm:"with:uid"`
	}

	type UserDetail3 struct {
		元数据类.Meta  `orm:"table:user_detail"`
		Uid         int           `json:"uid"`
		Address     string        `json:"address"`
		UserDetail2 *UserDetail2  `orm:"with:uid"`
		UserScores  []*UserScores `orm:"with:uid"`
	}

	type UserDetail struct {
		元数据类.Meta  `orm:"table:user_detail"`
		Uid         int           `json:"uid"`
		Address     string        `json:"address"`
		UserDetail3 *UserDetail3  `orm:"with:uid"`
		UserScores  []*UserScores `orm:"with:uid"`
	}

	type User struct {
		元数据类.Meta  `orm:"table:user"`
		*UserDetail `orm:"with:uid=id"`
		Id          int    `json:"id"`
		Name        string `json:"name"`
	}

	// 初始化数据
	var err error
	for i := 1; i <= 5; i++ {
		// User.
		_, err = db.X插入(ctx, tableUser, g.Map{
			"id":   i,
			"name": fmt.Sprintf(`name_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Detail.
		_, err = db.X插入(ctx, tableUserDetail, g.Map{
			"uid":     i,
			"address": fmt.Sprintf(`address_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Scores.
		for j := 1; j <= 5; j++ {
			_, err = db.X插入(ctx, tableUserScores, g.Map{
				"uid":   i,
				"score": j,
			})
			单元测试类.AssertNil(err)
		}
	}

	单元测试类.C(t, func(t *单元测试类.T) {
		var user *User
		err := db.X创建Model对象(tableUser).X关联全部对象().X条件("id", 3).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 3)
		t.AssertNE(user.UserDetail, nil)
		t.Assert(user.UserDetail.Uid, 3)
		t.Assert(user.UserDetail.UserDetail3.Uid, 3)
		t.Assert(user.UserDetail.UserDetail3.UserDetail2.Uid, 3)
		t.Assert(user.UserDetail.UserDetail3.UserDetail2.UserDetail1.Uid, 3)
		t.Assert(user.UserDetail.Address, `address_3`)
		t.Assert(len(user.UserDetail.UserScores), 5)
		t.Assert(user.UserDetail.UserScores[0].Uid, 3)
		t.Assert(user.UserDetail.UserScores[0].Score, 1)
		t.Assert(user.UserDetail.UserScores[4].Uid, 3)
		t.Assert(user.UserDetail.UserScores[4].Score, 5)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var user User
		err := db.X创建Model对象(tableUser).X关联全部对象().X条件("id", 4).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 4)
		t.AssertNE(user.UserDetail, nil)
		t.Assert(user.UserDetail.Uid, 4)
		t.Assert(user.UserDetail.UserDetail3.Uid, 4)
		t.Assert(user.UserDetail.UserDetail3.UserDetail2.Uid, 4)
		t.Assert(user.UserDetail.UserDetail3.UserDetail2.UserDetail1.Uid, 4)
		t.Assert(user.UserDetail.Address, `address_4`)
		t.Assert(len(user.UserDetail.UserScores), 5)
		t.Assert(user.UserDetail.UserScores[0].Uid, 4)
		t.Assert(user.UserDetail.UserScores[0].Score, 1)
		t.Assert(user.UserDetail.UserScores[4].Uid, 4)
		t.Assert(user.UserDetail.UserScores[4].Score, 5)
	})
}

func Test_Table_Relation_With_AttributeStructAlsoHasWithTag_MoreDeep(t *testing.T) {
	var (
		tableUser       = "user"
		tableUserDetail = "user_detail"
		tableUserScores = "user_scores"
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
name varchar(45) NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
`, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
uid int(10) unsigned NOT NULL AUTO_INCREMENT,
address varchar(45) NOT NULL,
PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
`, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
uid int(10) unsigned NOT NULL,
score int(10) unsigned NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
`, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type UserScores struct {
		元数据类.Meta `orm:"table:user_scores"`
		Id         int `json:"id"`
		Uid        int `json:"uid"`
		Score      int `json:"score"`
	}

	type UserDetail1 struct {
		元数据类.Meta `orm:"table:user_detail"`
		Uid        int           `json:"uid"`
		Address    string        `json:"address"`
		UserScores []*UserScores `orm:"with:uid"`
	}

	type UserDetail2 struct {
		元数据类.Meta  `orm:"table:user_detail"`
		Uid         int           `json:"uid"`
		Address     string        `json:"address"`
		UserDetail1 *UserDetail1  `orm:"with:uid"`
		UserScores  []*UserScores `orm:"with:uid"`
	}

	type UserDetail3 struct {
		元数据类.Meta  `orm:"table:user_detail"`
		Uid         int           `json:"uid"`
		Address     string        `json:"address"`
		UserDetail2 *UserDetail2  `orm:"with:uid"`
		UserScores  []*UserScores `orm:"with:uid"`
	}

	type UserDetail struct {
		元数据类.Meta  `orm:"table:user_detail"`
		Uid         int           `json:"uid"`
		Address     string        `json:"address"`
		UserDetail3 *UserDetail3  `orm:"with:uid"`
		UserScores  []*UserScores `orm:"with:uid"`
	}

	type User struct {
		元数据类.Meta  `orm:"table:user"`
		*UserDetail `orm:"with:uid=id"`
		Id          int    `json:"id"`
		Name        string `json:"name"`
	}

	// 初始化数据
	var err error
	for i := 1; i <= 5; i++ {
		// User.
		_, err = db.X插入(ctx, tableUser, g.Map{
			"id":   i,
			"name": fmt.Sprintf(`name_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Detail.
		_, err = db.X插入(ctx, tableUserDetail, g.Map{
			"uid":     i,
			"address": fmt.Sprintf(`address_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Scores.
		for j := 1; j <= 5; j++ {
			_, err = db.X插入(ctx, tableUserScores, g.Map{
				"uid":   i,
				"score": j,
			})
			单元测试类.AssertNil(err)
		}
	}

	单元测试类.C(t, func(t *单元测试类.T) {
		var user *User
		err := db.X创建Model对象(tableUser).X关联对象(UserDetail{}, UserDetail2{}, UserDetail3{}, UserScores{}).X条件("id", 3).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 3)
		t.AssertNE(user.UserDetail, nil)
		t.Assert(user.UserDetail.Uid, 3)
		t.Assert(user.UserDetail.UserDetail3.Uid, 3)
		t.Assert(user.UserDetail.UserDetail3.UserDetail2.Uid, 3)
		t.Assert(user.UserDetail.UserDetail3.UserDetail2.UserDetail1, nil)
		t.Assert(user.UserDetail.Address, `address_3`)
		t.Assert(len(user.UserDetail.UserScores), 5)
		t.Assert(user.UserDetail.UserScores[0].Uid, 3)
		t.Assert(user.UserDetail.UserScores[0].Score, 1)
		t.Assert(user.UserDetail.UserScores[4].Uid, 3)
		t.Assert(user.UserDetail.UserScores[4].Score, 5)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var user User
		err := db.X创建Model对象(tableUser).X关联对象(UserDetail{}, UserDetail2{}, UserDetail3{}, UserScores{}).X条件("id", 4).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.Id, 4)
		t.AssertNE(user.UserDetail, nil)
		t.Assert(user.UserDetail.Uid, 4)
		t.Assert(user.UserDetail.UserDetail3.Uid, 4)
		t.Assert(user.UserDetail.UserDetail3.UserDetail2.Uid, 4)
		t.Assert(user.UserDetail.UserDetail3.UserDetail2.UserDetail1, nil)
		t.Assert(user.UserDetail.Address, `address_4`)
		t.Assert(len(user.UserDetail.UserScores), 5)
		t.Assert(user.UserDetail.UserScores[0].Uid, 4)
		t.Assert(user.UserDetail.UserScores[0].Score, 1)
		t.Assert(user.UserDetail.UserScores[4].Uid, 4)
		t.Assert(user.UserDetail.UserScores[4].Score, 5)
	})
}

func Test_Table_Relation_With_MultipleDepends1(t *testing.T) {
	defer func() {
		dropTable("table_a")
		dropTable("table_b")
		dropTable("table_c")
	}()
	for _, v := range 文本类.X分割并忽略空值(文件类.X读文本(单元测试类.DataPath("with_multiple_depends.sql")), ";") {
		if _, err := db.X原生SQL执行(ctx, v); err != nil {
			单元测试类.Error(err)
		}
	}

	type TableC struct {
		元数据类.Meta `orm:"table_c"`
		Id         int `orm:"id,primary" json:"id"`
		TableBId   int `orm:"table_b_id" json:"table_b_id"`
	}

	type TableB struct {
		元数据类.Meta `orm:"table_b"`
		Id         int     `orm:"id,primary" json:"id"`
		TableAId   int     `orm:"table_a_id" json:"table_a_id"`
		TableC     *TableC `orm:"with:table_b_id=id"  json:"table_c"`
	}

	type TableA struct {
		元数据类.Meta `orm:"table_a"`
		Id         int     `orm:"id,primary" json:"id"`
		TableB     *TableB `orm:"with:table_a_id=id" json:"table_b"`
	}

	db.X设置调试模式(true)
	defer db.X设置调试模式(false)

	// Struct.
	单元测试类.C(t, func(t *单元测试类.T) {
		var tableA *TableA
		err := db.X创建Model对象("table_a").X关联全部对象().X查询到结构体指针(&tableA)
		// g.Dump(tableA)
		t.AssertNil(err)
		t.AssertNE(tableA, nil)
		t.Assert(tableA.Id, 1)

		t.AssertNE(tableA.TableB, nil)
		t.AssertNE(tableA.TableB.TableC, nil)
		t.Assert(tableA.TableB.TableAId, 1)
		t.Assert(tableA.TableB.TableC.Id, 100)
		t.Assert(tableA.TableB.TableC.TableBId, 10)
	})

	// Structs
	单元测试类.C(t, func(t *单元测试类.T) {
		var tableA []*TableA
		err := db.X创建Model对象("table_a").X关联全部对象().X排序ASC("id").X查询到结构体指针(&tableA)
		// g.Dump(tableA)
		t.AssertNil(err)
		t.Assert(len(tableA), 2)
		t.AssertNE(tableA[0].TableB, nil)
		t.AssertNE(tableA[1].TableB, nil)
		t.AssertNE(tableA[0].TableB.TableC, nil)
		t.AssertNE(tableA[1].TableB.TableC, nil)

		t.Assert(tableA[0].Id, 1)
		t.Assert(tableA[0].TableB.Id, 10)
		t.Assert(tableA[0].TableB.TableC.Id, 100)

		t.Assert(tableA[1].Id, 2)
		t.Assert(tableA[1].TableB.Id, 20)
		t.Assert(tableA[1].TableB.TableC.Id, 300)
	})
}

func Test_Table_Relation_With_MultipleDepends2(t *testing.T) {
	defer func() {
		dropTable("table_a")
		dropTable("table_b")
		dropTable("table_c")
	}()
	for _, v := range 文本类.X分割并忽略空值(文件类.X读文本(单元测试类.DataPath("with_multiple_depends.sql")), ";") {
		if _, err := db.X原生SQL执行(ctx, v); err != nil {
			单元测试类.Error(err)
		}
	}

	type TableC struct {
		元数据类.Meta `orm:"table_c"`
		Id         int `orm:"id,primary" json:"id"`
		TableBId   int `orm:"table_b_id" json:"table_b_id"`
	}

	type TableB struct {
		元数据类.Meta `orm:"table_b"`
		Id         int       `orm:"id,primary" json:"id"`
		TableAId   int       `orm:"table_a_id" json:"table_a_id"`
		TableC     []*TableC `orm:"with:table_b_id=id"  json:"table_c"`
	}

	type TableA struct {
		元数据类.Meta `orm:"table_a"`
		Id         int       `orm:"id,primary" json:"id"`
		TableB     []*TableB `orm:"with:table_a_id=id" json:"table_b"`
	}

	db.X设置调试模式(true)
	defer db.X设置调试模式(false)

	// Struct.
	单元测试类.C(t, func(t *单元测试类.T) {
		var tableA *TableA
		err := db.X创建Model对象("table_a").X关联全部对象().X查询到结构体指针(&tableA)
		// g.Dump(tableA)
		t.AssertNil(err)
		t.AssertNE(tableA, nil)
		t.Assert(tableA.Id, 1)

		t.Assert(len(tableA.TableB), 2)
		t.Assert(tableA.TableB[0].Id, 10)
		t.Assert(tableA.TableB[1].Id, 30)

		t.Assert(len(tableA.TableB[0].TableC), 2)
		t.Assert(len(tableA.TableB[1].TableC), 1)
		t.Assert(tableA.TableB[0].TableC[0].Id, 100)
		t.Assert(tableA.TableB[0].TableC[0].TableBId, 10)
		t.Assert(tableA.TableB[0].TableC[1].Id, 200)
		t.Assert(tableA.TableB[0].TableC[1].TableBId, 10)
		t.Assert(tableA.TableB[1].TableC[0].Id, 400)
		t.Assert(tableA.TableB[1].TableC[0].TableBId, 30)
	})

	// Structs
	单元测试类.C(t, func(t *单元测试类.T) {
		var tableA []*TableA
		err := db.X创建Model对象("table_a").X关联全部对象().X排序ASC("id").X查询到结构体指针(&tableA)
		// g.Dump(tableA)
		t.AssertNil(err)
		t.Assert(len(tableA), 2)

		t.Assert(len(tableA[0].TableB), 2)
		t.Assert(tableA[0].TableB[0].Id, 10)
		t.Assert(tableA[0].TableB[1].Id, 30)

		t.Assert(len(tableA[0].TableB[0].TableC), 2)
		t.Assert(len(tableA[0].TableB[1].TableC), 1)
		t.Assert(tableA[0].TableB[0].TableC[0].Id, 100)
		t.Assert(tableA[0].TableB[0].TableC[0].TableBId, 10)
		t.Assert(tableA[0].TableB[0].TableC[1].Id, 200)
		t.Assert(tableA[0].TableB[0].TableC[1].TableBId, 10)
		t.Assert(tableA[0].TableB[1].TableC[0].Id, 400)
		t.Assert(tableA[0].TableB[1].TableC[0].TableBId, 30)

		t.Assert(tableA[1].TableB[0].TableC[0].Id, 300)
		t.Assert(tableA[1].TableB[0].TableC[0].TableBId, 20)

		t.Assert(tableA[1].TableB[1].Id, 40)
		t.Assert(tableA[1].TableB[1].TableAId, 2)
		t.Assert(tableA[1].TableB[1].TableC, nil)
	})
}

func Test_Table_Relation_With_MultipleDepends_Embedded(t *testing.T) {
	defer func() {
		dropTable("table_a")
		dropTable("table_b")
		dropTable("table_c")
	}()
	for _, v := range 文本类.X分割并忽略空值(文件类.X读文本(单元测试类.DataPath("with_multiple_depends.sql")), ";") {
		if _, err := db.X原生SQL执行(ctx, v); err != nil {
			单元测试类.Error(err)
		}
	}

	type TableC struct {
		元数据类.Meta `orm:"table_c"`
		Id         int `orm:"id,primary" json:"id"`
		TableBId   int `orm:"table_b_id" json:"table_b_id"`
	}

	type TableB struct {
		元数据类.Meta `orm:"table_b"`
		Id         int `orm:"id,primary" json:"id"`
		TableAId   int `orm:"table_a_id" json:"table_a_id"`
		*TableC    `orm:"with:table_b_id=id"  json:"table_c"`
	}

	type TableA struct {
		元数据类.Meta `orm:"table_a"`
		Id         int `orm:"id,primary" json:"id"`
		*TableB    `orm:"with:table_a_id=id" json:"table_b"`
	}

	db.X设置调试模式(true)
	defer db.X设置调试模式(false)

	// Struct.
	单元测试类.C(t, func(t *单元测试类.T) {
		var tableA *TableA
		err := db.X创建Model对象("table_a").X关联全部对象().X查询到结构体指针(&tableA)
		// g.Dump(tableA)
		t.AssertNil(err)
		t.AssertNE(tableA, nil)
		t.Assert(tableA.Id, 1)

		t.AssertNE(tableA.TableB, nil)
		t.AssertNE(tableA.TableB.TableC, nil)
		t.Assert(tableA.TableB.TableAId, 1)
		t.Assert(tableA.TableB.TableC.Id, 100)
		t.Assert(tableA.TableB.TableC.TableBId, 10)
	})

	// Structs
	单元测试类.C(t, func(t *单元测试类.T) {
		var tableA []*TableA
		err := db.X创建Model对象("table_a").X关联全部对象().X排序ASC("id").X查询到结构体指针(&tableA)
		// g.Dump(tableA)
		t.AssertNil(err)
		t.Assert(len(tableA), 2)
		t.AssertNE(tableA[0].TableB, nil)
		t.AssertNE(tableA[1].TableB, nil)
		t.AssertNE(tableA[0].TableB.TableC, nil)
		t.AssertNE(tableA[1].TableB.TableC, nil)

		t.Assert(tableA[0].Id, 1)
		t.Assert(tableA[0].TableB.Id, 10)
		t.Assert(tableA[0].TableB.TableC.Id, 100)

		t.Assert(tableA[1].Id, 2)
		t.Assert(tableA[1].TableB.Id, 20)
		t.Assert(tableA[1].TableB.TableC.Id, 300)
	})
}

func Test_Table_Relation_WithAll_Embedded_Meta_NameMatchingRule(t *testing.T) {
	var (
		tableUser       = "user100"
		tableUserDetail = "user_detail100"
		tableUserScores = "user_scores100"
	)
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
name varchar(45) NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUser)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUser)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
user_id int(10) unsigned NOT NULL,
address varchar(45) NOT NULL,
PRIMARY KEY (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUserDetail)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserDetail)

	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
id int(10) unsigned NOT NULL AUTO_INCREMENT,
user_id int(10) unsigned NOT NULL,
score int(10) unsigned NOT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 `, tableUserScores)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(tableUserScores)

	type UserDetail struct {
		元数据类.Meta `orm:"table:user_detail100"`
		UserID     int    `json:"user_id"`
		Address    string `json:"address"`
	}

	type UserScores struct {
		元数据类.Meta `orm:"table:user_scores100"`
		ID         int `json:"id"`
		UserID     int `json:"user_id"`
		Score      int `json:"score"`
	}

	// For Test Only
	type UserEmbedded struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	type User struct {
		元数据类.Meta `orm:"table:user100"`
		UserEmbedded
		UserDetail UserDetail    `orm:"with:user_id=id"`
		UserScores []*UserScores `orm:"with:user_id=id"`
	}

	// 初始化数据
	var err error
	for i := 1; i <= 5; i++ {
		// User.
		_, err = db.X插入(ctx, tableUser, g.Map{
			"id":   i,
			"name": fmt.Sprintf(`name_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Detail.
		_, err = db.X插入(ctx, tableUserDetail, g.Map{
			"user_id": i,
			"address": fmt.Sprintf(`address_%d`, i),
		})
		单元测试类.AssertNil(err)
		// Scores.
		for j := 1; j <= 5; j++ {
			_, err = db.X插入(ctx, tableUserScores, g.Map{
				"user_id": i,
				"score":   j,
			})
			单元测试类.AssertNil(err)
		}
	}

// gtest.C(t, func(t *gtest.T) { // 使用gtest框架对代码进行单元测试
//	var user *User // 声明一个指向User类型的指针变量user
//	err := db.Model(tableUser).WithAll().Where("id", 3).Scan(&user) // 根据id为3查询tableUser表中的数据到user变量中
//	t.AssertNil(err) // 断言查询过程中无错误发生，即err应为nil
//	t.Assert(user.ID, 3) // 断言查询结果中user的ID属性为3
//	t.AssertNE(user.UserDetail, nil) // 断言user的UserDetail属性不为空（非nil）
//	t.Assert(user.UserDetail.UserID, 3) // 断言user的UserDetail结构体中的UserID属性为3
//	t.Assert(user.UserDetail.Address, `address_3`) // 断言user的UserDetail结构体中的Address属性为"address_3"
//	t.Assert(len(user.UserScores), 5) // 断言user的UserScores切片长度为5
//	t.Assert(user.UserScores[0].UserID, 3) // 断言user的UserScores切片中第一个元素的UserID属性为3
//	t.Assert(user.UserScores[0].Score, 1) // 断言user的UserScores切片中第一个元素的Score属性为1
//	t.Assert(user.UserScores[4].UserID, 3) // 断言user的UserScores切片中最后一个元素的UserID属性为3
//	t.Assert(user.UserScores[4].Score, 5) // 断言user的UserScores切片中最后一个元素的Score属性为5
// }) // 结束gtest.C()函数的测试用例定义
	单元测试类.C(t, func(t *单元测试类.T) {
		var user User
		err := db.X创建Model对象(tableUser).X关联全部对象().X条件("id", 4).X查询到结构体指针(&user)
		t.AssertNil(err)
		t.Assert(user.ID, 4)
		t.AssertNE(user.UserDetail, nil)
		t.Assert(user.UserDetail.UserID, 4)
		t.Assert(user.UserDetail.Address, `address_4`)
		t.Assert(len(user.UserScores), 5)
		t.Assert(user.UserScores[0].UserID, 4)
		t.Assert(user.UserScores[0].Score, 1)
		t.Assert(user.UserScores[4].UserID, 4)
		t.Assert(user.UserScores[4].Score, 5)
	})
}
