// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gini 提供了对 INI 内容的访问和转换功能。
package gini

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
)

// Decode将INI格式转换为map。
func Decode(data []byte) (res map[string]interface{}, err error) {
	res = make(map[string]interface{})
	var (
		fieldMap    = make(map[string]interface{})
		bytesReader = bytes.NewReader(data)
		bufioReader = bufio.NewReader(bytesReader)
		section     string
		lastSection string
		haveSection bool
		line        string
	)

	for {
		line, err = bufioReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			err = gerror.Wrapf(err, `bufioReader.ReadString failed`)
			return nil, err
		}
		if line = strings.TrimSpace(line); len(line) == 0 {
			continue
		}

		if line[0] == ';' || line[0] == '#' {
			continue
		}
		var (
			sectionBeginPos = strings.Index(line, "[")
			sectionEndPos   = strings.Index(line, "]")
		)
		if sectionBeginPos >= 0 && sectionEndPos >= 2 {
			section = line[sectionBeginPos+1 : sectionEndPos]
			if lastSection == "" {
				lastSection = section
			} else if lastSection != section {
				lastSection = section
				fieldMap = make(map[string]interface{})
			}
			haveSection = true
		} else if !haveSection {
			continue
		}

		if strings.Contains(line, "=") && haveSection {
			values := strings.Split(line, "=")
			fieldMap[strings.TrimSpace(values[0])] = strings.TrimSpace(strings.Join(values[1:], "="))
			res[section] = fieldMap
		}
	}

	if !haveSection {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "failed to parse INI file, section not found")
	}
	return res, nil
}

// Encode 将 map 转换为 INI 格式。
func Encode(data map[string]interface{}) (res []byte, err error) {
	var (
		n  int
		w  = new(bytes.Buffer)
		m  map[string]interface{}
		ok bool
	)
	for section, item := range data {
		// 配置项键值对。
		if m, ok = item.(map[string]interface{}); ok {
			n, err = w.WriteString(fmt.Sprintf("[%s]\n", section))
			if err != nil || n == 0 {
				return nil, gerror.Wrapf(err, "w.WriteString failed")
			}
			for k, v := range m {
				if n, err = w.WriteString(fmt.Sprintf("%s=%v\n", k, v)); err != nil || n == 0 {
					return nil, gerror.Wrapf(err, "w.WriteString failed")
				}
			}
			continue
		}
		// 简单的键值对。
		for k, v := range data {
			if n, err = w.WriteString(fmt.Sprintf("%s=%v\n", k, v)); err != nil || n == 0 {
				return nil, gerror.Wrapf(err, "w.WriteString failed")
			}
		}
		break
	}
	res = make([]byte, w.Len())
	if n, err = w.Read(res); err != nil || n == 0 {
		return nil, gerror.Wrapf(err, "w.Read failed")
	}
	return res, nil
}

// ToJson 将 INI 格式转换为 JSON。
func ToJson(data []byte) (res []byte, err error) {
	iniMap, err := Decode(data)
	if err != nil {
		return nil, err
	}
	return json.Marshal(iniMap)
}
