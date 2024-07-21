// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gjson

import (
	"github.com/gogf/gf/v2/encoding/gini"
	"github.com/gogf/gf/v2/encoding/gproperties"
	"github.com/gogf/gf/v2/encoding/gtoml"
	"github.com/gogf/gf/v2/encoding/gxml"
	"github.com/gogf/gf/v2/encoding/gyaml"
	"github.com/gogf/gf/v2/internal/json"
)

// ========================================================================
// JSON
// ========================================================================
// md5:9f636a538977ae4f

// ff:取json字节集
// j:
func (j *Json) ToJson() ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return Encode(*(j.p))
}

// ff:取json文本
// j:
func (j *Json) ToJsonString() (string, error) {
	b, e := j.ToJson()
	return string(b), e
}

// ff:取json字节集并格式化
// j:
func (j *Json) ToJsonIndent() ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return json.MarshalIndent(*(j.p), "", "\t")
}

// ff:取json文本并格式化
// j:
func (j *Json) ToJsonIndentString() (string, error) {
	b, e := j.ToJsonIndent()
	return string(b), e
}

// ff:取json字节集PANI
// j:
func (j *Json) MustToJson() []byte {
	result, err := j.ToJson()
	if err != nil {
		panic(err)
	}
	return result
}

// ff:取json文本PANI
// j:
func (j *Json) MustToJsonString() string {
	return string(j.MustToJson())
}

// ff:取json字节集并格式化PANI
// j:
func (j *Json) MustToJsonIndent() []byte {
	result, err := j.ToJsonIndent()
	if err != nil {
		panic(err)
	}
	return result
}

// ff:取json文本并格式化PANI
// j:
func (j *Json) MustToJsonIndentString() string {
	return string(j.MustToJsonIndent())
}

// ========================================================================
// XML
// ========================================================================
// md5:931c367389ad5867

// ff:取xml字节集
// j:
// rootTag:
func (j *Json) ToXml(rootTag ...string) ([]byte, error) {
	return gxml.Encode(j.Var().Map(), rootTag...)
}

// ff:取xml文本
// j:
// rootTag:
func (j *Json) ToXmlString(rootTag ...string) (string, error) {
	b, e := j.ToXml(rootTag...)
	return string(b), e
}

// ff:取xml字节集并格式化
// j:
// rootTag:
func (j *Json) ToXmlIndent(rootTag ...string) ([]byte, error) {
	return gxml.EncodeWithIndent(j.Var().Map(), rootTag...)
}

// ff:取xml文本并格式化
// j:
// rootTag:
func (j *Json) ToXmlIndentString(rootTag ...string) (string, error) {
	b, e := j.ToXmlIndent(rootTag...)
	return string(b), e
}

// ff:取xml字节集PANI
// j:
// rootTag:
func (j *Json) MustToXml(rootTag ...string) []byte {
	result, err := j.ToXml(rootTag...)
	if err != nil {
		panic(err)
	}
	return result
}

// ff:取xml文本PANI
// j:
// rootTag:
func (j *Json) MustToXmlString(rootTag ...string) string {
	return string(j.MustToXml(rootTag...))
}

// ff:取xml字节集并格式化PANI
// j:
// rootTag:
func (j *Json) MustToXmlIndent(rootTag ...string) []byte {
	result, err := j.ToXmlIndent(rootTag...)
	if err != nil {
		panic(err)
	}
	return result
}

// ff:取xml文本并格式化PANI
// j:
// rootTag:
func (j *Json) MustToXmlIndentString(rootTag ...string) string {
	return string(j.MustToXmlIndent(rootTag...))
}

// ========================================================================
// YAML
// ========================================================================
// md5:86131a4a0253d702

// ff:取YAML字节集
// j:
func (j *Json) ToYaml() ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return gyaml.Encode(*(j.p))
}

// ff:取YAML字节集并格式化
// j:
// indent:缩进
func (j *Json) ToYamlIndent(indent string) ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return gyaml.EncodeIndent(*(j.p), indent)
}

// ff:取YAML文本
// j:
func (j *Json) ToYamlString() (string, error) {
	b, e := j.ToYaml()
	return string(b), e
}

// ff:取YAML字节集PANI
// j:
func (j *Json) MustToYaml() []byte {
	result, err := j.ToYaml()
	if err != nil {
		panic(err)
	}
	return result
}

// ff:取YAML文本PANI
// j:
func (j *Json) MustToYamlString() string {
	return string(j.MustToYaml())
}

// ========================================================================
// TOML 配置文件格式
// ========================================================================
// md5:2a6d07eba917d4f3

// ff:取TOML字节集
// j:
func (j *Json) ToToml() ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return gtoml.Encode(*(j.p))
}

// ff:取TOML文本
// j:
func (j *Json) ToTomlString() (string, error) {
	b, e := j.ToToml()
	return string(b), e
}

// ff:取TOML字节集PANI
// j:
func (j *Json) MustToToml() []byte {
	result, err := j.ToToml()
	if err != nil {
		panic(err)
	}
	return result
}

// ff:取TOML文本PANI
// j:
func (j *Json) MustToTomlString() string {
	return string(j.MustToToml())
}

// ========================================================================
// INI
// ========================================================================
// md5:a7d46faaad75eec6

// ToIni json to ini
// ff:取ini字节集
// j:
func (j *Json) ToIni() ([]byte, error) {
	return gini.Encode(j.Map())
}

// ToIniString 将ini格式转换为字符串. md5:954c17725442fbb6
// ff:取ini文本
// j:
func (j *Json) ToIniString() (string, error) {
	b, e := j.ToIni()
	return string(b), e
}

// ff:取ini字节集PANI
// j:
func (j *Json) MustToIni() []byte {
	result, err := j.ToIni()
	if err != nil {
		panic(err)
	}
	return result
}

// MustToIniString .
// ff:取ini文本PANI
// j:
func (j *Json) MustToIniString() string {
	return string(j.MustToIni())
}

// ========================================================================
// 属性
// ========================================================================
// 将json格式的属性转换为properties格式
// md5:83a506c62c95394b
// ff:取properties字节集
// j:
func (j *Json) ToProperties() ([]byte, error) {
	return gproperties.Encode(j.Map())
}

// TopropertiesString 将属性转换为字符串. md5:4e7ae41f91f6945a
// ff:取properties文本
// j:
func (j *Json) ToPropertiesString() (string, error) {
	b, e := j.ToProperties()
	return string(b), e
}

// ff:取properties字节集PANI
// j:
func (j *Json) MustToProperties() []byte {
	result, err := j.ToProperties()
	if err != nil {
		panic(err)
	}
	return result
}

// MustTopropertiesString
// ff:取properties文本PANI
// j:
func (j *Json) MustToPropertiesString() string {
	return string(j.MustToProperties())
}
