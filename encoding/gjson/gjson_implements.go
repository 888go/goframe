// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gjson

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (j Json) MarshalJSON() ([]byte, error) {
	return j.ToJson()
}

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
func (j *Json) UnmarshalJSON(b []byte) error {
	r, err := loadContentWithOptions(b, Options{
		Type:      ContentTypeJson,
		StrNumber: true,
	})
	if r != nil {
		// Value copy.
		*j = *r
	}
	return err
}

// UnmarshalValue 是一个接口实现，用于为 Json 设置任意类型的值。
func (j *Json) UnmarshalValue(value interface{}) error {
	if r := NewWithOptions(value, Options{
		StrNumber: true,
	}); r != nil {
		// Value copy.
		*j = *r
	}
	return nil
}

// MapStrAny 实现了接口函数 MapStrAny()。
func (j *Json) MapStrAny() map[string]interface{} {
	if j == nil {
		return nil
	}
	return j.Map()
}

// Interfaces 实现了接口函数 Interfaces()。
func (j *Json) Interfaces() []interface{} {
	if j == nil {
		return nil
	}
	return j.Array()
}

// String 将当前Json对象以字符串形式返回。
func (j *Json) String() string {
	if j.IsNil() {
		return ""
	}
	return j.MustToJsonString()
}
