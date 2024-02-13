// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类_test

import (
	"testing"
	"time"
	
	"github.com/888go/goframe/crypto/gcrc32"
	"github.com/888go/goframe/encoding/gbinary"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

type MyTime struct {
	time.Time
}

type MyTimeSt struct {
	ServiceDate MyTime
}

func (st *MyTimeSt) UnmarshalValue(v interface{}) error {
	m := 转换类.X取Map(v)
	t, err := 时间类.X转换文本(转换类.String(m["ServiceDate"]))
	if err != nil {
		return err
	}
	st.ServiceDate = MyTime{t.Time}
	return nil
}

func Test_Struct_UnmarshalValue1(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		st := &MyTimeSt{}
		err := 转换类.Struct(g.Map{"ServiceDate": "2020-10-10 12:00:01"}, st)
		t.AssertNil(err)
		t.Assert(st.ServiceDate.Time.Format("2006-01-02 15:04:05"), "2020-10-10 12:00:01")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		st := &MyTimeSt{}
		err := 转换类.Struct(g.Map{"ServiceDate": nil}, st)
		t.AssertNil(err)
		t.Assert(st.ServiceDate.Time.IsZero(), true)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		st := &MyTimeSt{}
		err := 转换类.Struct(g.Map{"ServiceDate": "error"}, st)
		t.AssertNE(err, nil)
	})
}

type Pkg struct {
	Length uint16 // Total length.
	Crc32  uint32 // CRC32.
	Data   []byte
}

// NewPkg 根据给定的数据创建并返回一个包。
func NewPkg(data []byte) *Pkg {
	return &Pkg{
		Length: uint16(len(data) + 6),
		Crc32:  加密crc32类.X加密(data),
		Data:   data,
	}
}

// Marshal 将协议结构体编码为字节流。
func (p *Pkg) Marshal() []byte {
	b := make([]byte, 6+len(p.Data))
	copy(b, 字节集类.EncodeUint16(p.Length))
	copy(b[2:], 字节集类.EncodeUint32(p.Crc32))
	copy(b[6:], p.Data)
	return b
}

// UnmarshalValue 将字节解码为协议结构体。
func (p *Pkg) UnmarshalValue(v interface{}) error {
	b := 转换类.X取字节集(v)
	if len(b) < 6 {
		return 错误类.X创建("invalid package length")
	}
	p.Length = 字节集类.DecodeToUint16(b[:2])
	if len(b) < int(p.Length) {
		return 错误类.X创建("invalid data length")
	}
	p.Crc32 = 字节集类.DecodeToUint32(b[2:6])
	p.Data = b[6:]
	if 加密crc32类.X加密(p.Data) != p.Crc32 {
		return 错误类.X创建("crc32 validation failed")
	}
	return nil
}

func Test_Struct_UnmarshalValue2(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var p1, p2 *Pkg
		p1 = NewPkg([]byte("123"))
		err := 转换类.Struct(p1.Marshal(), &p2)
		t.AssertNil(err)
		t.Assert(p1, p2)
	})
}
