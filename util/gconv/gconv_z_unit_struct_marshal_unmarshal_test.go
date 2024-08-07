// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 转换类_test

import (
	"testing"
	"time"

	gcrc32 "github.com/888go/goframe/crypto/gcrc32"
	gbinary "github.com/888go/goframe/encoding/gbinary"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/frame/g"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gconv "github.com/888go/goframe/util/gconv"
)

type MyTime struct {
	time.Time
}

type MyTimeSt struct {
	ServiceDate MyTime
}

func (st *MyTimeSt) UnmarshalValue(v interface{}) error {
	m := gconv.X取Map(v)
	t, err := gtime.X转换文本(gconv.String(m["ServiceDate"]))
	if err != nil {
		return err
	}
	st.ServiceDate = MyTime{t.Time}
	return nil
}

func Test_Struct_UnmarshalValue1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		st := &MyTimeSt{}
		err := gconv.Struct(g.Map{"ServiceDate": "2020-10-10 12:00:01"}, st)
		t.AssertNil(err)
		t.Assert(st.ServiceDate.Time.Format("2006-01-02 15:04:05"), "2020-10-10 12:00:01")
	})
	gtest.C(t, func(t *gtest.T) {
		st := &MyTimeSt{}
		err := gconv.Struct(g.Map{"ServiceDate": nil}, st)
		t.AssertNil(err)
		t.Assert(st.ServiceDate.Time.IsZero(), true)
	})
	gtest.C(t, func(t *gtest.T) {
		st := &MyTimeSt{}
		err := gconv.Struct(g.Map{"ServiceDate": "error"}, st)
		t.AssertNE(err, nil)
	})
}

type Pkg struct {
	Length uint16 // Total length.
	Crc32  uint32 // CRC32.
	Data   []byte
}

// NewPkg 根据给定的数据创建并返回一个包。 md5:6d3c16ce53d93510
func NewPkg(data []byte) *Pkg {
	return &Pkg{
		Length: uint16(len(data) + 6),
		Crc32:  gcrc32.X加密(data),
		Data:   data,
	}
}

// Marshal 将协议结构体编码为字节。 md5:0d863a50fb456e64
func (p *Pkg) Marshal() []byte {
	b := make([]byte, 6+len(p.Data))
	copy(b, gbinary.EncodeUint16(p.Length))
	copy(b[2:], gbinary.EncodeUint32(p.Crc32))
	copy(b[6:], p.Data)
	return b
}

// UnmarshalValue 将字节解码为协议结构体。 md5:8eadcfee42fbaf11
func (p *Pkg) UnmarshalValue(v interface{}) error {
	b := gconv.X取字节集(v)
	if len(b) < 6 {
		return gerror.X创建("invalid package length")
	}
	p.Length = gbinary.DecodeToUint16(b[:2])
	if len(b) < int(p.Length) {
		return gerror.X创建("invalid data length")
	}
	p.Crc32 = gbinary.DecodeToUint32(b[2:6])
	p.Data = b[6:]
	if gcrc32.X加密(p.Data) != p.Crc32 {
		return gerror.X创建("crc32 validation failed")
	}
	return nil
}

func Test_Struct_UnmarshalValue2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var p1, p2 *Pkg
		p1 = NewPkg([]byte("123"))
		err := gconv.Struct(p1.Marshal(), &p2)
		t.AssertNil(err)
		t.Assert(p1, p2)
	})
}
