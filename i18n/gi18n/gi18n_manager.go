// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gi18n

import (
	"context"
	"fmt"
	"strings"
	"sync"
	
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gfsnotify"
	"github.com/888go/goframe/os/gres"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/util/gconv"
)

// pathType 是用于国际化文件路径的类型。
type pathType string

const (
	pathTypeNone   pathType = "none"
	pathTypeNormal pathType = "normal"
	pathTypeGres   pathType = "gres"
)

// i18n内容的管理器，它是并发安全的，并支持热重载。
type Manager struct {
	mu       sync.RWMutex
	data     map[string]map[string]string // Translating map.
	pattern  string                       // 正则表达式解析的模式。
	pathType pathType                     // Path 类型用于 i18n 文件。
	options  Options                      // 配置选项
}

// Options 用于i18n对象配置。
type Options struct {
	Path       string         // 国际化文件存储路径。
	Language   string         // 默认本地语言。
	Delimiters []string       // 变量解析的分隔符。
	Resource   *gres.Resource // i18n文件资源。
}

var (
	// defaultDelimiters 定义了如果用户在选项中未指定时的默认分隔符语言。
	defaultLanguage = "en"

	// defaultDelimiters 定义了默认的关键字变量分隔符。
	defaultDelimiters = []string{"{#", "}"}

	// i18n文件搜索目录
	searchFolders = []string{"manifest/i18n", "manifest/config/i18n", "i18n"}
)

// New 创建并返回一个新的 i18n 管理器。
// 可选参数 `option` 指定 i18n 管理器的自定义选项。
// 若未传递该参数，则使用默认选项。
func New(options ...Options) *Manager {
	var opts Options
	var pathType = pathTypeNone
	if len(options) > 0 {
		opts = options[0]
		pathType = opts.checkPathType(opts.Path)
	} else {
		opts = Options{}
		for _, folder := range searchFolders {
			pathType = opts.checkPathType(folder)
			if pathType != pathTypeNone {
				break
			}
		}
		if opts.Path != "" {
			// 为避免引用GoFrame的源路径：github.com/gogf/i18n/gi18n
			if gfile.Exists(opts.Path + gfile.Separator + "gi18n") {
				opts.Path = ""
				pathType = pathTypeNone
			}
		}
	}
	if len(opts.Language) == 0 {
		opts.Language = defaultLanguage
	}
	if len(opts.Delimiters) == 0 {
		opts.Delimiters = defaultDelimiters
	}
	m := &Manager{
		options: opts,
		pattern: fmt.Sprintf(
			`%s(.+?)%s`,
			gregex.Quote(opts.Delimiters[0]),
			gregex.Quote(opts.Delimiters[1]),
		),
		pathType: pathType,
	}
	intlog.Printf(context.TODO(), `New: %#v`, m)
	return m
}

// checkPathType 对给定的目录路径进行检查并返回其路径类型。
func (o *Options) checkPathType(dirPath string) pathType {
	if dirPath == "" {
		return pathTypeNone
	}

	if o.Resource == nil {
		o.Resource = gres.Instance()
	}

	if o.Resource.Contains(dirPath) {
		o.Path = dirPath
		return pathTypeGres
	}

	realPath, _ := gfile.Search(dirPath)
	if realPath != "" {
		o.Path = realPath
		return pathTypeNormal
	}

	return pathTypeNone
}

// SetPath 设置存储 i18n 文件的目录路径。
func (m *Manager) SetPath(path string) error {
	pathType := m.options.checkPathType(path)
	if pathType == pathTypeNone {
		return gerror.NewCodef(gcode.CodeInvalidParameter, `%s does not exist`, path)
	}

	m.pathType = pathType
	intlog.Printf(context.TODO(), `SetPath[%s]: %s`, m.pathType, m.options.Path)
	// 在路径改变后重置管理器。
	m.reset()
	return nil
}

// SetLanguage 设置翻译器的语言。
func (m *Manager) SetLanguage(language string) {
	m.options.Language = language
	intlog.Printf(context.TODO(), `SetLanguage: %s`, m.options.Language)
}

// SetDelimiters 设置翻译器的分隔符。
func (m *Manager) SetDelimiters(left, right string) {
	m.pattern = fmt.Sprintf(`%s(.+?)%s`, gregex.Quote(left), gregex.Quote(right))
	intlog.Printf(context.TODO(), `SetDelimiters: %v`, m.pattern)
}

// T 是 Translate 的别名，用于提供便利。
func (m *Manager) T(ctx context.Context, content string) string {
	return m.Translate(ctx, content)
}

// Tf 是 TranslateFormat 的别名，用于提供便利。
func (m *Manager) Tf(ctx context.Context, format string, values ...interface{}) string {
	return m.TranslateFormat(ctx, format, values...)
}

// TranslateFormat 将根据配置的语言和给定的 `values` 对 `format` 进行翻译、格式化并返回结果。
func (m *Manager) TranslateFormat(ctx context.Context, format string, values ...interface{}) string {
	return fmt.Sprintf(m.Translate(ctx, format), values...)
}

// Translate 使用配置的语言对`content`进行翻译。
func (m *Manager) Translate(ctx context.Context, content string) string {
	m.init(ctx)
	m.mu.RLock()
	defer m.mu.RUnlock()
	transLang := m.options.Language
	if lang := LanguageFromCtx(ctx); lang != "" {
		transLang = lang
	}
	data := m.data[transLang]
	if data == nil {
		return content
	}
	// 将内容解析为名称。
	if v, ok := data[content]; ok {
		return v
	}
	// 将内容解析为变量容器。
	result, _ := gregex.ReplaceStringFuncMatch(
		m.pattern, content,
		func(match []string) string {
			if v, ok := data[match[1]]; ok {
				return v
			}
// 返回match[1]将返回分隔符之间的内容
// 返回match[0]将返回原始内容
			return match[0]
		})
	intlog.Printf(ctx, `Translate for language: %s`, transLang)
	return result
}

// GetContent 函数根据给定的键和指定的语言获取并返回配置的内容。
// 如果未找到，则返回一个空字符串。
func (m *Manager) GetContent(ctx context.Context, key string) string {
	m.init(ctx)
	m.mu.RLock()
	defer m.mu.RUnlock()
	transLang := m.options.Language
	if lang := LanguageFromCtx(ctx); lang != "" {
		transLang = lang
	}
	if data, ok := m.data[transLang]; ok {
		return data[key]
	}
	return ""
}

// reset 重置管理器的数据。
func (m *Manager) reset() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data = nil
}

// init 用于实现延迟初始化设计，初始化i18n管理器。
// i18n管理器只初始化一次。
func (m *Manager) init(ctx context.Context) {
	m.mu.RLock()
	// 如果数据不为nil，表示它已经被初始化过。
	if m.data != nil {
		m.mu.RUnlock()
		return
	}
	m.mu.RUnlock()

	defer func() {
		intlog.Printf(ctx, `Manager init finish: %#v`, m)
	}()

	intlog.Printf(ctx, `init path: %s`, m.options.Path)

	m.mu.Lock()
	defer m.mu.Unlock()
	switch m.pathType {
	case pathTypeGres:
		files := m.options.Resource.ScanDirFile(m.options.Path, "*.*", true)
		if len(files) > 0 {
			var (
				path  string
				name  string
				lang  string
				array []string
			)
			m.data = make(map[string]map[string]string)
			for _, file := range files {
				name = file.Name()
				path = name[len(m.options.Path)+1:]
				array = strings.Split(path, "/")
				if len(array) > 1 {
					lang = array[0]
				} else if len(array) == 1 {
					lang = gfile.Name(array[0])
				}
				if m.data[lang] == nil {
					m.data[lang] = make(map[string]string)
				}
				if j, err := gjson.LoadContent(file.Content()); err == nil {
					for k, v := range j.Var().Map() {
						m.data[lang][k] = gconv.String(v)
					}
				} else {
					intlog.Errorf(ctx, "load i18n file '%s' failed: %+v", name, err)
				}
			}
		}
	case pathTypeNormal:
		files, _ := gfile.ScanDirFile(m.options.Path, "*.*", true)
		if len(files) == 0 {
			return
		}
		var (
			path  string
			lang  string
			array []string
		)
		m.data = make(map[string]map[string]string)
		for _, file := range files {
			path = file[len(m.options.Path)+1:]
			array = strings.Split(path, gfile.Separator)
			if len(array) > 1 {
				lang = array[0]
			} else if len(array) == 1 {
				lang = gfile.Name(array[0])
			}
			if m.data[lang] == nil {
				m.data[lang] = make(map[string]string)
			}
			if j, err := gjson.LoadContent(gfile.GetBytes(file)); err == nil {
				for k, v := range j.Var().Map() {
					m.data[lang][k] = gconv.String(v)
				}
			} else {
				intlog.Errorf(ctx, "load i18n file '%s' failed: %+v", file, err)
			}
		}
		intlog.Printf(ctx, "i18n files loaded in path: %s", m.options.Path)
		// 监控i18n文件的更改以实现热重载功能。
		_, _ = gfsnotify.Add(m.options.Path, func(event *gfsnotify.Event) {
			intlog.Printf(ctx, `i18n file changed: %s`, event.Path)
			// 如果i18n文件有任何更改，清空数据。
			m.reset()
			gfsnotify.Exit()
		})
	}
}
