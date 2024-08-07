// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package json类_test

import (
	"fmt"

	gjson "github.com/888go/goframe/encoding/gjson"
)

func ExampleConversionNormalFormats() {
	data :=
		`{
        "users" : {
            "count" : 1,
            "array" : ["John", "Ming"]
        }
    }`

	if j, err := gjson.X解码到json(data); err != nil {
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
	if j, err := gjson.X解码到json(data); err != nil {
		panic(err)
	} else {
		type Users struct {
			Count int
			Array []string
		}
		users := new(Users)
		if err := j.X取值("users").X取结构体指针(users); err != nil {
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
	if j, err := gjson.X解码到json(data); err != nil {
		panic(err)
	} else {
		type Users struct {
			Count int
			Array []string
		}
		users := new(Users)
		if err := j.X取泛型类().X取结构体指针(users); err != nil {
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
	fmt.Println(gjson.X是否为有效json(data1))
	fmt.Println(gjson.X是否为有效json(data2))

	// Output:
	// true
	// false
}

func ExampleMarshal() {
	data := map[string]interface{}{
		"name":  "john",
		"score": 100,
	}

	jsonData, _ := gjson.Marshal别名(data)
	fmt.Println(string(jsonData))

	type BaseInfo struct {
		Name string
		Age  int
	}

	info := BaseInfo{
		Name: "Guo Qiang",
		Age:  18,
	}

	infoData, _ := gjson.Marshal别名(info)
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

	infoData, _ := gjson.MarshalIndent别名(info, "", "\t")
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
	gjson.Unmarshal别名([]byte(jsonContent), &info)
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

	infoData, _ := gjson.X变量到json字节集(info)
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

	infoData := gjson.X变量到json字节集PANI(info)
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

	infoData, _ := gjson.X变量到json文本(info)
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

	infoData := gjson.X变量到json文本PANI(info)
	fmt.Println(infoData)

	// Output:
	// {"Name":"John","Age":18}
}

func ExampleDecode() {
	jsonContent := `{"name":"john","score":100}`
	info, _ := gjson.Json格式到变量([]byte(jsonContent))
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
	gjson.Json格式到变量指针([]byte(jsonContent), &info)
	fmt.Printf("%+v", info)

	// Output:
	// {Name:john Score:100}
}

func ExampleDecodeToJson() {
	jsonContent := `{"name":"john","score":100}"`
	j, _ := gjson.X解码到json([]byte(jsonContent))
	fmt.Println(j.X取Map())

	// May Output:
	// map[name:john score:100]
}
