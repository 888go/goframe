	// 版权归GoFrame作者(https:	//goframe.org)所有。保留所有权利。
	//
	// 本源代码形式受MIT许可证条款约束。
	// 如果未随本文件一同分发MIT许可证副本，
	// 您可以在https:	//github.com/gogf/gf处获取。
	// md5:a9832f33b234e3f3

package gtag_test

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/gogf/gf/v2/util/gtag"
)

func ExampleSet() {
	type User struct {
		g.Meta `name:"User Struct" description:"{UserDescription}"`
	}
	gtag.Sets(g.MapStrStr{
		`UserDescription`: `This is a demo struct named "User Struct"`,
	})
	fmt.Println(gmeta.Get(User{}, `description`))

	// Output:
	// This is a demo struct named "User Struct"
}
