// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gstr

import (
	"bytes"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	
	"github.com/gogf/gf/v2/util/grand"
)

var (
	// octReg 是用于检查八进制字符串的正则表达式对象。
	octReg = regexp.MustCompile(`\\[0-7]{3}`)
)

// Chr 返回一个数字（0-255）对应的ASCII字符串。
func Chr(ascii int) string {
	return string([]byte{byte(ascii % 256)})
}

// Ord将字符串的第一个字节转换为0到255之间的值。
func Ord(char string) int {
	return int(char[0])
}

// OctStr converts string container octal string to its original string,
// for example, to Chinese string.
// Eg: `\346\200\241` -> 怡
func OctStr(str string) string {
	return octReg.ReplaceAllStringFunc(
		str,
		func(s string) string {
			i, _ := strconv.ParseInt(s[1:], 8, 0)
			return string([]byte{byte(i)})
		},
	)
}

// Reverse 函数返回一个字符串，该字符串是 `str` 的反转字符串。
func Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// NumberFormat 对数字进行格式化，添加千位分隔符。
// `decimals`: 设置小数点后的位数。
// `decPoint`: 设置小数点的分隔符。
// `thousandsSep`: 设置千位之间的分隔符。
// 参考：http://php.net/manual/en/function.number-format.php
// 这段注释是Go语言代码中对NumberFormat函数功能和参数的中文解释。
func NumberFormat(number float64, decimals int, decPoint, thousandsSep string) string {
	neg := false
	if number < 0 {
		number = -number
		neg = true
	}
	// Will round off
	str := fmt.Sprintf("%."+strconv.Itoa(decimals)+"F", number)
	prefix, suffix := "", ""
	if decimals > 0 {
		prefix = str[:len(str)-(decimals+1)]
		suffix = str[len(str)-decimals:]
	} else {
		prefix = str
	}
	sep := []byte(thousandsSep)
	n, l1, l2 := 0, len(prefix), len(sep)
	// thousands sep num
	c := (l1 - 1) / 3
	tmp := make([]byte, l2*c+l1)
	pos := len(tmp) - 1
	for i := l1 - 1; i >= 0; i, n, pos = i-1, n+1, pos-1 {
		if l2 > 0 && n > 0 && n%3 == 0 {
			for j := range sep {
				tmp[pos] = sep[l2-j-1]
				pos--
			}
		}
		tmp[pos] = prefix[i]
	}
	s := string(tmp)
	if decimals > 0 {
		s += decPoint + suffix
	}
	if neg {
		s = "-" + s
	}

	return s
}

// Shuffle 随机打乱一个字符串。
// 它将参数 `str` 视为unicode字符串。
func Shuffle(str string) string {
	runes := []rune(str)
	s := make([]rune, len(runes))
	for i, v := range grand.Perm(len(runes)) {
		s[i] = runes[v]
	}
	return string(s)
}

// HideStr 将字符串 `str` 中间部分按 `percentage` 百分比替换为 `hide`。
// 此函数视参数 `str` 为 unicode 字符串。
func HideStr(str string, percent int, hide string) string {
	array := strings.Split(str, "@")
	if len(array) > 1 {
		str = array[0]
	}
	var (
		rs       = []rune(str)
		length   = len(rs)
		mid      = math.Floor(float64(length / 2))
		hideLen  = int(math.Floor(float64(length) * (float64(percent) / 100)))
		start    = int(mid - math.Floor(float64(hideLen)/2))
		hideStr  = []rune("")
		hideRune = []rune(hide)
	)
	for i := 0; i < hideLen; i++ {
		hideStr = append(hideStr, hideRune...)
	}
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString(string(rs[0:start]))
	buffer.WriteString(string(hideStr))
	buffer.WriteString(string(rs[start+hideLen:]))
	if len(array) > 1 {
		buffer.WriteString("@" + array[1])
	}
	return buffer.String()
}

// Nl2Br在字符串中的所有换行符前插入HTML换行标签(`br`|<br />)：
// 包括\n\r、\r\n、\r、\n。
// 此函数将参数`str`视为Unicode字符串处理。
func Nl2Br(str string, isXhtml ...bool) string {
	r, n, runes := '\r', '\n', []rune(str)
	var br []byte
	if len(isXhtml) > 0 && isXhtml[0] {
		br = []byte("<br />")
	} else {
		br = []byte("<br>")
	}
	skip := false
	length := len(runes)
	var buf bytes.Buffer
	for i, v := range runes {
		if skip {
			skip = false
			continue
		}
		switch v {
		case n, r:
			if (i+1 < length) && ((v == r && runes[i+1] == n) || (v == n && runes[i+1] == r)) {
				buf.Write(br)
				skip = true
				continue
			}
			buf.Write(br)
		default:
			buf.WriteRune(v)
		}
	}
	return buf.String()
}

// WordWrap 将字符串按照给定的字符数进行换行处理。
// 该函数支持对英文和中文标点符号进行截断处理。
// TODO: 实现自定义截断参数功能，参考 http://php.net/manual/en/function.wordwrap.php.
func WordWrap(str string, width int, br string) string {
	if br == "" {
		br = "\n"
	}
	var (
		current           int
		wordBuf, spaceBuf bytes.Buffer
		init              = make([]byte, 0, len(str))
		buf               = bytes.NewBuffer(init)
		strRunes          = []rune(str)
	)
	for _, char := range strRunes {
		switch {
		case char == '\n':
			if wordBuf.Len() == 0 {
				if current+spaceBuf.Len() > width {
					current = 0
				} else {
					current += spaceBuf.Len()
					_, _ = spaceBuf.WriteTo(buf)
				}
				spaceBuf.Reset()
			} else {
				current += spaceBuf.Len() + wordBuf.Len()
				_, _ = spaceBuf.WriteTo(buf)
				spaceBuf.Reset()
				_, _ = wordBuf.WriteTo(buf)
				wordBuf.Reset()
			}
			buf.WriteRune(char)
			current = 0

		case unicode.IsSpace(char):
			if spaceBuf.Len() == 0 || wordBuf.Len() > 0 {
				current += spaceBuf.Len() + wordBuf.Len()
				_, _ = spaceBuf.WriteTo(buf)
				spaceBuf.Reset()
				_, _ = wordBuf.WriteTo(buf)
				wordBuf.Reset()
			}
			spaceBuf.WriteRune(char)

		case isPunctuation(char):
			wordBuf.WriteRune(char)
			if spaceBuf.Len() == 0 || wordBuf.Len() > 0 {
				current += spaceBuf.Len() + wordBuf.Len()
				_, _ = spaceBuf.WriteTo(buf)
				spaceBuf.Reset()
				_, _ = wordBuf.WriteTo(buf)
				wordBuf.Reset()
			}

		default:
			wordBuf.WriteRune(char)
			if current+spaceBuf.Len()+wordBuf.Len() > width && wordBuf.Len() < width {
				buf.WriteString(br)
				current = 0
				spaceBuf.Reset()
			}
		}
	}

	if wordBuf.Len() == 0 {
		if current+spaceBuf.Len() <= width {
			_, _ = spaceBuf.WriteTo(buf)
		}
	} else {
		_, _ = spaceBuf.WriteTo(buf)
		_, _ = wordBuf.WriteTo(buf)
	}
	return buf.String()
}

func isPunctuation(char int32) bool {
	switch char {
	// 英文标点符号
	case ';', '.', ',', ':', '~':
		return true
	// 中文标点符号
	case '；', '，', '。', '：', '？', '！', '…', '、':
		return true
	default:
		return false
	}
}
