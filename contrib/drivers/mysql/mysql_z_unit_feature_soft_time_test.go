// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package mysql_test

import (
	"fmt"
	"testing"
	"time"

	gdb "github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

// 创建时间/更新时间/删除时间。 md5:19b2f7344b6c3216
func Test_SoftTime_CreateUpdateDelete1(t *testing.T) {
	table := "soft_time_test_table_" + gtime.X取文本时间戳纳秒()
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
		gtest.Error(err)
	}
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
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
		t.AssertGE(oneInsert["create_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)
		t.AssertGE(oneInsert["update_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)

				// 用于时间断言目的。 md5:9b80204747a3e820
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
		t.AssertGE(oneSave["update_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)

				// 用于时间断言目的。 md5:9b80204747a3e820
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
		t.AssertGE(oneUpdate["update_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)

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

				// 用于时间断言目的。 md5:9b80204747a3e820
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
		one5, err := db.X创建Model对象(table).Unscoped().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one5["id"].X取整数(), 1)
		t.AssertGE(one5["delete_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)
		// Delete Count
		i, err := db.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(i, 0)
		i, err = db.X创建Model对象(table).Unscoped().X查询行数()
		t.AssertNil(err)
		t.Assert(i, 1)

		// Delete Unscoped
		r, err = db.X创建Model对象(table).Unscoped().X删除("id", 1)
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)
		one6, err := db.X创建Model对象(table).Unscoped().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one6), 0)
		i, err = db.X创建Model对象(table).Unscoped().X查询行数()
		t.AssertNil(err)
		t.Assert(i, 0)
	})
}

// 创建时间/更新时间/删除时间。 md5:19b2f7344b6c3216
func Test_SoftTime_CreateUpdateDelete2(t *testing.T) {
	table := "soft_time_test_table_" + gtime.X取文本时间戳纳秒()
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
		gtest.Error(err)
	}
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
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
		t.AssertGE(oneInsert["create_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)
		t.AssertGE(oneInsert["update_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)

				// 用于时间断言目的。 md5:9b80204747a3e820
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
		t.AssertGE(oneSave["update_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)

				// 用于时间断言目的。 md5:9b80204747a3e820
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
		t.AssertGE(oneUpdate["update_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)

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

				// 用于时间断言目的。 md5:9b80204747a3e820
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
		one5, err := db.X创建Model对象(table).Unscoped().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one5["id"].X取整数(), 1)
		t.AssertGE(one5["delete_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)
		// Delete Count
		i, err := db.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(i, 0)
		i, err = db.X创建Model对象(table).Unscoped().X查询行数()
		t.AssertNil(err)
		t.Assert(i, 1)

		// Delete Unscoped
		r, err = db.X创建Model对象(table).Unscoped().X删除("id", 1)
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)
		one6, err := db.X创建Model对象(table).Unscoped().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one6), 0)
		i, err = db.X创建Model对象(table).Unscoped().X查询行数()
		t.AssertNil(err)
		t.Assert(i, 0)
	})
}

// 创建时间/更新时间/删除时间。 md5:e6765f0f8b916a37
func Test_SoftTime_CreatedUpdatedDeleted_Map(t *testing.T) {
	table := "soft_time_test_table_" + gtime.X取文本时间戳纳秒()
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
		gtest.Error(err)
	}
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
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
		t.AssertGE(oneInsert["created_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)
		t.AssertGE(oneInsert["updated_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)

				// 用于时间断言目的。 md5:9b80204747a3e820
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
		t.AssertGE(oneSave["updated_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)

				// 用于时间断言目的。 md5:9b80204747a3e820
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
		t.AssertGE(oneUpdate["updated_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)

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

				// 用于时间断言目的。 md5:9b80204747a3e820
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
		one5, err := db.X创建Model对象(table).Unscoped().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one5["id"].X取整数(), 1)
		t.AssertGE(one5["deleted_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)
		// Delete Count
		i, err := db.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(i, 0)
		i, err = db.X创建Model对象(table).Unscoped().X查询行数()
		t.AssertNil(err)
		t.Assert(i, 1)

		// Delete Unscoped
		r, err = db.X创建Model对象(table).Unscoped().X删除("id", 1)
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)
		one6, err := db.X创建Model对象(table).Unscoped().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one6), 0)
		i, err = db.X创建Model对象(table).Unscoped().X查询行数()
		t.AssertNil(err)
		t.Assert(i, 0)
	})
}

// 创建时间/更新时间/删除时间。 md5:e6765f0f8b916a37
func Test_SoftTime_CreatedUpdatedDeleted_Struct(t *testing.T) {
	table := "soft_time_test_table_" + gtime.X取文本时间戳纳秒()
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
		gtest.Error(err)
	}
	defer dropTable(table)

	type User struct {
		Id        int
		Name      string
		CreatedAT *gtime.Time
		UpdatedAT *gtime.Time
		DeletedAT *gtime.Time
	}
	gtest.C(t, func(t *gtest.T) {
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
		t.AssertGE(oneInsert["created_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)
		t.AssertGE(oneInsert["updated_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)

				// 用于时间断言目的。 md5:9b80204747a3e820
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
		t.AssertGE(oneSave["updated_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)

				// 用于时间断言目的。 md5:9b80204747a3e820
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
		t.AssertGE(oneUpdate["updated_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-4)

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

				// 用于时间断言目的。 md5:9b80204747a3e820
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
		one5, err := db.X创建Model对象(table).Unscoped().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one5["id"].X取整数(), 1)
		t.AssertGE(one5["deleted_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)
		// Delete Count
		i, err := db.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(i, 0)
		i, err = db.X创建Model对象(table).Unscoped().X查询行数()
		t.AssertNil(err)
		t.Assert(i, 1)

		// Delete Unscoped
		r, err = db.X创建Model对象(table).Unscoped().X删除("id", 1)
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)
		one6, err := db.X创建Model对象(table).Unscoped().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one6), 0)
		i, err = db.X创建Model对象(table).Unscoped().X查询行数()
		t.AssertNil(err)
		t.Assert(i, 0)
	})
}

func Test_SoftUpdateTime(t *testing.T) {
	table := "soft_time_test_table_" + gtime.X取文本时间戳纳秒()
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
		gtest.Error(err)
	}
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
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
	table := "soft_time_test_table_" + gtime.X取文本时间戳纳秒()
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
		gtest.Error(err)
	}
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
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
	table := "soft_time_test_table_" + gtime.X取文本时间戳纳秒()
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
		gtest.Error(err)
	}
	defer dropTable(table)
	// db.SetDebug(true)
	gtest.C(t, func(t *gtest.T) {
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
	gtest.C(t, func(t *gtest.T) {
		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.AssertNE(one["create_at"].String(), "")
		t.AssertNE(one["update_at"].String(), "")
		t.Assert(one["delete_at"].String(), "")
	})
	gtest.C(t, func(t *gtest.T) {
		one, err := db.X创建Model对象(table).X条件并识别主键(10).X查询一条()
		t.AssertNil(err)
		t.AssertNE(one["create_at"].String(), "")
		t.AssertNE(one["update_at"].String(), "")
		t.Assert(one["delete_at"].String(), "")
	})
	gtest.C(t, func(t *gtest.T) {
		ids := g.SliceInt别名{1, 3, 5}
		r, err := db.X创建Model对象(table).X条件("id", ids).X删除()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 3)

		count, err := db.X创建Model对象(table).X条件("id", ids).X查询行数()
		t.AssertNil(err)
		t.Assert(count, 0)

		all, err := db.X创建Model对象(table).Unscoped().X条件("id", ids).X查询()
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
		gtest.Error(err)
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
		gtest.Error(err)
	}
	defer dropTable(table2)

	gtest.C(t, func(t *gtest.T) {
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
	table := "soft_time_test_table_" + gtime.X取文本时间戳纳秒()
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
		gtest.Error(err)
	}
	defer dropTable(table)
	// db.SetDebug(true)
	// 添加数据。
	// md5:940469df021bcc3f
	gtest.C(t, func(t *gtest.T) {
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
	gtest.C(t, func(t *gtest.T) {
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
	table := "soft_time_test_table_" + gtime.X取文本时间戳纳秒()
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
		gtest.Error(err)
	}
	defer dropTable(table)

	// 将数据库设置为调试模式
	// 使用defer语句确保在函数返回前将数据库的调试模式重置为false
	// md5:b9225b2fca692b91

	type Entity struct {
		Id       uint64      `orm:"id,primary" json:"id"`
		Name     string      `orm:"name"       json:"name"`
		CreateAt *gtime.Time `orm:"create_at"  json:"create_at"`
		UpdateAt *gtime.Time `orm:"update_at"  json:"update_at"`
		DeleteAt *gtime.Time `orm:"delete_at"  json:"delete_at"`
	}
	gtest.C(t, func(t *gtest.T) {
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
		t.AssertGE(oneInsert["create_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)
		t.AssertGE(oneInsert["update_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)

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
		t.AssertGE(oneSave["update_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)

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
		t.AssertGE(oneUpdate["update_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)

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
		one5, err := db.X创建Model对象(table).Unscoped().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one5["id"].X取整数(), 1)
		t.AssertGE(one5["delete_at"].X取gtime时间类().X取时间戳秒(), gtime.X取时间戳秒()-2)
		// Delete Count
		i, err := db.X创建Model对象(table).X查询行数()
		t.AssertNil(err)
		t.Assert(i, 0)
		i, err = db.X创建Model对象(table).Unscoped().X查询行数()
		t.AssertNil(err)
		t.Assert(i, 1)

		// Delete Unscoped
		r, err = db.X创建Model对象(table).Unscoped().X删除("id", 1)
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)
		one6, err := db.X创建Model对象(table).Unscoped().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one6), 0)
		i, err = db.X创建Model对象(table).Unscoped().X查询行数()
		t.AssertNil(err)
		t.Assert(i, 0)
	})
}

func Test_SoftTime_CreateUpdateDelete_UnixTimestamp(t *testing.T) {
	table := "soft_time_test_table_" + gtime.X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id        int(11) NOT NULL,
  name      varchar(45) DEFAULT NULL,
  create_at int(11) DEFAULT NULL,
  update_at int(11) DEFAULT NULL,
  delete_at int(11) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, table)); err != nil {
		gtest.Error(err)
	}
	defer dropTable(table)

	// insert
	gtest.C(t, func(t *gtest.T) {
		dataInsert := g.Map{
			"id":   1,
			"name": "name_1",
		}
		r, err := db.X创建Model对象(table).X设置数据(dataInsert).X插入()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["name"].String(), "name_1")
		t.AssertGT(one["create_at"].X取整数64位(), 0)
		t.AssertGT(one["update_at"].X取整数64位(), 0)
		t.Assert(one["delete_at"].X取整数64位(), 0)
		t.Assert(len(one["create_at"].String()), 10)
		t.Assert(len(one["update_at"].String()), 10)
	})

		// 睡眠一些秒，以使更新时间大于创建时间。 md5:7c7908ddbd15e9a3
	time.Sleep(2 * time.Second)

	// update
	gtest.C(t, func(t *gtest.T) {
		// update: map
		dataInsert := g.Map{
			"name": "name_11",
		}
		r, err := db.X创建Model对象(table).X设置数据(dataInsert).X条件并识别主键(1).X更新()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["name"].String(), "name_11")
		t.AssertGT(one["create_at"].X取整数64位(), 0)
		t.AssertGT(one["update_at"].X取整数64位(), 0)
		t.Assert(one["delete_at"].X取整数64位(), 0)
		t.Assert(len(one["create_at"].String()), 10)
		t.Assert(len(one["update_at"].String()), 10)

		var (
			lastCreateTime = one["create_at"].X取整数64位()
			lastUpdateTime = one["update_at"].X取整数64位()
		)

		time.Sleep(2 * time.Second)

		// update: string
		r, err = db.X创建Model对象(table).X设置数据("name='name_111'").X条件并识别主键(1).X更新()
		t.AssertNil(err)
		n, _ = r.RowsAffected()
		t.Assert(n, 1)

		one, err = db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["name"].String(), "name_111")
		t.Assert(one["create_at"].X取整数64位(), lastCreateTime)
		t.AssertGT(one["update_at"].X取整数64位(), lastUpdateTime)
		t.Assert(one["delete_at"].X取整数64位(), 0)
	})

	// delete
	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table).X条件并识别主键(1).X删除()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one), 0)

		one, err = db.X创建Model对象(table).Unscoped().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["name"].String(), "name_111")
		t.AssertGT(one["create_at"].X取整数64位(), 0)
		t.AssertGT(one["update_at"].X取整数64位(), 0)
		t.AssertGT(one["delete_at"].X取整数64位(), 0)
	})
}

func Test_SoftTime_CreateUpdateDelete_Bool_Deleted(t *testing.T) {
	table := "soft_time_test_table_" + gtime.X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id        int(11) NOT NULL,
  name      varchar(45) DEFAULT NULL,
  create_at int(11) DEFAULT NULL,
  update_at int(11) DEFAULT NULL,
  delete_at bit(1) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, table)); err != nil {
		gtest.Error(err)
	}
	defer dropTable(table)

	//db.SetDebug(true) 	// 设置数据库调试模式为开启
	// insert 	// 插入数据操作
	// md5:43a7c855e5c6ebf5
	gtest.C(t, func(t *gtest.T) {
		dataInsert := g.Map{
			"id":   1,
			"name": "name_1",
		}
		r, err := db.X创建Model对象(table).X设置数据(dataInsert).X插入()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["name"].String(), "name_1")
		t.AssertGT(one["create_at"].X取整数64位(), 0)
		t.AssertGT(one["update_at"].X取整数64位(), 0)
		t.Assert(one["delete_at"].X取整数64位(), 0)
		t.Assert(len(one["create_at"].String()), 10)
		t.Assert(len(one["update_at"].String()), 10)
	})

	// delete
	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table).X条件并识别主键(1).X删除()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one), 0)

		one, err = db.X创建Model对象(table).Unscoped().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["name"].String(), "name_1")
		t.AssertGT(one["create_at"].X取整数64位(), 0)
		t.AssertGT(one["update_at"].X取整数64位(), 0)
		t.Assert(one["delete_at"].X取整数64位(), 1)
	})
}

func Test_SoftTime_CreateUpdateDelete_Option_SoftTimeTypeTimestampMilli(t *testing.T) {
	table := "soft_time_test_table_" + gtime.X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id        int(11) NOT NULL,
  name      varchar(45) DEFAULT NULL,
  create_at bigint(19) unsigned DEFAULT NULL,
  update_at bigint(19) unsigned DEFAULT NULL,
  delete_at bit(1) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, table)); err != nil {
		gtest.Error(err)
	}
	defer dropTable(table)

	var softTimeOption = gdb.SoftTimeOption{
		SoftTimeType: gdb.SoftTimeTypeTimestampMilli,
	}

	// insert
	gtest.C(t, func(t *gtest.T) {
		dataInsert := g.Map{
			"id":   1,
			"name": "name_1",
		}
		r, err := db.X创建Model对象(table).SoftTime(softTimeOption).X设置数据(dataInsert).X插入()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).SoftTime(softTimeOption).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["name"].String(), "name_1")
		t.Assert(len(one["create_at"].String()), 13)
		t.Assert(len(one["update_at"].String()), 13)
		t.Assert(one["delete_at"].X取整数64位(), 0)
	})

	// delete
	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table).SoftTime(softTimeOption).X条件并识别主键(1).X删除()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).SoftTime(softTimeOption).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one), 0)

		one, err = db.X创建Model对象(table).Unscoped().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["name"].String(), "name_1")
		t.AssertGT(one["create_at"].X取整数64位(), 0)
		t.AssertGT(one["update_at"].X取整数64位(), 0)
		t.Assert(one["delete_at"].X取整数64位(), 1)
	})
}

func Test_SoftTime_CreateUpdateDelete_Option_SoftTimeTypeTimestampNano(t *testing.T) {
	table := "soft_time_test_table_" + gtime.X取文本时间戳纳秒()
	if _, err := db.X原生SQL执行(ctx, fmt.Sprintf(`
CREATE TABLE %s (
  id        int(11) NOT NULL,
  name      varchar(45) DEFAULT NULL,
  create_at bigint(19) unsigned DEFAULT NULL,
  update_at bigint(19) unsigned DEFAULT NULL,
  delete_at bit(1) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, table)); err != nil {
		gtest.Error(err)
	}
	defer dropTable(table)

	var softTimeOption = gdb.SoftTimeOption{
		SoftTimeType: gdb.SoftTimeTypeTimestampNano,
	}

	// insert
	gtest.C(t, func(t *gtest.T) {
		dataInsert := g.Map{
			"id":   1,
			"name": "name_1",
		}
		r, err := db.X创建Model对象(table).SoftTime(softTimeOption).X设置数据(dataInsert).X插入()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).SoftTime(softTimeOption).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["name"].String(), "name_1")
		t.Assert(len(one["create_at"].String()), 19)
		t.Assert(len(one["update_at"].String()), 19)
		t.Assert(one["delete_at"].X取整数64位(), 0)
	})

	// delete
	gtest.C(t, func(t *gtest.T) {
		r, err := db.X创建Model对象(table).SoftTime(softTimeOption).X条件并识别主键(1).X删除()
		t.AssertNil(err)
		n, _ := r.RowsAffected()
		t.Assert(n, 1)

		one, err := db.X创建Model对象(table).SoftTime(softTimeOption).X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(len(one), 0)

		one, err = db.X创建Model对象(table).Unscoped().X条件并识别主键(1).X查询一条()
		t.AssertNil(err)
		t.Assert(one["name"].String(), "name_1")
		t.AssertGT(one["create_at"].X取整数64位(), 0)
		t.AssertGT(one["update_at"].X取整数64位(), 0)
		t.Assert(one["delete_at"].X取整数64位(), 1)
	})
}
