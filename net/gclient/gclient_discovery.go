// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gclient

import (
	"context"
	"net/http"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/net/gsel"
	"github.com/gogf/gf/v2/net/gsvc"
	"github.com/gogf/gf/v2/text/gstr"
)

type discoveryNode struct {
	service gsvc.Service
	address string
}

// Service是服务发现客户端。. md5:0022814c173cde3d
func (n *discoveryNode) Service() gsvc.Service {
	return n.service
}

// Address 返回节点的地址。. md5:299d6002549d1c50
func (n *discoveryNode) Address() string {
	return n.address
}

// 将服务前缀转换并存储到其选择器映射缓存中。. md5:96fd7ebfce9abb4a
var clientSelectorMap = gmap.New(true)

// internalMiddlewareDiscovery 是一个客户端中间件，它为客户端启用服务发现功能。. md5:d465f663ac6b0f25
func internalMiddlewareDiscovery(c *Client, r *http.Request) (response *Response, err error) {
	if c.discovery == nil && !isServiceName(r.URL.Host) {
		return c.Next(r)
	}
	var (
		ctx     = r.Context()
		service gsvc.Service
	)
	service, err = gsvc.GetAndWatchWithDiscovery(ctx, c.discovery, r.URL.Host, func(service gsvc.Service) {
		intlog.Printf(ctx, `http client watching service "%s" changed`, service.GetPrefix())
		if v := clientSelectorMap.Get(service.GetPrefix()); v != nil {
			if err = updateSelectorNodesByService(ctx, v.(gsel.Selector), service); err != nil {
				intlog.Errorf(ctx, `%+v`, err)
			}
		}
	})
	if err != nil {
		if gerror.Code(err) == gcode.CodeNotFound {
			intlog.Printf(
				ctx,
				`service discovery error with url "%s:%s":%s`,
				r.Method, r.URL.String(), err.Error(),
			)
			return c.Next(r)
		}
	}
	if service == nil {
		return c.Next(r)
	}
	// Balancer.
	var (
		selectorMapKey   = service.GetPrefix()
		selectorMapValue = clientSelectorMap.GetOrSetFuncLock(selectorMapKey, func() interface{} {
			intlog.Printf(ctx, `http client create selector for service "%s"`, selectorMapKey)
			selector := c.builder.Build()
			// Update selector nodes.
			if err = updateSelectorNodesByService(ctx, selector, service); err != nil {
				return nil
			}
			return selector
		})
	)
	if err != nil {
		return nil, err
	}
	selector := selectorMapValue.(gsel.Selector)
	// 从多个地址中选择一个节点。. md5:088c30a06a0dc6a9
	node, done, err := selector.Pick(ctx)
	if err != nil {
		return nil, err
	}
	if done != nil {
		defer done(ctx, gsel.DoneInfo{})
	}
	r.Host = node.Address()
	r.URL.Host = node.Address()
	return c.Next(r)
}

func updateSelectorNodesByService(ctx context.Context, selector gsel.Selector, service gsvc.Service) error {
	nodes := make(gsel.Nodes, 0)
	for _, endpoint := range service.GetEndpoints() {
		nodes = append(nodes, &discoveryNode{
			service: service,
			address: endpoint.String(),
		})
	}
	return selector.Update(ctx, nodes)
}

// isServiceName 检查并返回给定的输入参数是否为服务名称。
// 它通过判断参数中是否包含端口分隔符':'来检查。
//
// 如果使用服务发现，则它不包含任何端口号。
// md5:fefe525183a2ba4d
func isServiceName(serviceNameOrAddress string) bool {
	return !gstr.Contains(serviceNameOrAddress, gsvc.EndpointHostPortDelimiter)
}
