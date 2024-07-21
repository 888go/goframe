// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gproperties提供了对.properties内容的访问和转换。 md5:d7f77f0eb45bfdad
package gproperties

import (
	"bytes"
	"sort"
	"strings"

	"github.com/magiconair/properties"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/util/gconv"
)

// Decode将属性格式转换为映射。 md5:022b98c96d23f910
// ff:
// data:
// res:
// err:
func Decode(data []byte) (res map[string]interface{}, err error) {
	res = make(map[string]interface{})
	pr, err := properties.Load(data, properties.UTF8)
	if err != nil || pr == nil {
		err = gerror.Wrapf(err, `Lib magiconair load Properties data failed.`)
		return nil, err
	}
	for _, key := range pr.Keys() {
		// 忽略存在性检查：我们知道它就在那里. md5:ab11cdfb730ab02c
		value, _ := pr.Get(key)
		// 递归地构建嵌套映射. md5:1bc0faa7e615b22a
		path := strings.Split(key, ".")
		lastKey := strings.ToLower(path[len(path)-1])
		deepestMap := deepSearch(res, path[0:len(path)-1])

		// set innermost value
		deepestMap[lastKey] = value
	}
	return res, nil
}

// Encode 将映射转换为属性格式。 md5:d1876189b2478c4b
// ff:
// data:
// res:
// err:
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

// ToJson 将.properties格式转换为JSON。 md5:1575bc15c05b514f
// ff:
// data:
// res:
// err:
func ToJson(data []byte) (res []byte, err error) {
	prMap, err := Decode(data)
	if err != nil {
		return nil, err
	}
	return json.Marshal(prMap)
}

// deepSearch 用于深入扫描嵌套的映射，它会按照序列"path"中列出的键索引进行遍历。
// 预期最后一个值是另一个映射，并将其返回。
// md5:2b80516b778b8ffa
func deepSearch(m map[string]interface{}, path []string) map[string]interface{} {
	for _, k := range path {
		m2, ok := m[k]
		if !ok {
			// 中间键不存在
			// => 创建它并从此处继续
			// md5:ea01acf7f923de86
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
		// 从这里继续搜索. md5:fb1246c13ecceb40
		m = m3
	}
	return m
}

// flattenAndMergeMap递归地将给定的映射扁平化为一个新的映射. md5:4aca04c7957a8f20
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
											// 递归地将内容合并到阴影映射中. md5:89e72bf601f325cb
		shadow = flattenAndMergeMap(shadow, m2, fullKey, delimiter)
	}
	return shadow
}
