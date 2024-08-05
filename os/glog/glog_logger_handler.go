// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package glog

import (
	"bytes"
	"context"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
)

// Handler 是用于自定义日志内容输出的函数处理器。 md5:486a8db7f7dd8188
type Handler func(ctx context.Context, in *HandlerInput)

// HandlerInput是日志处理器的输入参数结构体。
// 
// 日志内容由以下部分组成：
// 时间格式 [级别格式] {跟踪ID} {上下文字符串} 前缀 调用函数 调用路径 内容 值 堆栈
// 
// 日志内容的头部是：
// 时间格式 [级别格式] {跟踪ID} {上下文字符串} 前缀 调用函数 调用路径
// md5:6213dd0ebb4e9188
type HandlerInput struct {
	internalHandlerInfo

	// Current Logger object.
	Logger *Logger

		// 用于存储日志输出内容的缓冲区。 md5:33224816c1505400
	Buffer *bytes.Buffer

		// （只读）记录时间，即触发日志记录的时间。 md5:5ce6aaa482dcea28
	Time time.Time

		// 格式化的输出时间字符串，如 "2016-01-09 12:00:00"。 md5:530cb544b3906631
	TimeFormat string

	//（只读）使用颜色常量值，如COLOR_RED，COLOR_BLUE等。
	// 示例：34
	// md5:e377684c0eb82b75
	Color int

	// （只读）使用级别，如LEVEL_INFO、LEVEL_ERROR等。
	// 例子：256
	// md5:20e8f648c34222c9
	Level int

	// 用于输出的格式化日志级别字符串，如 "DEBU"、"ERRO" 等。
	// 示例：ERRO
	// md5:0bad424894695e93
	LevelFormat string

		// 调用日志的源函数名称，仅在设置F_CALLER_FN时可用。 md5:2bfd8148853e8e4c
	CallerFunc string

	// 调用日志的源文件路径及其行号，只有在设置F_FILE_SHORT或F_FILE_LONG时可用。
	// md5:8e31a0cc592be662
	CallerPath string

	// 从上下文中获取的已配置的 context 值字符串。如果没有配置 Config.CtxKeys，它将为空。
	// md5:b854bd1bcad06fda
	CtxStr string

		// 跟踪ID，仅在启用OpenTelemetry时可用，否则为空字符串。 md5:0cd8e77f80286121
	TraceId string

	// 在日志内容头部自定义的前缀字符串。
	// 请注意，如果已禁用HeaderPrint，此设置将不会生效。
	// md5:004eed7afe3ca2dd
	Prefix string

		// 用于日志记录的自定义日志内容。 md5:9749c3bafd8e33d5
	Content string

		// 传递给日志记录器的未经格式化的值数组。 md5:854ab8e84e01371d
	Values []any

	// 由记录器生成的堆栈字符串，仅在配置了Config.StStatus时可用。
	// 注意，堆栈内容中通常包含多行。
	// md5:c36e69fdfae3ac16
	Stack string

		// IsAsync 标记为异步日志记录。 md5:e138a9a968506347
	IsAsync bool
}

type internalHandlerInfo struct {
	index    int       // 处理内部使用的索引的中间件。 md5:61d366e59aee7159
	handlers []Handler // 通过索引调用处理器数组。 md5:7cb772c2e129fd27
}

// defaultHandler 是包的默认处理程序。 md5:0f4cafed00a48af2
var defaultHandler Handler

// doFinalPrint 是用于记录内容打印的处理器。
// 此处理器将日志内容输出到文件/stdout/write，如果它们中有任何被配置的话。
// md5:794b81b9fa0a2bd6
func doFinalPrint(ctx context.Context, in *HandlerInput) {
	buffer := in.Logger.doFinalPrint(ctx, in)
	if in.Buffer.Len() == 0 {
		in.Buffer = buffer
	}
}

// SetDefaultHandler 设置包的默认处理器。 md5:33a213aebe83e5ed
func SetDefaultHandler(handler Handler) {
	defaultHandler = handler
}

// GetDefaultHandler 返回包的默认处理器。 md5:8812c42db1189f3b
func GetDefaultHandler() Handler {
	return defaultHandler
}

// Next 以中间件的方式调用下一个日志处理程序。 md5:ab91f9dfe65c4322
func (in *HandlerInput) Next(ctx context.Context) {
	in.index++
	if in.index < len(in.handlers) {
		in.handlers[in.index](ctx, in)
	}
}

// String 返回默认日志处理器格式化的日志内容。 md5:e78613962fe54276
func (in *HandlerInput) String(withColor ...bool) string {
	formatWithColor := false
	if len(withColor) > 0 {
		formatWithColor = withColor[0]
	}
	return in.getDefaultBuffer(formatWithColor).String()
}

// ValuesContent 将值转换为字符串内容并返回。 md5:da3a0fd9093d35c9
func (in *HandlerInput) ValuesContent() string {
	var (
		buffer       = bytes.NewBuffer(nil)
		valueContent string
	)
	for _, v := range in.Values {
		valueContent = gconv.String(v)
		if len(valueContent) == 0 {
			continue
		}
		if buffer.Len() == 0 {
			buffer.WriteString(valueContent)
			continue
		}
		if buffer.Bytes()[buffer.Len()-1] != '\n' {
			buffer.WriteString(" " + valueContent)
			continue
		}
				// 移除一个空行（\n\n）。 md5:777d73ee86014d2c
		if valueContent[0] == '\n' {
			valueContent = valueContent[1:]
		}
		buffer.WriteString(valueContent)
	}
	return buffer.String()
}

func (in *HandlerInput) getDefaultBuffer(withColor bool) *bytes.Buffer {
	buffer := bytes.NewBuffer(nil)
	if in.Logger.config.HeaderPrint {
		if in.TimeFormat != "" {
			buffer.WriteString(in.TimeFormat)
		}
		if in.Logger.config.LevelPrint && in.LevelFormat != "" {
			var levelStr = "[" + in.LevelFormat + "]"
			if withColor {
				in.addStringToBuffer(buffer, in.Logger.getColoredStr(
					in.Logger.getColorByLevel(in.Level), levelStr,
				))
			} else {
				in.addStringToBuffer(buffer, levelStr)
			}
		}
	}
	if in.TraceId != "" {
		in.addStringToBuffer(buffer, "{"+in.TraceId+"}")
	}
	if in.CtxStr != "" {
		in.addStringToBuffer(buffer, "{"+in.CtxStr+"}")
	}
	if in.Logger.config.HeaderPrint {
		if in.Prefix != "" {
			in.addStringToBuffer(buffer, in.Prefix)
		}
		if in.CallerFunc != "" {
			in.addStringToBuffer(buffer, in.CallerFunc)
		}
		if in.CallerPath != "" {
			in.addStringToBuffer(buffer, in.CallerPath)
		}
	}

	if in.Content != "" {
		in.addStringToBuffer(buffer, in.Content)
	}

	if len(in.Values) > 0 {
		in.addStringToBuffer(buffer, in.ValuesContent())
	}

	if in.Stack != "" {
		in.addStringToBuffer(buffer, "\nStack:\n"+in.Stack)
	}
		// 避免在行尾留下单个空格。 md5:f107ec37b9775773
	buffer.WriteString("\n")
	return buffer
}

func (in *HandlerInput) getRealBuffer(withColor bool) *bytes.Buffer {
	if in.Buffer.Len() > 0 {
		return in.Buffer
	}
	return in.getDefaultBuffer(withColor)
}

func (in *HandlerInput) addStringToBuffer(buffer *bytes.Buffer, strings ...string) {
	for _, s := range strings {
		if buffer.Len() > 0 {
			buffer.WriteByte(' ')
		}
		buffer.WriteString(s)
	}
}
