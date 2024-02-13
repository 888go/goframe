// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdebug

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"
	
	"github.com/888go/goframe/encoding/ghash"
	"github.com/888go/goframe/errors/gerror"
)

// BinVersion 返回当前运行二进制文件的版本。
// 它采用 ghash.BKDRHash+BASE36 算法来计算二进制文件的唯一版本。
func BinVersion() string {
	if binaryVersion == "" {
		binaryContent, _ := os.ReadFile(selfPath)
		binaryVersion = strconv.FormatInt(
			int64(哈希类.BKDR(binaryContent)),
			36,
		)
	}
	return binaryVersion
}

// BinVersionMd5 返回当前运行二进制文件的版本。
// 它使用MD5算法来计算该二进制文件的唯一版本。
func BinVersionMd5() string {
	if binaryVersionMd5 == "" {
		binaryVersionMd5, _ = md5File(selfPath)
	}
	return binaryVersionMd5
}

// md5File 使用MD5算法加密`path`文件内容。
func md5File(path string) (encrypt string, err error) {
	f, err := os.Open(path)
	if err != nil {
		err = 错误类.X多层错误并格式化(err, `os.Open failed for name "%s"`, path)
		return "", err
	}
	defer f.Close()
	h := md5.New()
	_, err = io.Copy(h, f)
	if err != nil {
		err = 错误类.X多层错误(err, `io.Copy failed`)
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
