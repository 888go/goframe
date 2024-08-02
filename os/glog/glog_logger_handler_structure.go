// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 日志类

import (
	"bytes"
	"context"
	"strconv"
	"unicode"
	"unicode/utf8"

	gconv "github.com/888go/goframe/util/gconv"
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

// 从encoding/json/tables.go复制。
//
// safeSet用于存储一个布尔值，表示具有给定数组位置的ASCII字符可以在JSON字符串中表示，而无需进一步转义。
//
// 除了ASCII控制字符（0-31）、双引号（"）和反斜杠字符（\）之外，所有值都为true。
// md5:2df5305c3a107923
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

// HandlerStructure 是一个处理器，用于将输出的日志内容以结构化字符串的形式处理。 md5:392f74b46dcdd1eb
func HandlerStructure(ctx context.Context, in *HandlerInput) {
	s := newStructuredBuffer(in)
	in.Buffer.Write(s.Bytes())
	in.Buffer.Write([]byte("\n"))
	in.Next(ctx)
}

func newStructuredBuffer(in *HandlerInput) *structuredBuffer {
	return &structuredBuffer{
		in:     in,
		buffer: bytes.NewBuffer(nil),
	}
}

func (buf *structuredBuffer) Bytes() []byte {
	buf.addValue(structureKeyTime, buf.in.TimeFormat)
	if buf.in.TraceId != "" {
		buf.addValue(structureKeyTraceId, buf.in.TraceId)
	}
	if buf.in.CtxStr != "" {
		buf.addValue(structureKeyCtxStr, buf.in.CtxStr)
	}
	if buf.in.LevelFormat != "" {
		buf.addValue(structureKeyLevel, buf.in.LevelFormat)
	}
	if buf.in.CallerPath != "" {
		buf.addValue(structureKeyCallerPath, buf.in.CallerPath)
	}
	if buf.in.CallerFunc != "" {
		buf.addValue(structureKeyCallerFunc, buf.in.CallerFunc)
	}
	if buf.in.Prefix != "" {
		buf.addValue(structureKeyPrefix, buf.in.Prefix)
	}
		// 如果这些值不能组成一对，将第一个移到content中。 md5:2bc1ae2ae5605225
	values := buf.in.Values
	if len(values)%2 != 0 {
		if buf.in.Content != "" {
			buf.in.Content += " "
		}
		buf.in.Content += gconv.String(values[0])
		values = values[1:]
	}
	if buf.in.Content != "" {
		buf.addValue(structureKeyContent, buf.in.Content)
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
		ks = gconv.String(k)
		vs = gconv.String(v)
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
			// 在JSON字符串中，除了需要转义的反斜杠、空格和'='之外，对任何内容进行引号包裹
			// md5:0202f0293260c21e
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
