// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gjson

import (
	"github.com/gogf/gf/v2/encoding/gini"
	"github.com/gogf/gf/v2/encoding/gproperties"
	"github.com/gogf/gf/v2/encoding/gtoml"
	"github.com/gogf/gf/v2/encoding/gxml"
	"github.com/gogf/gf/v2/encoding/gyaml"
	"github.com/888go/goframe/gjson/internal/json"
)

// ===========================================================================
// JSON
// ===========================================================================
// （此部分代码注释为标题性描述，翻译如下）
// ============================================================================
// JSON 部分
// ============================================================================
// （这里表示该段代码与 JSON 相关，可能是用于处理、解析或生成 JSON 的功能模块）

func (j *Json) ToJson() ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return Encode(*(j.p))
}

func (j *Json) ToJsonString() (string, error) {
	b, e := j.ToJson()
	return string(b), e
}

func (j *Json) ToJsonIndent() ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return json.MarshalIndent(*(j.p), "", "\t")
}

func (j *Json) ToJsonIndentString() (string, error) {
	b, e := j.ToJsonIndent()
	return string(b), e
}

func (j *Json) MustToJson() []byte {
	result, err := j.ToJson()
	if err != nil {
		panic(err)
	}
	return result
}

func (j *Json) MustToJsonString() string {
	return string(j.MustToJson())
}

func (j *Json) MustToJsonIndent() []byte {
	result, err := j.ToJsonIndent()
	if err != nil {
		panic(err)
	}
	return result
}

func (j *Json) MustToJsonIndentString() string {
	return string(j.MustToJsonIndent())
}

// ========================================================================
// XML
// ========================================================================
// ===中文注释开始===
// 此部分代码与XML相关功能实现
// ========================================================================

func (j *Json) ToXml(rootTag ...string) ([]byte, error) {
	return gxml.Encode(j.Var().Map(), rootTag...)
}

func (j *Json) ToXmlString(rootTag ...string) (string, error) {
	b, e := j.ToXml(rootTag...)
	return string(b), e
}

func (j *Json) ToXmlIndent(rootTag ...string) ([]byte, error) {
	return gxml.EncodeWithIndent(j.Var().Map(), rootTag...)
}

func (j *Json) ToXmlIndentString(rootTag ...string) (string, error) {
	b, e := j.ToXmlIndent(rootTag...)
	return string(b), e
}

func (j *Json) MustToXml(rootTag ...string) []byte {
	result, err := j.ToXml(rootTag...)
	if err != nil {
		panic(err)
	}
	return result
}

func (j *Json) MustToXmlString(rootTag ...string) string {
	return string(j.MustToXml(rootTag...))
}

func (j *Json) MustToXmlIndent(rootTag ...string) []byte {
	result, err := j.ToXmlIndent(rootTag...)
	if err != nil {
		panic(err)
	}
	return result
}

func (j *Json) MustToXmlIndentString(rootTag ...string) string {
	return string(j.MustToXmlIndent(rootTag...))
}

// ==========================================================================
// YAML
// ==========================================================================
// （这段代码的注释表明了该部分代码与YAML相关，但没有给出具体的功能描述，故无法精确翻译更多内容。）
// ========================================================================
// 表示分割线，用于组织和划分代码区域
// YAML
// 这行注释表示这部分代码与YAML（一种数据序列化格式）有关

func (j *Json) ToYaml() ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return gyaml.Encode(*(j.p))
}

func (j *Json) ToYamlIndent(indent string) ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return gyaml.EncodeIndent(*(j.p), indent)
}

func (j *Json) ToYamlString() (string, error) {
	b, e := j.ToYaml()
	return string(b), e
}

func (j *Json) MustToYaml() []byte {
	result, err := j.ToYaml()
	if err != nil {
		panic(err)
	}
	return result
}

func (j *Json) MustToYamlString() string {
	return string(j.MustToYaml())
}

// ===========================================================================
// TOML
// ===========================================================================
// （注释翻译：）
// ============================================================================
// TOML
// ============================================================================
// 这段代码的注释表明了该部分与TOML相关。TOML是一种配置文件格式，因此这部分可能是用于解析、生成或操作TOML格式数据的Go语言代码段。此处的“//”表示单行注释，中文翻译后含义不变，仅将原有的英文内容转换为中文展示。

func (j *Json) ToToml() ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return gtoml.Encode(*(j.p))
}

func (j *Json) ToTomlString() (string, error) {
	b, e := j.ToToml()
	return string(b), e
}

func (j *Json) MustToToml() []byte {
	result, err := j.ToToml()
	if err != nil {
		panic(err)
	}
	return result
}

func (j *Json) MustToTomlString() string {
	return string(j.MustToToml())
}

// ===========================================================================
// INI
// ===========================================================================
// （此部分代码的注释表明该部分与INI文件相关，但没有提供详细信息，故直接翻译为中文即可）
// ===========================================================================
// INI配置文件相关代码
// ===========================================================================

// ToIni json to ini
func (j *Json) ToIni() ([]byte, error) {
	return gini.Encode(j.Map())
}

// ToIniString 将ini转换为字符串
func (j *Json) ToIniString() (string, error) {
	b, e := j.ToIni()
	return string(b), e
}

func (j *Json) MustToIni() []byte {
	result, err := j.ToIni()
	if err != nil {
		panic(err)
	}
	return result
}

// MustToIniString .
func (j *Json) MustToIniString() string {
	return string(j.MustToIni())
}

// ==========================================================================
// 属性
// ==========================================================================
// 将JSON转换为属性
func (j *Json) ToProperties() ([]byte, error) {
	return gproperties.Encode(j.Map())
}

// TopropertiesString 将属性转换为字符串
func (j *Json) ToPropertiesString() (string, error) {
	b, e := j.ToProperties()
	return string(b), e
}

func (j *Json) MustToProperties() []byte {
	result, err := j.ToProperties()
	if err != nil {
		panic(err)
	}
	return result
}

// MustToPropertiesString
func (j *Json) MustToPropertiesString() string {
	return string(j.MustToProperties())
}
