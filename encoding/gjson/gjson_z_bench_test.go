// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package json类_test

import (
	json2 "encoding/json"
	"testing"
	
	"github.com/888go/goframe/encoding/gjson"
)

var (
	jsonStr1 = `{"name":"john","slice":[1,2,3]}`
	jsonStr2 = `{"CallbackCommand":"Group.CallbackAfterSendMsg","From_Account":"61934946","GroupId":"@TGS#2FLGX67FD","MsgBody":[{"MsgContent":{"Text":"是的"},"MsgType":"TIMTextElem"}],"MsgSeq":23,"MsgTime":1567032819,"Operator_Account":"61934946","Random":2804799576,"Type":"Public"}`
	jsonObj1 = json类.X创建(jsonStr1)
	jsonObj2 = json类.X创建(jsonStr2)
)

func Benchmark_Validate_Simple_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json类.X是否为有效json(jsonStr1)
	}
}

func Benchmark_Validate_Complicated_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json类.X是否为有效json(jsonStr2)
	}
}

func Benchmark_Get_Simple_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonObj1.X取值("name")
	}
}

func Benchmark_Get_Complicated_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonObj2.X取值("GroupId")
	}
}

func Benchmark_Stdlib_Json_Unmarshal_Simple_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m map[string]interface{}
		json2.Unmarshal([]byte(jsonStr1), &m)
	}
}

func Benchmark_Stdlib_Json_Unmarshal_Complicated_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m map[string]interface{}
		json2.Unmarshal([]byte(jsonStr2), &m)
	}
}

func Benchmark_New_Simple_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json类.X创建(jsonStr1)
	}
}

func Benchmark_New_Complicated_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json类.X创建(jsonStr2)
	}
}

func Benchmark_Remove_Simple_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonObj1.X删除("name")
	}
}

func Benchmark_Remove_Complicated_Json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonObj2.X删除("GroupId")
	}
}

func Benchmark_New_Nil_And_Set_Simple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := json类.X创建(nil)
		p.X设置值("k", "v")
	}
}

func Benchmark_New_Nil_And_Set_Multiple_Level(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := json类.X创建(nil)
		p.X设置值("0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0", []int{1, 2, 3})
	}
}
