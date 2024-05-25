// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package builtin

import (
	"errors"
	"strings"

	"github.com/gogf/gf/v2/internal/empty"
	"github.com/gogf/gf/v2/util/gutil"
)

// RuleRequiredWithAll 实现了 `required-with-all` 规则：
// 如果所有给定的字段都不为空，则此字段为必填。
//
// 格式：required-with-all:field1,field2,...
// 示例：required-with-all:id,name
// md5:1e341cc8965dfdc6
type RuleRequiredWithAll struct{}

func init() {
	Register(RuleRequiredWithAll{})
}

func (r RuleRequiredWithAll) Name() string {
	return "required-with-all"
}

func (r RuleRequiredWithAll) Message() string {
	return "The {field} field is required"
}

func (r RuleRequiredWithAll) Run(in RunInput) error {
	var (
		required   = true
		array      = strings.Split(in.RulePattern, ",")
		foundValue interface{}
		dataMap    = in.Data.Map()
	)

	for i := 0; i < len(array); i++ {
		_, foundValue = gutil.MapPossibleItemByKey(dataMap, array[i])
		if empty.IsEmpty(foundValue) {
			required = false
			break
		}
	}

	if required && isRequiredEmpty(in.Value.Val()) {
		return errors.New(in.Message)
	}
	return nil
}
