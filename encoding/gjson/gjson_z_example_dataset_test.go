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

func ExampleJson_Set_DataSetCreate1() {
	j := gjson.New(nil)
	j.Set("name", "John")
	j.Set("score", 99.5)
	fmt.Printf(
		"Name: %s, Score: %v\n",
		j.Get("name").String(),
		j.Get("score").Float32(),
	)
	fmt.Println(j.MustToJsonString())

	// Output:
	// Name: John, Score: 99.5
	// {"name":"John","score":99.5}
}

func ExampleJson_Set_DataSetCreate2() {
	j := gjson.New(nil)
	for i := 0; i < 5; i++ {
		j.Set(fmt.Sprintf(`%d.id`, i), i)
		j.Set(fmt.Sprintf(`%d.name`, i), fmt.Sprintf(`student-%d`, i))
	}
	fmt.Println(j.MustToJsonString())

	// Output:
	// [{"id":0,"name":"student-0"},{"id":1,"name":"student-1"},{"id":2,"name":"student-2"},{"id":3,"name":"student-3"},{"id":4,"name":"student-4"}]
}

func ExampleJson_DataSetRuntimeEdit() {
	data :=
		`{
        "users" : {
            "count" : 2,
            "list"  : [
                {"name" : "Ming", "score" : 60},
                {"name" : "John", "score" : 59}
            ]
        }
    }`
	if j, err := gjson.DecodeToJson(data); err != nil {
		panic(err)
	} else {
		j.Set("users.list.1.score", 100)
		fmt.Println("John Score:", j.Get("users.list.1.score").Float32())
		fmt.Println(j.MustToJsonString())
	}
	// Output:
	// John Score: 100
	// {"users":{"count":2,"list":[{"name":"Ming","score":60},{"name":"John","score":100}]}}
}
