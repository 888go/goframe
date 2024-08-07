// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 转换类_test

import (
	"encoding/json"
	"testing"
	"time"

	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gconv "github.com/888go/goframe/util/gconv"
)

func TestConverter_ConvertWithRefer(t *testing.T) {
	type tA struct {
		Val int
	}

	type tB struct {
		Val1 int32
		Val2 string
	}

	gtest.C(t, func(t *gtest.T) {
		err := gconv.X转换器注册(func(a tA) (b *tB, err error) {
			b = &tB{
				Val1: int32(a.Val),
				Val2: "abcd",
			}
			return
		})
		t.AssertNil(err)
	})

	gtest.C(t, func(t *gtest.T) {
		a := &tA{
			Val: 1,
		}
		var b tB
		result := gconv.X按参考值类型转换(a, &b)
		t.Assert(result.(*tB), &tB{
			Val1: 1,
			Val2: "abcd",
		})
	})

	gtest.C(t, func(t *gtest.T) {
		a := &tA{
			Val: 1,
		}
		var b tB
		result := gconv.X按参考值类型转换(a, b)
		t.Assert(result.(tB), tB{
			Val1: 1,
			Val2: "abcd",
		})
	})
}

func TestConverter_Struct(t *testing.T) {
	type tA struct {
		Val int
	}

	type tB struct {
		Val1 int32
		Val2 string
	}

	type tAA struct {
		ValTop int
		ValTA  tA
	}

	type tBB struct {
		ValTop int32
		ValTB  tB
	}

	type tCC struct {
		ValTop string
		ValTa  *tB
	}

	type tDD struct {
		ValTop string
		ValTa  tB
	}

	type tEE struct {
		Val1 time.Time  `json:"val1"`
		Val2 *time.Time `json:"val2"`
		Val3 *time.Time `json:"val3"`
	}

	type tFF struct {
		Val1 json.RawMessage            `json:"val1"`
		Val2 []json.RawMessage          `json:"val2"`
		Val3 map[string]json.RawMessage `json:"val3"`
	}

	gtest.C(t, func(t *gtest.T) {
		a := &tA{
			Val: 1,
		}
		var b *tB
		err := gconv.Scan(a, &b)
		t.AssertNil(err)
		t.AssertNE(b, nil)
		t.Assert(b.Val1, 0)
		t.Assert(b.Val2, "")
	})

	gtest.C(t, func(t *gtest.T) {
		err := gconv.X转换器注册(func(a tA) (b *tB, err error) {
			b = &tB{
				Val1: int32(a.Val),
				Val2: "abc",
			}
			return
		})
		t.AssertNil(err)
	})

	gtest.C(t, func(t *gtest.T) {
		a := &tA{
			Val: 1,
		}
		var b *tB
		err := gconv.Scan(a, &b)
		t.AssertNil(err)
		t.AssertNE(b, nil)
		t.Assert(b.Val1, 1)
		t.Assert(b.Val2, "abc")
	})

	gtest.C(t, func(t *gtest.T) {
		a := &tA{
			Val: 1,
		}
		var b *tB
		err := gconv.Scan(a, &b)
		t.AssertNil(err)
		t.AssertNE(b, nil)
		t.Assert(b.Val1, 1)
		t.Assert(b.Val2, "abc")
	})

	gtest.C(t, func(t *gtest.T) {
		a := &tA{
			Val: 1,
		}
		var b *tB
		err := gconv.Scan(a, &b)
		t.AssertNil(err)
		t.AssertNE(b, nil)
		t.Assert(b.Val1, 1)
		t.Assert(b.Val2, "abc")
	})

	gtest.C(t, func(t *gtest.T) {
		a := &tA{
			Val: 1,
		}
		var b *tB
		err := gconv.Scan(a, &b)
		t.AssertNil(err)
		t.AssertNE(b, nil)
		t.Assert(b.Val1, 1)
		t.Assert(b.Val2, "abc")
	})

	gtest.C(t, func(t *gtest.T) {
		aa := &tAA{
			ValTop: 123,
			ValTA:  tA{Val: 234},
		}
		var bb *tBB

		err := gconv.Scan(aa, &bb)
		t.AssertNil(err)
		t.AssertNE(bb, nil)
		t.Assert(bb.ValTop, 123)
		t.AssertNE(bb.ValTB.Val1, 234)

		err = gconv.X转换器注册(func(a tAA) (b *tBB, err error) {
			b = &tBB{
				ValTop: int32(a.ValTop) + 2,
			}
			err = gconv.Scan(a.ValTA, &b.ValTB)
			return
		})
		t.AssertNil(err)

		err = gconv.Scan(aa, &bb)
		t.AssertNil(err)
		t.AssertNE(bb, nil)
		t.Assert(bb.ValTop, 125)
		t.Assert(bb.ValTB.Val1, 234)
		t.Assert(bb.ValTB.Val2, "abc")

	})

	gtest.C(t, func(t *gtest.T) {
		aa := &tAA{
			ValTop: 123,
			ValTA:  tA{Val: 234},
		}
		var cc *tCC
		err := gconv.Scan(aa, &cc)
		t.AssertNil(err)
		t.AssertNE(cc, nil)
		t.Assert(cc.ValTop, "123")
		t.AssertNE(cc.ValTa, nil)
		t.Assert(cc.ValTa.Val1, 234)
		t.Assert(cc.ValTa.Val2, "abc")
	})

	gtest.C(t, func(t *gtest.T) {
		aa := &tAA{
			ValTop: 123,
			ValTA:  tA{Val: 234},
		}

		var dd *tDD
		err := gconv.Scan(aa, &dd)
		t.AssertNil(err)
		t.AssertNE(dd, nil)
		t.Assert(dd.ValTop, "123")
		t.Assert(dd.ValTa.Val1, 234)
		t.Assert(dd.ValTa.Val2, "abc")
	})

			//github.com/gogf/gf/issues/2665. md5:8667bbc7fc630f4b
	gtest.C(t, func(t *gtest.T) {
		aa := &tEE{}

		var tmp = map[string]any{
			"val1": "2023-04-15 19:10:00 +0800 CST",
			"val2": "2023-04-15 19:10:00 +0800 CST",
			"val3": "2006-01-02T15:04:05Z07:00",
		}
		err := gconv.Struct(tmp, aa)
		t.AssertNil(err)
		t.AssertNE(aa, nil)
		t.Assert(aa.Val1.Local(), gtime.X创建("2023-04-15 19:10:00 +0800 CST").X取本地时区().Time)
		t.Assert(aa.Val2.Local(), gtime.X创建("2023-04-15 19:10:00 +0800 CST").X取本地时区().Time)
		t.Assert(aa.Val3.Local(), gtime.X创建("2006-01-02T15:04:05Z07:00").X取本地时区().Time)
	})

			//github.com/gogf/gf/issues/3006. md5:8462c8d1487f0f90
	gtest.C(t, func(t *gtest.T) {
		ff := &tFF{}
		var tmp = map[string]any{
			"val1": map[string]any{"hello": "world"},
			"val2": []any{map[string]string{"hello": "world"}},
			"val3": map[string]map[string]string{"val3": {"hello": "world"}},
		}

		err := gconv.Struct(tmp, ff)
		t.AssertNil(err)
		t.AssertNE(ff, nil)
		t.Assert(ff.Val1, []byte(`{"hello":"world"}`))
		t.AssertEQ(len(ff.Val2), 1)
		t.Assert(ff.Val2[0], []byte(`{"hello":"world"}`))
		t.AssertEQ(len(ff.Val3), 1)
		t.Assert(ff.Val3["val3"], []byte(`{"hello":"world"}`))
	})
}

func TestConverter_CustomBasicType_ToStruct(t *testing.T) {
	type CustomString string
	type CustomStruct struct {
		S string
	}
	gtest.C(t, func(t *gtest.T) {
		var (
			a CustomString = "abc"
			b *CustomStruct
		)
		err := gconv.Scan(a, &b)
		t.AssertNE(err, nil)
		t.Assert(b, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		err := gconv.X转换器注册(func(a CustomString) (b *CustomStruct, err error) {
			b = &CustomStruct{
				S: string(a),
			}
			return
		})
		t.AssertNil(err)
	})
	gtest.C(t, func(t *gtest.T) {
		var (
			a CustomString = "abc"
			b *CustomStruct
		)
		err := gconv.Scan(a, &b)
		t.AssertNil(err)
		t.AssertNE(b, nil)
		t.Assert(b.S, a)
	})
	gtest.C(t, func(t *gtest.T) {
		var (
			a CustomString = "abc"
			b *CustomStruct
		)
		err := gconv.Scan(&a, &b)
		t.AssertNil(err)
		t.AssertNE(b, nil)
		t.Assert(b.S, a)
	})
}

//github.com/gogf/gf/issues/3099. md5:d217d1a2ab119603
func TestConverter_CustomTimeType_ToStruct(t *testing.T) {
	type timestamppb struct {
		S string
	}
	type CustomGTime struct {
		T *gtime.Time
	}
	type CustomPbTime struct {
		T *timestamppb
	}
	gtest.C(t, func(t *gtest.T) {
		var (
			a = CustomGTime{
				T: gtime.X创建并按给定格式文本("2023-10-26", "Y-m-d"),
			}
			b *CustomPbTime
		)
		err := gconv.Scan(a, &b)
		t.AssertNil(err)
		t.AssertNE(b, nil)
		t.Assert(b.T.S, "")
	})

	gtest.C(t, func(t *gtest.T) {
		err := gconv.X转换器注册(func(in gtime.Time) (*timestamppb, error) {
			return &timestamppb{
				S: in.X取本地时区().X取格式文本("Y-m-d"),
			}, nil
		})
		t.AssertNil(err)
		err = gconv.X转换器注册(func(in timestamppb) (*gtime.Time, error) {
			return gtime.X创建并从文本(in.S), nil
		})
		t.AssertNil(err)
	})
	gtest.C(t, func(t *gtest.T) {
		var (
			a = CustomGTime{
				T: gtime.X创建并按给定格式文本("2023-10-26", "Y-m-d"),
			}
			b *CustomPbTime
			c *CustomGTime
		)
		err := gconv.Scan(a, &b)
		t.AssertNil(err)
		t.AssertNE(b, nil)
		t.AssertNE(b.T, nil)

		err = gconv.Scan(b, &c)
		t.AssertNil(err)
		t.AssertNE(c, nil)
		t.AssertNE(c.T, nil)
		t.AssertEQ(a.T.X取时间戳秒(), c.T.X取时间戳秒())
	})
}
