// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gview
import (
	"bytes"
	"context"
	"fmt"
	htmltpl "html/template"
	"strconv"
	texttpl "text/template"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/encoding/ghash"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gfsnotify"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/os/gmlock"
	"github.com/888go/goframe/os/gres"
	"github.com/888go/goframe/os/gspath"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gutil"
	)
const (
	// 模板名称，用于内容解析。
	templateNameForContentParsing = "TemplateContent"
)

// fileCacheItem 是用于模板文件的缓存项。
type fileCacheItem struct {
	path    string
	folder  string
	content string
}

var (
// 模板文件夹的模板缓存映射。
// 注意，此映射没有设置过期逻辑。
	templates = gmap.NewStrAnyMap(true)

	// 尝试在以下文件夹中搜索资源模板文件。
	resourceTryFolders = []string{
		"template/", "template", "/template", "/template/",
		"resource/template/", "resource/template", "/resource/template", "/resource/template/",
	}

	// 前缀数组，用于尝试在本地系统中进行搜索。
	localSystemTryFolders = []string{"", "template/", "resource/template"}
)

// Parse函数用于解析给定的模板文件`file`，并使用给定的模板变量`params`进行解析，
// 然后返回解析后的模板内容。
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

// ParseDefault 通过给定的参数解析默认模板文件。
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

// ParseContent 函数用于解析给定的模板内容 `content`，同时使用模板变量 `params` 进行替换，
// 并将解析后的内容以 []byte 类型返回。
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

// 模板解析的选项。
type Option struct {
	File    string // 模板文件路径，可以是绝对路径，也可以相对于搜索路径。
	Content string // 模板内容，如果提供了`Content`，则会忽略`File`。
	Orphan  bool   // 如果为true，那么`File`被视为单个文件解析，不递归地从其所在文件夹中解析其他文件。
	Params  Params // 模板参数映射。
}

// ParseOption 实现了通过 Option 进行模板解析的功能。
func (view *View) ParseOption(ctx context.Context, option Option) (result string, err error) {
	if option.Content != "" {
		return view.doParseContent(ctx, option.Content, option.Params)
	}
	if option.File == "" {
		return "", gerror.New(`template file cannot be empty`)
	}
	// 它缓存文件、文件夹及其内容以提高性能。
	r := view.fileCacheMap.GetOrSetFuncLock(option.File, func() interface{} {
		var (
			path     string
			folder   string
			content  string
			resource *gres.File
		)
		// 搜索`file`的绝对文件路径。
		path, folder, resource, err = view.searchFile(ctx, option.File)
		if err != nil {
			return nil
		}
		if resource != nil {
			content = string(resource.Content())
		} else {
			content = gfile.GetContentsWithCache(path)
		}
		// 使用fsnotify异步监控模板文件的变更。
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
	// 如果模板内容为空，则没有必要继续解析。
	if item.content == "" {
		return "", nil
	}
	// 如果是Orphan选项，它仅通过ParseContent解析单个文件。
	if option.Orphan {
		return view.doParseContent(ctx, item.content, option.Params)
	}
	// 获取`folder`对应的模板对象实例。
	var tpl interface{}
	tpl, err = view.getTemplate(item.path, item.folder, fmt.Sprintf(`*%s`, gfile.Ext(item.path)))
	if err != nil {
		return "", err
	}
	// 使用内存锁以确保模板解析过程中的并发安全性。
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
// 注意，模板变量赋值无法改变已存在的`params`或view.data的值，
// 因为两者都是指针变量。它需要将两个映射的值合并到一个新的映射中。
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

	// TODO 是否有优雅的方案来替换 "<无值>"？
	result = gstr.Replace(buffer.String(), "<no value>", "")
	result = view.i18nTranslate(ctx, result, variables)
	return result, nil
}

// doParseContent 函数用于解析给定的模板内容 `content`，并使用模板变量 `params` 进行替换，
// 然后返回已解析内容的 []byte 类型数据。
func (view *View) doParseContent(ctx context.Context, content string, params Params) (string, error) {
	// 如果模板内容为空，则没有必要继续解析。
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
	// 使用内存锁以确保内容解析的并发安全。
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
// 注意，模板变量赋值无法改变已存在的`params`或view.data的值，
// 因为两者都是指针变量。它需要将两个映射的值合并到一个新的映射中。
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
	// TODO 是否有优雅的方案来替换 "<无值>"？
	result := gstr.Replace(buffer.String(), "<no value>", "")
	result = view.i18nTranslate(ctx, result, variables)
	return result, nil
}

// getTemplate 返回与给定模板文件`path`关联的模板对象。
// 它利用模板缓存来提升性能，即对于相同的给定`path`，它将返回相同的模板对象。
// 同时，如果`path`路径下的模板文件发生变化（递归检测），它会自动刷新模板缓存。
// 这段代码注释翻译成中文为：
// ```go
// getTemplate 函数用于获取与指定模板文件 `path` 相关联的模板对象。
// 为了提高性能，它使用了模板缓存技术。这意味着当传入相同的 `path` 时，它会返回同一模板对象。
// 另外，该函数会自动监测并更新缓存：一旦发现 `path` 路径下（包括子目录）的模板文件发生变动，就会自动刷新模板缓存。
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
			// 首先检查资源管理器。
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

// 其次检查文件系统，
// 然后递归地自动解析其所有子文件。
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

// formatTemplateObjectCreatingError 格式化创建模板对象时产生的错误信息。
func (view *View) formatTemplateObjectCreatingError(filePath, tplName string, err error) error {
	if err != nil {
		return gerror.NewSkip(1, gstr.Replace(err.Error(), tplName, filePath))
	}
	return nil
}

// searchFile 函数返回文件 `file` 找到的绝对路径及其对应的模板文件夹路径。
// 注意，返回的 `folder` 是模板文件夹路径，并不是返回的模板文件 `path` 的所在文件夹路径。
func (view *View) searchFile(ctx context.Context, file string) (path string, folder string, resource *gres.File, err error) {
	var tempPath string
	// 首先检查资源管理器。
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

	// 第二步检查文件系统。
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
