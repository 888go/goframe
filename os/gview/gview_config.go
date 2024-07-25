// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gview

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/os/gspath"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

// Config是模板引擎的配置对象。 md5:0c7a20a5c1f534d4
type Config struct {
	Paths       []string               `json:"paths"`       // 在数组中搜索路径，为了性能原因，非并发安全。 md5:536357ec68a07213
	Data        map[string]interface{} `json:"data"`        // 全局模板变量，包括配置信息。 md5:5f96c7a35c11b4b2
	DefaultFile string                 `json:"defaultFile"` // 默认的模板文件用于解析。 md5:41607c84f42fcf9d
	Delimiters  []string               `json:"delimiters"`  // 自定义模板分隔符。 md5:0a97ca0eda8842d4
	AutoEncode  bool                   `json:"autoEncode"`  // 自动进行编码并提供安全的HTML输出，这对于防止XSS攻击很有帮助。 md5:ec33e2ef01aaf3d3
	I18nManager *gi18n.Manager         `json:"-"`           // 视图的国际化管理器。 md5:7c90b657f5c4c28b
}

const (
	// 默认的模板文件用于解析。 md5:41607c84f42fcf9d
	defaultParsingFile = "index.html"
)

// DefaultConfig 创建并返回一个使用默认配置的配置对象。 md5:27f0cf63ebd5dd9e
func DefaultConfig() Config {
	return Config{
		DefaultFile: defaultParsingFile,
		I18nManager: gi18n.Instance(),
		Delimiters:  make([]string, 2),
	}
}

// SetConfig 设置视图的配置。 md5:44d304b99e74e865
func (view *View) SetConfig(config Config) error {
	var err error
	if len(config.Paths) > 0 {
		for _, v := range config.Paths {
			if err = view.AddPath(v); err != nil {
				return err
			}
		}
	}
	if len(config.Data) > 0 {
		view.Assigns(config.Data)
	}
	if config.DefaultFile != "" {
		view.SetDefaultFile(config.DefaultFile)
	}
	if len(config.Delimiters) > 1 {
		view.SetDelimiters(config.Delimiters[0], config.Delimiters[1])
	}
	view.config = config
	// 清除全局模板对象缓存。
	// 这只是一个缓存，不要犹豫清空它。 md5:51c51fe68d143dd8
	templates.Clear()

	intlog.Printf(context.TODO(), "SetConfig: %+v", view.config)
	return nil
}

// SetConfigWithMap 使用映射为视图设置配置。 md5:1e1d667c3b2ace2b
func (view *View) SetConfigWithMap(m map[string]interface{}) error {
	if len(m) == 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "configuration cannot be empty")
	}
	// m 现在是 m 的浅拷贝。
	// 对 m 的任何修改都不会影响原始对象。
	// 这有点巧妙，不是吗？ md5:4d1dd38c4db57a79
	m = gutil.MapCopy(m)
	// 最常用的单视图路径配置支持。 md5:4ebc24cd15a30d35
	_, v1 := gutil.MapPossibleItemByKey(m, "paths")
	_, v2 := gutil.MapPossibleItemByKey(m, "path")
	if v1 == nil && v2 != nil {
		switch v2.(type) {
		case string:
			m["paths"] = []string{v2.(string)}
		case []string:
			m["paths"] = v2
		}
	}
	err := gconv.Struct(m, &view.config)
	if err != nil {
		return err
	}
	return view.SetConfig(view.config)
}

// SetPath 设置模板文件搜索的目录路径。参数 `path` 可以是绝对路径或相对路径，但建议使用绝对路径。 md5:abd751ab819d28b6
func (view *View) SetPath(path string) error {
	var (
		ctx      = context.TODO()
		isDir    = false
		realPath = ""
	)
	if file := gres.Get(path); file != nil {
		realPath = path
		isDir = file.FileInfo().IsDir()
	} else {
		// Absolute path.
		realPath = gfile.RealPath(path)
		if realPath == "" {
			// Relative path.
			view.searchPaths.RLockFunc(func(array []string) {
				for _, v := range array {
					if path, _ := gspath.Search(v, path); path != "" {
						realPath = path
						break
					}
				}
			})
		}
		if realPath != "" {
			isDir = gfile.IsDir(realPath)
		}
	}
	// Path not exist.
	if realPath == "" {
		err := gerror.NewCodef(gcode.CodeInvalidParameter, `View.SetPath failed: path "%s" does not exist`, path)
		if errorPrint() {
			glog.Error(ctx, err)
		}
		return err
	}
	// Should be a directory.
	if !isDir {
		err := gerror.NewCodef(gcode.CodeInvalidParameter, `View.SetPath failed: path "%s" should be directory type`, path)
		if errorPrint() {
			glog.Error(ctx, err)
		}
		return err
	}
	// 重复路径添加检查。 md5:e210e91d65ec4857
	if view.searchPaths.Search(realPath) != -1 {
		return nil
	}
	view.searchPaths.Clear()
	view.searchPaths.Append(realPath)
	view.fileCacheMap.Clear()
	return nil
}

// AddPath 向搜索路径中添加一个绝对或相对路径。 md5:d279479528c86f4e
func (view *View) AddPath(path string) error {
	var (
		ctx      = context.TODO()
		isDir    = false
		realPath = ""
	)
	if file := gres.Get(path); file != nil {
		realPath = path
		isDir = file.FileInfo().IsDir()
	} else {
		// Absolute path.
		if realPath = gfile.RealPath(path); realPath == "" {
			// Relative path.
			view.searchPaths.RLockFunc(func(array []string) {
				for _, v := range array {
					if searchedPath, _ := gspath.Search(v, path); searchedPath != "" {
						realPath = searchedPath
						break
					}
				}
			})
		}
		if realPath != "" {
			isDir = gfile.IsDir(realPath)
		}
	}
	// Path not exist.
	if realPath == "" {
		err := gerror.NewCodef(gcode.CodeInvalidParameter, `View.AddPath failed: path "%s" does not exist`, path)
		if errorPrint() {
			glog.Error(ctx, err)
		}
		return err
	}
	// realPath 应该是文件夹类型的路径。 md5:8b57fae1c1158ae9
	if !isDir {
		err := gerror.NewCodef(gcode.CodeInvalidParameter, `View.AddPath failed: path "%s" should be directory type`, path)
		if errorPrint() {
			glog.Error(ctx, err)
		}
		return err
	}
	// 重复路径添加检查。 md5:e210e91d65ec4857
	if view.searchPaths.Search(realPath) != -1 {
		return nil
	}
	view.searchPaths.Append(realPath)
	view.fileCacheMap.Clear()
	return nil
}

// 将多个全局模板变量绑定到当前视图对象。需要注意的是，它不是并发安全的，这意味着如果在运行时从多个goroutine中调用它，会导致panic。 md5:b31929b349e74390
func (view *View) Assigns(data Params) {
	for k, v := range data {
		view.data[k] = v
	}
}

// Assign 将全局模板变量绑定到当前视图对象。需要注意的是，它不是线程安全的，这意味着如果在运行时从多个goroutine中调用它，会导致panic。 md5:7043c41fc2b3a0c3
func (view *View) Assign(key string, value interface{}) {
	view.data[key] = value
}

// SetDefaultFile 为解析设置默认的模板文件。 md5:17f210ece0d189f6
func (view *View) SetDefaultFile(file string) {
	view.config.DefaultFile = file
}

// GetDefaultFile 返回默认的模板文件，用于解析。 md5:f72bb2dc04f3d4a4
func (view *View) GetDefaultFile() string {
	return view.config.DefaultFile
}

// SetDelimiters 设置模板解析的自定义分隔符。 md5:a09465c3518f1023
func (view *View) SetDelimiters(left, right string) {
	view.config.Delimiters = []string{left, right}
}

// SetAutoEncode 启用/禁用自动 HTML 编码功能。
// 当 AutoEncode 功能启用时，视图引擎会自动编码并提供安全的 HTML 输出，这对于防止 XSS 攻击很有好处。 md5:cd0107f5d2170f4f
func (view *View) SetAutoEncode(enable bool) {
	view.config.AutoEncode = enable
}

// BindFunc 向当前视图对象注册一个名为 `name` 的自定义全局模板函数，
// 使用提供的 `function` 函数。其中，`name` 是在模板内容中可被调用的函数名。 md5:20f79a4c8d0ba97a
func (view *View) BindFunc(name string, function interface{}) {
	view.funcMap[name] = function
	// Clear global template object cache.
	templates.Clear()
}

// BindFuncMap 将自定义的全局模板函数通过映射注册到当前视图对象中。
// 映射的键是模板函数名称，
// 映射的值是自定义函数的地址。 md5:2fe9bab0463cef27
func (view *View) BindFuncMap(funcMap FuncMap) {
	for k, v := range funcMap {
		view.funcMap[k] = v
	}
	// Clear global template object cache.
	templates.Clear()
}

// SetI18n 将i18n管理器绑定到当前视图引擎。 md5:8d1b88bd87c041ba
func (view *View) SetI18n(manager *gi18n.Manager) {
	view.config.I18nManager = manager
}
