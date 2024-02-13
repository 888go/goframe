// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package session类_test

import (
	"fmt"
	"time"
	
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gctx"
	"github.com/888go/goframe/os/gsession"
)

func ExampleNew() {
	manager := session类.New(time.Second)
	fmt.Println(manager.GetTTL())

	// Output:
	// 1s
}

func ExampleManager_SetStorage() {
	manager := session类.New(time.Second)
	manager.SetStorage(session类.NewStorageMemory())
	fmt.Println(manager.GetTTL())

	// Output:
	// 1s
}

func ExampleManager_GetStorage() {
	manager := session类.New(time.Second, session类.NewStorageMemory())
	size, _ := manager.GetStorage().GetSize(上下文类.X创建(), "id")
	fmt.Println(size)

	// Output:
	// 0
}

func ExampleManager_SetTTL() {
	manager := session类.New(time.Second)
	manager.SetTTL(time.Minute)
	fmt.Println(manager.GetTTL())

	// Output:
	// 1m0s
}

func ExampleSession_Set() {
	storage := session类.NewStorageFile("", time.Second)
	manager := session类.New(time.Second, storage)
	s := manager.New(上下文类.X创建())
	fmt.Println(s.X设置值("key", "val") == nil)

	// Output:
	// true
}

func ExampleSession_SetMap() {
	storage := session类.NewStorageFile("", time.Second)
	manager := session类.New(time.Second, storage)
	s := manager.New(上下文类.X创建())
	fmt.Println(s.SetMap(map[string]interface{}{}) == nil)

	// Output:
	// true
}

func ExampleSession_Remove() {
	storage := session类.NewStorageFile("", time.Second)
	manager := session类.New(time.Second, storage)
	s1 := manager.New(上下文类.X创建())
	fmt.Println(s1.Remove("key"))

	s2 := manager.New(上下文类.X创建(), "Remove")
	fmt.Println(s2.Remove("key"))

	// Output:
	// <nil>
	// <nil>
}

func ExampleSession_RemoveAll() {
	storage := session类.NewStorageFile("", time.Second)
	manager := session类.New(time.Second, storage)
	s1 := manager.New(上下文类.X创建())
	fmt.Println(s1.RemoveAll())

	s2 := manager.New(上下文类.X创建(), "Remove")
	fmt.Println(s2.RemoveAll())

	// Output:
	// <nil>
	// <nil>
}

func ExampleSession_Id() {
	storage := session类.NewStorageFile("", time.Second)
	manager := session类.New(time.Second, storage)
	s := manager.New(上下文类.X创建(), "Id")
	id, _ := s.Id()
	fmt.Println(id)

	// Output:
	// Id
}

func ExampleSession_SetId() {
	nilSession := &session类.Session{}
	fmt.Println(nilSession.SetId("id"))

	storage := session类.NewStorageFile("", time.Second)
	manager := session类.New(time.Second, storage)
	s := manager.New(上下文类.X创建())
	s.Id()
	fmt.Println(s.SetId("id"))

	// Output:
	// <nil>
	// session already started
}

func ExampleSession_SetIdFunc() {
	nilSession := &session类.Session{}
	fmt.Println(nilSession.SetIdFunc(func(ttl time.Duration) string {
		return "id"
	}))

	storage := session类.NewStorageFile("", time.Second)
	manager := session类.New(time.Second, storage)
	s := manager.New(上下文类.X创建())
	s.Id()
	fmt.Println(s.SetIdFunc(func(ttl time.Duration) string {
		return "id"
	}))

	// Output:
	// <nil>
	// session already started
}

func ExampleSession_Data() {
	storage := session类.NewStorageFile("", time.Second)
	manager := session类.New(time.Second, storage)

	s1 := manager.New(上下文类.X创建())
	data1, _ := s1.Data()
	fmt.Println(data1)

	s2 := manager.New(上下文类.X创建(), "id_data")
	data2, _ := s2.Data()
	fmt.Println(data2)

	// Output:
	// map[]
	// map[]
}

func ExampleSession_Size() {
	storage := session类.NewStorageFile("", time.Second)
	manager := session类.New(time.Second, storage)

	s1 := manager.New(上下文类.X创建())
	size1, _ := s1.Size()
	fmt.Println(size1)

	s2 := manager.New(上下文类.X创建(), "Size")
	size2, _ := s2.Size()
	fmt.Println(size2)

	// Output:
	// 0
	// 0
}

func ExampleSession_Contains() {
	storage := session类.NewStorageFile("", time.Second)
	manager := session类.New(time.Second, storage)

	s1 := manager.New(上下文类.X创建())
	notContains, _ := s1.Contains("Contains")
	fmt.Println(notContains)

	s2 := manager.New(上下文类.X创建(), "Contains")
	contains, _ := s2.Contains("Contains")
	fmt.Println(contains)

	// Output:
	// false
	// false
}

func ExampleStorageFile_SetCryptoKey() {
	storage := session类.NewStorageFile("", time.Second)
	storage.SetCryptoKey([]byte("key"))

	size, _ := storage.GetSize(上下文类.X创建(), "id")
	fmt.Println(size)

	// Output:
	// 0
}

func ExampleStorageFile_SetCryptoEnabled() {
	storage := session类.NewStorageFile("", time.Second)
	storage.SetCryptoEnabled(true)

	size, _ := storage.GetSize(上下文类.X创建(), "id")
	fmt.Println(size)

	// Output:
	// 0
}

func ExampleStorageFile_UpdateTTL() {
	var (
		ctx = 上下文类.X创建()
	)

	storage := session类.NewStorageFile("", time.Second)
	fmt.Println(storage.UpdateTTL(ctx, "id", time.Second*15))

	time.Sleep(time.Second * 11)

	// Output:
	// <nil>
}

func ExampleStorageRedis_Get() {
	storage := session类.NewStorageRedis(g.Redis类())
	val, _ := storage.Get(上下文类.X创建(), "id", "key")
	fmt.Println(val)

	// May Output:
	// <nil>
}

func ExampleStorageRedis_Data() {
	storage := session类.NewStorageRedis(g.Redis类())
	val, _ := storage.Data(上下文类.X创建(), "id")
	fmt.Println(val)

	// May Output:
	// map[]
}

func ExampleStorageRedis_GetSize() {
	storage := session类.NewStorageRedis(g.Redis类())
	val, _ := storage.GetSize(上下文类.X创建(), "id")
	fmt.Println(val)

	// May Output:
	// 0
}

func ExampleStorageRedis_Remove() {
	storage := session类.NewStorageRedis(g.Redis类())
	err := storage.Remove(上下文类.X创建(), "id", "key")
	fmt.Println(err != nil)

	// May Output:
	// true
}

func ExampleStorageRedis_RemoveAll() {
	storage := session类.NewStorageRedis(g.Redis类())
	err := storage.RemoveAll(上下文类.X创建(), "id")
	fmt.Println(err != nil)

	// May Output:
	// true
}

func ExampleStorageRedis_UpdateTTL() {
	storage := session类.NewStorageRedis(g.Redis类())
	err := storage.UpdateTTL(上下文类.X创建(), "id", time.Second*15)
	fmt.Println(err)

	time.Sleep(time.Second * 11)

	// May Output:
	// <nil>
}

func ExampleStorageRedisHashTable_Get() {
	storage := session类.NewStorageRedisHashTable(g.Redis类())

	v, err := storage.Get(上下文类.X创建(), "id", "key")

	fmt.Println(v)
	fmt.Println(err)

	// May Output:
	// <nil>
	// redis adapter is not set, missing configuration or adapter register? possible reference: https://github.com/gogf/gf/tree/master/contrib/nosql/redis
}

func ExampleStorageRedisHashTable_Data() {
	storage := session类.NewStorageRedisHashTable(g.Redis类())

	data, err := storage.Data(上下文类.X创建(), "id")

	fmt.Println(data)
	fmt.Println(err)

	// May Output:
	// map[]
	// redis adapter is not set, missing configuration or adapter register? possible reference: https://github.com/gogf/gf/tree/master/contrib/nosql/redis
}

func ExampleStorageRedisHashTable_GetSize() {
	storage := session类.NewStorageRedisHashTable(g.Redis类())

	size, err := storage.GetSize(上下文类.X创建(), "id")

	fmt.Println(size)
	fmt.Println(err)

	// May Output:
	// 0
	// redis adapter is not set, missing configuration or adapter register? possible reference: https://github.com/gogf/gf/tree/master/contrib/nosql/redis
}

func ExampleStorageRedisHashTable_Remove() {
	storage := session类.NewStorageRedisHashTable(g.Redis类())

	err := storage.Remove(上下文类.X创建(), "id", "key")

	fmt.Println(err)

	// May Output:
	// redis adapter is not set, missing configuration or adapter register? possible reference: https://github.com/gogf/gf/tree/master/contrib/nosql/redis
}

func ExampleStorageRedisHashTable_RemoveAll() {
	storage := session类.NewStorageRedisHashTable(g.Redis类())

	err := storage.RemoveAll(上下文类.X创建(), "id")

	fmt.Println(err)

	// May Output:
	// redis adapter is not set, missing configuration or adapter register? possible reference: https://github.com/gogf/gf/tree/master/contrib/nosql/redis
}

func ExampleStorageRedisHashTable_GetSession() {
	storage := session类.NewStorageRedisHashTable(g.Redis类())
	data, err := storage.GetSession(上下文类.X创建(), "id", time.Second)

	fmt.Println(data)
	fmt.Println(err)

	// May Output:
	//
	// redis adapter is not set, missing configuration or adapter register? possible reference: https://github.com/gogf/gf/tree/master/contrib/nosql/redis
}

func ExampleStorageRedisHashTable_SetSession() {
	storage := session类.NewStorageRedisHashTable(g.Redis类())

	strAnyMap := map类.StrAnyMap{}

	err := storage.SetSession(上下文类.X创建(), "id", &strAnyMap, time.Second)

	fmt.Println(err)

	// May Output:
	// redis adapter is not set, missing configuration or adapter register? possible reference: https://github.com/gogf/gf/tree/master/contrib/nosql/redis
}

func ExampleStorageRedisHashTable_UpdateTTL() {
	storage := session类.NewStorageRedisHashTable(g.Redis类())

	err := storage.UpdateTTL(上下文类.X创建(), "id", time.Second)

	fmt.Println(err)

	// May Output:
	// redis adapter is not set, missing configuration or adapter register? possible reference: https://github.com/gogf/gf/tree/master/contrib/nosql/redis
}
