// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 日志类

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
	X中间件             []Handler      `json:"-"`                    // 日志处理器实现类似中间件的功能。
	Writer               io.Writer      `json:"-"`                    // 自定义 io.Writer。
	X日志标识                int            `json:"flags"`                // 附加日志输出功能的标志。
	X时间格式           string         `json:"timeFormat"`           // 日志时间格式
	X文件路径                 string         `json:"path"`                 // 日志目录路径。
	X文件名格式                 string         `json:"file"`                 // 日志文件的格式化模式。
	X级别                int            `json:"level"`                // Output level.
	X前缀               string         `json:"prefix"`               // Prefix 字符串，用于作为每条日志内容的前缀。
	X堆栈偏移量               int            `json:"stSkip"`               // 跳过堆栈的计数。
	X堆栈状态             int            `json:"stStatus"`             // 栈状态(1: 启用 - 默认值; 0: 禁用)
	X堆栈过滤             string         `json:"stFilter"`             // 字符串过滤栈
	X上下文名称              []interface{}  `json:"ctxKeys"`              // Context keys 用于日志记录，它们被用于从context中检索值。
	X是否输出头信息          bool           `json:"header"`               // 是否打印头部信息（默认为true）
	X是否同时输出到终端          bool           `json:"stdout"`               // 是否输出到标准输出（默认为true）
	X是否输出级别           bool           `json:"levelPrint"`           // 是否打印级别格式化字符串（默认为true）
	X日志级别名称映射        map[int]string `json:"levelPrefixes"`        // 日志级别到其前缀字符串的映射。
	X文件分割大小           int64          `json:"rotateSize"`           // 如果日志文件大小 > 0字节，则旋转日志文件。
	X文件分割周期         time.Duration  `json:"rotateExpire"`         // 如果日志文件的修改时间超过这个持续时间，则旋转日志文件。
	X文件分割保留数量    int            `json:"rotateBackupLimit"`    // Max 备份文件数量，默认为0，表示不进行备份。
	X文件分割过期时间   time.Duration  `json:"rotateBackupExpire"`   // Max 为轮转文件设置的过期时间，默认为0，表示永不过期。
	X文件压缩级别 int            `json:"rotateBackupCompress"` // Compress level for rotated files using gzip algorithm. It's 0 in default, means no compression.
	X文件分割检查间隔  time.Duration  `json:"rotateCheckInterval"`  // Asynchronously checks the backups and expiration at intervals. It's 1 hour in default.
	X关闭终端颜色输出  bool           `json:"stdoutColorDisabled"`  // 是否（默认为false）在向writer输出时，以颜色前缀形式记录日志级别
	X文件是否输出颜色    bool           `json:"writerColorEnable"`    // 是否（默认为false）在向writer输出时，以颜色前缀形式记录日志级别
	internalConfig
}

type internalConfig struct {
	rotatedHandlerInitialized *安全变量类.Bool // 是否已初始化旋转功能
}

// DefaultConfig 返回日志器的默认配置。
func X生成默认配置() Config {
	c := Config{
		X文件名格式:                defaultFileFormat,
		X日志标识:               F_TIME_STD,
		X时间格式:          "",
		X级别:               LEVEL_ALL,
		X上下文名称:             []interface{}{},
		X堆栈状态:            1,
		X是否输出头信息:         true,
		X是否同时输出到终端:         true,
		X是否输出级别:          true,
		X日志级别名称映射:       make(map[int]string, len(defaultLevelPrefixes)),
		X文件分割检查间隔: time.Hour,
		internalConfig: internalConfig{
			rotatedHandlerInitialized: 安全变量类.NewBool(),
		},
	}
	for k, v := range defaultLevelPrefixes {
		c.X日志级别名称映射[k] = v
	}
	if !defaultDebug {
		c.X级别 = c.X级别 & ^LEVEL_DEBU
	}
	return c
}

// GetConfig 返回当前 Logger 的配置。
func (l *Logger) X取配置项() Config {
	return l.config
}

// SetConfig 为日志器设置配置。
func (l *Logger) X设置配置项(配置项 Config) error {
	l.config = 配置项
	// 必要的验证
	if 配置项.X文件路径 != "" {
		if err := l.X设置文件路径(配置项.X文件路径); err != nil {
			intlog.Errorf(context.TODO(), `%+v`, err)
			return err
		}
	}
	intlog.Printf(context.TODO(), "SetConfig: %+v", l.config)
	return nil
}

// SetConfigWithMap 通过map设置日志器的配置。
func (l *Logger) X设置配置Map(m map[string]interface{}) error {
	if len(m) == 0 {
		return 错误类.X创建错误码(错误码类.CodeInvalidParameter, "configuration cannot be empty")
	}
// 现在的m是m的一个浅拷贝。
// 这有点巧妙，不是吗？
	m = 工具类.MapCopy(m)
	// 将字符串配置转换为级别对应的整数值。
	levelKey, levelValue := 工具类.MapPossibleItemByKey(m, "Level")
	if levelValue != nil {
		if level, ok := levelStringMap[strings.ToUpper(转换类.String(levelValue))]; ok {
			m[levelKey] = level
		} else {
			return 错误类.X创建错误码并格式化(错误码类.CodeInvalidParameter, `invalid level string: %v`, levelValue)
		}
	}
	// 将字符串配置转换为文件旋转大小的整数值。
	rotateSizeKey, rotateSizeValue := 工具类.MapPossibleItemByKey(m, "RotateSize")
	if rotateSizeValue != nil {
		m[rotateSizeKey] = 文件类.X易读格式转字节长度(转换类.String(rotateSizeValue))
		if m[rotateSizeKey] == -1 {
			return 错误类.X创建错误码并格式化(错误码类.CodeInvalidConfiguration, `invalid rotate size: %v`, rotateSizeValue)
		}
	}
	if err := 转换类.Struct(m, &l.config); err != nil {
		return err
	}
	return l.X设置配置项(l.config)
}

// SetDebug 用于开启或关闭日志器的调试级别。
// 默认情况下，调试级别是启用的。
func (l *Logger) X设置debug(开启 bool) {
	if 开启 {
		l.config.X级别 = l.config.X级别 | LEVEL_DEBU
	} else {
		l.config.X级别 = l.config.X级别 & ^LEVEL_DEBU
	}
}

// SetAsync 启用/禁用异步日志输出功能。
func (l *Logger) X设置异步输出(开启 bool) {
	if 开启 {
		l.config.X日志标识 = l.config.X日志标识 | F_ASYNC
	} else {
		l.config.X日志标识 = l.config.X日志标识 & ^F_ASYNC
	}
}

// SetFlags 设置日志输出功能的额外标志。
func (l *Logger) X设置额外标识(标识 int) {
	l.config.X日志标识 = 标识
}

// GetFlags 返回日志器的标志。
func (l *Logger) X取标识() int {
	return l.config.X日志标识
}

// SetStack 启用/禁用失败日志输出中的堆栈跟踪功能。
func (l *Logger) X设置堆栈跟踪(开启 bool) {
	if 开启 {
		l.config.X堆栈状态 = 1
	} else {
		l.config.X堆栈状态 = 0
	}
}

// SetStackSkip 设置从终点开始的堆栈偏移量。
func (l *Logger) X设置堆栈偏移量(偏移量 int) {
	l.config.X堆栈偏移量 = 偏移量
}

// SetStackFilter 从终点设置堆栈过滤器。
func (l *Logger) X设置堆栈过滤(过滤器 string) {
	l.config.X堆栈过滤 = 过滤器
}

// SetCtxKeys 设置日志器的上下文键。这些键用于从上下文中检索值并将其打印到日志内容中。
//
// 注意，多次调用此函数将覆盖之前设置的上下文键。
func (l *Logger) X设置上下文名称(名称 ...interface{}) {
	l.config.X上下文名称 = 名称
}

// AppendCtxKeys 向日志器追加额外键。
// 如果该键之前已向日志器追加过，则忽略此次操作。
func (l *Logger) AppendCtxKeys(keys ...interface{}) {
	var isExist bool
	for _, key := range keys {
		isExist = false
		for _, ctxKey := range l.config.X上下文名称 {
			if ctxKey == key {
				isExist = true
				break
			}
		}
		if !isExist {
			l.config.X上下文名称 = append(l.config.X上下文名称, key)
		}
	}
}

// GetCtxKeys 获取并返回用于日志记录的上下文键。
func (l *Logger) X取上下文名称() []interface{} {
	return l.config.X上下文名称
}

// SetWriter 设置自定义的日志 `writer` 用于日志记录。
// `writer` 对象应实现 io.Writer 接口。
// 开发者可以使用自定义的日志 `writer` 将日志输出重定向到其他服务，
// 例如：kafka、mysql、mongodb 等。
func (l *Logger) X设置Writer(writer io.Writer) {
	l.config.Writer = writer
}

// GetWriter 返回自定义的writer对象，该对象实现了io.Writer接口。
// 如果之前未设置过writer，则返回nil。
func (l *Logger) X取Writer() io.Writer {
	return l.config.Writer
}

// SetPath 设置文件日志的目录路径。
func (l *Logger) X设置文件路径(文件路径 string) error {
	if 文件路径 == "" {
		return 错误类.X创建错误码(错误码类.CodeInvalidParameter, "logging path is empty")
	}
	if !文件类.X是否存在(文件路径) {
		if err := 文件类.X创建目录(文件路径); err != nil {
			return 错误类.X多层错误并格式化(err, `Mkdir "%s" failed in PWD "%s"`, 文件路径, 文件类.X取当前工作目录())
		}
	}
	l.config.X文件路径 = strings.TrimRight(文件路径, 文件类.Separator)
	return nil
}

// GetPath 返回用于文件日志记录的日志目录路径。
// 如果未设置目录路径，则返回空字符串。
func (l *Logger) X取文件路径() string {
	return l.config.X文件路径
}

// SetFile 设置文件日志的文件名`pattern`。
// 在`pattern`中可以使用日期时间模式，例如：access-{Ymd}.log。
// 默认的文件名模式是：Y-m-d.log，例如：2018-01-01.log
// 这段Go语言代码注释翻译成中文为：
// 设置文件日志的文件名为 `pattern`。
// 可以在 `pattern` 中使用日期时间格式化字符串，例如：access-{Ymd}.log（表示按年月日生成不同文件）。
// 默认的文件名格式是：Y-m-d.log，例如：2018-01-01.log
func (l *Logger) X设置文件名格式(文件名格式 string) {
	l.config.X文件名格式 = 文件名格式
}

// SetTimeFormat 设置日志时间的时间格式。
func (l *Logger) X设置时间格式(时间格式 string) {
	l.config.X时间格式 = 时间格式
}

// SetStdoutPrint 设置是否将日志内容输出到标准输出(stdout)，默认为true。
func (l *Logger) X设置是否同时输出到终端(开启 bool) {
	l.config.X是否同时输出到终端 = 开启
}

// SetHeaderPrint 设置是否输出日志内容的头部，默认为true。
func (l *Logger) X设置是否输出头信息(开启 bool) {
	l.config.X是否输出头信息 = 开启
}

// SetLevelPrint 设置是否输出日志内容的级别字符串，默认为true。
func (l *Logger) X设置是否输出级别(开启 bool) {
	l.config.X是否输出级别 = 开启
}

// SetPrefix 设置每个日志内容的前缀字符串。
// 前缀是头部的一部分，这意味着如果关闭了头部输出，则不会输出任何前缀。
func (l *Logger) X设置前缀(前缀 string) {
	l.config.X前缀 = 前缀
}

// SetHandlers 设置当前日志器的处理程序。
func (l *Logger) X设置中间件(处理函数 ...Handler) {
	l.config.X中间件 = 处理函数
}

// SetWriterColorEnable 开启文件/写入器日志的彩色输出功能。
func (l *Logger) X设置文件是否输出颜色(开启 bool) {
	l.config.X文件是否输出颜色 = 开启
}

// SetStdoutColorDisabled 禁用 stdout 日志颜色输出。
func (l *Logger) X设置关闭终端颜色输出(关闭 bool) {
	l.config.X关闭终端颜色输出 = 关闭
}
