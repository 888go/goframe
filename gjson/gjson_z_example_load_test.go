// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package json类_test

import (
	"fmt"
	
	"github.com/888go/goframe/gjson"
	"github.com/gogf/gf/v2/test/gtest"
)

func ExampleLoad() {
	jsonFilePath := gtest.DataPath("json", "data1.json")
	j, _ := json类.X加载文件(jsonFilePath)
	fmt.Println(j.X取值("name"))
	fmt.Println(j.X取值("score"))

	notExistFilePath := gtest.DataPath("json", "data2.json")
	j2, _ := json类.X加载文件(notExistFilePath)
	fmt.Println(j2.X取值("name"))

	// Output:
	// john
	// 100
}

func ExampleLoadJson() {
	jsonContent := `{"name":"john", "score":"100"}`
	j, _ := json类.X加载json(jsonContent)
	fmt.Println(j.X取值("name"))
	fmt.Println(j.X取值("score"))

	// Output:
	// john
	// 100
}

func ExampleLoadXml() {
	xmlContent := `<?xml version="1.0" encoding="UTF-8"?>
	<base>
		<name>john</name>
		<score>100</score>
	</base>`
	j, _ := json类.X加载xml(xmlContent)
	fmt.Println(j.X取值("base.name"))
	fmt.Println(j.X取值("base.score"))

	// Output:
	// john
	// 100
}

func ExampleLoadIni() {
	iniContent := `
	[base]
	name = john
	score = 100
	`
	j, _ := json类.X加载ini(iniContent)
	fmt.Println(j.X取值("base.name"))
	fmt.Println(j.X取值("base.score"))

	// Output:
	// john
	// 100
}

func ExampleLoadYaml() {
	yamlContent :=
		`base:
  name: john
  score: 100`

	j, _ := json类.X加载Yaml(yamlContent)
	fmt.Println(j.X取值("base.name"))
	fmt.Println(j.X取值("base.score"))

	// Output:
	// john
	// 100
}

func ExampleLoadToml() {
	tomlContent :=
		`[base]
  name = "john"
  score = 100`

	j, _ := json类.X加载Toml(tomlContent)
	fmt.Println(j.X取值("base.name"))
	fmt.Println(j.X取值("base.score"))

	// Output:
	// john
	// 100
}

func ExampleLoadContent() {
	jsonContent := `{"name":"john", "score":"100"}`

	j, _ := json类.X加载并自动识别格式(jsonContent)

	fmt.Println(j.X取值("name"))
	fmt.Println(j.X取值("score"))

	// Output:
	// john
	// 100
}

func ExampleLoadContent_UTF8BOM() {
	jsonContent := `{"name":"john", "score":"100"}`

	content := make([]byte, 3, len(jsonContent)+3)
	content[0] = 0xEF
	content[1] = 0xBB
	content[2] = 0xBF
	content = append(content, jsonContent...)

	j, _ := json类.X加载并自动识别格式(content)

	fmt.Println(j.X取值("name"))
	fmt.Println(j.X取值("score"))

	// Output:
	// john
	// 100
}

func ExampleLoadContent_Xml() {
	xmlContent := `<?xml version="1.0" encoding="UTF-8"?>
	<base>
		<name>john</name>
		<score>100</score>
	</base>`

	x, _ := json类.X加载并自动识别格式(xmlContent)

	fmt.Println(x.X取值("base.name"))
	fmt.Println(x.X取值("base.score"))

	// Output:
	// john
	// 100
}

func ExampleLoadContentType() {
	jsonContent := `{"name":"john", "score":"100"}`
	xmlContent := `<?xml version="1.0" encoding="UTF-8"?>
	<base>
		<name>john</name>
		<score>100</score>
	</base>`

	j, _ := json类.X加载并按格式("json", jsonContent)
	x, _ := json类.X加载并按格式("xml", xmlContent)
	j1, _ := json类.X加载并按格式("json", "")

	fmt.Println(j.X取值("name"))
	fmt.Println(j.X取值("score"))
	fmt.Println(x.X取值("base.name"))
	fmt.Println(x.X取值("base.score"))
	fmt.Println(j1.X取值(""))

	// Output:
	// john
	// 100
	// john
	// 100
}

func ExampleIsValidDataType() {
	fmt.Println(json类.X检查类型("json"))
	fmt.Println(json类.X检查类型("yml"))
	fmt.Println(json类.X检查类型("js"))
	fmt.Println(json类.X检查类型("mp4"))
	fmt.Println(json类.X检查类型("xsl"))
	fmt.Println(json类.X检查类型("txt"))
	fmt.Println(json类.X检查类型(""))
	fmt.Println(json类.X检查类型(".json"))
	fmt.Println(json类.X检查类型(".properties"))

	// Output:
	// true
	// true
	// true
	// false
	// false
	// false
	// false
	// true
	// true
}

func ExampleLoad_Xml() {
	jsonFilePath := gtest.DataPath("xml", "data1.xml")
	j, _ := json类.X加载文件(jsonFilePath)
	fmt.Println(j.X取值("doc.name"))
	fmt.Println(j.X取值("doc.score"))
}

func ExampleLoad_Properties() {
	jsonFilePath := gtest.DataPath("properties", "data1.properties")
	j, _ := json类.X加载文件(jsonFilePath)
	fmt.Println(j.X取值("pr.name"))
	fmt.Println(j.X取值("pr.score"))
	fmt.Println(j.X取值("pr.sex"))

	//Output:
	// john
	// 100
	// 0
}
