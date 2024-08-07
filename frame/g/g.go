// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包g提供了常用类型/函数的定义及耦合调用，以便创建常用的对象。 md5:ef4a19ad9323813a
package g

import (
	"context"

	gvar "github.com/888go/goframe/container/gvar"
	gmeta "github.com/888go/goframe/util/gmeta"
)

type (
	Var  = gvar.Var        // Var 是一个通用变量接口，类似于泛型。 md5:2d8d391465845592
	Ctx  = context.Context // Ctx是常用类型context.Context的别名。 md5:b60c4b141eb5549c
	Meta = gmeta.Meta      // Meta 是常用类型 gmeta.Meta 的别名。 md5:5a38da34ec0eacb2
)

type (
	Map        = map[string]interface{}      // Map 是常用映射类型 map[string]interface{} 的别名。 md5:8b62aec3a70d17d1
	MapAnyAny  = map[interface{}]interface{} // MapAnyAny 是常用 map 类型 map[interface{}]interface{} 的别名。 md5:9a81681bf8ebfc85
	MapAnyStr  = map[interface{}]string      // MapAnyStr是常用map类型map[interface{}]string的别名。 md5:30d02524cd1788b7
	MapAnyInt  = map[interface{}]int         // MapAnyInt 是一个常用的映射类型，键为interface{}，值为int的别名。 md5:decce2b63961b555
	MapStrAny  = map[string]interface{}      // MapStrAny 是常用映射类型 map[string]interface{} 的别名。 md5:cc0e9426f8d44672
	MapStrStr  = map[string]string           // MapStrStr 是常用 map 类型 map[string]string 的别名。 md5:93033a49680d75b9
	MapStrInt  = map[string]int              // MapStrInt是常用map类型map[string]int的别名。 md5:90201f6646f5919e
	MapIntAny  = map[int]interface{}         // MapIntAny是常用的map类型map[int]interface{}的别名。 md5:ce4daf18697d3f31
	MapIntStr  = map[int]string              // MapIntStr 是常用映射类型 map[int]string 的别名。 md5:9f579e30a314d76f
	MapIntInt  = map[int]int                 // MapIntInt 是常用 map 类型 map[int]int 的别名。 md5:1624c134c49d8be9
	MapAnyBool = map[interface{}]bool        // MapAnyBool 是一个常用的映射类型，键为interface{}，值为bool的别名。 md5:d7ade515deebc839
	MapStrBool = map[string]bool             // MapStrBool是常用的map类型map[string]bool的别名。 md5:9b58348f2bdc3832
	MapIntBool = map[int]bool                // MapIntBool 是常用映射类型 map[int]bool 的别名。 md5:fae965f49030c4f0
)

type (
	Map切片        = []Map        // List 是常用切片类型 []Map 的别名。 md5:ff03e40655e3f35f
	MapAnyAny切片  = []MapAnyAny  // ListAnyAny是常用切片类型[]MapAnyAny的别名。 md5:f6b168f11392abb0
	MapAnyStr切片  = []MapAnyStr  // ListAnyStr 是常用的切片类型 []MapAnyStr 的别名。 md5:d7812e32db355aa6
	MapAnyInt切片  = []MapAnyInt  // ListAnyInt 是常用切片类型 []MapAnyInt 的别名。 md5:8562af0324cd1e5a
	MapStrAny切片  = []MapStrAny  // ListStrAny 是常用切片类型 []MapStrAny 的别名。 md5:144077caa161cdf6
	MapStrStr切片  = []MapStrStr  // ListStrStr是常用的切片类型[]MapStrStr的别名。 md5:2045a3a5bab66ff3
	MapStrInt切片  = []MapStrInt  // ListStrInt 是一个常用的切片类型 []MapStrInt 的别名。 md5:31e1179a1af9c57e
	MapIntAny切片  = []MapIntAny  // ListIntAny 是常用切片类型 []MapIntAny 的别名。 md5:5683714d5eda3cde
	MapIntStr切片  = []MapIntStr  // ListIntStr 是常用切片类型 []MapIntStr 的别名。 md5:383a98d6f17d278d
	MapIntInt切片  = []MapIntInt  // ListIntInt 是一个常用的切片类型 []MapIntInt 的别名。 md5:9fd8565ffb13d5b2
	MapAnyBool切片 = []MapAnyBool // ListAnyBool 是一个常用的切片类型 []MapAnyBool 的别名。 md5:8dcdb92f2ce053c8
	MapStrBool切片 = []MapStrBool // ListStrBool 是常用切片类型 []MapStrBool 的别名。 md5:de56ef83a16ba4a2
	MapIntBool切片 = []MapIntBool // ListIntBool 是常用切片类型 []MapIntBool 的别名。 md5:44e8bb52de3fc953
)

type (
	Slice别名    = []interface{} // Slice 是对常用切片类型 []interface{} 的别名。 md5:9a6ecfa92fac8f3a
	SliceAny别名 = []interface{} // SliceAny 是一个别名，用于频繁使用的切片类型 []interface{}。 md5:0b565fb84a9bac93
	SliceStr别名 = []string      // SliceStr 是经常使用的切片类型 []string 的别名。 md5:19d8a28bce60c7d8
	SliceInt别名 = []int         // SliceInt 是常用切片类型 []int 的别名。 md5:2923a4699e7e68da
)

type (
	X切片    = []interface{} // Array 是一个常用的切片类型 []interface{} 的别名。 md5:83e56770b618f21f
	Any切片 = []interface{} // ArrayAny是常用的切片类型[]interface{}的别名。 md5:6dfa6ee832fcc9f4
	X文本切片 = []string      // ArrayStr 是经常使用的切片类型 []string 的别名。 md5:55e2175bd1bb667d
	X整数切片 = []int         // ArrayInt 是常用切片类型 []int 的别名。 md5:f524bf88e64fcf6e
)
