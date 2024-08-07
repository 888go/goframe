// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// gcompress 包提供了针对二进制/字节数据的各种压缩算法。 md5:eb6f1752b091375d
package 压缩类

import (
	"bytes"
	"compress/zlib"
	"io"

	gerror "github.com/888go/goframe/errors/gerror"
)

// Zlib压缩字节集使用zlib算法对`data`进行压缩。 md5:ed5cf5943e81e6a3
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
		err = gerror.X多层错误并格式化(err, `zlib.Writer.Write failed`)
		return nil, err
	}
	if err = writer.Close(); err != nil {
		err = gerror.X多层错误并格式化(err, `zlib.Writer.Close failed`)
		return in.Bytes(), err
	}
	return in.Bytes(), nil
}

// Zlib解压字节集使用zlib算法对`data`进行解压缩。 md5:e5713bb3c9724494
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
		err = gerror.X多层错误并格式化(err, `zlib.NewReader failed`)
		return nil, err
	}
	if _, err = io.Copy(&out, zlibReader); err != nil {
		err = gerror.X多层错误并格式化(err, `io.Copy failed`)
		return nil, err
	}
	return out.Bytes(), nil
}
