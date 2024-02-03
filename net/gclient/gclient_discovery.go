// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gclient

import (
	"context"
	"net/http"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/net/gsel"
	"github.com/888go/goframe/net/gsvc"
)

type discoveryNode struct {
	service gsvc.Service
	address string
}

// Service 是客户端发现服务。
func (n *discoveryNode) Service() gsvc.Service {
	return n.service
}

// Address 返回节点的地址。
func (n *discoveryNode) Address() string {
	return n.address
}

var clientSelectorMap = gmap.New(true)

// internalMiddlewareDiscovery 是一个客户端中间件，用于为客户端启用服务发现功能。
func internalMiddlewareDiscovery(c *Client, r *http.Request) (response *Response, err error) {
	if c.discovery == nil {
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
		return nil, err
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
			// 更新选择器节点。
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
	// 从多个地址中选择一个节点。
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
