// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类_test

import (
	"fmt"
	"os"
	
	"github.com/888go/goframe/gfile"
)

func ExampleMkdir() {
	// init
	var (
		path = 文件类.X取临时目录("gfile_example_basic_dir")
	)

	// Creates directory
	文件类.X创建目录(path)

	// 检查目录是否存在
	fmt.Println(文件类.X是否存在目录(path))

	// Output:
	// true
}

func ExampleCreate() {
	// init
	var (
		path     = 文件类.X路径生成(文件类.X取临时目录("gfile_example_basic_dir"), "file1")
		dataByte = make([]byte, 50)
	)
	// 检查文件是否存在
	isFile := 文件类.X是否为文件(path)

	fmt.Println(isFile)

	// 根据给定的`path`递归创建文件
	fileHandle, _ := 文件类.X创建文件与目录(path)
	defer fileHandle.Close()

	// 向文件写入一些内容
	n, _ := fileHandle.WriteString("hello goframe")

	// 检查文件是否存在
	isFile = 文件类.X是否为文件(path)

	fmt.Println(isFile)

	// 从File中读取len(b)个字节
	fileHandle.ReadAt(dataByte, 0)

	fmt.Println(string(dataByte[:n]))

	// Output:
	// false
	// true
	// hello goframe
}

func ExampleOpen() {
	// init
	var (
		path     = 文件类.X路径生成(文件类.X取临时目录("gfile_example_basic_dir"), "file1")
		dataByte = make([]byte, 4096)
	)
	// 以只读模式打开文件或目录
	file, _ := 文件类.X打开并按只读模式(path)
	defer file.Close()

	// Read data
	n, _ := file.Read(dataByte)

	fmt.Println(string(dataByte[:n]))

	// Output:
	// hello goframe
}

func ExampleOpenFile() {
	// init
	var (
		path     = 文件类.X路径生成(文件类.X取临时目录("gfile_example_basic_dir"), "file1")
		dataByte = make([]byte, 4096)
	)
// 通过自定义的`flag`和`perm`打开文件/目录
// 如果文件不存在，则创建，创建时默认为可读写模式，权限为0777
	openFile, _ := 文件类.X打开(path, os.O_CREATE|os.O_RDWR, 文件类.DefaultPermCopy)
	defer openFile.Close()

	// 向文件写入一些内容
	writeLength, _ := openFile.WriteString("hello goframe test open file")

	fmt.Println(writeLength)

	// Read data
	n, _ := openFile.ReadAt(dataByte, 0)

	fmt.Println(string(dataByte[:n]))

	// Output:
	// 28
	// hello goframe test open file
}

func ExampleOpenWithFlag() {
	// init
	var (
		path     = 文件类.X路径生成(文件类.X取临时目录("gfile_example_basic_dir"), "file1")
		dataByte = make([]byte, 4096)
	)

// 使用自定义`flag`打开文件/目录
// 如果文件不存在，则创建，以默认权限`perm`为0666的可读写模式创建
	openFile, _ := 文件类.X打开并按默认权限(path, os.O_CREATE|os.O_RDWR)
	defer openFile.Close()

	// 向文件写入一些内容
	writeLength, _ := openFile.WriteString("hello goframe test open file with flag")

	fmt.Println(writeLength)

	// Read data
	n, _ := openFile.ReadAt(dataByte, 0)

	fmt.Println(string(dataByte[:n]))

	// Output:
	// 38
	// hello goframe test open file with flag
}

func ExampleJoin() {
	// init
	var (
		dirPath  = 文件类.X取临时目录("gfile_example_basic_dir")
		filePath = "file1"
	)

	// 使用当前系统的文件分隔符连接字符串数组路径。
	joinString := 文件类.X路径生成(dirPath, filePath)

	fmt.Println(joinString)

	// May Output:
	// /tmp/gfile_example_basic_dir/file1
}

func ExampleExists() {
	// init
	var (
		path = 文件类.X路径生成(文件类.X取临时目录("gfile_example_basic_dir"), "file1")
	)
	// 检查给定的 `path` 是否存在。
	joinString := 文件类.X是否存在(path)

	fmt.Println(joinString)

	// Output:
	// true
}

func ExampleIsDir() {
	// init
	var (
		path     = 文件类.X取临时目录("gfile_example_basic_dir")
		filePath = 文件类.X路径生成(文件类.X取临时目录("gfile_example_basic_dir"), "file1")
	)
	// 检查给定的`path`是否为一个目录。
	fmt.Println(文件类.X是否存在目录(path))
	fmt.Println(文件类.X是否存在目录(filePath))

	// Output:
	// true
	// false
}

func ExamplePwd() {
	// 获取当前工作目录的绝对路径。
	fmt.Println(文件类.X取当前工作目录())

	// May Output:
	// xxx/gf/os/gfile
}

func ExampleChdir() {
	// init
	var (
		path = 文件类.X路径生成(文件类.X取临时目录("gfile_example_basic_dir"), "file1")
	)
	// Get current working directory
	fmt.Println(文件类.X取当前工作目录())

	// 将当前工作目录更改为指定的目录。
	文件类.X设置当前工作目录(path)

	// Get current working directory
	fmt.Println(文件类.X取当前工作目录())

	// May Output:
	// xxx/gf/os/gfile
	// /tmp/gfile_example_basic_dir/file1
}

func ExampleIsFile() {
	// init
	var (
		filePath = 文件类.X路径生成(文件类.X取临时目录("gfile_example_basic_dir"), "file1")
		dirPath  = 文件类.X取临时目录("gfile_example_basic_dir")
	)
	// 检查给定的`path`是否为文件，也就是说它不是一个目录。
	fmt.Println(文件类.X是否为文件(filePath))
	fmt.Println(文件类.X是否为文件(dirPath))

	// Output:
	// true
	// false
}

func ExampleStat() {
	// init
	var (
		path = 文件类.X路径生成(文件类.X取临时目录("gfile_example_basic_dir"), "file1")
	)
	// 获取描述指定文件的FileInfo对象。
	stat, _ := 文件类.X取详情(path)

	fmt.Println(stat.Name())
	fmt.Println(stat.IsDir())
	fmt.Println(stat.Mode())
	fmt.Println(stat.ModTime())
	fmt.Println(stat.Size())
	fmt.Println(stat.Sys())

	// May Output:
	// file1
	// false
	// -rwxr-xr-x
	// 2021-12-02 11:01:27.261441694 +0800 CST
	// &{16777220 33261 1 8597857090 501 20 0 [0 0 0 0] {1638414088 192363490} {1638414087 261441694} {1638414087 261441694} {1638413480 485068275} 38 8 4096 0 0 0 [0 0]}
}

func ExampleMove() {
	// init
	var (
		srcPath = 文件类.X路径生成(文件类.X取临时目录("gfile_example_basic_dir"), "file1")
		dstPath = 文件类.X路径生成(文件类.X取临时目录("gfile_example_basic_dir"), "file2")
	)
	// Check is file
	fmt.Println(文件类.X是否为文件(dstPath))

// 将`src`移动到`dst`路径。
// 如果`dst`已存在且不是一个目录，它将会被替换。
	文件类.X移动(srcPath, dstPath)

	fmt.Println(文件类.X是否为文件(srcPath))
	fmt.Println(文件类.X是否为文件(dstPath))

	// Output:
	// false
	// false
	// true
}

func ExampleRename() {
	// init
	var (
		srcPath = 文件类.X路径生成(文件类.X取临时目录("gfile_example_basic_dir"), "file2")
		dstPath = 文件类.X路径生成(文件类.X取临时目录("gfile_example_basic_dir"), "file1")
	)
	// Check is file
	fmt.Println(文件类.X是否为文件(dstPath))

// 将`src`重命名为（移动到）`dst`路径。
// 如果`dst`已存在且不是一个目录，则会被替换。
	文件类.Rename别名(srcPath, dstPath)

	fmt.Println(文件类.X是否为文件(srcPath))
	fmt.Println(文件类.X是否为文件(dstPath))

	// Output:
	// false
	// false
	// true
}

func ExampleDirNames() {
	// init
	var (
		path = 文件类.X取临时目录("gfile_example_basic_dir")
	)
	// 获取给定目录 `path` 下的子文件名列表。
	dirNames, _ := 文件类.X取文件列表(path)

	fmt.Println(dirNames)

	// May Output:
	// [file1]
}

func ExampleGlob() {
	// init
	var (
		path = 文件类.X取当前工作目录() + 文件类.Separator + "*_example_basic_test.go"
	)
	// 获取给定目录 `path` 下的子文件名列表。
	// Only show file name
	matchNames, _ := 文件类.X模糊查找(path, true)

	fmt.Println(matchNames)

	// 显示文件的完整路径
	matchNames, _ = 文件类.X模糊查找(path, false)

	fmt.Println(matchNames)

	// May Output:
	// [gfile_z_example_basic_test.go]
	// [xxx/gf/os/gfile/gfile_z_example_basic_test.go]
}

func ExampleIsReadable() {
	// init
	var (
		path = 文件类.X取当前工作目录() + 文件类.Separator + "testdata/readline/file.log"
	)

	// 检查给定的`path`是否可读。
	fmt.Println(文件类.X是否可读(path))

	// Output:
	// true
}

func ExampleIsWritable() {
	// init
	var (
		path = 文件类.X取当前工作目录() + 文件类.Separator + "testdata/readline/"
		file = "file.log"
	)

	// 检查给定的`path`是否可写。
	fmt.Println(文件类.X是否可写(path))
	fmt.Println(文件类.X是否可写(path + file))

	// Output:
	// true
	// true
}

func ExampleChmod() {
	// init
	var (
		path = 文件类.X路径生成(文件类.X取临时目录("gfile_example_basic_dir"), "file1")
	)

	// 获取描述指定文件的FileInfo对象。
	stat, err := 文件类.X取详情(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 显示原始模式
	fmt.Println(stat.Mode())

	// Change file model
	文件类.X更改权限(path, 文件类.DefaultPermCopy)

	// 获取描述指定文件的FileInfo对象。
	stat, _ = 文件类.X取详情(path)
	// 显示修改后的模式
	fmt.Println(stat.Mode())

	// Output:
	// -rw-r--r--
	// -rwxr-xr-x
}

func ExampleAbs() {
	// init
	var (
		path = 文件类.X路径生成(文件类.X取临时目录("gfile_example_basic_dir"), "file1")
	)

	// 获取path的绝对表示形式。
	fmt.Println(文件类.X取绝对路径(path))

	// May Output:
	// /tmp/gfile_example_basic_dir/file1
}

func ExampleRealPath() {
	// init
	var (
		realPath  = 文件类.X路径生成(文件类.X取临时目录("gfile_example_basic_dir"), "file1")
		worryPath = 文件类.X路径生成(文件类.X取临时目录("gfile_example_basic_dir"), "worryFile")
	)

	// fetch an absolute representation of path.
	fmt.Println(文件类.X取绝对路径且效验(realPath))
	fmt.Println(文件类.X取绝对路径且效验(worryPath))

	// May Output:
	// /tmp/gfile_example_basic_dir/file1
}

func ExampleSelfPath() {

	// 获取当前运行进程的绝对文件路径
	fmt.Println(文件类.X取当前进程路径())

	// May Output:
	// xxx/___github_com_gogf_gf_v2_os_gfile__ExampleSelfPath
}

func ExampleSelfName() {

	// 获取当前运行进程的文件名
	fmt.Println(文件类.X取当前进程名())

	// May Output:
	// ___github_com_gogf_gf_v2_os_gfile__ExampleSelfName
}

func ExampleSelfDir() {

	// 获取当前运行进程的绝对目录路径
	fmt.Println(文件类.X取当前进程目录())

	// May Output:
	// /private/var/folders/p6/gc_9mm3j229c0mjrjp01gqn80000gn/T
}

func ExampleBasename() {
	// init
	var (
		path = 文件类.X取当前工作目录() + 文件类.Separator + "testdata/readline/file.log"
	)

	// 获取路径中的最后一个元素，该元素包含文件扩展名。
	fmt.Println(文件类.X路径取文件名(path))

	// Output:
	// file.log
}

func ExampleName() {
	// init
	var (
		path = 文件类.X取当前工作目录() + 文件类.Separator + "testdata/readline/file.log"
	)

	// 获取路径中最后一个元素，不包括文件扩展名。
	fmt.Println(文件类.X路径取文件名且不含扩展名(path))

	// Output:
	// file
}

func ExampleDir() {
	// init
	var (
		path = 文件类.X路径生成(文件类.X取临时目录("gfile_example_basic_dir"), "file1")
	)

	// 获取路径中除最后一个元素之外的所有元素，通常是指路径的目录部分。
	fmt.Println(文件类.X路径取父目录(path))

	// May Output:
	// /tmp/gfile_example_basic_dir
}

func ExampleIsEmpty() {
	// init
	var (
		path = 文件类.X路径生成(文件类.X取临时目录("gfile_example_basic_dir"), "file1")
	)

	// 检查 `path` 是否为空
	fmt.Println(文件类.X是否为空(path))

	// Truncate file
	文件类.X截断(path, 0)

	// 检查 `path` 是否为空
	fmt.Println(文件类.X是否为空(path))

	// Output:
	// false
	// true
}

func ExampleExt() {
	// init
	var (
		path = 文件类.X取当前工作目录() + 文件类.Separator + "testdata/readline/file.log"
	)

	// 获取路径中使用的文件名扩展名。
	fmt.Println(文件类.X路径取扩展名(path))

	// Output:
	// .log
}

func ExampleExtName() {
	// init
	var (
		path = 文件类.X取当前工作目录() + 文件类.Separator + "testdata/readline/file.log"
	)

	// 获取路径path所使用的文件名扩展名，但结果中不包含符号'.'。
	fmt.Println(文件类.X路径取扩展名且不含点号(path))

	// Output:
	// log
}

func ExampleTempDir() {
	// init
	var (
		fileName = "gfile_example_basic_dir"
	)

	// fetch an absolute representation of path.
	path := 文件类.X取临时目录(fileName)

	fmt.Println(path)

	// May Output:
	// /tmp/gfile_example_basic_dir
}

func ExampleRemove() {
	// init
	var (
		path = 文件类.X路径生成(文件类.X取临时目录("gfile_example_basic_dir"), "file1")
	)

	// 检查给定的`path`是否为文件，也就是说它不是一个目录。
	fmt.Println(文件类.X是否为文件(path))

	// 删除具有`path`参数的所有文件/目录。
	文件类.X删除(path)

	// Check again
	fmt.Println(文件类.X是否为文件(path))

	// Output:
	// true
	// false
}
