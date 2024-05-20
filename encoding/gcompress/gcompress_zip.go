// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gcompress

import (
	"archive/zip"
	"bytes"
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

// ZipPath 使用zip压缩算法将`fileOrFolderPaths`压缩到`dstFilePath`。
//
// 参数`paths`可以是目录或文件，支持使用`,`连接多个路径。参数`prefix`（可选）表示zip文件的路径前缀。
// md5:6754e1656d2dfc22
func ZipPath(fileOrFolderPaths, dstFilePath string, prefix ...string) error {
	writer, err := os.Create(dstFilePath)
	if err != nil {
		err = gerror.Wrapf(err, `os.Create failed for name "%s"`, dstFilePath)
		return err
	}
	defer writer.Close()
	zipWriter := zip.NewWriter(writer)
	defer zipWriter.Close()
	for _, path := range strings.Split(fileOrFolderPaths, ",") {
		path = strings.TrimSpace(path)
		if err = doZipPathWriter(path, gfile.RealPath(dstFilePath), zipWriter, prefix...); err != nil {
			return err
		}
	}
	return nil
}

// ZipPathWriter 使用zip压缩算法将`fileOrFolderPaths`压缩到`writer`中。
// 
// 注意，参数`fileOrFolderPaths`可以是目录或文件，支持使用','连接多个路径。
// 参数`prefix`（可选）表示zip文件的路径前缀。
// md5:0e6a4ca6fdf7a9d7
func ZipPathWriter(fileOrFolderPaths string, writer io.Writer, prefix ...string) error {
	zipWriter := zip.NewWriter(writer)
	defer zipWriter.Close()
	for _, path := range strings.Split(fileOrFolderPaths, ",") {
		path = strings.TrimSpace(path)
		if err := doZipPathWriter(path, "", zipWriter, prefix...); err != nil {
			return err
		}
	}
	return nil
}

// ZipPathContent 使用zip压缩算法将`fileOrFolderPaths`压缩为[]byte。
//
// 注意，参数`fileOrFolderPaths`可以是目录或文件，支持使用逗号','连接多个路径。
// 不强制要求的参数`prefix`表示zip文件中的路径前缀。
// md5:6700858e8ecb32a5
func ZipPathContent(fileOrFolderPaths string, prefix ...string) ([]byte, error) {
	var (
		err    error
		buffer = bytes.NewBuffer(nil)
	)
	if err = ZipPathWriter(fileOrFolderPaths, buffer, prefix...); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// doZipPathWriter 将给定的 `fileOrFolderPaths` 压缩，并将内容写入 `zipWriter`。
//
// 参数 `fileOrFolderPath` 可以是一个单一的文件或文件夹路径。
// 参数 `exclude` 指定了不应被压缩到 `zipWriter` 的排除文件路径，通常为目标zip文件路径。
// 参数 `prefix` 是用于zip文件的路径前缀，一般不需要。
// md5:491b5e660bfd8ac9
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

// UnZipFile 使用 ZIP 压缩算法将 `archive` 解压缩到 `dstFolderPath`。
//
// 参数 `dstFolderPath` 应该是一个目录。可选参数 `zippedPrefix` 指定了 `zippedFilePath` 的解压缩路径部分，可以用来指定要解压缩的归档文件的一部分。
// md5:4ef9114de36ab1d8
func UnZipFile(zippedFilePath, dstFolderPath string, zippedPrefix ...string) error {
	readerCloser, err := zip.OpenReader(zippedFilePath)
	if err != nil {
		err = gerror.Wrapf(err, `zip.OpenReader failed for name "%s"`, dstFolderPath)
		return err
	}
	defer readerCloser.Close()
	return unZipFileWithReader(&readerCloser.Reader, dstFolderPath, zippedPrefix...)
}

// UnZipContent 使用zip压缩算法将`zippedContent`解压缩到`dstFolderPath`。
//
// 参数`dstFolderPath`应该是一个目录。参数`zippedPrefix`指定了`zippedContent`的解压路径，可以用来指定要解压的归档文件的一部分。
// md5:808f21381d5e3681
func UnZipContent(zippedContent []byte, dstFolderPath string, zippedPrefix ...string) error {
	reader, err := zip.NewReader(bytes.NewReader(zippedContent), int64(len(zippedContent)))
	if err != nil {
		err = gerror.Wrapf(err, `zip.NewReader failed`)
		return err
	}
	return unZipFileWithReader(reader, dstFolderPath, zippedPrefix...)
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
		// 文件读取器在函数doCopyForUnZipFileWithReader中被关闭。. md5:bdbed60d16aa0ca2
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

// zipFile 将给定的 `filePath` 文件压缩，并将内容写入 `zw`。
// 参数 `prefix` 用于表示在压缩文件中的路径前缀。
// md5:69f2856c4cb49f38
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
