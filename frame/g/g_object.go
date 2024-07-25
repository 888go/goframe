// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

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

// Client 是一个便利函数，它创建并返回一个新的 HTTP 客户端。 md5:70c95e835d086085
func Client() *gclient.Client {
	return gclient.New()
}

// Server 函数返回一个指定名称的 http 服务器实例。 md5:8613002877ca5b89
func Server(name ...interface{}) *ghttp.Server {
	return gins.Server(name...)
}

// TCPServer 返回一个具有指定名称的TCP服务器实例。 md5:57eae18c70a9e8c2
func TCPServer(name ...interface{}) *gtcp.Server {
	return gtcp.GetServer(name...)
}

// UDPServer 返回一个指定名称的UDP服务器实例。 md5:bffbc3c34f72d706
func UDPServer(name ...interface{}) *gudp.Server {
	return gudp.GetServer(name...)
}

// View 返回指定名称的模板引擎对象实例。 md5:090a094475513c16
func View(name ...string) *gview.View {
	return gins.View(name...)
}

// Config返回具有指定名称的config对象实例。 md5:68a54cf8b0afa4f0
func Config(name ...string) *gcfg.Config {
	return gins.Config(name...)
}

// Cfg 是 Config 的别名。
// 参见 Config。 md5:1b7d381469f7e212
func Cfg(name ...string) *gcfg.Config {
	return Config(name...)
}

// Resource 返回一个 Resource 类型的实例。
// 参数 `name` 为该实例的名称。 md5:42e664c4b3a2bb54
func Resource(name ...string) *gres.Resource {
	return gins.Resource(name...)
}

// I18n 返回一个 gi18n.Manager 的实例。
// 参数 `name` 是实例的名称。 md5:cb8fb8e2c93c597b
func I18n(name ...string) *gi18n.Manager {
	return gins.I18n(name...)
}

// Res是Resource的别名。
// 请参阅Resource。 md5:f2fc8f778d75b07c
func Res(name ...string) *gres.Resource {
	return Resource(name...)
}

// Log 函数返回一个 glog.Logger 实例。
// 参数 `name` 用于指定该实例的名称。 md5:01b7637de189b37b
func Log(name ...string) *glog.Logger {
	return gins.Log(name...)
}

// DB 根据指定的配置组名称返回一个数据库ORM对象实例。 md5:229506bfec6bd334
func DB(name ...string) gdb.DB {
	return gins.Database(name...)
}

// Model 根据默认数据库组的配置创建并返回一个模型。 md5:544c43fcc039bdfd
func Model(tableNameOrStruct ...interface{}) *gdb.Model {
	return DB().Model(tableNameOrStruct...)
}

// ModelRaw根据原始SQL（而非表）创建并返回一个模型。 md5:674c1e89391cbf02
func ModelRaw(rawSql string, args ...interface{}) *gdb.Model {
	return DB().Raw(rawSql, args...)
}

// Redis 根据给定的配置组名称返回一个redis客户端实例。 md5:3fab97150e392125
func Redis(name ...string) *gredis.Redis {
	return gins.Redis(name...)
}

// Validator 是一个便利的函数，它创建并返回一个新的验证管理器对象。 md5:de213201906601b7
func Validator() *gvalid.Validator {
	return gvalid.New()
}
