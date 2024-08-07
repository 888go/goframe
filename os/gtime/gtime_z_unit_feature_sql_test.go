package 时间类_test

import (
	"testing"

	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

func TestTime_Scan(t1 *testing.T) {
	gtest.C(t1, func(t *gtest.T) {
		tt := gtime.Time{}
		// test string
		s := gtime.X创建并按当前时间().String()
		t.Assert(tt.Scan(s), nil)
		t.Assert(tt.String(), s)
		// test nano
		n := gtime.X取时间戳纳秒()
		t.Assert(tt.Scan(n), nil)
		t.Assert(tt.X取时间戳纳秒(), n)
		// test nil
		none := (*gtime.Time)(nil)
		t.Assert(none.Scan(nil), nil)
		t.Assert(none, nil)
	})

}

func TestTime_Value(t1 *testing.T) {
	gtest.C(t1, func(t *gtest.T) {
		tt := gtime.X创建并按当前时间()
		s, err := tt.Value()
		t.AssertNil(err)
		t.Assert(s, tt.Time)
		// test nil
		none := (*gtime.Time)(nil)
		s, err = none.Value()
		t.AssertNil(err)
		t.Assert(s, nil)

	})
}
