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
	if j, err := gjson.X解码到json(data); err != nil {
		panic(err)
	} else {
		fmt.Println("John Score:", j.X取值("users.list.1.score").X取小数32位())
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
	if j, err := gjson.X解码到json(data); err != nil {
		panic(err)
	} else {
		j.X设置分层冲突检查(true)
		fmt.Println("Users Count:", j.X取值("users.count").X取整数())
	}
	// Output:
	// Users Count: 101
}

func ExampleJson_Get_MapSliceChange() {
	jsonContent := `{"map":{"key":"value"}, "slice":[59,90]}`
	j, _ := gjson.X加载json(jsonContent)
	m := j.X取值("map").X取Map()
	fmt.Println(m)

		// 修改键值对。 md5:3e65afa62ae7277a
	m["key"] = "john"

		// 它会修改底层的键值对。 md5:ab28b164ebbac7ef
	fmt.Println(j.X取值("map").X取Map())

	s := j.X取值("slice").Array别名()
	fmt.Println(s)

		// 修改指定索引的值。 md5:2a9fd01566dd74e8
	s[0] = 100

		// 它会改变底层的切片。 md5:ee9cf84e999339cf
	fmt.Println(j.X取值("slice").Array别名())

	// output:
	// map[key:value]
	// map[key:john]
	// [59 90]
	// [100 90]
}
