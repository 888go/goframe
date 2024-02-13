// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 配置类

import (
	"bytes"
	"context"
	"fmt"
	"os"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gres"
	"github.com/888go/goframe/os/gspath"
	"github.com/888go/goframe/text/gstr"
)

// SetPath 设置配置文件搜索的 `directory` 路径。
// 参数 `path` 可以是绝对路径或相对 `directory` 路径，
// 但强烈建议使用绝对 `directory` 路径。
//
// 注意，此参数是一个目录而非文件的路径。
func (a *AdapterFile) SetPath(directoryPath string) (err error) {
	var (
		isDir    = false
		realPath = ""
	)
	if file := 资源类.Get(directoryPath); file != nil {
		realPath = directoryPath
		isDir = file.FileInfo().IsDir()
	} else {
		// Absolute path.
		realPath = 文件类.X取绝对路径且效验(directoryPath)
		if realPath == "" {
			// Relative path.
			a.searchPaths.X遍历读锁定(func(array []string) {
				for _, v := range array {
					if searchedPath, _ := 文件搜索类.Search(v, directoryPath); searchedPath != "" {
						realPath = searchedPath
						break
					}
				}
			})
		}
		if realPath != "" {
			isDir = 文件类.X是否存在目录(realPath)
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
		return 错误类.X创建(buffer.String())
	}
	// 应该是一个目录。
	if !isDir {
		return 错误类.X创建错误码并格式化(
			错误码类.CodeInvalidParameter,
			`SetPath failed: path "%s" should be directory type`,
			directoryPath,
		)
	}
	// 重复路径检查。
	if a.searchPaths.X查找(realPath) != -1 {
		return nil
	}
	a.jsonMap.X清空()
	a.searchPaths.X清空()
	a.searchPaths.Append别名(realPath)
	intlog.Print(context.TODO(), "SetPath:", realPath)
	return nil
}

// AddPath 将绝对或相对的 `directory` 路径添加到搜索路径中。
//
// 注意，此参数是目录而非文件的路径。
func (a *AdapterFile) AddPath(directoryPaths ...string) (err error) {
	for _, directoryPath := range directoryPaths {
		if err = a.doAddPath(directoryPath); err != nil {
			return err
		}
	}
	return nil
}

// doAddPath 将绝对或相对 `directory` 路径添加到搜索路径中。
func (a *AdapterFile) doAddPath(directoryPath string) (err error) {
	var (
		isDir    = false
		realPath = ""
	)
// 首先检查资源管理器，
// 然后在文件系统中检查路径。
	if file := 资源类.Get(directoryPath); file != nil {
		realPath = directoryPath
		isDir = file.FileInfo().IsDir()
	} else {
		// Absolute path.
		realPath = 文件类.X取绝对路径且效验(directoryPath)
		if realPath == "" {
			// Relative path.
			a.searchPaths.X遍历读锁定(func(array []string) {
				for _, v := range array {
					if searchedPath, _ := 文件搜索类.Search(v, directoryPath); searchedPath != "" {
						realPath = searchedPath
						break
					}
				}
			})
		}
		if realPath != "" {
			isDir = 文件类.X是否存在目录(realPath)
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
		return 错误类.X创建(buffer.String())
	}
	if !isDir {
		return 错误类.X创建错误码并格式化(
			错误码类.CodeInvalidParameter,
			`AddPath failed: path "%s" should be directory type`,
			directoryPath,
		)
	}
	// 重复路径检查。
	if a.searchPaths.X查找(realPath) != -1 {
		return nil
	}
	a.searchPaths.Append别名(realPath)
	intlog.Print(context.TODO(), "AddPath:", realPath)
	return nil
}

// GetPaths 返回当前配置管理器的搜索目录路径数组。
func (a *AdapterFile) GetPaths() []string {
	return a.searchPaths.X取切片()
}

// doGetFilePath 根据传入的文件名 `file` 返回其绝对配置文件路径。
// 若未传入 `file`，则返回默认文件名的配置文件路径。
// 若给定的 `file` 不存在，则返回空字符串 `path` 及错误信息。
func (a *AdapterFile) doGetFilePath(fileName string) (filePath string) {
	var (
		tempPath string
		resFile  *资源类.File
		fileInfo os.FileInfo
	)
	// 搜索资源管理器
	if !资源类.IsEmpty() {
		for _, tryFolder := range resourceTryFolders {
			tempPath = tryFolder + fileName
			if resFile = 资源类.Get(tempPath); resFile != nil {
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
					if resFile = 资源类.Get(tempPath); resFile != nil {
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

	// 正在搜索本地文件系统。
	if filePath == "" {
		// Absolute path.
		if filePath = 文件类.X取绝对路径且效验(fileName); filePath != "" && !文件类.X是否存在目录(filePath) {
			return
		}
		a.searchPaths.X遍历读锁定(func(array []string) {
			for _, searchPath := range array {
				searchPath = 文本类.X过滤尾字符并含空白(searchPath, `\/`)
				for _, tryFolder := range localSystemTryFolders {
					relativePath := 文本类.X过滤尾字符并含空白(
						文件类.X路径生成(tryFolder, fileName),
						`\/`,
					)
					if filePath, _ = 文件搜索类.Search(searchPath, relativePath); filePath != "" &&
						!文件类.X是否存在目录(filePath) {
						return
					}
				}
			}
		})
	}
	return
}

// GetFilePath 函数根据传入的文件名`file`返回该文件的绝对配置文件路径。
// 若未传递`file`参数，则返回默认文件名的配置文件路径。
// 如果给定的`file`不存在，则返回一个空字符串`path`及错误信息。
func (a *AdapterFile) GetFilePath(fileName ...string) (filePath string, err error) {
	var (
		fileExtName  string
		tempFileName string
		usedFileName = a.defaultName
	)
	if len(fileName) > 0 {
		usedFileName = fileName[0]
	}
	fileExtName = 文件类.X路径取扩展名且不含点号(usedFileName)
	if filePath = a.doGetFilePath(usedFileName); (filePath == "" || 文件类.X是否存在目录(filePath)) && !文本类.X数组是否存在(supportedFileTypes, fileExtName) {
// 如果不是使用默认配置，或者其配置文件不可用，
// 则根据名称和所有支持的文件类型搜索可能的配置文件。
		for _, fileType := range supportedFileTypes {
			tempFileName = fmt.Sprintf(`%s.%s`, usedFileName, fileType)
			if filePath = a.doGetFilePath(tempFileName); filePath != "" {
				break
			}
		}
	}
	// 如果无法找到`file`的filePath，它将格式化并返回一个详细的错误。
	if filePath == "" {
		var buffer = bytes.NewBuffer(nil)
		if a.searchPaths.X取长度() > 0 {
			if !文本类.X数组是否存在(supportedFileTypes, fileExtName) {
				buffer.WriteString(fmt.Sprintf(
					`possible config files "%s" or "%s" not found in resource manager or following system searching paths:`,
					usedFileName, fmt.Sprintf(`%s.%s`, usedFileName, 文本类.X连接(supportedFileTypes, "/")),
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
					searchPath = 文本类.X过滤尾字符并含空白(searchPath, `\/`)
					for _, tryFolder := range localSystemTryFolders {
						buffer.WriteString(fmt.Sprintf(
							"\n%d. %s",
							index, 文件类.X路径生成(searchPath, tryFolder),
						))
						index++
					}
				}
			})
		} else {
			buffer.WriteString(fmt.Sprintf(`cannot find config file "%s" with no filePath configured`, usedFileName))
		}
		err = 错误类.X创建错误码(错误码类.CodeNotFound, buffer.String())
	}
	return
}
