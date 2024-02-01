// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtag
import (
	"github.com/888go/goframe/internal/json"
	)
var (
	// 类型名称 => 枚举值的JSON表示。
	enumsMap = make(map[string]json.RawMessage)
)

// SetGlobalEnums 将全局枚举设置到包中。
// 注意，此操作不是并发安全的。
func SetGlobalEnums(enumsJson string) error {
	return json.Unmarshal([]byte(enumsJson), &enumsMap)
}

// GetGlobalEnums 获取并返回全局枚举。
func GetGlobalEnums() (string, error) {
	enumsBytes, err := json.Marshal(enumsMap)
	if err != nil {
		return "", err
	}
	return string(enumsBytes), nil
}

// GetEnumsByType 通过类型名称检索并返回已存储的枚举json。
// 类型名称格式如：github.com/gogf/gf/v2/encoding/gjson.ContentType
func GetEnumsByType(typeName string) string {
	return string(enumsMap[typeName])
}
