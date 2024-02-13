// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package tcp类

import (
	"encoding/binary"
	"time"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
)

const (
	pkgHeaderSizeDefault = 2 // 简单包协议的头部大小
	pkgHeaderSizeMax     = 4 // 简单包协议的最大头部大小
)

// PkgOption 是用于简单协议的包选项。
type PkgOption struct {
// HeaderSize 用于标记下一段数据的长度，用于接收数据时进行判断。
// 默认情况下占2字节，最大为4字节，可表示的数据长度范围从65535字节到4294967295字节。
	HeaderSize int

// MaxDataSize 是用于数据长度验证的数据字段大小，单位为字节。
// 若未手动设置，它将根据 HeaderSize 自动进行相应的设置。
	MaxDataSize int

	// 当操作失败时的重试策略。
	Retry Retry
}

// SendPkg send data using simple package protocol.
//
// Simple package protocol: DataLength(24bit)|DataField(variant)。
//
// Note that,
// 1. The DataLength is the length of DataField, which does not contain the header size.
// 2. The integer bytes of the package are encoded using BigEndian order.
func (c *Conn) SendPkg(data []byte, option ...PkgOption) error {
	pkgOption, err := getPkgOption(option...)
	if err != nil {
		return err
	}
	length := len(data)
	if length > pkgOption.MaxDataSize {
		return 错误类.X创建错误码并格式化(
			错误码类.CodeInvalidParameter,
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

// SendPkgWithTimeout 使用简单的包协议并设置超时，向连接写入数据。
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

// SendRecvPkg 使用简单的包协议将数据写入连接，并阻塞等待读取响应。
func (c *Conn) SendRecvPkg(data []byte, option ...PkgOption) ([]byte, error) {
	if err := c.SendPkg(data, option...); err == nil {
		return c.RecvPkg(option...)
	} else {
		return nil, err
	}
}

// SendRecvPkgWithTimeout 使用简单包协议，以超时机制向连接写入数据并读取响应。
func (c *Conn) SendRecvPkgWithTimeout(data []byte, timeout time.Duration, option ...PkgOption) ([]byte, error) {
	if err := c.SendPkg(data, option...); err == nil {
		return c.RecvPkgWithTimeout(timeout, option...)
	} else {
		return nil, err
	}
}

// RecvPkg 通过简单的包协议从连接中接收数据。
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
		// 如果头部大小小于4字节（uint32），则用零填充。
		length = int(binary.BigEndian.Uint32([]byte{0, 0, 0, buffer[0]}))
	case 2:
		length = int(binary.BigEndian.Uint32([]byte{0, 0, buffer[0], buffer[1]}))
	case 3:
		length = int(binary.BigEndian.Uint32([]byte{0, buffer[0], buffer[1], buffer[2]}))
	default:
		length = int(binary.BigEndian.Uint32([]byte{buffer[0], buffer[1], buffer[2], buffer[3]}))
	}
// 在这里，它验证包的大小。
// 如果验证失败，它会立即清除缓冲区并返回错误。
	if length < 0 || length > pkgOption.MaxDataSize {
		return nil, 错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, `invalid package size %d`, length)
	}
	// Empty package.
	if length == 0 {
		return nil, nil
	}
	// Data field.
	return c.Recv(length, pkgOption.Retry)
}

// RecvPkgWithTimeout 使用简单包协议，以超时方式从连接中读取数据。
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
// 如果未提供选项，则返回一个具有默认值的新选项。
func getPkgOption(option ...PkgOption) (*PkgOption, error) {
	pkgOption := PkgOption{}
	if len(option) > 0 {
		pkgOption = option[0]
	}
	if pkgOption.HeaderSize == 0 {
		pkgOption.HeaderSize = pkgHeaderSizeDefault
	}
	if pkgOption.HeaderSize > pkgHeaderSizeMax {
		return nil, 错误类.X创建错误码并格式化(
			错误码类.CodeInvalidParameter,
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
			// math.MaxInt32 不是 math.MaxUint32
// （这段注释是在强调或纠正某个变量或函数的使用，指出这里使用的是32位整数的最大值常量 `math.MaxInt32`，而不是32位无符号整数的最大值常量 `math.MaxUint32`。）
			pkgOption.MaxDataSize = 0x7FFFFFFF
		}
	}
	if pkgOption.MaxDataSize > 0x7FFFFFFF {
		return nil, 错误类.X创建错误码并格式化(
			错误码类.CodeInvalidParameter,
			`package data size %d definition exceeds allowed max data size %d`,
			pkgOption.MaxDataSize, 0x7FFFFFFF,
		)
	}
	return &pkgOption, nil
}
