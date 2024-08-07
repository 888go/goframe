// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文件类

import (
	"io"
	"os"
	"path/filepath"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
)

// CopyOption 是 Copy* 函数的选项。 md5:1863c87f867e036e
type CopyOption struct {
		// 在源文件内容复制到目标文件后，自动调用文件同步。 md5:ef1f9250b5fdabe3
	Sync bool

	// 保留源文件的模式到目标文件。如果为true，Mode属性将没有意义。
	// md5:681b0704991c814c
	PreserveMode bool

	// 创建目标文件的模式。
	// 如果PreserveMode为false，默认的文件模式是DefaultPermCopy。
	// md5:e495278ff0787785
	Mode os.FileMode
}

// 将源`src`文件/目录复制到目标`dst`。
//
// 如果`src`是文件，它会调用CopyFile来实现复制功能，
// 否则，它会调用CopyDir。
//
// 如果`src`是文件，但`dst`已经存在并且是一个文件夹，
// 那么它会在`dst`文件夹中创建一个与`src`同名的文件。
//
// 例如：
// X复制("/tmp/file1", "/tmp/file2") => 将/tmp/file1复制到/tmp/file2
// X复制("/tmp/dir1",  "/tmp/dir2")  => 将/tmp/dir1复制到/tmp/dir2
// X复制("/tmp/file1", "/tmp/dir2")  => 将/tmp/file1复制到/tmp/dir2/file1
// X复制("/tmp/dir1",  "/tmp/file2") => 出错
// md5:51c6598025f6b135
func X复制(文件或目录路径 string, 复制到 string, 选项 ...CopyOption) error {
	if 文件或目录路径 == "" {
		return gerror.X创建错误码(gcode.CodeInvalidParameter, "source path cannot be empty")
	}
	if 复制到 == "" {
		return gerror.X创建错误码(gcode.CodeInvalidParameter, "destination path cannot be empty")
	}
	srcStat, srcStatErr := os.Stat(文件或目录路径)
	if srcStatErr != nil {
		if os.IsNotExist(srcStatErr) {
			return gerror.X多层错误码并格式化(
				gcode.CodeInvalidParameter,
				srcStatErr,
				`the src path "%s" does not exist`,
				文件或目录路径,
			)
		}
		return gerror.X多层错误码并格式化(
			gcode.CodeInternalError, srcStatErr, `call os.Stat on "%s" failed`, 文件或目录路径,
		)
	}
	dstStat, dstStatErr := os.Stat(复制到)
	if dstStatErr != nil && !os.IsNotExist(dstStatErr) {
		return gerror.X多层错误码并格式化(
			gcode.CodeInternalError, dstStatErr, `call os.Stat on "%s" failed`, 复制到)
	}

	if X是否为文件(文件或目录路径) {
		var isDstExist = false
		if dstStat != nil && !os.IsNotExist(dstStatErr) {
			isDstExist = true
		}
		if isDstExist && dstStat.IsDir() {
			var (
				srcName = X路径取文件名(文件或目录路径)
				dstPath = X路径生成(复制到, srcName)
			)
			return X复制文件(文件或目录路径, dstPath, 选项...)
		}
		return X复制文件(文件或目录路径, 复制到, 选项...)
	}
	if !srcStat.IsDir() && dstStat != nil && dstStat.IsDir() {
		return gerror.X创建错误码并格式化(
			gcode.CodeInvalidParameter,
			`Copy failed: the src path "%s" is file, but the dst path "%s" is folder`,
			文件或目录路径, 复制到,
		)
	}
	return X复制目录(文件或目录路径, 复制到, 选项...)
}

// X复制文件 将名为 `src` 的文件的内容复制到由 `dst` 指定的文件中。如果目标文件不存在，它将被创建。如果目标文件已存在，其所有内容将被源文件的内容替换。文件权限将从源文件复制，并且复制的数据会被同步/刷新到稳定的存储中。
// 谢谢：https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04
// md5:e2fc3c25ff06fa5b
func X复制文件(路径, 复制到 string, 选项 ...CopyOption) (错误 error) {
	var usedOption = getCopyOption(选项...)
	if 路径 == "" {
		return gerror.X创建错误码(gcode.CodeInvalidParameter, "source file cannot be empty")
	}
	if 复制到 == "" {
		return gerror.X创建错误码(gcode.CodeInvalidParameter, "destination file cannot be empty")
	}
		// 如果src和dst是相同的路径，它不会做任何事情。 md5:1ad6359456a4bebc
	if 路径 == 复制到 {
		return nil
	}
	// file state check.
	srcStat, srcStatErr := os.Stat(路径)
	if srcStatErr != nil {
		if os.IsNotExist(srcStatErr) {
			return gerror.X多层错误码并格式化(
				gcode.CodeInvalidParameter,
				srcStatErr,
				`the src path "%s" does not exist`,
				路径,
			)
		}
		return gerror.X多层错误码并格式化(
			gcode.CodeInternalError, srcStatErr, `call os.Stat on "%s" failed`, 路径,
		)
	}
	dstStat, dstStatErr := os.Stat(复制到)
	if dstStatErr != nil && !os.IsNotExist(dstStatErr) {
		return gerror.X多层错误码并格式化(
			gcode.CodeInternalError, dstStatErr, `call os.Stat on "%s" failed`, 复制到,
		)
	}
	if !srcStat.IsDir() && dstStat != nil && dstStat.IsDir() {
		return gerror.X创建错误码并格式化(
			gcode.CodeInvalidParameter,
			`CopyFile failed: the src path "%s" is file, but the dst path "%s" is folder`,
			路径, 复制到,
		)
	}
	// copy file logic.
	var inFile *os.File
	inFile, 错误 = X打开并按只读模式(路径)
	if 错误 != nil {
		return
	}
	defer func() {
		if e := inFile.Close(); e != nil {
			错误 = gerror.X多层错误并格式化(e, `file close failed for "%s"`, 路径)
		}
	}()
	var outFile *os.File
	outFile, 错误 = X创建文件与目录(复制到)
	if 错误 != nil {
		return
	}
	defer func() {
		if e := outFile.Close(); e != nil {
			错误 = gerror.X多层错误并格式化(e, `file close failed for "%s"`, 复制到)
		}
	}()
	if _, 错误 = io.Copy(outFile, inFile); 错误 != nil {
		错误 = gerror.X多层错误并格式化(错误, `io.Copy failed from "%s" to "%s"`, 路径, 复制到)
		return
	}
	if usedOption.Sync {
		if 错误 = outFile.Sync(); 错误 != nil {
			错误 = gerror.X多层错误并格式化(错误, `file sync failed for file "%s"`, 复制到)
			return
		}
	}
	if usedOption.PreserveMode {
		usedOption.Mode = srcStat.Mode().Perm()
	}
	if 错误 = X更改权限(复制到, usedOption.Mode); 错误 != nil {
		return
	}
	return
}

// X复制目录 递归地复制目录树，尝试保留权限。
//
// 注意，源目录必须存在，并且符号链接将被忽略和跳过。
// md5:4dd9167e563fa997
func X复制目录(目录路径 string, 复制到 string, 选项 ...CopyOption) (错误 error) {
	var usedOption = getCopyOption(选项...)
	if 目录路径 == "" {
		return gerror.X创建错误码(gcode.CodeInvalidParameter, "source directory cannot be empty")
	}
	if 复制到 == "" {
		return gerror.X创建错误码(gcode.CodeInvalidParameter, "destination directory cannot be empty")
	}
		// 如果src和dst是相同的路径，它不会做任何事情。 md5:1ad6359456a4bebc
	if 目录路径 == 复制到 {
		return nil
	}
	目录路径 = filepath.Clean(目录路径)
	复制到 = filepath.Clean(复制到)
	si, 错误 := X取详情(目录路径)
	if 错误 != nil {
		return 错误
	}
	if !si.IsDir() {
		return gerror.X创建错误码(gcode.CodeInvalidParameter, "source is not a directory")
	}
	if usedOption.PreserveMode {
		usedOption.Mode = si.Mode().Perm()
	}
	if !X是否存在(复制到) {
		if 错误 = os.MkdirAll(复制到, usedOption.Mode); 错误 != nil {
			错误 = gerror.X多层错误并格式化(
				错误,
				`create directory failed for path "%s", perm "%s"`,
				复制到,
				usedOption.Mode,
			)
			return
		}
	}
	entries, 错误 := os.ReadDir(目录路径)
	if 错误 != nil {
		错误 = gerror.X多层错误并格式化(错误, `read directory failed for path "%s"`, 目录路径)
		return
	}
	for _, entry := range entries {
		srcPath := filepath.Join(目录路径, entry.Name())
		dstPath := filepath.Join(复制到, entry.Name())
		if entry.IsDir() {
			if 错误 = X复制目录(srcPath, dstPath); 错误 != nil {
				return
			}
		} else {
			// Skip symlinks.
			if entry.Type()&os.ModeSymlink != 0 {
				continue
			}
			if 错误 = X复制文件(srcPath, dstPath, 选项...); 错误 != nil {
				return
			}
		}
	}
	return
}

func getCopyOption(option ...CopyOption) CopyOption {
	var usedOption CopyOption
	if len(option) > 0 {
		usedOption = option[0]
	}
	if usedOption.Mode == 0 {
		usedOption.Mode = DefaultPermCopy
	}
	return usedOption
}
