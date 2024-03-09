// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package json类_test

import (
	"fmt"
	
	"github.com/888go/goframe/gjson"
)

func ExampleConversionNormalFormats() {
	data :=
		`{
        "users" : {
            "count" : 1,
            "array" : ["John", "Ming"]
        }
    }`

	if j, err := json类.X解码到json(data); err != nil {
		panic(err)
	} else {
		fmt.Println("JSON:")
		fmt.Println(j.X取json文本PANI())
		fmt.Println("======================")

		fmt.Println("XML:")
		fmt.Println(j.X取xml文本PANI())
		fmt.Println("======================")

		fmt.Println("YAML:")
		fmt.Println(j.X取YAML文本PANI())
		fmt.Println("======================")

		fmt.Println("TOML:")
		fmt.Println(j.X取TOML文本PANI())
	}

	// Output:
	// JSON:
	// {"users":{"array":["John","Ming"],"count":1}}
	// ======================
	// XML:
	// <users><array>John</array><array>Ming</array><count>1</count></users>
	// ======================
	// YAML:
	// users:
	//     array:
	//         - John
	//         - Ming
	//     count: 1
	//
	// ======================
	// TOML:
	// [users]
	//   array = ["John", "Ming"]
	//   count = 1.0
}

func ExampleJson_ConversionGetStruct() {
	data :=
		`{
        "users" : {
            "count" : 1,
            "array" : ["John", "Ming"]
        }
    }`
	if j, err := json类.X解码到json(data); err != nil {
		panic(err)
	} else {
		type Users struct {
			Count int
			Array []string
		}
		users := new(Users)
		if err := j.X取值("users").Scan(users); err != nil {
			panic(err)
		}
		fmt.Printf(`%+v`, users)
	}

	// Output:
	// &{Count:1 Array:[John Ming]}
}

func ExampleJson_ConversionToStruct() {
	data :=
		`
	{
        "count" : 1,
        "array" : ["John", "Ming"]
    }`
	if j, err := json类.X解码到json(data); err != nil {
		panic(err)
	} else {
		type Users struct {
			Count int
			Array []string
		}
		users := new(Users)
		if err := j.X取泛型类().Scan(users); err != nil {
			panic(err)
		}
		fmt.Printf(`%+v`, users)
	}

	// Output:
	// &{Count:1 Array:[John Ming]}
}

func ExampleValid() {
	data1 := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	data2 := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]`)
	fmt.Println(json类.X是否为有效json(data1))
	fmt.Println(json类.X是否为有效json(data2))

	// Output:
	// true
	// false
}

func ExampleMarshal() {
	data := map[string]interface{}{
		"name":  "john",
		"score": 100,
	}

	jsonData, _ := json类.Marshal别名(data)
	fmt.Println(string(jsonData))

	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "Guo Qiang",
		Age:  18,
	}

	infoData, _ := json类.Marshal别名(info)
	fmt.Println(string(infoData))

	// Output:
	// {"name":"john","score":100}
	// {"Name":"Guo Qiang","Age":18}
}

func ExampleMarshalIndent() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	infoData, _ := json类.MarshalIndent别名(info, "", "\t")
	fmt.Println(string(infoData))

	// Output:
	// {
	//	"Name": "John",
	//	"Age": 18
	// }
}

func ExampleUnmarshal() {
	type BaseInfo struct {
		Name  string
		Score int
	}

	var info BaseInfo

	jsonContent := "{\"name\":\"john\",\"score\":100}"
	json类.Unmarshal别名([]byte(jsonContent), &info)
	fmt.Printf("%+v", info)

	// Output:
	// {Name:john Score:100}
}

func ExampleEncode() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	infoData, _ := json类.X变量到json字节集(info)
	fmt.Println(string(infoData))

	// Output:
	// {"Name":"John","Age":18}
}

func ExampleMustEncode() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	infoData := json类.X变量到json字节集PANI(info)
	fmt.Println(string(infoData))

	// Output:
	// {"Name":"John","Age":18}
}

func ExampleEncodeString() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	infoData, _ := json类.X变量到json文本(info)
	fmt.Println(infoData)

	// Output:
	// {"Name":"John","Age":18}
}

func ExampleMustEncodeString() {
	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "John",
		Age:  18,
	}

	infoData := json类.X变量到json文本PANI(info)
	fmt.Println(infoData)

	// Output:
	// {"Name":"John","Age":18}
}

func ExampleDecode() {
	jsonContent := `{"name":"john","score":100}`
	info, _ := json类.Json格式到变量([]byte(jsonContent))
	fmt.Println(info)

	// Output:
	// map[name:john score:100]
}

func ExampleDecodeTo() {
	type BaseInfo struct {
		Name  string
		Score int
	}

	var info BaseInfo

	jsonContent := "{\"name\":\"john\",\"score\":100}"
	json类.Json格式到变量指针([]byte(jsonContent), &info)
	fmt.Printf("%+v", info)

	// Output:
	// {Name:john Score:100}
}

func ExampleDecodeToJson() {
	jsonContent := `{"name":"john","score":100}"`
	j, _ := json类.X解码到json([]byte(jsonContent))
	fmt.Println(j.X取Map())

	// May Output:
	// map[name:john score:100]
}
