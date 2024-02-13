// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package db类

import (
	"context"
	"fmt"
	"testing"
	
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gregex"
)

var (
	db  DB
	ctx = context.TODO()
)

func Test_HookSelect_Regex(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			err              error
			toBeCommittedSql = `select * from "user" where 1=1`
		)
		toBeCommittedSql, err = 正则类.ReplaceStringFuncMatch(
			`(?i) FROM ([\S]+)`,
			toBeCommittedSql,
			func(match []string) string {

				return fmt.Sprintf(` FROM "%s"`, "user_1")
			},
		)
		t.AssertNil(err)
		t.Assert(toBeCommittedSql, `select * FROM "user_1" where 1=1`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			err              error
			toBeCommittedSql = `select * from user`
		)
		toBeCommittedSql, err = 正则类.ReplaceStringFuncMatch(
			`(?i) FROM ([\S]+)`,
			toBeCommittedSql,
			func(match []string) string {
				return fmt.Sprintf(` FROM %s`, "user_1")
			},
		)
		t.AssertNil(err)
		t.Assert(toBeCommittedSql, `select * FROM user_1`)
	})
}

func Test_parseConfigNodeLink_WithType(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		node := &X配置项{
			X自定义链接信息: `mysql:root:CxzhD*624:27jh@tcp(9.135.69.119:3306)/khaos_oss?loc=Local&parseTime=true&charset=latin`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.X类型, `mysql`)
		t.Assert(newNode.X账号, `root`)
		t.Assert(newNode.X密码, `CxzhD*624:27jh`)
		t.Assert(newNode.X地址, `9.135.69.119`)
		t.Assert(newNode.X端口, `3306`)
		t.Assert(newNode.X名称, `khaos_oss`)
		t.Assert(newNode.X额外, `loc=Local&parseTime=true&charset=latin`)
		t.Assert(newNode.X字符集, `latin`)
		t.Assert(newNode.X协议, `tcp`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		node := &X配置项{
			X自定义链接信息: `mysql:root:CxzhD*624:27jh@tcp(9.135.69.119:3306)/khaos_oss?`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.X类型, `mysql`)
		t.Assert(newNode.X账号, `root`)
		t.Assert(newNode.X密码, `CxzhD*624:27jh`)
		t.Assert(newNode.X地址, `9.135.69.119`)
		t.Assert(newNode.X端口, `3306`)
		t.Assert(newNode.X名称, `khaos_oss`)
		t.Assert(newNode.X额外, ``)
		t.Assert(newNode.X字符集, defaultCharset)
		t.Assert(newNode.X协议, `tcp`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		node := &X配置项{
			X自定义链接信息: `mysql:root:CxzhD*624:27jh@tcp(9.135.69.119:3306)/khaos_oss`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.X类型, `mysql`)
		t.Assert(newNode.X账号, `root`)
		t.Assert(newNode.X密码, `CxzhD*624:27jh`)
		t.Assert(newNode.X地址, `9.135.69.119`)
		t.Assert(newNode.X端口, `3306`)
		t.Assert(newNode.X名称, `khaos_oss`)
		t.Assert(newNode.X额外, ``)
		t.Assert(newNode.X字符集, defaultCharset)
		t.Assert(newNode.X协议, `tcp`)
	})
	// 空数据库预选择。
	单元测试类.C(t, func(t *单元测试类.T) {
		node := &X配置项{
			X自定义链接信息: `mysql:root:CxzhD*624:27jh@tcp(9.135.69.119:3306)/?loc=Local&parseTime=true&charset=latin`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.X类型, `mysql`)
		t.Assert(newNode.X账号, `root`)
		t.Assert(newNode.X密码, `CxzhD*624:27jh`)
		t.Assert(newNode.X地址, `9.135.69.119`)
		t.Assert(newNode.X端口, `3306`)
		t.Assert(newNode.X名称, ``)
		t.Assert(newNode.X额外, `loc=Local&parseTime=true&charset=latin`)
		t.Assert(newNode.X字符集, `latin`)
		t.Assert(newNode.X协议, `tcp`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		node := &X配置项{
			X自定义链接信息: `mysql:root:CxzhD*624:27jh@tcp(9.135.69.119:3306)?loc=Local&parseTime=true&charset=latin`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.X类型, `mysql`)
		t.Assert(newNode.X账号, `root`)
		t.Assert(newNode.X密码, `CxzhD*624:27jh`)
		t.Assert(newNode.X地址, `9.135.69.119`)
		t.Assert(newNode.X端口, `3306`)
		t.Assert(newNode.X名称, ``)
		t.Assert(newNode.X额外, `loc=Local&parseTime=true&charset=latin`)
		t.Assert(newNode.X字符集, `latin`)
		t.Assert(newNode.X协议, `tcp`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		node := &X配置项{
			X自定义链接信息: `mysql:root:CxzhD*624:27jh@tcp(9.135.69.119:3306)/`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.X类型, `mysql`)
		t.Assert(newNode.X账号, `root`)
		t.Assert(newNode.X密码, `CxzhD*624:27jh`)
		t.Assert(newNode.X地址, `9.135.69.119`)
		t.Assert(newNode.X端口, `3306`)
		t.Assert(newNode.X名称, ``)
		t.Assert(newNode.X额外, ``)
		t.Assert(newNode.X字符集, defaultCharset)
		t.Assert(newNode.X协议, `tcp`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		node := &X配置项{
			X自定义链接信息: `mysql:root:CxzhD*624:27jh@tcp(9.135.69.119:3306)`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.X类型, `mysql`)
		t.Assert(newNode.X账号, `root`)
		t.Assert(newNode.X密码, `CxzhD*624:27jh`)
		t.Assert(newNode.X地址, `9.135.69.119`)
		t.Assert(newNode.X端口, `3306`)
		t.Assert(newNode.X名称, ``)
		t.Assert(newNode.X额外, ``)
		t.Assert(newNode.X字符集, defaultCharset)
		t.Assert(newNode.X协议, `tcp`)
	})
	// udp.
	单元测试类.C(t, func(t *单元测试类.T) {
		node := &X配置项{
			X自定义链接信息: `mysql:root:CxzhD*624:27jh@udp(9.135.69.119:3306)`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.X类型, `mysql`)
		t.Assert(newNode.X账号, `root`)
		t.Assert(newNode.X密码, `CxzhD*624:27jh`)
		t.Assert(newNode.X地址, `9.135.69.119`)
		t.Assert(newNode.X端口, `3306`)
		t.Assert(newNode.X名称, ``)
		t.Assert(newNode.X额外, ``)
		t.Assert(newNode.X字符集, defaultCharset)
		t.Assert(newNode.X协议, `udp`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		node := &X配置项{
			X自定义链接信息: `sqlite:root:CxzhD*624:27jh@file(/var/data/db.sqlite3)?local=Local&parseTime=true`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.X类型, `sqlite`)
		t.Assert(newNode.X账号, `root`)
		t.Assert(newNode.X密码, `CxzhD*624:27jh`)
		t.Assert(newNode.X地址, ``)
		t.Assert(newNode.X端口, ``)
		t.Assert(newNode.X名称, `/var/data/db.sqlite3`)
		t.Assert(newNode.X额外, `local=Local&parseTime=true`)
		t.Assert(newNode.X字符集, defaultCharset)
		t.Assert(newNode.X协议, `file`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		node := &X配置项{
			X自定义链接信息: `sqlite::CxzhD*624:2@7jh@file(/var/data/db.sqlite3)`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.X类型, `sqlite`)
		t.Assert(newNode.X账号, ``)
		t.Assert(newNode.X密码, `CxzhD*624:2@7jh`)
		t.Assert(newNode.X地址, ``)
		t.Assert(newNode.X端口, ``)
		t.Assert(newNode.X名称, `/var/data/db.sqlite3`)
		t.Assert(newNode.X额外, ``)
		t.Assert(newNode.X字符集, defaultCharset)
		t.Assert(newNode.X协议, `file`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		node := &X配置项{
			X自定义链接信息: `sqlite::@file(/var/data/db.sqlite3)`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.X类型, `sqlite`)
		t.Assert(newNode.X账号, ``)
		t.Assert(newNode.X密码, ``)
		t.Assert(newNode.X地址, ``)
		t.Assert(newNode.X端口, ``)
		t.Assert(newNode.X名称, `/var/data/db.sqlite3`)
		t.Assert(newNode.X额外, ``)
		t.Assert(newNode.X字符集, defaultCharset)
		t.Assert(newNode.X协议, `file`)
	})
	// #3146
	单元测试类.C(t, func(t *单元测试类.T) {
		node := &X配置项{
			X自定义链接信息: `pgsql:BASIC$xxxx:123456@tcp(xxxx.hologres.aliyuncs.com:80)/xxx`,
		}
		newNode := parseConfigNodeLink(node)
		t.Assert(newNode.X类型, `pgsql`)
		t.Assert(newNode.X账号, `BASIC$xxxx`)
		t.Assert(newNode.X密码, `123456`)
		t.Assert(newNode.X地址, `xxxx.hologres.aliyuncs.com`)
		t.Assert(newNode.X端口, `80`)
		t.Assert(newNode.X名称, `xxx`)
		t.Assert(newNode.X额外, ``)
		t.Assert(newNode.X字符集, defaultCharset)
		t.Assert(newNode.X协议, `tcp`)
	})
}

func Test_Func_doQuoteWord(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := map[string]string{
			"user":                   "`user`",
			"user u":                 "user u",
			"user_detail":            "`user_detail`",
			"user,user_detail":       "user,user_detail",
			"user u, user_detail ut": "user u, user_detail ut",
			"u.id asc":               "u.id asc",
			"u.id asc, ut.uid desc":  "u.id asc, ut.uid desc",
		}
		for k, v := range array {
			t.Assert(doQuoteWord(k, "`", "`"), v)
		}
	})
}

func Test_Func_doQuoteString(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		array := map[string]string{
			"user":                             "`user`",
			"user u":                           "`user` u",
			"user,user_detail":                 "`user`,`user_detail`",
			"user u, user_detail ut":           "`user` u,`user_detail` ut",
			"u.id, u.name, u.age":              "`u`.`id`,`u`.`name`,`u`.`age`",
			"u.id asc":                         "`u`.`id` asc",
			"u.id asc, ut.uid desc":            "`u`.`id` asc,`ut`.`uid` desc",
			"user.user u, user.user_detail ut": "`user`.`user` u,`user`.`user_detail` ut",
			// 使用双点操作符进行全局访问mssql模式
			"user..user u, user.user_detail ut": "`user`..`user` u,`user`.`user_detail` ut",
		}
		for k, v := range array {
			t.Assert(doQuoteString(k, "`", "`"), v)
		}
	})
}

func Test_Func_addTablePrefix(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := ""
		array := map[string]string{
			"user":                         "`user`",
			"user u":                       "`user` u",
			"user as u":                    "`user` as u",
			"user,user_detail":             "`user`,`user_detail`",
			"user u, user_detail ut":       "`user` u,`user_detail` ut",
			"`user`.user_detail":           "`user`.`user_detail`",
			"`user`.`user_detail`":         "`user`.`user_detail`",
			"user as u, user_detail as ut": "`user` as u,`user_detail` as ut",
			"UserCenter.user as u, UserCenter.user_detail as ut": "`UserCenter`.`user` as u,`UserCenter`.`user_detail` as ut",
			// 使用双点操作符进行全局访问mssql模式
			"UserCenter..user as u, user_detail as ut": "`UserCenter`..`user` as u,`user_detail` as ut",
		}
		for k, v := range array {
			t.Assert(doQuoteTableName(k, prefix, "`", "`"), v)
		}
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		prefix := "gf_"
		array := map[string]string{
			"user":                         "`gf_user`",
			"user u":                       "`gf_user` u",
			"user as u":                    "`gf_user` as u",
			"user,user_detail":             "`gf_user`,`gf_user_detail`",
			"user u, user_detail ut":       "`gf_user` u,`gf_user_detail` ut",
			"`user`.user_detail":           "`user`.`gf_user_detail`",
			"`user`.`user_detail`":         "`user`.`gf_user_detail`",
			"user as u, user_detail as ut": "`gf_user` as u,`gf_user_detail` as ut",
			"UserCenter.user as u, UserCenter.user_detail as ut": "`UserCenter`.`gf_user` as u,`UserCenter`.`gf_user_detail` as ut",
			// 使用双点操作符进行全局访问mssql模式
			"UserCenter..user as u, user_detail as ut": "`UserCenter`..`gf_user` as u,`gf_user_detail` as ut",
		}
		for k, v := range array {
			t.Assert(doQuoteTableName(k, prefix, "`", "`"), v)
		}
	})
}

func Test_isSubQuery(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(isSubQuery("user"), false)
		t.Assert(isSubQuery("user.uid"), false)
		t.Assert(isSubQuery("u, user.uid"), false)
		t.Assert(isSubQuery("select 1"), true)
	})
}
