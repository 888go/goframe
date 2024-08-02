// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 配置类

import (
	"context"
)

// Adapter是用于获取配置的接口。 md5:5c3d613bea87d056
type Adapter interface {
	// 可用性检查并返回后端配置服务是否可用。
	// 可选参数 `resource` 指定特定的配置资源。
	// 
	// 请注意，此函数不会返回错误，因为它只是简单地检查后端配置服务。
	// md5:8c240c72c0849cd7
	Available(ctx context.Context, resource ...string) (ok bool)

	// Get 通过在当前资源中指定的`pattern`获取并返回值。
	// 模式示例：
	// 对于映射项，使用 "x.y.z"。
	// 对于切片项，使用 "x.0.y"。
	// md5:821429a92b84150c
	Get(ctx context.Context, pattern string) (value interface{}, err error)

	// Data 获取并返回当前资源中的所有配置数据作为映射。
	// 请注意，如果配置数据过大，此函数可能导致大量内存使用。
	// 如果需要，你可以自行实现这个函数。
	// md5:7eaedd1a7f099a23
	Data(ctx context.Context) (data map[string]interface{}, err error)
}
