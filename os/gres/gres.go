// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gres提供了文件和字节之间资源管理及打包/解包的功能。 md5:29e79f40a11fe941
package gres//bm:资源类

const (
	// 目录分隔符。 md5:a4832545f002edfe
	Separator = "/"
)

var (
	// 默认资源对象。 md5:f02aee71ab2f8fc2
	defaultResource = Instance()
)

// Add 方法将 'content' 解包并添加到默认资源对象中。
// 不需要的参数 'prefix' 表示存储在当前资源对象中的每个文件的前缀。
// md5:3b5da05501708d4d
// ff:
// content:
// prefix:
func Add(content string, prefix ...string) error {
	return defaultResource.Add(content, prefix...)
}

// Load 从 `path` 加载、解包并将数据添加到默认资源对象中。
// 不必要的参数 `prefix` 表示存储到当前资源对象中的每个文件的前缀。
// md5:901d4f7e4d8bf0cf
// ff:
// path:
// prefix:
func Load(path string, prefix ...string) error {
	return defaultResource.Load(path, prefix...)
}

// Get返回给定路径的文件。 md5:f4989a4832cde2d2
// ff:
// path:
func Get(path string) *File {
	return defaultResource.Get(path)
}

// GetWithIndex 在给定路径`path`下搜索文件。如果找到的是一个目录，它会在这个目录下索引文件进行搜索。
//
// GetWithIndex 通常用于HTTP静态文件服务中。
// md5:bfb61cc8920b4633
// ff:
// path:
// indexFiles:
func GetWithIndex(path string, indexFiles []string) *File {
	return defaultResource.GetWithIndex(path, indexFiles)
}

// GetContent 直接返回默认资源对象中 `path` 的内容。 md5:6446043ef668020c
// ff:
// path:
func GetContent(path string) []byte {
	return defaultResource.GetContent(path)
}

// Contains 检查默认资源对象中是否存在 `path`。 md5:f69f9f792a33a089
// ff:
// path:
func Contains(path string) bool {
	return defaultResource.Contains(path)
}

// IsEmpty 检查资源管理器是否为空，并返回结果。 md5:3aaae27781ad4e8c
// ff:
func IsEmpty() bool {
	return defaultResource.tree.IsEmpty()
}

// ScanDir 在给定路径下返回文件，参数 `path` 应该是一个文件夹类型。
//
// `pattern` 参数支持多个文件名模式，使用逗号 `,` 来分隔多个模式。
//
// 如果 `recursive` 参数为真，它会递归扫描目录。
// md5:4726ded4e00ca75f
// ff:
// path:
// pattern:
// recursive:
func ScanDir(path string, pattern string, recursive ...bool) []*File {
	return defaultResource.ScanDir(path, pattern, recursive...)
}

// ScanDirFile 返回给定`path`下的所有子文件的绝对路径，
// 如果给定的参数`recursive`为true，它会递归地扫描目录。
//
// 注意，它只返回文件，不包括目录。
// md5:0f3154c32271652b
// ff:
// path:
// pattern:
// recursive:
func ScanDirFile(path string, pattern string, recursive ...bool) []*File {
	return defaultResource.ScanDirFile(path, pattern, recursive...)
}

// Export 将指定的路径 `src` 及其所有子文件递归导出并保存到指定的系统路径 `dst`。 md5:944ad6e86342817b
// ff:
// src:
// dst:
// option:
func Export(src, dst string, option ...ExportOption) error {
	return defaultResource.Export(src, dst, option...)
}

// Dump 打印默认资源对象的文件。 md5:fc090361befff87e
// ff:
func Dump() {
	defaultResource.Dump()
}
