// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcfg
import (
	"context"
	)
// Adapter 是用于获取配置的接口。
type Adapter interface {
// Available 检查并返回配置服务是否可用。
// 可选参数 `resource` 用于指定特定的配置资源。
//
// 注意，此函数不会返回错误，因为它只是简单地检查后端配置服务。
	Available(ctx context.Context, resource ...string) (ok bool)

// Get 通过指定的 `pattern` 从当前资源中获取并返回值。
// 模式示例：
// "x.y.z" 用于获取映射项。
// "x.0.y" 用于获取切片项。
	Get(ctx context.Context, pattern string) (value interface{}, err error)

// Data 函数从当前资源中获取并返回所有的配置数据，以 map 的形式。
// 注意，如果配置数据过大，此函数可能会导致大量内存使用，
// 如有必要，您可以自行实现这个函数。
	Data(ctx context.Context) (data map[string]interface{}, err error)
}
