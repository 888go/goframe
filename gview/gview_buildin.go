// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gview

import (
	"bytes"
	"context"
	"fmt"
	htmltpl "html/template"
	"strings"
	
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
	"github.com/gogf/gf/v2/util/gutil"
)

// buildInFuncDump 实现了内置模板函数：dump
func (view *View) buildInFuncDump(values ...interface{}) string {
	buffer := bytes.NewBuffer(nil)
	buffer.WriteString("\n")
	buffer.WriteString("<!--\n")
	if gmode.IsDevelop() {
		for _, v := range values {
			gutil.DumpTo(buffer, v, gutil.DumpOption{})
			buffer.WriteString("\n")
		}
	} else {
		buffer.WriteString("dump feature is disabled as process is not running in develop mode\n")
	}
	buffer.WriteString("-->\n")
	return buffer.String()
}

// buildInFuncMap 实现了内置模板函数：map
func (view *View) buildInFuncMap(value ...interface{}) map[string]interface{} {
	if len(value) > 0 {
		return gconv.Map(value[0])
	}
	return map[string]interface{}{}
}

// buildInFuncMaps 实现了内置模板函数：maps
func (view *View) buildInFuncMaps(value ...interface{}) []map[string]interface{} {
	if len(value) > 0 {
		return gconv.Maps(value[0])
	}
	return []map[string]interface{}{}
}

// buildInFuncEq 实现了内建模板函数：eq
func (view *View) buildInFuncEq(value interface{}, others ...interface{}) bool {
	s := gconv.String(value)
	for _, v := range others {
		if strings.Compare(s, gconv.String(v)) == 0 {
			return true
		}
	}
	return false
}

// buildInFuncNe 实现了内建模板函数：ne
func (view *View) buildInFuncNe(value, other interface{}) bool {
	return strings.Compare(gconv.String(value), gconv.String(other)) != 0
}

// buildInFuncLt 实现了内置模板函数：lt
func (view *View) buildInFuncLt(value, other interface{}) bool {
	s1 := gconv.String(value)
	s2 := gconv.String(other)
	if gstr.IsNumeric(s1) && gstr.IsNumeric(s2) {
		return gconv.Int64(value) < gconv.Int64(other)
	}
	return strings.Compare(s1, s2) < 0
}

// buildInFuncLe 实现了内置模板函数：le
func (view *View) buildInFuncLe(value, other interface{}) bool {
	s1 := gconv.String(value)
	s2 := gconv.String(other)
	if gstr.IsNumeric(s1) && gstr.IsNumeric(s2) {
		return gconv.Int64(value) <= gconv.Int64(other)
	}
	return strings.Compare(s1, s2) <= 0
}

// buildInFuncGt 实现了内置模板函数：gt
func (view *View) buildInFuncGt(value, other interface{}) bool {
	s1 := gconv.String(value)
	s2 := gconv.String(other)
	if gstr.IsNumeric(s1) && gstr.IsNumeric(s2) {
		return gconv.Int64(value) > gconv.Int64(other)
	}
	return strings.Compare(s1, s2) > 0
}

// buildInFuncGe 实现了内置模板函数：ge
func (view *View) buildInFuncGe(value, other interface{}) bool {
	s1 := gconv.String(value)
	s2 := gconv.String(other)
	if gstr.IsNumeric(s1) && gstr.IsNumeric(s2) {
		return gconv.Int64(value) >= gconv.Int64(other)
	}
	return strings.Compare(s1, s2) >= 0
}

// buildInFuncInclude 实现了内建模板函数：include
// 注意，配置项 AutoEncode 不会影响此函数的输出结果。
func (view *View) buildInFuncInclude(file interface{}, data ...map[string]interface{}) htmltpl.HTML {
	var m map[string]interface{} = nil
	if len(data) > 0 {
		m = data[0]
	}
	path := gconv.String(file)
	if path == "" {
		return ""
	}
	// 它将会在内部搜索文件。
	content, err := view.Parse(context.TODO(), path, m)
	if err != nil {
		return htmltpl.HTML(err.Error())
	}
	return htmltpl.HTML(content)
}

// buildInFuncText 实现了内置模板函数：text
func (view *View) buildInFuncText(html interface{}) string {
	return ghtml.StripTags(gconv.String(html))
}

// buildInFuncHtmlEncode 实现了内置模板函数：html Encode
// 此函数用于对输入内容进行 HTML 编码，以确保在网页中安全显示，防止 XSS 攻击
func (view *View) buildInFuncHtmlEncode(html interface{}) string {
	return ghtml.Entities(gconv.String(html))
}

// buildInFuncHtmlDecode 实现了内置模板函数：htmldecode
func (view *View) buildInFuncHtmlDecode(html interface{}) string {
	return ghtml.EntitiesDecode(gconv.String(html))
}

// buildInFuncUrlEncode 实现了内置模板函数：url编码
// ```go
// 这个函数（buildInFuncUrlEncode）是 Go 语言中用于实现模板引擎内置功能的，
// 具体实现了对字符串进行 URL 编码的处理。
// 当在模板中调用 url 函数时，会调用此函数进行实际的 URL 编码操作。
func (view *View) buildInFuncUrlEncode(url interface{}) string {
	return gurl.Encode(gconv.String(url))
}

// buildInFuncUrlDecode 实现了内置模板函数：urldecode
func (view *View) buildInFuncUrlDecode(url interface{}) string {
	if content, err := gurl.Decode(gconv.String(url)); err == nil {
		return content
	} else {
		return err.Error()
	}
}

// buildInFuncDate 实现了内置模板函数：date
func (view *View) buildInFuncDate(format interface{}, timestamp ...interface{}) string {
	t := int64(0)
	if len(timestamp) > 0 {
		t = gconv.Int64(timestamp[0])
	}
	if t == 0 {
		t = gtime.Timestamp()
	}
	return gtime.NewFromTimeStamp(t).Format(gconv.String(format))
}

// buildInFuncCompare 实现了内置模板函数：compare
func (view *View) buildInFuncCompare(value1, value2 interface{}) int {
	return strings.Compare(gconv.String(value1), gconv.String(value2))
}

// buildInFuncSubStr 实现了内置模板函数：substr
func (view *View) buildInFuncSubStr(start, end, str interface{}) string {
	return gstr.SubStrRune(gconv.String(str), gconv.Int(start), gconv.Int(end))
}

// buildInFuncStrLimit 实现了内置模板函数：strlimit
func (view *View) buildInFuncStrLimit(length, suffix, str interface{}) string {
	return gstr.StrLimitRune(gconv.String(str), gconv.Int(length), gconv.String(suffix))
}

// buildInFuncConcat 实现了内置模板函数：concat
func (view *View) buildInFuncConcat(str ...interface{}) string {
	var s string
	for _, v := range str {
		s += gconv.String(v)
	}
	return s
}

// buildInFuncReplace 实现了内置模板函数：replace
func (view *View) buildInFuncReplace(search, replace, str interface{}) string {
	return gstr.Replace(gconv.String(str), gconv.String(search), gconv.String(replace), -1)
}

// buildInFuncHighlight 实现了内置模板函数：highlight
func (view *View) buildInFuncHighlight(key, color, str interface{}) string {
	return gstr.Replace(gconv.String(str), gconv.String(key), fmt.Sprintf(`<span style="color:%v;">%v</span>`, color, key))
}

// buildInFuncHideStr 实现了内建模板函数：hidestr
func (view *View) buildInFuncHideStr(percent, hide, str interface{}) string {
	return gstr.HideStr(gconv.String(str), gconv.Int(percent), gconv.String(hide))
}

// buildInFuncToUpper 实现了内置模板函数：toupper
func (view *View) buildInFuncToUpper(str interface{}) string {
	return gstr.ToUpper(gconv.String(str))
}

// buildInFuncToLower 实现了内置模板函数：tolower
func (view *View) buildInFuncToLower(str interface{}) string {
	return gstr.ToLower(gconv.String(str))
}

// buildInFuncNl2Br 实现了内置模板函数：nl2br
func (view *View) buildInFuncNl2Br(str interface{}) string {
	return gstr.Nl2Br(gconv.String(str))
}

// buildInFuncJson 实现了内建模板函数：json ，
// 它将 `value` 编码并以 JSON 字符串的形式返回。
func (view *View) buildInFuncJson(value interface{}) (string, error) {
	b, err := gjson.Marshal(value)
	return string(b), err
}

// buildInFuncXml 实现了内建模板函数：xml ，
// 它将`value`编码并以 XML 字符串形式返回。
func (view *View) buildInFuncXml(value interface{}, rootTag ...string) (string, error) {
	b, err := gjson.New(value).ToXml(rootTag...)
	return string(b), err
}

// buildInFuncIni 实现了内置模板函数：ini，
// 它将`value`编码并以XML字符串形式返回。
func (view *View) buildInFuncIni(value interface{}) (string, error) {
	b, err := gjson.New(value).ToIni()
	return string(b), err
}

// buildInFuncYaml 实现了内置模板函数：yaml，
// 它将`value`编码并以 YAML 字符串形式返回。
func (view *View) buildInFuncYaml(value interface{}) (string, error) {
	b, err := gjson.New(value).ToYaml()
	return string(b), err
}

// buildInFuncYamlIndent 实现了内置模板函数：yamli，
// 它将`value`按照自定义缩进字符串编码并以 YAML 字符串形式返回。
func (view *View) buildInFuncYamlIndent(value, indent interface{}) (string, error) {
	b, err := gjson.New(value).ToYamlIndent(gconv.String(indent))
	return string(b), err
}

// buildInFuncToml 实现了内置模板函数：toml，
// 它将`value`编码并以TOML字符串的形式返回。
func (view *View) buildInFuncToml(value interface{}) (string, error) {
	b, err := gjson.New(value).ToToml()
	return string(b), err
}

// buildInFuncPlus 实现了内建模板函数：plus ，
// 它返回将所有 `deltas` 加到 `value` 后的结果。
func (view *View) buildInFuncPlus(value interface{}, deltas ...interface{}) string {
	result := gconv.Float64(value)
	for _, v := range deltas {
		result += gconv.Float64(v)
	}
	return gconv.String(result)
}

// buildInFuncMinus 实现了内置模板函数：minus，
// 它返回从 `value` 中减去所有 `deltas` 后的结果。
func (view *View) buildInFuncMinus(value interface{}, deltas ...interface{}) string {
	result := gconv.Float64(value)
	for _, v := range deltas {
		result -= gconv.Float64(v)
	}
	return gconv.String(result)
}

// buildInFuncTimes 实现了内置模板函数：times ，
// 它返回将 `value` 与 `values` 中的所有值相乘的结果。
func (view *View) buildInFuncTimes(value interface{}, values ...interface{}) string {
	result := gconv.Float64(value)
	for _, v := range values {
		result *= gconv.Float64(v)
	}
	return gconv.String(result)
}

// buildInFuncDivide 实现了内置模板函数：divide，
// 它将 `value` 除以 `values` 中的所有值并返回结果。
func (view *View) buildInFuncDivide(value interface{}, values ...interface{}) string {
	result := gconv.Float64(value)
	for _, v := range values {
		value2Float64 := gconv.Float64(v)
		if value2Float64 == 0 {
			// Invalid `value2`.
			return "0"
		}
		result /= value2Float64
	}
	return gconv.String(result)
}
