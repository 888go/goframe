// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gfile

import (
	"github.com/gogf/gf/v2/text/gstr"
)

// ReplaceFile 替换文件 `path` 的内容。 md5:70e99046a619416c
// ff:子文本替换
// search:欲被替换的子文本
// replace:用作替换的子文本
// path:文件路径
func ReplaceFile(search, replace, path string) error {
	return PutContents(path, gstr.Replace(GetContents(path), search, replace))
}

// ReplaceFileFunc 使用回调函数 `f` 替换文件 `path` 中的内容。 md5:033d4157195d29fc
// ff:子文本替换_函数
// f:回调函数
// path:文件路径
// content:
// path:文件路径
func ReplaceFileFunc(f func(path, content string) string, path string) error {
	data := GetContents(path)
	result := f(path, data)
	if len(data) != len(result) || data != result {
		return PutContents(path, result)
	}
	return nil
}

// ReplaceDir 替换给定路径下文件的内容。
// 参数 `pattern` 指定了需要被替换的文件模式。
// 如果参数 `recursive` 为真，它将递归地进行替换操作。
// md5:d9b10978d6db3bce
// ff:目录子文本替换
// search:欲被替换的子文本
// replace:用作替换的子文本
// path:目录
// pattern:匹配文件模式
// recursive:是否递归替换
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

// ReplaceDirFunc 使用回调函数 `f` 替换路径 `path` 下的文件内容。
// 参数 `pattern` 指定了将被替换的文件模式。
// 如果 `recursive` 为 true，它会递归地进行替换。
// md5:9bff4662f6d662c1
// ff:目录子文本替换_函数
// f:回调函数
// path:目录
// content:
// path:目录
// pattern:匹配文件模式
// recursive:是否递归替换
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
