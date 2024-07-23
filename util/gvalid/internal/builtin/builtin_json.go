// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package builtin

import (
	"errors"

	"github.com/gogf/gf/v2/internal/json"
)

// RuleJson 实现了 `json` 规则：
// JSON。
//
// 格式：json
// md5:9f015465c19625e6
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
	if json.Valid(in.Value.Bytes()) {
		return nil
	}
	return errors.New(in.Message)
}
