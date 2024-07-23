// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gview

import (
	"bytes"
	"context"
	"fmt"
	htmltpl "html/template"
	"strconv"
	texttpl "text/template"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/ghash"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gfsnotify"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gmlock"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/os/gspath"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gutil"
)

const (
	// 内容解析的模板名称。 md5:6afae8a8a8ed33e4
	templateNameForContentParsing = "TemplateContent"
)

// fileCacheItem 是用于模板文件的缓存项。 md5:0b67edc82216beb0
type fileCacheItem struct {
	path    string
	folder  string
	content string
}

var (
	// 模板缓存映射，用于模板文件夹。
	// 注意，这个映射没有过期逻辑。
	// md5:23e4c8f42fd00704
	templates = gmap.NewStrAnyMap(true)

	// 资源模板文件搜索的尝试文件夹。 md5:17efa863e4db400f
	resourceTryFolders = []string{
		"template/", "template", "/template", "/template/",
		"resource/template/", "resource/template", "/resource/template", "/resource/template/",
	}

	// 前缀数组，用于在本地系统中尝试搜索。 md5:51a8f1255f95f3fc
	localSystemTryFolders = []string{"", "template/", "resource/template"}
)

// Parse 使用给定的模板变量`params`解析给定的模板文件`file`，并返回解析后的模板内容。
// md5:4b41bf3f848a2345
func (view *View) Parse(ctx context.Context, file string, params ...Params) (result string, err error) {
	var usedParams Params
	if len(params) > 0 {
		usedParams = params[0]
	}
	return view.ParseOption(ctx, Option{
		File:    file,
		Content: "",
		Orphan:  false,
		Params:  usedParams,
	})
}

// ParseDefault 使用params解析默认模板文件。 md5:32a43fbd413f5a4e
func (view *View) ParseDefault(ctx context.Context, params ...Params) (result string, err error) {
	var usedParams Params
	if len(params) > 0 {
		usedParams = params[0]
	}
	return view.ParseOption(ctx, Option{
		File:    view.config.DefaultFile,
		Content: "",
		Orphan:  false,
		Params:  usedParams,
	})
}

// ParseContent 使用模板变量 `params` 解析给定的模板内容 `content`，并返回解析后的字节切片。
// md5:26fcffe5c26897e5
func (view *View) ParseContent(ctx context.Context, content string, params ...Params) (string, error) {
	var usedParams Params
	if len(params) > 0 {
		usedParams = params[0]
	}
	return view.ParseOption(ctx, Option{
		Content: content,
		Orphan:  false,
		Params:  usedParams,
	})
}

// 用于模板解析的选项。 md5:cdeffab407011a88
type Option struct {
	File    string // 模板文件的路径，可以是绝对路径或相对于搜索路径的相对路径。 md5:6be52fee4d922970
	Content string // 模板内容，如果提供了`Content`，则忽略`File`。 md5:ca0535d67c8790ea
	Orphan  bool   // 如果为真，将`File`视为单个文件解析，不会递归地从其文件夹中解析其他文件。 md5:33ef5ff5d5c82177
	Params  Params // 模板参数映射。 md5:1ffdb0c9f199a7a3
}

// ParseOption 使用 Option 实现模板解析。 md5:ffb69e45da51ff4f
func (view *View) ParseOption(ctx context.Context, option Option) (result string, err error) {
	if option.Content != "" {
		return view.doParseContent(ctx, option.Content, option.Params)
	}
	if option.File == "" {
		return "", gerror.New(`template file cannot be empty`)
	}
	// 它缓存文件、文件夹和内容以提高性能。 md5:18ed1889fbe8ba22
	r := view.fileCacheMap.GetOrSetFuncLock(option.File, func() interface{} {
		var (
			path     string
			folder   string
			content  string
			resource *gres.File
		)
		// 在`file`的绝对路径下进行搜索。 md5:769fae837e95c873
		path, folder, resource, err = view.searchFile(ctx, option.File)
		if err != nil {
			return nil
		}
		if resource != nil {
			content = string(resource.Content())
		} else {
			content = gfile.GetContentsWithCache(path)
		}
		// 异步使用fsnotify监视模板文件的更改。 md5:e8a79bcdc9b5c5a4
		if resource == nil {
			if _, err = gfsnotify.AddOnce("gview.Parse:"+folder, folder, func(event *gfsnotify.Event) {
				// CLEAR THEM ALL.
				view.fileCacheMap.Clear()
				templates.Clear()
				gfsnotify.Exit()
			}); err != nil {
				intlog.Errorf(ctx, `%+v`, err)
			}
		}
		return &fileCacheItem{
			path:    path,
			folder:  folder,
			content: content,
		}
	})
	if r == nil {
		return
	}
	item := r.(*fileCacheItem)
	// 如果模板内容为空，没有必要继续解析。 md5:59270c3283cce903
	if item.content == "" {
		return "", nil
	}
	// 如果它是孤儿选项，它只是通过ParseContent解析单个文件。 md5:bd95b1f5616b7fce
	if option.Orphan {
		return view.doParseContent(ctx, item.content, option.Params)
	}
	// 获取`folder`的模板对象实例。 md5:850769d5264084fa
	var tpl interface{}
	tpl, err = view.getTemplate(item.path, item.folder, fmt.Sprintf(`*%s`, gfile.Ext(item.path)))
	if err != nil {
		return "", err
	}
	// 使用内存锁确保模板解析的并发安全性。 md5:b64152a6d03ebce0
	gmlock.LockFunc("gview.Parse:"+item.path, func() {
		if view.config.AutoEncode {
			tpl, err = tpl.(*htmltpl.Template).Parse(item.content)
		} else {
			tpl, err = tpl.(*texttpl.Template).Parse(item.content)
		}
		if err != nil && item.path != "" {
			err = gerror.Wrap(err, item.path)
		}
	})
	if err != nil {
		return "", err
	}
	// 请注意，模板变量赋值不能改变现有`params`或view.data的值，
	// 因为这两个变量都是指针。它需要将两个映射的值合并到一个新的映射中。
	// md5:07678aa51c871b54
	variables := gutil.MapMergeCopy(option.Params)
	if len(view.data) > 0 {
		gutil.MapMerge(variables, view.data)
	}
	view.setI18nLanguageFromCtx(ctx, variables)

	buffer := bytes.NewBuffer(nil)
	if view.config.AutoEncode {
		newTpl, err := tpl.(*htmltpl.Template).Clone()
		if err != nil {
			return "", err
		}
		if err = newTpl.Execute(buffer, variables); err != nil {
			return "", err
		}
	} else {
		if err = tpl.(*texttpl.Template).Execute(buffer, variables); err != nil {
			return "", err
		}
	}

	// TODO 有没有一种优雅的计划来替换 "<无值>"？. md5:b722bf3a8104fe3b
	result = gstr.Replace(buffer.String(), "<no value>", "")
	result = view.i18nTranslate(ctx, result, variables)
	return result, nil
}

// doParseContent 使用模板变量 `params` 解析给定的模板内容 `content`，并返回解析后的内容作为 []byte 类型。
// md5:9fcc7059fb505864
func (view *View) doParseContent(ctx context.Context, content string, params Params) (string, error) {
	// 如果模板内容为空，没有必要继续解析。 md5:59270c3283cce903
	if content == "" {
		return "", nil
	}
	var (
		err error
		key = fmt.Sprintf("%s_%v_%v", templateNameForContentParsing, view.config.Delimiters, view.config.AutoEncode)
		tpl = templates.GetOrSetFuncLock(key, func() interface{} {
			if view.config.AutoEncode {
				return htmltpl.New(templateNameForContentParsing).Delims(
					view.config.Delimiters[0],
					view.config.Delimiters[1],
				).Funcs(view.funcMap)
			}
			return texttpl.New(templateNameForContentParsing).Delims(
				view.config.Delimiters[0],
				view.config.Delimiters[1],
			).Funcs(view.funcMap)
		})
	)
	// 使用内存锁确保内容解析的并发安全性。 md5:d526d1fe96e88c9d
	hash := strconv.FormatUint(ghash.DJB64([]byte(content)), 10)
	gmlock.LockFunc("gview.ParseContent:"+hash, func() {
		if view.config.AutoEncode {
			tpl, err = tpl.(*htmltpl.Template).Parse(content)
		} else {
			tpl, err = tpl.(*texttpl.Template).Parse(content)
		}
	})
	if err != nil {
		err = gerror.Wrapf(err, `template parsing failed`)
		return "", err
	}
	// 请注意，模板变量赋值不能改变现有`params`或view.data的值，
	// 因为这两个变量都是指针。它需要将两个映射的值合并到一个新的映射中。
	// md5:07678aa51c871b54
	variables := gutil.MapMergeCopy(params)
	if len(view.data) > 0 {
		gutil.MapMerge(variables, view.data)
	}
	view.setI18nLanguageFromCtx(ctx, variables)

	buffer := bytes.NewBuffer(nil)
	if view.config.AutoEncode {
		var newTpl *htmltpl.Template
		newTpl, err = tpl.(*htmltpl.Template).Clone()
		if err != nil {
			err = gerror.Wrapf(err, `template clone failed`)
			return "", err
		}
		if err = newTpl.Execute(buffer, variables); err != nil {
			err = gerror.Wrapf(err, `template parsing failed`)
			return "", err
		}
	} else {
		if err = tpl.(*texttpl.Template).Execute(buffer, variables); err != nil {
			err = gerror.Wrapf(err, `template parsing failed`)
			return "", err
		}
	}
	// TODO 有没有一种优雅的计划来替换 "<无值>"？. md5:b722bf3a8104fe3b
	result := gstr.Replace(buffer.String(), "<no value>", "")
	result = view.i18nTranslate(ctx, result, variables)
	return result, nil
}

// getTemplate 根据给定的模板文件路径 `path` 返回关联的模板对象。
// 它使用模板缓存来提高性能，即对于相同的 `path`，它将返回相同的模板对象。
// 当`path`下的模板文件发生改变（递归检查）时，它会自动刷新模板缓存。
// md5:c5cd3094a5634faa
func (view *View) getTemplate(filePath, folderPath, pattern string) (tpl interface{}, err error) {
	var (
		mapKey  = fmt.Sprintf("%s_%v", filePath, view.config.Delimiters)
		mapFunc = func() interface{} {
			tplName := filePath
			if view.config.AutoEncode {
				tpl = htmltpl.New(tplName).Delims(
					view.config.Delimiters[0],
					view.config.Delimiters[1],
				).Funcs(view.funcMap)
			} else {
				tpl = texttpl.New(tplName).Delims(
					view.config.Delimiters[0],
					view.config.Delimiters[1],
				).Funcs(view.funcMap)
			}
			// 首先检查资源管理器。 md5:da6f8b6e01c9081c
			if !gres.IsEmpty() {
				if files := gres.ScanDirFile(folderPath, pattern, true); len(files) > 0 {
					if view.config.AutoEncode {
						var t = tpl.(*htmltpl.Template)
						for _, v := range files {
							_, err = t.New(v.FileInfo().Name()).Parse(string(v.Content()))
							if err != nil {
								err = view.formatTemplateObjectCreatingError(v.Name(), tplName, err)
								return nil
							}
						}
					} else {
						var t = tpl.(*texttpl.Template)
						for _, v := range files {
							_, err = t.New(v.FileInfo().Name()).Parse(string(v.Content()))
							if err != nil {
								err = view.formatTemplateObjectCreatingError(v.Name(), tplName, err)
								return nil
							}
						}
					}
					return tpl
				}
			}

			// 其次，检查文件系统，
			// 然后递归地自动解析所有子文件。
			// md5:46d132de94281d12
			var files []string
			files, err = gfile.ScanDir(folderPath, pattern, true)
			if err != nil {
				return nil
			}
			if view.config.AutoEncode {
				t := tpl.(*htmltpl.Template)
				for _, file := range files {
					if _, err = t.Parse(gfile.GetContents(file)); err != nil {
						err = view.formatTemplateObjectCreatingError(file, tplName, err)
						return nil
					}
				}
			} else {
				t := tpl.(*texttpl.Template)
				for _, file := range files {
					if _, err = t.Parse(gfile.GetContents(file)); err != nil {
						err = view.formatTemplateObjectCreatingError(file, tplName, err)
						return nil
					}
				}
			}
			return tpl
		}
	)
	result := templates.GetOrSetFuncLock(mapKey, mapFunc)
	if result != nil {
		return result, nil
	}
	return
}

// formatTemplateObjectCreatingError 格式化从创建模板对象中产生的错误。 md5:896510b4d17d39d6
func (view *View) formatTemplateObjectCreatingError(filePath, tplName string, err error) error {
	if err != nil {
		return gerror.NewSkip(1, gstr.Replace(err.Error(), tplName, filePath))
	}
	return nil
}

// searchFile 返回找到的文件`file`的绝对路径以及其模板文件夹路径。
// 请注意，返回的`folder`是模板文件夹路径，而不是返回的模板文件`path`所在的文件夹。
// md5:a3bcfce2f1e0e878
func (view *View) searchFile(ctx context.Context, file string) (path string, folder string, resource *gres.File, err error) {
	var tempPath string
	// 首先检查资源管理器。 md5:da6f8b6e01c9081c
	if !gres.IsEmpty() {
		// Try folders.
		for _, tryFolder := range resourceTryFolders {
			tempPath = tryFolder + file
			if resource = gres.Get(tempPath); resource != nil {
				path = resource.Name()
				folder = tryFolder
				return
			}
		}
		// Search folders.
		view.searchPaths.RLockFunc(func(array []string) {
			for _, searchPath := range array {
				for _, tryFolder := range resourceTryFolders {
					tempPath = searchPath + tryFolder + file
					if resFile := gres.Get(tempPath); resFile != nil {
						path = resFile.Name()
						folder = searchPath + tryFolder
						return
					}
				}
			}
		})
	}

	// 其次，检查文件系统。 md5:1afe55a17dac6b06
	if path == "" {
		// Absolute path.
		path = gfile.RealPath(file)
		if path != "" {
			folder = gfile.Dir(path)
			return
		}
		// In search paths.
		view.searchPaths.RLockFunc(func(array []string) {
			for _, searchPath := range array {
				searchPath = gstr.TrimRight(searchPath, `\/`)
				for _, tryFolder := range localSystemTryFolders {
					relativePath := gstr.TrimRight(
						gfile.Join(tryFolder, file),
						`\/`,
					)
					if path, _ = gspath.Search(searchPath, relativePath); path != "" {
						folder = gfile.Join(searchPath, tryFolder)
						return
					}
				}
			}
		})
	}

	// Error checking.
	if path == "" {
		buffer := bytes.NewBuffer(nil)
		if view.searchPaths.Len() > 0 {
			buffer.WriteString(fmt.Sprintf("cannot find template file \"%s\" in following paths:", file))
			view.searchPaths.RLockFunc(func(array []string) {
				index := 1
				for _, searchPath := range array {
					searchPath = gstr.TrimRight(searchPath, `\/`)
					for _, tryFolder := range localSystemTryFolders {
						buffer.WriteString(fmt.Sprintf(
							"\n%d. %s",
							index, gfile.Join(searchPath, tryFolder),
						))
						index++
					}
				}
			})
		} else {
			buffer.WriteString(fmt.Sprintf("cannot find template file \"%s\" with no path set/add", file))
		}
		if errorPrint() {
			glog.Error(ctx, buffer.String())
		}
		err = gerror.NewCodef(gcode.CodeInvalidParameter, `template file "%s" not found`, file)
	}
	return
}
