// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 配置类

import (
	"bytes"
	"context"
	"fmt"
	"os"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	gfile "github.com/888go/goframe/os/gfile"
	gres "github.com/888go/goframe/os/gres"
	gspath "github.com/888go/goframe/os/gspath"
	gstr "github.com/888go/goframe/text/gstr"
)

// SetPath 设置文件搜索的配置`目录`路径。
// 参数 `path` 可以是绝对或相对的`目录`路径，
// 但强烈建议使用绝对`目录`路径。
//
// 注意，此参数是一个指向目录的路径，而不是指向文件的路径。
// md5:56f162e4bbfc634d
func (a *AdapterFile) SetPath(directoryPath string) (err error) {
	var (
		isDir    = false
		realPath = ""
	)
	if file := gres.Get(directoryPath); file != nil {
		realPath = directoryPath
		isDir = file.FileInfo().IsDir()
	} else {
		// Absolute path.
		realPath = gfile.X取绝对路径且效验(directoryPath)
		if realPath == "" {
			// Relative path.
			a.searchPaths.X遍历读锁定(func(array []string) {
				for _, v := range array {
					if searchedPath, _ := gspath.Search(v, directoryPath); searchedPath != "" {
						realPath = searchedPath
						break
					}
				}
			})
		}
		if realPath != "" {
			isDir = gfile.X是否存在目录(realPath)
		}
	}
	// Path not exist.
	if realPath == "" {
		buffer := bytes.NewBuffer(nil)
		if a.searchPaths.X取长度() > 0 {
			buffer.WriteString(fmt.Sprintf(
				`SetPath failed: cannot find directory "%s" in following paths:`,
				directoryPath,
			))
			a.searchPaths.X遍历读锁定(func(array []string) {
				for k, v := range array {
					buffer.WriteString(fmt.Sprintf("\n%d. %s", k+1, v))
				}
			})
		} else {
			buffer.WriteString(fmt.Sprintf(
				`SetPath failed: path "%s" does not exist`,
				directoryPath,
			))
		}
		return gerror.X创建(buffer.String())
	}
	// Should be a directory.
	if !isDir {
		return gerror.X创建错误码并格式化(
			gcode.CodeInvalidParameter,
			`SetPath failed: path "%s" should be directory type`,
			directoryPath,
		)
	}
	// Repeated path check.
	if a.searchPaths.X查找(realPath) != -1 {
		return nil
	}
	a.jsonMap.X清空()
	a.searchPaths.X清空()
	a.searchPaths.Append别名(realPath)
	intlog.Print(context.TODO(), "SetPath:", realPath)
	return nil
}

// AddPath 向搜索路径中添加一个绝对或相对的`目录`路径。
//
// 请注意，此参数是目录路径，而不是文件路径。
// md5:25c79c7444dc4e16
func (a *AdapterFile) AddPath(directoryPaths ...string) (err error) {
	for _, directoryPath := range directoryPaths {
		if err = a.doAddPath(directoryPath); err != nil {
			return err
		}
	}
	return nil
}

// doAddPath 将绝对或相对的 `directory` 路径添加到搜索路径中。 md5:43115dba5403276a
func (a *AdapterFile) doAddPath(directoryPath string) (err error) {
	var (
		isDir    = false
		realPath = ""
	)
	// 首先检查资源管理器，然后在文件系统中查找路径。
	// md5:deb5a0d060375b57
	if file := gres.Get(directoryPath); file != nil {
		realPath = directoryPath
		isDir = file.FileInfo().IsDir()
	} else {
		// Absolute path.
		realPath = gfile.X取绝对路径且效验(directoryPath)
		if realPath == "" {
			// Relative path.
			a.searchPaths.X遍历读锁定(func(array []string) {
				for _, v := range array {
					if searchedPath, _ := gspath.Search(v, directoryPath); searchedPath != "" {
						realPath = searchedPath
						break
					}
				}
			})
		}
		if realPath != "" {
			isDir = gfile.X是否存在目录(realPath)
		}
	}
	if realPath == "" {
		buffer := bytes.NewBuffer(nil)
		if a.searchPaths.X取长度() > 0 {
			buffer.WriteString(fmt.Sprintf(
				`AddPath failed: cannot find directory "%s" in following paths:`,
				directoryPath,
			))
			a.searchPaths.X遍历读锁定(func(array []string) {
				for k, v := range array {
					buffer.WriteString(fmt.Sprintf("\n%d. %s", k+1, v))
				}
			})
		} else {
			buffer.WriteString(fmt.Sprintf(
				`AddPath failed: path "%s" does not exist`,
				directoryPath,
			))
		}
		return gerror.X创建(buffer.String())
	}
	if !isDir {
		return gerror.X创建错误码并格式化(
			gcode.CodeInvalidParameter,
			`AddPath failed: path "%s" should be directory type`,
			directoryPath,
		)
	}
	// Repeated path check.
	if a.searchPaths.X查找(realPath) != -1 {
		return nil
	}
	a.searchPaths.Append别名(realPath)
	intlog.Print(context.TODO(), "AddPath:", realPath)
	return nil
}

// GetPaths 返回当前配置管理器的搜索目录路径数组。 md5:c77738d1ef96cc99
func (a *AdapterFile) GetPaths() []string {
	return a.searchPaths.X取切片()
}

// doGetFilePath 根据`file`返回绝对配置文件路径。
// 如果未传递`file`，则返回默认名称的配置文件路径。
// 如果给定的`file`不存在，它将返回一个空的`path`字符串和一个错误。
// md5:4044ef5a7532d997
func (a *AdapterFile) doGetFilePath(fileName string) (filePath string) {
	var (
		tempPath string
		resFile  *gres.File
		fileInfo os.FileInfo
	)
		// 在搜索资源管理器。 md5:52083f8252a4c319
	if !gres.IsEmpty() {
		for _, tryFolder := range resourceTryFolders {
			tempPath = tryFolder + fileName
			if resFile = gres.Get(tempPath); resFile != nil {
				fileInfo, _ = resFile.Stat()
				if fileInfo != nil && !fileInfo.IsDir() {
					filePath = resFile.Name()
					return
				}
			}
		}
		a.searchPaths.X遍历读锁定(func(array []string) {
			for _, searchPath := range array {
				for _, tryFolder := range resourceTryFolders {
					tempPath = searchPath + tryFolder + fileName
					if resFile = gres.Get(tempPath); resFile != nil {
						fileInfo, _ = resFile.Stat()
						if fileInfo != nil && !fileInfo.IsDir() {
							filePath = resFile.Name()
							return
						}
					}
				}
			}
		})
	}

	a.autoCheckAndAddMainPkgPathToSearchPaths()

		// 在本地文件系统中搜索。 md5:a557bf6cadf8eec7
	if filePath == "" {
		// Absolute path.
		if filePath = gfile.X取绝对路径且效验(fileName); filePath != "" && !gfile.X是否存在目录(filePath) {
			return
		}
		a.searchPaths.X遍历读锁定(func(array []string) {
			for _, searchPath := range array {
				searchPath = gstr.X过滤尾字符并含空白(searchPath, `\/`)
				for _, tryFolder := range localSystemTryFolders {
					relativePath := gstr.X过滤尾字符并含空白(
						gfile.X路径生成(tryFolder, fileName),
						`\/`,
					)
					if filePath, _ = gspath.Search(searchPath, relativePath); filePath != "" &&
						!gfile.X是否存在目录(filePath) {
						return
					}
				}
			}
		})
	}
	return
}

// GetFilePath 通过 `file` 参数返回给定文件名的绝对配置文件路径。
// 如果没有传递 `file`，则返回默认名称的配置文件路径。
// 如果给定的 `file` 不存在，它将返回一个空的 `path` 字符串和一个错误。
// md5:b116b9d063e12bc9
func (a *AdapterFile) GetFilePath(fileName ...string) (filePath string, err error) {
	var (
		fileExtName  string
		tempFileName string
		usedFileName = a.defaultName
	)
	if len(fileName) > 0 {
		usedFileName = fileName[0]
	}
	fileExtName = gfile.X路径取扩展名且不含点号(usedFileName)
	if filePath = a.doGetFilePath(usedFileName); (filePath == "" || gfile.X是否存在目录(filePath)) && !gstr.X切片是否存在(supportedFileTypes, fileExtName) {
		// 如果它没有使用默认配置，或者其配置文件不可用，
		// 它将根据名称和所有支持的文件类型搜索可能的配置文件。
		// md5:421551127aec1652
		for _, fileType := range supportedFileTypes {
			tempFileName = fmt.Sprintf(`%s.%s`, usedFileName, fileType)
			if filePath = a.doGetFilePath(tempFileName); filePath != "" {
				break
			}
		}
	}
		// 如果无法找到`file`的filePath，它会格式化并返回一个详细的错误。 md5:4aed299684f45971
	if filePath == "" {
		var buffer = bytes.NewBuffer(nil)
		if a.searchPaths.X取长度() > 0 {
			if !gstr.X切片是否存在(supportedFileTypes, fileExtName) {
				buffer.WriteString(fmt.Sprintf(
					`possible config files "%s" or "%s" not found in resource manager or following system searching paths:`,
					usedFileName, fmt.Sprintf(`%s.%s`, usedFileName, gstr.X连接(supportedFileTypes, "/")),
				))
			} else {
				buffer.WriteString(fmt.Sprintf(
					`specified config file "%s" not found in resource manager or following system searching paths:`,
					usedFileName,
				))
			}
			a.searchPaths.X遍历读锁定(func(array []string) {
				index := 1
				for _, searchPath := range array {
					searchPath = gstr.X过滤尾字符并含空白(searchPath, `\/`)
					for _, tryFolder := range localSystemTryFolders {
						buffer.WriteString(fmt.Sprintf(
							"\n%d. %s",
							index, gfile.X路径生成(searchPath, tryFolder),
						))
						index++
					}
				}
			})
		} else {
			buffer.WriteString(fmt.Sprintf(`cannot find config file "%s" with no filePath configured`, usedFileName))
		}
		err = gerror.X创建错误码(gcode.CodeNotFound, buffer.String())
	}
	return
}
