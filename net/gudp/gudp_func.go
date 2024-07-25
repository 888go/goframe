// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gudp

import (
	"net"

	"github.com/gogf/gf/v2/errors/gerror"
)

// NewNetConn 创建并返回一个具有给定地址的 *net.UDPConn。 md5:7327f361f04568ff
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

// Send 使用UDP连接向`address`发送数据，然后关闭连接。
// 注意，它用于短连接用途。 md5:3d373f3db04ae03d
func Send(address string, data []byte, retry ...Retry) error {
	conn, err := NewConn(address)
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.Send(data, retry...)
}

// SendRecv 使用UDP连接向`address`写入数据，读取响应后关闭连接。
// 注意，它用于短暂连接的场景。 md5:f88304194d59603d
func SendRecv(address string, data []byte, receive int, retry ...Retry) ([]byte, error) {
	conn, err := NewConn(address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.SendRecv(data, receive, retry...)
}

// MustGetFreePort 的行为与 GetFreePort 相似，但如果发生任何错误，则会引发恐慌。 md5:20b57c89fd162890
func MustGetFreePort() (port int) {
	port, err := GetFreePort()
	if err != nil {
		panic(err)
	}
	return port
}

// GetFreePort 获取并返回一个空闲的端口号。 md5:52dbf7a2d6e71da6
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

// GetFreePorts 获取并返回指定数量的空闲端口。 md5:ea99fb15b5bbc0fb
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
