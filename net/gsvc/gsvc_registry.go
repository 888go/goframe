// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gsvc
import (
	"context"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	)
// Register 将`service`注册到默认注册中心。
func Register(ctx context.Context, service Service) (Service, error) {
	if defaultRegistry == nil {
		return nil, gerror.NewCodef(gcode.CodeNotImplemented, `no Registry is registered`)
	}
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	return defaultRegistry.Register(ctx, service)
}

// Deregister 从默认注册表中移除`service`。
func Deregister(ctx context.Context, service Service) error {
	if defaultRegistry == nil {
		return gerror.NewCodef(gcode.CodeNotImplemented, `no Registry is registered`)
	}
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	return defaultRegistry.Deregister(ctx, service)
}
