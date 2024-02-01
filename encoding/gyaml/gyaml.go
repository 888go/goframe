// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gyaml 提供对 YAML 内容的访问和转换功能。
package gyaml
import (
	"bytes"
	"strings"
	
	"gopkg.in/yaml.v3"
	
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/util/gconv"
	)
// Encode 将`value`编码为YAML格式的字节内容。
func Encode(value interface{}) (out []byte, err error) {
	if out, err = yaml.Marshal(value); err != nil {
		err = gerror.Wrap(err, `yaml.Marshal failed`)
	}
	return
}

// EncodeIndent 将`value`编码为带缩进的YAML格式内容，并以字节形式输出。
func EncodeIndent(value interface{}, indent string) (out []byte, err error) {
	out, err = Encode(value)
	if err != nil {
		return
	}
	if indent != "" {
		var (
			buffer = bytes.NewBuffer(nil)
			array  = strings.Split(strings.TrimSpace(string(out)), "\n")
		)
		for _, v := range array {
			buffer.WriteString(indent)
			buffer.WriteString(v)
			buffer.WriteString("\n")
		}
		out = buffer.Bytes()
	}
	return
}

// Decode 解析 `content` 并将其转换为 map 后返回。
func Decode(content []byte) (map[string]interface{}, error) {
	var (
		result map[string]interface{}
		err    error
	)
	if err = yaml.Unmarshal(content, &result); err != nil {
		err = gerror.Wrap(err, `yaml.Unmarshal failed`)
		return nil, err
	}
	return gconv.MapDeep(result), nil
}

// DecodeTo 将 `content` 解析到 `result` 中。
func DecodeTo(value []byte, result interface{}) (err error) {
	err = yaml.Unmarshal(value, result)
	if err != nil {
		err = gerror.Wrap(err, `yaml.Unmarshal failed`)
	}
	return
}

// ToJson 将 `content` 转换为 JSON 格式的内容。
func ToJson(content []byte) (out []byte, err error) {
	var (
		result interface{}
	)
	if result, err = Decode(content); err != nil {
		return nil, err
	} else {
		return json.Marshal(result)
	}
}
