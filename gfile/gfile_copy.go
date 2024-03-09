// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类

import (
	"io"
	"os"
	"path/filepath"
	
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// CopyOption 是用于 Copy* 函数的选项。
type CopyOption struct {
	// 在源文件内容复制到目标文件后自动调用文件同步
	Sync bool

// 保留原始文件的模式到目标文件。
// 如果为 true，则 Mode 属性将不起作用。
	PreserveMode bool

// 目标文件创建时的模式
// 若PreserveMode为false，默认的文件模式为DefaultPermCopy
	Mode os.FileMode
}

// 将文件/目录从`src`复制到`dst`。
//
// 如果`src`是文件，它将调用CopyFile实现复制功能，
// 否则调用CopyDir。
//
// 如果`src`是文件，但`dst`已存在且是一个文件夹，
// 则在`dst`目录下创建一个与`src`同名的文件。
//
// 示例：
// Copy("/tmp/file1", "/tmp/file2") => 将/tmp/file1复制到/tmp/file2
// Copy("/tmp/dir1",  "/tmp/dir2")  => 将/tmp/dir1复制到/tmp/dir2
// Copy("/tmp/file1", "/tmp/dir2")  => 将/tmp/file1复制到/tmp/dir2/file1
// Copy("/tmp/dir1",  "/tmp/file2") => 报错
func X复制(文件或目录路径 string, 复制到 string, 选项 ...CopyOption) error {
	if 文件或目录路径 == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "source path cannot be empty")
	}
	if 复制到 == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "destination path cannot be empty")
	}
	srcStat, srcStatErr := os.Stat(文件或目录路径)
	if srcStatErr != nil {
		if os.IsNotExist(srcStatErr) {
			return gerror.WrapCodef(
				gcode.CodeInvalidParameter,
				srcStatErr,
				`the src path "%s" does not exist`,
				文件或目录路径,
			)
		}
		return gerror.WrapCodef(
			gcode.CodeInternalError, srcStatErr, `call os.Stat on "%s" failed`, 文件或目录路径,
		)
	}
	dstStat, dstStatErr := os.Stat(复制到)
	if dstStatErr != nil && !os.IsNotExist(dstStatErr) {
		return gerror.WrapCodef(
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
		return gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`Copy failed: the src path "%s" is file, but the dst path "%s" is folder`,
			文件或目录路径, 复制到,
		)
	}
	return X复制目录(文件或目录路径, 复制到, 选项...)
}

// CopyFile 将名为 `src` 的文件内容复制到名为 `dst` 的文件中。如果目标文件不存在，将会创建该文件。如果目标文件已存在，则其所有内容将被源文件内容替换。文件模式将从源文件复制，并且复制的数据将同步/刷新到稳定的存储设备中。
// 感谢：https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04
func X复制文件(路径, 复制到 string, 选项 ...CopyOption) (错误 error) {
	var usedOption = getCopyOption(选项...)
	if 路径 == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "source file cannot be empty")
	}
	if 复制到 == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "destination file cannot be empty")
	}
	// 如果src和dst是相同的路径，则不做任何操作。
	if 路径 == 复制到 {
		return nil
	}
	// file state check.
	srcStat, srcStatErr := os.Stat(路径)
	if srcStatErr != nil {
		if os.IsNotExist(srcStatErr) {
			return gerror.WrapCodef(
				gcode.CodeInvalidParameter,
				srcStatErr,
				`the src path "%s" does not exist`,
				路径,
			)
		}
		return gerror.WrapCodef(
			gcode.CodeInternalError, srcStatErr, `call os.Stat on "%s" failed`, 路径,
		)
	}
	dstStat, dstStatErr := os.Stat(复制到)
	if dstStatErr != nil && !os.IsNotExist(dstStatErr) {
		return gerror.WrapCodef(
			gcode.CodeInternalError, dstStatErr, `call os.Stat on "%s" failed`, 复制到,
		)
	}
	if !srcStat.IsDir() && dstStat != nil && dstStat.IsDir() {
		return gerror.NewCodef(
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
			错误 = gerror.Wrapf(e, `file close failed for "%s"`, 路径)
		}
	}()
	var outFile *os.File
	outFile, 错误 = X创建文件与目录(复制到)
	if 错误 != nil {
		return
	}
	defer func() {
		if e := outFile.Close(); e != nil {
			错误 = gerror.Wrapf(e, `file close failed for "%s"`, 复制到)
		}
	}()
	if _, 错误 = io.Copy(outFile, inFile); 错误 != nil {
		错误 = gerror.Wrapf(错误, `io.Copy failed from "%s" to "%s"`, 路径, 复制到)
		return
	}
	if usedOption.Sync {
		if 错误 = outFile.Sync(); 错误 != nil {
			错误 = gerror.Wrapf(错误, `file sync failed for file "%s"`, 复制到)
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

// CopyDir递归地复制一个目录树，尝试保持原有的权限设置。
//
// 注意：源目录必须存在，并且符号链接会被忽略并跳过。
func X复制目录(目录路径 string, 复制到 string, 选项 ...CopyOption) (错误 error) {
	var usedOption = getCopyOption(选项...)
	if 目录路径 == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "source directory cannot be empty")
	}
	if 复制到 == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "destination directory cannot be empty")
	}
	// 如果src和dst是相同的路径，则不做任何操作。
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
		return gerror.NewCode(gcode.CodeInvalidParameter, "source is not a directory")
	}
	if usedOption.PreserveMode {
		usedOption.Mode = si.Mode().Perm()
	}
	if !X是否存在(复制到) {
		if 错误 = os.MkdirAll(复制到, usedOption.Mode); 错误 != nil {
			错误 = gerror.Wrapf(
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
		错误 = gerror.Wrapf(错误, `read directory failed for path "%s"`, 目录路径)
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
