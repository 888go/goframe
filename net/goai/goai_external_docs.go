// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package goai

import (
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// ExternalDocs 是由 OpenAPI/Swagger 标准版本 3.0 定义的。 md5:351fbd5301af5e56
type ExternalDocs struct {
	URL         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
}

func (ed *ExternalDocs) UnmarshalValue(value interface{}) error {
	var valueBytes = gconv.Bytes(value)
	if json.Valid(valueBytes) {
		return json.UnmarshalUseNumber(valueBytes, ed)
	}
	var (
		valueString = string(valueBytes)
		valueArray  = gstr.Split(valueString, "|")
	)
	ed.URL = valueArray[0]
	if len(valueArray) > 1 {
		ed.Description = valueArray[1]
	}
	return nil
}
