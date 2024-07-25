// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gtime_test

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Json_Pointer(t *testing.T) {
	// Marshal
	gtest.C(t, func(t *gtest.T) {
		type MyTime struct {
			MyTime *gtime.Time
		}
		b, err := json.Marshal(MyTime{
			MyTime: gtime.NewFromStr("2006-01-02 15:04:05"),
		})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":"2006-01-02 15:04:05"}`)
	})
	gtest.C(t, func(t *gtest.T) {
		b, err := json.Marshal(g.Map{
			"MyTime": gtime.NewFromStr("2006-01-02 15:04:05"),
		})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":"2006-01-02 15:04:05"}`)
	})
	gtest.C(t, func(t *gtest.T) {
		b, err := json.Marshal(g.Map{
			"MyTime": *gtime.NewFromStr("2006-01-02 15:04:05"),
		})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":"2006-01-02 15:04:05"}`)
	})
	// Marshal nil
	gtest.C(t, func(t *gtest.T) {
		type MyTime struct {
			MyTime *gtime.Time
		}
		b, err := json.Marshal(&MyTime{})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":null}`)
	})
	// 使用json omitempty标签处理nil值的序列化. md5:fbdcc2e2dbd298f0
	gtest.C(t, func(t *gtest.T) {
		type MyTime struct {
			MyTime *gtime.Time `json:"time,omitempty"`
		}
		b, err := json.Marshal(&MyTime{})
		t.AssertNil(err)
		t.Assert(b, `{}`)
	})
	// Unmarshal
	gtest.C(t, func(t *gtest.T) {
		var (
			myTime gtime.Time
			err    = json.UnmarshalUseNumber([]byte(`"2006-01-02 15:04:05"`), &myTime)
		)
		t.AssertNil(err)
		t.Assert(myTime.String(), "2006-01-02 15:04:05")
	})
}

func Test_Json_Struct(t *testing.T) {
	// Marshal struct.
	gtest.C(t, func(t *gtest.T) {
		type MyTime struct {
			MyTime gtime.Time
		}
		b, err := json.Marshal(MyTime{
			MyTime: *gtime.NewFromStr("2006-01-02 15:04:05"),
		})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":"2006-01-02 15:04:05"}`)
	})
	// Marshal pointer.
	gtest.C(t, func(t *gtest.T) {
		type MyTime struct {
			MyTime gtime.Time
		}
		b, err := json.Marshal(&MyTime{
			MyTime: *gtime.NewFromStr("2006-01-02 15:04:05"),
		})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":"2006-01-02 15:04:05"}`)
	})
	// Marshal nil
	gtest.C(t, func(t *gtest.T) {
		type MyTime struct {
			MyTime gtime.Time
		}
		b, err := json.Marshal(MyTime{})
		t.AssertNil(err)
		t.Assert(b, `{"MyTime":""}`)
	})
	// Marshal nil omitempty
	gtest.C(t, func(t *gtest.T) {
		type MyTime struct {
			MyTime gtime.Time `json:"time,omitempty"`
		}
		b, err := json.Marshal(MyTime{})
		t.AssertNil(err)
		t.Assert(b, `{"time":""}`)
	})

}
