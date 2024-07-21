// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtcp

import (
	"encoding/binary"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

const (
	_                    = iota << 1
	pkgHeaderSizeDefault // 简单包协议的头部大小。 md5:6e3294120717ff1f
	pkgHeaderSizeMax     // 简单包协议的最大头部大小。 md5:b1e4f447dc182fa5
)

// PkgOption是简单协议的包选项。 md5:f49ee7556a39be3e
type PkgOption struct {
// HeaderSize 用于标记接下来接收数据的长度。
// 它默认为2字节，最大为4字节，代表数据的最大长度可以从65535字节到4294967295字节。
// md5:cc02a98c94fddd76
	HeaderSize int

// MaxDataSize 是数据字段的字节大小，用于数据长度验证。
// 如果未手动设置，它将根据HeaderSize自动相应设置。
// md5:a47982162ce5ef01
	MaxDataSize int

	// 操作失败时的重试策略。 md5:cd672b1b96adbbdd
	Retry Retry
}

// SendPkg 使用简单包协议发送数据。
//
// 简单包协议：DataLength(24位)|DataField(variant)。
//
// 注意，
// 1. DataLength 是 DataField 的长度，不包含头大小。
// 2. 包的整数字节使用大端序编码。
// md5:daa39f4e32227d79
// ff:
// c:
// data:
// option:
func (c *Conn) SendPkg(data []byte, option ...PkgOption) error {
	pkgOption, err := getPkgOption(option...)
	if err != nil {
		return err
	}
	length := len(data)
	if length > pkgOption.MaxDataSize {
		return gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`data too long, data size %d exceeds allowed max data size %d`,
			length, pkgOption.MaxDataSize,
		)
	}
	offset := pkgHeaderSizeMax - pkgOption.HeaderSize
	buffer := make([]byte, pkgHeaderSizeMax+len(data))
	binary.BigEndian.PutUint32(buffer[0:], uint32(length))
	copy(buffer[pkgHeaderSizeMax:], data)
	if pkgOption.Retry.Count > 0 {
		return c.Send(buffer[offset:], pkgOption.Retry)
	}
	return c.Send(buffer[offset:])
}

// 使用简单包协议带超时时间地向连接发送数据。 md5:3f89f6011aed63bc
// ff:
// c:
// data:
// timeout:
// option:
// err:
func (c *Conn) SendPkgWithTimeout(data []byte, timeout time.Duration, option ...PkgOption) (err error) {
	if err := c.SetDeadlineSend(time.Now().Add(timeout)); err != nil {
		return err
	}
	defer func() {
		_ = c.SetDeadlineSend(time.Time{})
	}()
	err = c.SendPkg(data, option...)
	return
}

// SendRecvPkg 使用简单的包协议将数据写入连接，并阻塞读取响应。 md5:c157760431f11896
// ff:
// c:
// data:
// option:
func (c *Conn) SendRecvPkg(data []byte, option ...PkgOption) ([]byte, error) {
	if err := c.SendPkg(data, option...); err == nil {
		return c.RecvPkg(option...)
	} else {
		return nil, err
	}
}

// SendRecvPkgWithTimeout 使用简单包协议向连接写入数据，并在超时后读取响应。 md5:6da9109d534f7729
// ff:
// c:
// data:
// timeout:
// option:
func (c *Conn) SendRecvPkgWithTimeout(data []byte, timeout time.Duration, option ...PkgOption) ([]byte, error) {
	if err := c.SendPkg(data, option...); err == nil {
		return c.RecvPkgWithTimeout(timeout, option...)
	} else {
		return nil, err
	}
}

// RecvPkg 使用简单包协议从连接接收数据。 md5:cf1329c5df27539a
// ff:
// c:
// option:
// result:
// err:
func (c *Conn) RecvPkg(option ...PkgOption) (result []byte, err error) {
	var (
		buffer []byte
		length int
	)
	pkgOption, err := getPkgOption(option...)
	if err != nil {
		return nil, err
	}
	// Header field.
	buffer, err = c.Recv(pkgOption.HeaderSize, pkgOption.Retry)
	if err != nil {
		return nil, err
	}
	switch pkgOption.HeaderSize {
	case 1:
		// 如果头部大小小于4字节（uint32），则用零填充。 md5:5e9e147401796703
		length = int(binary.BigEndian.Uint32([]byte{0, 0, 0, buffer[0]}))
	case 2:
		length = int(binary.BigEndian.Uint32([]byte{0, 0, buffer[0], buffer[1]}))
	case 3:
		length = int(binary.BigEndian.Uint32([]byte{0, buffer[0], buffer[1], buffer[2]}))
	default:
		length = int(binary.BigEndian.Uint32([]byte{buffer[0], buffer[1], buffer[2], buffer[3]}))
	}
	// 此处校验包的大小。
	// 如果校验失败，会立即清空缓冲区并返回错误。
	// md5:0871405b30986628
	if length < 0 || length > pkgOption.MaxDataSize {
		return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `invalid package size %d`, length)
	}
	// Empty package.
	if length == 0 {
		return nil, nil
	}
	// Data field.
	return c.Recv(length, pkgOption.Retry)
}

// RecvPkgWithTimeout 使用简单包协议，从连接中读取数据，同时设置超时。 md5:5e1d4882f4476862
// ff:
// c:
// timeout:
// option:
// data:
// err:
func (c *Conn) RecvPkgWithTimeout(timeout time.Duration, option ...PkgOption) (data []byte, err error) {
	if err = c.SetDeadlineRecv(time.Now().Add(timeout)); err != nil {
		return nil, err
	}
	defer func() {
		_ = c.SetDeadlineRecv(time.Time{})
	}()
	data, err = c.RecvPkg(option...)
	return
}

// getPkgOption 包装并返回 PkgOption。
// 如果没有提供选项，则返回一个具有默认值的新选项。
// md5:752809cff379479d
func getPkgOption(option ...PkgOption) (*PkgOption, error) {
	pkgOption := PkgOption{}
	if len(option) > 0 {
		pkgOption = option[0]
	}
	if pkgOption.HeaderSize == 0 {
		pkgOption.HeaderSize = pkgHeaderSizeDefault
	}
	if pkgOption.HeaderSize > pkgHeaderSizeMax {
		return nil, gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`package header size %d definition exceeds max header size %d`,
			pkgOption.HeaderSize, pkgHeaderSizeMax,
		)
	}
	if pkgOption.MaxDataSize == 0 {
		switch pkgOption.HeaderSize {
		case 1:
			pkgOption.MaxDataSize = 0xFF
		case 2:
			pkgOption.MaxDataSize = 0xFFFF
		case 3:
			pkgOption.MaxDataSize = 0xFFFFFF
		case 4:
			// math.MaxInt32 而不是 math.MaxUint32. md5:11ed9a0830ca2d39
			pkgOption.MaxDataSize = 0x7FFFFFFF
		}
	}
	if pkgOption.MaxDataSize > 0x7FFFFFFF {
		return nil, gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`package data size %d definition exceeds allowed max data size %d`,
			pkgOption.MaxDataSize, 0x7FFFFFFF,
		)
	}
	return &pkgOption, nil
}
