// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 泛型类_test

import (
	"bytes"
	"encoding/binary"
	"testing"
	"time"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_Set(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var v 泛型类.Var
		v.X设置值(123.456)
		t.Assert(v.X取值(), 123.456)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var v 泛型类.Var
		v.X设置值(123.456)
		t.Assert(v.X取值(), 123.456)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		objOne := 泛型类.X创建("old", true)
		objOneOld, _ := objOne.X设置值("new").(string)
		t.Assert(objOneOld, "old")

		objTwo := 泛型类.X创建("old", false)
		objTwoOld, _ := objTwo.X设置值("new").(string)
		t.Assert(objTwoOld, "old")
	})
}

func Test_Val(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		objOne := 泛型类.X创建(1, true)
		objOneOld, _ := objOne.X取值().(int)
		t.Assert(objOneOld, 1)

		objTwo := 泛型类.X创建(1, false)
		objTwoOld, _ := objTwo.X取值().(int)
		t.Assert(objTwoOld, 1)

		objOne = nil
		t.Assert(objOne.X取值(), nil)
	})
}

func Test_Interface(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		objOne := 泛型类.X创建(1, true)
		objOneOld, _ := objOne.Interface().(int)
		t.Assert(objOneOld, 1)

		objTwo := 泛型类.X创建(1, false)
		objTwoOld, _ := objTwo.Interface().(int)
		t.Assert(objTwoOld, 1)
	})
}

func Test_IsNil(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		objOne := 泛型类.X创建(nil, true)
		t.Assert(objOne.X是否为Nil(), true)

		objTwo := 泛型类.X创建("noNil", false)
		t.Assert(objTwo.X是否为Nil(), false)

	})
}

func Test_Bytes(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		x := int32(1)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, x)

		objOne := 泛型类.X创建(bytesBuffer.Bytes(), true)

		bBuf := bytes.NewBuffer(objOne.X取字节集())
		var y int32
		binary.Read(bBuf, binary.BigEndian, &y)

		t.Assert(x, y)

	})
}

func Test_String(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var str string = "hello"
		objOne := 泛型类.X创建(str, true)
		t.Assert(objOne.String(), str)

	})
}

func Test_Bool(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var ok bool = true
		objOne := 泛型类.X创建(ok, true)
		t.Assert(objOne.X取布尔(), ok)

		ok = false
		objTwo := 泛型类.X创建(ok, true)
		t.Assert(objTwo.X取布尔(), ok)

	})
}

func Test_Int(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var num int = 1
		objOne := 泛型类.X创建(num, true)
		t.Assert(objOne.X取整数(), num)

	})
}

func Test_Int8(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var num int8 = 1
		objOne := 泛型类.X创建(num, true)
		t.Assert(objOne.X取整数8位(), num)

	})
}

func Test_Int16(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var num int16 = 1
		objOne := 泛型类.X创建(num, true)
		t.Assert(objOne.X取整数16位(), num)

	})
}

func Test_Int32(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var num int32 = 1
		objOne := 泛型类.X创建(num, true)
		t.Assert(objOne.X取整数32位(), num)

	})
}

func Test_Int64(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var num int64 = 1
		objOne := 泛型类.X创建(num, true)
		t.Assert(objOne.X取整数64位(), num)

	})
}

func Test_Uint(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var num uint = 1
		objOne := 泛型类.X创建(num, true)
		t.Assert(objOne.X取正整数(), num)

	})
}

func Test_Uint8(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var num uint8 = 1
		objOne := 泛型类.X创建(num, true)
		t.Assert(objOne.X取正整数8位(), num)

	})
}

func Test_Uint16(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var num uint16 = 1
		objOne := 泛型类.X创建(num, true)
		t.Assert(objOne.X取正整数16位(), num)

	})
}

func Test_Uint32(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var num uint32 = 1
		objOne := 泛型类.X创建(num, true)
		t.Assert(objOne.X取正整数32位(), num)

	})
}

func Test_Uint64(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var num uint64 = 1
		objOne := 泛型类.X创建(num, true)
		t.Assert(objOne.X取正整数64位(), num)

	})
}

func Test_Float32(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var num float32 = 1.1
		objOne := 泛型类.X创建(num, true)
		t.Assert(objOne.X取小数32位(), num)

	})
}

func Test_Float64(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var num float64 = 1.1
		objOne := 泛型类.X创建(num, true)
		t.Assert(objOne.X取小数64位(), num)

	})
}

func Test_Time(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var timeUnix int64 = 1556242660
		objOne := 泛型类.X创建(timeUnix, true)
		t.Assert(objOne.X取时间类().Unix(), timeUnix)
	})
}

func Test_GTime(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var timeUnix int64 = 1556242660
		objOne := 泛型类.X创建(timeUnix, true)
		t.Assert(objOne.X取gtime时间类().Unix(), timeUnix)
	})
}

func Test_Duration(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var timeUnix int64 = 1556242660
		objOne := 泛型类.X创建(timeUnix, true)
		t.Assert(objOne.X取时长(), time.Duration(timeUnix))
	})
}

func Test_UnmarshalJson(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type V struct {
			Name string
			Var  *泛型类.Var
		}
		var v *V
		err := 转换类.Struct(map[string]interface{}{
			"name": "john",
			"var":  "v",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.String(), "v")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type V struct {
			Name string
			Var  泛型类.Var
		}
		var v *V
		err := 转换类.Struct(map[string]interface{}{
			"name": "john",
			"var":  "v",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.String(), "v")
	})
}

func Test_UnmarshalValue(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type V struct {
			Name string
			Var  *泛型类.Var
		}
		var v *V
		err := 转换类.Struct(map[string]interface{}{
			"name": "john",
			"var":  "v",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.String(), "v")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type V struct {
			Name string
			Var  泛型类.Var
		}
		var v *V
		err := 转换类.Struct(map[string]interface{}{
			"name": "john",
			"var":  "v",
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Var.String(), "v")
	})
}

func Test_Copy(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		src := g.Map{
			"k1": "v1",
			"k2": "v2",
		}
		srcVar := 泛型类.X创建(src)
		dstVar := srcVar.X深拷贝()
		t.Assert(srcVar.X取Map(), src)
		t.Assert(dstVar.X取Map(), src)

		dstVar.X取Map()["k3"] = "v3"
		t.Assert(srcVar.X取Map(), g.Map{
			"k1": "v1",
			"k2": "v2",
		})
		t.Assert(dstVar.X取Map(), g.Map{
			"k1": "v1",
			"k2": "v2",
			"k3": "v3",
		})
	})
}

func Test_DeepCopy(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		src := g.Map{
			"k1": "v1",
			"k2": "v2",
		}
		srcVar := 泛型类.X创建(src)
		copyVar := srcVar.DeepCopy().(*泛型类.Var)
		copyVar.X设置值(g.Map{
			"k3": "v3",
			"k4": "v4",
		})
		t.AssertNE(srcVar, copyVar)

		srcVar = nil
		t.AssertNil(srcVar.DeepCopy())
	})
}
