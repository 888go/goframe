// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gres

import (
	"archive/zip"
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/encoding/gcompress"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

const (
	packedGoSourceTemplate = `
package %s

import "github.com/gogf/gf/v2/os/gres"

func init() {
	if err := gres.Add("%s"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
`
)

// Option包含Pack函数的额外选项。 md5:1aecf45a2bd621ac
type Option struct {
	Prefix   string // 在资源管理器中每个文件项的文件路径前缀。 md5:54cf09b52af7353f
	KeepPath bool   // 在打包时保留传递的路径，通常用于相对路径。 md5:78a556a27d461bea
}

// Pack 将由 `srcPaths` 指定的路径打包成字节。不必要的参数 `keyPrefix` 表示每个文件打包到结果字节中的前缀。
// 
// 注意，参数 `srcPaths` 支持用逗号分隔多个路径。
// 
// 警告：请使用 PackWithOption 替代此方法。
// md5:bba941587b4a7962
func Pack(srcPaths string, keyPrefix ...string) ([]byte, error) {
	option := Option{}
	if len(keyPrefix) > 0 && keyPrefix[0] != "" {
		option.Prefix = keyPrefix[0]
	}
	return PackWithOption(srcPaths, option)
}

// PackWithOption 将由 `srcPaths` 指定的路径打包成字节。
// 
// 注意，参数 `srcPaths` 支持使用逗号分隔多个路径。
// md5:15ee3362e7cd91a0
func PackWithOption(srcPaths string, option Option) ([]byte, error) {
	var buffer = bytes.NewBuffer(nil)
	err := zipPathWriter(srcPaths, buffer, option)
	if err != nil {
		return nil, err
	}
	// 使用Gzip压缩数据字节以减小大小。 md5:d15c5898ab8d9408
	return gcompress.Gzip(buffer.Bytes(), 9)
}

// PackToFile 将`srcPaths`指定的路径打包到目标文件`dstPath`中。
// 不必要的参数`keyPrefix`表示打包到结果字节中的每个文件的前缀。
//
// 注意，参数`srcPaths`支持使用','连接的多个路径。
//
// 已弃用：请改用PackToFileWithOption。
// md5:222d6d9ef38edd09
func PackToFile(srcPaths, dstPath string, keyPrefix ...string) error {
	data, err := Pack(srcPaths, keyPrefix...)
	if err != nil {
		return err
	}
	return gfile.PutBytes(dstPath, data)
}

// PackToFileWithOption 将由 `srcPaths` 指定的路径打包到目标文件 `dstPath` 中。
// 
// 注意，参数 `srcPaths` 支持使用逗号分隔多个路径。
// md5:5daf8e107f124634
func PackToFileWithOption(srcPaths, dstPath string, option Option) error {
	data, err := PackWithOption(srcPaths, option)
	if err != nil {
		return err
	}
	return gfile.PutBytes(dstPath, data)
}

// PackToGoFile 将由 `srcPaths` 指定的路径打包成目标 Go 文件 `goFilePath`，并使用给定的包名 `pkgName`。
//
// 参数 `keyPrefix`（可选）表示打包到结果字节中的每个文件的前缀。
//
// 注意，`srcPaths` 参数支持用逗号分隔多个路径。
//
// 警告：请改用 PackToGoFileWithOption。
// md5:99701ca10a176f76
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

// PackToGoFileWithOption 将由 `srcPaths` 指定的路径打包到目标Go文件 `goFilePath` 中，
// 使用给定的包名 `pkgName`。
//
// 注意，参数 `srcPaths` 支持使用逗号`,`连接多个路径。
// md5:0e7ba248d1ba0543
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

// Unpack 将由 `path` 指定的内容解压缩到 []*File 中。 md5:c88b5e566f58802e
func Unpack(path string) ([]*File, error) {
	realPath, err := gfile.Search(path)
	if err != nil {
		return nil, err
	}
	return UnpackContent(gfile.GetContents(realPath))
}

// UnpackContent 将内容解包为 []*File。 md5:a49a123f27175e6d
func UnpackContent(content string) ([]*File, error) {
	var (
		err  error
		data []byte
	)
	if isHexStr(content) {
		// 这里是为了保持与旧版本使用十六进制字符串打包字符串的兼容性。
		// TODO：未来移除这个支持。
		// md5:5253278930daad11
		data, err = gcompress.UnGzip(hexStrToBytes(content))
		if err != nil {
			return nil, err
		}
	} else if isBase64(content) {
		// 使用base64的新版本打包字符串。 md5:c884a25b1e4334ae
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

// isBase64 检查并返回给定内容 `s` 是否为 Base64 编码的字符串。
// 如果 `s` 是 Base64 字符串，它将返回 true，否则返回 false。
// md5:314047c834f3cf6c
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

// isHexStr 检查并返回给定内容 `s` 是否为十六进制字符串。如果 `s` 是十六进制字符串，它将返回 true，否则返回 false。
// md5:ca395ed524f01122
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

// hexStrToBytes 将十六进制字符串内容转换为[]byte。 md5:0b3c7f4ed4b490fb
func hexStrToBytes(s string) []byte {
	src := []byte(s)
	dst := make([]byte, hex.DecodedLen(len(src)))
	_, _ = hex.Decode(dst, src)
	return dst
}
