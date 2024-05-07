// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 模板类

import (
	"context"

	"github.com/888go/goframe/gview/internal/intlog"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/os/gspath"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

// Config 是用于模板引擎的配置对象。
type Config struct {
	Paths       []string               `json:"paths"`       // 为了性能考虑，以下代码在切片中搜索路径，但并不保证并发安全。
	Data        map[string]interface{} `json:"data"`        // 全局模板变量，包括配置信息。
	DefaultFile string                 `json:"defaultFile"` // 默认用于解析的模板文件。
	Delimiters  []string               `json:"delimiters"`  // 自定义模板分隔符。
	AutoEncode  bool                   `json:"autoEncode"`  // 自动进行编码并提供安全的HTML输出，有助于避免XSS攻击。
	I18nManager *gi18n.Manager         `json:"-"`           // 视图的国际化管理器。
}

const (
	// 默认用于解析的模板文件。
	defaultParsingFile = "index.html"
)

// DefaultConfig 创建并返回一个包含默认配置的配置对象。
func DefaultConfig() Config {
	return Config{
		DefaultFile: defaultParsingFile,
		I18nManager: gi18n.Instance(),
		Delimiters:  make([]string, 2),
	}
}

// SetConfig 设置视图的配置。
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
	// 这只是缓存，不必犹豫去清除它。
	templates.Clear()

	intlog.Printf(context.TODO(), "SetConfig: %+v", view.config)
	return nil
}

// SetConfigWithMap 使用map设置视图的相关配置。
func (view *View) SetConfigWithMap(m map[string]interface{}) error {
	if len(m) == 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "configuration cannot be empty")
	}
	// 现在的m是m的一个浅拷贝。
	// 对m的任何改动都不会影响原始的那个m。
	// 有点小巧妙，不是吗？
	m = gutil.MapCopy(m)
	// 最常用的单视图路径配置支持。
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

// SetPath 设置模板文件搜索的目录路径。
// 参数 `path` 可以是绝对路径或相对路径，但建议使用绝对路径。
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
	// 应该是一个目录。
	if !isDir {
		err := gerror.NewCodef(gcode.CodeInvalidParameter, `View.SetPath failed: path "%s" should be directory type`, path)
		if errorPrint() {
			glog.Error(ctx, err)
		}
		return err
	}
	// 重复路径添加检查。
	if view.searchPaths.Search(realPath) != -1 {
		return nil
	}
	view.searchPaths.Clear()
	view.searchPaths.Append(realPath)
	view.fileCacheMap.Clear()
	return nil
}

// AddPath 将一个绝对路径或相对路径添加到搜索路径中。
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
	// realPath 应为文件夹类型。
	if !isDir {
		err := gerror.NewCodef(gcode.CodeInvalidParameter, `View.AddPath failed: path "%s" should be directory type`, path)
		if errorPrint() {
			glog.Error(ctx, err)
		}
		return err
	}
	// 重复路径添加检查。
	if view.searchPaths.Search(realPath) != -1 {
		return nil
	}
	view.searchPaths.Append(realPath)
	view.fileCacheMap.Clear()
	return nil
}

// Assigns 将多个全局模板变量绑定到当前视图对象。
// 注意，它不是并发安全的，这意味着如果在运行时多个goroutine中调用它，将会引发panic。
func (view *View) Assigns(data Params) {
	for k, v := range data {
		view.data[k] = v
	}
}

// Assign 将全局模板变量绑定到当前视图对象。
// 注意，它不是并发安全的，这意味着如果在运行时多个goroutine中调用它，将会导致panic。
func (view *View) Assign(key string, value interface{}) {
	view.data[key] = value
}

// SetDefaultFile 设置用于解析的默认模板文件。
func (view *View) SetDefaultFile(file string) {
	view.config.DefaultFile = file
}

// GetDefaultFile 返回用于解析的默认模板文件。
func (view *View) GetDefaultFile() string {
	return view.config.DefaultFile
}

// SetDelimiters 设置用于模板解析的自定义分隔符。
func (view *View) SetDelimiters(left, right string) {
	view.config.Delimiters = []string{left, right}
}

// SetAutoEncode 用于开启或关闭自动HTML编码功能。
// 当自动编码功能开启时，视图引擎会自动进行编码并提供安全的HTML输出，
// 这有助于避免XSS（跨站脚本攻击）漏洞。
func (view *View) SetAutoEncode(enable bool) {
	view.config.AutoEncode = enable
}

// BindFunc 注册一个名为 `name` 的自定义全局模板函数到当前视图对象，
// 使用给定的 `function` 函数。在模板内容中，`name` 是可以被调用的函数名。
func (view *View) BindFunc(name string, function interface{}) {
	view.funcMap[name] = function
	// Clear global template object cache.
	templates.Clear()
}

// BindFuncMap 通过映射注册自定义的全局模板函数到当前视图对象。
// 映射的键是模板函数名称
// 映射的值是自定义函数的地址。
func (view *View) BindFuncMap(funcMap FuncMap) {
	for k, v := range funcMap {
		view.funcMap[k] = v
	}
	// Clear global template object cache.
	templates.Clear()
}

// SetI18n 将 i18n 管理器绑定到当前视图引擎。
func (view *View) SetI18n(manager *gi18n.Manager) {
	view.config.I18nManager = manager
}
