// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package clickhouse

import (
	"context"
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	gtime "github.com/888go/goframe/os/gtime"
)

// X底层ConvertValueForField 将值转换为记录字段的类型。 md5:8da3e2d9dc99c3ab
func (d *Driver) X底层ConvertValueForField(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error) {
	switch itemValue := fieldValue.(type) {
	case time.Time:
		// 如果时间是零值，它将更新为nil，
		// 这样在数据库中插入或更新的值将会是"null"。
		// md5:058aebae61025f37
		if itemValue.IsZero() {
			return nil, nil
		}

	case uuid.UUID:
		return itemValue, nil

	case *time.Time:
		// 如果时间是零值，它将更新为nil，
		// 这样在数据库中插入或更新的值将会是"null"。
		// md5:058aebae61025f37
		if itemValue == nil || itemValue.IsZero() {
			return nil, nil
		}
		return itemValue, nil

	case gtime.Time:
		// 如果时间是零值，它将更新为nil，
		// 这样在数据库中插入或更新的值将会是"null"。
		// md5:058aebae61025f37
		if itemValue.IsZero() {
			return nil, nil
		}
						// 用于gtime类型，需要获取time.Time. md5:cbd653f9cf62963a
		return itemValue.Time, nil

	case *gtime.Time:
						// 用于gtime类型，需要获取time.Time. md5:cbd653f9cf62963a
		if itemValue != nil {
			return itemValue.Time, nil
		}
		// 如果时间是零值，它将更新为nil，
		// 这样在数据库中插入或更新的值将会是"null"。
		// md5:058aebae61025f37
		if itemValue == nil || itemValue.IsZero() {
			return nil, nil
		}

	case decimal.Decimal:
		return itemValue, nil

	case *decimal.Decimal:
		if itemValue != nil {
			return *itemValue, nil
		}
		return nil, nil

	default:
		// 如果其他类型实现了driver包的valuer接口
		// 则使用转换后的结果
		// 否则，提交接口数据
		// md5:a04dad650b0b5d2a
		valuer, ok := itemValue.(driver.Valuer)
		if !ok {
			return itemValue, nil
		}
		convertedValue, err := valuer.Value()
		if err != nil {
			return nil, err
		}
		return convertedValue, nil
	}
	return fieldValue, nil
}
