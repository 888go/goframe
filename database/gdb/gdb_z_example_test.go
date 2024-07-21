				// 版权归GoFrame作者(https:				//goframe.org)所有。保留所有权利。
				//
				// 本源代码形式受MIT许可证条款约束。
				// 如果未随本文件一同分发MIT许可证副本，
				// 您可以在https:				//github.com/gogf/gf处获取。
				// md5:a9832f33b234e3f3

package gdb_test//bm:db类_test

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
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
