// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gfile 提供了对文件系统进行便捷操作的功能。
package 文件类

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
	
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
// 文件系统的分隔符
// 这里将分隔符定义为变量
// 以便在必要时允许开发者进行修改
	Separator = string(filepath.Separator)

	// DefaultPermOpen 是文件打开时的默认权限。
	DefaultPermOpen = os.FileMode(0666)

	// DefaultPermCopy 是文件/文件夹复制时的默认权限。
	DefaultPermCopy = os.FileMode(0755)
)

var (
// 主包的绝对文件路径。
// 它只能被检查和设置一次。
	mainPkgPath = gtype.NewString()

// selfPath 是当前运行的二进制文件路径。
// 由于它被广泛使用，因此将其定义为内部包变量。
	selfPath = ""
)

func init() {
	// 初始化内部包变量：selfPath
	selfPath, _ = exec.LookPath(os.Args[0])
	if selfPath != "" {
		selfPath, _ = filepath.Abs(selfPath)
	}
	if selfPath == "" {
		selfPath, _ = filepath.Abs(os.Args[0])
	}
}

// Mkdir 通过给定的 `path` 参数递归创建目录。
// 建议 `path` 参数使用绝对路径而非相对路径。
func X创建目录(目录 string) (错误 error) {
	if 错误 = os.MkdirAll(目录, os.ModePerm); 错误 != nil {
		错误 = gerror.Wrapf(错误, `os.MkdirAll failed for path "%s" with perm "%d"`, 目录, os.ModePerm)
		return 错误
	}
	return nil
}

// Create 以给定的 `path` 创建文件并递归创建其所在目录。
// 建议参数 `path` 使用绝对路径。
func X创建文件与目录(文件路径 string) (*os.File, error) {
	dir := X路径取父目录(文件路径)
	if !X是否存在(dir) {
		if err := X创建目录(dir); err != nil {
			return nil, err
		}
	}
	file, err := os.Create(文件路径)
	if err != nil {
		err = gerror.Wrapf(err, `os.Create failed for name "%s"`, 文件路径)
	}
	return file, err
}

// Open 以只读方式打开文件/目录。
func X打开并按只读模式(路径 string) (*os.File, error) {
	file, err := os.Open(路径)
	if err != nil {
		err = gerror.Wrapf(err, `os.Open failed for name "%s"`, 路径)
	}
	return file, err
}

// OpenFile函数以自定义的`flag`和`perm`打开文件/目录。
// 参数`flag`类似于：O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC等。
func X打开(路径 string, 读写模式 int, 权限模式 os.FileMode) (*os.File, error) {
	file, err := os.OpenFile(路径, 读写模式, 权限模式)
	if err != nil {
		err = gerror.Wrapf(err, `os.OpenFile failed with name "%s", flag "%d", perm "%d"`, 路径, 读写模式, 权限模式)
	}
	return file, err
}

// OpenWithFlag 函数以默认权限和自定义标志 `flag` 打开文件/目录。
// 默认的 `perm` 为 0666。
// 参数 `flag` 如：O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC 等。
func X打开并按默认权限(路径 string, 读写模式 int) (*os.File, error) {
	file, err := X打开(路径, 读写模式, DefaultPermOpen)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// OpenWithFlagPerm 使用自定义`flag`和`perm`打开文件/目录。
// 参数`flag`类似于：O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC等。
// 参数`perm`类似于：0600, 0666, 0777等。
func OpenWithFlagPerm别名(path string, flag int, perm os.FileMode) (*os.File, error) {
	file, err := X打开(path, flag, perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Join 使用当前系统的文件分隔符连接字符串数组paths。
func X路径生成(路径s ...string) string {
	var s string
	for _, path := range 路径s {
		if s != "" {
			s += Separator
		}
		s += gstr.TrimRight(path, Separator)
	}
	return s
}

// Exists 检查给定的 `path` 是否存在。
func X是否存在(路径 string) bool {
	if stat, err := os.Stat(路径); stat != nil && !os.IsNotExist(err) {
		return true
	}
	return false
}

// IsDir 检查给定的 `path` 是否为一个目录。
// 注意，如果 `path` 不存在，则返回 false。
func X是否存在目录(路径 string) bool {
	s, err := os.Stat(路径)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// Pwd 返回当前工作目录的绝对路径。
// 注意，如果获取当前工作目录失败，它将返回一个空字符串。
func X取当前工作目录() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path
}

// Chdir函数将当前工作目录更改为指定的目录名称。
// 如果出现错误，其类型将会是*PathError。
func X设置当前工作目录(目录 string) (错误 error) {
	错误 = os.Chdir(目录)
	if 错误 != nil {
		错误 = gerror.Wrapf(错误, `os.Chdir failed with dir "%s"`, 目录)
	}
	return
}

// IsFile 检查给定的 `path` 是否为文件，也就是说它不是一个目录。
// 注意，如果 `path` 不存在，则返回 false。
func X是否为文件(路径 string) bool {
	s, err := X取详情(路径)
	if err != nil {
		return false
	}
	return !s.IsDir()
}

// Stat返回一个FileInfo，用于描述指定名称的文件。
// 如果出现错误，其类型将会是*PathError。
func X取详情(路径 string) (os.FileInfo, error) {
	info, err := os.Stat(路径)
	if err != nil {
		err = gerror.Wrapf(err, `os.Stat failed for file "%s"`, 路径)
	}
	return info, err
}

// Move 将`src`重命名为（移动到）`dst`路径。
// 如果`dst`已存在且不是一个目录，它将会被替换。
func X移动(路径 string, 新路径 string) (错误 error) {
	错误 = os.Rename(路径, 新路径)
	if 错误 != nil {
		错误 = gerror.Wrapf(错误, `os.Rename failed from "%s" to "%s"`, 路径, 新路径)
	}
	return
}

// Rename 是 Move 的别名。
// 请参阅 Move。
func Rename别名(src string, dst string) error {
	return X移动(src, dst)
}

// DirNames 返回给定目录 `path` 下的子文件名。
// 注意，返回的名称不是绝对路径。
func X取文件列表(路径 string) ([]string, error) {
	f, err := X打开并按只读模式(路径)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdirnames(-1)
	_ = f.Close()
	if err != nil {
		err = gerror.Wrapf(err, `Read dir files failed from path "%s"`, 路径)
		return nil, err
	}
	return list, nil
}

// Glob返回所有与pattern匹配的文件名，如果没有匹配的文件，则返回nil。模式语法与Match函数中相同。
// 模式可以描述层级式的文件名，例如（假设分隔符为'/'）/usr/*/bin/ed。
//
// Glob会忽略文件系统错误，如读取目录时的I/O错误。唯一可能返回的错误是ErrBadPattern，仅当模式格式不正确时发生。
func X模糊查找(路径 string, 返回绝对路径 ...bool) ([]string, error) {
	list, err := filepath.Glob(路径)
	if err != nil {
		err = gerror.Wrapf(err, `filepath.Glob failed for pattern "%s"`, 路径)
		return nil, err
	}
	if len(返回绝对路径) > 0 && 返回绝对路径[0] && len(list) > 0 {
		array := make([]string, len(list))
		for k, v := range list {
			array[k] = X路径取文件名(v)
		}
		return array, nil
	}
	return list, nil
}

// Remove 函数用于删除参数 `path` 指定的文件或目录。
// 若参数 `path` 为目录，该函数会递归地删除整个目录及其包含的所有文件和子目录。
//
// 如果给定的 `path` 不存在或者为空，则该函数不做任何操作。
func X删除(路径或文件夹 string) (错误 error) {
	// 如果`path`为空，则此操作无任何效果。
	if 路径或文件夹 == "" {
		return nil
	}
	if 错误 = os.RemoveAll(路径或文件夹); 错误 != nil {
		错误 = gerror.Wrapf(错误, `os.RemoveAll failed for path "%s"`, 路径或文件夹)
	}
	return
}

// IsReadable 检查给定的 `path` 是否可读。
func X是否可读(路径 string) bool {
	result := true
	file, err := os.OpenFile(路径, os.O_RDONLY, DefaultPermOpen)
	if err != nil {
		result = false
	}
	file.Close()
	return result
}

// IsWritable 检查给定的 `path` 是否可写。
//
// TODO 提高性能；使用 golang.org/x/sys 以实现跨平台
func X是否可写(路径 string) bool {
	result := true
	if X是否存在目录(路径) {
		// 如果是目录，则创建一个临时文件以测试其是否可写。
		tmpFile := strings.TrimRight(路径, Separator) + Separator + gconv.String(time.Now().UnixNano())
		if f, err := X创建文件与目录(tmpFile); err != nil || !X是否存在(tmpFile) {
			result = false
		} else {
			_ = f.Close()
			_ = X删除(tmpFile)
		}
	} else {
		// 如果它是一个文件，检查是否可以打开它。
		file, err := os.OpenFile(路径, os.O_WRONLY, DefaultPermOpen)
		if err != nil {
			result = false
		}
		_ = file.Close()
	}
	return result
}

// Chmod 是 os.Chmod 的别名。
// 请参阅 os.Chmod。
func X更改权限(路径 string, 权限模式 os.FileMode) (错误 error) {
	错误 = os.Chmod(路径, 权限模式)
	if 错误 != nil {
		错误 = gerror.Wrapf(错误, `os.Chmod failed with path "%s" and mode "%s"`, 路径, 权限模式)
	}
	return
}

// Abs 返回 path 的绝对路径表示。
// 如果 path 不是绝对路径，则将其与当前工作目录连接，
// 以转换为绝对路径。对于给定的文件，其绝对路径名称不保证唯一性。
// Abs 会对结果调用 Clean 函数。
// 这段代码是关于 Go 语言（Golang）中的 `path/filepath` 包中 Abs 函数的注释翻译：
// ```go
// Abs 函数返回路径 path 的绝对路径形式。
// 如果路径 path 不是绝对路径，该函数会将其与当前工作目录拼接，
// 从而生成一个绝对路径。需要注意的是，即使对于同一个文件，其计算出的绝对路径也未必是唯一的。
// Abs 函数还会对处理后的结果调用 Clean 函数进行规范化。
func X取绝对路径(路径 string) string {
	p, _ := filepath.Abs(路径)
	return p
}

// RealPath 将给定的 `path` 转换为绝对路径
// 并检查文件路径是否存在。
// 如果文件不存在，则返回一个空字符串。
func X取绝对路径且效验(路径 string) string {
	p, err := filepath.Abs(路径)
	if err != nil {
		return ""
	}
	if !X是否存在(p) {
		return ""
	}
	return p
}

// SelfPath 返回当前运行进程（二进制文件）的绝对文件路径。
func X取当前进程路径() string {
	return selfPath
}

// SelfName 返回当前运行进程（二进制文件）的文件名。
func X取当前进程名() string {
	return X路径取文件名(X取当前进程路径())
}

// SelfDir 返回当前运行进程（二进制文件）的绝对目录路径。
func X取当前进程目录() string {
	return filepath.Dir(X取当前进程路径())
}

// Basename 返回路径的最后一个元素，其中包含文件扩展名。
// 在提取最后一个元素之前会移除尾部的路径分隔符。
// 如果路径为空，Base 返回 "."。
// 如果路径完全由分隔符组成，Basename 将返回一个单个分隔符。
// 示例：
// /var/www/file.js -> file.js
// file.js          -> file.js
func X路径取文件名(路径 string) string {
	return filepath.Base(路径)
}

// Name函数返回路径中最后一个元素的文件名部分，不包括文件扩展名。
// 示例：
// /var/www/file.js -> file
// file.js          -> file
func X路径取文件名且不含扩展名(路径 string) string {
	base := filepath.Base(路径)
	if i := strings.LastIndexByte(base, '.'); i != -1 {
		return base[:i]
	}
	return base
}

// Dir 返回路径path去掉最后一个元素后的部分，通常为路径的目录部分。
// 在去掉最后一个元素之后，Dir会对路径进行Clean操作，并移除尾部的斜杠。
// 如果 `path` 为空，Dir 返回"."。
// 如果 `path` 为".", Dir 将路径视为当前工作目录。
// 如果 `path` 完全由分隔符组成，Dir 返回一个单独的分隔符。
// 返回的路径除非是根目录，否则不会以分隔符结尾。
func X路径取父目录(路径 string) string {
	if 路径 == "." {
		return filepath.Dir(X取绝对路径且效验(路径))
	}
	return filepath.Dir(路径)
}

// IsEmpty 检查给定的 `path` 是否为空。
// 如果 `path` 是一个文件夹，它会检查该文件夹下是否存在任何文件。
// 如果 `path` 是一个文件，它会检查该文件的大小是否为零。
//
// 注意，如果 `path` 不存在，此函数也将返回 true。
func X是否为空(路径 string) bool {
	stat, err := X取详情(路径)
	if err != nil {
		return true
	}
	if stat.IsDir() {
		file, err := os.Open(路径)
		if err != nil {
			return true
		}
		defer file.Close()
		names, err := file.Readdirnames(-1)
		if err != nil {
			return true
		}
		return len(names) == 0
	} else {
		return stat.Size() == 0
	}
}

// Ext 返回路径 path 使用的文件名扩展名。
// 扩展名是从路径中最后一个元素的最后一个点开始的后缀；
// 如果没有点，则为空。
// 注意：结果中包含符号 '.'。
// 示例：
// main.go  => .go
// api.json => .json
func X路径取扩展名(路径 string) string {
	ext := filepath.Ext(路径)
	if p := strings.IndexByte(ext, '?'); p != -1 {
		ext = ext[0:p]
	}
	return ext
}

// ExtName 类似于函数 Ext，它返回路径中使用的文件名扩展名，
// 但是结果中不包含符号'.'。
// 示例：
// main.go  => go
// api.json => json
func X路径取扩展名且不含点号(路径 string) string {
	return strings.TrimLeft(X路径取扩展名(路径), ".")
}

// Temp 函数获取并返回当前系统的临时目录路径。
//
// 可选参数 `names` 指定的是子文件夹或子文件名，
// 这些名称会与当前系统的路径分隔符拼接，并将最终生成的完整路径返回。
func X取临时目录(可选路径 ...string) string {
	path := os.TempDir()
	for _, name := range 可选路径 {
		path = X路径生成(path, name)
	}
	return path
}
