// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package json类

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (j Json) MarshalJSON() ([]byte, error) {
	return j.X取json字节集()
}

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
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

// UnmarshalValue 是一个接口实现，用于将任何类型的值设置为 Json。 md5:4cedac254443f4c3
func (j *Json) UnmarshalValue(value interface{}) error {
	if r := X创建并按选项(value, Options{
		StrNumber: true,
	}); r != nil {
		// Value copy.
		*j = *r
	}
	return nil
}

// X取MapStrAny 实现了接口函数 X取MapStrAny()。 md5:e07c56a48f2ed248
func (j *Json) X取MapStrAny() map[string]interface{} {
	if j == nil {
		return nil
	}
	return j.X取Map()
}

// X取any切片 实现了接口函数 X取any切片()。 md5:b4dff925202b8b35
func (j *Json) X取any切片() []interface{} {
	if j == nil {
		return nil
	}
	return j.X取切片()
}

// String 返回当前Json对象作为字符串。 md5:741610441066450a
func (j *Json) String() string {
	if j.X是否为Nil() {
		return ""
	}
	return j.X取json文本PANI()
}


// zj:
func (j *Json) X取文本() string { // 字符串返回当前Json对象作为字符串
	return j.String()
}

//zj: