// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gres

import (
	"archive/zip"
	"io"
	"os"
	"strings"
	"time"
	
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/fileinfo"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/text/gregex"
)

// ZipPathWriter 使用zip压缩算法将`paths`压缩到`writer`。
// 不必要的参数`prefix`表示zip文件的路径前缀。
//
// 注意，参数`paths`可以是目录或文件，支持使用','连接的多个路径。
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

// doZipPathWriter 函数用于压缩指定 `path` 的文件并将压缩内容写入到 `zipWriter`。
// 参数 `exclude` 指定不需要被压缩到 `zipWriter` 中的文件路径，通常是指定的目标 zip 文件路径本身。
// 参数 `prefix`（非必需）表示 zip 文件中的路径前缀。
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
// 它在资源管理器中保留从文件系统到zip信息的路径。
// 通常对于相对路径有意义，但对于绝对路径意义不大。
			headerPrefix = srcPath
		} else {
			headerPrefix = gfile.Basename(absolutePath)
		}
	}
	headerPrefix = strings.ReplaceAll(headerPrefix, `//`, `/`)
	for _, file := range files {
// 这里计算文件名前缀，特别是打包目录。
// 例如：
// 路径：dir1
// 文件：dir1/dir2/file
// file[len(absolutePath):] => /dir2/file （取绝对路径后缀部分）
// gfile.Dir(subFilePath)   => /dir2 （获取子文件路径的目录部分）
		var subFilePath string
		// 正常处理：从文件中移除`absolutePath`(源目录路径)。
		subFilePath = file[len(absolutePath):]
		if subFilePath != "" {
			subFilePath = gfile.Dir(subFilePath)
		}
		if err = zipFile(file, headerPrefix+subFilePath, zipWriter); err != nil {
			return err
		}
	}
	// 将所有目录添加到zip归档中。
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

// zipFile 将给定 `path` 的文件压缩，并将内容写入 `zw`。
// 参数 `prefix` 表示 zip 文件的路径前缀。
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
		// 默认压缩级别。
		header.Method = zip.Deflate
	}
	// Zip头包含zip文件的信息。
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
