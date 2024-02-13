// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package redis类

import (
	"context"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/internal/intlog"
)

var (
	// localInstances 用于 Redis 客户端实例的管理。
	localInstances = map类.X创建StrAny(true)
)

// Instance 返回指定组的 Redis 客户端实例。
// `name` 参数不是必需的，如果未传递 `name`，
// 则返回一个使用默认配置组的 Redis 实例。
func Instance(name ...string) *Redis {
	group := DefaultGroupName
	if len(name) > 0 && name[0] != "" {
		group = name[0]
	}
	v := localInstances.X取值或设置值_函数带锁(group, func() interface{} {
		if config, ok := GetConfig(group); ok {
			r, err := New(config)
			if err != nil {
				intlog.Errorf(context.TODO(), `%+v`, err)
				return nil
			}
			return r
		}
		return nil
	})
	if v != nil {
		return v.(*Redis)
	}
	return nil
}
