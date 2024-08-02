// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 模板类

import (
	"bytes"
	"context"
	"fmt"
	htmltpl "html/template"
	"strings"

	ghtml "github.com/888go/goframe/encoding/ghtml"
	gjson "github.com/888go/goframe/encoding/gjson"
	gurl "github.com/888go/goframe/encoding/gurl"
	gtime "github.com/888go/goframe/os/gtime"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
	gmode "github.com/888go/goframe/util/gmode"
	gutil "github.com/888go/goframe/util/gutil"
)

// buildInFuncDump 实现了内置模板函数：dump. md5:1f15207852fa685e
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

// buildInFuncMap 实现了内置的模板函数：map. md5:a2fd9c249b8323ee
func (view *View) buildInFuncMap(value ...interface{}) map[string]interface{} {
	if len(value) > 0 {
		return gconv.Map(value[0])
	}
	return map[string]interface{}{}
}

// buildInFuncMaps 实现了内置的模板函数：maps. md5:883c67613b885970
func (view *View) buildInFuncMaps(value ...interface{}) []map[string]interface{} {
	if len(value) > 0 {
		return gconv.Maps(value[0])
	}
	return []map[string]interface{}{}
}

// buildInFuncEq 实现了内置模板函数：eq. md5:e7da41f3d71aaeaa
func (view *View) buildInFuncEq(value interface{}, others ...interface{}) bool {
	s := gconv.String(value)
	for _, v := range others {
		if strings.Compare(s, gconv.String(v)) == 0 {
			return true
		}
	}
	return false
}

// buildInFuncNe 实现了内置模板函数：ne. md5:cb7ab3f138a25f49
func (view *View) buildInFuncNe(value, other interface{}) bool {
	return strings.Compare(gconv.String(value), gconv.String(other)) != 0
}

// buildInFuncLt 实现了内置的模板函数：lt. md5:d893f944e85d6c9d
func (view *View) buildInFuncLt(value, other interface{}) bool {
	s1 := gconv.String(value)
	s2 := gconv.String(other)
	if gstr.IsNumeric(s1) && gstr.IsNumeric(s2) {
		return gconv.Int64(value) < gconv.Int64(other)
	}
	return strings.Compare(s1, s2) < 0
}

// buildInFuncLe 实现了内置模板函数：le. md5:d064dd2a61308c66
func (view *View) buildInFuncLe(value, other interface{}) bool {
	s1 := gconv.String(value)
	s2 := gconv.String(other)
	if gstr.IsNumeric(s1) && gstr.IsNumeric(s2) {
		return gconv.Int64(value) <= gconv.Int64(other)
	}
	return strings.Compare(s1, s2) <= 0
}

// buildInFuncGt 实现了内置模板函数：gt. md5:3726333feaaed038
func (view *View) buildInFuncGt(value, other interface{}) bool {
	s1 := gconv.String(value)
	s2 := gconv.String(other)
	if gstr.IsNumeric(s1) && gstr.IsNumeric(s2) {
		return gconv.Int64(value) > gconv.Int64(other)
	}
	return strings.Compare(s1, s2) > 0
}

// buildInFuncGe 实现了内置模板函数：ge. md5:e78013901a4cdefd
func (view *View) buildInFuncGe(value, other interface{}) bool {
	s1 := gconv.String(value)
	s2 := gconv.String(other)
	if gstr.IsNumeric(s1) && gstr.IsNumeric(s2) {
		return gconv.Int64(value) >= gconv.Int64(other)
	}
	return strings.Compare(s1, s2) >= 0
}

// buildInFuncInclude 实现了内置模板函数：include
// 注意，配置自动编码(AutoEncode)不会影响此函数的输出。
// md5:3741767b68e0d6cc
func (view *View) buildInFuncInclude(file interface{}, data ...map[string]interface{}) htmltpl.HTML {
	var m map[string]interface{} = nil
	if len(data) > 0 {
		m = data[0]
	}
	path := gconv.String(file)
	if path == "" {
		return ""
	}
		// 它会内部搜索文件。 md5:f9f7bcb705f25f28
	content, err := view.Parse(context.TODO(), path, m)
	if err != nil {
		return htmltpl.HTML(err.Error())
	}
	return htmltpl.HTML(content)
}

// buildInFuncText 实现了内置的模板函数：text. md5:def5d05fa8935495
func (view *View) buildInFuncText(html interface{}) string {
	return ghtml.StripTags(gconv.String(html))
}

// buildInFuncHtmlEncode 实现了内置模板函数：html编码. md5:28b73e10863aa821
func (view *View) buildInFuncHtmlEncode(html interface{}) string {
	return ghtml.Entities(gconv.String(html))
}

// buildInFuncHtmlDecode 实现了内置模板函数：htmldecode. md5:989afb1f98599297
func (view *View) buildInFuncHtmlDecode(html interface{}) string {
	return ghtml.EntitiesDecode(gconv.String(html))
}

// buildInFuncUrlEncode 实现内置模板函数：url编码. md5:abeb3ab9af4cbddc
func (view *View) buildInFuncUrlEncode(url interface{}) string {
	return gurl.Encode(gconv.String(url))
}

// buildInFuncUrlDecode 实现了内置模板函数：urldecode. md5:b672e7cf3e2a329f
func (view *View) buildInFuncUrlDecode(url interface{}) string {
	if content, err := gurl.Decode(gconv.String(url)); err == nil {
		return content
	} else {
		return err.Error()
	}
}

// buildInFuncDate 实现了内置的模板函数：date. md5:cb730bbf7ec749d5
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

// buildInFuncCompare 实现了内置模板函数：compare. md5:b64f353a2f21fd26
func (view *View) buildInFuncCompare(value1, value2 interface{}) int {
	return strings.Compare(gconv.String(value1), gconv.String(value2))
}

// buildInFuncSubStr 实现了内置模板函数：substr. md5:f9a0424bba321aa4
func (view *View) buildInFuncSubStr(start, end, str interface{}) string {
	return gstr.SubStrRune(gconv.String(str), gconv.Int(start), gconv.Int(end))
}

// buildInFuncStrLimit 实现了内置的模板函数：strlimit. md5:da5d54a81cf4013d
func (view *View) buildInFuncStrLimit(length, suffix, str interface{}) string {
	return gstr.StrLimitRune(gconv.String(str), gconv.Int(length), gconv.String(suffix))
}

// buildInFuncConcat 实现了内置模板函数：concat. md5:e6a5bbbe2f1fb9e6
func (view *View) buildInFuncConcat(str ...interface{}) string {
	var s string
	for _, v := range str {
		s += gconv.String(v)
	}
	return s
}

// buildInFuncReplace 实现了内置模板函数：replace. md5:0b6b5a9cd71b2968
func (view *View) buildInFuncReplace(search, replace, str interface{}) string {
	return gstr.Replace(gconv.String(str), gconv.String(search), gconv.String(replace), -1)
}

// buildInFuncHighlight 实现内置模板函数：highlight. md5:a48873ff349fcc97
func (view *View) buildInFuncHighlight(key, color, str interface{}) string {
	return gstr.Replace(gconv.String(str), gconv.String(key), fmt.Sprintf(`<span style="color:%v;">%v</span>`, color, key))
}

// buildInFuncHideStr 实现了内置的模板函数：hidestr. md5:bdb6684409108de6
func (view *View) buildInFuncHideStr(percent, hide, str interface{}) string {
	return gstr.HideStr(gconv.String(str), gconv.Int(percent), gconv.String(hide))
}

// buildInFuncToUpper 实现了内置模板函数：toupper. md5:b67b1e5ebf474cbb
func (view *View) buildInFuncToUpper(str interface{}) string {
	return gstr.ToUpper(gconv.String(str))
}

// buildInFuncToLower 实现内置模板函数：toupper. md5:137e000a2677ea04
func (view *View) buildInFuncToLower(str interface{}) string {
	return gstr.ToLower(gconv.String(str))
}

// buildInFuncNl2Br 实现了内置模板函数：nl2br. md5:a75c85e699d18527
func (view *View) buildInFuncNl2Br(str interface{}) string {
	return gstr.Nl2Br(gconv.String(str))
}

// buildInFuncJson 实现了模板内置函数：json，
// 该函数将 `value` 编码为JSON字符串并返回。
// md5:fb0da1ee66d4c2a4
func (view *View) buildInFuncJson(value interface{}) (string, error) {
	b, err := gjson.Marshal(value)
	return string(b), err
}

// buildInFuncXml 实现了内置模板函数：xml，它将`value`编码为XML字符串并返回。
// md5:c1a5971d91e92b28
func (view *View) buildInFuncXml(value interface{}, rootTag ...string) (string, error) {
	b, err := gjson.New(value).ToXml(rootTag...)
	return string(b), err
}

// buildInFuncIni 实现了内置模板函数：ini，它将 `value` 编码为 XML 字符串并返回。
// md5:e58c98b4d09ac61d
func (view *View) buildInFuncIni(value interface{}) (string, error) {
	b, err := gjson.New(value).ToIni()
	return string(b), err
}

// buildInFuncYaml 实现了内置模板函数：yaml，它将 `value` 编码为 YAML 字符串并返回。
// md5:04183ea12eb16c2c
func (view *View) buildInFuncYaml(value interface{}) (string, error) {
	b, err := gjson.New(value).ToYaml()
	return string(b), err
}

// buildInFuncYamlIndent 实现了构建内置模板函数：yamli，
// 该函数将 `value` 编码为具有自定义缩进字符串的 YAML 格式字符串并返回。
// md5:2b1af65bb7bac809
func (view *View) buildInFuncYamlIndent(value, indent interface{}) (string, error) {
	b, err := gjson.New(value).ToYamlIndent(gconv.String(indent))
	return string(b), err
}

// buildInFuncToml 实现了内建的模板函数：toml，该函数将 `value` 编码为 TOML 字符串并返回。
// md5:fc418b3314bd75fd
func (view *View) buildInFuncToml(value interface{}) (string, error) {
	b, err := gjson.New(value).ToToml()
	return string(b), err
}

// buildInFuncPlus 实现了内置模板函数：plus，它将所有 `deltas` 加到 `value` 上并返回结果。
// md5:66a5ee3d0a30fd00
func (view *View) buildInFuncPlus(value interface{}, deltas ...interface{}) string {
	result := gconv.Float64(value)
	for _, v := range deltas {
		result += gconv.Float64(v)
	}
	return gconv.String(result)
}

// buildInFuncMinus 实现了内置模板函数：minus，它从"value"中减去所有 "deltas" 并返回结果。
// md5:3a8c7bc3d577d854
func (view *View) buildInFuncMinus(value interface{}, deltas ...interface{}) string {
	result := gconv.Float64(value)
	for _, v := range deltas {
		result -= gconv.Float64(v)
	}
	return gconv.String(result)
}

// buildInFuncTimes 实现了内置模板函数：times，
// 该函数返回 `value` 与所有 `values` 元素相乘的结果。
// md5:5e5ba3a1856c3b44
func (view *View) buildInFuncTimes(value interface{}, values ...interface{}) string {
	result := gconv.Float64(value)
	for _, v := range values {
		result *= gconv.Float64(v)
	}
	return gconv.String(result)
}

// buildInFuncDivide 实现了内置模板函数：divide，它返回将`value`除以所有`values`的结果。
// md5:55b84c767c41e466
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
