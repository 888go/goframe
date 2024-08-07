// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"context"
	"fmt"

	gipv4 "github.com/888go/goframe/net/gipv4"
	"github.com/888go/goframe/net/gsvc"
	gctx "github.com/888go/goframe/os/gctx"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

// doServiceRegister 将当前服务注册到Registry。 md5:fbfe784fd5718e65
func (s *X服务) doServiceRegister() {
	if s.registrar == nil {
		return
	}
	s.serviceMu.Lock()
	defer s.serviceMu.Unlock()
	var (
		ctx      = gctx.X取初始化上下文()
		protocol = gsvc.DefaultProtocol
		insecure = true
		err      error
	)
	if s.config.TLSConfig != nil {
		protocol = `https`
		insecure = false
	}
	metadata := gsvc.Metadata{
		gsvc.MDProtocol: protocol,
		gsvc.MDInsecure: insecure,
	}
	s.service = &gsvc.LocalService{
		Name:      s.X取服务名称(),
		Endpoints: s.calculateListenedEndpoints(ctx),
		Metadata:  metadata,
	}
	s.Logger别名().X输出并格式化DEBU(ctx, `service register: %+v`, s.service)
	if len(s.service.GetEndpoints()) == 0 {
		s.Logger别名().X输出并格式化WARN(ctx, `no endpoints found to register service, abort service registering`)
		return
	}
	if s.service, err = s.registrar.Register(ctx, s.service); err != nil {
		s.Logger别名().X输出并格式化FATA(ctx, `%+v`, err)
	}
}

// doServiceDeregister 从注册表中注销当前服务。 md5:bc984d9e06c1ce87
func (s *X服务) doServiceDeregister() {
	if s.registrar == nil {
		return
	}
	s.serviceMu.Lock()
	defer s.serviceMu.Unlock()
	if s.service == nil {
		return
	}
	var ctx = gctx.X取初始化上下文()
	s.Logger别名().X输出并格式化DEBU(ctx, `service deregister: %+v`, s.service)
	if err := s.registrar.Deregister(ctx, s.service); err != nil {
		s.Logger别名().X输出并格式化ERR(ctx, `%+v`, err)
	}
	s.service = nil
}

func (s *X服务) calculateListenedEndpoints(ctx context.Context) gsvc.Endpoints {
	var (
		configAddr = s.config.Address
		endpoints  = make(gsvc.Endpoints, 0)
		addresses  = s.config.Endpoints
	)
	if configAddr == "" {
		configAddr = s.config.HTTPSAddr
	}
	if len(addresses) == 0 {
		addresses = gstr.X分割并忽略空值(configAddr, ",")
	}
	for _, address := range addresses {
		var (
			addrArray     = gstr.X分割(address, ":")
			listenedIps   []string
			listenedPorts []int
		)
		if len(addrArray) == 1 {
			addrArray = append(addrArray, gconv.String(defaultEndpointPort))
		}
		// IPs.
		switch addrArray[0] {
		case "127.0.0.1":
			// Nothing to do.
		case "0.0.0.0", "":
			intranetIps, err := gipv4.GetIntranetIpArray()
			if err != nil {
				s.Logger别名().X输出并格式化ERR(ctx, `error retrieving intranet ip: %+v`, err)
				return nil
			}
			// 如果找不到内网IP，它将使用能够获取到的所有IP，这可能包括公网IP。
			// md5:17448379a248fd2b
			if len(intranetIps) == 0 {
				allIps, err := gipv4.GetIpArray()
				if err != nil {
					s.Logger别名().X输出并格式化ERR(ctx, `error retrieving ip from current node: %+v`, err)
					return nil
				}
				s.Logger别名().X输出并格式化NOTI(
					ctx,
					`no intranet ip found, using internet ip to register service: %v`,
					allIps,
				)
				listenedIps = allIps
				break
			}
			listenedIps = intranetIps
		default:
			listenedIps = []string{addrArray[0]}
		}
		// Ports.
		switch addrArray[1] {
		case "0":
			listenedPorts = s.X取所有监听已端口()
		default:
			listenedPorts = []int{gconv.X取整数(addrArray[1])}
		}
		for _, ip := range listenedIps {
			for _, port := range listenedPorts {
				endpoints = append(endpoints, gsvc.NewEndpoint(fmt.Sprintf(`%s:%d`, ip, port)))
			}
		}
	}
	return endpoints
}
