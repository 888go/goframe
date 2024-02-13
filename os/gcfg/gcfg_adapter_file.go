// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 配置类

import (
	"context"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/command"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gfsnotify"
	"github.com/888go/goframe/os/gres"
	"github.com/888go/goframe/util/gmode"
	"github.com/888go/goframe/util/gutil"
)

// AdapterFile实现了使用文件的Adapter接口。
type AdapterFile struct {
	defaultName   string           // 默认配置文件名称。
	searchPaths   *数组类.StrArray // 搜索路径数组
	jsonMap       *map类.StrAnyMap  // 这是用于配置文件的解析后的JSON对象。
	violenceCheck bool             // 是否在值索引搜索时进行暴力检查。当设置为true（默认为false）时，会影响性能。
}

const (
	commandEnvKeyForFile = "gf.gcfg.file" // commandEnvKeyForFile 是用于配置命令行参数或环境配置文件名的配置键。
	commandEnvKeyForPath = "gf.gcfg.path" // commandEnvKeyForPath 是用于配置命令参数或环境目录路径的配置键。
)

var (
	supportedFileTypes     = []string{"toml", "yaml", "yml", "json", "ini", "xml", "properties"} // 所有支持的文件类型后缀。
	localInstances         = map类.X创建StrAny(true)                                             // Instances：包含配置实例的映射（map）。
	customConfigContentMap = map类.X创建StrStr(true)                                             // 自定义配置内容

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
		name = X默认配置文件名称
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
		searchPaths: 数组类.X创建文本(true),
		jsonMap:     map类.X创建StrAny(true),
	}
	// 从环境变量/命令行自定义目录路径。
	if customPath := command.GetOptWithEnv(commandEnvKeyForPath); customPath != "" {
		if 文件类.X是否存在(customPath) {
			if err = config.SetPath(customPath); err != nil {
				return nil, err
			}
		} else {
			return nil, 错误类.X创建并格式化(`configuration directory path "%s" does not exist`, customPath)
		}
	} else {
// ================================================================================
// 自动搜索目录
// 如果这些目录不存在，也不会影响适配器对象的创建。
// ================================================================================

		// Dir：工作目录的路径。
		if err = config.AddPath(文件类.X取当前工作目录()); err != nil {
			intlog.Errorf(context.TODO(), `%+v`, err)
		}

		// Dir：主包的路径。
		if mainPath := 文件类.X取main路径(); mainPath != "" && 文件类.X是否存在(mainPath) {
			if err = config.AddPath(mainPath); err != nil {
				intlog.Errorf(context.TODO(), `%+v`, err)
			}
		}

		// Dir 二进制文件的路径。
		if selfPath := 文件类.X取当前进程目录(); selfPath != "" && 文件类.X是否存在(selfPath) {
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
		return j.X取值(pattern).X取值(), nil
	}
	return nil, nil
}

// Set 通过指定的`pattern`设置值。
// 它支持使用字符分隔符（默认为'. '）进行层级数据访问。
// 通常用于在运行时更新特定配置值。
// 注意，不建议在运行时使用`Set`方法来配置，因为如果底层配置文件发生更改，
// 配置将会自动刷新。因此，直接运行时设置可能不会持久生效。
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

// Data 函数获取并以 map 类型返回所有配置数据。
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

// MustGet 行为类似于函数 Get，但在发生错误时会触发 panic。
func (a *AdapterFile) MustGet(ctx context.Context, pattern string) *泛型类.Var {
	v, err := a.Get(ctx, pattern)
	if err != nil {
		panic(err)
	}
	return 泛型类.X创建(v)
}

// 清除所有已解析的配置文件内容缓存，
// 这将强制从文件重新加载配置内容。
func (a *AdapterFile) Clear() {
	a.jsonMap.X清空()
}

// Dump 打印当前Json对象，使其更易于人工阅读。
func (a *AdapterFile) Dump() {
	if j, _ := a.getJson(); j != nil {
		j.X调试输出()
	}
}

// Available 检查并返回给定 `file` 配置是否可用。
func (a *AdapterFile) Available(ctx context.Context, fileName ...string) bool {
	checkFileName := 工具类.X取文本值或取默认值(a.defaultName, fileName...)
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
	if 环境类.IsDevelop() {
		mainPkgPath := 文件类.X取main路径()
		if mainPkgPath != "" {
			if !a.searchPaths.X是否存在(mainPkgPath) {
				a.searchPaths.Append别名(mainPkgPath)
			}
		}
	}
}

// getJson 函数返回指定 `file` 内容对应的 *gjson.Json 对象。
// 如果文件读取失败，会打印错误信息。若发生任何错误，将返回 nil。
func (a *AdapterFile) getJson(fileName ...string) (configJson *json类.Json, err error) {
	var (
		usedFileName = a.defaultName
	)
	if len(fileName) > 0 && fileName[0] != "" {
		usedFileName = fileName[0]
	} else {
		usedFileName = a.defaultName
	}
	// 它使用json映射来缓存指定配置文件的内容。
	result := a.jsonMap.X取值或设置值_函数带锁(usedFileName, func() interface{} {
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
			if file := 资源类.Get(filePath); file != nil {
				content = string(file.Content())
			} else {
				content = 文件类.X读文本(filePath)
			}
		}
		// 注意，底层的配置json对象操作是线程安全的。
		dataType := json类.ContentType(文件类.X路径取扩展名且不含点号(filePath))
		if json类.X检查类型(dataType) && !isFromConfigContent {
			configJson, err = json类.X加载并按格式(dataType, content, true)
		} else {
			configJson, err = json类.X加载并自动识别格式(content, true)
		}
		if err != nil {
			if filePath != "" {
				err = 错误类.X多层错误并格式化(err, `load config file "%s" failed`, filePath)
			} else {
				err = 错误类.X多层错误(err, `load configuration failed`)
			}
			return nil
		}
		configJson.X设置分层冲突检查(a.violenceCheck)
// 添加对这个配置文件的监控，
// 当该文件有任何变化时，都会在Config对象中刷新其缓存。
		if filePath != "" && !资源类.Contains(filePath) {
			_, err = 文件监控类.Add(filePath, func(event *文件监控类.Event) {
				a.jsonMap.X删除(usedFileName)
			})
			if err != nil {
				return nil
			}
		}
		return configJson
	})
	if result != nil {
		return result.(*json类.Json), err
	}
	return
}
