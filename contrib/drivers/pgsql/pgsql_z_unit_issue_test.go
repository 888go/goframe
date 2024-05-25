//go:build 屏蔽单元测试

// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package pgsql_test

import (
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
)

// https://github.com/gogf/gf/issues/3330
// 
// 这段注释引用的是一个GitHub问题或拉取请求，它来自于 "gogf/gf"（Go GF）项目，编号为3330。"gf" 是一个Go语言的框架（framework）。这个链接指向的是关于该框架的一个 issue（通常指代项目中遇到的问题、错误或特性请求），可能是用户报告了一个问题或者提出了一个改进的建议。. md5:9f7312fda44501ce
func Test_Issue3330(t *testing.T) {
	var (
		table      = fmt.Sprintf(`%s_%d`, TablePrefix+"test", gtime.TimestampNano())
		uniqueName = fmt.Sprintf(`%s_%d`, TablePrefix+"test_unique", gtime.TimestampNano())
	)
	if _, err := db.Exec(ctx, fmt.Sprintf(`
		CREATE TABLE %s (
		   	id bigserial  NOT NULL,
		   	passport varchar(45) NOT NULL,
		   	password varchar(32) NOT NULL,
		   	nickname varchar(45) NOT NULL,
		   	create_time timestamp NOT NULL,
		   	PRIMARY KEY (id),
			CONSTRAINT %s unique ("password")
		) ;`, table, uniqueName,
	)); err != nil {
		gtest.Fatal(err)
	}
	defer dropTable(table)

	gtest.C(t, func(t *gtest.T) {
		var (
			list []map[string]interface{}
			one  gdb.Record
			err  error
		)

		fields, err := db.TableFields(ctx, table)
		t.AssertNil(err)

		t.Assert(fields["id"].Key, "pri")
		t.Assert(fields["password"].Key, "uni")

		for i := 1; i <= 10; i++ {
			list = append(list, g.Map{
				"id":          i,
				"passport":    fmt.Sprintf("p%d", i),
				"password":    fmt.Sprintf("pw%d", i),
				"nickname":    fmt.Sprintf("n%d", i),
				"create_time": "2016-06-01 00:00:00",
			})
		}

		_, err = db.Model(table).Data(list).Insert()
		t.AssertNil(err)

		for i := 1; i <= 10; i++ {
			one, err = db.Model(table).WherePri(i).One()
			t.AssertNil(err)
			t.Assert(one["id"], list[i-1]["id"])
			t.Assert(one["passport"], list[i-1]["passport"])
			t.Assert(one["password"], list[i-1]["password"])
			t.Assert(one["nickname"], list[i-1]["nickname"])
		}
	})
}
