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
// Copy("/tmp/file1", "/tmp/file2") => 将/tmp/file1复制到/tmp/file2
// Copy("/tmp/dir1",  "/tmp/dir2")  => 将/tmp/dir1复制到/tmp/dir2
// Copy("/tmp/file1", "/tmp/dir2")  => 将/tmp/file1复制到/tmp/dir2/file1
// Copy("/tmp/dir1",  "/tmp/file2") => 出错
// md5:51c6598025f6b135
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

// CopyFile 将名为 `src` 的文件的内容复制到由 `dst` 指定的文件中。如果目标文件不存在，它将被创建。如果目标文件已存在，其所有内容将被源文件的内容替换。文件权限将从源文件复制，并且复制的数据会被同步/刷新到稳定的存储中。
// 谢谢：https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04
// md5:e2fc3c25ff06fa5b
func CopyFile(src, dst string, option ...CopyOption) (err error) {
	var usedOption = getCopyOption(option...)
	if src == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "source file cannot be empty")
	}
	if dst == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "destination file cannot be empty")
	}
		// 如果src和dst是相同的路径，它不会做任何事情。 md5:1ad6359456a4bebc
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

// CopyDir 递归地复制目录树，尝试保留权限。
//
// 注意，源目录必须存在，并且符号链接将被忽略和跳过。
// md5:4dd9167e563fa997
func CopyDir(src string, dst string, option ...CopyOption) (err error) {
	var usedOption = getCopyOption(option...)
	if src == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "source directory cannot be empty")
	}
	if dst == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "destination directory cannot be empty")
	}
		// 如果src和dst是相同的路径，它不会做任何事情。 md5:1ad6359456a4bebc
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
