// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gdb_test

import (
	"context"
	
	"github.com/888go/goframe/database/gdb"
	"github.com/888go/goframe/frame/g"
)

func ExampleTransaction() {
	g.DB().Transaction(context.TODO(), func(ctx context.Context, tx gdb.TX) error {
		// user
		result, err := tx.Insert("user", g.Map{
			"passport": "john",
			"password": "12345678",
			"nickname": "JohnGuo",
		})
		if err != nil {
			return err
		}
		// user_detail
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		_, err = tx.Insert("user_detail", g.Map{
			"uid":       id,
			"site":      "https://johng.cn",
			"true_name": "GuoQiang",
		})
		if err != nil {
			return err
		}
		return nil
	})
}
