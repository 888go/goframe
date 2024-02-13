// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package mysql_test

import (
	"testing"
	
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gmeta"
)

func Test_Model_Builder(t *testing.T) {
	table := createInitTable()
	defer dropTable(table)

	单元测试类.C(t, func(t *单元测试类.T) {
		m := db.X创建Model对象(table)
		b := m.X创建组合条件()

		all, err := m.X条件(
			b.X条件("id", g.Slice别名{1, 2, 3}).X条件或("id", g.Slice别名{4, 5, 6}),
		).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 6)
	})

	// Where And
	单元测试类.C(t, func(t *单元测试类.T) {
		m := db.X创建Model对象(table)
		b := m.X创建组合条件()

		all, err := m.X条件(
			b.X条件("id", g.Slice别名{1, 2, 3}).X条件或("id", g.Slice别名{4, 5, 6}),
		).X条件(
			b.X条件("id", g.Slice别名{2, 3}).X条件或("id", g.Slice别名{5, 6}),
		).X条件(
			b.X条件("id", g.Slice别名{3}).X条件("id", g.Slice别名{1, 2, 3}),
		).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 1)
	})

	// Where Or
	单元测试类.C(t, func(t *单元测试类.T) {
		m := db.X创建Model对象(table)
		b := m.X创建组合条件()

		all, err := m.X条件或(
			b.X条件("id", g.Slice别名{1, 2, 3}).X条件或("id", g.Slice别名{4, 5, 6}),
		).X条件或(
			b.X条件("id", g.Slice别名{2, 3}).X条件或("id", g.Slice别名{5, 6}),
		).X条件或(
			b.X条件("id", g.Slice别名{3}).X条件("id", g.Slice别名{1, 2, 3}),
		).X查询()
		t.AssertNil(err)
		t.Assert(len(all), 6)
	})

	// Where 针对具有 *gtime.Time 类型字段的结构体
	单元测试类.C(t, func(t *单元测试类.T) {
		m := db.X创建Model对象(table)
		b := m.X创建组合条件()

		type Query struct {
			Id       interface{}
			Nickname *时间类.Time
		}

		where, args := b.X条件(&Query{Id: 1}).X生成条件字符串及参数()
		t.Assert(where, "`id`=? AND `nickname` IS NULL")
		t.Assert(args, []interface{}{1})
	})

	// Where 函数，用于处理具有字段类型为 *gjson.Json 的结构体
	单元测试类.C(t, func(t *单元测试类.T) {
		m := db.X创建Model对象(table)
		b := m.X创建组合条件()

		type Query struct {
			Id       interface{}
			Nickname *json类.Json
		}

		where, args := b.X条件(&Query{Id: 1}).X生成条件字符串及参数()
		t.Assert(where, "`id`=? AND `nickname` IS NULL")
		t.Assert(args, []interface{}{1})
	})

	// Where 函数配合使用具有 *gtime.Time 类型字段的结构体，该结构体由 gf cli 生成
	单元测试类.C(t, func(t *单元测试类.T) {
		m := db.X创建Model对象(table)
		b := m.X创建组合条件()

		type Query struct {
			元数据类.Meta `orm:"do:true"`
			Id         interface{}
			Nickname   *时间类.Time
		}

		where, args := b.X条件(&Query{Id: 1}).X生成条件字符串及参数()
		t.Assert(where, "`id`=?")
		t.Assert(args, []interface{}{1})
	})

	// Where 针对由 gf cli 生成的拥有 *gjson.Json 类型字段的 do 结构体
	单元测试类.C(t, func(t *单元测试类.T) {
		m := db.X创建Model对象(table)
		b := m.X创建组合条件()

		type Query struct {
			元数据类.Meta `orm:"do:true"`
			Id         interface{}
			Nickname   *json类.Json
		}

		where, args := b.X条件(&Query{Id: 1}).X生成条件字符串及参数()
		t.Assert(where, "`id`=?")
		t.Assert(args, []interface{}{1})
	})
}

func Test_Safe_Builder(t *testing.T) {
	// 测试m.Builder()是否支持链式安全调用
	单元测试类.C(t, func(t *单元测试类.T) {
		b := db.X创建Model对象().X创建组合条件()
		b.X条件("id", 1)
		_, args := b.X生成条件字符串及参数()
		t.AssertNil(args)

		b = b.X条件("id", 1)
		_, args = b.X生成条件字符串及参数()
		t.Assert(args, g.Slice别名{1})
	})
}
