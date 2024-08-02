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

	"github.com/888go/goframe/internal/empty"
	gutil "github.com/888go/goframe/util/gutil"
)

// RuleRequiredWithoutAll 实现了 `required-without-all` 规则：
// 如果所有给定字段都为空，那么这个是必需的。
//
// 格式：  required-without-all:field1,field2,...
// 示例： required-without-all:id,name
// md5:faeb08fb4e97c2b7
type RuleRequiredWithoutAll struct{}

func init() {
	Register(RuleRequiredWithoutAll{})
}

func (r RuleRequiredWithoutAll) Name() string {
	return "required-without-all"
}

func (r RuleRequiredWithoutAll) Message() string {
	return "The {field} field is required"
}

func (r RuleRequiredWithoutAll) Run(in RunInput) error {
	var (
		required   = true
		array      = strings.Split(in.RulePattern, ",")
		foundValue interface{}
		dataMap    = in.Data.Map()
	)

	for i := 0; i < len(array); i++ {
		_, foundValue = gutil.MapPossibleItemByKey(dataMap, array[i])
		if !empty.IsEmpty(foundValue) {
			required = false
			break
		}
	}

	if required && isRequiredEmpty(in.Value.Val()) {
		return errors.New(in.Message)
	}
	return nil
}
