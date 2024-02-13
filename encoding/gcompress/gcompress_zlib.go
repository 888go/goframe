// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包gcompress提供了多种压缩算法，适用于二进制/字节数据。
package 压缩类

import (
	"bytes"
	"compress/zlib"
	"io"
	
	"github.com/888go/goframe/errors/gerror"
)

// Zlib 使用zlib算法压缩`data`。
func Zlib压缩字节集(字节集 []byte) ([]byte, error) {
	if 字节集 == nil || len(字节集) < 13 {
		return 字节集, nil
	}
	var (
		err    error
		in     bytes.Buffer
		writer = zlib.NewWriter(&in)
	)

	if _, err = writer.Write(字节集); err != nil {
		err = 错误类.X多层错误并格式化(err, `zlib.Writer.Write failed`)
		return nil, err
	}
	if err = writer.Close(); err != nil {
		err = 错误类.X多层错误并格式化(err, `zlib.Writer.Close failed`)
		return in.Bytes(), err
	}
	return in.Bytes(), nil
}

// UnZlib 使用zlib算法对`data`进行解压缩。
func Zlib解压字节集(字节集 []byte) ([]byte, error) {
	if 字节集 == nil || len(字节集) < 13 {
		return 字节集, nil
	}
	var (
		out             bytes.Buffer
		bytesReader     = bytes.NewReader(字节集)
		zlibReader, err = zlib.NewReader(bytesReader)
	)
	if err != nil {
		err = 错误类.X多层错误并格式化(err, `zlib.NewReader failed`)
		return nil, err
	}
	if _, err = io.Copy(&out, zlibReader); err != nil {
		err = 错误类.X多层错误并格式化(err, `io.Copy failed`)
		return nil, err
	}
	return out.Bytes(), nil
}
