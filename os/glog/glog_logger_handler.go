// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 日志类

import (
	"bytes"
	"context"
	"time"
	
	"github.com/888go/goframe/util/gconv"
)

// Handler 是用于自定义日志内容输出的函数处理器。
type Handler func(ctx context.Context, in *HandlerInput)

// HandlerInput 是 logging Handler 的输入参数结构体。
type HandlerInput struct {
	internalHandlerInfo
	Logger      *Logger       // 当前日志器对象。
	X缓冲区      *bytes.Buffer // Buffer，用于日志内容输出。
	X时间        time.Time     // 日志时间，即触发日志记录的时间。
	X格式化时间  string        // 格式化的时间字符串，如 "2016-01-09 12:00:00"。
	X颜色       int           // 使用颜色，如COLOR_RED、COLOR_BLUE等。例如：34
	X级别       int           // 使用级别，如 LEVEL_INFO, LEVEL_ERRO 等。例如：256
	X文本级别 string        // 格式化的级别字符串，如 "DEBU", "ERRO" 等。例如：ERRO
	X源文件函数名  string        // 如果设置了F_CALLER_FN，该变量记录调用日志函数的源函数名。
	X源文件路径与行号  string        // 调用日志记录的源文件路径及其行号，仅在设置了 F_FILE_SHORT 或 F_FILE_LONG 时可用。
	X上下文值      string        // 从context中获取的字符串类型的上下文值，但只有在配置了Config.CtxKeys时才可用。
	X链路跟踪ID     string        // 跟踪ID，仅在启用OpenTelemetry时可用。
	X前缀      string        // 自定义日志内容前缀字符串。
	X日志内容     string        // Content 是由 logger 生成的、不包含错误堆栈信息的主要日志内容。
	X未格式化数组      []any         // 传递给 logger 的未格式化的值数组。
	Stack       string        // Stack 字符串由 logger 生成，仅在配置了 Config.StStatus 时可用。
	X是否为异步     bool          // IsAsync 标记它处于异步日志记录状态。
}

type internalHandlerInfo struct {
	index    int       // 此中间件用于内部使用，处理索引功能。
	handlers []Handler // 通过索引调用处理器数组
}

// defaultHandler 是该包的默认处理器。
var defaultHandler Handler

// doFinalPrint 是一个用于记录内容打印的处理器。
// 如果其中任意一项被配置，此处理器会将日志内容输出到文件、标准输出(stdout)或写入指定位置。
func doFinalPrint(ctx context.Context, in *HandlerInput) {
	buffer := in.Logger.doFinalPrint(ctx, in)
	if in.X缓冲区.Len() == 0 {
		in.X缓冲区 = buffer
	}
}

// SetDefaultHandler 设置包的默认处理器。
func X设置默认中间件(处理函数 Handler) {
	defaultHandler = 处理函数
}

// GetDefaultHandler 返回该包的默认处理器。
func X取默认中间件() Handler {
	return defaultHandler
}

// Next 以中间件方式调用下一个日志处理程序。
func (in *HandlerInput) Next(ctx context.Context) {
	in.index++
	if in.index < len(in.handlers) {
		in.handlers[in.index](ctx, in)
	}
}

// String 返回由默认日志处理程序格式化的日志内容。
func (in *HandlerInput) String(withColor ...bool) string {
	formatWithColor := false
	if len(withColor) > 0 {
		formatWithColor = withColor[0]
	}
	return in.getDefaultBuffer(formatWithColor).String()
}

func (in *HandlerInput) getDefaultBuffer(withColor bool) *bytes.Buffer {
	buffer := bytes.NewBuffer(nil)
	if in.Logger.config.X是否输出头信息 {
		if in.X格式化时间 != "" {
			buffer.WriteString(in.X格式化时间)
		}
		if in.Logger.config.X是否输出级别 && in.X文本级别 != "" {
			var levelStr = "[" + in.X文本级别 + "]"
			if withColor {
				in.addStringToBuffer(buffer, in.Logger.getColoredStr(
					in.Logger.getColorByLevel(in.X级别), levelStr,
				))
			} else {
				in.addStringToBuffer(buffer, levelStr)
			}
		}
	}
	if in.X链路跟踪ID != "" {
		in.addStringToBuffer(buffer, "{"+in.X链路跟踪ID+"}")
	}
	if in.X上下文值 != "" {
		in.addStringToBuffer(buffer, "{"+in.X上下文值+"}")
	}
	if in.Logger.config.X是否输出头信息 {
		if in.X前缀 != "" {
			in.addStringToBuffer(buffer, in.X前缀)
		}
		if in.X源文件函数名 != "" {
			in.addStringToBuffer(buffer, in.X源文件函数名)
		}
		if in.X源文件路径与行号 != "" {
			in.addStringToBuffer(buffer, in.X源文件路径与行号)
		}
	}

	if in.X日志内容 != "" {
		in.addStringToBuffer(buffer, in.X日志内容)
	}

	// 将values字符串内容进行转换
	var valueContent string
	for _, v := range in.X未格式化数组 {
		valueContent = 转换类.String(v)
		if len(valueContent) == 0 {
			continue
		}
		if buffer.Len() > 0 {
			if buffer.Bytes()[buffer.Len()-1] == '\n' {
				// 删除一个空行（\n\n）
				if valueContent[0] == '\n' {
					valueContent = valueContent[1:]
				}
				buffer.WriteString(valueContent)
			} else {
				buffer.WriteString(" " + valueContent)
			}
		} else {
			buffer.WriteString(valueContent)
		}
	}

	if in.Stack != "" {
		in.addStringToBuffer(buffer, "\nStack:\n"+in.Stack)
	}
	// 避免在行尾出现单个空格。
	buffer.WriteString("\n")
	return buffer
}

func (in *HandlerInput) getRealBuffer(withColor bool) *bytes.Buffer {
	if in.X缓冲区.Len() > 0 {
		return in.X缓冲区
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
