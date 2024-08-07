// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 日志类

import (
	"context"
	"io"
	"strings"
	"time"

	gtype "github.com/888go/goframe/container/gtype"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	gfile "github.com/888go/goframe/os/gfile"
	gconv "github.com/888go/goframe/util/gconv"
	gutil "github.com/888go/goframe/util/gutil"
)

// Config 是记录器的配置对象。 md5:df2a8ab047bea305
type Config struct {
	Handlers             []Handler      `json:"-"`                    // Logger handlers 实现了类似于中间件的功能。 md5:dba4d3d0c7f592b9
	Writer               io.Writer      `json:"-"`                    // Customized io.Writer.
	Flags                int            `json:"flags"`                // 用于日志输出功能的额外标志。 md5:6b323bf0cac304e0
	TimeFormat           string         `json:"timeFormat"`           // Logging time format
	Path                 string         `json:"path"`                 // 日志目录路径。 md5:a9b74f93138f8158
	File                 string         `json:"file"`                 // 日志文件的格式化模式。 md5:88ed9324a5afc4c8
	Level                int            `json:"level"`                // Output level.
	Prefix               string         `json:"prefix"`               // 每个日志内容的前缀字符串。 md5:9318d5ac0b1e3e0e
	StSkip               int            `json:"stSkip"`               // 跳过栈的计数。 md5:dd0842336cee717c
	StStatus             int            `json:"stStatus"`             // 栈的状态（1：启用 - 默认；0：禁用）. md5:3a07964ef563d1f6
	StFilter             string         `json:"stFilter"`             // Stack string filter.
	CtxKeys              []interface{}  `json:"ctxKeys"`              // 用于日志记录的上下文键，用于从上下文中检索值。 md5:d4a4f8b7c1027e23
	HeaderPrint          bool           `json:"header"`               // 是否打印头部信息，默认为true。 md5:8b587e739278ffe3
	StdoutPrint          bool           `json:"stdout"`               // 是否将输出写入stdout（默认为true）。 md5:4f790cec19c3aa5a
	LevelPrint           bool           `json:"levelPrint"`           // 是否打印级别格式字符串（默认为 true）。 md5:2261e6b6d1cccb2d
	LevelPrefixes        map[int]string `json:"levelPrefixes"`        // 日志级别与其前缀字符串的映射。 md5:c4a5a8857bc37946
	RotateSize           int64          `json:"rotateSize"`           // 如果日志文件的大小大于0字节，则进行日志文件轮换。 md5:9fb4614dcea49823
	RotateExpire         time.Duration  `json:"rotateExpire"`         // 如果日志文件的修改时间超过这个持续时间，就旋转日志文件。 md5:0832ad6a5113efe9
	RotateBackupLimit    int            `json:"rotateBackupLimit"`    // 旋转文件的最大备份数量，默认为0，表示不备份。 md5:67b4908c6a850b93
	RotateBackupExpire   time.Duration  `json:"rotateBackupExpire"`   // Max 为旋转文件的过期时间，默认为0，表示永不过期。 md5:eedec2e3ee56fc5d
	RotateBackupCompress int            `json:"rotateBackupCompress"` // Compress level for rotated files using gzip algorithm. It's 0 in default, means no compression.
	RotateCheckInterval  time.Duration  `json:"rotateCheckInterval"`  // Asynchronously checks the backups and expiration at intervals. It's 1 hour in default.
	StdoutColorDisabled  bool           `json:"stdoutColorDisabled"`  // 是否向写入器输出带有颜色的日志级别前缀（默认为false）。 md5:cd5684396601fdfc
	WriterColorEnable    bool           `json:"writerColorEnable"`    // 是否向写入器输出带有颜色的日志级别前缀（默认为false）。 md5:cd5684396601fdfc
	internalConfig
}

type internalConfig struct {
	rotatedHandlerInitialized *gtype.Bool // 是否启用了旋转功能。 md5:32f779f6bf9c7aee
}

// X生成默认配置 返回日志记录器的默认配置。 md5:307781636b8ca142
func X生成默认配置() Config {
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

// X取配置项 返回当前Logger的配置。 md5:34aac9175b86a456
func (l *Logger) X取配置项() Config {
	return l.config
}

// X设置配置项 为logger设置配置。 md5:d219673b9a3ec8b0
func (l *Logger) X设置配置项(配置项 Config) error {
	l.config = 配置项
	// Necessary validation.
	if 配置项.Path != "" {
		if err := l.X设置文件路径(配置项.Path); err != nil {
			intlog.Errorf(context.TODO(), `%+v`, err)
			return err
		}
	}
	intlog.Printf(context.TODO(), "SetConfig: %+v", l.config)
	return nil
}

// X设置配置Map 使用映射为日志器设置配置。 md5:a4d4197c666898a3
func (l *Logger) X设置配置Map(m map[string]interface{}) error {
	if len(m) == 0 {
		return gerror.X创建错误码(gcode.CodeInvalidParameter, "configuration cannot be empty")
	}
	// 现在m是一个浅拷贝 of m。
	// 有点巧妙，不是吗？
	// md5:644970336da24c9d
	m = gutil.MapCopy(m)
		// 将字符串配置更改为级别的整数值。 md5:e990c1dc64df6943
	levelKey, levelValue := gutil.MapPossibleItemByKey(m, "Level")
	if levelValue != nil {
		if level, ok := levelStringMap[strings.ToUpper(gconv.String(levelValue))]; ok {
			m[levelKey] = level
		} else {
			return gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, `invalid level string: %v`, levelValue)
		}
	}
		// 将字符串配置转换为文件轮转大小的整数值。 md5:b9efebe5c9999270
	rotateSizeKey, rotateSizeValue := gutil.MapPossibleItemByKey(m, "RotateSize")
	if rotateSizeValue != nil {
		m[rotateSizeKey] = gfile.X易读格式转字节长度(gconv.String(rotateSizeValue))
		if m[rotateSizeKey] == -1 {
			return gerror.X创建错误码并格式化(gcode.CodeInvalidConfiguration, `invalid rotate size: %v`, rotateSizeValue)
		}
	}
	if err := gconv.Struct(m, &l.config); err != nil {
		return err
	}
	return l.X设置配置项(l.config)
}

// X设置debug 用于启用/禁用日志记录器的调试级别。默认情况下，调试级别是启用的。
// md5:72f0f67e25416b8e
func (l *Logger) X设置debug(开启 bool) {
	if 开启 {
		l.config.Level = l.config.Level | LEVEL_DEBU
	} else {
		l.config.Level = l.config.Level & ^LEVEL_DEBU
	}
}

// X设置异步输出 启用/禁用异步日志输出功能。 md5:10096a3a0860346e
func (l *Logger) X设置异步输出(开启 bool) {
	if 开启 {
		l.config.Flags = l.config.Flags | F_ASYNC
	} else {
		l.config.Flags = l.config.Flags & ^F_ASYNC
	}
}

// X设置额外标识 为日志输出功能设置额外的标志。 md5:40253d4ed662de77
func (l *Logger) X设置额外标识(标识 int) {
	l.config.Flags = 标识
}

// X取标识 返回记录器的标志。 md5:b9a17daa74081d07
func (l *Logger) X取标识() int {
	return l.config.Flags
}

// X设置堆栈跟踪 启用/禁用失败日志输出中的堆栈功能。 md5:3c80a664fff650de
func (l *Logger) X设置堆栈跟踪(开启 bool) {
	if 开启 {
		l.config.StStatus = 1
	} else {
		l.config.StStatus = 0
	}
}

// X设置堆栈偏移量 设置从终点开始的堆栈偏移量。 md5:98a83cd0e38dc56c
func (l *Logger) X设置堆栈偏移量(偏移量 int) {
	l.config.StSkip = 偏移量
}

// X设置堆栈过滤 从端点设置堆栈过滤器。 md5:7eabd577c24907f2
func (l *Logger) X设置堆栈过滤(过滤器 string) {
	l.config.StFilter = 过滤器
}

// X设置上下文名称 为日志器设置上下文键。这些键用于从上下文中检索值并将其打印到日志内容中。
// 
// 注意，多次调用此函数会覆盖之前设置的上下文键。
// md5:f7244f6c7fa79db2
func (l *Logger) X设置上下文名称(名称 ...interface{}) {
	l.config.CtxKeys = 名称
}

// AppendCtxKeys 向记录器添加额外的键。
// 如果该键已先前被添加到记录器中，则此操作会忽略该键。
// md5:f989e696d285ffc1
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

// X取上下文名称检索并返回日志记录的上下文键。 md5:1e780871ada8c59c
func (l *Logger) X取上下文名称() []interface{} {
	return l.config.CtxKeys
}

// X设置Writer 设置自定义的日志记录器`writer`。
// `writer`对象应实现io.Writer接口。
// 开发者可以使用自定义的`writer`将日志输出重定向到其他服务，
// 例如：kafka，mysql，mongodb等。
// md5:8f650a69c1fe2b4b
func (l *Logger) X设置Writer(writer io.Writer) {
	l.config.Writer = writer
}

// X取Writer 返回自定义的 writer 对象，该对象实现了 io.Writer 接口。
// 如果之前未设置 writer，则返回 nil。
// md5:cce0a2679c717d75
func (l *Logger) X取Writer() io.Writer {
	return l.config.Writer
}

// X设置文件路径 设置文件日志的目录路径。 md5:817e6d2802241584
func (l *Logger) X设置文件路径(文件路径 string) error {
	if 文件路径 == "" {
		return gerror.X创建错误码(gcode.CodeInvalidParameter, "logging path is empty")
	}
	if !gfile.X是否存在(文件路径) {
		if err := gfile.X创建目录(文件路径); err != nil {
			return gerror.X多层错误并格式化(err, `Mkdir "%s" failed in PWD "%s"`, 文件路径, gfile.X取当前工作目录())
		}
	}
	l.config.Path = strings.TrimRight(文件路径, gfile.Separator)
	return nil
}

// X取文件路径 返回用于文件日志记录的目录路径。
// 如果未设置目录路径，它将返回空字符串。
// md5:f69da996992ffd9e
func (l *Logger) X取文件路径() string {
	return l.config.Path
}

// X设置文件名格式 设置文件日志的文件名模式为 `pattern`。
// 在 `pattern` 中可以使用日期时间模式，例如：access-YYYYMMDD.log。
// 默认的文件名模式为：Y-m-d.log，例如：2018-01-01.log
// md5:03b3a973ce783b24
func (l *Logger) X设置文件名格式(文件名格式 string) {
	l.config.File = 文件名格式
}

// X设置时间格式 设置日志时间的时间格式。 md5:258a98926fba4588
func (l *Logger) X设置时间格式(时间格式 string) {
	l.config.TimeFormat = 时间格式
}

// X设置是否同时输出到终端 设置是否将日志内容输出到标准输出，默认为true。 md5:b212437cebfd423a
func (l *Logger) X设置是否同时输出到终端(开启 bool) {
	l.config.StdoutPrint = 开启
}

// X设置是否输出头信息 设置日志输出的头部是否打印，默认为 true。 md5:3e71cb67564384cc
func (l *Logger) X设置是否输出头信息(开启 bool) {
	l.config.HeaderPrint = 开启
}

// X设置是否输出级别 设置是否输出日志内容的级别字符串，默认为true。 md5:6ba8899e4d3d1c1b
func (l *Logger) X设置是否输出级别(开启 bool) {
	l.config.LevelPrint = 开启
}

// X设置前缀 设置日志内容的前缀字符串。
// 前缀是日志头的一部分，如果关闭了头部输出，就不会显示前缀。
// md5:31d8e3c101c1eea6
func (l *Logger) X设置前缀(前缀 string) {
	l.config.Prefix = 前缀
}

// X设置中间件 设置当前日志记录器的处理程序。 md5:7b876afcd04a669e
func (l *Logger) X设置中间件(处理函数 ...Handler) {
	l.config.Handlers = 处理函数
}

// X设置文件是否输出颜色 启用带有颜色的文件/写入器日志记录。 md5:deef19b9707bd4df
func (l *Logger) X设置文件是否输出颜色(开启 bool) {
	l.config.WriterColorEnable = 开启
}

// X设置关闭终端颜色输出 禁用带有颜色的stdout日志记录。 md5:aed9b0e4a2ba0f72
func (l *Logger) X设置关闭终端颜色输出(关闭 bool) {
	l.config.StdoutColorDisabled = 关闭
}
