package gtime
import (
	"database/sql/driver"
	)
// Scan 实现了由 database/sql 包中 Scan 方法使用的接口，用于将数据库中的值扫描到本地 Go 语言变量中。
func (t *Time) Scan(value interface{}) error {
	if t == nil {
		return nil
	}
	newTime := New(value)
	*t = *newTime
	return nil
}

// Value 是一个接口，为 package database/sql/driver 提供 Value 方法，用于从 Go 语言变量向数据库获取值。
func (t *Time) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	if t.IsZero() {
		return nil, nil
	}
	return t.Time, nil
}
