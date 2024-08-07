// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文本类

import (
	"net/url"
	"strings"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
)

// X参数解析 将字符串解析为 map[string]interface{} 类型。
//
// v1=m&v2=n           -> map[v1:m, v2:n]
// v[a]=m&v[b]=n       -> map[v:map[a:m, b:n]]
// v[a][a]=m&v[a][b]=n -> map[v:map[a:map[a:m, b:n]]]
// v[]=m&v[]=n         -> map[v:[m, n]]
// v[a][]=m&v[a][]=n   -> map[v:map[a:[m, n]]]
// v[][]=m&v[][]=n     -> map[v:[map[]]] // 当前不支持嵌套切片。
// v=m&v[a]=n          -> 错误
// a .[[b=c            -> 无法解析，缺少有效的键值对格式。
// md5:28f985708060eab0
func X参数解析(文本 string) (map结果 map[string]interface{}, 错误 error) {
	if 文本 == "" {
		return nil, nil
	}
	map结果 = make(map[string]interface{})
	parts := strings.Split(文本, "&")
	for _, part := range parts {
		pos := strings.Index(part, "=")
		if pos <= 0 {
			continue
		}
		key, err := url.QueryUnescape(part[:pos])
		if err != nil {
			err = gerror.X多层错误并格式化(err, `url.QueryUnescape failed for string "%s"`, part[:pos])
			return nil, err
		}

		for len(key) > 0 && key[0] == ' ' {
			key = key[1:]
		}

		if key == "" || key[0] == '[' {
			continue
		}
		value, err := url.QueryUnescape(part[pos+1:])
		if err != nil {
			err = gerror.X多层错误并格式化(err, `url.QueryUnescape failed for string "%s"`, part[pos+1:])
			return nil, err
		}
						// 分割成多个键. md5:3bdb5e68a953321c
		var keys []string
		left := 0
		for i, k := range key {
			if k == '[' && left == 0 {
				left = i
			} else if k == ']' {
				if left > 0 {
					if len(keys) == 0 {
						keys = append(keys, key[:left])
					}
					keys = append(keys, key[left+1:i])
					left = 0
					if i+1 < len(key) && key[i+1] != '[' {
						break
					}
				}
			}
		}
		if len(keys) == 0 {
			keys = append(keys, key)
		}
		// first key
		first := ""
		for i, chr := range keys[0] {
			if chr == ' ' || chr == '.' || chr == '[' {
				first += "_"
			} else {
				first += string(chr)
			}
			if chr == '[' {
				first += keys[0][i+1:]
				break
			}
		}
		keys[0] = first

		// build nested map
		if err = build(map结果, keys, value); err != nil {
			return nil, err
		}
	}
	return map结果, nil
}

// build nested map.
func build(result map[string]interface{}, keys []string, value interface{}) error {
	var (
		length = len(keys)
		key    = strings.Trim(keys[0], "'\"")
	)
	if length == 1 {
		result[key] = value
		return nil
	}

		// "end" 是一个切片，类似于 f[] 或者 f[a][]. md5:41e332252e9d2da1
	if keys[1] == "" && length == 2 {
		// TODO nested slice
		if key == "" {
			return nil
		}
		val, ok := result[key]
		if !ok {
			result[key] = []interface{}{value}
			return nil
		}
		children, ok := val.([]interface{})
		if !ok {
			return gerror.X创建错误码并格式化(
				gcode.CodeInvalidParameter,
				"expected type '[]interface{}' for key '%s', but got '%T'",
				key, val,
			)
		}
		result[key] = append(children, value)
		return nil
	}
		// 结束是切片和映射。就像 v[][][a]. md5:79444379ed8ddfc4
	if keys[1] == "" && length > 2 && keys[2] != "" {
		val, ok := result[key]
		if !ok {
			result[key] = []interface{}{}
			val = result[key]
		}
		children, ok := val.([]interface{})
		if !ok {
			return gerror.X创建错误码并格式化(
				gcode.CodeInvalidParameter,
				"expected type '[]interface{}' for key '%s', but got '%T'",
				key, val,
			)
		}
		if l := len(children); l > 0 {
			if child, ok := children[l-1].(map[string]interface{}); ok {
				if _, ok := child[keys[2]]; !ok {
					_ = build(child, keys[2:], value)
					return nil
				}
			}
		}
		child := map[string]interface{}{}
		_ = build(child, keys[2:], value)
		result[key] = append(children, child)
		return nil
	}

		// 类似于 v[a]，v[a][b] 的映射. md5:e8aa555b3543c9ea
	val, ok := result[key]
	if !ok {
		result[key] = map[string]interface{}{}
		val = result[key]
	}
	children, ok := val.(map[string]interface{})
	if !ok {
		return gerror.X创建错误码并格式化(
			gcode.CodeInvalidParameter,
			"expected type 'map[string]interface{}' for key '%s', but got '%T'",
			key, val,
		)
	}
	if err := build(children, keys[1:], value); err != nil {
		return err
	}
	return nil
}
