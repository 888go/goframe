// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 时间类_test

import (
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
)

func Test_Json_Pointer(t *testing.T) {
	// Marshal
	单元测试类.C(t, func(t *单元测试类.T) {
		type MyTime struct {
			MyTime *时间类.Time
		}
		b, err := json.Marshal(MyTime{
			MyTime: 时间类.X创建并从文本("2006-01-02 15:04:05"),
		})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":"2006-01-02 15:04:05"}`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		b, err := json.Marshal(g.Map{
			"MyTime": 时间类.X创建并从文本("2006-01-02 15:04:05"),
		})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":"2006-01-02 15:04:05"}`)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		b, err := json.Marshal(g.Map{
			"MyTime": *时间类.X创建并从文本("2006-01-02 15:04:05"),
		})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":"2006-01-02 15:04:05"}`)
	})
	// Marshal nil
	单元测试类.C(t, func(t *单元测试类.T) {
		type MyTime struct {
			MyTime *时间类.Time
		}
		b, err := json.Marshal(&MyTime{})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":null}`)
	})
	// 使用json标签omitempty对nil进行序列化
	单元测试类.C(t, func(t *单元测试类.T) {
		type MyTime struct {
			MyTime *时间类.Time `json:"time,omitempty"`
		}
		b, err := json.Marshal(&MyTime{})
		t.AssertNil(err)
		t.Assert(b, `{}`)
	})
	// Unmarshal
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			myTime 时间类.Time
			err    = json.UnmarshalUseNumber([]byte(`"2006-01-02 15:04:05"`), &myTime)
		)
		t.AssertNil(err)
		t.Assert(myTime.String(), "2006-01-02 15:04:05")
	})
}

func Test_Json_Struct(t *testing.T) {
	// Marshal struct.
	单元测试类.C(t, func(t *单元测试类.T) {
		type MyTime struct {
			MyTime 时间类.Time
		}
		b, err := json.Marshal(MyTime{
			MyTime: *时间类.X创建并从文本("2006-01-02 15:04:05"),
		})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":"2006-01-02 15:04:05"}`)
	})
	// Marshal pointer.
	单元测试类.C(t, func(t *单元测试类.T) {
		type MyTime struct {
			MyTime 时间类.Time
		}
		b, err := json.Marshal(&MyTime{
			MyTime: *时间类.X创建并从文本("2006-01-02 15:04:05"),
		})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":"2006-01-02 15:04:05"}`)
	})
	// Marshal nil
	单元测试类.C(t, func(t *单元测试类.T) {
		type MyTime struct {
			MyTime 时间类.Time
		}
		b, err := json.Marshal(MyTime{})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":""}`)
	})
	// 对空值进行序列化时忽略（omitempty）
	单元测试类.C(t, func(t *单元测试类.T) {
		type MyTime struct {
			MyTime 时间类.Time `json:"time,omitempty"`
		}
		b, err := json.Marshal(MyTime{})
		t.AssertNil(err)
		t.Assert(b, `{"time":""}`)
	})

}
