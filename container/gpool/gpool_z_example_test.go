// 版权归GoFrame作者所有（https://goframe.org）。保留所有权利。
//
// 本源代码形式受MIT许可证条款的约束。如果gm文件中未附带MIT许可证的副本，
// 您可以从https://github.com/gogf/gf获取。
// md5:1d281c30cdc3423b

package 对象复用类_test

import (
	"database/sql"
	"fmt"
	"time"

	gpool "github.com/888go/goframe/container/gpool"
)

func ExampleNew() {
	type DBConn struct {
		Conn *sql.Conn
	}

	dbConnPool := gpool.New(time.Hour,
		func() (interface{}, error) {
			dbConn := new(DBConn)
			return dbConn, nil
		},
		func(i interface{}) {
			// 示例：关闭数据库连接
			// 调用 i.(DBConn) 的 Conn 关闭方法
			// md5:1207f4943d8a98dc
		})

	fmt.Println(dbConnPool.TTL)

	// Output:
	// 1h0m0s
}

func ExamplePool_Put() {
	type DBConn struct {
		Conn  *sql.Conn
		Limit int
	}

	dbConnPool := gpool.New(time.Hour,
		func() (interface{}, error) {
			dbConn := new(DBConn)
			dbConn.Limit = 10
			return dbConn, nil
		},
		func(i interface{}) {
			// 示例：关闭数据库连接
			// 调用 i.(DBConn) 的 Conn 关闭方法
			// md5:1207f4943d8a98dc
		})

	// get db conn
	conn, _ := dbConnPool.Get()
		// 修改这个连接的限制. md5:fbc2b791ac0ae7a0
	conn.(*DBConn).Limit = 20

	// 示例：执行相同的数据库操作
	// 使用conn指向的*DBConn的Conn方法，以context.Background()为上下文，执行SQL查询"select * from user"
	// md5:92af4813b4267108

	// put back conn
	dbConnPool.MustPut(conn)

	fmt.Println(conn.(*DBConn).Limit)

	// Output:
	// 20
}

func ExamplePool_Clear() {
	type DBConn struct {
		Conn  *sql.Conn
		Limit int
	}

	dbConnPool := gpool.New(time.Hour,
		func() (interface{}, error) {
			dbConn := new(DBConn)
			dbConn.Limit = 10
			return dbConn, nil
		},
		func(i interface{}) {
			i.(*DBConn).Limit = 0
			// 示例：关闭数据库连接
			// 调用 i.(DBConn) 的 Conn 关闭方法
			// md5:1207f4943d8a98dc
		})

	conn, _ := dbConnPool.Get()
	dbConnPool.MustPut(conn)
	dbConnPool.MustPut(conn)
	fmt.Println(dbConnPool.Size())
	dbConnPool.Clear()
	fmt.Println(dbConnPool.Size())

	// Output:
	// 2
	// 0
}

func ExamplePool_Get() {
	type DBConn struct {
		Conn  *sql.Conn
		Limit int
	}

	dbConnPool := gpool.New(time.Hour,
		func() (interface{}, error) {
			dbConn := new(DBConn)
			dbConn.Limit = 10
			return dbConn, nil
		},
		func(i interface{}) {
			// 示例：关闭数据库连接
			// 调用 i.(DBConn) 的 Conn 关闭方法
			// md5:1207f4943d8a98dc
		})

	conn, err := dbConnPool.Get()
	if err == nil {
		fmt.Println(conn.(*DBConn).Limit)
	}

	// Output:
	// 10
}

func ExamplePool_Size() {
	type DBConn struct {
		Conn  *sql.Conn
		Limit int
	}

	dbConnPool := gpool.New(time.Hour,
		func() (interface{}, error) {
			dbConn := new(DBConn)
			dbConn.Limit = 10
			return dbConn, nil
		},
		func(i interface{}) {
			// 示例：关闭数据库连接
			// 调用 i.(DBConn) 的 Conn 关闭方法
			// md5:1207f4943d8a98dc
		})

	conn, _ := dbConnPool.Get()
	fmt.Println(dbConnPool.Size())
	dbConnPool.MustPut(conn)
	dbConnPool.MustPut(conn)
	fmt.Println(dbConnPool.Size())

	// Output:
	// 0
	// 2
}

func ExamplePool_Close() {
	type DBConn struct {
		Conn  *sql.Conn
		Limit int
	}
	var (
		newFunc = func() (interface{}, error) {
			dbConn := new(DBConn)
			dbConn.Limit = 10
			return dbConn, nil
		}
		closeFunc = func(i interface{}) {
			fmt.Println("Close The Pool")
			// 示例：关闭数据库连接
			// 调用 i.(DBConn) 的 Conn 关闭方法
			// md5:1207f4943d8a98dc
		}
	)
	dbConnPool := gpool.New(time.Hour, newFunc, closeFunc)

	conn, _ := dbConnPool.Get()
	dbConnPool.MustPut(conn)

	dbConnPool.Close()

	// wait for pool close
	time.Sleep(time.Second * 1)

	// May Output:
	// Close The Pool
}
