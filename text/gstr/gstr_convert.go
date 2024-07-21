// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

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
	// octReg 是用于检查八进制字符串的正则表达式对象。 md5:5c5c93db5da71e18
	octReg = regexp.MustCompile(`\\[0-7]{3}`)
)

// Chr return the ascii string of a number(0-255).
//
// Chr(65) -> "A"
// ff:整数到ascii
// ascii:整数
func Chr(ascii int) string {
	return string([]byte{byte(ascii % 256)})
}

// Ord converts the first byte of a string to a value between 0 and 255.
//
// Chr("A") -> 65
// ff:
// char:
func Ord(char string) int {
	return int(char[0])
}

// OctStr converts string container octal string to its original string,
// for example, to Chinese string.
//
// OctStr("\346\200\241") -> 怡
// ff:八进制到文本
// str:文本
func OctStr(str string) string {
	return octReg.ReplaceAllStringFunc(
		str,
		func(s string) string {
			i, _ := strconv.ParseInt(s[1:], 8, 0)
			return string([]byte{byte(i)})
		},
	)
}

// Reverse returns a string which is the reverse of `str`.
//
// Reverse("123456") -> "654321"
// ff:反转字符
// str:文本
func Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// NumberFormat formats a number with grouped thousands.
// Parameter `decimals`: Sets the number of decimal points.
// Parameter `decPoint`: Sets the separator for the decimal point.
// Parameter `thousandsSep`: Sets the thousands' separator.
// See http://php.net/manual/en/function.number-format.php.
//
// NumberFormat(1234.56, 2, ".", "")  -> 1234,56
// NumberFormat(1234.56, 2, ",", " ") -> 1 234,56
// ff:格式化数值
// number:数值
// decimals:小数点个数
// decPoint:小数点分隔符
// thousandsSep:千位分隔符
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

// Shuffle randomly shuffles a string.
// It considers parameter `str` as unicode string.
//
// Shuffle("123456") -> "325164"
// Shuffle("123456") -> "231546"
// ...
// ff:随机打散字符
// str:文本
func Shuffle(str string) string {
	runes := []rune(str)
	s := make([]rune, len(runes))
	for i, v := range grand.Perm(len(runes)) {
		s[i] = runes[v]
	}
	return string(s)
}

// HideStr 函数将字符串 `str` 的从中间开始按 `percentage` 比例部分内容替换为 `hide`。
// 此函数将参数 `str` 视为Unicode字符串处理。
// md5:f9986962939bb788
// ff:替换中间字符
// str:文本
// percent:替换百分比
// hide:替换符
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

// Nl2Br 在字符串中的所有换行符(\n\r, \r\n, \r, \n)前插入HTML换行标签(`br`|<br />)。
// 它将参数`str`视为Unicode字符串。
// md5:6cad5f70848065d0
// ff:替换换行符
// str:文本
// isXhtml:是否html
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

// WordWrap wraps a string to a given number of characters.
// This function supports cut parameters of both english and chinese punctuations.
// ff:按字符数量换行
// str:文本
// width:字符数
// br:换行符
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
	// English Punctuations.
	case ';', '.', ',', ':', '~':
		return true
	// Chinese Punctuations.
	case '；', '，', '。', '：', '？', '！', '…', '、':
		return true
	default:
		return false
	}
}
