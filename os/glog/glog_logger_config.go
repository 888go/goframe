// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package glog

import (
	"context"
	"io"
	"strings"
	"time"
	
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
)

// Config 是 logger 的配置对象。
type Config struct {
	Handlers             []Handler      `json:"-"`                    // 日志处理器实现类似中间件的功能。
	Writer               io.Writer      `json:"-"`                    // 自定义 io.Writer。
	Flags                int            `json:"flags"`                // 附加日志输出功能的标志。
	TimeFormat           string         `json:"timeFormat"`           // 日志时间格式
	Path                 string         `json:"path"`                 // 日志目录路径。
	File                 string         `json:"file"`                 // 日志文件的格式化模式。
	Level                int            `json:"level"`                // Output level.
	Prefix               string         `json:"prefix"`               // Prefix 字符串，用于作为每条日志内容的前缀。
	StSkip               int            `json:"stSkip"`               // 跳过堆栈的计数。
	StStatus             int            `json:"stStatus"`             // 栈状态(1: 启用 - 默认值; 0: 禁用)
	StFilter             string         `json:"stFilter"`             // 字符串过滤栈
	CtxKeys              []interface{}  `json:"ctxKeys"`              // Context keys 用于日志记录，它们被用于从context中检索值。
	HeaderPrint          bool           `json:"header"`               // 是否打印头部信息（默认为true）
	StdoutPrint          bool           `json:"stdout"`               // 是否输出到标准输出（默认为true）
	LevelPrint           bool           `json:"levelPrint"`           // 是否打印级别格式化字符串（默认为true）
	LevelPrefixes        map[int]string `json:"levelPrefixes"`        // 日志级别到其前缀字符串的映射。
	RotateSize           int64          `json:"rotateSize"`           // 如果日志文件大小 > 0字节，则旋转日志文件。
	RotateExpire         time.Duration  `json:"rotateExpire"`         // 如果日志文件的修改时间超过这个持续时间，则旋转日志文件。
	RotateBackupLimit    int            `json:"rotateBackupLimit"`    // Max 备份文件数量，默认为0，表示不进行备份。
	RotateBackupExpire   time.Duration  `json:"rotateBackupExpire"`   // Max 为轮转文件设置的过期时间，默认为0，表示永不过期。
	RotateBackupCompress int            `json:"rotateBackupCompress"` // Compress level for rotated files using gzip algorithm. It's 0 in default, means no compression.
	RotateCheckInterval  time.Duration  `json:"rotateCheckInterval"`  // Asynchronously checks the backups and expiration at intervals. It's 1 hour in default.
	StdoutColorDisabled  bool           `json:"stdoutColorDisabled"`  // 是否（默认为false）在向writer输出时，以颜色前缀形式记录日志级别
	WriterColorEnable    bool           `json:"writerColorEnable"`    // 是否（默认为false）在向writer输出时，以颜色前缀形式记录日志级别
	internalConfig
}

type internalConfig struct {
	rotatedHandlerInitialized *gtype.Bool // 是否已初始化旋转功能
}

// DefaultConfig 返回日志器的默认配置。
func DefaultConfig() Config {
	c := Config{
		File:                defaultFileFormat,
		Flags:               F_TIME_STD,
		TimeFormat:          "",
		Level:               LEVEL_ALL,
		CtxKeys:             []interface{}{},
		StStatus:            1,
		HeaderPrint:         true,
		StdoutPrint:         true,
		LevelPrint:          true,
		LevelPrefixes:       make(map[int]string, len(defaultLevelPrefixes)),
		RotateCheckInterval: time.Hour,
		internalConfig: internalConfig{
			rotatedHandlerInitialized: gtype.NewBool(),
		},
	}
	for k, v := range defaultLevelPrefixes {
		c.LevelPrefixes[k] = v
	}
	if !defaultDebug {
		c.Level = c.Level & ^LEVEL_DEBU
	}
	return c
}

// GetConfig 返回当前 Logger 的配置。
func (l *Logger) GetConfig() Config {
	return l.config
}

// SetConfig 为日志器设置配置。
func (l *Logger) SetConfig(config Config) error {
	l.config = config
	// 必要的验证
	if config.Path != "" {
		if err := l.SetPath(config.Path); err != nil {
			intlog.Errorf(context.TODO(), `%+v`, err)
			return err
		}
	}
	intlog.Printf(context.TODO(), "SetConfig: %+v", l.config)
	return nil
}

// SetConfigWithMap 通过map设置日志器的配置。
func (l *Logger) SetConfigWithMap(m map[string]interface{}) error {
	if len(m) == 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "configuration cannot be empty")
	}
// 现在的m是m的一个浅拷贝。
// 这有点巧妙，不是吗？
	m = gutil.MapCopy(m)
	// 将字符串配置转换为级别对应的整数值。
	levelKey, levelValue := gutil.MapPossibleItemByKey(m, "Level")
	if levelValue != nil {
		if level, ok := levelStringMap[strings.ToUpper(gconv.String(levelValue))]; ok {
			m[levelKey] = level
		} else {
			return gerror.NewCodef(gcode.CodeInvalidParameter, `invalid level string: %v`, levelValue)
		}
	}
	// 将字符串配置转换为文件旋转大小的整数值。
	rotateSizeKey, rotateSizeValue := gutil.MapPossibleItemByKey(m, "RotateSize")
	if rotateSizeValue != nil {
		m[rotateSizeKey] = gfile.StrToSize(gconv.String(rotateSizeValue))
		if m[rotateSizeKey] == -1 {
			return gerror.NewCodef(gcode.CodeInvalidConfiguration, `invalid rotate size: %v`, rotateSizeValue)
		}
	}
	if err := gconv.Struct(m, &l.config); err != nil {
		return err
	}
	return l.SetConfig(l.config)
}

// SetDebug 用于开启或关闭日志器的调试级别。
// 默认情况下，调试级别是启用的。
func (l *Logger) SetDebug(debug bool) {
	if debug {
		l.config.Level = l.config.Level | LEVEL_DEBU
	} else {
		l.config.Level = l.config.Level & ^LEVEL_DEBU
	}
}

// SetAsync 启用/禁用异步日志输出功能。
func (l *Logger) SetAsync(enabled bool) {
	if enabled {
		l.config.Flags = l.config.Flags | F_ASYNC
	} else {
		l.config.Flags = l.config.Flags & ^F_ASYNC
	}
}

// SetFlags 设置日志输出功能的额外标志。
func (l *Logger) SetFlags(flags int) {
	l.config.Flags = flags
}

// GetFlags 返回日志器的标志。
func (l *Logger) GetFlags() int {
	return l.config.Flags
}

// SetStack 启用/禁用失败日志输出中的堆栈跟踪功能。
func (l *Logger) SetStack(enabled bool) {
	if enabled {
		l.config.StStatus = 1
	} else {
		l.config.StStatus = 0
	}
}

// SetStackSkip 设置从终点开始的堆栈偏移量。
func (l *Logger) SetStackSkip(skip int) {
	l.config.StSkip = skip
}

// SetStackFilter 从终点设置堆栈过滤器。
func (l *Logger) SetStackFilter(filter string) {
	l.config.StFilter = filter
}

// SetCtxKeys 设置日志器的上下文键。这些键用于从上下文中检索值并将其打印到日志内容中。
//
// 注意，多次调用此函数将覆盖之前设置的上下文键。
func (l *Logger) SetCtxKeys(keys ...interface{}) {
	l.config.CtxKeys = keys
}

// AppendCtxKeys 向日志器追加额外键。
// 如果该键之前已向日志器追加过，则忽略此次操作。
func (l *Logger) AppendCtxKeys(keys ...interface{}) {
	var isExist bool
	for _, key := range keys {
		isExist = false
		for _, ctxKey := range l.config.CtxKeys {
			if ctxKey == key {
				isExist = true
				break
			}
		}
		if !isExist {
			l.config.CtxKeys = append(l.config.CtxKeys, key)
		}
	}
}

// GetCtxKeys 获取并返回用于日志记录的上下文键。
func (l *Logger) GetCtxKeys() []interface{} {
	return l.config.CtxKeys
}

// SetWriter 设置自定义的日志 `writer` 用于日志记录。
// `writer` 对象应实现 io.Writer 接口。
// 开发者可以使用自定义的日志 `writer` 将日志输出重定向到其他服务，
// 例如：kafka、mysql、mongodb 等。
func (l *Logger) SetWriter(writer io.Writer) {
	l.config.Writer = writer
}

// GetWriter 返回自定义的writer对象，该对象实现了io.Writer接口。
// 如果之前未设置过writer，则返回nil。
func (l *Logger) GetWriter() io.Writer {
	return l.config.Writer
}

// SetPath 设置文件日志的目录路径。
func (l *Logger) SetPath(path string) error {
	if path == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "logging path is empty")
	}
	if !gfile.Exists(path) {
		if err := gfile.Mkdir(path); err != nil {
			return gerror.Wrapf(err, `Mkdir "%s" failed in PWD "%s"`, path, gfile.Pwd())
		}
	}
	l.config.Path = strings.TrimRight(path, gfile.Separator)
	return nil
}

// GetPath 返回用于文件日志记录的日志目录路径。
// 如果未设置目录路径，则返回空字符串。
func (l *Logger) GetPath() string {
	return l.config.Path
}

// SetFile 设置文件日志的文件名`pattern`。
// 在`pattern`中可以使用日期时间模式，例如：access-{Ymd}.log。
// 默认的文件名模式是：Y-m-d.log，例如：2018-01-01.log
// 这段Go语言代码注释翻译成中文为：
// 设置文件日志的文件名为 `pattern`。
// 可以在 `pattern` 中使用日期时间格式化字符串，例如：access-{Ymd}.log（表示按年月日生成不同文件）。
// 默认的文件名格式是：Y-m-d.log，例如：2018-01-01.log
func (l *Logger) SetFile(pattern string) {
	l.config.File = pattern
}

// SetTimeFormat 设置日志时间的时间格式。
func (l *Logger) SetTimeFormat(timeFormat string) {
	l.config.TimeFormat = timeFormat
}

// SetStdoutPrint 设置是否将日志内容输出到标准输出(stdout)，默认为true。
func (l *Logger) SetStdoutPrint(enabled bool) {
	l.config.StdoutPrint = enabled
}

// SetHeaderPrint 设置是否输出日志内容的头部，默认为true。
func (l *Logger) SetHeaderPrint(enabled bool) {
	l.config.HeaderPrint = enabled
}

// SetLevelPrint 设置是否输出日志内容的级别字符串，默认为true。
func (l *Logger) SetLevelPrint(enabled bool) {
	l.config.LevelPrint = enabled
}

// SetPrefix 设置每个日志内容的前缀字符串。
// 前缀是头部的一部分，这意味着如果关闭了头部输出，则不会输出任何前缀。
func (l *Logger) SetPrefix(prefix string) {
	l.config.Prefix = prefix
}

// SetHandlers 设置当前日志器的处理程序。
func (l *Logger) SetHandlers(handlers ...Handler) {
	l.config.Handlers = handlers
}

// SetWriterColorEnable 开启文件/写入器日志的彩色输出功能。
func (l *Logger) SetWriterColorEnable(enabled bool) {
	l.config.WriterColorEnable = enabled
}

// SetStdoutColorDisabled 禁用 stdout 日志颜色输出。
func (l *Logger) SetStdoutColorDisabled(disabled bool) {
	l.config.StdoutColorDisabled = disabled
}
