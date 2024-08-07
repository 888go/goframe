// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 配置类

import (
	"context"

	garray "github.com/888go/goframe/container/garray"
	gmap "github.com/888go/goframe/container/gmap"
	gvar "github.com/888go/goframe/container/gvar"
	gjson "github.com/888go/goframe/encoding/gjson"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/command"
	"github.com/888go/goframe/internal/intlog"
	gfile "github.com/888go/goframe/os/gfile"
	gfsnotify "github.com/888go/goframe/os/gfsnotify"
	gres "github.com/888go/goframe/os/gres"
	gmode "github.com/888go/goframe/util/gmode"
	gutil "github.com/888go/goframe/util/gutil"
)

// AdapterFile 实现了使用文件的 Adapter 接口。 md5:c0f0e0b1d4b217fd
type AdapterFile struct {
	defaultName   string           // 默认配置文件名。 md5:30af4aed147cf623
	searchPaths   *garray.StrArray // Searching path array.
	jsonMap       *gmap.StrAnyMap  // 配置文件中使用的简化JSON对象。 md5:dc2b385b92cd7edc
	violenceCheck bool             // 是否在值索引搜索中进行暴力检查。设置为true（默认为false）会影响性能。 md5:b2ea5ca1ded97be3
}

const (
	commandEnvKeyForFile = "gf.gcfg.file" // commandEnvKeyForFile是用于命令参数或环境配置文件名的配置键。 md5:32e2ea36b81b5269
	commandEnvKeyForPath = "gf.gcfg.path" // commandEnvKeyForPath 是用于命令参数或配置目录路径的配置键。 md5:0e1e71d5290a8c3c
)

var (
	supportedFileTypes     = []string{"toml", "yaml", "yml", "json", "ini", "xml", "properties"} // 所支持的文件类型后缀。 md5:3609c8928b780170
	localInstances         = gmap.X创建StrAny(true)                                             // Instances映射，其中包含配置实例。 md5:df7e552f8e970f97
	customConfigContentMap = gmap.X创建StrStr(true)                                             // 定制化配置内容。 md5:e408d212ab61e310

		// 用于在资源管理器中尝试搜索的前缀数组。 md5:f69485b110ee7be3
	resourceTryFolders = []string{
		"", "/", "config/", "config", "/config", "/config/",
		"manifest/config/", "manifest/config", "/manifest/config", "/manifest/config/",
	}

		// 前缀数组，用于在本地系统中尝试搜索。 md5:51a8f1255f95f3fc
	localSystemTryFolders = []string{"", "config/", "manifest/config"}
)

// NewAdapterFile 返回一个新的配置管理对象。
// 参数 `file` 指定了默认的配置文件读取名称。
// md5:52ab633a98562ceb
func NewAdapterFile(file ...string) (*AdapterFile, error) {
	var (
		err  error
		name = X默认配置文件名称
	)
	if len(file) > 0 {
		name = file[0]
	} else {
				// 从命令行或环境变量中获取自定义的默认配置文件名。 md5:d43279fee761ac4d
		if customFile := command.GetOptWithEnv(commandEnvKeyForFile); customFile != "" {
			name = customFile
		}
	}
	config := &AdapterFile{
		defaultName: name,
		searchPaths: garray.X创建文本(true),
		jsonMap:     gmap.X创建StrAny(true),
	}
		// 从环境变量或命令行自定义的目录路径。 md5:8cfcbca968e23c5b
	if customPath := command.GetOptWithEnv(commandEnvKeyForPath); customPath != "" {
		if gfile.X是否存在(customPath) {
			if err = config.SetPath(customPath); err != nil {
				return nil, err
			}
		} else {
			return nil, gerror.X创建并格式化(`configuration directory path "%s" does not exist`, customPath)
		}
	} else {
		// =================================================================================
		// 自动搜索目录。
		// 如果这些目录不存在，不影响适配器对象的创建。
		// =================================================================================
		// md5:08a226598ce0311e

				// Dir 是工作目录的路径。 md5:0fba211853ea97a0
		if err = config.AddPath(gfile.X取当前工作目录()); err != nil {
			intlog.Errorf(context.TODO(), `%+v`, err)
		}

				// 主包的目录路径。 md5:a4d2802779172abe
		if mainPath := gfile.X取main路径(); mainPath != "" && gfile.X是否存在(mainPath) {
			if err = config.AddPath(mainPath); err != nil {
				intlog.Errorf(context.TODO(), `%+v`, err)
			}
		}

		// Dir path of binary.
		if selfPath := gfile.X取当前进程目录(); selfPath != "" && gfile.X是否存在(selfPath) {
			if err = config.AddPath(selfPath); err != nil {
				intlog.Errorf(context.TODO(), `%+v`, err)
			}
		}
	}
	return config, nil
}

// SetViolenceCheck 设置是否执行层次冲突检查。
// 当键名中包含级别符号时，需要启用此功能。默认情况下禁用。
// 
// 注意，开启此功能的开销较大，并不建议在键名中使用分隔符。最好在应用层面上避免这种情况。
// md5:5009f694ccd4efc0
func (a *AdapterFile) SetViolenceCheck(check bool) {
	a.violenceCheck = check
	a.Clear()
}

// SetFileName 设置默认的配置文件名。 md5:b540171ead70ddf8
func (a *AdapterFile) SetFileName(name string) {
	a.defaultName = name
}

// GetFileName 返回默认的配置文件名。 md5:d13e3bd27526f03d
func (a *AdapterFile) GetFileName() string {
	return a.defaultName
}

// Get通过指定的`pattern`获取并返回值。如果`pattern`为空或为"."，则返回当前Json对象的所有值。如果没有找到匹配`pattern`的值，它将返回nil。
// 
// 我们也可以通过在`pattern`中使用索引来访问切片项，例如："list.10"，"array.0.name"，"array.0.1.id"。
// 
// 如果没有找到与`pattern`匹配的值，它将返回由`def`指定的默认值。
// md5:8a88d01912ac6218
func (a *AdapterFile) Get(ctx context.Context, pattern string) (value interface{}, err error) {
	j, err := a.getJson()
	if err != nil {
		return nil, err
	}
	if j != nil {
		return j.X取值(pattern).X取值(), nil
	}
	return nil, nil
}

// X设置值 使用指定的 `pattern` 设置值。
// 它支持通过字符分隔符（默认为`.`）进行层次数据访问。
// 这通常用于在运行时更新特定配置值。
// 请注意，不建议在运行时使用 `X设置值` 配置，因为如果底层配置文件更改，配置会自动刷新。
// md5:65992c2815af747e
func (a *AdapterFile) X设置值(pattern string, value interface{}) error {
	j, err := a.getJson()
	if err != nil {
		return err
	}
	if j != nil {
		return j.X设置值(pattern, value)
	}
	return nil
}

// Data 获取并以映射类型返回所有配置数据。 md5:2a92e8bbe7388f01
func (a *AdapterFile) Data(ctx context.Context) (data map[string]interface{}, err error) {
	j, err := a.getJson()
	if err != nil {
		return nil, err
	}
	if j != nil {
		return j.X取泛型类().X取Map(), nil
	}
	return nil, nil
}

// MustGet 行为类似于函数 Get，但如果发生错误时会引发 panic。 md5:b1d3af83a52fd248
func (a *AdapterFile) MustGet(ctx context.Context, pattern string) *gvar.Var {
	v, err := a.Get(ctx, pattern)
	if err != nil {
		panic(err)
	}
	return gvar.X创建(v)
}

// Clear 清除所有解析的配置文件内容缓存，这将强制重新从文件加载配置内容。
// md5:5868c636ce62cb14
func (a *AdapterFile) Clear() {
	a.jsonMap.X清空()
}

// Dump 打印当前的Json对象，使其更便于人工阅读。 md5:c8c6bbdb40fa6383
func (a *AdapterFile) Dump() {
	if j, _ := a.getJson(); j != nil {
		j.X调试输出()
	}
}

// 可用检查并返回给定`file`的配置是否可用。 md5:d915d3cb575cbd5b
func (a *AdapterFile) Available(ctx context.Context, fileName ...string) bool {
	checkFileName := gutil.X取文本值或取默认值(a.defaultName, fileName...)
		// 存在自定义配置内容。 md5:50d226a12b07427d
	if a.GetContent(checkFileName) != "" {
		return true
	}
		// 配置文件存在于系统路径中。 md5:a32283fd4eff7ddf
	if path, _ := a.GetFilePath(checkFileName); path != "" {
		return true
	}
	return false
}

// autoCheckAndAddMainPkgPathToSearchPaths 自动检查并添加当前开发环境中的"main"包目录路径到搜索路径列表中。
// md5:4a1366fa2d1d98ab
func (a *AdapterFile) autoCheckAndAddMainPkgPathToSearchPaths() {
	if gmode.IsDevelop() {
		mainPkgPath := gfile.X取main路径()
		if mainPkgPath != "" {
			if !a.searchPaths.X是否存在(mainPkgPath) {
				a.searchPaths.Append别名(mainPkgPath)
			}
		}
	}
}

// getJson 为指定的`file`内容返回一个*gjson.Json*对象。
// 如果文件读取失败，它会打印错误。如果发生任何错误，它将返回nil。
// md5:ffbc3e1a6ff12753
func (a *AdapterFile) getJson(fileName ...string) (configJson *gjson.Json, err error) {
	var (
		usedFileName = a.defaultName
	)
	if len(fileName) > 0 && fileName[0] != "" {
		usedFileName = fileName[0]
	} else {
		usedFileName = a.defaultName
	}
		// 它使用json映射来缓存指定的配置文件内容。 md5:70b9eac1f3ac38b4
	result := a.jsonMap.X取值或设置值_函数带锁(usedFileName, func() interface{} {
		var (
			content  string
			filePath string
		)
				// 配置的内容可以是与其文件类型不同的任何数据类型。 md5:11fb8ecd6511ef10
		isFromConfigContent := true
		if content = a.GetContent(usedFileName); content == "" {
			isFromConfigContent = false
			filePath, err = a.GetFilePath(usedFileName)
			if err != nil {
				return nil
			}
			if filePath == "" {
				return nil
			}
			if file := gres.Get(filePath); file != nil {
				content = string(file.Content())
			} else {
				content = gfile.X读文本(filePath)
			}
		}
				// 注意，底层的配置json对象操作是并发安全的。 md5:2cd371ca691286f9
		dataType := gjson.ContentType(gfile.X路径取扩展名且不含点号(filePath))
		if gjson.X检查类型(dataType) && !isFromConfigContent {
			configJson, err = gjson.X加载并按格式(dataType, content, true)
		} else {
			configJson, err = gjson.X加载并自动识别格式(content, true)
		}
		if err != nil {
			if filePath != "" {
				err = gerror.X多层错误并格式化(err, `load config file "%s" failed`, filePath)
			} else {
				err = gerror.X多层错误(err, `load configuration failed`)
			}
			return nil
		}
		configJson.X设置分层冲突检查(a.violenceCheck)
		// 为这个配置文件添加监控，
		// 该文件的任何更改都会刷新Config对象中的缓存。
		// md5:8520fe419f2d8cc1
		if filePath != "" && !gres.Contains(filePath) {
			_, err = gfsnotify.Add(filePath, func(event *gfsnotify.Event) {
				a.jsonMap.X删除(usedFileName)
			})
			if err != nil {
				return nil
			}
		}
		return configJson
	})
	if result != nil {
		return result.(*gjson.Json), err
	}
	return
}
