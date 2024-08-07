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

func ExampleJson_Set_DataSetCreate1() {
	j := gjson.X创建(nil)
	j.X设置值("name", "John")
	j.X设置值("score", 99.5)
	fmt.Printf(
		"Name: %s, Score: %v\n",
		j.X取值("name").String(),
		j.X取值("score").X取小数32位(),
	)
	fmt.Println(j.X取json文本PANI())

	// Output:
	// Name: John, Score: 99.5
	// {"name":"John","score":99.5}
}

func ExampleJson_Set_DataSetCreate2() {
	j := gjson.X创建(nil)
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
	if j, err := gjson.X解码到json(data); err != nil {
		panic(err)
	} else {
		j.X设置值("users.list.1.score", 100)
		fmt.Println("John Score:", j.X取值("users.list.1.score").X取小数32位())
		fmt.Println(j.X取json文本PANI())
	}
	// Output:
	// John Score: 100
	// {"users":{"count":2,"list":[{"name":"Ming","score":60},{"name":"John","score":100}]}}
}
