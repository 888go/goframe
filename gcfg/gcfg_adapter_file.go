// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcfg

import (
	"context"

	"github.com/888go/goframe/gcfg/internal/command"
	"github.com/888go/goframe/gcfg/internal/intlog"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gfsnotify"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/util/gmode"
	"github.com/gogf/gf/v2/util/gutil"
)

// AdapterFile实现了使用文件的Adapter接口。
type AdapterFile struct {
	defaultName   string           // 默认配置文件名称。
	searchPaths   *garray.StrArray // 搜索路径数组
	jsonMap       *gmap.StrAnyMap  // 这是用于配置文件的解析后的JSON对象。
	violenceCheck bool             // 是否在值索引搜索时进行暴力检查。当设置为true（默认为false）时，会影响性能。
}

const (
	commandEnvKeyForFile = "gf.gcfg.file" // commandEnvKeyForFile 是用于配置命令行参数或环境配置文件名的配置键。
	commandEnvKeyForPath = "gf.gcfg.path" // commandEnvKeyForPath 是用于配置命令参数或环境目录路径的配置键。
)

var (
	supportedFileTypes     = []string{"toml", "yaml", "yml", "json", "ini", "xml", "properties"} // 所有支持的文件类型后缀。
	localInstances         = gmap.NewStrAnyMap(true)                                             // Instances：包含配置实例的映射（map）。
	customConfigContentMap = gmap.NewStrStrMap(true)                                             // 自定义配置内容

	// 前缀数组，用于在资源管理器中尝试搜索。
	resourceTryFolders = []string{
		"", "/", "config/", "config", "/config", "/config/",
		"manifest/config/", "manifest/config", "/manifest/config", "/manifest/config/",
	}

	// 前缀数组，用于尝试在本地系统中搜索。
	localSystemTryFolders = []string{"", "config/", "manifest/config"}
)

// NewAdapterFile 返回一个新的配置管理对象。
// 参数`file`指定了用于读取的默认配置文件名。
func NewAdapterFile(file ...string) (*AdapterFile, error) {
	var (
		err  error
		name = DefaultConfigFileName
	)
	if len(file) > 0 {
		name = file[0]
	} else {
		// 从命令行或环境变量自定义默认配置文件名。
		if customFile := command.GetOptWithEnv(commandEnvKeyForFile); customFile != "" {
			name = customFile
		}
	}
	config := &AdapterFile{
		defaultName: name,
		searchPaths: garray.NewStrArray(true),
		jsonMap:     gmap.NewStrAnyMap(true),
	}
	// 从环境变量/命令行自定义目录路径。
	if customPath := command.GetOptWithEnv(commandEnvKeyForPath); customPath != "" {
		if gfile.Exists(customPath) {
			if err = config.SetPath(customPath); err != nil {
				return nil, err
			}
		} else {
			return nil, gerror.Newf(`configuration directory path "%s" does not exist`, customPath)
		}
	} else {
// ================================================================================
// 自动搜索目录
// 如果这些目录不存在，也不会影响适配器对象的创建。
// ================================================================================

		// Dir：工作目录的路径。
		if err = config.AddPath(gfile.Pwd()); err != nil {
			intlog.Errorf(context.TODO(), `%+v`, err)
		}

		// Dir：主包的路径。
		if mainPath := gfile.MainPkgPath(); mainPath != "" && gfile.Exists(mainPath) {
			if err = config.AddPath(mainPath); err != nil {
				intlog.Errorf(context.TODO(), `%+v`, err)
			}
		}

		// Dir 二进制文件的路径。
		if selfPath := gfile.SelfDir(); selfPath != "" && gfile.Exists(selfPath) {
			if err = config.AddPath(selfPath); err != nil {
				intlog.Errorf(context.TODO(), `%+v`, err)
			}
		}
	}
	return config, nil
}

// SetViolenceCheck 设置是否进行层级冲突检查。
// 当键名中存在层级符号时，需要开启此功能。默认情况下，该功能是关闭的。
//
// 注意，开启此特性代价较高，且不建议在键名中允许分隔符。最好在应用层面避免这种情况。
func (a *AdapterFile) SetViolenceCheck(check bool) {
	a.violenceCheck = check
	a.Clear()
}

// SetFileName 设置默认的配置文件名。
func (a *AdapterFile) SetFileName(name string) {
	a.defaultName = name
}

// GetFileName 返回默认配置文件名称。
func (a *AdapterFile) GetFileName() string {
	return a.defaultName
}

// Get 方法通过指定的`pattern`获取并返回值。
// 如果`pattern`为空字符串或"."，则返回当前Json对象的所有值。
// 若通过`pattern`未找到任何值，则返回nil。
//
// 我们还可以通过在`pattern`中使用索引号访问切片元素，例如：
// "list.10", "array.0.name", "array.0.1.id"。
//
// 如果根据`pattern`未能找到对应的值，则返回由`def`指定的默认值。
func (a *AdapterFile) Get(ctx context.Context, pattern string) (value interface{}, err error) {
	j, err := a.getJson()
	if err != nil {
		return nil, err
	}
	if j != nil {
		return j.Get(pattern).Val(), nil
	}
	return nil, nil
}

// Set 通过指定的`pattern`设置值。
// 它支持使用字符分隔符（默认为'. '）进行层级数据访问。
// 通常用于在运行时更新特定配置值。
// 注意，不建议在运行时使用`Set`方法来配置，因为如果底层配置文件发生更改，
// 配置将会自动刷新。因此，直接运行时设置可能不会持久生效。
func (a *AdapterFile) Set(pattern string, value interface{}) error {
	j, err := a.getJson()
	if err != nil {
		return err
	}
	if j != nil {
		return j.Set(pattern, value)
	}
	return nil
}

// Data 函数获取并以 map 类型返回所有配置数据。
func (a *AdapterFile) Data(ctx context.Context) (data map[string]interface{}, err error) {
	j, err := a.getJson()
	if err != nil {
		return nil, err
	}
	if j != nil {
		return j.Var().Map(), nil
	}
	return nil, nil
}

// MustGet 行为类似于函数 Get，但在发生错误时会触发 panic。
func (a *AdapterFile) MustGet(ctx context.Context, pattern string) *gvar.Var {
	v, err := a.Get(ctx, pattern)
	if err != nil {
		panic(err)
	}
	return gvar.New(v)
}

// 清除所有已解析的配置文件内容缓存，
// 这将强制从文件重新加载配置内容。
func (a *AdapterFile) Clear() {
	a.jsonMap.Clear()
}

// Dump 打印当前Json对象，使其更易于人工阅读。
func (a *AdapterFile) Dump() {
	if j, _ := a.getJson(); j != nil {
		j.Dump()
	}
}

// Available 检查并返回给定 `file` 配置是否可用。
func (a *AdapterFile) Available(ctx context.Context, fileName ...string) bool {
	checkFileName := gutil.GetOrDefaultStr(a.defaultName, fileName...)
	// 自定义配置内容存在。
	if a.GetContent(checkFileName) != "" {
		return true
	}
	// 配置文件存在于系统路径中。
	if path, _ := a.GetFilePath(checkFileName); path != "" {
		return true
	}
	return false
}

// autoCheckAndAddMainPkgPathToSearchPaths 自动检测并添加 main 包的目录路径到搜索路径列表中，
// 如果当前处于开发环境的话。
func (a *AdapterFile) autoCheckAndAddMainPkgPathToSearchPaths() {
	if gmode.IsDevelop() {
		mainPkgPath := gfile.MainPkgPath()
		if mainPkgPath != "" {
			if !a.searchPaths.Contains(mainPkgPath) {
				a.searchPaths.Append(mainPkgPath)
			}
		}
	}
}

// getJson 函数返回指定 `file` 内容对应的 *gjson.Json 对象。
// 如果文件读取失败，会打印错误信息。若发生任何错误，将返回 nil。
func (a *AdapterFile) getJson(fileName ...string) (configJson *gjson.Json, err error) {
	var (
		usedFileName = a.defaultName
	)
	if len(fileName) > 0 && fileName[0] != "" {
		usedFileName = fileName[0]
	} else {
		usedFileName = a.defaultName
	}
	// 它使用json映射来缓存指定配置文件的内容。
	result := a.jsonMap.GetOrSetFuncLock(usedFileName, func() interface{} {
		var (
			content  string
			filePath string
		)
		// 配置的内容可以是与文件类型不同的任何数据类型。
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
				content = gfile.GetContents(filePath)
			}
		}
		// 注意，底层的配置json对象操作是线程安全的。
		dataType := gjson.ContentType(gfile.ExtName(filePath))
		if gjson.IsValidDataType(dataType) && !isFromConfigContent {
			configJson, err = gjson.LoadContentType(dataType, content, true)
		} else {
			configJson, err = gjson.LoadContent(content, true)
		}
		if err != nil {
			if filePath != "" {
				err = gerror.Wrapf(err, `load config file "%s" failed`, filePath)
			} else {
				err = gerror.Wrap(err, `load configuration failed`)
			}
			return nil
		}
		configJson.SetViolenceCheck(a.violenceCheck)
// 添加对这个配置文件的监控，
// 当该文件有任何变化时，都会在Config对象中刷新其缓存。
		if filePath != "" && !gres.Contains(filePath) {
			_, err = gfsnotify.Add(filePath, func(event *gfsnotify.Event) {
				a.jsonMap.Remove(usedFileName)
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
