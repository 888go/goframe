// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	
	"github.com/888go/goframe/internal/json"
)

// RuleJson 实现了 `json` 规则：
// JSON。
//
// 格式：json
type RuleJson struct{}

func init() {
	Register(RuleJson{})
}

func (r RuleJson) Name() string {
	return "json"
}

func (r RuleJson) Message() string {
	return "The {field} value `{value}` is not a valid JSON string"
}

func (r RuleJson) Run(in RunInput) error {
	if json.Valid(in.Value.X取字节集()) {
		return nil
	}
	return errors.New(in.Message)
}
