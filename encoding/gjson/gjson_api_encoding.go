// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package json类

import (
	gini "github.com/888go/goframe/encoding/gini"
	"github.com/888go/goframe/encoding/gproperties"
	gtoml "github.com/888go/goframe/encoding/gtoml"
	gxml "github.com/888go/goframe/encoding/gxml"
	gyaml "github.com/888go/goframe/encoding/gyaml"
	"github.com/888go/goframe/internal/json"
)

// ========================================================================
// JSON
// ========================================================================
// md5:9f636a538977ae4f

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
// md5:931c367389ad5867

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

// ========================================================================
// YAML
// ========================================================================
// md5:86131a4a0253d702

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

// ========================================================================
// TOML 配置文件格式
// ========================================================================
// md5:2a6d07eba917d4f3

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

// ========================================================================
// INI
// ========================================================================
// md5:a7d46faaad75eec6

// ToIni json to ini
func (j *Json) ToIni() ([]byte, error) {
	return gini.Encode(j.Map())
}

// ToIniString 将ini格式转换为字符串. md5:954c17725442fbb6
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

// ========================================================================
// 属性
// ========================================================================
// 将json格式的属性转换为properties格式
// md5:83a506c62c95394b
func (j *Json) ToProperties() ([]byte, error) {
	return gproperties.Encode(j.Map())
}

// TopropertiesString 将属性转换为字符串. md5:4e7ae41f91f6945a
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

// MustTopropertiesString
func (j *Json) MustToPropertiesString() string {
	return string(j.MustToProperties())
}
