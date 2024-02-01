// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包gcompress提供了多种压缩算法，适用于二进制/字节数据。
package gcompress
import (
	"bytes"
	"compress/zlib"
	"io"
	
	"github.com/888go/goframe/errors/gerror"
	)
// Zlib 使用zlib算法压缩`data`。
func Zlib(data []byte) ([]byte, error) {
	if data == nil || len(data) < 13 {
		return data, nil
	}
	var (
		err    error
		in     bytes.Buffer
		writer = zlib.NewWriter(&in)
	)

	if _, err = writer.Write(data); err != nil {
		err = gerror.Wrapf(err, `zlib.Writer.Write failed`)
		return nil, err
	}
	if err = writer.Close(); err != nil {
		err = gerror.Wrapf(err, `zlib.Writer.Close failed`)
		return in.Bytes(), err
	}
	return in.Bytes(), nil
}

// UnZlib 使用zlib算法对`data`进行解压缩。
func UnZlib(data []byte) ([]byte, error) {
	if data == nil || len(data) < 13 {
		return data, nil
	}
	var (
		out             bytes.Buffer
		bytesReader     = bytes.NewReader(data)
		zlibReader, err = zlib.NewReader(bytesReader)
	)
	if err != nil {
		err = gerror.Wrapf(err, `zlib.NewReader failed`)
		return nil, err
	}
	if _, err = io.Copy(&out, zlibReader); err != nil {
		err = gerror.Wrapf(err, `io.Copy failed`)
		return nil, err
	}
	return out.Bytes(), nil
}
