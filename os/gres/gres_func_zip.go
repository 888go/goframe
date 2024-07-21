// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gres

import (
	"archive/zip"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/fileinfo"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
)

// ZipPathWriter 使用zip压缩算法将`paths`压缩到`writer`中。
// 不需要的参数`prefix`表示zip文件中的路径前缀。
//
// 注意，参数`paths`可以是目录或文件，支持使用逗号','连接多个路径。
// md5:d392a5d80ec973d9
func zipPathWriter(paths string, writer io.Writer, option ...Option) error {
	zipWriter := zip.NewWriter(writer)
	defer zipWriter.Close()
	for _, path := range strings.Split(paths, ",") {
		path = strings.TrimSpace(path)
		if err := doZipPathWriter(path, zipWriter, option...); err != nil {
			return err
		}
	}
	return nil
}

// doZipPathWriter 将给定路径`path`的文件压缩，并将内容写入`zipWriter`。
// 参数`exclude`指定了不被压缩到`zipWriter`的排除文件路径，通常为目标zip文件路径。
// 参数`prefix`表示zip文件的路径前缀，可选。
// md5:46c9d23dcfa03c25
func doZipPathWriter(srcPath string, zipWriter *zip.Writer, option ...Option) error {
	var (
		err          error
		files        []string
		usedOption   Option
		absolutePath string
	)
	if len(option) > 0 {
		usedOption = option[0]
	}
	absolutePath, err = gfile.Search(srcPath)
	if err != nil {
		return err
	}
	if gfile.IsDir(absolutePath) {
		files, err = gfile.ScanDir(absolutePath, "*", true)
		if err != nil {
			return err
		}
	} else {
		files = []string{absolutePath}
	}
	headerPrefix := strings.TrimRight(usedOption.Prefix, `\/`)
	if headerPrefix != "" && gfile.IsDir(absolutePath) {
		headerPrefix += "/"
	}

	if headerPrefix == "" {
		if usedOption.KeepPath {
// 它在资源管理器中保存从文件系统到zip信息的路径。通常对于相对路径，绝对路径意义不大。
// md5:bba8ee186d063506
			headerPrefix = srcPath
		} else {
			headerPrefix = gfile.Basename(absolutePath)
		}
	}
	headerPrefix = strings.ReplaceAll(headerPrefix, `//`, `/`)
	for _, file := range files {
// 它在这里计算文件名前缀，特别是打包目录。
// 例如：
// 路径：dir1
// 文件：dir1/dir2/file
// file[字符串长度(absolutePath)：] => /dir2/file
// gfile.Dir(subFilePath) => /dir2
// md5:80c4920a234839ce
		var subFilePath string
		// 正常处理：移除文件的`absolutePath`（源目录路径）。 md5:66bfc67471cf5f63
		subFilePath = file[len(absolutePath):]
		if subFilePath != "" {
			subFilePath = gfile.Dir(subFilePath)
		}
		if err = zipFile(file, headerPrefix+subFilePath, zipWriter); err != nil {
			return err
		}
	}
	// 将所有目录添加到zip归档中。 md5:f8910528d8dda79d
	if headerPrefix != "" {
		var (
			name    string
			tmpPath = headerPrefix
		)
		for {
			name = strings.ReplaceAll(gfile.Basename(tmpPath), `\`, `/`)
			err = zipFileVirtual(fileinfo.New(name, 0, os.ModeDir|os.ModePerm, time.Now()), tmpPath, zipWriter)
			if err != nil {
				return err
			}
			if tmpPath == `/` || !strings.Contains(tmpPath, `/`) {
				break
			}
			tmpPath = gfile.Dir(tmpPath)
		}
	}
	return nil
}

// zipFile 压缩给定路径 `path` 的文件，并将内容写入 `zw`。参数 `prefix` 表示zip文件的路径前缀。
// md5:bb4064703bf6d8ad
func zipFile(path string, prefix string, zw *zip.Writer) error {
	prefix = strings.ReplaceAll(prefix, `//`, `/`)
	file, err := os.Open(path)
	if err != nil {
		err = gerror.Wrapf(err, `os.Open failed for path "%s"`, path)
		return nil
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		err = gerror.Wrapf(err, `read file stat failed for path "%s"`, path)
		return err
	}

	header, err := createFileHeader(info, prefix)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		// 默认压缩级别。 md5:27fa604e26eb1270
		header.Method = zip.Deflate
	}
	// 包含ZIP文件信息的ZIP头。 md5:df2d788fe836a2e5
	writer, err := zw.CreateHeader(header)
	if err != nil {
		err = gerror.Wrapf(err, `create zip header failed for %#v`, header)
		return err
	}
	if !info.IsDir() {
		if _, err = io.Copy(writer, file); err != nil {
			err = gerror.Wrapf(err, `io.Copy failed for file "%s"`, path)
			return err
		}
	}
	return nil
}

func zipFileVirtual(info os.FileInfo, path string, zw *zip.Writer) error {
	header, err := createFileHeader(info, "")
	if err != nil {
		return err
	}
	header.Name = path
	if _, err = zw.CreateHeader(header); err != nil {
		err = gerror.Wrapf(err, `create zip header failed for %#v`, header)
		return err
	}
	return nil
}

func createFileHeader(info os.FileInfo, prefix string) (*zip.FileHeader, error) {
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		err = gerror.Wrapf(err, `create file header failed for name "%s"`, info.Name())
		return nil, err
	}
	if len(prefix) > 0 {
		header.Name = prefix + `/` + header.Name
		header.Name = strings.ReplaceAll(header.Name, `\`, `/`)
		header.Name, _ = gregex.ReplaceString(`/{2,}`, `/`, header.Name)
	}
	return header, nil
}
