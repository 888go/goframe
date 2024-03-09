// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类

import (
	"github.com/gogf/gf/v2/text/gstr"
)

// ReplaceFile将文件`path`的内容替换为新内容。
func X子文本替换(欲被替换的子文本, 用作替换的子文本, 文件路径 string) error {
	return X写入文本(文件路径, gstr.Replace(X读文本(文件路径), 欲被替换的子文本, 用作替换的子文本))
}

// ReplaceFileFunc 通过回调函数 `f` 替换文件 `path` 的内容。
func X子文本替换_函数(回调函数 func(路径, 内容 string) string, 文件路径 string) error {
	data := X读文本(文件路径)
	result := 回调函数(文件路径, data)
	if len(data) != len(result) && data != result {
		return X写入文本(文件路径, result)
	}
	return nil
}

// ReplaceDir 函数用于替换 `path` 路径下文件的内容。
// 参数 `pattern` 指定需要进行替换的文件匹配模式。
// 如果给定的参数 `recursive` 为 true，则会递归地进行替换操作。
func X目录子文本替换(欲被替换的子文本, 用作替换的子文本, 目录, 匹配文件模式 string, 是否递归替换 ...bool) error {
	files, err := X枚举(目录, 匹配文件模式, 是否递归替换...)
	if err != nil {
		return err
	}
	for _, file := range files {
		if err = X子文本替换(欲被替换的子文本, 用作替换的子文本, file); err != nil {
			return err
		}
	}
	return err
}

// ReplaceDirFunc 函数使用回调函数 `f` 替换路径 `path` 下符合模式 `pattern` 的文件内容。
// 参数 `pattern` 指定需要匹配并替换的文件模式。
// 若给定参数 `recursive` 为真，则会递归地进行替换操作。
func X目录子文本替换_函数(回调函数 func(路径, 内容 string) string, 目录, 匹配文件模式 string, 是否递归替换 ...bool) error {
	files, err := X枚举(目录, 匹配文件模式, 是否递归替换...)
	if err != nil {
		return err
	}
	for _, file := range files {
		if err = X子文本替换_函数(回调函数, file); err != nil {
			return err
		}
	}
	return err
}
