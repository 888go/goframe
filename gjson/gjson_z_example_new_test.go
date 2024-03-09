// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package json类_test

import (
	"fmt"
	
	"github.com/888go/goframe/gjson"
)

func ExampleNew() {
	jsonContent := `{"name":"john", "score":"100"}`
	j := json类.X创建(jsonContent)
	fmt.Println(j.X取值("name"))
	fmt.Println(j.X取值("score"))

	// Output:
	// john
	// 100
}

func ExampleNewWithTag() {
	type Me struct {
		Name  string `tag:"name"`
		Score int    `tag:"score"`
		Title string
	}
	me := Me{
		Name:  "john",
		Score: 100,
		Title: "engineer",
	}
	j := json类.X创建并按类型标签(me, "tag", true)
	fmt.Println(j.X取值("name"))
	fmt.Println(j.X取值("score"))
	fmt.Println(j.X取值("Title"))

	// Output:
	// john
	// 100
	// engineer
}

func ExampleNewWithOptions() {
	type Me struct {
		Name  string `tag:"name"`
		Score int    `tag:"score"`
		Title string
	}
	me := Me{
		Name:  "john",
		Score: 100,
		Title: "engineer",
	}

	j := json类.X创建并按选项(me, json类.Options{
		Tags: "tag",
	})
	fmt.Println(j.X取值("name"))
	fmt.Println(j.X取值("score"))
	fmt.Println(j.X取值("Title"))

	// Output:
	// john
	// 100
	// engineer
}

func ExampleNewWithOptions_UTF8BOM() {
	jsonContent := `{"name":"john", "score":"100"}`

	content := make([]byte, 3, len(jsonContent)+3)
	content[0] = 0xEF
	content[1] = 0xBB
	content[2] = 0xBF
	content = append(content, jsonContent...)

	j := json类.X创建并按选项(content, json类.Options{
		Tags: "tag",
	})
	fmt.Println(j.X取值("name"))
	fmt.Println(j.X取值("score"))

	// Output:
	// john
	// 100
}

func ExampleNew_Xml() {
	jsonContent := `<?xml version="1.0" encoding="UTF-8"?><doc><name>john</name><score>100</score></doc>`
	j := json类.X创建(jsonContent)
	// 注意：在XML内容中存在一个根节点。
	fmt.Println(j.X取值("doc.name"))
	fmt.Println(j.X取值("doc.score"))
	// Output:
	// john
	// 100
}

func ExampleNew_Struct() {
	type Me struct {
		Name  string `json:"name"`
		Score int    `json:"score"`
	}
	me := Me{
		Name:  "john",
		Score: 100,
	}
	j := json类.X创建(me)
	fmt.Println(j.X取值("name"))
	fmt.Println(j.X取值("score"))
	// Output:
	// john
	// 100
}
