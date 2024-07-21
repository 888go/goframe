// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gredis

import (
	"context"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/internal/intlog"
)

var (
	// localInstances 用于管理redis客户端的实例。 md5:3e8c8ec6661fd69a
	localInstances = gmap.NewStrAnyMap(true)
)

// Instance 返回指定分组的 redis 客户端实例。
// 如果未传入 `name` 参数，它将返回一个使用默认配置分组的 redis 实例。
// md5:22f72aaad455f11e
func Instance(name ...string) *Redis {
	group := DefaultGroupName
	if len(name) > 0 && name[0] != "" {
		group = name[0]
	}
	v := localInstances.GetOrSetFuncLock(group, func() interface{} {
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
