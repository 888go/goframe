// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gproperties 提供了对 .properties 文件内容的访问和转换功能。
package gproperties

import (
	"bytes"
	"sort"
	"strings"
	
	"github.com/magiconair/properties"
	
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/util/gconv"
)

// Decode将属性格式转换为映射。
func Decode(data []byte) (res map[string]interface{}, err error) {
	res = make(map[string]interface{})
	pr, err := properties.Load(data, properties.UTF8)
	if err != nil || pr == nil {
		err = gerror.Wrapf(err, `Lib magiconair load Properties data failed.`)
		return nil, err
	}
	for _, key := range pr.Keys() {
		// 忽略存在性检查：我们知道它在那里
		value, _ := pr.Get(key)
		// 递归构建嵌套映射
		path := strings.Split(key, ".")
		lastKey := strings.ToLower(path[len(path)-1])
		deepestMap := deepSearch(res, path[0:len(path)-1])

		// 设置最内层的值
		deepestMap[lastKey] = value
	}
	return res, nil
}

// Encode 将map转换为属性格式。
func Encode(data map[string]interface{}) (res []byte, err error) {
	pr := properties.NewProperties()

	flattened := map[string]interface{}{}

	flattened = flattenAndMergeMap(flattened, data, "", ".")

	keys := make([]string, 0, len(flattened))

	for key := range flattened {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		_, _, err := pr.Set(key, gconv.String(flattened[key]))
		if err != nil {
			err = gerror.Wrapf(err, `Sets the property key to the corresponding value failed.`)
			return nil, err
		}
	}

	var buf bytes.Buffer

	_, err = pr.Write(&buf, properties.UTF8)
	if err != nil {
		err = gerror.Wrapf(err, `Properties Write buf failed.`)
		return nil, err
	}

	return buf.Bytes(), nil
}

// ToJson 将 .properties 格式转换为 JSON。
func ToJson(data []byte) (res []byte, err error) {
	prMap, err := Decode(data)
	if err != nil {
		return nil, err
	}
	return json.Marshal(prMap)
}

// deepSearch 深度扫描映射，根据序列"path"中列出的键索引进行查找。
// 期望最后一个值为另一个映射，并将其返回。
func deepSearch(m map[string]interface{}, path []string) map[string]interface{} {
	for _, k := range path {
		m2, ok := m[k]
		if !ok {
// 中间键不存在
// => 创建它并从此处继续
			m3 := make(map[string]interface{})
			m[k] = m3
			m = m3
			continue
		}
		m3, ok := m2.(map[string]interface{})
		if !ok {
			m3 = make(map[string]interface{})
			m[k] = m3
		}
		// 从这里继续搜索
		m = m3
	}
	return m
}

// flattenAndMergeMap 递归地将给定的映射扁平化并合并到一个新的映射中
func flattenAndMergeMap(shadow map[string]interface{}, m map[string]interface{}, prefix string, delimiter string) map[string]interface{} {
	if shadow != nil && prefix != "" && shadow[prefix] != nil {
		return shadow
	}

	var m2 map[string]interface{}
	if prefix != "" {
		prefix += delimiter
	}
	for k, val := range m {
		fullKey := prefix + k
		switch val.(type) {
		case map[string]interface{}:
			m2 = val.(map[string]interface{})
		case map[interface{}]interface{}:
			m2 = gconv.Map(val)
		default:
			// immediate value
			shadow[strings.ToLower(fullKey)] = val
			continue
		}
		// 递归合并至影子映射
		shadow = flattenAndMergeMap(shadow, m2, fullKey, delimiter)
	}
	return shadow
}
