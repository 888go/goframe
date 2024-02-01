// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gres 提供了资源管理功能以及文件和字节之间的打包/解包功能。
package gres

const (
	// 目录分隔符。
	Separator = "/"
)

var (
	// 默认资源对象。
	defaultResource = Instance()
)

// Add 函数将解包并把 `content` 添加到默认资源对象中。
// 不必要的参数 `prefix` 表示存储到当前资源对象中每个文件的前缀。
func Add(content string, prefix ...string) error {
	return defaultResource.Add(content, prefix...)
}

// Load 从`path`加载、解压并将数据添加到默认资源对象中。
// 不必要的参数`prefix`表示存储到当前资源对象时每个文件的前缀。
func Load(path string, prefix ...string) error {
	return defaultResource.Load(path, prefix...)
}

// Get 返回指定路径的文件。
func Get(path string) *File {
	return defaultResource.Get(path)
}

// GetWithIndex 搜索指定 `path` 的文件，如果该文件是一个目录，
// 则进一步在该目录下进行索引文件的搜索。
//
// GetWithIndex 通常用于 HTTP 静态文件服务。
func GetWithIndex(path string, indexFiles []string) *File {
	return defaultResource.GetWithIndex(path, indexFiles)
}

// GetContent 直接返回默认资源对象中 `path` 的内容。
func GetContent(path string) []byte {
	return defaultResource.GetContent(path)
}

// Contains 检查默认资源对象中是否存在 `path`。
func Contains(path string) bool {
	return defaultResource.Contains(path)
}

// IsEmpty 检查并返回资源管理器是否为空。
func IsEmpty() bool {
	return defaultResource.tree.IsEmpty()
}

// ScanDir 返回给定路径下的文件，参数 `path` 应为文件夹类型。
//
// 参数 `pattern` 支持多个文件名模式，
// 使用 ',' 符号分隔多个模式。
//
// 若给定的参数 `recursive` 为 true，则会递归扫描目录。
func ScanDir(path string, pattern string, recursive ...bool) []*File {
	return defaultResource.ScanDir(path, pattern, recursive...)
}

// ScanDirFile 返回给定 `path` 下所有子文件的绝对路径，
// 若给定参数 `recursive` 为 true，则会递归扫描目录。
//
// 注意，该函数仅返回文件，不包括目录。
func ScanDirFile(path string, pattern string, recursive ...bool) []*File {
	return defaultResource.ScanDirFile(path, pattern, recursive...)
}

// Export 函数会递归地导出并保存指定路径 `src` 及其所有子文件到指定的系统路径 `dst`。
func Export(src, dst string, option ...ExportOption) error {
	return defaultResource.Export(src, dst, option...)
}

// Dump 打印默认资源对象中的文件。
func Dump() {
	defaultResource.Dump()
}
