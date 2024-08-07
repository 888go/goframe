// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtag

import (
	"github.com/888go/goframe/internal/json"
)

var (
		// 类型名称 => 枚举值的json。 md5:08ddee0638398fd6
	enumsMap = make(map[string]json.RawMessage)
)

// SetGlobalEnums 将全局枚举设置到包中。
// 注意，此操作不具备并发安全性。
// md5:1967d957ac1f393c
func SetGlobalEnums(enumsJson string) error {
	return json.Unmarshal([]byte(enumsJson), &enumsMap)
}

// GetGlobalEnums 获取并返回全局枚举。 md5:9652b5705ec1767f
func GetGlobalEnums() (string, error) {
	enumsBytes, err := json.Marshal(enumsMap)
	if err != nil {
		return "", err
	}
	return string(enumsBytes), nil
}

// GetEnumsByType 根据类型名称检索并返回存储的枚举json。
// 类型名称格式如：github.com/gogf/gf/v2/encoding/gjson.ContentType
// md5:51961ee0c0c68589
func GetEnumsByType(typeName string) string {
	return string(enumsMap[typeName])
}
