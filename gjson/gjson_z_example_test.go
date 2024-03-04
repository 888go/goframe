package gjson_test

import (
	"fmt"
	
	"github.com/888go/goframe/gjson"
)

func ExampleJson_SetSplitChar() {
	data :=
		`{
        "users" : {
            "count" : 2,
            "list"  : [
                {"name" : "Ming",  "score" : 60},
                {"name" : "John", "score" : 99.5}
            ]
        }
    }`
	if j, err := gjson.DecodeToJson(data); err != nil {
		panic(err)
	} else {
		j.SetSplitChar('#')
		fmt.Println("John Score:", j.Get("users#list#1#score").Float32())
	}
	// Output:
	// John Score: 99.5
}

func ExampleJson_SetViolenceCheck() {
	data :=
		`{
        "users" : {
            "count" : 100
        },
        "users.count" : 101
    }`
	if j, err := gjson.DecodeToJson(data); err != nil {
		fmt.Println(err)
	} else {
		j.SetViolenceCheck(false)
		fmt.Println("Users Count:", j.Get("users.count"))
		j.SetViolenceCheck(true)
		fmt.Println("Users Count:", j.Get("users.count"))
	}
	// Output:
	// Users Count: 100
	// Users Count: 101
}

// ===========================================================================
// JSON
// ===========================================================================
// （此部分代码注释为标题性描述，翻译如下）
// ============================================================================
// JSON 部分
// ============================================================================
// （这里表示该段代码与 JSON 相关，可能是用于处理、解析或生成 JSON 的功能模块）
func ExampleJson_ToJson() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	jsonBytes, _ := j.ToJson()
	fmt.Println(string(jsonBytes))

	// Output:
	// {"Age":18,"Name":"John"}
}

func ExampleJson_ToJsonString() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	jsonStr, _ := j.ToJsonString()
	fmt.Println(jsonStr)

	// Output:
	// {"Age":18,"Name":"John"}
}

func ExampleJson_ToJsonIndent() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	jsonBytes, _ := j.ToJsonIndent()
	fmt.Println(string(jsonBytes))

	// Output:
	// {
	//	"Age": 18,
	//	"Name": "John"
	// }
}

func ExampleJson_ToJsonIndentString() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	jsonStr, _ := j.ToJsonIndentString()
	fmt.Println(jsonStr)

	// Output:
	// {
	//	"Age": 18,
	//	"Name": "John"
	// }
}

func ExampleJson_MustToJson() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	jsonBytes := j.MustToJson()
	fmt.Println(string(jsonBytes))

	// Output:
	// {"Age":18,"Name":"John"}
}

func ExampleJson_MustToJsonString() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	jsonStr := j.MustToJsonString()
	fmt.Println(jsonStr)

	// Output:
	// {"Age":18,"Name":"John"}
}

func ExampleJson_MustToJsonIndent() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	jsonBytes := j.MustToJsonIndent()
	fmt.Println(string(jsonBytes))

	// Output:
	// {
	//	"Age": 18,
	//	"Name": "John"
	// }
}

func ExampleJson_MustToJsonIndentString() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	jsonStr := j.MustToJsonIndentString()
	fmt.Println(jsonStr)

	// Output:
	// {
	//	"Age": 18,
	//	"Name": "John"
	// }
}

// ========================================================================
// XML
// ========================================================================
// ===中文注释开始===
// 此部分代码与XML相关功能实现
// ========================================================================
func ExampleJson_ToXml() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	xmlBytes, _ := j.ToXml()
	fmt.Println(string(xmlBytes))

	// Output:
	// <doc><Age>18</Age><Name>John</Name></doc>
}

func ExampleJson_ToXmlString() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	xmlStr, _ := j.ToXmlString()
	fmt.Println(string(xmlStr))

	// Output:
	// <doc><Age>18</Age><Name>John</Name></doc>
}

func ExampleJson_ToXmlIndent() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	xmlBytes, _ := j.ToXmlIndent()
	fmt.Println(string(xmlBytes))

	// Output:
	// <doc>
	//	<Age>18</Age>
	//	<Name>John</Name>
	// </doc>
}

func ExampleJson_ToXmlIndentString() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	xmlStr, _ := j.ToXmlIndentString()
	fmt.Println(string(xmlStr))

	// Output:
	// <doc>
	//	<Age>18</Age>
	//	<Name>John</Name>
	// </doc>
}

func ExampleJson_MustToXml() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	xmlBytes := j.MustToXml()
	fmt.Println(string(xmlBytes))

	// Output:
	// <doc><Age>18</Age><Name>John</Name></doc>
}

func ExampleJson_MustToXmlString() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	xmlStr := j.MustToXmlString()
	fmt.Println(string(xmlStr))

	// Output:
	// <doc><Age>18</Age><Name>John</Name></doc>
}

func ExampleJson_MustToXmlIndent() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	xmlBytes := j.MustToXmlIndent()
	fmt.Println(string(xmlBytes))

	// Output:
	// <doc>
	//	<Age>18</Age>
	//	<Name>John</Name>
	// </doc>
}

func ExampleJson_MustToXmlIndentString() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	xmlStr := j.MustToXmlIndentString()
	fmt.Println(string(xmlStr))

	// Output:
	// <doc>
	//	<Age>18</Age>
	//	<Name>John</Name>
	// </doc>
}

// ==========================================================================
// YAML
// ==========================================================================
// （这段代码的注释表明了该部分代码与YAML相关，但没有给出具体的功能描述，故无法精确翻译更多内容。）
// ========================================================================
// 表示分割线，用于组织和划分代码区域
// YAML
// 这行注释表示这部分代码与YAML（一种数据序列化格式）有关
func ExampleJson_ToYaml() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	YamlBytes, _ := j.ToYaml()
	fmt.Println(string(YamlBytes))

	// Output:
	// Age: 18
	// Name: John
}

func ExampleJson_ToYamlString() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	YamlStr, _ := j.ToYamlString()
	fmt.Println(string(YamlStr))

	// Output:
	// Age: 18
	// Name: John
}

func ExampleJson_ToYamlIndent() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	YamlBytes, _ := j.ToYamlIndent("")
	fmt.Println(string(YamlBytes))

	// Output:
	// Age: 18
	// Name: John
}

func ExampleJson_MustToYaml() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	YamlBytes := j.MustToYaml()
	fmt.Println(string(YamlBytes))

	// Output:
	// Age: 18
	// Name: John
}

func ExampleJson_MustToYamlString() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	YamlStr := j.MustToYamlString()
	fmt.Println(string(YamlStr))

	// Output:
	// Age: 18
	// Name: John
}

// ===========================================================================
// TOML
// ===========================================================================
// （注释翻译：）
// ============================================================================
// TOML
// ============================================================================
// 这段代码的注释表明了该部分与TOML相关。TOML是一种配置文件格式，因此这部分可能是用于解析、生成或操作TOML格式数据的Go语言代码段。此处的“//”表示单行注释，中文翻译后含义不变，仅将原有的英文内容转换为中文展示。
func ExampleJson_ToToml() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	TomlBytes, _ := j.ToToml()
	fmt.Println(string(TomlBytes))

	// Output:
	// Age = 18
	// Name = "John"
}

func ExampleJson_ToTomlString() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	TomlStr, _ := j.ToTomlString()
	fmt.Println(string(TomlStr))

	// Output:
	// Age = 18
	// Name = "John"
}

func ExampleJson_MustToToml() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	TomlBytes := j.MustToToml()
	fmt.Println(string(TomlBytes))

	// Output:
	// Age = 18
	// Name = "John"
}

func ExampleJson_MustToTomlString() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	TomlStr := j.MustToTomlString()
	fmt.Println(string(TomlStr))

	// Output:
	// Age = 18
	// Name = "John"
}

// ===========================================================================
// INI
// ===========================================================================
// （此部分代码的注释表明该部分与INI文件相关，但没有提供详细信息，故直接翻译为中文即可）
// ===========================================================================
// INI配置文件相关代码
// ===========================================================================
func ExampleJson_ToIni() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	IniBytes, _ := j.ToIni()
	fmt.Println(string(IniBytes))

	// May Output:
	// Name=John
	// Age=18
}

func ExampleJson_ToIniString() {
	type BaseInfo struct {
		Name string
	}

	info := BaseInfo{
		Name: "John",
	}

	j := gjson.New(info)
	IniStr, _ := j.ToIniString()
	fmt.Println(string(IniStr))

	// Output:
	// Name=John
}

func ExampleJson_MustToIni() {
	type BaseInfo struct {
		Name string
	}

	info := BaseInfo{
		Name: "John",
	}

	j := gjson.New(info)
	IniBytes := j.MustToIni()
	fmt.Println(string(IniBytes))

	// Output:
	// Name=John
}

func ExampleJson_MustToIniString() {
	type BaseInfo struct {
		Name string
	}

	info := BaseInfo{
		Name: "John",
	}

	j := gjson.New(info)
	IniStr := j.MustToIniString()
	fmt.Println(string(IniStr))

	// Output:
	// Name=John
}

// ========================================================================
// 属性
// ========================================================================
// 这段 Go 语言代码的注释表明该部分是关于“属性”的，翻译成中文后即为：
// ========================================================================
// 属性
// ========================================================================
// 这里的“Properties”在编程中通常指的是程序、类或对象所具有的各种配置项、特性或变量等。
func ExampleJson_ToProperties() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	pr, _ := j.ToProperties()
	fmt.Println(string(pr))

	// May Output:
	// name = John
	// age = 18
}

func ExampleJson_ToPropertiesString() {
	type BaseInfo struct {
		Name string
	}

	info := BaseInfo{
		Name: "John",
	}

	j := gjson.New(info)
	pr, _ := j.ToPropertiesString()
	fmt.Println(pr)

	// Output:
	// name = John
}

func ExampleJson_MustToProperties() {
	type BaseInfo struct {
		Name string
	}

	info := BaseInfo{
		Name: "John",
	}

	j := gjson.New(info)
	pr := j.MustToProperties()
	fmt.Println(string(pr))

	// Output:
	// name = John
}

func ExampleJson_MustToPropertiesString() {
	type BaseInfo struct {
		Name string
	}

	info := BaseInfo{
		Name: "John",
	}

	j := gjson.New(info)
	pr := j.MustToPropertiesString()
	fmt.Println(pr)

	// Output:
	// name = John
}

func ExampleJson_MarshalJSON() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	jsonBytes, _ := j.MarshalJSON()
	fmt.Println(string(jsonBytes))

	// Output:
	// {"Age":18,"Name":"John"}
}

func ExampleJson_UnmarshalJSON() {
	jsonStr := `{"Age":18,"Name":"John"}`

	j := gjson.New("")
	j.UnmarshalJSON([]byte(jsonStr))
	fmt.Println(j.Map())

	// Output:
	// map[Age:18 Name:John]
}

func ExampleJson_UnmarshalValue_Yaml() {
	yamlContent :=
		`base:
  name: john
  score: 100`

	j := gjson.New("")
	j.UnmarshalValue([]byte(yamlContent))
	fmt.Println(j.Var().String())

	// Output:
	// {"base":{"name":"john","score":100}}
}

func ExampleJson_UnmarshalValue_Xml() {
	xmlStr := `<?xml version="1.0" encoding="UTF-8"?><doc><name>john</name><score>100</score></doc>`

	j := gjson.New("")
	j.UnmarshalValue([]byte(xmlStr))
	fmt.Println(j.Var().String())

	// Output:
	// {"doc":{"name":"john","score":"100"}}
}

func ExampleJson_MapStrAny() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	fmt.Println(j.MapStrAny())

	// Output:
	// map[Age:18 Name:John]
}

func ExampleJson_Interfaces() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	infoList := []BaseInfo{
		{
			Name: "John",
			Age:  18,
		},
		{
			Name: "Tom",
			Age:  20,
		},
	}

	j := gjson.New(infoList)
	fmt.Println(j.Interfaces())

	// Output:
	// [{John 18} {Tom 20}]
}

func ExampleJson_Interface() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	fmt.Println(j.Interface())

	var nilJ *gjson.Json = nil
	fmt.Println(nilJ.Interface())

	// Output:
	// map[Age:18 Name:John]
	// <nil>
}

func ExampleJson_Var() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	fmt.Println(j.Var().String())
	fmt.Println(j.Var().Map())

	// Output:
	// {"Age":18,"Name":"John"}
	// map[Age:18 Name:John]
}

func ExampleJson_IsNil() {
	data1 := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	data2 := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]`)

	j1, _ := gjson.LoadContent(data1)
	fmt.Println(j1.IsNil())

	j2, _ := gjson.LoadContent(data2)
	fmt.Println(j2.IsNil())

	// Output:
	// false
	// true
}

func ExampleJson_Get() {
	data :=
		`{
        "users" : {
            "count" : 1,
            "array" : ["John", "Ming"]
        }
    }`

	j, _ := gjson.LoadContent(data)
	fmt.Println(j.Get("."))
	fmt.Println(j.Get("users"))
	fmt.Println(j.Get("users.count"))
	fmt.Println(j.Get("users.array"))

	var nilJ *gjson.Json = nil
	fmt.Println(nilJ.Get("."))

	// Output:
	// {"users":{"array":["John","Ming"],"count":1}}
	// {"array":["John","Ming"],"count":1}
	// 1
	// ["John","Ming"]
}

func ExampleJson_GetJson() {
	data :=
		`{
        "users" : {
            "count" : 1,
            "array" : ["John", "Ming"]
        }
    }`

	j, _ := gjson.LoadContent(data)

	fmt.Println(j.GetJson("users.array").Array())

	// Output:
	// [John Ming]
}

func ExampleJson_GetJsons() {
	data :=
		`{
        "users" : {
            "count" : 3,
            "array" : [{"Age":18,"Name":"John"}, {"Age":20,"Name":"Tom"}]
        }
    }`

	j, _ := gjson.LoadContent(data)

	jsons := j.GetJsons("users.array")
	for _, json := range jsons {
		fmt.Println(json.Interface())
	}

	// Output:
	// map[Age:18 Name:John]
	// map[Age:20 Name:Tom]
}

func ExampleJson_GetJsonMap() {
	data :=
		`{
        "users" : {
            "count" : 1,
			"array" : {
				"info" : {"Age":18,"Name":"John"},
				"addr" : {"City":"Chengdu","Company":"Tencent"}
			}
        }
    }`

	j, _ := gjson.LoadContent(data)

	jsonMap := j.GetJsonMap("users.array")

	for _, json := range jsonMap {
		fmt.Println(json.Interface())
	}

	// May Output:
	// map[City:Chengdu Company:Tencent]
	// map[Age:18 Name:John]
}

func ExampleJson_Set() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	j.Set("Addr", "ChengDu")
	j.Set("Friends.0", "Tom")
	fmt.Println(j.Var().String())

	// Output:
	// {"Addr":"ChengDu","Age":18,"Friends":["Tom"],"Name":"John"}
}

func ExampleJson_MustSet() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	j.MustSet("Addr", "ChengDu")
	fmt.Println(j.Var().String())

	// Output:
	// {"Addr":"ChengDu","Age":18,"Name":"John"}
}

func ExampleJson_Remove() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	j.Remove("Age")
	fmt.Println(j.Var().String())

	// Output:
	// {"Name":"John"}
}

func ExampleJson_MustRemove() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	j.MustRemove("Age")
	fmt.Println(j.Var().String())

	// Output:
	// {"Name":"John"}
}

func ExampleJson_Contains() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.New(info)
	fmt.Println(j.Contains("Age"))
	fmt.Println(j.Contains("Addr"))

	// Output:
	// true
	// false
}

func ExampleJson_Len() {
	data :=
		`{
        "users" : {
            "count" : 1,
			"nameArray" : ["Join", "Tom"],
			"infoMap" : {
				"name" : "Join",
				"age" : 18,
				"addr" : "ChengDu"
			}
        }
    }`

	j, _ := gjson.LoadContent(data)

	fmt.Println(j.Len("users.nameArray"))
	fmt.Println(j.Len("users.infoMap"))

	// Output:
	// 2
	// 3
}

func ExampleJson_Append() {
	data :=
		`{
        "users" : {
            "count" : 1,
            "array" : ["John", "Ming"]
        }
    }`

	j, _ := gjson.LoadContent(data)

	j.Append("users.array", "Lily")

	fmt.Println(j.Get("users.array").Array())

	// Output:
	// [John Ming Lily]
}

func ExampleJson_MustAppend() {
	data :=
		`{
        "users" : {
            "count" : 1,
            "array" : ["John", "Ming"]
        }
    }`

	j, _ := gjson.LoadContent(data)

	j.MustAppend("users.array", "Lily")

	fmt.Println(j.Get("users.array").Array())

	// Output:
	// [John Ming Lily]
}

func ExampleJson_Map() {
	data :=
		`{
        "users" : {
            "count" : 1,
            "info" : {
				"name" : "John",
				"age" : 18,
				"addr" : "ChengDu"
			}
        }
    }`

	j, _ := gjson.LoadContent(data)

	fmt.Println(j.Get("users.info").Map())

	// Output:
	// map[addr:ChengDu age:18 name:John]
}

func ExampleJson_Array() {
	data :=
		`{
        "users" : {
            "count" : 1,
            "array" : ["John", "Ming"]
        }
    }`

	j, _ := gjson.LoadContent(data)

	fmt.Println(j.Get("users.array"))

	// Output:
	// ["John","Ming"]
}

func ExampleJson_Scan() {
	data := `{"name":"john","age":"18"}`

	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{}

	j, _ := gjson.LoadContent(data)
	j.Scan(&info)

	fmt.Println(info)

	// May Output:
	// {john 18}
}

func ExampleJson_Dump() {
	data := `{"name":"john","age":"18"}`

	j, _ := gjson.LoadContent(data)
	j.Dump()

	// May Output:
	// {
	//	"name": "john",
	//	"age":  "18",
	// }
}
