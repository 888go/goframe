// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// gfile 包提供了易于使用的文件系统操作。 md5:51d18e994a768bb4
package 文件类

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	gtype "github.com/888go/goframe/container/gtype"
	gerror "github.com/888go/goframe/errors/gerror"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
)

const (
	// 文件系统分隔符。
	// 这里将分隔符定义为变量，
	// 以便开发人员在必要时进行修改。
	// md5:ec0b6e47ec28478f
	Separator = string(filepath.Separator)

		// DefaultPermOpen 是文件打开的默认权限。 md5:dc57341030d46a11
	DefaultPermOpen = os.FileMode(0666)

		// DefaultPermCopy 是用于文件/文件夹复制的默认权限。 md5:fdef6c133c375aa4
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
		// 初始化内部包变量：selfPath。 md5:8d1168ac7361bb54
	selfPath, _ = exec.LookPath(os.Args[0])
	if selfPath != "" {
		selfPath, _ = filepath.Abs(selfPath)
	}
	if selfPath == "" {
		selfPath, _ = filepath.Abs(os.Args[0])
	}
}

// X创建目录 递归创建给定的 `path` 所表示的目录。建议使用绝对路径而非相对路径作为参数。
// md5:e78abb40a45c2886
func X创建目录(目录 string) (错误 error) {
	if 错误 = os.MkdirAll(目录, os.ModePerm); 错误 != nil {
		错误 = gerror.X多层错误并格式化(错误, `os.MkdirAll failed for path "%s" with perm "%d"`, 目录, os.ModePerm)
		return 错误
	}
	return nil
}

// X创建文件与目录 递归地创建具有给定`path`的文件。
// 建议参数`path`使用绝对路径。
// md5:163accaf36969b42
func X创建文件与目录(文件路径 string) (*os.File, error) {
	dir := X路径取父目录(文件路径)
	if !X是否存在(dir) {
		if err := X创建目录(dir); err != nil {
			return nil, err
		}
	}
	file, err := os.Create(文件路径)
	if err != nil {
		err = gerror.X多层错误并格式化(err, `os.Create failed for name "%s"`, 文件路径)
	}
	return file, err
}

// X打开并按只读模式以只读方式打开文件/目录。 md5:7f50cf0f63b9e34e
func X打开并按只读模式(路径 string) (*os.File, error) {
	file, err := os.Open(路径)
	if err != nil {
		err = gerror.X多层错误并格式化(err, `os.Open failed for name "%s"`, 路径)
	}
	return file, err
}

// X打开 用自定义的 `flag` 和 `perm` 打开文件或目录。参数 `flag` 类似于：O_RDONLY（只读），O_RDWR（读写），O_RDWR|O_CREATE|O_TRUNC 等。
// md5:0cef38d8408ed250
func X打开(路径 string, 读写模式 int, 权限模式 os.FileMode) (*os.File, error) {
	file, err := os.OpenFile(路径, 读写模式, 权限模式)
	if err != nil {
		err = gerror.X多层错误并格式化(err, `os.OpenFile failed with name "%s", flag "%d", perm "%d"`, 路径, 读写模式, 权限模式)
	}
	return file, err
}

// X打开并按默认权限 使用默认权限和自定义`flag`打开文件/目录。默认的`perm`为0666。
// 参数`flag`类似于：O_RDONLY（只读），O_RDWR（读写），O_RDWR|O_CREATE|O_TRUNC（读写并创建或截断），等等。
// md5:2e77d9a0acc43298
func X打开并按默认权限(路径 string, 读写模式 int) (*os.File, error) {
	file, err := X打开(路径, 读写模式, DefaultPermOpen)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// OpenWithFlagPerm别名 使用自定义的 `flag` 和 `perm` 来打开文件/目录。
// 参数 `flag` 例如：O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC 等。
// 参数 `perm` 例如：0600, 0666, 0777 等。
// md5:1f44c05cc68654d8
func OpenWithFlagPerm别名(path string, flag int, perm os.FileMode) (*os.File, error) {
	file, err := X打开(path, flag, perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// X路径生成 使用当前系统的文件分隔符将字符串数组路径连接起来。 md5:349ebcda51de0442
func X路径生成(路径s ...string) string {
	var s string
	for _, path := range 路径s {
		if s != "" {
			s += Separator
		}
		s += gstr.X过滤尾字符并含空白(path, Separator)
	}
	return s
}

// X是否存在检查给定的`path`是否存在。 md5:523f33d374bd2841
func X是否存在(路径 string) bool {
	if stat, err := os.Stat(路径); stat != nil && !os.IsNotExist(err) {
		return true
	}
	return false
}

// X是否存在目录 检查给定的 `path` 是否为目录。
// 注意，如果 `path` 不存在，它将返回 false。
// md5:c5b2468307c9c9e2
func X是否存在目录(路径 string) bool {
	s, err := os.Stat(路径)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// X取当前工作目录 返回当前工作目录的绝对路径。
// 注意，如果获取当前工作目录失败，它将返回一个空字符串。
// md5:90f41f1bfdd61dba
func X取当前工作目录() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path
}

// X设置当前工作目录 将当前工作目录更改为指定的目录。
// 如果发生错误，该错误将为 *PathError 类型。
// md5:7bafb79d47f0c3b3
func X设置当前工作目录(目录 string) (错误 error) {
	错误 = os.Chdir(目录)
	if 错误 != nil {
		错误 = gerror.X多层错误并格式化(错误, `os.Chdir failed with dir "%s"`, 目录)
	}
	return
}

// X是否为文件 检查给定的 `path` 是否为文件，即不是目录。
// 注意，如果 `path` 不存在，它将返回 false。
// md5:38595d733f36d367
func X是否为文件(路径 string) bool {
	s, err := X取详情(路径)
	if err != nil {
		return false
	}
	return !s.IsDir()
}

// X取详情 返回一个描述命名文件的 FileInfo。如果出现错误，错误类型为 *PathError。
// md5:f4ee45de3278c17f
func X取详情(路径 string) (os.FileInfo, error) {
	info, err := os.Stat(路径)
	if err != nil {
		err = gerror.X多层错误并格式化(err, `os.Stat failed for file "%s"`, 路径)
	}
	return info, err
}

// X移动 将路径 `src` 重命名（移动）到 `dst`。
// 如果 `dst` 已经存在且不是一个目录，它将被替换。
// md5:4bc635341db78f64
func X移动(路径 string, 新路径 string) (错误 error) {
	错误 = os.Rename(路径, 新路径)
	if 错误 != nil {
		错误 = gerror.X多层错误并格式化(错误, `os.Rename failed from "%s" to "%s"`, 路径, 新路径)
	}
	return
}

// Rename别名 是 Move 的别名。
// 参考 Move。
// md5:f235456881c23527
func Rename别名(src string, dst string) error {
	return X移动(src, dst)
}

// X取文件列表 返回给定目录 `path` 的子文件名。请注意，返回的名称不是绝对路径。
// md5:62471f4b0e0bb389
func X取文件列表(路径 string) ([]string, error) {
	f, err := X打开并按只读模式(路径)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdirnames(-1)
	_ = f.Close()
	if err != nil {
		err = gerror.X多层错误并格式化(err, `Read dir files failed from path "%s"`, 路径)
		return nil, err
	}
	return list, nil
}

// X模糊查找 返回匹配 pattern 的所有文件名，如果没有匹配的文件，则返回 nil。pattern 的语法与 Match 中的相同。pattern 可以描述像 /usr/*/bin/ed（假设分隔符为 '/'）这样的层次名称。
// 
// X模糊查找 忽略读取目录时发生的文件系统错误，如 I/O 错误。可能返回的唯一错误是 ErrBadPattern，当 pattern 格式不正确时。
// md5:0baeeb8710df5d67
func X模糊查找(路径 string, 返回绝对路径 ...bool) ([]string, error) {
	list, err := filepath.Glob(路径)
	if err != nil {
		err = gerror.X多层错误并格式化(err, `filepath.Glob failed for pattern "%s"`, 路径)
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

// X删除 删除具有`path`参数的所有文件/目录。
// 如果参数`path`是一个目录，它会递归地删除该目录。
//
// 如果给定的`path`不存在或为空，它不做任何操作。
// md5:ae9ec60d038f0ebd
func X删除(路径或文件夹 string) (错误 error) {
		// 如果`path`为空，该函数不执行任何操作。 md5:5ea1be4b22dde448
	if 路径或文件夹 == "" {
		return nil
	}
	if 错误 = os.RemoveAll(路径或文件夹); 错误 != nil {
		错误 = gerror.X多层错误并格式化(错误, `os.RemoveAll failed for path "%s"`, 路径或文件夹)
	}
	return
}

// X是否可读 检查给定的 `path` 是否可读。 md5:1b38deb3c4c35233
func X是否可读(路径 string) bool {
	result := true
	file, err := os.OpenFile(路径, os.O_RDONLY, DefaultPermOpen)
	if err != nil {
		result = false
	}
	file.Close()
	return result
}

// X是否可写 检查给定的 `path` 是否可写。
//
// TODO 优化性能；使用 golang.org/x/sys 进行跨平台处理
// md5:2b947cc78310d3f1
func X是否可写(路径 string) bool {
	result := true
	if X是否存在目录(路径) {
				// 如果是一个目录，则创建一个临时文件来测试是否可写。 md5:171566b92b9fb098
		tmpFile := strings.TrimRight(路径, Separator) + Separator + gconv.String(time.Now().UnixNano())
		if f, err := X创建文件与目录(tmpFile); err != nil || !X是否存在(tmpFile) {
			result = false
		} else {
			_ = f.Close()
			_ = X删除(tmpFile)
		}
	} else {
				// 如果它是一个文件，检查是否可以打开它。 md5:48e1a1f6b6b7d3aa
		file, err := os.OpenFile(路径, os.O_WRONLY, DefaultPermOpen)
		if err != nil {
			result = false
		}
		_ = file.Close()
	}
	return result
}

// X更改权限是os.X更改权限的别名。
// 请参阅os.X更改权限。
// md5:edb0528fe01cdccd
func X更改权限(路径 string, 权限模式 os.FileMode) (错误 error) {
	错误 = os.Chmod(路径, 权限模式)
	if 错误 != nil {
		错误 = gerror.X多层错误并格式化(错误, `os.Chmod failed with path "%s" and mode "%s"`, 路径, 权限模式)
	}
	return
}

// X取绝对路径返回一个路径的绝对表示。
// 如果路径不是绝对路径，它将与当前工作目录连接起来，使其成为一个绝对路径。对于给定的文件，其绝对路径名称并不保证是唯一的。
// X取绝对路径会调用Clean方法处理结果。
// md5:4cb8146c59de22da
func X取绝对路径(路径 string) string {
	p, _ := filepath.Abs(路径)
	return p
}

// X取绝对路径且效验 将给定的`path`转换为其绝对路径
// 并检查文件路径是否存在。
// 如果文件不存在，则返回空字符串。
// md5:125663d904f8d81e
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

// X取当前进程路径 返回当前运行进程（二进制文件）的绝对文件路径。 md5:87c861104977f515
func X取当前进程路径() string {
	return selfPath
}

// X取当前进程名 返回当前运行进程（二进制文件）的文件名。 md5:1dea5b20c2c13ef6
func X取当前进程名() string {
	return X路径取文件名(X取当前进程路径())
}

// X取当前进程目录返回当前运行进程（二进制文件）的绝对目录路径。 md5:36d8d88a7947606c
func X取当前进程目录() string {
	return filepath.Dir(X取当前进程路径())
}

// X路径取文件名 返回路径中的最后一个元素，该元素包含文件扩展名。
// 在提取最后一个元素之前，会移除尾随的路径分隔符。
// 如果路径为空，Base 返回 "."。
// 如果路径完全由分隔符组成，X路径取文件名 返回一个单个的分隔符。
//
// 示例：
// X路径取文件名("/var/www/file.js") -> file.js
// X路径取文件名("file.js")          -> file.js
// md5:0601675e20751381
func X路径取文件名(路径 string) string {
	return filepath.Base(路径)
}

// X路径取文件名且不含扩展名 返回路径中不包含文件扩展名的最后一个元素。
//
// 示例：
// X路径取文件名且不含扩展名("/var/www/file.js") -> file
// X路径取文件名且不含扩展名("file.js")          -> file
// md5:231670418efd9216
func X路径取文件名且不含扩展名(路径 string) string {
	base := filepath.Base(路径)
	if i := strings.LastIndexByte(base, '.'); i != -1 {
		return base[:i]
	}
	return base
}

// X路径取父目录 函数返回 path 中除最后一个元素之外的所有内容，通常为路径的目录。在丢弃最后一个元素后，X路径取父目录 对路径调用 Clean 函数，并移除尾随的斜杠。
// 如果 `path` 为空，X路径取父目录 返回"."。
// 如果 `path` 为"."，X路径取父目录 将路径视为当前工作目录。
// 如果 `path` 仅由分隔符组成，X路径取父目录 返回一个单独的分隔符。
// 返回的路径除非是根目录，否则不会以分隔符结尾。
// 
// 示例：
// X路径取父目录("/var/www/file.js") -> "/var/www"
// X路径取父目录("file.js")          -> "."
// md5:03710913db229986
func X路径取父目录(路径 string) string {
	if 路径 == "." {
		return filepath.Dir(X取绝对路径且效验(路径))
	}
	return filepath.Dir(路径)
}

// X是否为空 检查给定的 `path` 是否为空。
// 如果 `path` 是一个文件夹，它会检查其下是否存在任何文件。
// 如果 `path` 是一个文件，它会检查文件大小是否为零。
// 
// 注意，如果 `path` 不存在，它将返回 true。
// md5:1b96bff377b05eac
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

// X路径取扩展名 返回路径使用的文件名扩展名。
// 扩展名是路径最后一个元素中从最后一个点开始的后缀；
// 如果没有点，则扩展名为空。
// 注意：结果包含符号'.'。
//
// 示例：
// X路径取扩展名("main.go")  => .go
// X路径取扩展名("api.json") => .json
// md5:63a13ee69ce09cec
func X路径取扩展名(路径 string) string {
	ext := filepath.Ext(路径)
	if p := strings.IndexByte(ext, '?'); p != -1 {
		ext = ext[0:p]
	}
	return ext
}

// X路径取扩展名且不含点号 类似于函数 Ext，它返回路径中使用的文件扩展名，
// 但结果不包含符号'.'。
//
// 示例：
// X路径取扩展名且不含点号("main.go")  => "go"
// X路径取扩展名且不含点号("api.json") => "json"
// md5:d508af455375f787
func X路径取扩展名且不含点号(路径 string) string {
	return strings.TrimLeft(X路径取扩展名(路径), ".")
}

// X取临时目录获取并返回当前系统的临时目录。
//
// 可选参数`names`指定了要与当前系统分隔符连接的子文件夹/子文件，将与路径一起返回。
// md5:8db9471945246517
func X取临时目录(可选路径 ...string) string {
	path := os.TempDir()
	for _, name := range 可选路径 {
		path = X路径生成(path, name)
	}
	return path
}
