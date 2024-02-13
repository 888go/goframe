// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_IsDir(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		paths := "/testfile"
		createDir(paths)
		defer delTestFiles(paths)

		t.Assert(文件类.X是否存在目录(testpath()+paths), true)
		t.Assert(文件类.X是否存在目录("./testfile2"), false)
		t.Assert(文件类.X是否存在目录("./testfile/tt.txt"), false)
		t.Assert(文件类.X是否存在目录(""), false)
	})
}

func Test_IsEmpty(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		path := "/testdir_" + 转换类.String(时间类.X取时间戳纳秒())
		createDir(path)
		defer delTestFiles(path)

		t.Assert(文件类.X是否为空(testpath()+path), true)
		t.Assert(文件类.X是否为空(testpath()+path+文件类.Separator+"test.txt"), true)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		path := "/testfile_" + 转换类.String(时间类.X取时间戳纳秒())
		createTestFile(path, "")
		defer delTestFiles(path)

		t.Assert(文件类.X是否为空(testpath()+path), true)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		path := "/testfile_" + 转换类.String(时间类.X取时间戳纳秒())
		createTestFile(path, "1")
		defer delTestFiles(path)

		t.Assert(文件类.X是否为空(testpath()+path), false)
	})
}

func Test_Create(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			err       error
			filepaths []string
			fileobj   *os.File
		)
		filepaths = append(filepaths, "/testfile_cc1.txt")
		filepaths = append(filepaths, "/testfile_cc2.txt")
		for _, v := range filepaths {
			fileobj, err = 文件类.X创建文件与目录(testpath() + v)
			defer delTestFiles(v)
			fileobj.Close()
			t.AssertNil(err)
		}
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		tmpPath := 文件类.X路径生成(文件类.X取临时目录(), "test/testfile_cc1.txt")
		fileobj, err := 文件类.X创建文件与目录(tmpPath)
		defer 文件类.X删除(tmpPath)
		t.AssertNE(fileobj, nil)
		t.AssertNil(err)
		fileobj.Close()
	})
}

func Test_Open(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			err     error
			files   []string
			flags   []bool
			fileobj *os.File
		)

		file1 := "/testfile_nc1.txt"
		createTestFile(file1, "")
		defer delTestFiles(file1)

		files = append(files, file1)
		flags = append(flags, true)

		files = append(files, "./testfile/file1/c1.txt")
		flags = append(flags, false)

		for k, v := range files {
			fileobj, err = 文件类.X打开并按只读模式(testpath() + v)
			fileobj.Close()
			if flags[k] {
				t.AssertNil(err)
			} else {
				t.AssertNE(err, nil)
			}

		}

	})
}

func Test_OpenFile(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			err     error
			files   []string
			flags   []bool
			fileobj *os.File
		)

		files = append(files, "./testfile/file1/nc1.txt")
		flags = append(flags, false)

		f1 := "/testfile_tt.txt"
		createTestFile(f1, "")
		defer delTestFiles(f1)

		files = append(files, f1)
		flags = append(flags, true)

		for k, v := range files {
			fileobj, err = 文件类.X打开(testpath()+v, os.O_RDWR, 0666)
			fileobj.Close()
			if flags[k] {
				t.AssertNil(err)
			} else {
				t.AssertNE(err, nil)
			}

		}

	})
}

func Test_OpenWithFlag(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			err     error
			files   []string
			flags   []bool
			fileobj *os.File
		)

		file1 := "/testfile_t1.txt"
		createTestFile(file1, "")
		defer delTestFiles(file1)
		files = append(files, file1)
		flags = append(flags, true)

		files = append(files, "/testfiless/dirfiles/t1_no.txt")
		flags = append(flags, false)

		for k, v := range files {
			fileobj, err = 文件类.X打开并按默认权限(testpath()+v, os.O_RDWR)
			fileobj.Close()
			if flags[k] {
				t.AssertNil(err)
			} else {
				t.AssertNE(err, nil)
			}

		}

	})
}

func Test_OpenWithFlagPerm(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			err     error
			files   []string
			flags   []bool
			fileobj *os.File
		)
		file1 := "/testfile_nc1.txt"
		createTestFile(file1, "")
		defer delTestFiles(file1)
		files = append(files, file1)
		flags = append(flags, true)

		files = append(files, "/testfileyy/tt.txt")
		flags = append(flags, false)

		for k, v := range files {
			fileobj, err = 文件类.OpenWithFlagPerm别名(testpath()+v, os.O_RDWR, 0666)
			fileobj.Close()
			if flags[k] {
				t.AssertNil(err)
			} else {
				t.AssertNE(err, nil)
			}

		}

	})
}

func Test_Exists(t *testing.T) {

	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			flag  bool
			files []string
			flags []bool
		)

		file1 := "/testfile_GetContents.txt"
		createTestFile(file1, "")
		defer delTestFiles(file1)

		files = append(files, file1)
		flags = append(flags, true)

		files = append(files, "./testfile/havefile1/tt_no.txt")
		flags = append(flags, false)

		for k, v := range files {
			flag = 文件类.X是否存在(testpath() + v)
			if flags[k] {
				t.Assert(flag, true)
			} else {
				t.Assert(flag, false)
			}

		}

	})
}

func Test_Pwd(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		paths, err := os.Getwd()
		t.AssertNil(err)
		t.Assert(文件类.X取当前工作目录(), paths)

	})
}

func Test_IsFile(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			flag  bool
			files []string
			flags []bool
		)

		file1 := "/testfile_tt.txt"
		createTestFile(file1, "")
		defer delTestFiles(file1)
		files = append(files, file1)
		flags = append(flags, true)

		dir1 := "/testfiless"
		createDir(dir1)
		defer delTestFiles(dir1)
		files = append(files, dir1)
		flags = append(flags, false)

		files = append(files, "./testfiledd/tt1.txt")
		flags = append(flags, false)

		for k, v := range files {
			flag = 文件类.X是否为文件(testpath() + v)
			if flags[k] {
				t.Assert(flag, true)
			} else {
				t.Assert(flag, false)
			}

		}

	})
}

func Test_Info(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			err    error
			paths  string = "/testfile_t1.txt"
			files  os.FileInfo
			files2 os.FileInfo
		)

		createTestFile(paths, "")
		defer delTestFiles(paths)
		files, err = 文件类.X取详情(testpath() + paths)
		t.AssertNil(err)

		files2, err = os.Stat(testpath() + paths)
		t.AssertNil(err)

		t.Assert(files, files2)

	})
}

func Test_Move(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			paths     string = "/ovetest"
			filepaths string = "/testfile_ttn1.txt"
			topath    string = "/testfile_ttn2.txt"
		)
		createDir("/ovetest")
		createTestFile(paths+filepaths, "a")

		defer delTestFiles(paths)

		yfile := testpath() + paths + filepaths
		tofile := testpath() + paths + topath

		t.Assert(文件类.X移动(yfile, tofile), nil)

		// 检查移动后的文件是否真实存在
		_, err := os.Stat(tofile)
		t.Assert(os.IsNotExist(err), false)

	})
}

func Test_Rename(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			paths  string = "/testfiles"
			ypath  string = "/testfilettm1.txt"
			topath string = "/testfilettm2.txt"
		)
		createDir(paths)
		createTestFile(paths+ypath, "a")
		defer delTestFiles(paths)

		ypath = testpath() + paths + ypath
		topath = testpath() + paths + topath

		t.Assert(文件类.Rename别名(ypath, topath), nil)
		t.Assert(文件类.X是否为文件(topath), true)

		t.AssertNE(文件类.Rename别名("", ""), nil)

	})

}

func Test_DirNames(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			paths    string = "/testdirs"
			err      error
			readlist []string
		)
		havelist := []string{
			"t1.txt",
			"t2.txt",
		}

		// 创建测试文件
		createDir(paths)
		for _, v := range havelist {
			createTestFile(paths+"/"+v, "")
		}
		defer delTestFiles(paths)

		readlist, err = 文件类.X取文件列表(testpath() + paths)

		t.AssertNil(err)
		t.AssertIN(readlist, havelist)

		_, err = 文件类.X取文件列表("")
		t.AssertNE(err, nil)

	})
}

func Test_Glob(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			paths      string = "/testfiles/*.txt"
			dirpath    string = "/testfiles"
			err        error
			resultlist []string
		)

		havelist1 := []string{
			"t1.txt",
			"t2.txt",
		}

		havelist2 := []string{
			testpath() + "/testfiles/t1.txt",
			testpath() + "/testfiles/t2.txt",
		}

		// ===============================构建测试文件
		createDir(dirpath)
		for _, v := range havelist1 {
			createTestFile(dirpath+"/"+v, "")
		}
		defer delTestFiles(dirpath)

		resultlist, err = 文件类.X模糊查找(testpath()+paths, true)
		t.AssertNil(err)
		t.Assert(resultlist, havelist1)

		resultlist, err = 文件类.X模糊查找(testpath()+paths, false)

		t.AssertNil(err)
		t.Assert(formatpaths(resultlist), formatpaths(havelist2))

		_, err = 文件类.X模糊查找("", true)
		t.AssertNil(err)

	})
}

func Test_Remove(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			paths string = "/testfile_t1.txt"
		)
		createTestFile(paths, "")
		t.Assert(文件类.X删除(testpath()+paths), nil)

		t.Assert(文件类.X删除(""), nil)

		defer delTestFiles(paths)

	})
}

func Test_IsReadable(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			paths1 string = "/testfile_GetContents.txt"
			paths2 string = "./testfile_GetContents_no.txt"
		)

		createTestFile(paths1, "")
		defer delTestFiles(paths1)

		t.Assert(文件类.X是否可读(testpath()+paths1), true)
		t.Assert(文件类.X是否可读(paths2), false)

	})
}

func Test_IsWritable(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			paths1 string = "/testfile_GetContents.txt"
			paths2 string = "./testfile_GetContents_no.txt"
		)

		createTestFile(paths1, "")
		defer delTestFiles(paths1)
		t.Assert(文件类.X是否可写(testpath()+paths1), true)
		t.Assert(文件类.X是否可写(paths2), false)

	})
}

func Test_Chmod(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			paths1 string = "/testfile_GetContents.txt"
			paths2 string = "./testfile_GetContents_no.txt"
		)
		createTestFile(paths1, "")
		defer delTestFiles(paths1)

		t.Assert(文件类.X更改权限(testpath()+paths1, 0777), nil)
		t.AssertNE(文件类.X更改权限(paths2, 0777), nil)

	})
}

// 获取绝对目录地址
func Test_RealPath(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			paths1    string = "/testfile_files"
			readlPath string

			tempstr string
		)

		createDir(paths1)
		defer delTestFiles(paths1)

		readlPath = 文件类.X取绝对路径且效验("./")

		tempstr, _ = filepath.Abs("./")

		t.Assert(readlPath, tempstr)

		t.Assert(文件类.X取绝对路径且效验("./nodirs"), "")

	})
}

// 获取当前执行文件的目录
func Test_SelfPath(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			paths1    string
			readlPath string
			tempstr   string
		)
		readlPath = 文件类.X取当前进程路径()
		readlPath = filepath.ToSlash(readlPath)

		tempstr, _ = filepath.Abs(os.Args[0])
		paths1 = filepath.ToSlash(tempstr)
		paths1 = strings.Replace(paths1, "./", "/", 1)

		t.Assert(readlPath, paths1)

	})
}

func Test_SelfDir(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			paths1    string
			readlPath string
			tempstr   string
		)
		readlPath = 文件类.X取当前进程目录()

		tempstr, _ = filepath.Abs(os.Args[0])
		paths1 = filepath.Dir(tempstr)

		t.Assert(readlPath, paths1)

	})
}

func Test_Basename(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			paths1    string = "/testfilerr_GetContents.txt"
			readlPath string
		)

		createTestFile(paths1, "")
		defer delTestFiles(paths1)

		readlPath = 文件类.X路径取文件名(testpath() + paths1)
		t.Assert(readlPath, "testfilerr_GetContents.txt")

	})
}

func Test_Dir(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			paths1    string = "/testfiless"
			readlPath string
		)
		createDir(paths1)
		defer delTestFiles(paths1)

		readlPath = 文件类.X路径取父目录(testpath() + paths1)

		t.Assert(readlPath, testpath())

		t.Assert(len(文件类.X路径取父目录(".")) > 0, true)
	})
}

func Test_Ext(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			paths1   string = "/testfile_GetContents.txt"
			dirpath1        = "/testdirs"
		)
		createTestFile(paths1, "")
		defer delTestFiles(paths1)

		createDir(dirpath1)
		defer delTestFiles(dirpath1)

		t.Assert(文件类.X路径取扩展名(testpath()+paths1), ".txt")
		t.Assert(文件类.X路径取扩展名(testpath()+dirpath1), "")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文件类.X路径取扩展名("/var/www/test.js"), ".js")
		t.Assert(文件类.X路径取扩展名("/var/www/test.min.js"), ".js")
		t.Assert(文件类.X路径取扩展名("/var/www/test.js?1"), ".js")
		t.Assert(文件类.X路径取扩展名("/var/www/test.min.js?v1"), ".js")
	})
}

func Test_ExtName(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文件类.X路径取扩展名且不含点号("/var/www/test.js"), "js")
		t.Assert(文件类.X路径取扩展名且不含点号("/var/www/test.min.js"), "js")
		t.Assert(文件类.X路径取扩展名且不含点号("/var/www/test.js?v=1"), "js")
		t.Assert(文件类.X路径取扩展名且不含点号("/var/www/test.min.js?v=1"), "js")
	})
}

func Test_TempDir(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文件类.X取临时目录(), os.TempDir())
	})
}

func Test_Mkdir(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			tpath string = "/testfile/createdir"
			err   error
		)

		defer delTestFiles("/testfile")

		err = 文件类.X创建目录(testpath() + tpath)
		t.AssertNil(err)

		err = 文件类.X创建目录("")
		t.AssertNE(err, nil)

		err = 文件类.X创建目录(testpath() + tpath + "2/t1")
		t.AssertNil(err)

	})
}

func Test_Stat(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			tpath1   = "/testfile_t1.txt"
			tpath2   = "./testfile_t1_no.txt"
			err      error
			fileiofo os.FileInfo
		)

		createTestFile(tpath1, "a")
		defer delTestFiles(tpath1)

		fileiofo, err = 文件类.X取详情(testpath() + tpath1)
		t.AssertNil(err)

		t.Assert(fileiofo.Size(), 1)

		_, err = 文件类.X取详情(tpath2)
		t.AssertNE(err, nil)

	})
}

func Test_MainPkgPath(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		reads := 文件类.X取main路径()
		t.Assert(reads, "")
	})
}

func Test_SelfName(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(len(文件类.X取当前进程名()) > 0, true)
	})
}

func Test_MTimestamp(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文件类.X取修改时间戳秒(文件类.X取临时目录()) > 0, true)
	})
}
