
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Template name for content parsing.
<原文结束>

# <翻译开始>
// 模板名称，用于内容解析。
# <翻译结束>


<原文开始>
// fileCacheItem is the cache item for template file.
<原文结束>

# <翻译开始>
// fileCacheItem 是用于模板文件的缓存项。
# <翻译结束>


<原文开始>
	// Templates cache map for template folder.
	// Note that there's no expiring logic for this map.
<原文结束>

# <翻译开始>
// 模板文件夹的模板缓存映射。
// 注意，此映射没有设置过期逻辑。
# <翻译结束>


<原文开始>
// Try-folders for resource template file searching.
<原文结束>

# <翻译开始>
// 尝试在以下文件夹中搜索资源模板文件。
# <翻译结束>


<原文开始>
// Prefix array for trying searching in local system.
<原文结束>

# <翻译开始>
// 前缀切片，用于尝试在本地系统中进行搜索。
# <翻译结束>


<原文开始>
// Parse parses given template file `file` with given template variables `params`
// and returns the parsed template content.
<原文结束>

# <翻译开始>
// Parse函数用于解析给定的模板文件`file`，并使用给定的模板变量`params`进行解析，
// 然后返回解析后的模板内容。
# <翻译结束>


<原文开始>
// ParseDefault parses the default template file with params.
<原文结束>

# <翻译开始>
// ParseDefault 通过给定的参数解析默认模板文件。
# <翻译结束>


<原文开始>
// ParseContent parses given template content `content`  with template variables `params`
// and returns the parsed content in []byte.
<原文结束>

# <翻译开始>
// ParseContent 函数用于解析给定的模板内容 `content`，同时使用模板变量 `params` 进行替换，
// 并将解析后的内容以 []byte 类型返回。
# <翻译结束>


<原文开始>
// Option for template parsing.
<原文结束>

# <翻译开始>
// 模板解析的选项。
# <翻译结束>


<原文开始>
// Template file path in absolute or relative to searching paths.
<原文结束>

# <翻译开始>
// 模板文件路径，可以是绝对路径，也可以相对于搜索路径。
# <翻译结束>


<原文开始>
// Template content, it ignores `File` if `Content` is given.
<原文结束>

# <翻译开始>
// 模板内容，如果提供了`Content`，则会忽略`File`。
# <翻译结束>


<原文开始>
// If true, the `File` is considered as a single file parsing without files recursively parsing from its folder.
<原文结束>

# <翻译开始>
// 如果为true，那么`File`被视为单个文件解析，不递归地从其所在文件夹中解析其他文件。
# <翻译结束>


<原文开始>
// ParseOption implements template parsing using Option.
<原文结束>

# <翻译开始>
// ParseOption 实现了通过 Option 进行模板解析的功能。
# <翻译结束>


<原文开始>
// It caches the file, folder and content to enhance performance.
<原文结束>

# <翻译开始>
// 它缓存文件、文件夹及其内容以提高性能。
# <翻译结束>


<原文开始>
// Searching the absolute file path for `file`.
<原文结束>

# <翻译开始>
// 搜索`file`的绝对文件路径。
# <翻译结束>


<原文开始>
// Monitor template files changes using fsnotify asynchronously.
<原文结束>

# <翻译开始>
// 使用fsnotify异步监控模板文件的变更。
# <翻译结束>


<原文开始>
// It's not necessary continuing parsing if template content is empty.
<原文结束>

# <翻译开始>
// 如果模板内容为空，则没有必要继续解析。
# <翻译结束>


<原文开始>
// If it's Orphan option, it just parses the single file by ParseContent.
<原文结束>

# <翻译开始>
// 如果是Orphan选项，它仅通过ParseContent解析单个文件。
# <翻译结束>


<原文开始>
// Get the template object instance for `folder`.
<原文结束>

# <翻译开始>
// 获取`folder`对应的模板对象实例。
# <翻译结束>


<原文开始>
// Using memory lock to ensure concurrent safety for template parsing.
<原文结束>

# <翻译开始>
// 使用内存锁以确保模板解析过程中的并发安全性。
# <翻译结束>


<原文开始>
	// Note that the template variable assignment cannot change the value
	// of the existing `params` or view.data because both variables are pointers.
	// It needs to merge the values of the two maps into a new map.
<原文结束>

# <翻译开始>
// 注意，模板变量赋值无法改变已存在的`params`或view.data的值，
// 因为两者都是指针变量。它需要将两个映射的值合并到一个新的映射中。
# <翻译结束>


<原文开始>
// TODO any graceful plan to replace "<no value>"?
<原文结束>

# <翻译开始>
// TODO 是否有优雅的方案来替换 "<无值>"？
# <翻译结束>


<原文开始>
// doParseContent parses given template content `content`  with template variables `params`
// and returns the parsed content in []byte.
<原文结束>

# <翻译开始>
// doParseContent 函数用于解析给定的模板内容 `content`，并使用模板变量 `params` 进行替换，
// 然后返回已解析内容的 []byte 类型数据。
# <翻译结束>


<原文开始>
// Using memory lock to ensure concurrent safety for content parsing.
<原文结束>

# <翻译开始>
// 使用内存锁以确保内容解析的并发安全。
# <翻译结束>


<原文开始>
// getTemplate returns the template object associated with given template file `path`.
// It uses template cache to enhance performance, that is, it will return the same template object
// with the same given `path`. It will also automatically refresh the template cache
// if the template files under `path` changes (recursively).
<原文结束>

# <翻译开始>
// getTemplate 返回与给定模板文件`path`关联的模板对象。
// 它利用模板缓存来提升性能，即对于相同的给定`path`，它将返回相同的模板对象。
// 同时，如果`path`路径下的模板文件发生变化（递归检测），它会自动刷新模板缓存。
// 这段代码注释翻译成中文为：
// ```go
// getTemplate 函数用于获取与指定模板文件 `path` 相关联的模板对象。
// 为了提高性能，它使用了模板缓存技术。这意味着当传入相同的 `path` 时，它会返回同一模板对象。
// 另外，该函数会自动监测并更新缓存：一旦发现 `path` 路径下（包括子目录）的模板文件发生变动，就会自动刷新模板缓存。
# <翻译结束>


<原文开始>
// Firstly checking the resource manager.
<原文结束>

# <翻译开始>
// 首先检查资源管理器。
# <翻译结束>


<原文开始>
			// Secondly checking the file system,
			// and then automatically parsing all its sub-files recursively.
<原文结束>

# <翻译开始>
// 其次检查文件系统，
// 然后递归地自动解析其所有子文件。
# <翻译结束>


<原文开始>
// formatTemplateObjectCreatingError formats the error that created from creating template object.
<原文结束>

# <翻译开始>
// formatTemplateObjectCreatingError 格式化创建模板对象时产生的错误信息。
# <翻译结束>


<原文开始>
// searchFile returns the found absolute path for `file` and its template folder path.
// Note that, the returned `folder` is the template folder path, but not the folder of
// the returned template file `path`.
<原文结束>

# <翻译开始>
// searchFile 函数返回文件 `file` 找到的绝对路径及其对应的模板文件夹路径。
// 注意，返回的 `folder` 是模板文件夹路径，并不是返回的模板文件 `path` 的所在文件夹路径。
# <翻译结束>


<原文开始>
// Secondly checking the file system.
<原文结束>

# <翻译开始>
// 第二步检查文件系统。
# <翻译结束>


<原文开始>
// Template parameters map.
<原文结束>

# <翻译开始>
// 模板参数映射。
# <翻译结束>

