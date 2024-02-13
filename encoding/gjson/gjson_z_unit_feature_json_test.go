// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package json类_test

import (
	"testing"
	
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
)

func Test_ToJson(t *testing.T) {
	type ModifyFieldInfoType struct {
		Id  int64  `json:"id"`
		New string `json:"new"`
	}
	type ModifyFieldInfosType struct {
		Duration ModifyFieldInfoType `json:"duration"`
		OMLevel  ModifyFieldInfoType `json:"om_level"`
	}

	type MediaRequestModifyInfo struct {
		Modify ModifyFieldInfosType `json:"modifyFieldInfos"`
		Field  ModifyFieldInfosType `json:"fieldInfos"`
		FeedID string               `json:"feed_id"`
		Vid    string               `json:"id"`
	}

	单元测试类.C(t, func(t *单元测试类.T) {
		jsonContent := `{"dataSetId":2001,"fieldInfos":{"duration":{"id":80079,"value":"59"},"om_level":{"id":2409,"value":"4"}},"id":"g0936lt1u0f","modifyFieldInfos":{"om_level":{"id":2409,"new":"4","old":""}},"timeStamp":1584599734}`
		var info MediaRequestModifyInfo
		err := json类.Json格式到变量指针(jsonContent, &info)
		t.AssertNil(err)
		content := json类.X创建(info).X取json文本PANI()
		t.Assert(文本类.X是否包含(content, `"feed_id":""`), true)
		t.Assert(文本类.X是否包含(content, `"fieldInfos":{`), true)
		t.Assert(文本类.X是否包含(content, `"id":80079`), true)
		t.Assert(文本类.X是否包含(content, `"om_level":{`), true)
		t.Assert(文本类.X是否包含(content, `"id":2409,`), true)
		t.Assert(文本类.X是否包含(content, `"id":"g0936lt1u0f"`), true)
		t.Assert(文本类.X是否包含(content, `"new":"4"`), true)
	})
}

func Test_MapAttributeConvert(t *testing.T) {
	var data = `
 {
   "title": {"l1":"标签1","l2":"标签2"}
}
`
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X加载并自动识别格式(data)
		单元测试类.AssertNil(err)

		tx := struct {
			Title map[string]interface{}
		}{}

		err = j.X取泛型类().X取结构体指针(&tx)
		单元测试类.AssertNil(err)
		t.Assert(tx.Title, g.Map{
			"l1": "标签1", "l2": "标签2",
		})

		j.X调试输出()

		var nilJ *json类.Json = nil
		nilJ.X调试输出()
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X加载并自动识别格式(data)
		单元测试类.AssertNil(err)

		tx := struct {
			Title map[string]string
		}{}

		err = j.X取泛型类().X取结构体指针(&tx)
		单元测试类.AssertNil(err)
		t.Assert(tx.Title, g.Map{
			"l1": "标签1", "l2": "标签2",
		})
	})
}
