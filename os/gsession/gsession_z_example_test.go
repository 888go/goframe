// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package session类_test

import (
	"fmt"
	"time"

	gmap "github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/frame/g"
	gctx "github.com/888go/goframe/os/gctx"
	gsession "github.com/888go/goframe/os/gsession"
)

func ExampleNew() {
	manager := gsession.New(time.Second)
	fmt.Println(manager.GetTTL())

	// Output:
	// 1s
}

func ExampleManager_SetStorage() {
	manager := gsession.New(time.Second)
	manager.SetStorage(gsession.NewStorageMemory())
	fmt.Println(manager.GetTTL())

	// Output:
	// 1s
}

func ExampleManager_GetStorage() {
	manager := gsession.New(time.Second, gsession.NewStorageMemory())
	size, _ := manager.GetStorage().GetSize(gctx.X创建(), "id")
	fmt.Println(size)

	// Output:
	// 0
}

func ExampleManager_SetTTL() {
	manager := gsession.New(time.Second)
	manager.SetTTL(time.Minute)
	fmt.Println(manager.GetTTL())

	// Output:
	// 1m0s
}

func ExampleSession_Set() {
	storage := gsession.NewStorageFile("", time.Second)
	manager := gsession.New(time.Second, storage)
	s := manager.New(gctx.X创建())
	fmt.Println(s.X设置值("key", "val") == nil)

	// Output:
	// true
}

func ExampleSession_SetMap() {
	storage := gsession.NewStorageFile("", time.Second)
	manager := gsession.New(time.Second, storage)
	s := manager.New(gctx.X创建())
	fmt.Println(s.SetMap(map[string]interface{}{}) == nil)

	// Output:
	// true
}

func ExampleSession_Remove() {
	storage := gsession.NewStorageFile("", time.Second)
	manager := gsession.New(time.Second, storage)
	s1 := manager.New(gctx.X创建())
	fmt.Println(s1.Remove("key"))

	s2 := manager.New(gctx.X创建(), "Remove")
	fmt.Println(s2.Remove("key"))

	// Output:
	// <nil>
	// <nil>
}

func ExampleSession_RemoveAll() {
	storage := gsession.NewStorageFile("", time.Second)
	manager := gsession.New(time.Second, storage)
	s1 := manager.New(gctx.X创建())
	fmt.Println(s1.RemoveAll())

	s2 := manager.New(gctx.X创建(), "Remove")
	fmt.Println(s2.RemoveAll())

	// Output:
	// <nil>
	// <nil>
}

func ExampleSession_Id() {
	storage := gsession.NewStorageFile("", time.Second)
	manager := gsession.New(time.Second, storage)
	s := manager.New(gctx.X创建(), "Id")
	id, _ := s.Id()
	fmt.Println(id)

	// Output:
	// Id
}

func ExampleSession_SetId() {
	nilSession := &gsession.Session{}
	fmt.Println(nilSession.SetId("id"))

	storage := gsession.NewStorageFile("", time.Second)
	manager := gsession.New(time.Second, storage)
	s := manager.New(gctx.X创建())
	s.Id()
	fmt.Println(s.SetId("id"))

	// Output:
	// <nil>
	// session already started
}

func ExampleSession_SetIdFunc() {
	nilSession := &gsession.Session{}
	fmt.Println(nilSession.SetIdFunc(func(ttl time.Duration) string {
		return "id"
	}))

	storage := gsession.NewStorageFile("", time.Second)
	manager := gsession.New(time.Second, storage)
	s := manager.New(gctx.X创建())
	s.Id()
	fmt.Println(s.SetIdFunc(func(ttl time.Duration) string {
		return "id"
	}))

	// Output:
	// <nil>
	// session already started
}

func ExampleSession_Data() {
	storage := gsession.NewStorageFile("", time.Second)
	manager := gsession.New(time.Second, storage)

	s1 := manager.New(gctx.X创建())
	data1, _ := s1.Data()
	fmt.Println(data1)

	s2 := manager.New(gctx.X创建(), "id_data")
	data2, _ := s2.Data()
	fmt.Println(data2)

	// Output:
	// map[]
	// map[]
}

func ExampleSession_Size() {
	storage := gsession.NewStorageFile("", time.Second)
	manager := gsession.New(time.Second, storage)

	s1 := manager.New(gctx.X创建())
	size1, _ := s1.Size()
	fmt.Println(size1)

	s2 := manager.New(gctx.X创建(), "Size")
	size2, _ := s2.Size()
	fmt.Println(size2)

	// Output:
	// 0
	// 0
}

func ExampleSession_Contains() {
	storage := gsession.NewStorageFile("", time.Second)
	manager := gsession.New(time.Second, storage)

	s1 := manager.New(gctx.X创建())
	notContains, _ := s1.Contains("Contains")
	fmt.Println(notContains)

	s2 := manager.New(gctx.X创建(), "Contains")
	contains, _ := s2.Contains("Contains")
	fmt.Println(contains)

	// Output:
	// false
	// false
}

func ExampleStorageFile_SetCryptoKey() {
	storage := gsession.NewStorageFile("", time.Second)
	storage.SetCryptoKey([]byte("key"))

	size, _ := storage.GetSize(gctx.X创建(), "id")
	fmt.Println(size)

	// Output:
	// 0
}

func ExampleStorageFile_SetCryptoEnabled() {
	storage := gsession.NewStorageFile("", time.Second)
	storage.SetCryptoEnabled(true)

	size, _ := storage.GetSize(gctx.X创建(), "id")
	fmt.Println(size)

	// Output:
	// 0
}

func ExampleStorageFile_UpdateTTL() {
	var (
		ctx = gctx.X创建()
	)

	storage := gsession.NewStorageFile("", time.Second)
	fmt.Println(storage.UpdateTTL(ctx, "id", time.Second*15))

	time.Sleep(time.Second * 11)

	// Output:
	// <nil>
}

func ExampleStorageRedis_Get() {
	storage := gsession.NewStorageRedis(g.Redis类())
	val, _ := storage.Get(gctx.X创建(), "id", "key")
	fmt.Println(val)

	// May Output:
	// <nil>
}

func ExampleStorageRedis_Data() {
	storage := gsession.NewStorageRedis(g.Redis类())
	val, _ := storage.Data(gctx.X创建(), "id")
	fmt.Println(val)

	// May Output:
	// map[]
}

func ExampleStorageRedis_GetSize() {
	storage := gsession.NewStorageRedis(g.Redis类())
	val, _ := storage.GetSize(gctx.X创建(), "id")
	fmt.Println(val)

	// May Output:
	// 0
}

func ExampleStorageRedis_Remove() {
	storage := gsession.NewStorageRedis(g.Redis类())
	err := storage.Remove(gctx.X创建(), "id", "key")
	fmt.Println(err != nil)

	// May Output:
	// true
}

func ExampleStorageRedis_RemoveAll() {
	storage := gsession.NewStorageRedis(g.Redis类())
	err := storage.RemoveAll(gctx.X创建(), "id")
	fmt.Println(err != nil)

	// May Output:
	// true
}

func ExampleStorageRedis_UpdateTTL() {
	storage := gsession.NewStorageRedis(g.Redis类())
	err := storage.UpdateTTL(gctx.X创建(), "id", time.Second*15)
	fmt.Println(err)

	time.Sleep(time.Second * 11)

	// May Output:
	// <nil>
}

func ExampleStorageRedisHashTable_Get() {
	storage := gsession.NewStorageRedisHashTable(g.Redis类())

	v, err := storage.Get(gctx.X创建(), "id", "key")

	fmt.Println(v)
	fmt.Println(err)

	// May Output:
	// <nil>
	// redis adapter is not set, missing configuration or adapter register? possible reference: https://github.com/gogf/gf/tree/master/contrib/nosql/redis
}

func ExampleStorageRedisHashTable_Data() {
	storage := gsession.NewStorageRedisHashTable(g.Redis类())

	data, err := storage.Data(gctx.X创建(), "id")

	fmt.Println(data)
	fmt.Println(err)

	// May Output:
	// map[]
	// redis adapter is not set, missing configuration or adapter register? possible reference: https://github.com/gogf/gf/tree/master/contrib/nosql/redis
}

func ExampleStorageRedisHashTable_GetSize() {
	storage := gsession.NewStorageRedisHashTable(g.Redis类())

	size, err := storage.GetSize(gctx.X创建(), "id")

	fmt.Println(size)
	fmt.Println(err)

	// May Output:
	// 0
	// redis adapter is not set, missing configuration or adapter register? possible reference: https://github.com/gogf/gf/tree/master/contrib/nosql/redis
}

func ExampleStorageRedisHashTable_Remove() {
	storage := gsession.NewStorageRedisHashTable(g.Redis类())

	err := storage.Remove(gctx.X创建(), "id", "key")

	fmt.Println(err)

	// May Output:
	// redis adapter is not set, missing configuration or adapter register? possible reference: https://github.com/gogf/gf/tree/master/contrib/nosql/redis
}

func ExampleStorageRedisHashTable_RemoveAll() {
	storage := gsession.NewStorageRedisHashTable(g.Redis类())

	err := storage.RemoveAll(gctx.X创建(), "id")

	fmt.Println(err)

	// May Output:
	// redis adapter is not set, missing configuration or adapter register? possible reference: https://github.com/gogf/gf/tree/master/contrib/nosql/redis
}

func ExampleStorageRedisHashTable_GetSession() {
	storage := gsession.NewStorageRedisHashTable(g.Redis类())
	data, err := storage.GetSession(gctx.X创建(), "id", time.Second)

	fmt.Println(data)
	fmt.Println(err)

	// May Output:
	//
	// redis adapter is not set, missing configuration or adapter register? possible reference: https://github.com/gogf/gf/tree/master/contrib/nosql/redis
}

func ExampleStorageRedisHashTable_SetSession() {
	storage := gsession.NewStorageRedisHashTable(g.Redis类())

	strAnyMap := gmap.StrAnyMap{}

	err := storage.SetSession(gctx.X创建(), "id", &strAnyMap, time.Second)

	fmt.Println(err)

	// May Output:
	// redis adapter is not set, missing configuration or adapter register? possible reference: https://github.com/gogf/gf/tree/master/contrib/nosql/redis
}

func ExampleStorageRedisHashTable_UpdateTTL() {
	storage := gsession.NewStorageRedisHashTable(g.Redis类())

	err := storage.UpdateTTL(gctx.X创建(), "id", time.Second)

	fmt.Println(err)

	// May Output:
	// redis adapter is not set, missing configuration or adapter register? possible reference: https://github.com/gogf/gf/tree/master/contrib/nosql/redis
}
