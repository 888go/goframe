// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package g

import (
	gdb "github.com/888go/goframe/database/gdb"
	gredis "github.com/888go/goframe/database/gredis"
	"github.com/888go/goframe/frame/gins"
	"github.com/888go/goframe/i18n/gi18n"
	gclient "github.com/888go/goframe/net/gclient"
	ghttp "github.com/888go/goframe/net/ghttp"
	gtcp "github.com/888go/goframe/net/gtcp"
	gudp "github.com/888go/goframe/net/gudp"
	gcfg "github.com/888go/goframe/os/gcfg"
	glog "github.com/888go/goframe/os/glog"
	gres "github.com/888go/goframe/os/gres"
	gview "github.com/888go/goframe/os/gview"
	gvalid "github.com/888go/goframe/util/gvalid"
)

// X网页类 是一个便利函数，它创建并返回一个新的 HTTP 客户端。 md5:70c95e835d086085
func X网页类() *gclient.Client {
	return gclient.X创建()
}

// Http类 函数返回一个指定名称的 http 服务器实例。 md5:8613002877ca5b89
func Http类(名称 ...interface{}) *ghttp.X服务 {
	return gins.Server(名称...)
}

// TCP类 返回一个具有指定名称的TCP服务器实例。 md5:57eae18c70a9e8c2
func TCP类(名称 ...interface{}) *gtcp.Server {
	return gtcp.GetServer(名称...)
}

// UDP类 返回一个指定名称的UDP服务器实例。 md5:bffbc3c34f72d706
func UDP类(名称 ...interface{}) *gudp.Server {
	return gudp.GetServer(名称...)
}

// X模板类 返回指定名称的模板引擎对象实例。 md5:090a094475513c16
func X模板类(名称 ...string) *gview.View {
	return gins.View(名称...)
}

// X配置类返回具有指定名称的config对象实例。 md5:68a54cf8b0afa4f0
func X配置类(名称 ...string) *gcfg.Config {
	return gins.Config(名称...)
}

// Cfg别名 是 Config 的别名。
// 参见 Config。
// md5:1b7d381469f7e212
func Cfg别名(名称 ...string) *gcfg.Config {
	return X配置类(名称...)
}

// X资源类 返回一个 X资源类 类型的实例。
// 参数 `name` 为该实例的名称。
// md5:42e664c4b3a2bb54
func X资源类(名称 ...string) *gres.Resource {
	return gins.Resource(名称...)
}

// X多语言类 返回一个 gi18n.Manager 的实例。
// 参数 `name` 是实例的名称。
// md5:cb8fb8e2c93c597b
func X多语言类(名称 ...string) *gi18n.Manager {
	return gins.I18n(名称...)
}

// Res别名是Resource的别名。
// 请参阅Resource。
// md5:f2fc8f778d75b07c
func Res别名(名称 ...string) *gres.Resource {
	return X资源类(名称...)
}

// X日志类 函数返回一个 glog.Logger 实例。
// 参数 `name` 用于指定该实例的名称。
// md5:01b7637de189b37b
func X日志类(名称 ...string) *glog.Logger {
	return gins.Log(名称...)
}

// DB类 根据指定的配置组名称返回一个数据库ORM对象实例。 md5:229506bfec6bd334
func DB类(名称 ...string) gdb.DB {
	return gins.Database(名称...)
}

// DB类Model 根据默认数据库组的配置创建并返回一个模型。 md5:544c43fcc039bdfd
func DB类Model(表名或结构体 ...interface{}) *gdb.Model {
	return DB类().X创建Model对象(表名或结构体...)
}

// DB类原生SQL根据原始SQL（而非表）创建并返回一个模型。 md5:674c1e89391cbf02
func DB类原生SQL(原生Sql string, 参数 ...interface{}) *gdb.Model {
	return DB类().X原生SQL(原生Sql, 参数...)
}

// Redis类 根据给定的配置组名称返回一个redis客户端实例。 md5:3fab97150e392125
func Redis类(名称 ...string) *gredis.Redis {
	return gins.Redis(名称...)
}

// X效验类 是一个便利的函数，它创建并返回一个新的验证管理器对象。 md5:de213201906601b7
func X效验类() *gvalid.Validator {
	return gvalid.New()
}
