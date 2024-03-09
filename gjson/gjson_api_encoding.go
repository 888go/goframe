// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package json类

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

func (j *Json) X取json字节集() ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return X变量到json字节集(*(j.p))
}

func (j *Json) X取json文本() (string, error) {
	b, e := j.X取json字节集()
	return string(b), e
}

func (j *Json) X取json字节集并格式化() ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return json.MarshalIndent(*(j.p), "", "\t")
}

func (j *Json) X取json文本并格式化() (string, error) {
	b, e := j.X取json字节集并格式化()
	return string(b), e
}

func (j *Json) X取json字节集PANI() []byte {
	result, err := j.X取json字节集()
	if err != nil {
		panic(err)
	}
	return result
}

func (j *Json) X取json文本PANI() string {
	return string(j.X取json字节集PANI())
}

func (j *Json) X取json字节集并格式化PANI() []byte {
	result, err := j.X取json字节集并格式化()
	if err != nil {
		panic(err)
	}
	return result
}

func (j *Json) X取json文本并格式化PANI() string {
	return string(j.X取json字节集并格式化PANI())
}

// ========================================================================
// XML
// ========================================================================
// ===中文注释开始===
// 此部分代码与XML相关功能实现
// ========================================================================

func (j *Json) X取xml字节集(rootTag ...string) ([]byte, error) {
	return gxml.Encode(j.X取泛型类().Map(), rootTag...)
}

func (j *Json) X取xml文本(rootTag ...string) (string, error) {
	b, e := j.X取xml字节集(rootTag...)
	return string(b), e
}

func (j *Json) X取xml字节集并格式化(rootTag ...string) ([]byte, error) {
	return gxml.EncodeWithIndent(j.X取泛型类().Map(), rootTag...)
}

func (j *Json) X取xml文本并格式化(rootTag ...string) (string, error) {
	b, e := j.X取xml字节集并格式化(rootTag...)
	return string(b), e
}

func (j *Json) X取xml字节集PANI(rootTag ...string) []byte {
	result, err := j.X取xml字节集(rootTag...)
	if err != nil {
		panic(err)
	}
	return result
}

func (j *Json) X取xml文本PANI(rootTag ...string) string {
	return string(j.X取xml字节集PANI(rootTag...))
}

func (j *Json) X取xml字节集并格式化PANI(rootTag ...string) []byte {
	result, err := j.X取xml字节集并格式化(rootTag...)
	if err != nil {
		panic(err)
	}
	return result
}

func (j *Json) X取xml文本并格式化PANI(rootTag ...string) string {
	return string(j.X取xml字节集并格式化PANI(rootTag...))
}

// ==========================================================================
// YAML
// ==========================================================================
// （这段代码的注释表明了该部分代码与YAML相关，但没有给出具体的功能描述，故无法精确翻译更多内容。）
// ========================================================================
// 表示分割线，用于组织和划分代码区域
// YAML
// 这行注释表示这部分代码与YAML（一种数据序列化格式）有关

func (j *Json) X取YAML字节集() ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return gyaml.Encode(*(j.p))
}

func (j *Json) X取YAML字节集并格式化(缩进 string) ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return gyaml.EncodeIndent(*(j.p), 缩进)
}

func (j *Json) X取YAML文本() (string, error) {
	b, e := j.X取YAML字节集()
	return string(b), e
}

func (j *Json) X取YAML字节集PANI() []byte {
	result, err := j.X取YAML字节集()
	if err != nil {
		panic(err)
	}
	return result
}

func (j *Json) X取YAML文本PANI() string {
	return string(j.X取YAML字节集PANI())
}

// ===========================================================================
// TOML
// ===========================================================================
// （注释翻译：）
// ============================================================================
// TOML
// ============================================================================
// 这段代码的注释表明了该部分与TOML相关。TOML是一种配置文件格式，因此这部分可能是用于解析、生成或操作TOML格式数据的Go语言代码段。此处的“//”表示单行注释，中文翻译后含义不变，仅将原有的英文内容转换为中文展示。

func (j *Json) X取TOML字节集() ([]byte, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return gtoml.Encode(*(j.p))
}

func (j *Json) X取TOML文本() (string, error) {
	b, e := j.X取TOML字节集()
	return string(b), e
}

func (j *Json) X取TOML字节集PANI() []byte {
	result, err := j.X取TOML字节集()
	if err != nil {
		panic(err)
	}
	return result
}

func (j *Json) X取TOML文本PANI() string {
	return string(j.X取TOML字节集PANI())
}

// ===========================================================================
// INI
// ===========================================================================
// （此部分代码的注释表明该部分与INI文件相关，但没有提供详细信息，故直接翻译为中文即可）
// ===========================================================================
// INI配置文件相关代码
// ===========================================================================

// ToIni json to ini
func (j *Json) X取ini字节集() ([]byte, error) {
	return gini.Encode(j.X取Map())
}

// ToIniString 将ini转换为字符串
func (j *Json) X取ini文本() (string, error) {
	b, e := j.X取ini字节集()
	return string(b), e
}

func (j *Json) X取ini字节集PANI() []byte {
	result, err := j.X取ini字节集()
	if err != nil {
		panic(err)
	}
	return result
}

// MustToIniString .
func (j *Json) X取ini文本PANI() string {
	return string(j.X取ini字节集PANI())
}

// ==========================================================================
// 属性
// ==========================================================================
// 将JSON转换为属性
func (j *Json) X取properties字节集() ([]byte, error) {
	return gproperties.Encode(j.X取Map())
}

// TopropertiesString 将属性转换为字符串
func (j *Json) X取properties文本() (string, error) {
	b, e := j.X取properties字节集()
	return string(b), e
}

func (j *Json) X取properties字节集PANI() []byte {
	result, err := j.X取properties字节集()
	if err != nil {
		panic(err)
	}
	return result
}

// MustToPropertiesString
func (j *Json) X取properties文本PANI() string {
	return string(j.X取properties字节集PANI())
}
