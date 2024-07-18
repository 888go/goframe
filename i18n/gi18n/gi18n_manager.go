// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gi18n

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gfsnotify"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/util/gconv"
)

// pathType 是用于i18n文件路径的类型。 md5:1aa056f2406cd3a6
type pathType string

const (
	pathTypeNone   pathType = "none"
	pathTypeNormal pathType = "normal"
	pathTypeGres   pathType = "gres"
)

// i18n内容的管理器，它是并发安全的，支持热重载。 md5:9c519435bec8f5ad
type Manager struct {
	mu       sync.RWMutex
	data     map[string]map[string]string // Translating map.
	pattern  string                       // 正则表达式解析的模式。 md5:1d0109d4850fd141
	pathType pathType                     // i18n 文件的路径类型。 md5:5b086fc46e36729e
	options  Options                      // configuration options.
}

// Options 用于国际化对象的配置。 md5:029f81136c5c3e6a
type Options struct {
	Path       string         // 国际化文件的存储路径。 md5:67cec25950dc6464
	Language   string         // 默认本地语言。 md5:41ccf6a6028cf49d
	Delimiters []string       // 变量解析的定界符。 md5:355db5afb17acaf5
	Resource   *gres.Resource // i18n文件的资源。 md5:611cd5c408223400
}

var (
	// defaultLanguage 定义了如果用户在选项中未指定，默认的语言。 md5:37b426a695c48d49
	defaultLanguage = "en"

	// defaultDelimiters 定义了默认的键变量分隔符。 md5:98706258206bfd9a
	defaultDelimiters = []string{"{#", "}"}

	// 国际化文件搜索目录。 md5:cf8914abf6ec0557
	searchFolders = []string{"manifest/i18n", "manifest/config/i18n", "i18n"}
)

// New 创建并返回一个新的国际化管理器。
// 可选参数 `option` 用于指定国际化管理器的自定义选项。
// 如果未传递该参数，它将使用默认选项。
// md5:79f31dcd2ff8cf56
// ff:
// options:
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
			// 为了避免GoFrame的源路径：github.com/gogf/i18n/gi18n. md5:2eecc4478ca65bd7
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

// checkPathType 检查并返回给定目录路径的路径类型。 md5:101af7b8de6f50f8
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

// SetPath 设置存储i18n文件的目录路径。 md5:b39e1d244949dcf8
// ff:
// m:
// path:
func (m *Manager) SetPath(path string) error {
	pathType := m.options.checkPathType(path)
	if pathType == pathTypeNone {
		return gerror.NewCodef(gcode.CodeInvalidParameter, `%s does not exist`, path)
	}

	m.pathType = pathType
	intlog.Printf(context.TODO(), `SetPath[%s]: %s`, m.pathType, m.options.Path)
	// 路径改变后重置管理器。 md5:1f0260d8d112184d
	m.reset()
	return nil
}

// SetLanguage 设置翻译器的语言。 md5:50b09b0bb0944dc1
// ff:
// m:
// language:
func (m *Manager) SetLanguage(language string) {
	m.options.Language = language
	intlog.Printf(context.TODO(), `SetLanguage: %s`, m.options.Language)
}

// SetDelimiters 为翻译器设置分隔符。 md5:f84b046b11204dc7
// ff:
// m:
// left:
// right:
func (m *Manager) SetDelimiters(left, right string) {
	m.pattern = fmt.Sprintf(`%s(.+?)%s`, gregex.Quote(left), gregex.Quote(right))
	intlog.Printf(context.TODO(), `SetDelimiters: %v`, m.pattern)
}

// T 是为了方便而对 Translate 的别名。 md5:c07a6fa99a429eb3
// ff:
// m:
// ctx:
// content:
func (m *Manager) T(ctx context.Context, content string) string {
	return m.Translate(ctx, content)
}

// Tf是TranslateFormat的别名，为了方便起见。 md5:bdb209b24c669f5a
// ff:
// m:
// ctx:
// format:
// values:
func (m *Manager) Tf(ctx context.Context, format string, values ...interface{}) string {
	return m.TranslateFormat(ctx, format, values...)
}

// TranslateFormat 使用配置的语言和给定的 `values` 对 `format` 进行翻译、格式化并返回结果。
// md5:2806a81d6db86c7f
// ff:
// m:
// ctx:
// format:
// values:
func (m *Manager) TranslateFormat(ctx context.Context, format string, values ...interface{}) string {
	return fmt.Sprintf(m.Translate(ctx, format), values...)
}

// Translate 使用配置的语言翻译`content`。 md5:8f8b7d32e0b26a99
// ff:
// m:
// ctx:
// content:
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
	// Parse content as name.
	if v, ok := data[content]; ok {
		return v
	}
	// 解析内容作为变量容器。 md5:6fab6ca886fe327a
	result, _ := gregex.ReplaceStringFuncMatch(
		m.pattern, content,
		func(match []string) string {
			if v, ok := data[match[1]]; ok {
				return v
			}
// 返回match[1] 将返回分隔符之间的内容
// 返回match[0] 将返回原始内容
// md5:3dd48230b02f1348
			return match[0]
		})
	intlog.Printf(ctx, `Translate for language: %s`, transLang)
	return result
}

// GetContent 获取并返回给定键和指定语言的配置内容。
// 如果未找到，将返回一个空字符串。
// md5:c64a3a803ac07e38
// ff:
// m:
// ctx:
// key:
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

// reset 重置管理器的数据。 md5:582ac65a0b066583
func (m *Manager) reset() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data = nil
}

// init 初始化管理器，用于延迟初始化设计。
// 国际化(i18n)管理器仅初始化一次。
// md5:b3e5cf7f018d1485
func (m *Manager) init(ctx context.Context) {
	m.mu.RLock()
	// 如果数据不为nil，表示它已经初始化。 md5:8d4d8b324fc9951a
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
		// 监控i18n文件的变化，以实现热重载功能。 md5:feeb4e0abb048a7b
		_, _ = gfsnotify.Add(m.options.Path, func(event *gfsnotify.Event) {
			intlog.Printf(ctx, `i18n file changed: %s`, event.Path)
			// 对i18n文件的任何更改，都会清空数据。 md5:fbcc6de55a881c92
			m.reset()
			gfsnotify.Exit()
		})
	}
}
