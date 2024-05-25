// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gfile_test

import (
	"fmt"
	"os"

	"github.com/gogf/gf/v2/os/gfile"
)

func ExampleMkdir() {
	// init
	var (
		path = gfile.Temp("gfile_example_basic_dir")
	)

	// Creates directory
	gfile.Mkdir(path)

	// 检查目录是否存在. md5:0c502e5e10c3d1bc
	fmt.Println(gfile.IsDir(path))

	// Output:
	// true
}

func ExampleCreate() {
	// init
	var (
		path     = gfile.Join(gfile.Temp("gfile_example_basic_dir"), "file1")
		dataByte = make([]byte, 50)
	)
	// 检查文件是否存在. md5:d1455b2a0aa17f63
	isFile := gfile.IsFile(path)

	fmt.Println(isFile)

	// 递归创建具有给定`path`的文件. md5:587a4af68c8bc5ac
	fileHandle, _ := gfile.Create(path)
	defer fileHandle.Close()

	// 向文件中写入一些内容. md5:856ea5269b5be5ff
	n, _ := fileHandle.WriteString("hello goframe")

	// 检查文件是否存在. md5:d1455b2a0aa17f63
	isFile = gfile.IsFile(path)

	fmt.Println(isFile)

	// 从File中读取len(b)字节. md5:a14d5883b14d9063
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
		path     = gfile.Join(gfile.Temp("gfile_example_basic_dir"), "file1")
		dataByte = make([]byte, 4096)
	)
	// 使用只读模式打开文件或目录. md5:78e9e881c189899d
	file, _ := gfile.Open(path)
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
		path     = gfile.Join(gfile.Temp("gfile_example_basic_dir"), "file1")
		dataByte = make([]byte, 4096)
	)
// 使用自定义的 `flag` 和 `perm` 打开文件/目录
// 如果文件不存在，则创建一个可读写模式的文件，权限默认为 0777
// md5:77b0a10407d251c5
	openFile, _ := gfile.OpenFile(path, os.O_CREATE|os.O_RDWR, gfile.DefaultPermCopy)
	defer openFile.Close()

	// 向文件中写入一些内容. md5:856ea5269b5be5ff
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
		path     = gfile.Join(gfile.Temp("gfile_example_basic_dir"), "file1")
		dataByte = make([]byte, 4096)
	)

// 使用自定义的`flag`打开文件/目录
// 如果文件不存在，将创建文件，并以可读写模式打开，默认的`perm`权限为0666
// md5:510ad8864d50d6b6
	openFile, _ := gfile.OpenWithFlag(path, os.O_CREATE|os.O_RDWR)
	defer openFile.Close()

	// 向文件中写入一些内容. md5:856ea5269b5be5ff
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
		dirPath  = gfile.Temp("gfile_example_basic_dir")
		filePath = "file1"
	)

	// 使用当前系统文件分隔符将字符串数组路径连接起来。. md5:729553e2f763ca20
	joinString := gfile.Join(dirPath, filePath)

	fmt.Println(joinString)

	// May Output:
	// /tmp/gfile_example_basic_dir/file1
}

func ExampleExists() {
	// init
	var (
		path = gfile.Join(gfile.Temp("gfile_example_basic_dir"), "file1")
	)
	// 检查给定的`path`是否存在。. md5:801440e91778729a
	joinString := gfile.Exists(path)

	fmt.Println(joinString)

	// Output:
	// true
}

func ExampleIsDir() {
	// init
	var (
		path     = gfile.Temp("gfile_example_basic_dir")
		filePath = gfile.Join(gfile.Temp("gfile_example_basic_dir"), "file1")
	)
	// 检查给定的`path`是否为目录。. md5:5744f2242b1a0948
	fmt.Println(gfile.IsDir(path))
	fmt.Println(gfile.IsDir(filePath))

	// Output:
	// true
	// false
}

func ExamplePwd() {
	// 获取当前工作目录的绝对路径。. md5:02d8656598c3d01b
	fmt.Println(gfile.Pwd())

	// May Output:
	// xxx/gf/os/gfile
}

func ExampleChdir() {
	// init
	var (
		path = gfile.Join(gfile.Temp("gfile_example_basic_dir"), "file1")
	)
	// 获取当前工作目录. md5:87642df8d64a090c
	fmt.Println(gfile.Pwd())

	// 将当前工作目录更改为指定的目录。. md5:c7ba95b4405caafe
	gfile.Chdir(path)

	// 获取当前工作目录. md5:87642df8d64a090c
	fmt.Println(gfile.Pwd())

	// May Output:
	// xxx/gf/os/gfile
	// /tmp/gfile_example_basic_dir/file1
}

func ExampleIsFile() {
	// init
	var (
		filePath = gfile.Join(gfile.Temp("gfile_example_basic_dir"), "file1")
		dirPath  = gfile.Temp("gfile_example_basic_dir")
	)
	// 检查给定的`path`是否为文件，这意味着它不是目录。. md5:cb0ae2363ad14139
	fmt.Println(gfile.IsFile(filePath))
	fmt.Println(gfile.IsFile(dirPath))

	// Output:
	// true
	// false
}

func ExampleStat() {
	// init
	var (
		path = gfile.Join(gfile.Temp("gfile_example_basic_dir"), "file1")
	)
	// 获取关于指定文件的FileInfo信息。. md5:189ffdaf06730055
	stat, _ := gfile.Stat(path)

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
		srcPath = gfile.Join(gfile.Temp("gfile_example_basic_dir"), "file1")
		dstPath = gfile.Join(gfile.Temp("gfile_example_basic_dir"), "file2")
	)
	// Check is file
	fmt.Println(gfile.IsFile(dstPath))

// 将`src`移动到`dst`路径。
// 如果`dst`已经存在且不是目录，它将被替换。
// md5:3401f06a2c8ccd49
	gfile.Move(srcPath, dstPath)

	fmt.Println(gfile.IsFile(srcPath))
	fmt.Println(gfile.IsFile(dstPath))

	// Output:
	// false
	// false
	// true
}

func ExampleRename() {
	// init
	var (
		srcPath = gfile.Join(gfile.Temp("gfile_example_basic_dir"), "file2")
		dstPath = gfile.Join(gfile.Temp("gfile_example_basic_dir"), "file1")
	)
	// Check is file
	fmt.Println(gfile.IsFile(dstPath))

// 将`src`路径重命名（移动）到`dst`。
// 如果`dst`已经存在且不是目录，它将被替换。
// md5:b028a167dc2de1d0
	gfile.Rename(srcPath, dstPath)

	fmt.Println(gfile.IsFile(srcPath))
	fmt.Println(gfile.IsFile(dstPath))

	// Output:
	// false
	// false
	// true
}

func ExampleDirNames() {
	// init
	var (
		path = gfile.Temp("gfile_example_basic_dir")
	)
	// 获取给定目录`path`下的子文件名。. md5:a7ba80d33218bf78
	dirNames, _ := gfile.DirNames(path)

	fmt.Println(dirNames)

	// May Output:
	// [file1]
}

func ExampleGlob() {
	// init
	var (
		path = gfile.Pwd() + gfile.Separator + "*_example_basic_test.go"
	)
	// 获取给定目录`path`下的子文件名。. md5:a7ba80d33218bf78
	// Only show file name
	matchNames, _ := gfile.Glob(path, true)

	fmt.Println(matchNames)

	// 显示文件的完整路径. md5:d246b83579c32f8a
	matchNames, _ = gfile.Glob(path, false)

	fmt.Println(matchNames)

	// May Output:
	// [gfile_z_example_basic_test.go]
	// [xxx/gf/os/gfile/gfile_z_example_basic_test.go]
}

func ExampleIsReadable() {
	// init
	var (
		path = gfile.Pwd() + gfile.Separator + "testdata/readline/file.log"
	)

	// 检查给定的`path`是否可读。. md5:fda74ad537c20ca3
	fmt.Println(gfile.IsReadable(path))

	// Output:
	// true
}

func ExampleIsWritable() {
	// init
	var (
		path = gfile.Pwd() + gfile.Separator + "testdata/readline/"
		file = "file.log"
	)

	// 检查给定的`path`是否可写。. md5:cbf170ef62b28ee0
	fmt.Println(gfile.IsWritable(path))
	fmt.Println(gfile.IsWritable(path + file))

	// Output:
	// true
	// true
}

func ExampleChmod() {
	// init
	var (
		path = gfile.Join(gfile.Temp("gfile_example_basic_dir"), "file1")
	)

	// 获取关于指定文件的FileInfo信息。. md5:189ffdaf06730055
	stat, err := gfile.Stat(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Show original mode
	fmt.Println(stat.Mode())

	// Change file model
	gfile.Chmod(path, gfile.DefaultPermCopy)

	// 获取关于指定文件的FileInfo信息。. md5:189ffdaf06730055
	stat, _ = gfile.Stat(path)
	// Show the modified mode
	fmt.Println(stat.Mode())

	// Output:
	// -rw-r--r--
	// -rwxr-xr-x
}

func ExampleAbs() {
	// init
	var (
		path = gfile.Join(gfile.Temp("gfile_example_basic_dir"), "file1")
	)

	// 获取path的绝对表示形式。. md5:9e6cadaac30f8871
	fmt.Println(gfile.Abs(path))

	// May Output:
	// /tmp/gfile_example_basic_dir/file1
}

func ExampleRealPath() {
	// init
	var (
		realPath  = gfile.Join(gfile.Temp("gfile_example_basic_dir"), "file1")
		worryPath = gfile.Join(gfile.Temp("gfile_example_basic_dir"), "worryFile")
	)

	// 获取path的绝对表示形式。. md5:cca3127b33ff195c
	fmt.Println(gfile.RealPath(realPath))
	fmt.Println(gfile.RealPath(worryPath))

	// May Output:
	// /tmp/gfile_example_basic_dir/file1
}

func ExampleSelfPath() {

	// 获取当前运行进程的绝对文件路径. md5:976eb91d29aba4fd
	fmt.Println(gfile.SelfPath())

	// May Output:
	// xxx/___github_com_gogf_gf_v2_os_gfile__ExampleSelfPath
}

func ExampleSelfName() {

	// 获取当前正在运行进程的文件名. md5:d2f55580550d36cc
	fmt.Println(gfile.SelfName())

	// May Output:
	// ___github_com_gogf_gf_v2_os_gfile__ExampleSelfName
}

func ExampleSelfDir() {

	// 获取当前运行进程的绝对目录路径. md5:f0b7c37862a2865b
	fmt.Println(gfile.SelfDir())

	// May Output:
	// /private/var/folders/p6/gc_9mm3j229c0mjrjp01gqn80000gn/T
}

func ExampleBasename() {
	// init
	var (
		path = gfile.Pwd() + gfile.Separator + "testdata/readline/file.log"
	)

	// 获取路径中的最后一个元素，该元素包含文件扩展名。. md5:4868d5ea79029f54
	fmt.Println(gfile.Basename(path))

	// Output:
	// file.log
}

func ExampleName() {
	// init
	var (
		path = gfile.Pwd() + gfile.Separator + "testdata/readline/file.log"
	)

	// 获取路径中不包括文件扩展名的最后一个元素。. md5:8291b4d785e21395
	fmt.Println(gfile.Name(path))

	// Output:
	// file
}

func ExampleDir() {
	// init
	var (
		path = gfile.Join(gfile.Temp("gfile_example_basic_dir"), "file1")
	)

	// 获取路径中除最后一个元素外的所有部分，通常是路径的目录部分。. md5:21ab4b575c298060
	fmt.Println(gfile.Dir(path))

	// May Output:
	// /tmp/gfile_example_basic_dir
}

func ExampleIsEmpty() {
	// init
	var (
		path = gfile.Join(gfile.Temp("gfile_example_basic_dir"), "file1")
	)

	// 检查`path`是否为空. md5:87c020da5f9bc2aa
	fmt.Println(gfile.IsEmpty(path))

	// Truncate file
	gfile.Truncate(path, 0)

	// 检查`path`是否为空. md5:87c020da5f9bc2aa
	fmt.Println(gfile.IsEmpty(path))

	// Output:
	// false
	// true
}

func ExampleExt() {
	// init
	var (
		path = gfile.Pwd() + gfile.Separator + "testdata/readline/file.log"
	)

	// 获取path所使用的文件扩展名。. md5:5a50317e9cb8596e
	fmt.Println(gfile.Ext(path))

	// Output:
	// .log
}

func ExampleExtName() {
	// init
	var (
		path = gfile.Pwd() + gfile.Separator + "testdata/readline/file.log"
	)

	// 从路径中获取文件扩展名，但结果不包含'.'符号。. md5:0a63ac6fbba1d676
	fmt.Println(gfile.ExtName(path))

	// Output:
	// log
}

func ExampleTempDir() {
	// init
	var (
		fileName = "gfile_example_basic_dir"
	)

	// 获取path的绝对表示形式。. md5:cca3127b33ff195c
	path := gfile.Temp(fileName)

	fmt.Println(path)

	// May Output:
	// /tmp/gfile_example_basic_dir
}

func ExampleRemove() {
	// init
	var (
		path = gfile.Join(gfile.Temp("gfile_example_basic_dir"), "file1")
	)

	// 检查给定的`path`是否为文件，这意味着它不是目录。. md5:cb0ae2363ad14139
	fmt.Println(gfile.IsFile(path))

	// 使用`path`参数删除所有文件/目录。. md5:8d2699993a255ec6
	gfile.Remove(path)

	// Check again
	fmt.Println(gfile.IsFile(path))

	// Output:
	// true
	// false
}
