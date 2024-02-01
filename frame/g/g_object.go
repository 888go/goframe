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
func Client() *gclient.Client {
	return gclient.New()
}

// Server 返回一个具有指定名称的 http 服务器实例。
func Server(name ...interface{}) *ghttp.Server {
	return gins.Server(name...)
}

// TCPServer 返回一个具有指定名称的TCP服务器实例。
func TCPServer(name ...interface{}) *gtcp.Server {
	return gtcp.GetServer(name...)
}

// UDPServer 返回一个具有指定名称的 UDP 服务器实例。
func UDPServer(name ...interface{}) *gudp.Server {
	return gudp.GetServer(name...)
}

// View 返回指定名称的模板引擎对象实例。
func View(name ...string) *gview.View {
	return gins.View(name...)
}

// Config 返回一个具有指定名称的配置对象实例。
func Config(name ...string) *gcfg.Config {
	return gins.Config(name...)
}

// Cfg 是 Config 的别名。
// 请参阅 Config。
func Cfg(name ...string) *gcfg.Config {
	return Config(name...)
}

// Resource 返回一个 Resource 实例。
// 参数 `name` 是该实例的名称。
func Resource(name ...string) *gres.Resource {
	return gins.Resource(name...)
}

// I18n 返回一个 gi18n.Manager 的实例。
// 参数 `name` 是该实例的名称。
func I18n(name ...string) *gi18n.Manager {
	return gins.I18n(name...)
}

// Res 是 Resource 的别名。
// 请参阅 Resource。
func Res(name ...string) *gres.Resource {
	return Resource(name...)
}

// Log 返回一个 glog.Logger 的实例。
// 参数 `name` 是该实例的名称。
func Log(name ...string) *glog.Logger {
	return gins.Log(name...)
}

// DB 返回一个指定配置组名称的数据库 ORM 对象实例。
func DB(name ...string) gdb.DB {
	return gins.Database(name...)
}

// Model 根据默认数据库组的配置创建并返回一个模型。
func Model(tableNameOrStruct ...interface{}) *gdb.Model {
	return DB().Model(tableNameOrStruct...)
}

// ModelRaw 根据原始SQL（非表）创建并返回一个模型。
func ModelRaw(rawSql string, args ...interface{}) *gdb.Model {
	return DB().Raw(rawSql, args...)
}

// Redis 返回一个具有指定配置组名称的 redis 客户端实例。
func Redis(name ...string) *gredis.Redis {
	return gins.Redis(name...)
}

// Validator 是一个便捷函数，用于创建并返回一个新的验证管理器对象。
func Validator() *gvalid.Validator {
	return gvalid.New()
}
