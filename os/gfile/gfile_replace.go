// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfile

import (
	"github.com/888go/goframe/text/gstr"
)

// ReplaceFile将文件`path`的内容替换为新内容。
func ReplaceFile(search, replace, path string) error {
	return PutContents(path, gstr.Replace(GetContents(path), search, replace))
}

// ReplaceFileFunc 通过回调函数 `f` 替换文件 `path` 的内容。
func ReplaceFileFunc(f func(path, content string) string, path string) error {
	data := GetContents(path)
	result := f(path, data)
	if len(data) != len(result) && data != result {
		return PutContents(path, result)
	}
	return nil
}

// ReplaceDir 函数用于替换 `path` 路径下文件的内容。
// 参数 `pattern` 指定需要进行替换的文件匹配模式。
// 如果给定的参数 `recursive` 为 true，则会递归地进行替换操作。
func ReplaceDir(search, replace, path, pattern string, recursive ...bool) error {
	files, err := ScanDirFile(path, pattern, recursive...)
	if err != nil {
		return err
	}
	for _, file := range files {
		if err = ReplaceFile(search, replace, file); err != nil {
			return err
		}
	}
	return err
}

// ReplaceDirFunc 函数使用回调函数 `f` 替换路径 `path` 下符合模式 `pattern` 的文件内容。
// 参数 `pattern` 指定需要匹配并替换的文件模式。
// 若给定参数 `recursive` 为真，则会递归地进行替换操作。
func ReplaceDirFunc(f func(path, content string) string, path, pattern string, recursive ...bool) error {
	files, err := ScanDirFile(path, pattern, recursive...)
	if err != nil {
		return err
	}
	for _, file := range files {
		if err = ReplaceFileFunc(f, file); err != nil {
			return err
		}
	}
	return err
}
