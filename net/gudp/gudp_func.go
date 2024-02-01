// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gudp
import (
	"net"
	
	"github.com/888go/goframe/errors/gerror"
	)
// NewNetConn 根据给定的地址创建并返回一个 *net.UDPConn 类型的实例。
func NewNetConn(remoteAddress string, localAddress ...string) (*net.UDPConn, error) {
	var (
		err        error
		remoteAddr *net.UDPAddr
		localAddr  *net.UDPAddr
		network    = `udp`
	)
	remoteAddr, err = net.ResolveUDPAddr(network, remoteAddress)
	if err != nil {
		return nil, gerror.Wrapf(
			err,
			`net.ResolveUDPAddr failed for network "%s", address "%s"`,
			network, remoteAddress,
		)
	}
	if len(localAddress) > 0 {
		localAddr, err = net.ResolveUDPAddr(network, localAddress[0])
		if err != nil {
			return nil, gerror.Wrapf(
				err,
				`net.ResolveUDPAddr failed for network "%s", address "%s"`,
				network, localAddress[0],
			)
		}
	}
	conn, err := net.DialUDP(network, localAddr, remoteAddr)
	if err != nil {
		return nil, gerror.Wrapf(
			err,
			`net.DialUDP failed for network "%s", local "%s", remote "%s"`,
			network, localAddr.String(), remoteAddr.String(),
		)
	}
	return conn, nil
}

// Send通过UDP连接将数据写入`address`，然后关闭连接。
// 注意，它适用于短连接场景。
func Send(address string, data []byte, retry ...Retry) error {
	conn, err := NewConn(address)
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.Send(data, retry...)
}

// SendRecv 使用UDP连接将数据写入`address`，读取响应后再关闭连接。
// 注意，它适用于短连接场景。
func SendRecv(address string, data []byte, receive int, retry ...Retry) ([]byte, error) {
	conn, err := NewConn(address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.SendRecv(data, receive, retry...)
}

// MustGetFreePort 的行为与 GetFreePort 相同，但是当出现任何错误时它会触发 panic。
func MustGetFreePort() (port int) {
	port, err := GetFreePort()
	if err != nil {
		panic(err)
	}
	return port
}

// GetFreePort 获取并返回一个可用的端口。
func GetFreePort() (port int, err error) {
	var (
		network = `udp`
		address = `:0`
	)
	resolvedAddr, err := net.ResolveUDPAddr(network, address)
	if err != nil {
		return 0, gerror.Wrapf(
			err,
			`net.ResolveUDPAddr failed for network "%s", address "%s"`,
			network, address,
		)
	}
	l, err := net.ListenUDP(network, resolvedAddr)
	if err != nil {
		return 0, gerror.Wrapf(
			err,
			`net.ListenUDP failed for network "%s", address "%s"`,
			network, resolvedAddr.String(),
		)
	}
	port = l.LocalAddr().(*net.UDPAddr).Port
	_ = l.Close()
	return
}

// GetFreePorts 获取并返回指定数量的空闲端口。
func GetFreePorts(count int) (ports []int, err error) {
	var (
		network = `udp`
		address = `:0`
	)
	for i := 0; i < count; i++ {
		resolvedAddr, err := net.ResolveUDPAddr(network, address)
		if err != nil {
			return nil, gerror.Wrapf(
				err,
				`net.ResolveUDPAddr failed for network "%s", address "%s"`,
				network, address,
			)
		}
		l, err := net.ListenUDP(network, resolvedAddr)
		if err != nil {
			return nil, gerror.Wrapf(
				err,
				`net.ListenUDP failed for network "%s", address "%s"`,
				network, resolvedAddr.String(),
			)
		}
		ports = append(ports, l.LocalAddr().(*net.UDPAddr).Port)
		_ = l.Close()
	}
	return ports, nil
}
