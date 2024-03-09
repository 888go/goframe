// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 压缩类

import (
	"archive/zip"
	"bytes"
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/888go/goframe/gcompress/internal/intlog"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

// ZipPath 使用zip压缩算法将`fileOrFolderPaths`压缩到`dstFilePath`。
//
// 参数`paths`可以是目录或文件，支持使用','连接的多个路径。
// 可选参数`prefix`表示zip文件中的路径前缀。
func Zip压缩文件(目录或文件, 压缩文件路径 string, 可选路径前缀 ...string) error {
	writer, err := os.Create(压缩文件路径)
	if err != nil {
		err = gerror.Wrapf(err, `os.Create failed for name "%s"`, 压缩文件路径)
		return err
	}
	defer writer.Close()
	zipWriter := zip.NewWriter(writer)
	defer zipWriter.Close()
	for _, path := range strings.Split(目录或文件, ",") {
		path = strings.TrimSpace(path)
		if err = doZipPathWriter(path, gfile.RealPath(压缩文件路径), zipWriter, 可选路径前缀...); err != nil {
			return err
		}
	}
	return nil
}

// ZipPathWriter 使用zip压缩算法将`fileOrFolderPaths`压缩到`writer`。
//
// 注意参数`fileOrFolderPaths`可以是目录或文件，支持使用','连接的多个路径。
// 可选参数`prefix`表示zip文件中的路径前缀。
func Zip压缩文件到Writer(目录或文件 string, writer io.Writer, 可选路径前缀 ...string) error {
	zipWriter := zip.NewWriter(writer)
	defer zipWriter.Close()
	for _, path := range strings.Split(目录或文件, ",") {
		path = strings.TrimSpace(path)
		if err := doZipPathWriter(path, "", zipWriter, 可选路径前缀...); err != nil {
			return err
		}
	}
	return nil
}

// ZipPathContent 使用zip压缩算法将`fileOrFolderPaths`压缩为[]byte。
//
// 注意，参数`fileOrFolderPaths`可以是目录或文件，支持使用','连接多个路径。
// 可选参数`prefix`表示zip文件中的路径前缀。
func Zip压缩文件到字节集(目录或文件 string, 可选路径前缀 ...string) ([]byte, error) {
	var (
		err    error
		buffer = bytes.NewBuffer(nil)
	)
	if err = Zip压缩文件到Writer(目录或文件, buffer, 可选路径前缀...); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// doZipPathWriter 将给定的 `fileOrFolderPaths` 进行压缩，并将内容写入 `zipWriter`。
//
// 参数 `fileOrFolderPath` 可以是单个文件或文件夹路径。
// 参数 `exclude` 指定了不被压缩到 `zipWriter` 中的排除文件路径，通常是指定的目标 zip 文件路径。
// 非必需参数 `prefix` 表示 zip 文件的路径前缀。
func doZipPathWriter(fileOrFolderPath string, exclude string, zipWriter *zip.Writer, prefix ...string) error {
	var (
		err   error
		files []string
	)
	fileOrFolderPath, err = gfile.Search(fileOrFolderPath)
	if err != nil {
		return err
	}
	if gfile.IsDir(fileOrFolderPath) {
		files, err = gfile.ScanDir(fileOrFolderPath, "*", true)
		if err != nil {
			return err
		}
	} else {
		files = []string{fileOrFolderPath}
	}
	headerPrefix := ""
	if len(prefix) > 0 && prefix[0] != "" {
		headerPrefix = prefix[0]
	}
	headerPrefix = strings.TrimRight(headerPrefix, "\\/")
	if gfile.IsDir(fileOrFolderPath) {
		if len(headerPrefix) > 0 {
			headerPrefix += "/"
		} else {
			headerPrefix = gfile.Basename(fileOrFolderPath)
		}
	}
	headerPrefix = strings.ReplaceAll(headerPrefix, "//", "/")
	for _, file := range files {
		if exclude == file {
			intlog.Printf(context.TODO(), `exclude file path: %s`, file)
			continue
		}
		dir := gfile.Dir(file[len(fileOrFolderPath):])
		if dir == "." {
			dir = ""
		}
		if err = zipFile(file, headerPrefix+dir, zipWriter); err != nil {
			return err
		}
	}
	return nil
}

// UnZipFile 使用zip压缩算法将`archive`解压到`dstFolderPath`。
//
// 参数`dstFolderPath`应为一个目录。
// 可选参数`zippedPrefix`用于指定`zippedFilePath`解压后的路径前缀，
// 该参数可用于指定只解压归档文件中的部分内容。
func Zip解压文件(压缩包路径, 解压目录 string, 可选路径前缀 ...string) error {
	readerCloser, err := zip.OpenReader(压缩包路径)
	if err != nil {
		err = gerror.Wrapf(err, `zip.OpenReader failed for name "%s"`, 解压目录)
		return err
	}
	defer readerCloser.Close()
	return unZipFileWithReader(&readerCloser.Reader, 解压目录, 可选路径前缀...)
}

// UnZipContent 使用zip压缩算法将`zippedContent`解压到`dstFolderPath`。
//
// 参数`dstFolderPath`应该是一个目录。
// 参数`zippedPrefix`指定了`zippedContent`解压后的路径，
// 可用于指定要解压的归档文件的部分。
// 进一步细化翻译：
// ```go
// UnZipContent 函数负责使用ZIP压缩算法将压缩内容 `zippedContent` 解压到目标文件夹 `dstFolderPath`。
//
// 参数 `dstFolderPath` 必须是一个存在的目录，解压后的文件将存放在此目录下。
// 参数 `zippedPrefix` 指定 `zippedContent` 中待解压内容的相对路径前缀，
// 通过此参数可以选择性地解压归档文件中的特定部分。
func Zip解压字节集(zip字节集 []byte, 解压目录 string, 可选路径前缀 ...string) error {
	reader, err := zip.NewReader(bytes.NewReader(zip字节集), int64(len(zip字节集)))
	if err != nil {
		err = gerror.Wrapf(err, `zip.NewReader failed`)
		return err
	}
	return unZipFileWithReader(reader, 解压目录, 可选路径前缀...)
}

func unZipFileWithReader(reader *zip.Reader, dstFolderPath string, zippedPrefix ...string) error {
	prefix := ""
	if len(zippedPrefix) > 0 {
		prefix = gstr.Replace(zippedPrefix[0], `\`, `/`)
	}
	if err := os.MkdirAll(dstFolderPath, 0755); err != nil {
		return err
	}
	var (
		name    string
		dstPath string
		dstDir  string
	)
	for _, file := range reader.File {
		name = gstr.Replace(file.Name, `\`, `/`)
		name = gstr.Trim(name, "/")
		if prefix != "" {
			if !strings.HasPrefix(name, prefix) {
				continue
			}
			name = name[len(prefix):]
		}
		dstPath = filepath.Join(dstFolderPath, name)
		if file.FileInfo().IsDir() {
			_ = os.MkdirAll(dstPath, file.Mode())
			continue
		}
		dstDir = filepath.Dir(dstPath)
		if len(dstDir) > 0 {
			if _, err := os.Stat(dstDir); os.IsNotExist(err) {
				if err = os.MkdirAll(dstDir, 0755); err != nil {
					err = gerror.Wrapf(err, `os.MkdirAll failed for path "%s"`, dstDir)
					return err
				}
			}
		}
		fileReader, err := file.Open()
		if err != nil {
			err = gerror.Wrapf(err, `file.Open failed`)
			return err
		}
		// 文件读取器在函数doCopyForUnZipFileWithReader中关闭。
		if err = doCopyForUnZipFileWithReader(file, fileReader, dstPath); err != nil {
			return err
		}
	}
	return nil
}

func doCopyForUnZipFileWithReader(file *zip.File, fileReader io.ReadCloser, dstPath string) error {
	defer fileReader.Close()
	targetFile, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
		err = gerror.Wrapf(err, `os.OpenFile failed for name "%s"`, dstPath)
		return err
	}
	defer targetFile.Close()

	if _, err = io.Copy(targetFile, fileReader); err != nil {
		err = gerror.Wrapf(err, `io.Copy failed from "%s" to "%s"`, file.Name, dstPath)
		return err
	}
	return nil
}

// zipFile 将指定 `filePath` 的文件进行压缩，并将压缩内容写入 `zw`。
// 参数 `prefix` 表示压缩文件路径的前缀。
func zipFile(filePath string, prefix string, zw *zip.Writer) error {
	file, err := os.Open(filePath)
	if err != nil {
		err = gerror.Wrapf(err, `os.Open failed for name "%s"`, filePath)
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		err = gerror.Wrapf(err, `file.Stat failed for name "%s"`, filePath)
		return err
	}

	header, err := createFileHeader(info, prefix)
	if err != nil {
		return err
	}

	if info.IsDir() {
		header.Name += "/"
	} else {
		header.Method = zip.Deflate
	}

	writer, err := zw.CreateHeader(header)
	if err != nil {
		err = gerror.Wrapf(err, `zip.Writer.CreateHeader failed for header "%#v"`, header)
		return err
	}
	if !info.IsDir() {
		if _, err = io.Copy(writer, file); err != nil {
			err = gerror.Wrapf(err, `io.Copy failed from "%s" to "%s"`, filePath, header.Name)
			return err
		}
	}
	return nil
}

func createFileHeader(info os.FileInfo, prefix string) (*zip.FileHeader, error) {
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		err = gerror.Wrapf(err, `zip.FileInfoHeader failed for info "%#v"`, info)
		return nil, err
	}

	if len(prefix) > 0 {
		prefix = strings.ReplaceAll(prefix, `\`, `/`)
		prefix = strings.TrimRight(prefix, `/`)
		header.Name = prefix + `/` + header.Name
	}
	return header, nil
}
