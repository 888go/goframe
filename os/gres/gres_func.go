// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gres

import (
	"archive/zip"
	"bytes"
	"encoding/hex"
	"fmt"
	
	"github.com/888go/goframe/encoding/gbase64"
	"github.com/888go/goframe/encoding/gcompress"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/text/gstr"
)

const (
	packedGoSourceTemplate = `
package %s

import "github.com/888go/goframe/os/gres"

func init() {
	if err := gres.Add("%s"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
`
)

// Option 包含 Pack 函数的额外选项。
type Option struct {
	Prefix   string // 在资源管理器中，每个文件项的文件路径前缀。
	KeepPath bool   // 在打包时保留传递的路径，通常用于相对路径。
}

// Pack 将由 `srcPaths` 指定的路径打包成字节形式。
// 不必要的参数 `keyPrefix` 表示每个文件被打包到结果字节中时的前缀。
//
// 注意，参数 `srcPaths` 支持使用 ',' 连接的多个路径。
//
// 已弃用：请改用 PackWithOption。
func Pack(srcPaths string, keyPrefix ...string) ([]byte, error) {
	option := Option{}
	if len(keyPrefix) > 0 && keyPrefix[0] != "" {
		option.Prefix = keyPrefix[0]
	}
	return PackWithOption(srcPaths, option)
}

// PackWithOption 函数将由 `srcPaths` 指定的路径打包成字节形式。
//
// 注意，参数 `srcPaths` 支持使用 ',' 连接的多个路径。
func PackWithOption(srcPaths string, option Option) ([]byte, error) {
	var buffer = bytes.NewBuffer(nil)
	err := zipPathWriter(srcPaths, buffer, option)
	if err != nil {
		return nil, err
	}
	// 使用Gzip压缩数据字节以减少其大小。
	return gcompress.Gzip(buffer.Bytes(), 9)
}

// PackToFile 将由`srcPaths`指定的路径打包到目标文件`dstPath`中。
// 不必要的参数`keyPrefix`表示每个被打包到结果字节中的文件前缀。
//
// 注意，参数`srcPaths`支持使用','连接的多个路径。
//
// 已弃用：请改用PackToFileWithOption。
func PackToFile(srcPaths, dstPath string, keyPrefix ...string) error {
	data, err := Pack(srcPaths, keyPrefix...)
	if err != nil {
		return err
	}
	return gfile.PutBytes(dstPath, data)
}

// PackToFileWithOption 将由 `srcPaths` 指定的路径打包到目标文件 `dstPath`。
//
// 注意，参数 `srcPaths` 支持通过 ',' 连接的多个路径。
func PackToFileWithOption(srcPaths, dstPath string, option Option) error {
	data, err := PackWithOption(srcPaths, option)
	if err != nil {
		return err
	}
	return gfile.PutBytes(dstPath, data)
}

// PackToGoFile 将由 `srcPaths` 指定的路径打包到目标 Go 文件 `goFilePath`，
// 并使用给定的包名 `pkgName`。
//
// 非必需参数 `keyPrefix` 表示每个打包进结果字节流文件的前缀。
//
// 注意，参数 `srcPaths` 支持以 ',' 连接的多个路径。
//
// 已弃用：请改用 PackToGoFileWithOption。
func PackToGoFile(srcPath, goFilePath, pkgName string, keyPrefix ...string) error {
	data, err := Pack(srcPath, keyPrefix...)
	if err != nil {
		return err
	}
	return gfile.PutContents(
		goFilePath,
		fmt.Sprintf(gstr.TrimLeft(packedGoSourceTemplate), pkgName, gbase64.EncodeToString(data)),
	)
}

// PackToGoFileWithOption 根据指定的 `srcPaths` 路径将文件打包到目标 Go 文件 `goFilePath`，
// 同时使用给定的包名 `pkgName`。
//
// 注意，参数 `srcPaths` 支持通过 ',' 连接多个路径。
func PackToGoFileWithOption(srcPath, goFilePath, pkgName string, option Option) error {
	data, err := PackWithOption(srcPath, option)
	if err != nil {
		return err
	}
	return gfile.PutContents(
		goFilePath,
		fmt.Sprintf(gstr.TrimLeft(packedGoSourceTemplate), pkgName, gbase64.EncodeToString(data)),
	)
}

// Unpack 将由 `path` 指定的内容解包为 []*File 类型的切片。
func Unpack(path string) ([]*File, error) {
	realPath, err := gfile.Search(path)
	if err != nil {
		return nil, err
	}
	return UnpackContent(gfile.GetContents(realPath))
}

// UnpackContent 解析内容到 []*File 类型的切片。
func UnpackContent(content string) ([]*File, error) {
	var (
		err  error
		data []byte
	)
	if isHexStr(content) {
// 这里保留了使用十六进制字符串进行旧版本打包字符串的兼容性。
// TODO：未来将删除对此种支持。
		data, err = gcompress.UnGzip(hexStrToBytes(content))
		if err != nil {
			return nil, err
		}
	} else if isBase64(content) {
		// 新版本使用base64对字符串进行打包。
		b, err := gbase64.DecodeString(content)
		if err != nil {
			return nil, err
		}
		data, err = gcompress.UnGzip(b)
		if err != nil {
			return nil, err
		}
	} else {
		data, err = gcompress.UnGzip([]byte(content))
		if err != nil {
			return nil, err
		}
	}
	reader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		err = gerror.Wrapf(err, `create zip reader failed`)
		return nil, err
	}
	array := make([]*File, len(reader.File))
	for i, file := range reader.File {
		array[i] = &File{file: file}
	}
	return array, nil
}

// isBase64 检查并返回给定内容 `s` 是否为 base64 字符串。
// 如果 `s` 是 base64 字符串，则返回 true；否则返回 false。
func isBase64(s string) bool {
	var r bool
	for i := 0; i < len(s); i++ {
		r = (s[i] >= '0' && s[i] <= '9') ||
			(s[i] >= 'a' && s[i] <= 'z') ||
			(s[i] >= 'A' && s[i] <= 'Z') ||
			(s[i] == '+' || s[i] == '-') ||
			(s[i] == '_' || s[i] == '/') || s[i] == '='
		if !r {
			return false
		}
	}
	return true
}

// isHexStr 检查并返回给定内容 `s` 是否为十六进制字符串。
// 如果 `s` 是十六进制字符串，则返回 true；否则返回 false。
func isHexStr(s string) bool {
	var r bool
	for i := 0; i < len(s); i++ {
		r = (s[i] >= '0' && s[i] <= '9') ||
			(s[i] >= 'a' && s[i] <= 'f') ||
			(s[i] >= 'A' && s[i] <= 'F')
		if !r {
			return false
		}
	}
	return true
}

// hexStrToBytes 将十六进制字符串内容转换为 []byte 类型。
func hexStrToBytes(s string) []byte {
	src := []byte(s)
	dst := make([]byte, hex.DecodedLen(len(src)))
	_, _ = hex.Decode(dst, src)
	return dst
}
