// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gini提供了访问和转换INI内容的功能。 md5:3e0e37cb2af85941
package ini类

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
)

// X取Map将INI格式转换为映射。 md5:355a2d8ee06f84fe
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
			错误 = gerror.X多层错误并格式化(错误, `bufioReader.ReadString failed`)
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
		return nil, gerror.X创建错误码(gcode.CodeInvalidParameter, "failed to parse INI file, section not found")
	}
	return map值, nil
}

// Map到ini将映射转换为INI格式。 md5:2b1bb156815e46bd
func Map到ini(map值 map[string]interface{}) (字节集 []byte, 错误 error) {
	var (
		n  int
		w  = new(bytes.Buffer)
		m  map[string]interface{}
		ok bool
	)
	for section, item := range map值 {
				// 部分键值对。 md5:4d0c7048f054d3df
		if m, ok = item.(map[string]interface{}); ok {
			n, 错误 = w.WriteString(fmt.Sprintf("[%s]\n", section))
			if 错误 != nil || n == 0 {
				return nil, gerror.X多层错误并格式化(错误, "w.WriteString failed")
			}
			for k, v := range m {
				if n, 错误 = w.WriteString(fmt.Sprintf("%s=%v\n", k, v)); 错误 != nil || n == 0 {
					return nil, gerror.X多层错误并格式化(错误, "w.WriteString failed")
				}
			}
			continue
		}
				// 简单的键值对。 md5:4ddd5708336bef92
		for k, v := range map值 {
			if n, 错误 = w.WriteString(fmt.Sprintf("%s=%v\n", k, v)); 错误 != nil || n == 0 {
				return nil, gerror.X多层错误并格式化(错误, "w.WriteString failed")
			}
		}
		break
	}
	字节集 = make([]byte, w.Len())
	if n, 错误 = w.Read(字节集); 错误 != nil || n == 0 {
		return nil, gerror.X多层错误并格式化(错误, "w.Read failed")
	}
	return 字节集, nil
}

// X取json 将INI格式转换为JSON。 md5:760a6629bda12608
func X取json(字节集 []byte) (json字节集 []byte, 错误 error) {
	iniMap, 错误 := X取Map(字节集)
	if 错误 != nil {
		return nil, 错误
	}
	return json.Marshal(iniMap)
}
