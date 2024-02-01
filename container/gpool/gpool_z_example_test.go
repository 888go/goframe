// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随 gm 文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一个。

package gpool_test
import (
	"database/sql"
	"fmt"
	"time"
	
	"github.com/888go/goframe/container/gpool"
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
// i.(DBConn).Conn.Close()
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
// i.(DBConn).Conn.Close()
		})

	// get db conn
	conn, _ := dbConnPool.Get()
	// 修改此连接的限制
	conn.(*DBConn).Limit = 20

// 示例：执行相同的数据库操作
// conn.(*DBConn).Conn.QueryContext(context.Background(), "select * from user")
// 翻译：
// 示例：进行同样的数据库操作
// conn.(*DBConn).Conn.QueryContext 使用 context.Background() 作为上下文，执行 SQL 查询语句："从 user 表中选择所有列"

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
// i.(DBConn).Conn.Close()
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
// i.(DBConn).Conn.Close()
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
// i.(DBConn).Conn.Close()
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
// i.(DBConn).Conn.Close()
		}
	)
	dbConnPool := gpool.New(time.Hour, newFunc, closeFunc)

	conn, _ := dbConnPool.Get()
	dbConnPool.MustPut(conn)

	dbConnPool.Close()

	// 等待连接池关闭
	time.Sleep(time.Second * 1)

	// May Output:
	// Close The Pool
}
