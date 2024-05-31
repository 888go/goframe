// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gfile

import (
	"github.com/gogf/gf/v2/text/gstr"
)

// ReplaceFile replaces content for file `path`.

// ff:子文本替换
// path:文件路径
// replace:用作替换的子文本
// search:欲被替换的子文本
func ReplaceFile(search, replace, path string) error {
	return PutContents(path, gstr.Replace(GetContents(path), search, replace))
}

// ReplaceFileFunc replaces content for file `path` with callback function `f`.

// ff:子文本替换_函数
// path:文件路径
// f:回调函数
// content:
// path:
func ReplaceFileFunc(f func(path, content string) string, path string) error {
	data := GetContents(path)
	result := f(path, data)
	if len(data) != len(result) || data != result {
		return PutContents(path, result)
	}
	return nil
}

// ReplaceDir replaces content for files under `path`.
// The parameter `pattern` specifies the file pattern which matches to be replaced.
// It does replacement recursively if given parameter `recursive` is true.

// ff:目录子文本替换
// recursive:是否递归替换
// pattern:匹配文件模式
// path:目录
// replace:用作替换的子文本
// search:欲被替换的子文本
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

// ReplaceDirFunc replaces content for files under `path` with callback function `f`.
// The parameter `pattern` specifies the file pattern which matches to be replaced.
// It does replacement recursively if given parameter `recursive` is true.

// ff:目录子文本替换_函数
// recursive:是否递归替换
// pattern:匹配文件模式
// path:目录
// f:回调函数
// content:
// path:
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
