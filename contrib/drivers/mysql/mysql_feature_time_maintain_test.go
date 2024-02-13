// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package mysql_test

import (
	"fmt"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
)

// CreateAt/UpdateAt/DeleteAt.
// 创建时间/更新时间/删除时间。
func Test_SoftCreateUpdateDeleteTimeMicroSecond(t *testing.T) {
	table := "time_test_table_" + 时间类.X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id        int(11) NOT NULL,
  name      varchar(45) DEFAULT NULL,
  create_at datetime(6) DEFAULT NULL,
  update_at datetime(6) DEFAULT NULL,
  delete_at datetime(6) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, table)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		// Insert
		dataInsert := g.Map{
			"id":   1,
			"name": "name_1",
		}
		r, err := db.X创建Model对象(table).X设置数据(dataInsert).X插入()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		oneInsert, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneInsert["id"].X取整数(), 1)
		t.Assert(oneInsert["name"].String(), "name_1")
		t.Assert(oneInsert["delete_at"].String(), "")
		t.AssertGE(oneInsert["create_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)
		t.AssertGE(oneInsert["update_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)

		// 用于时间断言的目的。
		time.Sleep(2 * time.Second)

		// Save
		dataSave := g.Map{
			"id":   1,
			"name": "name_10",
		}
		r, err = db.X创建Model对象(table).X设置数据(dataSave).X插入并更新已存在()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 2)

		oneSave, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneSave["id"].X取整数(), 1)
		t.Assert(oneSave["name"].String(), "name_10")
		t.Assert(oneSave["delete_at"].String(), "")
		t.Assert(oneSave["create_at"].X取gtime时间类().X取时间戳秒(), oneInsert["create_at"].X取gtime时间类().X取时间戳秒())
		t.AssertNE(oneSave["update_at"].X取gtime时间类().X取时间戳秒(), oneInsert["update_at"].X取gtime时间类().X取时间戳秒())
		t.AssertGE(oneSave["update_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)

		// 用于时间断言的目的。
		time.Sleep(2 * time.Second)

		// Update
		dataUpdate := g.Map{
			"name": "name_1000",
		}
		r, err = db.X创建Model对象(table).X设置数据(dataUpdate).X条件并识别主键(1).X更新()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)

		oneUpdate, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneUpdate["id"].X取整数(), 1)
		t.Assert(oneUpdate["name"].String(), "name_1000")
		t.Assert(oneUpdate["delete_at"].String(), "")
		t.Assert(oneUpdate["create_at"].X取gtime时间类().X取时间戳秒(), oneInsert["create_at"].X取gtime时间类().X取时间戳秒())
		t.AssertGE(oneUpdate["update_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)

		// Replace
		dataReplace := g.Map{
			"id":   1,
			"name": "name_100",
		}
		r, err = db.X创建Model对象(table).X设置数据(dataReplace).X插入并替换已存在()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 2)

		oneReplace, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneReplace["id"].X取整数(), 1)
		t.Assert(oneReplace["name"].String(), "name_100")
		t.Assert(oneReplace["delete_at"].String(), "")
		t.AssertGE(oneReplace["create_at"].X取gtime时间类().X取时间戳秒(), oneInsert["create_at"].X取gtime时间类().X取时间戳秒())
		t.AssertGE(oneReplace["update_at"].X取gtime时间类().X取时间戳秒(), oneInsert["update_at"].X取gtime时间类().X取时间戳秒())

		// 用于时间断言的目的。
		time.Sleep(2 * time.Second)

		// Delete
		r, err = db.X创建Model对象(table).X删除("id", 1)
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)
		// Delete Select
		one4, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one4), 0)
		one5, err := db.X创建Model对象(table).X禁用时间自动更新特性().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one5["id"].X取整数(), 1)
		t.AssertGE(one5["delete_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)
		// Delete Count
		i, err := db.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(i, 0)
		i, err = db.X创建Model对象(table).X禁用时间自动更新特性().X查询行数()
		t.AssertNil(err)
		t.Assert(i, 1)

		// Delete Unscoped
		r, err = db.X创建Model对象(table).X禁用时间自动更新特性().X删除("id", 1)
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)
		one6, err := db.X创建Model对象(table).X禁用时间自动更新特性().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one6), 0)
		i, err = db.X创建Model对象(table).X禁用时间自动更新特性().X查询行数()
		t.AssertNil(err)
		t.Assert(i, 0)
	})
}

// CreateAt/UpdateAt/DeleteAt.
// 创建时间/更新时间/删除时间。
func Test_SoftCreateUpdateDeleteTimeSecond(t *testing.T) {
	table := "time_test_table_" + 时间类.X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id        int(11) NOT NULL,
  name      varchar(45) DEFAULT NULL,
  create_at datetime(0) DEFAULT NULL,
  update_at datetime(0) DEFAULT NULL,
  delete_at datetime(0) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, table)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		// Insert
		dataInsert := g.Map{
			"id":   1,
			"name": "name_1",
		}
		r, err := db.X创建Model对象(table).X设置数据(dataInsert).X插入()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		oneInsert, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneInsert["id"].X取整数(), 1)
		t.Assert(oneInsert["name"].String(), "name_1")
		t.Assert(oneInsert["delete_at"].String(), "")
		t.AssertGE(oneInsert["create_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)
		t.AssertGE(oneInsert["update_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)

		// 用于时间断言的目的。
		time.Sleep(2 * time.Second)

		// Save
		dataSave := g.Map{
			"id":   1,
			"name": "name_10",
		}
		r, err = db.X创建Model对象(table).X设置数据(dataSave).X插入并更新已存在()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 2)

		oneSave, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneSave["id"].X取整数(), 1)
		t.Assert(oneSave["name"].String(), "name_10")
		t.Assert(oneSave["delete_at"].String(), "")
		t.Assert(oneSave["create_at"].X取gtime时间类().X取时间戳秒(), oneInsert["create_at"].X取gtime时间类().X取时间戳秒())
		t.AssertNE(oneSave["update_at"].X取gtime时间类().X取时间戳秒(), oneInsert["update_at"].X取gtime时间类().X取时间戳秒())
		t.AssertGE(oneSave["update_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)

		// 用于时间断言的目的。
		time.Sleep(2 * time.Second)

		// Update
		dataUpdate := g.Map{
			"name": "name_1000",
		}
		r, err = db.X创建Model对象(table).X设置数据(dataUpdate).X条件并识别主键(1).X更新()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)

		oneUpdate, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneUpdate["id"].X取整数(), 1)
		t.Assert(oneUpdate["name"].String(), "name_1000")
		t.Assert(oneUpdate["delete_at"].String(), "")
		t.Assert(oneUpdate["create_at"].X取gtime时间类().X取时间戳秒(), oneInsert["create_at"].X取gtime时间类().X取时间戳秒())
		t.AssertGE(oneUpdate["update_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)

		// Replace
		dataReplace := g.Map{
			"id":   1,
			"name": "name_100",
		}
		r, err = db.X创建Model对象(table).X设置数据(dataReplace).X插入并替换已存在()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 2)

		oneReplace, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneReplace["id"].X取整数(), 1)
		t.Assert(oneReplace["name"].String(), "name_100")
		t.Assert(oneReplace["delete_at"].String(), "")
		t.AssertGE(oneReplace["create_at"].X取gtime时间类().X取时间戳秒(), oneInsert["create_at"].X取gtime时间类().X取时间戳秒())
		t.AssertGE(oneReplace["update_at"].X取gtime时间类().X取时间戳秒(), oneInsert["update_at"].X取gtime时间类().X取时间戳秒())

		// 用于时间断言的目的。
		time.Sleep(2 * time.Second)

		// Delete
		r, err = db.X创建Model对象(table).X删除("id", 1)
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)
		// Delete Select
		one4, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one4), 0)
		one5, err := db.X创建Model对象(table).X禁用时间自动更新特性().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one5["id"].X取整数(), 1)
		t.AssertGE(one5["delete_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)
		// Delete Count
		i, err := db.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(i, 0)
		i, err = db.X创建Model对象(table).X禁用时间自动更新特性().X查询行数()
		t.AssertNil(err)
		t.Assert(i, 1)

		// Delete Unscoped
		r, err = db.X创建Model对象(table).X禁用时间自动更新特性().X删除("id", 1)
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)
		one6, err := db.X创建Model对象(table).X禁用时间自动更新特性().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one6), 0)
		i, err = db.X创建Model对象(table).X禁用时间自动更新特性().X查询行数()
		t.AssertNil(err)
		t.Assert(i, 0)
	})
}

// CreatedAt/UpdatedAt/DeletedAt.
//这三个字段分别代表：
// CreatedAt：记录创建的时间
// UpdatedAt：记录最后一次更新的时间
// DeletedAt：记录删除的时间（若该记录已被逻辑删除）
func Test_SoftCreatedUpdatedDeletedTime_Map(t *testing.T) {
	table := "time_test_table_" + 时间类.X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id        int(11) NOT NULL,
  name      varchar(45) DEFAULT NULL,
  created_at datetime(6) DEFAULT NULL,
  updated_at datetime(6) DEFAULT NULL,
  deleted_at datetime(6) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, table)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		// Insert
		dataInsert := g.Map{
			"id":   1,
			"name": "name_1",
		}
		r, err := db.X创建Model对象(table).X设置数据(dataInsert).X插入()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		oneInsert, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneInsert["id"].X取整数(), 1)
		t.Assert(oneInsert["name"].String(), "name_1")
		t.Assert(oneInsert["deleted_at"].String(), "")
		t.AssertGE(oneInsert["created_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)
		t.AssertGE(oneInsert["updated_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)

		// 用于时间断言的目的。
		time.Sleep(2 * time.Second)

		// Save
		dataSave := g.Map{
			"id":   1,
			"name": "name_10",
		}
		r, err = db.X创建Model对象(table).X设置数据(dataSave).X插入并更新已存在()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 2)

		oneSave, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneSave["id"].X取整数(), 1)
		t.Assert(oneSave["name"].String(), "name_10")
		t.Assert(oneSave["deleted_at"].String(), "")
		t.Assert(oneSave["created_at"].X取gtime时间类().X取时间戳秒(), oneInsert["created_at"].X取gtime时间类().X取时间戳秒())
		t.AssertNE(oneSave["updated_at"].X取gtime时间类().X取时间戳秒(), oneInsert["updated_at"].X取gtime时间类().X取时间戳秒())
		t.AssertGE(oneSave["updated_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)

		// 用于时间断言的目的。
		time.Sleep(2 * time.Second)

		// Update
		dataUpdate := g.Map{
			"name": "name_1000",
		}
		r, err = db.X创建Model对象(table).X设置数据(dataUpdate).X条件并识别主键(1).X更新()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)

		oneUpdate, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneUpdate["id"].X取整数(), 1)
		t.Assert(oneUpdate["name"].String(), "name_1000")
		t.Assert(oneUpdate["deleted_at"].String(), "")
		t.Assert(oneUpdate["created_at"].X取gtime时间类().X取时间戳秒(), oneInsert["created_at"].X取gtime时间类().X取时间戳秒())
		t.AssertGE(oneUpdate["updated_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)

		// Replace
		dataReplace := g.Map{
			"id":   1,
			"name": "name_100",
		}
		r, err = db.X创建Model对象(table).X设置数据(dataReplace).X插入并替换已存在()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 2)

		oneReplace, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneReplace["id"].X取整数(), 1)
		t.Assert(oneReplace["name"].String(), "name_100")
		t.Assert(oneReplace["deleted_at"].String(), "")
		t.AssertGE(oneReplace["created_at"].X取gtime时间类().X取时间戳秒(), oneInsert["created_at"].X取gtime时间类().X取时间戳秒())
		t.AssertGE(oneReplace["updated_at"].X取gtime时间类().X取时间戳秒(), oneInsert["updated_at"].X取gtime时间类().X取时间戳秒())

		// 用于时间断言的目的。
		time.Sleep(2 * time.Second)

		// Delete
		r, err = db.X创建Model对象(table).X删除("id", 1)
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)
		// Delete Select
		one4, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one4), 0)
		one5, err := db.X创建Model对象(table).X禁用时间自动更新特性().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one5["id"].X取整数(), 1)
		t.AssertGE(one5["deleted_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)
		// Delete Count
		i, err := db.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(i, 0)
		i, err = db.X创建Model对象(table).X禁用时间自动更新特性().X查询行数()
		t.AssertNil(err)
		t.Assert(i, 1)

		// Delete Unscoped
		r, err = db.X创建Model对象(table).X禁用时间自动更新特性().X删除("id", 1)
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)
		one6, err := db.X创建Model对象(table).X禁用时间自动更新特性().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one6), 0)
		i, err = db.X创建Model对象(table).X禁用时间自动更新特性().X查询行数()
		t.AssertNil(err)
		t.Assert(i, 0)
	})
}

// CreatedAt/UpdatedAt/DeletedAt.
//这三个字段分别代表：
// CreatedAt：记录创建的时间
// UpdatedAt：记录最后一次更新的时间
// DeletedAt：记录删除的时间（若该记录已被逻辑删除）
func Test_SoftCreatedUpdatedDeletedTime_Struct(t *testing.T) {
	table := "time_test_table_" + 时间类.X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id        int(11) NOT NULL,
  name      varchar(45) DEFAULT NULL,
  created_at datetime(6) DEFAULT NULL,
  updated_at datetime(6) DEFAULT NULL,
  deleted_at datetime(6) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, table)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(table)

	type User struct {
		Id        int
		Name      string
		CreatedAT *时间类.Time
		UpdatedAT *时间类.Time
		DeletedAT *时间类.Time
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		// Insert
		dataInsert := User{
			Id:   1,
			Name: "name_1",
		}
		r, err := db.X创建Model对象(table).X设置数据(dataInsert).X插入()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		oneInsert, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneInsert["id"].X取整数(), 1)
		t.Assert(oneInsert["name"].String(), "name_1")
		t.Assert(oneInsert["deleted_at"].String(), "")
		t.AssertGE(oneInsert["created_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)
		t.AssertGE(oneInsert["updated_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)

		// 用于时间断言的目的。
		time.Sleep(2 * time.Second)

		// Save
		dataSave := User{
			Id:   1,
			Name: "name_10",
		}
		r, err = db.X创建Model对象(table).X设置数据(dataSave).X过滤空值().X插入并更新已存在()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 2)

		oneSave, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneSave["id"].X取整数(), 1)
		t.Assert(oneSave["name"].String(), "name_10")
		t.Assert(oneSave["deleted_at"].String(), "")
		t.Assert(oneSave["created_at"].X取gtime时间类().X取时间戳秒(), oneInsert["created_at"].X取gtime时间类().X取时间戳秒())
		t.AssertNE(oneSave["updated_at"].X取gtime时间类().X取时间戳秒(), oneInsert["updated_at"].X取gtime时间类().X取时间戳秒())
		t.AssertGE(oneSave["updated_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)

		// 用于时间断言的目的。
		time.Sleep(2 * time.Second)

		// Update
		dataUpdate := User{
			Name: "name_1000",
		}
		r, err = db.X创建Model对象(table).X设置数据(dataUpdate).X过滤空值().X条件并识别主键(1).X更新()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)

		oneUpdate, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneUpdate["id"].X取整数(), 1)
		t.Assert(oneUpdate["name"].String(), "name_1000")
		t.Assert(oneUpdate["deleted_at"].String(), "")
		t.Assert(oneUpdate["created_at"].X取gtime时间类().X取时间戳秒(), oneInsert["created_at"].X取gtime时间类().X取时间戳秒())
		t.AssertGE(oneUpdate["updated_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-4)

		// Replace
		dataReplace := User{
			Id:   1,
			Name: "name_100",
		}
		r, err = db.X创建Model对象(table).X设置数据(dataReplace).X过滤空值().X插入并替换已存在()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 2)

		oneReplace, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneReplace["id"].X取整数(), 1)
		t.Assert(oneReplace["name"].String(), "name_100")
		t.Assert(oneReplace["deleted_at"].String(), "")
		t.AssertGE(oneReplace["created_at"].X取gtime时间类().X取时间戳秒(), oneInsert["created_at"].X取gtime时间类().X取时间戳秒())
		t.AssertGE(oneReplace["updated_at"].X取gtime时间类().X取时间戳秒(), oneInsert["updated_at"].X取gtime时间类().X取时间戳秒())

		// 用于时间断言的目的。
		time.Sleep(2 * time.Second)

		// Delete
		r, err = db.X创建Model对象(table).X删除("id", 1)
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)
		// Delete Select
		one4, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one4), 0)
		one5, err := db.X创建Model对象(table).X禁用时间自动更新特性().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one5["id"].X取整数(), 1)
		t.AssertGE(one5["deleted_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)
		// Delete Count
		i, err := db.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(i, 0)
		i, err = db.X创建Model对象(table).X禁用时间自动更新特性().X查询行数()
		t.AssertNil(err)
		t.Assert(i, 1)

		// Delete Unscoped
		r, err = db.X创建Model对象(table).X禁用时间自动更新特性().X删除("id", 1)
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)
		one6, err := db.X创建Model对象(table).X禁用时间自动更新特性().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one6), 0)
		i, err = db.X创建Model对象(table).X禁用时间自动更新特性().X查询行数()
		t.AssertNil(err)
		t.Assert(i, 0)
	})
}

func Test_SoftUpdateTime(t *testing.T) {
	table := "time_test_table_" + 时间类.X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id        int(11) NOT NULL,
  num       int(11) DEFAULT NULL,
  create_at datetime(6) DEFAULT NULL,
  update_at datetime(6) DEFAULT NULL,
  delete_at datetime(6) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, table)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		// Insert
		dataInsert := g.Map{
			"id":  1,
			"num": 10,
		}
		r, err := db.X创建Model对象(table).X设置数据(dataInsert).X插入()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		oneInsert, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneInsert["id"].X取整数(), 1)
		t.Assert(oneInsert["num"].X取整数(), 10)

		// Update.
		r, err = db.X创建Model对象(table).X设置数据("num=num+1").X条件("id=?", 1).X更新()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)
	})
}

func Test_SoftUpdateTime_WithDO(t *testing.T) {
	table := "time_test_table_" + 时间类.X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id        int(11) NOT NULL,
  num       int(11) DEFAULT NULL,
  created_at datetime(6) DEFAULT NULL,
  updated_at datetime(6) DEFAULT NULL,
  deleted_at datetime(6) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, table)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		// Insert
		dataInsert := g.Map{
			"id":  1,
			"num": 10,
		}
		r, err := db.X创建Model对象(table).X设置数据(dataInsert).X插入()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		oneInserted, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneInserted["id"].X取整数(), 1)
		t.Assert(oneInserted["num"].X取整数(), 10)

		// Update.
		time.Sleep(2 * time.Second)
		type User struct {
			g.Meta    `orm:"do:true"`
			Id        interface{}
			Num       interface{}
			CreatedAt interface{}
			UpdatedAt interface{}
			DeletedAt interface{}
		}
		r, err = db.X创建Model对象(table).X设置数据(User{
			Num: 100,
		}).X条件("id=?", 1).X更新()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)

		oneUpdated, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneUpdated["num"].X取整数(), 100)
		t.Assert(oneUpdated["created_at"].String(), oneInserted["created_at"].String())
		t.AssertNE(oneUpdated["updated_at"].String(), oneInserted["updated_at"].String())
	})
}

func Test_SoftDelete(t *testing.T) {
	table := "time_test_table_" + 时间类.X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id        int(11) NOT NULL,
  name      varchar(45) DEFAULT NULL,
  create_at datetime(6) DEFAULT NULL,
  update_at datetime(6) DEFAULT NULL,
  delete_at datetime(6) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, table)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(table)
	// db.SetDebug(true)
	单元测试类.C(t, func(t *单元测试类.T) {
		for i := 1; i <= 10; i++ {
			data := g.Map{
				"id":   i,
				"name": fmt.Sprintf("name_%d", i),
			}
			r, err := db.X创建Model对象(table).X设置数据(data).X插入()
			t.AssertNil(err)
			n, _ := r.RowsAffected()
			t.Assert(n, 1)
		}
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.AssertNE(one["create_at"].String(), "")
		t.AssertNE(one["update_at"].String(), "")
		t.Assert(one["delete_at"].String(), "")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		one, err := db.X创建Model对象(table).X条件并识别主键(10).X查询一条()
		t.AssertNil(err)
		t.AssertNE(one["create_at"].String(), "")
		t.AssertNE(one["update_at"].String(), "")
		t.Assert(one["delete_at"].String(), "")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		ids := g.SliceInt别名{1, 3, 5}
		r, err := db.X创建Model对象(table).X条件("id", ids).X删除()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 3)

		count, err := db.X创建Model对象(table).X条件("id", ids).X查询行数()
		t.AssertNil(err)
		t.Assert(count, 0)

		all, err := db.X创建Model对象(table).X禁用时间自动更新特性().X条件("id", ids).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 3)
		t.AssertNE(all[0]["create_at"].String(), "")
		t.AssertNE(all[0]["update_at"].String(), "")
		t.AssertNE(all[0]["delete_at"].String(), "")
		t.AssertNE(all[1]["create_at"].String(), "")
		t.AssertNE(all[1]["update_at"].String(), "")
		t.AssertNE(all[1]["delete_at"].String(), "")
		t.AssertNE(all[2]["create_at"].String(), "")
		t.AssertNE(all[2]["update_at"].String(), "")
		t.AssertNE(all[2]["delete_at"].String(), "")
	})
}

func Test_SoftDelete_Join(t *testing.T) {
	table1 := "time_test_table1"
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id        int(11) NOT NULL,
  name      varchar(45) DEFAULT NULL,
  create_at datetime(6) DEFAULT NULL,
  update_at datetime(6) DEFAULT NULL,
  delete_at datetime(6) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, table1)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(table1)

	table2 := "time_test_table2"
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id        int(11) NOT NULL,
  name      varchar(45) DEFAULT NULL,
  createat datetime(6) DEFAULT NULL,
  updateat datetime(6) DEFAULT NULL,
  deleteat datetime(6) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, table2)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(table2)

	单元测试类.C(t, func(t *单元测试类.T) {
		// db.SetDebug(true)
		dataInsert1 := g.Map{
			"id":   1,
			"name": "name_1",
		}
		r, err := db.X创建Model对象(table1).X设置数据(dataInsert1).X插入()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		dataInsert2 := g.Map{
			"id":   1,
			"name": "name_2",
		}
		r, err = db.X创建Model对象(table2).X设置数据(dataInsert2).X插入()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table1, "t1").X左连接(table2, "t2", "t2.id=t1.id").X字段保留过滤("t1.name").X查询一条()
		t.AssertNil(err)
		t.Assert(one["name"], "name_1")

		// Soft deleting.
		r, err = db.X创建Model对象(table1).X条件(1).X删除()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)

		one, err = db.X创建Model对象(table1, "t1").X左连接(table2, "t2", "t2.id=t1.id").X字段保留过滤("t1.name").X查询一条()
		t.AssertNil(err)
		t.Assert(one.X是否为空(), true)

		one, err = db.X创建Model对象(table2, "t2").X左连接(table1, "t1", "t2.id=t1.id").X字段保留过滤("t2.name").X查询一条()
		t.AssertNil(err)
		t.Assert(one.X是否为空(), true)
	})
}

func Test_SoftDelete_WhereAndOr(t *testing.T) {
	table := "time_test_table_" + 时间类.X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id        int(11) NOT NULL,
  name      varchar(45) DEFAULT NULL,
  create_at datetime(6) DEFAULT NULL,
  update_at datetime(6) DEFAULT NULL,
  delete_at datetime(6) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, table)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(table)
	// 设置数据库调试模式为开启状态。
// db.SetDebug(true)
// 添加数据。
// Add datas.
// 这里的"datas"在中文中一般表示为“数据”，但在Go语言编程中，根据变量或方法的具体含义，可能会有不同的翻译。如果这里的"Add datas"是向数据库添加数据的意思，那么可以翻译为：
// ```go
// 开启数据库调试模式。
// db.SetDebug(true)
// 添加数据。
// db.AddData(...)
// 但请注意，具体的翻译需要根据上下文和实际代码逻辑进行准确理解。
	单元测试类.C(t, func(t *单元测试类.T) {
		for i := 1; i <= 10; i++ {
			data := g.Map{
				"id":   i,
				"name": fmt.Sprintf("name_%d", i),
			}
			r, err := db.X创建Model对象(table).X设置数据(data).X插入()
			t.AssertNil(err)
			n, _ := r.RowsAffected()
			t.Assert(n, 1)
		}
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		ids := g.SliceInt别名{1, 3, 5}
		r, err := db.X创建Model对象(table).X条件("id", ids).X删除()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 3)

		count, err := db.X创建Model对象(table).X条件("id", 1).X条件或("id", 3).X查询行数()
		t.AssertNil(err)
		t.Assert(count, 0)
	})
}

func Test_CreateUpdateTime_Struct(t *testing.T) {
	table := "time_test_table_" + 时间类.X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id        int(11) NOT NULL,
  name      varchar(45) DEFAULT NULL,
  create_at datetime(6) DEFAULT NULL,
  update_at datetime(6) DEFAULT NULL,
  delete_at datetime(6) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, table)); err != nil {
		单元测试类.Error(err)
	}
	defer dropTable(table)

	// 设置数据库调试模式为开启状态
// db.SetDebug(true)
// 在函数结束时，确保关闭数据库调试模式
// defer db.SetDebug(false)

	type Entity struct {
		Id       uint64      `orm:"id,primary" json:"id"`
		Name     string      `orm:"name"       json:"name"`
		CreateAt *时间类.Time `orm:"create_at"  json:"create_at"`
		UpdateAt *时间类.Time `orm:"update_at"  json:"update_at"`
		DeleteAt *时间类.Time `orm:"delete_at"  json:"delete_at"`
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		// Insert
		dataInsert := &Entity{
			Id:       1,
			Name:     "name_1",
			CreateAt: nil,
			UpdateAt: nil,
			DeleteAt: nil,
		}
		r, err := db.X创建Model对象(table).X设置数据(dataInsert).X过滤空值().X插入()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		oneInsert, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneInsert["id"].X取整数(), 1)
		t.Assert(oneInsert["name"].String(), "name_1")
		t.Assert(oneInsert["delete_at"].String(), "")
		t.AssertGE(oneInsert["create_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)
		t.AssertGE(oneInsert["update_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)

		time.Sleep(2 * time.Second)

		// Save
		dataSave := &Entity{
			Id:       1,
			Name:     "name_10",
			CreateAt: nil,
			UpdateAt: nil,
			DeleteAt: nil,
		}
		r, err = db.X创建Model对象(table).X设置数据(dataSave).X过滤空值().X插入并更新已存在()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 2)

		oneSave, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneSave["id"].X取整数(), 1)
		t.Assert(oneSave["name"].String(), "name_10")
		t.Assert(oneSave["delete_at"].String(), "")
		t.Assert(oneSave["create_at"].X取gtime时间类().X取时间戳秒(), oneInsert["create_at"].X取gtime时间类().X取时间戳秒())
		t.AssertNE(oneSave["update_at"].X取gtime时间类().X取时间戳秒(), oneInsert["update_at"].X取gtime时间类().X取时间戳秒())
		t.AssertGE(oneSave["update_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)

		time.Sleep(2 * time.Second)

		// Update
		dataUpdate := &Entity{
			Id:       1,
			Name:     "name_1000",
			CreateAt: nil,
			UpdateAt: nil,
			DeleteAt: nil,
		}
		r, err = db.X创建Model对象(table).X设置数据(dataUpdate).X条件并识别主键(1).X过滤空值().X更新()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)

		oneUpdate, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneUpdate["id"].X取整数(), 1)
		t.Assert(oneUpdate["name"].String(), "name_1000")
		t.Assert(oneUpdate["delete_at"].String(), "")
		t.Assert(oneUpdate["create_at"].X取gtime时间类().X取时间戳秒(), oneInsert["create_at"].X取gtime时间类().X取时间戳秒())
		t.AssertGE(oneUpdate["update_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)

		// Replace
		dataReplace := &Entity{
			Id:       1,
			Name:     "name_100",
			CreateAt: nil,
			UpdateAt: nil,
			DeleteAt: nil,
		}
		r, err = db.X创建Model对象(table).X设置数据(dataReplace).X过滤空值().X插入并替换已存在()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 2)

		oneReplace, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(oneReplace["id"].X取整数(), 1)
		t.Assert(oneReplace["name"].String(), "name_100")
		t.Assert(oneReplace["delete_at"].String(), "")
		t.AssertGE(oneReplace["create_at"].X取gtime时间类().X取时间戳秒(), oneInsert["create_at"].X取gtime时间类().X取时间戳秒())
		t.AssertGE(oneReplace["update_at"].X取gtime时间类().X取时间戳秒(), oneInsert["update_at"].X取gtime时间类().X取时间戳秒())

		time.Sleep(2 * time.Second)

		// Delete
		r, err = db.X创建Model对象(table).X删除("id", 1)
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)
		// Delete Select
		one4, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one4), 0)
		one5, err := db.X创建Model对象(table).X禁用时间自动更新特性().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one5["id"].X取整数(), 1)
		t.AssertGE(one5["delete_at"].X取gtime时间类().X取时间戳秒(), 时间类.X取时间戳秒()-2)
		// Delete Count
		i, err := db.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(i, 0)
		i, err = db.X创建Model对象(table).X禁用时间自动更新特性().X查询行数()
		t.AssertNil(err)
		t.Assert(i, 1)

		// Delete Unscoped
		r, err = db.X创建Model对象(table).X禁用时间自动更新特性().X删除("id", 1)
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)
		one6, err := db.X创建Model对象(table).X禁用时间自动更新特性().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one6), 0)
		i, err = db.X创建Model对象(table).X禁用时间自动更新特性().X查询行数()
		t.AssertNil(err)
		t.Assert(i, 0)
	})
}
