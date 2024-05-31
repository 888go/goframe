// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package g

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/gins"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/net/gudp"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/util/gvalid"
)

// Client is a convenience function, which creates and returns a new HTTP client.

// ff:网页类
func Client() *gclient.Client {
	return gclient.New()
}

// Server returns an instance of http server with specified name.

// ff:Http类
// name:名称
func Server(name ...interface{}) *ghttp.Server {
	return gins.Server(name...)
}

// TCPServer returns an instance of tcp server with specified name.

// ff:TCP类
// name:名称
func TCPServer(name ...interface{}) *gtcp.Server {
	return gtcp.GetServer(name...)
}

// UDPServer returns an instance of udp server with specified name.

// ff:UDP类
// name:名称
func UDPServer(name ...interface{}) *gudp.Server {
	return gudp.GetServer(name...)
}

// View returns an instance of template engine object with specified name.

// ff:模板类
// name:名称
func View(name ...string) *gview.View {
	return gins.View(name...)
}

// Config returns an instance of config object with specified name.

// ff:配置类
// name:名称
func Config(name ...string) *gcfg.Config {
	return gins.Config(name...)
}

// Cfg is alias of Config.
// See Config.

// ff:Cfg别名
// name:名称
func Cfg(name ...string) *gcfg.Config {
	return Config(name...)
}

// Resource returns an instance of Resource.
// The parameter `name` is the name for the instance.

// ff:资源类
// name:名称
func Resource(name ...string) *gres.Resource {
	return gins.Resource(name...)
}

// I18n returns an instance of gi18n.Manager.
// The parameter `name` is the name for the instance.

// ff:多语言类
// name:名称
func I18n(name ...string) *gi18n.Manager {
	return gins.I18n(name...)
}

// Res is alias of Resource.
// See Resource.

// ff:Res别名
// name:名称
func Res(name ...string) *gres.Resource {
	return Resource(name...)
}

// Log returns an instance of glog.Logger.
// The parameter `name` is the name for the instance.

// ff:日志类
// name:名称
func Log(name ...string) *glog.Logger {
	return gins.Log(name...)
}

// DB returns an instance of database ORM object with specified configuration group name.

// ff:DB类
// name:名称
func DB(name ...string) gdb.DB {
	return gins.Database(name...)
}

// Model creates and returns a model based on configuration of default database group.

// ff:DB类Model
// tableNameOrStruct:表名或结构体
func Model(tableNameOrStruct ...interface{}) *gdb.Model {
	return DB().Model(tableNameOrStruct...)
}

// ModelRaw creates and returns a model based on a raw sql not a table.

// ff:DB类原生SQL
// args:参数
// rawSql:原生Sql
func ModelRaw(rawSql string, args ...interface{}) *gdb.Model {
	return DB().Raw(rawSql, args...)
}

// Redis returns an instance of redis client with specified configuration group name.

// ff:Redis类
// name:名称
func Redis(name ...string) *gredis.Redis {
	return gins.Redis(name...)
}

// Validator is a convenience function, which creates and returns a new validation manager object.

// ff:效验类
func Validator() *gvalid.Validator {
	return gvalid.New()
}
