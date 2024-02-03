// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfile

import (
	"io"
	"os"
	"path/filepath"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
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
func Copy(src string, dst string, option ...CopyOption) error {
	if src == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "source path cannot be empty")
	}
	if dst == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "destination path cannot be empty")
	}
	srcStat, srcStatErr := os.Stat(src)
	if srcStatErr != nil {
		if os.IsNotExist(srcStatErr) {
			return gerror.WrapCodef(
				gcode.CodeInvalidParameter,
				srcStatErr,
				`the src path "%s" does not exist`,
				src,
			)
		}
		return gerror.WrapCodef(
			gcode.CodeInternalError, srcStatErr, `call os.Stat on "%s" failed`, src,
		)
	}
	dstStat, dstStatErr := os.Stat(dst)
	if dstStatErr != nil && !os.IsNotExist(dstStatErr) {
		return gerror.WrapCodef(
			gcode.CodeInternalError, dstStatErr, `call os.Stat on "%s" failed`, dst)
	}

	if IsFile(src) {
		var isDstExist = false
		if dstStat != nil && !os.IsNotExist(dstStatErr) {
			isDstExist = true
		}
		if isDstExist && dstStat.IsDir() {
			var (
				srcName = Basename(src)
				dstPath = Join(dst, srcName)
			)
			return CopyFile(src, dstPath, option...)
		}
		return CopyFile(src, dst, option...)
	}
	if !srcStat.IsDir() && dstStat != nil && dstStat.IsDir() {
		return gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`Copy failed: the src path "%s" is file, but the dst path "%s" is folder`,
			src, dst,
		)
	}
	return CopyDir(src, dst, option...)
}

// CopyFile 将名为 `src` 的文件内容复制到名为 `dst` 的文件中。如果目标文件不存在，将会创建该文件。如果目标文件已存在，则其所有内容将被源文件内容替换。文件模式将从源文件复制，并且复制的数据将同步/刷新到稳定的存储设备中。
// 感谢：https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04
func CopyFile(src, dst string, option ...CopyOption) (err error) {
	var usedOption = getCopyOption(option...)
	if src == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "source file cannot be empty")
	}
	if dst == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "destination file cannot be empty")
	}
	// 如果src和dst是相同的路径，则不做任何操作。
	if src == dst {
		return nil
	}
	// file state check.
	srcStat, srcStatErr := os.Stat(src)
	if srcStatErr != nil {
		if os.IsNotExist(srcStatErr) {
			return gerror.WrapCodef(
				gcode.CodeInvalidParameter,
				srcStatErr,
				`the src path "%s" does not exist`,
				src,
			)
		}
		return gerror.WrapCodef(
			gcode.CodeInternalError, srcStatErr, `call os.Stat on "%s" failed`, src,
		)
	}
	dstStat, dstStatErr := os.Stat(dst)
	if dstStatErr != nil && !os.IsNotExist(dstStatErr) {
		return gerror.WrapCodef(
			gcode.CodeInternalError, dstStatErr, `call os.Stat on "%s" failed`, dst,
		)
	}
	if !srcStat.IsDir() && dstStat != nil && dstStat.IsDir() {
		return gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`CopyFile failed: the src path "%s" is file, but the dst path "%s" is folder`,
			src, dst,
		)
	}
	// copy file logic.
	var inFile *os.File
	inFile, err = Open(src)
	if err != nil {
		return
	}
	defer func() {
		if e := inFile.Close(); e != nil {
			err = gerror.Wrapf(e, `file close failed for "%s"`, src)
		}
	}()
	var outFile *os.File
	outFile, err = Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := outFile.Close(); e != nil {
			err = gerror.Wrapf(e, `file close failed for "%s"`, dst)
		}
	}()
	if _, err = io.Copy(outFile, inFile); err != nil {
		err = gerror.Wrapf(err, `io.Copy failed from "%s" to "%s"`, src, dst)
		return
	}
	if usedOption.Sync {
		if err = outFile.Sync(); err != nil {
			err = gerror.Wrapf(err, `file sync failed for file "%s"`, dst)
			return
		}
	}
	if usedOption.PreserveMode {
		usedOption.Mode = srcStat.Mode().Perm()
	}
	if err = Chmod(dst, usedOption.Mode); err != nil {
		return
	}
	return
}

// CopyDir递归地复制一个目录树，尝试保持原有的权限设置。
//
// 注意：源目录必须存在，并且符号链接会被忽略并跳过。
func CopyDir(src string, dst string, option ...CopyOption) (err error) {
	var usedOption = getCopyOption(option...)
	if src == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "source directory cannot be empty")
	}
	if dst == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "destination directory cannot be empty")
	}
	// 如果src和dst是相同的路径，则不做任何操作。
	if src == dst {
		return nil
	}
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)
	si, err := Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return gerror.NewCode(gcode.CodeInvalidParameter, "source is not a directory")
	}
	if usedOption.PreserveMode {
		usedOption.Mode = si.Mode().Perm()
	}
	if !Exists(dst) {
		if err = os.MkdirAll(dst, usedOption.Mode); err != nil {
			err = gerror.Wrapf(
				err,
				`create directory failed for path "%s", perm "%s"`,
				dst,
				usedOption.Mode,
			)
			return
		}
	}
	entries, err := os.ReadDir(src)
	if err != nil {
		err = gerror.Wrapf(err, `read directory failed for path "%s"`, src)
		return
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		if entry.IsDir() {
			if err = CopyDir(srcPath, dstPath); err != nil {
				return
			}
		} else {
			// Skip symlinks.
			if entry.Type()&os.ModeSymlink != 0 {
				continue
			}
			if err = CopyFile(srcPath, dstPath, option...); err != nil {
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
