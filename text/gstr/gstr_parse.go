// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gstr
import (
	"net/url"
	"strings"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	)
// Parse 将字符串解析为 map[string]interface{} 类型。
//
// v1=m&v2=n           -> 解析得到的映射：map[v1:m v2:n]
// v[a]=m&v[b]=n       -> 解析得到的映射：map[v:map[a:m b:n]]
// v[a][a]=m&v[a][b]=n -> 解析得到的映射：map[v:map[a:map[a:m b:n]]]
// v[]=m&v[]=n         -> 解析得到的映射：map[v:[m n]]
// v[a][]=m&v[a][]=n   -> 解析得到的映射：map[v:map[a:[m n]]]
// v[][]=m&v[][]=n     -> 解析得到的映射：map[v:[map[]]]  // 目前不支持嵌套切片
// v=m&v[a]=n          -> 报错
// a .[[b=c            -> 解析得到的映射：map[a___[b:c]]
// 注意，上述代码注释描述了一个将查询字符串形式的数据解析成 Go 语言中的 map 的功能。在处理嵌套结构时，它会根据键名包含的中括号 `[]` 和方括号 `[]` 来构建嵌套的 map 或 slice。不过需要注意的是，对于 "v[][]=m&v[][]=n" 这种情况，当前实现并不支持嵌套的 slice 结构。
func Parse(s string) (result map[string]interface{}, err error) {
	if s == "" {
		return nil, nil
	}
	result = make(map[string]interface{})
	parts := strings.Split(s, "&")
	for _, part := range parts {
		pos := strings.Index(part, "=")
		if pos <= 0 {
			continue
		}
		key, err := url.QueryUnescape(part[:pos])
		if err != nil {
			err = gerror.Wrapf(err, `url.QueryUnescape failed for string "%s"`, part[:pos])
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
			err = gerror.Wrapf(err, `url.QueryUnescape failed for string "%s"`, part[pos+1:])
			return nil, err
		}
		// 将其拆分为多个键
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
		if err = build(result, keys, value); err != nil {
			return nil, err
		}
	}
	return result, nil
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

	// 结尾是切片。例如 f[]，f[a][]
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
			return gerror.NewCodef(
				gcode.CodeInvalidParameter,
				"expected type '[]interface{}' for key '%s', but got '%T'",
				key, val,
			)
		}
		result[key] = append(children, value)
		return nil
	}
	// 结尾是切片加映射，形式如 v[][a]
	if keys[1] == "" && length > 2 && keys[2] != "" {
		val, ok := result[key]
		if !ok {
			result[key] = []interface{}{}
			val = result[key]
		}
		children, ok := val.([]interface{})
		if !ok {
			return gerror.NewCodef(
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

	// map，类似于 v[a]、v[a][b] 的用法
	val, ok := result[key]
	if !ok {
		result[key] = map[string]interface{}{}
		val = result[key]
	}
	children, ok := val.(map[string]interface{})
	if !ok {
		return gerror.NewCodef(
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
