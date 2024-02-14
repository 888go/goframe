// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"context"
	"fmt"
	
	"github.com/888go/goframe/net/gipv4"
	"github.com/888go/goframe/net/gsvc"
	"github.com/888go/goframe/os/gctx"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
)

// doServiceRegister 将当前服务注册到注册中心。
func (s *X服务) doServiceRegister() {
	if s.registrar == nil {
		return
	}
	s.serviceMu.Lock()
	defer s.serviceMu.Unlock()
	var (
		ctx      = 上下文类.X取初始化上下文()
		protocol = gsvc.DefaultProtocol
		insecure = true
		err      error
	)
	if s.config.TLS配置 != nil {
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

// doServiceDeregister 从注册中心注销当前服务。
func (s *X服务) doServiceDeregister() {
	if s.registrar == nil {
		return
	}
	s.serviceMu.Lock()
	defer s.serviceMu.Unlock()
	if s.service == nil {
		return
	}
	var ctx = 上下文类.X取初始化上下文()
	s.Logger别名().X输出并格式化DEBU(ctx, `service deregister: %+v`, s.service)
	if err := s.registrar.Deregister(ctx, s.service); err != nil {
		s.Logger别名().X输出并格式化ERR(ctx, `%+v`, err)
	}
	s.service = nil
}

func (s *X服务) calculateListenedEndpoints(ctx context.Context) gsvc.Endpoints {
	var (
		configAddr = s.config.X监听地址
		endpoints  = make(gsvc.Endpoints, 0)
		addresses  = s.config.Endpoints
	)
	if configAddr == "" {
		configAddr = s.config.HTTPS监听地址
	}
	if len(addresses) == 0 {
		addresses = 文本类.X分割并忽略空值(configAddr, ",")
	}
	for _, address := range addresses {
		var (
			addrArray     = 文本类.X分割(address, ":")
			listenedIps   []string
			listenedPorts []int
		)
		if len(addrArray) == 1 {
			addrArray = append(addrArray, 转换类.String(defaultEndpointPort))
		}
		// IPs.
		switch addrArray[0] {
		case "127.0.0.1":
			// Nothing to do.
		case "0.0.0.0", "":
			intranetIps, err := ipv4类.GetIntranetIpArray()
			if err != nil {
				s.Logger别名().X输出并格式化ERR(ctx, `error retrieving intranet ip: %+v`, err)
				return nil
			}
// 如果没有找到内网IP，它将使用所有能够获取到的IP，这其中可能包括公网IP。
			if len(intranetIps) == 0 {
				allIps, err := ipv4类.GetIpArray()
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
			listenedPorts = []int{转换类.X取整数(addrArray[1])}
		}
		for _, ip := range listenedIps {
			for _, port := range listenedPorts {
				endpoints = append(endpoints, gsvc.NewEndpoint(fmt.Sprintf(`%s:%d`, ip, port)))
			}
		}
	}
	return endpoints
}
