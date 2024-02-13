// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package g

import (
	"github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/database/gredis"
	"github.com/888go/goframe/frame/gins"
	"github.com/888go/goframe/i18n/gi18n"
	"github.com/888go/goframe/net/gclient"
	"github.com/888go/goframe/net/ghttp"
	"github.com/888go/goframe/net/gtcp"
	"github.com/888go/goframe/net/gudp"
	"github.com/888go/goframe/os/gcfg"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/os/gres"
	"github.com/888go/goframe/os/gview"
	"github.com/888go/goframe/util/gvalid"
)

// Client 是一个便捷函数，它会创建并返回一个新的 HTTP 客户端。
func X网页类() *网页类.Client {
	return 网页类.X创建()
}

// Server 返回一个具有指定名称的 http 服务器实例。
func Http类(名称 ...interface{}) *http类.Server {
	return gins.Server(名称...)
}

// TCPServer 返回一个具有指定名称的TCP服务器实例。
func TCP类(名称 ...interface{}) *tcp类.Server {
	return tcp类.GetServer(名称...)
}

// UDPServer 返回一个具有指定名称的 UDP 服务器实例。
func UDP类(名称 ...interface{}) *udp类.Server {
	return udp类.GetServer(名称...)
}

// View 返回指定名称的模板引擎对象实例。
func X模板类(名称 ...string) *模板类.View {
	return gins.View(名称...)
}

// Config 返回一个具有指定名称的配置对象实例。
func X配置类(名称 ...string) *配置类.Config {
	return gins.Config(名称...)
}

// Cfg 是 Config 的别名。
// 请参阅 Config。
func Cfg别名(名称 ...string) *配置类.Config {
	return X配置类(名称...)
}

// Resource 返回一个 Resource 实例。
// 参数 `name` 是该实例的名称。
func X资源类(名称 ...string) *资源类.Resource {
	return gins.Resource(名称...)
}

// I18n 返回一个 gi18n.Manager 的实例。
// 参数 `name` 是该实例的名称。
func X多语言类(名称 ...string) *gi18n.Manager {
	return gins.I18n(名称...)
}

// Res 是 Resource 的别名。
// 请参阅 Resource。
func Res别名(名称 ...string) *资源类.Resource {
	return X资源类(名称...)
}

// Log 返回一个 glog.Logger 的实例。
// 参数 `name` 是该实例的名称。
func X日志类(名称 ...string) *日志类.Logger {
	return gins.Log(名称...)
}

// DB 返回一个指定配置组名称的数据库 ORM 对象实例。
func DB类(名称 ...string) db类.DB {
	return gins.Database(名称...)
}

// Model 根据默认数据库组的配置创建并返回一个模型。
func DB类Model(表名或结构体 ...interface{}) *db类.Model {
	return DB类().X创建Model对象(表名或结构体...)
}

// ModelRaw 根据原始SQL（非表）创建并返回一个模型。
func DB类原生SQL(原生Sql string, 参数 ...interface{}) *db类.Model {
	return DB类().X原生SQL(原生Sql, 参数...)
}

// Redis 返回一个具有指定配置组名称的 redis 客户端实例。
func Redis类(名称 ...string) *redis类.Redis {
	return gins.Redis(名称...)
}

// Validator 是一个便捷函数，用于创建并返回一个新的验证管理器对象。
func X效验类() *效验类.Validator {
	return 效验类.New()
}
