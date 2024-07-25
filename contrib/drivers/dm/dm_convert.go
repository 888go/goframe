// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package dm

import (
	"context"

	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

// ConvertValueForField 将值转换为记录字段的类型。 md5:8da3e2d9dc99c3ab
func (d *Driver) ConvertValueForField(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error) {
	switch itemValue := fieldValue.(type) {
		// dm 不支持 time.Time 类型，所以这里将其转换为它支持的时间字符串。 md5:afbc1f9b897fc589
	case time.Time:
		// 如果时间是零值，它将更新为nil，
		// 这样在数据库中插入或更新的值将会是"null"。
		// md5:058aebae61025f37
		if itemValue.IsZero() {
			return nil, nil
		}
		return gtime.New(itemValue).String(), nil

		// dm 不支持 time.Time 类型，所以这里将其转换为它支持的时间字符串。 md5:afbc1f9b897fc589
	case *time.Time:
		// 如果时间是零值，它将更新为nil，
		// 这样在数据库中插入或更新的值将会是"null"。
		// md5:058aebae61025f37
		if itemValue == nil || itemValue.IsZero() {
			return nil, nil
		}
		return gtime.New(itemValue).String(), nil
	}

	return fieldValue, nil
}
