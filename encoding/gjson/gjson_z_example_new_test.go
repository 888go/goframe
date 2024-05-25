// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gjson_test

import (
	"fmt"

	"github.com/gogf/gf/v2/encoding/gjson"
)

func ExampleNew() {
	jsonContent := `{"name":"john", "score":"100"}`
	j := gjson.New(jsonContent)
	fmt.Println(j.Get("name"))
	fmt.Println(j.Get("score"))

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
	j := gjson.NewWithTag(me, "tag", true)
	fmt.Println(j.Get("name"))
	fmt.Println(j.Get("score"))
	fmt.Println(j.Get("Title"))

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

	j := gjson.NewWithOptions(me, gjson.Options{
		Tags: "tag",
	})
	fmt.Println(j.Get("name"))
	fmt.Println(j.Get("score"))
	fmt.Println(j.Get("Title"))

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

	j := gjson.NewWithOptions(content, gjson.Options{
		Tags: "tag",
	})
	fmt.Println(j.Get("name"))
	fmt.Println(j.Get("score"))

	// Output:
	// john
	// 100
}

func ExampleNew_Xml() {
	jsonContent := `<?xml version="1.0" encoding="UTF-8"?><doc><name>john</name><score>100</score></doc>`
	j := gjson.New(jsonContent)
	// 注意XML内容中存在根节点。. md5:8fff88d0927f7c85
	fmt.Println(j.Get("doc.name"))
	fmt.Println(j.Get("doc.score"))
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
	j := gjson.New(me)
	fmt.Println(j.Get("name"))
	fmt.Println(j.Get("score"))
	// Output:
	// john
	// 100
}
