// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gini提供了访问和转换INI内容的功能。 md5:3e0e37cb2af85941
package gini

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/json"
)

// Decode将INI格式转换为映射。 md5:355a2d8ee06f84fe
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

// Encode将映射转换为INI格式。 md5:2b1bb156815e46bd
func Encode(data map[string]interface{}) (res []byte, err error) {
	var (
		n  int
		w  = new(bytes.Buffer)
		m  map[string]interface{}
		ok bool
	)
	for section, item := range data {
				// 部分键值对。 md5:4d0c7048f054d3df
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
				// 简单的键值对。 md5:4ddd5708336bef92
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

// ToJson 将INI格式转换为JSON。 md5:760a6629bda12608
func ToJson(data []byte) (res []byte, err error) {
	iniMap, err := Decode(data)
	if err != nil {
		return nil, err
	}
	return json.Marshal(iniMap)
}
