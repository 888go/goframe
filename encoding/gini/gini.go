// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gini 提供了对 INI 内容的访问和转换功能。
package ini类

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
func X取Map(字节集 []byte) (map值 map[string]interface{}, 错误 error) {
	map值 = make(map[string]interface{})
	var (
		fieldMap    = make(map[string]interface{})
		bytesReader = bytes.NewReader(字节集)
		bufioReader = bufio.NewReader(bytesReader)
		section     string
		lastSection string
		haveSection bool
		line        string
	)

	for {
		line, 错误 = bufioReader.ReadString('\n')
		if 错误 != nil {
			if 错误 == io.EOF {
				break
			}
			错误 = 错误类.X多层错误并格式化(错误, `bufioReader.ReadString failed`)
			return nil, 错误
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
			map值[section] = fieldMap
		}
	}

	if !haveSection {
		return nil, 错误类.X创建错误码(错误码类.CodeInvalidParameter, "failed to parse INI file, section not found")
	}
	return map值, nil
}

// Encode 将 map 转换为 INI 格式。
func Map到ini(map值 map[string]interface{}) (字节集 []byte, 错误 error) {
	var (
		n  int
		w  = new(bytes.Buffer)
		m  map[string]interface{}
		ok bool
	)
	for section, item := range map值 {
		// 配置项键值对。
		if m, ok = item.(map[string]interface{}); ok {
			n, 错误 = w.WriteString(fmt.Sprintf("[%s]\n", section))
			if 错误 != nil || n == 0 {
				return nil, 错误类.X多层错误并格式化(错误, "w.WriteString failed")
			}
			for k, v := range m {
				if n, 错误 = w.WriteString(fmt.Sprintf("%s=%v\n", k, v)); 错误 != nil || n == 0 {
					return nil, 错误类.X多层错误并格式化(错误, "w.WriteString failed")
				}
			}
			continue
		}
		// 简单的键值对。
		for k, v := range map值 {
			if n, 错误 = w.WriteString(fmt.Sprintf("%s=%v\n", k, v)); 错误 != nil || n == 0 {
				return nil, 错误类.X多层错误并格式化(错误, "w.WriteString failed")
			}
		}
		break
	}
	字节集 = make([]byte, w.Len())
	if n, 错误 = w.Read(字节集); 错误 != nil || n == 0 {
		return nil, 错误类.X多层错误并格式化(错误, "w.Read failed")
	}
	return 字节集, nil
}

// ToJson 将 INI 格式转换为 JSON。
func X取json(字节集 []byte) (json字节集 []byte, 错误 error) {
	iniMap, 错误 := X取Map(字节集)
	if 错误 != nil {
		return nil, 错误
	}
	return json.Marshal(iniMap)
}
