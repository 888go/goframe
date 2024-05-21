// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

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


// ff:取json字节集
func (j *Json) ToJson() ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return Encode(*(j.p))
}


// ff:取json文本
func (j *Json) ToJsonString() (string, error) {
	b, e := j.ToJson()
	return string(b), e
}


// ff:取json字节集并格式化
func (j *Json) ToJsonIndent() ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return json.MarshalIndent(*(j.p), "", "\t")
}


// ff:取json文本并格式化
func (j *Json) ToJsonIndentString() (string, error) {
	b, e := j.ToJsonIndent()
	return string(b), e
}


// ff:取json字节集PANI
func (j *Json) MustToJson() []byte {
	result, err := j.ToJson()
	if err != nil {
		panic(err)
	}
	return result
}


// ff:取json文本PANI
func (j *Json) MustToJsonString() string {
	return string(j.MustToJson())
}


// ff:取json字节集并格式化PANI
func (j *Json) MustToJsonIndent() []byte {
	result, err := j.ToJsonIndent()
	if err != nil {
		panic(err)
	}
	return result
}


// ff:取json文本并格式化PANI
func (j *Json) MustToJsonIndentString() string {
	return string(j.MustToJsonIndent())
}

// ========================================================================
// XML
// ========================================================================


// ff:取xml字节集
// rootTag:
func (j *Json) ToXml(rootTag ...string) ([]byte, error) {
	return gxml.Encode(j.Var().Map(), rootTag...)
}


// ff:取xml文本
// rootTag:
func (j *Json) ToXmlString(rootTag ...string) (string, error) {
	b, e := j.ToXml(rootTag...)
	return string(b), e
}


// ff:取xml字节集并格式化
// rootTag:
func (j *Json) ToXmlIndent(rootTag ...string) ([]byte, error) {
	return gxml.EncodeWithIndent(j.Var().Map(), rootTag...)
}


// ff:取xml文本并格式化
// rootTag:
func (j *Json) ToXmlIndentString(rootTag ...string) (string, error) {
	b, e := j.ToXmlIndent(rootTag...)
	return string(b), e
}


// ff:取xml字节集PANI
// rootTag:
func (j *Json) MustToXml(rootTag ...string) []byte {
	result, err := j.ToXml(rootTag...)
	if err != nil {
		panic(err)
	}
	return result
}


// ff:取xml文本PANI
// rootTag:
func (j *Json) MustToXmlString(rootTag ...string) string {
	return string(j.MustToXml(rootTag...))
}


// ff:取xml字节集并格式化PANI
// rootTag:
func (j *Json) MustToXmlIndent(rootTag ...string) []byte {
	result, err := j.ToXmlIndent(rootTag...)
	if err != nil {
		panic(err)
	}
	return result
}


// ff:取xml文本并格式化PANI
// rootTag:
func (j *Json) MustToXmlIndentString(rootTag ...string) string {
	return string(j.MustToXmlIndent(rootTag...))
}

// ========================================================================
// YAML
// ========================================================================


// ff:取YAML字节集
func (j *Json) ToYaml() ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return gyaml.Encode(*(j.p))
}


// ff:取YAML字节集并格式化
// indent:缩进
func (j *Json) ToYamlIndent(indent string) ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return gyaml.EncodeIndent(*(j.p), indent)
}


// ff:取YAML文本
func (j *Json) ToYamlString() (string, error) {
	b, e := j.ToYaml()
	return string(b), e
}


// ff:取YAML字节集PANI
func (j *Json) MustToYaml() []byte {
	result, err := j.ToYaml()
	if err != nil {
		panic(err)
	}
	return result
}


// ff:取YAML文本PANI
func (j *Json) MustToYamlString() string {
	return string(j.MustToYaml())
}

// ========================================================================
// TOML
// ========================================================================


// ff:取TOML字节集
func (j *Json) ToToml() ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return gtoml.Encode(*(j.p))
}


// ff:取TOML文本
func (j *Json) ToTomlString() (string, error) {
	b, e := j.ToToml()
	return string(b), e
}


// ff:取TOML字节集PANI
func (j *Json) MustToToml() []byte {
	result, err := j.ToToml()
	if err != nil {
		panic(err)
	}
	return result
}


// ff:取TOML文本PANI
func (j *Json) MustToTomlString() string {
	return string(j.MustToToml())
}

// ========================================================================
// INI
// ========================================================================

// ToIni json to ini

// ff:取ini字节集
func (j *Json) ToIni() ([]byte, error) {
	return gini.Encode(j.Map())
}

// ToIniString ini to string

// ff:取ini文本
func (j *Json) ToIniString() (string, error) {
	b, e := j.ToIni()
	return string(b), e
}


// ff:取ini字节集PANI
func (j *Json) MustToIni() []byte {
	result, err := j.ToIni()
	if err != nil {
		panic(err)
	}
	return result
}

// MustToIniString .

// ff:取ini文本PANI
func (j *Json) MustToIniString() string {
	return string(j.MustToIni())
}

// ========================================================================
// properties
// ========================================================================
// Toproperties json to properties

// ff:取properties字节集
func (j *Json) ToProperties() ([]byte, error) {
	return gproperties.Encode(j.Map())
}

// TopropertiesString properties to string

// ff:取properties文本
func (j *Json) ToPropertiesString() (string, error) {
	b, e := j.ToProperties()
	return string(b), e
}


// ff:取properties字节集PANI
func (j *Json) MustToProperties() []byte {
	result, err := j.ToProperties()
	if err != nil {
		panic(err)
	}
	return result
}

// MustTopropertiesString

// ff:取properties文本PANI
func (j *Json) MustToPropertiesString() string {
	return string(j.MustToProperties())
}
