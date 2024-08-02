// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gyaml提供了访问和转换YAML内容的功能。 md5:f1323f1f471201c0
package yaml类

import (
	"bytes"
	"strings"

	"gopkg.in/yaml.v3"

	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	gconv "github.com/888go/goframe/util/gconv"
)

// Encode 将 `value` 编码为字节形式的 YAML 格式内容。 md5:df5adae4088970d4
func Encode(value interface{}) (out []byte, err error) {
	if out, err = yaml.Marshal(value); err != nil {
		err = gerror.Wrap(err, `yaml.Marshal failed`)
	}
	return
}

// EncodeIndent 将 `value` 编码为带有缩进的 YAML 格式内容，返回字节切片。 md5:08a35501baf3e352
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

// Decode将`content`解析为映射并返回。 md5:09afa737da32e1d3
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

// DecodeTo 将 content 解析到 result 中。 md5:39ca50eceb6c38b3
func DecodeTo(value []byte, result interface{}) (err error) {
	err = yaml.Unmarshal(value, result)
	if err != nil {
		err = gerror.Wrap(err, `yaml.Unmarshal failed`)
	}
	return
}

// ToJson 将 `content` 转换为 JSON 格式的内容。 md5:8ea19c8cff7962ab
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
