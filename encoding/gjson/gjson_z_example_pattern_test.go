// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gjson_test

import (
	"fmt"
	
	"github.com/888go/goframe/encoding/gjson"
)

func ExampleDecodeToJson_PatternGet() {
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
		fmt.Println("John Score:", j.Get("users.list.1.score").Float32())
	}
	// Output:
	// John Score: 99.5
}

func ExampleDecodeToJson_PatternViolenceCheck() {
	data :=
		`{
        "users" : {
            "count" : 100
        },
        "users.count" : 101
    }`
	if j, err := gjson.DecodeToJson(data); err != nil {
		panic(err)
	} else {
		j.SetViolenceCheck(true)
		fmt.Println("Users Count:", j.Get("users.count").Int())
	}
	// Output:
	// Users Count: 101
}

func ExampleJson_Get_MapSliceChange() {
	jsonContent := `{"map":{"key":"value"}, "slice":[59,90]}`
	j, _ := gjson.LoadJson(jsonContent)
	m := j.Get("map").Map()
	fmt.Println(m)

	// 修改键值对。
	m["key"] = "john"

	// 它会改变底层的键值对。
	fmt.Println(j.Get("map").Map())

	s := j.Get("slice").Array()
	fmt.Println(s)

	// 修改指定索引处的值。
	s[0] = 100

	// 它会改变底层的切片。
	fmt.Println(j.Get("slice").Array())

	// output:
	// map[key:value]
	// map[key:john]
	// [59 90]
	// [100 90]
}
