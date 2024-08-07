// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文件类

import (
	gstr "github.com/888go/goframe/text/gstr"
)

// X子文本替换 替换文件 `path` 的内容。 md5:70e99046a619416c
func X子文本替换(欲被替换的子文本, 用作替换的子文本, 文件路径 string) error {
	return X写入文本(文件路径, gstr.X替换(X读文本(文件路径), 欲被替换的子文本, 用作替换的子文本))
}

// X子文本替换_函数 使用回调函数 `f` 替换文件 `path` 中的内容。 md5:033d4157195d29fc
func X子文本替换_函数(回调函数 func(文件路径, content string) string, 文件路径 string) error {
	data := X读文本(文件路径)
	result := 回调函数(文件路径, data)
	if len(data) != len(result) || data != result {
		return X写入文本(文件路径, result)
	}
	return nil
}

// X目录子文本替换 替换给定路径下文件的内容。
// 参数 `pattern` 指定了需要被替换的文件模式。
// 如果参数 `recursive` 为真，它将递归地进行替换操作。
// md5:d9b10978d6db3bce
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

// X目录子文本替换_函数 使用回调函数 `f` 替换路径 `path` 下的文件内容。
// 参数 `pattern` 指定了将被替换的文件模式。
// 如果 `recursive` 为 true，它会递归地进行替换。
// md5:9bff4662f6d662c1
func X目录子文本替换_函数(回调函数 func(目录, content string) string, 目录, 匹配文件模式 string, 是否递归替换 ...bool) error {
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
