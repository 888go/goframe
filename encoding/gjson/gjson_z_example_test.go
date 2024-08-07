package json类_test

import (
	"fmt"

	gjson "github.com/888go/goframe/encoding/gjson"
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
	if j, err := gjson.X解码到json(data); err != nil {
		panic(err)
	} else {
		j.X设置参数分隔符('#')
		fmt.Println("John Score:", j.X取值("users#list#1#score").X取小数32位())
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
	if j, err := gjson.X解码到json(data); err != nil {
		fmt.Println(err)
	} else {
		j.X设置分层冲突检查(false)
		fmt.Println("Users Count:", j.X取值("users.count"))
		j.X设置分层冲突检查(true)
		fmt.Println("Users Count:", j.X取值("users.count"))
	}
	// Output:
	// Users Count: 100
	// Users Count: 101
}

// ========================================================================
// JSON
// ========================================================================
// 这段注释是关于JSON的，可能表示下面的代码与JSON（JavaScript Object Notation）处理有关，JSON是一种轻量级的数据交换格式。
// md5:9f636a538977ae4f
func ExampleJson_ToJson() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.X创建(info)
	jsonBytes, _ := j.X取json字节集()
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

	j := gjson.X创建(info)
	jsonStr, _ := j.X取json文本()
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

	j := gjson.X创建(info)
	jsonBytes, _ := j.X取json字节集并格式化()
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

	j := gjson.X创建(info)
	jsonStr, _ := j.X取json文本并格式化()
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

	j := gjson.X创建(info)
	jsonBytes := j.X取json字节集PANI()
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

	j := gjson.X创建(info)
	jsonStr := j.X取json文本PANI()
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

	j := gjson.X创建(info)
	jsonBytes := j.X取json字节集并格式化PANI()
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

	j := gjson.X创建(info)
	jsonStr := j.X取json文本并格式化PANI()
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
// 这段代码的注释翻译成中文是：
// 
// ========================================================================
// XML
// ========================================================================
// 这是对XML部分的注释，表示接下来的内容与XML（可扩展标记语言）相关。
// md5:931c367389ad5867
func ExampleJson_ToXml() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.X创建(info)
	xmlBytes, _ := j.X取xml字节集()
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

	j := gjson.X创建(info)
	xmlStr, _ := j.X取xml文本()
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

	j := gjson.X创建(info)
	xmlBytes, _ := j.X取xml字节集并格式化()
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

	j := gjson.X创建(info)
	xmlStr, _ := j.X取xml文本并格式化()
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

	j := gjson.X创建(info)
	xmlBytes := j.X取xml字节集PANI()
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

	j := gjson.X创建(info)
	xmlStr := j.X取xml文本PANI()
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

	j := gjson.X创建(info)
	xmlBytes := j.X取xml字节集并格式化PANI()
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

	j := gjson.X创建(info)
	xmlStr := j.X取xml文本并格式化PANI()
	fmt.Println(string(xmlStr))

	// Output:
	// <doc>
	//	<Age>18</Age>
	//	<Name>John</Name>
	// </doc>
}

// ========================================================================
// YAML
// ========================================================================
// md5:86131a4a0253d702
func ExampleJson_ToYaml() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.X创建(info)
	YamlBytes, _ := j.X取YAML字节集()
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

	j := gjson.X创建(info)
	YamlStr, _ := j.X取YAML文本()
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

	j := gjson.X创建(info)
	YamlBytes, _ := j.X取YAML字节集并格式化("")
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

	j := gjson.X创建(info)
	YamlBytes := j.X取YAML字节集PANI()
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

	j := gjson.X创建(info)
	YamlStr := j.X取YAML文本PANI()
	fmt.Println(string(YamlStr))

	// Output:
	// Age: 18
	// Name: John
}

// ========================================================================
// TOML 配置文件格式
// ========================================================================
// md5:2a6d07eba917d4f3
func ExampleJson_ToToml() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.X创建(info)
	TomlBytes, _ := j.X取TOML字节集()
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

	j := gjson.X创建(info)
	TomlStr, _ := j.X取TOML文本()
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

	j := gjson.X创建(info)
	TomlBytes := j.X取TOML字节集PANI()
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

	j := gjson.X创建(info)
	TomlStr := j.X取TOML文本PANI()
	fmt.Println(string(TomlStr))

	// Output:
	// Age = 18
	// Name = "John"
}

// ========================================================================
// INI
// ========================================================================
// md5:a7d46faaad75eec6
func ExampleJson_ToIni() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.X创建(info)
	IniBytes, _ := j.X取ini字节集()
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

	j := gjson.X创建(info)
	IniStr, _ := j.X取ini文本()
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

	j := gjson.X创建(info)
	IniBytes := j.X取ini字节集PANI()
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

	j := gjson.X创建(info)
	IniStr := j.X取ini文本PANI()
	fmt.Println(string(IniStr))

	// Output:
	// Name=John
}

// ========================================================================
// 属性
// ========================================================================
// md5:b684439963732a7f
func ExampleJson_ToProperties() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	j := gjson.X创建(info)
	pr, _ := j.X取properties字节集()
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

	j := gjson.X创建(info)
	pr, _ := j.X取properties文本()
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

	j := gjson.X创建(info)
	pr := j.X取properties字节集PANI()
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

	j := gjson.X创建(info)
	pr := j.X取properties文本PANI()
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

	j := gjson.X创建(info)
	jsonBytes, _ := j.MarshalJSON()
	fmt.Println(string(jsonBytes))

	// Output:
	// {"Age":18,"Name":"John"}
}

func ExampleJson_UnmarshalJSON() {
	jsonStr := `{"Age":18,"Name":"John"}`

	j := gjson.X创建("")
	j.UnmarshalJSON([]byte(jsonStr))
	fmt.Println(j.X取Map())

	// Output:
	// map[Age:18 Name:John]
}

func ExampleJson_UnmarshalValue_Yaml() {
	yamlContent :=
		`base:
  name: john
  score: 100`

	j := gjson.X创建("")
	j.UnmarshalValue([]byte(yamlContent))
	fmt.Println(j.X取泛型类().String())

	// Output:
	// {"base":{"name":"john","score":100}}
}

func ExampleJson_UnmarshalValue_Xml() {
	xmlStr := `<?xml version="1.0" encoding="UTF-8"?><doc><name>john</name><score>100</score></doc>`

	j := gjson.X创建("")
	j.UnmarshalValue([]byte(xmlStr))
	fmt.Println(j.X取泛型类().String())

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

	j := gjson.X创建(info)
	fmt.Println(j.X取MapStrAny())

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

	j := gjson.X创建(infoList)
	fmt.Println(j.X取any切片())

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

	j := gjson.X创建(info)
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

	j := gjson.X创建(info)
	fmt.Println(j.X取泛型类().String())
	fmt.Println(j.X取泛型类().X取Map())

	// Output:
	// {"Age":18,"Name":"John"}
	// map[Age:18 Name:John]
}

func ExampleJson_IsNil() {
	data1 := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	data2 := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]`)

	j1, _ := gjson.X加载并自动识别格式(data1)
	fmt.Println(j1.X是否为Nil())

	j2, _ := gjson.X加载并自动识别格式(data2)
	fmt.Println(j2.X是否为Nil())

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

	j, _ := gjson.X加载并自动识别格式(data)
	fmt.Println(j.X取值("."))
	fmt.Println(j.X取值("users"))
	fmt.Println(j.X取值("users.count"))
	fmt.Println(j.X取值("users.array"))

	var nilJ *gjson.Json = nil
	fmt.Println(nilJ.X取值("."))

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

	j, _ := gjson.X加载并自动识别格式(data)

	fmt.Println(j.X取对象("users.array").X取切片())

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

	j, _ := gjson.X加载并自动识别格式(data)

	jsons := j.X取对象切片("users.array")
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

	j, _ := gjson.X加载并自动识别格式(data)

	jsonMap := j.X取对象Map("users.array")

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

	j := gjson.X创建(info)
	j.X设置值("Addr", "ChengDu")
	j.X设置值("Friends.0", "Tom")
	fmt.Println(j.X取泛型类().String())

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

	j := gjson.X创建(info)
	j.X设置值PANI("Addr", "ChengDu")
	fmt.Println(j.X取泛型类().String())

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

	j := gjson.X创建(info)
	j.X删除("Age")
	fmt.Println(j.X取泛型类().String())

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

	j := gjson.X创建(info)
	j.X删除PANI("Age")
	fmt.Println(j.X取泛型类().String())

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

	j := gjson.X创建(info)
	fmt.Println(j.X是否存在("Age"))
	fmt.Println(j.X是否存在("Addr"))

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

	j, _ := gjson.X加载并自动识别格式(data)

	fmt.Println(j.X取长度("users.nameArray"))
	fmt.Println(j.X取长度("users.infoMap"))

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

	j, _ := gjson.X加载并自动识别格式(data)

	j.X加入("users.array", "Lily")

	fmt.Println(j.X取值("users.array").Array别名())

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

	j, _ := gjson.X加载并自动识别格式(data)

	j.X加入PANI("users.array", "Lily")

	fmt.Println(j.X取值("users.array").Array别名())

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

	j, _ := gjson.X加载并自动识别格式(data)

	fmt.Println(j.X取值("users.info").X取Map())

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

	j, _ := gjson.X加载并自动识别格式(data)

	fmt.Println(j.X取值("users.array"))

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

	j, _ := gjson.X加载并自动识别格式(data)
	j.X取结构体指针(&info)

	fmt.Println(info)

	// May Output:
	// {john 18}
}

func ExampleJson_Dump() {
	data := `{"name":"john","age":"18"}`

	j, _ := gjson.X加载并自动识别格式(data)
	j.X调试输出()

	// May Output:
	// {
	//	"name": "john",
	//	"age":  "18",
	// }
}
