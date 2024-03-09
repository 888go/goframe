// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package json类_test

import (
	"fmt"
	
	"github.com/888go/goframe/gjson"
)

func ExampleJson_Set_DataSetCreate1() {
	j := json类.X创建(nil)
	j.X设置值("name", "John")
	j.X设置值("score", 99.5)
	fmt.Printf(
		"Name: %s, Score: %v\n",
		j.X取值("name").String(),
		j.X取值("score").Float32(),
	)
	fmt.Println(j.X取json文本PANI())

	// Output:
	// Name: John, Score: 99.5
	// {"name":"John","score":99.5}
}

func ExampleJson_Set_DataSetCreate2() {
	j := json类.X创建(nil)
	for i := 0; i < 5; i++ {
		j.X设置值(fmt.Sprintf(`%d.id`, i), i)
		j.X设置值(fmt.Sprintf(`%d.name`, i), fmt.Sprintf(`student-%d`, i))
	}
	fmt.Println(j.X取json文本PANI())

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
	if j, err := json类.X解码到json(data); err != nil {
		panic(err)
	} else {
		j.X设置值("users.list.1.score", 100)
		fmt.Println("John Score:", j.X取值("users.list.1.score").Float32())
		fmt.Println(j.X取json文本PANI())
	}
	// Output:
	// John Score: 100
	// {"users":{"count":2,"list":[{"name":"Ming","score":60},{"name":"John","score":100}]}}
}
