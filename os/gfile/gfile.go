// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// gfile 包提供了易于使用的文件系统操作。. md5:51d18e994a768bb4
package gfile

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
// 文件系统分隔符。
// 这里将分隔符定义为变量，
// 以便开发人员在必要时进行修改。
// md5:ec0b6e47ec28478f
	Separator = string(filepath.Separator)

	// DefaultPermOpen 是文件打开的默认权限。. md5:dc57341030d46a11
	DefaultPermOpen = os.FileMode(0666)

	// DefaultPermCopy 是用于文件/文件夹复制的默认权限。. md5:fdef6c133c375aa4
	DefaultPermCopy = os.FileMode(0755)
)

var (
// 主包的绝对文件路径。
// 只能检查和设置一次。
// md5:4a0d292a2835bc99
	mainPkgPath = gtype.NewString()

// selfPath 是当前运行二进制文件的路径。
// 由于它最常被使用，因此作为内部包变量进行定义。
// md5:0e75acfdf2f4b9e7
	selfPath = ""
)

func init() {
	// 初始化内部包变量：selfPath。. md5:8d1168ac7361bb54
	selfPath, _ = exec.LookPath(os.Args[0])
	if selfPath != "" {
		selfPath, _ = filepath.Abs(selfPath)
	}
	if selfPath == "" {
		selfPath, _ = filepath.Abs(os.Args[0])
	}
}

// Mkdir 递归创建给定的 `path` 所表示的目录。建议使用绝对路径而非相对路径作为参数。
// md5:e78abb40a45c2886
func Mkdir(path string) (err error) {
	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		err = gerror.Wrapf(err, `os.MkdirAll failed for path "%s" with perm "%d"`, path, os.ModePerm)
		return err
	}
	return nil
}

// Create 递归地创建具有给定`path`的文件。
// 建议参数`path`使用绝对路径。
// md5:163accaf36969b42
func Create(path string) (*os.File, error) {
	dir := Dir(path)
	if !Exists(dir) {
		if err := Mkdir(dir); err != nil {
			return nil, err
		}
	}
	file, err := os.Create(path)
	if err != nil {
		err = gerror.Wrapf(err, `os.Create failed for name "%s"`, path)
	}
	return file, err
}

// Open以只读方式打开文件/目录。. md5:7f50cf0f63b9e34e
func Open(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		err = gerror.Wrapf(err, `os.Open failed for name "%s"`, path)
	}
	return file, err
}

// OpenFile 用自定义的 `flag` 和 `perm` 打开文件或目录。参数 `flag` 类似于：O_RDONLY（只读），O_RDWR（读写），O_RDWR|O_CREATE|O_TRUNC 等。
// md5:0cef38d8408ed250
func OpenFile(path string, flag int, perm os.FileMode) (*os.File, error) {
	file, err := os.OpenFile(path, flag, perm)
	if err != nil {
		err = gerror.Wrapf(err, `os.OpenFile failed with name "%s", flag "%d", perm "%d"`, path, flag, perm)
	}
	return file, err
}

// OpenWithFlag 使用默认权限和自定义`flag`打开文件/目录。默认的`perm`为0666。
// 参数`flag`类似于：O_RDONLY（只读），O_RDWR（读写），O_RDWR|O_CREATE|O_TRUNC（读写并创建或截断），等等。
// md5:2e77d9a0acc43298
func OpenWithFlag(path string, flag int) (*os.File, error) {
	file, err := OpenFile(path, flag, DefaultPermOpen)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// OpenWithFlagPerm 使用自定义的 `flag` 和 `perm` 来打开文件/目录。
// 参数 `flag` 例如：O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC 等。
// 参数 `perm` 例如：0600, 0666, 0777 等。
// md5:1f44c05cc68654d8
func OpenWithFlagPerm(path string, flag int, perm os.FileMode) (*os.File, error) {
	file, err := OpenFile(path, flag, perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Join 使用当前系统的文件分隔符将字符串数组路径连接起来。. md5:349ebcda51de0442
func Join(paths ...string) string {
	var s string
	for _, path := range paths {
		if s != "" {
			s += Separator
		}
		s += gstr.TrimRight(path, Separator)
	}
	return s
}

// Exists检查给定的`path`是否存在。. md5:523f33d374bd2841
func Exists(path string) bool {
	if stat, err := os.Stat(path); stat != nil && !os.IsNotExist(err) {
		return true
	}
	return false
}

// IsDir 检查给定的 `path` 是否为目录。
// 注意，如果 `path` 不存在，它将返回 false。
// md5:c5b2468307c9c9e2
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// Pwd 返回当前工作目录的绝对路径。
// 注意，如果获取当前工作目录失败，它将返回一个空字符串。
// md5:90f41f1bfdd61dba
func Pwd() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path
}

// Chdir 将当前工作目录更改为指定的目录。
// 如果发生错误，该错误将为 *PathError 类型。
// md5:7bafb79d47f0c3b3
func Chdir(dir string) (err error) {
	err = os.Chdir(dir)
	if err != nil {
		err = gerror.Wrapf(err, `os.Chdir failed with dir "%s"`, dir)
	}
	return
}

// IsFile 检查给定的 `path` 是否为文件，即不是目录。
// 注意，如果 `path` 不存在，它将返回 false。
// md5:38595d733f36d367
func IsFile(path string) bool {
	s, err := Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}

// Stat 返回一个描述命名文件的 FileInfo。如果出现错误，错误类型为 *PathError。
// md5:f4ee45de3278c17f
func Stat(path string) (os.FileInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		err = gerror.Wrapf(err, `os.Stat failed for file "%s"`, path)
	}
	return info, err
}

// Move 将路径 `src` 重命名（移动）到 `dst`。
// 如果 `dst` 已经存在且不是一个目录，它将被替换。
// md5:4bc635341db78f64
func Move(src string, dst string) (err error) {
	err = os.Rename(src, dst)
	if err != nil {
		err = gerror.Wrapf(err, `os.Rename failed from "%s" to "%s"`, src, dst)
	}
	return
}

// Rename 是 Move 的别名。
// 参考 Move。
// md5:f235456881c23527
func Rename(src string, dst string) error {
	return Move(src, dst)
}

// DirNames 返回给定目录 `path` 的子文件名。请注意，返回的名称不是绝对路径。
// md5:62471f4b0e0bb389
func DirNames(path string) ([]string, error) {
	f, err := Open(path)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdirnames(-1)
	_ = f.Close()
	if err != nil {
		err = gerror.Wrapf(err, `Read dir files failed from path "%s"`, path)
		return nil, err
	}
	return list, nil
}

// Glob 返回匹配 pattern 的所有文件名，如果没有匹配的文件，则返回 nil。pattern 的语法与 Match 中的相同。pattern 可以描述像 /usr/*/bin/ed（假设分隔符为 '/'）这样的层次名称。
// 
// Glob 忽略读取目录时发生的文件系统错误，如 I/O 错误。可能返回的唯一错误是 ErrBadPattern，当 pattern 格式不正确时。
// md5:0baeeb8710df5d67
func Glob(pattern string, onlyNames ...bool) ([]string, error) {
	list, err := filepath.Glob(pattern)
	if err != nil {
		err = gerror.Wrapf(err, `filepath.Glob failed for pattern "%s"`, pattern)
		return nil, err
	}
	if len(onlyNames) > 0 && onlyNames[0] && len(list) > 0 {
		array := make([]string, len(list))
		for k, v := range list {
			array[k] = Basename(v)
		}
		return array, nil
	}
	return list, nil
}

// Remove 删除具有`path`参数的所有文件/目录。
// 如果参数`path`是一个目录，它会递归地删除该目录。
//
// 如果给定的`path`不存在或为空，它不做任何操作。
// md5:ae9ec60d038f0ebd
func Remove(path string) (err error) {
	// 如果`path`为空，该函数不执行任何操作。. md5:5ea1be4b22dde448
	if path == "" {
		return nil
	}
	if err = os.RemoveAll(path); err != nil {
		err = gerror.Wrapf(err, `os.RemoveAll failed for path "%s"`, path)
	}
	return
}

// IsReadable 检查给定的 `path` 是否可读。. md5:1b38deb3c4c35233
func IsReadable(path string) bool {
	result := true
	file, err := os.OpenFile(path, os.O_RDONLY, DefaultPermOpen)
	if err != nil {
		result = false
	}
	file.Close()
	return result
}

// IsWritable 检查给定的 `path` 是否可写。
//
// TODO 优化性能；使用 golang.org/x/sys 进行跨平台处理
// md5:2b947cc78310d3f1
func IsWritable(path string) bool {
	result := true
	if IsDir(path) {
		// 如果是一个目录，则创建一个临时文件来测试是否可写。. md5:171566b92b9fb098
		tmpFile := strings.TrimRight(path, Separator) + Separator + gconv.String(time.Now().UnixNano())
		if f, err := Create(tmpFile); err != nil || !Exists(tmpFile) {
			result = false
		} else {
			_ = f.Close()
			_ = Remove(tmpFile)
		}
	} else {
		// 如果它是一个文件，检查是否可以打开它。. md5:48e1a1f6b6b7d3aa
		file, err := os.OpenFile(path, os.O_WRONLY, DefaultPermOpen)
		if err != nil {
			result = false
		}
		_ = file.Close()
	}
	return result
}

// Chmod是os.Chmod的别名。
// 请参阅os.Chmod。
// md5:edb0528fe01cdccd
func Chmod(path string, mode os.FileMode) (err error) {
	err = os.Chmod(path, mode)
	if err != nil {
		err = gerror.Wrapf(err, `os.Chmod failed with path "%s" and mode "%s"`, path, mode)
	}
	return
}

// Abs返回一个路径的绝对表示。
// 如果路径不是绝对路径，它将与当前工作目录连接起来，使其成为一个绝对路径。对于给定的文件，其绝对路径名称并不保证是唯一的。
// Abs会调用Clean方法处理结果。
// md5:4cb8146c59de22da
func Abs(path string) string {
	p, _ := filepath.Abs(path)
	return p
}

// RealPath 将给定的`path`转换为其绝对路径
// 并检查文件路径是否存在。
// 如果文件不存在，则返回空字符串。
// md5:125663d904f8d81e
func RealPath(path string) string {
	p, err := filepath.Abs(path)
	if err != nil {
		return ""
	}
	if !Exists(p) {
		return ""
	}
	return p
}

// SelfPath 返回当前运行进程（二进制文件）的绝对文件路径。. md5:87c861104977f515
func SelfPath() string {
	return selfPath
}

// SelfName 返回当前运行进程（二进制文件）的文件名。. md5:1dea5b20c2c13ef6
func SelfName() string {
	return Basename(SelfPath())
}

// SelfDir返回当前运行进程（二进制文件）的绝对目录路径。. md5:36d8d88a7947606c
func SelfDir() string {
	return filepath.Dir(SelfPath())
}

// Basename 返回路径中的最后一个元素，该元素包含文件扩展名。
// 在提取最后一个元素之前，会移除尾随的路径分隔符。
// 如果路径为空，Base 返回 "."。
// 如果路径完全由分隔符组成，Basename 返回一个单个的分隔符。
//
// 示例：
// Basename("/var/www/file.js") -> file.js
// Basename("file.js")          -> file.js
// md5:0601675e20751381
func Basename(path string) string {
	return filepath.Base(path)
}

// Name 返回路径中不包含文件扩展名的最后一个元素。
//
// 示例：
// Name("/var/www/file.js") -> file
// Name("file.js")          -> file
// md5:231670418efd9216
func Name(path string) string {
	base := filepath.Base(path)
	if i := strings.LastIndexByte(base, '.'); i != -1 {
		return base[:i]
	}
	return base
}

// Dir 函数返回 path 中除最后一个元素之外的所有内容，通常为路径的目录。在丢弃最后一个元素后，Dir 对路径调用 Clean 函数，并移除尾随的斜杠。
// 如果 `path` 为空，Dir 返回"."。
// 如果 `path` 为"."，Dir 将路径视为当前工作目录。
// 如果 `path` 仅由分隔符组成，Dir 返回一个单独的分隔符。
// 返回的路径除非是根目录，否则不会以分隔符结尾。
// 
// 示例：
// Dir("/var/www/file.js") -> "/var/www"
// Dir("file.js")          -> "."
// md5:03710913db229986
func Dir(path string) string {
	if path == "." {
		return filepath.Dir(RealPath(path))
	}
	return filepath.Dir(path)
}

// IsEmpty 检查给定的 `path` 是否为空。
// 如果 `path` 是一个文件夹，它会检查其下是否存在任何文件。
// 如果 `path` 是一个文件，它会检查文件大小是否为零。
// 
// 注意，如果 `path` 不存在，它将返回 true。
// md5:1b96bff377b05eac
func IsEmpty(path string) bool {
	stat, err := Stat(path)
	if err != nil {
		return true
	}
	if stat.IsDir() {
		file, err := os.Open(path)
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

// Ext 返回路径使用的文件名扩展名。
// 扩展名是路径最后一个元素中从最后一个点开始的后缀；
// 如果没有点，则扩展名为空。
// 注意：结果包含符号'.'。
//
// 示例：
// Ext("main.go")  => .go
// Ext("api.json") => .json
// md5:63a13ee69ce09cec
func Ext(path string) string {
	ext := filepath.Ext(path)
	if p := strings.IndexByte(ext, '?'); p != -1 {
		ext = ext[0:p]
	}
	return ext
}

// ExtName 类似于函数 Ext，它返回路径中使用的文件扩展名，
// 但结果不包含符号'.'。
//
// 示例：
// ExtName("main.go")  => "go"
// ExtName("api.json") => "json"
// md5:d508af455375f787
func ExtName(path string) string {
	return strings.TrimLeft(Ext(path), ".")
}

// Temp获取并返回当前系统的临时目录。
//
// 可选参数`names`指定了要与当前系统分隔符连接的子文件夹/子文件，将与路径一起返回。
// md5:8db9471945246517
func Temp(names ...string) string {
	path := os.TempDir()
	for _, name := range names {
		path = Join(path, name)
	}
	return path
}
