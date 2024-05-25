
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Template name for content parsing.
<原文结束>

# <翻译开始>
// 内容解析的模板名称。 md5:6afae8a8a8ed33e4
# <翻译结束>


<原文开始>
// fileCacheItem is the cache item for template file.
<原文结束>

# <翻译开始>
// fileCacheItem 是用于模板文件的缓存项。 md5:0b67edc82216beb0
# <翻译结束>


<原文开始>
	// Templates cache map for template folder.
	// Note that there's no expiring logic for this map.
<原文结束>

# <翻译开始>
// 模板缓存映射，用于模板文件夹。
// 注意，这个映射没有过期逻辑。
// md5:23e4c8f42fd00704
# <翻译结束>


<原文开始>
// Try-folders for resource template file searching.
<原文结束>

# <翻译开始>
// 资源模板文件搜索的尝试文件夹。 md5:17efa863e4db400f
# <翻译结束>


<原文开始>
// Prefix array for trying searching in local system.
<原文结束>

# <翻译开始>
// 前缀数组，用于在本地系统中尝试搜索。 md5:51a8f1255f95f3fc
# <翻译结束>


<原文开始>
// Parse parses given template file `file` with given template variables `params`
// and returns the parsed template content.
<原文结束>

# <翻译开始>
// Parse 使用给定的模板变量`params`解析给定的模板文件`file`，并返回解析后的模板内容。
// md5:4b41bf3f848a2345
# <翻译结束>


<原文开始>
// ParseDefault parses the default template file with params.
<原文结束>

# <翻译开始>
// ParseDefault 使用params解析默认模板文件。 md5:32a43fbd413f5a4e
# <翻译结束>


<原文开始>
// ParseContent parses given template content `content`  with template variables `params`
// and returns the parsed content in []byte.
<原文结束>

# <翻译开始>
// ParseContent 使用模板变量 `params` 解析给定的模板内容 `content`，并返回解析后的字节切片。
// md5:26fcffe5c26897e5
# <翻译结束>


<原文开始>
// Option for template parsing.
<原文结束>

# <翻译开始>
// 用于模板解析的选项。 md5:cdeffab407011a88
# <翻译结束>


<原文开始>
// Template file path in absolute or relative to searching paths.
<原文结束>

# <翻译开始>
// 模板文件的路径，可以是绝对路径或相对于搜索路径的相对路径。 md5:6be52fee4d922970
# <翻译结束>


<原文开始>
// Template content, it ignores `File` if `Content` is given.
<原文结束>

# <翻译开始>
// 模板内容，如果提供了`Content`，则忽略`File`。 md5:ca0535d67c8790ea
# <翻译结束>


<原文开始>
// If true, the `File` is considered as a single file parsing without files recursively parsing from its folder.
<原文结束>

# <翻译开始>
// 如果为真，将`File`视为单个文件解析，不会递归地从其文件夹中解析其他文件。 md5:33ef5ff5d5c82177
# <翻译结束>


<原文开始>
// Template parameters map.
<原文结束>

# <翻译开始>
// 模板参数映射。 md5:1ffdb0c9f199a7a3
# <翻译结束>


<原文开始>
// ParseOption implements template parsing using Option.
<原文结束>

# <翻译开始>
// ParseOption 使用 Option 实现模板解析。 md5:ffb69e45da51ff4f
# <翻译结束>


<原文开始>
// It caches the file, folder and content to enhance performance.
<原文结束>

# <翻译开始>
// 它缓存文件、文件夹和内容以提高性能。 md5:18ed1889fbe8ba22
# <翻译结束>


<原文开始>
// Searching the absolute file path for `file`.
<原文结束>

# <翻译开始>
// 在`file`的绝对路径下进行搜索。 md5:769fae837e95c873
# <翻译结束>


<原文开始>
// Monitor template files changes using fsnotify asynchronously.
<原文结束>

# <翻译开始>
// 异步使用fsnotify监视模板文件的更改。 md5:e8a79bcdc9b5c5a4
# <翻译结束>


<原文开始>
// It's not necessary continuing parsing if template content is empty.
<原文结束>

# <翻译开始>
// 如果模板内容为空，没有必要继续解析。 md5:59270c3283cce903
# <翻译结束>


<原文开始>
// If it's Orphan option, it just parses the single file by ParseContent.
<原文结束>

# <翻译开始>
// 如果它是孤儿选项，它只是通过ParseContent解析单个文件。 md5:bd95b1f5616b7fce
# <翻译结束>


<原文开始>
// Get the template object instance for `folder`.
<原文结束>

# <翻译开始>
// 获取`folder`的模板对象实例。 md5:850769d5264084fa
# <翻译结束>


<原文开始>
// Using memory lock to ensure concurrent safety for template parsing.
<原文结束>

# <翻译开始>
// 使用内存锁确保模板解析的并发安全性。 md5:b64152a6d03ebce0
# <翻译结束>


<原文开始>
	// Note that the template variable assignment cannot change the value
	// of the existing `params` or view.data because both variables are pointers.
	// It needs to merge the values of the two maps into a new map.
<原文结束>

# <翻译开始>
// 请注意，模板变量赋值不能改变现有`params`或view.data的值，
// 因为这两个变量都是指针。它需要将两个映射的值合并到一个新的映射中。
// md5:07678aa51c871b54
# <翻译结束>


<原文开始>
// TODO any graceful plan to replace "<no value>"?
<原文结束>

# <翻译开始>
// TODO 有没有一种优雅的计划来替换 "<无值>"？. md5:b722bf3a8104fe3b
# <翻译结束>


<原文开始>
// doParseContent parses given template content `content`  with template variables `params`
// and returns the parsed content in []byte.
<原文结束>

# <翻译开始>
// doParseContent 使用模板变量 `params` 解析给定的模板内容 `content`，并返回解析后的内容作为 []byte 类型。
// md5:9fcc7059fb505864
# <翻译结束>


<原文开始>
// Using memory lock to ensure concurrent safety for content parsing.
<原文结束>

# <翻译开始>
// 使用内存锁确保内容解析的并发安全性。 md5:d526d1fe96e88c9d
# <翻译结束>


<原文开始>
// getTemplate returns the template object associated with given template file `path`.
// It uses template cache to enhance performance, that is, it will return the same template object
// with the same given `path`. It will also automatically refresh the template cache
// if the template files under `path` changes (recursively).
<原文结束>

# <翻译开始>
// getTemplate 根据给定的模板文件路径 `path` 返回关联的模板对象。
// 它使用模板缓存来提高性能，即对于相同的 `path`，它将返回相同的模板对象。
// 当`path`下的模板文件发生改变（递归检查）时，它会自动刷新模板缓存。
// md5:c5cd3094a5634faa
# <翻译结束>


<原文开始>
// Firstly checking the resource manager.
<原文结束>

# <翻译开始>
// 首先检查资源管理器。 md5:da6f8b6e01c9081c
# <翻译结束>


<原文开始>
			// Secondly checking the file system,
			// and then automatically parsing all its sub-files recursively.
<原文结束>

# <翻译开始>
// 其次，检查文件系统，
// 然后递归地自动解析所有子文件。
// md5:46d132de94281d12
# <翻译结束>


<原文开始>
// formatTemplateObjectCreatingError formats the error that created from creating template object.
<原文结束>

# <翻译开始>
// formatTemplateObjectCreatingError 格式化从创建模板对象中产生的错误。 md5:896510b4d17d39d6
# <翻译结束>


<原文开始>
// searchFile returns the found absolute path for `file` and its template folder path.
// Note that, the returned `folder` is the template folder path, but not the folder of
// the returned template file `path`.
<原文结束>

# <翻译开始>
// searchFile 返回找到的文件`file`的绝对路径以及其模板文件夹路径。
// 请注意，返回的`folder`是模板文件夹路径，而不是返回的模板文件`path`所在的文件夹。
// md5:a3bcfce2f1e0e878
# <翻译结束>


<原文开始>
// Secondly checking the file system.
<原文结束>

# <翻译开始>
// 其次，检查文件系统。 md5:1afe55a17dac6b06
# <翻译结束>

