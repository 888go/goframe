// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gsvc

import (
	"context"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
)

// Register 将 `服务` 注册到默认注册表中。 md5:07ed2154ce52f6df
func Register(ctx context.Context, service Service) (Service, error) {
	if defaultRegistry == nil {
		return nil, gerror.NewCodef(gcode.CodeNotImplemented, `no Registry is registered`)
	}
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	return defaultRegistry.Register(ctx, service)
}

// Deregister 将 `service` 从默认注册表中移除。 md5:21de7624550ef4ed
func Deregister(ctx context.Context, service Service) error {
	if defaultRegistry == nil {
		return gerror.NewCodef(gcode.CodeNotImplemented, `no Registry is registered`)
	}
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	return defaultRegistry.Deregister(ctx, service)
}
