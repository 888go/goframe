// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gjson

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
// ff:
// j:
func (j Json) MarshalJSON() ([]byte, error) {
	return j.ToJson()
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
// ff:
// j:
// b:
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

// UnmarshalValue is an interface implement which sets any type of value for Json.
// ff:
// j:
// value:
func (j *Json) UnmarshalValue(value interface{}) error {
	if r := NewWithOptions(value, Options{
		StrNumber: true,
	}); r != nil {
		// Value copy.
		*j = *r
	}
	return nil
}

// MapStrAny implements interface function MapStrAny().
// yx:true
// ff:取MapStrAny
// j:
func (j *Json) MapStrAny() map[string]interface{} {
	if j == nil {
		return nil
	}
	return j.Map()
}

// Interfaces implements interface function Interfaces().
// yx:true
// ff:取any切片
// j:
func (j *Json) Interfaces() []interface{} {
	if j == nil {
		return nil
	}
	return j.Array()
}

// String returns current Json object as string.
// ff:
// j:
func (j *Json) String() string {
	if j.IsNil() {
		return ""
	}
	return j.MustToJsonString()
}
