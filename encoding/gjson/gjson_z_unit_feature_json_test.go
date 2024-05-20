// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gjson_test

import (
	"testing"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/text/gstr"
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
		err := gjson.DecodeTo(jsonContent, &info)
		t.AssertNil(err)
		content := gjson.New(info).MustToJsonString()
		t.Assert(gstr.Contains(content, `"feed_id":""`), true)
		t.Assert(gstr.Contains(content, `"fieldInfos":{`), true)
		t.Assert(gstr.Contains(content, `"id":80079`), true)
		t.Assert(gstr.Contains(content, `"om_level":{`), true)
		t.Assert(gstr.Contains(content, `"id":2409,`), true)
		t.Assert(gstr.Contains(content, `"id":"g0936lt1u0f"`), true)
		t.Assert(gstr.Contains(content, `"new":"4"`), true)
	})
}

func Test_MapAttributeConvert(t *testing.T) {
	var data = `
 {
   "title": {"l1":"标签1","l2":"标签2"}
}
`
	gtest.C(t, func(t *gtest.T) {
		j, err := gjson.LoadContent(data)
		gtest.AssertNil(err)

		tx := struct {
			Title map[string]interface{}
		}{}

		err = j.Var().Scan(&tx)
		gtest.AssertNil(err)
		t.Assert(tx.Title, g.Map{
			"l1": "标签1", "l2": "标签2",
		})

		j.Dump()

		var nilJ *gjson.Json = nil
		nilJ.Dump()
	})

	gtest.C(t, func(t *gtest.T) {
		j, err := gjson.LoadContent(data)
		gtest.AssertNil(err)

		tx := struct {
			Title map[string]string
		}{}

		err = j.Var().Scan(&tx)
		gtest.AssertNil(err)
		t.Assert(tx.Title, g.Map{
			"l1": "标签1", "l2": "标签2",
		})
	})
}
