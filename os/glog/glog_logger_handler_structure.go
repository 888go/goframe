// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 日志类

import (
	"bytes"
	"context"
	"strconv"
	"unicode"
	"unicode/utf8"
	
	"github.com/888go/goframe/util/gconv"
)

type structuredBuffer struct {
	in     *HandlerInput
	buffer *bytes.Buffer
}

const (
	structureKeyTime       = "Time"
	structureKeyLevel      = "Level"
	structureKeyPrefix     = "Prefix"
	structureKeyContent    = "Content"
	structureKeyTraceId    = "TraceId"
	structureKeyCallerFunc = "CallerFunc"
	structureKeyCallerPath = "CallerPath"
	structureKeyCtxStr     = "CtxStr"
	structureKeyStack      = "Stack"
)

// 从encoding/json/tables.go复制而来。
//
// safeSet 记录了如果ASCII字符在给定数组位置的值能在JSON字符串中无须额外转义即可表示，则其值为真（true）。
//
// 除了ASCII控制字符（0-31）、双引号（"）和反斜杠字符（\）之外，其余所有字符的值都为真。
var safeSet = [utf8.RuneSelf]bool{
	' ':      true,
	'!':      true,
	'"':      false,
	'#':      true,
	'$':      true,
	'%':      true,
	'&':      true,
	'\'':     true,
	'(':      true,
	')':      true,
	'*':      true,
	'+':      true,
	',':      true,
	'-':      true,
	'.':      true,
	'/':      true,
	'0':      true,
	'1':      true,
	'2':      true,
	'3':      true,
	'4':      true,
	'5':      true,
	'6':      true,
	'7':      true,
	'8':      true,
	'9':      true,
	':':      true,
	';':      true,
	'<':      true,
	'=':      true,
	'>':      true,
	'?':      true,
	'@':      true,
	'A':      true,
	'B':      true,
	'C':      true,
	'D':      true,
	'E':      true,
	'F':      true,
	'G':      true,
	'H':      true,
	'I':      true,
	'J':      true,
	'K':      true,
	'L':      true,
	'M':      true,
	'N':      true,
	'O':      true,
	'P':      true,
	'Q':      true,
	'R':      true,
	'S':      true,
	'T':      true,
	'U':      true,
	'V':      true,
	'W':      true,
	'X':      true,
	'Y':      true,
	'Z':      true,
	'[':      true,
	'\\':     false,
	']':      true,
	'^':      true,
	'_':      true,
	'`':      true,
	'a':      true,
	'b':      true,
	'c':      true,
	'd':      true,
	'e':      true,
	'f':      true,
	'g':      true,
	'h':      true,
	'i':      true,
	'j':      true,
	'k':      true,
	'l':      true,
	'm':      true,
	'n':      true,
	'o':      true,
	'p':      true,
	'q':      true,
	'r':      true,
	's':      true,
	't':      true,
	'u':      true,
	'v':      true,
	'w':      true,
	'x':      true,
	'y':      true,
	'z':      true,
	'{':      true,
	'|':      true,
	'}':      true,
	'~':      true,
	'\u007f': true,
}

// HandlerStructure 是一个处理器，用于将输出的日志内容以结构化字符串形式记录。
func X中间件函数文本结构化输出(上下文 context.Context, in *HandlerInput) {
	s := newStructuredBuffer(in)
	in.X缓冲区.Write(s.X取字节集())
	in.X缓冲区.Write([]byte("\n"))
	in.Next(上下文)
}

func newStructuredBuffer(in *HandlerInput) *structuredBuffer {
	return &structuredBuffer{
		in:     in,
		buffer: bytes.NewBuffer(nil),
	}
}

func (buf *structuredBuffer) X取字节集() []byte {
	buf.addValue(structureKeyTime, buf.in.X格式化时间)
	if buf.in.X链路跟踪ID != "" {
		buf.addValue(structureKeyTraceId, buf.in.X链路跟踪ID)
	}
	if buf.in.X上下文值 != "" {
		buf.addValue(structureKeyCtxStr, buf.in.X上下文值)
	}
	if buf.in.X文本级别 != "" {
		buf.addValue(structureKeyLevel, buf.in.X文本级别)
	}
	if buf.in.X源文件路径与行号 != "" {
		buf.addValue(structureKeyCallerPath, buf.in.X源文件路径与行号)
	}
	if buf.in.X源文件函数名 != "" {
		buf.addValue(structureKeyCallerFunc, buf.in.X源文件函数名)
	}
	if buf.in.X前缀 != "" {
		buf.addValue(structureKeyPrefix, buf.in.X前缀)
	}
	// 如果这些值不能构成一对，则将第一个值移动到content中。
	values := buf.in.X未格式化数组
	if len(values)%2 != 0 {
		if buf.in.X日志内容 != "" {
			buf.in.X日志内容 += " "
		}
		buf.in.X日志内容 += 转换类.String(values[0])
		values = values[1:]
	}
	if buf.in.X日志内容 != "" {
		buf.addValue(structureKeyContent, buf.in.X日志内容)
	}
	// Values pairs.
	for i := 0; i < len(values); i += 2 {
		buf.addValue(values[i], values[i+1])
	}
	if buf.in.Stack != "" {
		buf.addValue(structureKeyStack, buf.in.Stack)
	}
	contentBytes := buf.buffer.Bytes()
	buf.buffer.Reset()
	contentBytes = bytes.ReplaceAll(contentBytes, []byte{'\n'}, []byte{' '})
	return contentBytes
}

func (buf *structuredBuffer) addValue(k, v any) {
	var (
		ks = 转换类.String(k)
		vs = 转换类.String(v)
	)
	if buf.buffer.Len() > 0 {
		buf.buffer.WriteByte(' ')
	}
	buf.appendString(ks)
	buf.buffer.WriteByte('=')
	buf.appendString(vs)
}

func (buf *structuredBuffer) appendString(s string) {
	if buf.needsQuoting(s) {
		s = strconv.Quote(s)
	}
	buf.buffer.WriteString(s)
}

func (buf *structuredBuffer) needsQuoting(s string) bool {
	if len(s) == 0 {
		return true
	}
	for i := 0; i < len(s); {
		b := s[i]
		if b < utf8.RuneSelf {
// 将JSON字符串中需要转义的除反斜杠以外的任何字符，以及空格和'='进行引用
			if b != '\\' && (b == ' ' || b == '=' || !safeSet[b]) {
				return true
			}
			i++
			continue
		}
		r, size := utf8.DecodeRuneInString(s[i:])
		if r == utf8.RuneError || unicode.IsSpace(r) || !unicode.IsPrint(r) {
			return true
		}
		i += size
	}
	return false
}
