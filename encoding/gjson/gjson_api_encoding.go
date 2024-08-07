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
// md5:931c367389ad5867

func (j *Json) X取xml字节集(rootTag ...string) ([]byte, error) {
	return gxml.Encode(j.X取泛型类().X取Map(), rootTag...)
}

func (j *Json) X取xml文本(rootTag ...string) (string, error) {
	b, e := j.X取xml字节集(rootTag...)
	return string(b), e
}

func (j *Json) X取xml字节集并格式化(rootTag ...string) ([]byte, error) {
	return gxml.EncodeWithIndent(j.X取泛型类().X取Map(), rootTag...)
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

// ========================================================================
// YAML
// ========================================================================
// md5:86131a4a0253d702

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

// ========================================================================
// TOML 配置文件格式
// ========================================================================
// md5:2a6d07eba917d4f3

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

// ========================================================================
// INI
// ========================================================================
// md5:a7d46faaad75eec6

// X取ini字节集 json to ini
func (j *Json) X取ini字节集() ([]byte, error) {
	return gini.Map到ini(j.X取Map())
}

// X取ini文本 将ini格式转换为字符串. md5:954c17725442fbb6
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

// X取ini文本PANI .
func (j *Json) X取ini文本PANI() string {
	return string(j.X取ini字节集PANI())
}

// ========================================================================
// 属性
// ========================================================================
// 将json格式的属性转换为properties格式
// md5:83a506c62c95394b
func (j *Json) X取properties字节集() ([]byte, error) {
	return gproperties.Encode(j.X取Map())
}

// TopropertiesString 将属性转换为字符串. md5:4e7ae41f91f6945a
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

// MustTopropertiesString
func (j *Json) X取properties文本PANI() string {
	return string(j.X取properties字节集PANI())
}
