package gtime

import (
	"database/sql/driver"
)

// Scan实现了database/sql包中Scan接口，用于从数据库扫描值到本地Golang变量。 md5:1869fcbf2844c5d7
func (t *Time) Scan(value interface{}) error {
	if t == nil {
		return nil
	}
	newTime := New(value)
	*t = *newTime
	return nil
}

// Value是database/sql/driver包中提供Value方法的接口，用于从golang变量获取值并将其插入数据库。 md5:f99b99a521257946
func (t *Time) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	if t.IsZero() {
		return nil, nil
	}
	return t.Time, nil
}
