// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package json类_test

import (
	"testing"

	gjson "github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/frame/g"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
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

	gtest.C(t, func(t *gtest.T) {
		jsonContent := `{"dataSetId":2001,"fieldInfos":{"duration":{"id":80079,"value":"59"},"om_level":{"id":2409,"value":"4"}},"id":"g0936lt1u0f","modifyFieldInfos":{"om_level":{"id":2409,"new":"4","old":""}},"timeStamp":1584599734}`
		var info MediaRequestModifyInfo
		err := gjson.Json格式到变量指针(jsonContent, &info)
		t.AssertNil(err)
		content := gjson.X创建(info).X取json文本PANI()
		t.Assert(gstr.X是否包含(content, `"feed_id":""`), true)
		t.Assert(gstr.X是否包含(content, `"fieldInfos":{`), true)
		t.Assert(gstr.X是否包含(content, `"id":80079`), true)
		t.Assert(gstr.X是否包含(content, `"om_level":{`), true)
		t.Assert(gstr.X是否包含(content, `"id":2409,`), true)
		t.Assert(gstr.X是否包含(content, `"id":"g0936lt1u0f"`), true)
		t.Assert(gstr.X是否包含(content, `"new":"4"`), true)
	})
}

func Test_MapAttributeConvert(t *testing.T) {
	var data = `
 {
   "title": {"l1":"标签1","l2":"标签2"}
}
`
	gtest.C(t, func(t *gtest.T) {
		j, err := gjson.X加载并自动识别格式(data)
		gtest.AssertNil(err)

		tx := struct {
			Title map[string]interface{}
		}{}

		err = j.X取泛型类().X取结构体指针(&tx)
		gtest.AssertNil(err)
		t.Assert(tx.Title, g.Map{
			"l1": "标签1", "l2": "标签2",
		})

		j.X调试输出()

		var nilJ *gjson.Json = nil
		nilJ.X调试输出()
	})

	gtest.C(t, func(t *gtest.T) {
		j, err := gjson.X加载并自动识别格式(data)
		gtest.AssertNil(err)

		tx := struct {
			Title map[string]string
		}{}

		err = j.X取泛型类().X取结构体指针(&tx)
		gtest.AssertNil(err)
		t.Assert(tx.Title, g.Map{
			"l1": "标签1", "l2": "标签2",
		})
	})
}
