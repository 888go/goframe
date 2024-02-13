// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 包g提供了常用类型/函数定义以及创建常用对象的配套调用功能。
package g

import (
	"context"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/util/gmeta"
)

type (
	Var  = 泛型类.Var        // Var 是一个通用变量接口，类似于泛型。
	Ctx  = context.Context // Ctx 是 context.Context 类型的别名，通常会被频繁使用。
	Meta = 元数据类.Meta      // Meta 是 gmeta.Meta 这一常用类型的别名。
)

type (
	Map        = map[string]interface{}      // Map 是对常用 map 类型 map[string]interface{} 的别名。
	MapAnyAny  = map[interface{}]interface{} // MapAnyAny 是频繁使用的 map 类型 map[interface{}]interface{} 的别名。
	MapAnyStr  = map[interface{}]string      // MapAnyStr 是频繁使用的 map 类型 map[interface{}]string 的别名。
	MapAnyInt  = map[interface{}]int         // MapAnyInt 是对常用映射类型 map[interface{}]int 的别名。
	MapStrAny  = map[string]interface{}      // MapStrAny 是对常用映射类型 map[string]interface{} 的别名。
	MapStrStr  = map[string]string           // MapStrStr 是频繁使用的 map 类型 map[string]string 的别名。
	MapStrInt  = map[string]int              // MapStrInt 是频繁使用的 map 类型 map[string]int 的别名。
	MapIntAny  = map[int]interface{}         // MapIntAny 是 map[int]interface{} 这一常用映射类型的别名。
	MapIntStr  = map[int]string              // MapIntStr 是频繁使用的 map 类型 map[int]string 的别名。
	MapIntInt  = map[int]int                 // MapIntInt 是 map[int]int 这种常用映射类型的别名。
	MapAnyBool = map[interface{}]bool        // MapAnyBool 是一个频繁使用的 map 类型 map[interface{}]bool 的别名。
	MapStrBool = map[string]bool             // MapStrBool 是频繁使用的 map 类型 map[string]bool 的别名。
	MapIntBool = map[int]bool                // MapIntBool 是对常用 map 类型 map[int]bool 的别名。
)

type (
	Map数组        = []Map        // List 是 []Map 类型的别名，用于表示常用切片类型。
	MapAnyAny数组  = []MapAnyAny  // ListAnyAny 是频繁使用的切片类型 []MapAnyAny 的别名。
	MapAnyStr数组  = []MapAnyStr  // ListAnyStr 是对常用切片类型 []MapAnyStr 的别名。
	MapAnyInt数组  = []MapAnyInt  // ListAnyInt 是 []MapAnyInt 这一常用切片类型的别名。
	MapStrAny数组  = []MapStrAny  // ListStrAny 是对频繁使用的切片类型 []MapStrAny 的别名。
	MapStrStr数组  = []MapStrStr  // ListStrStr 是 []MapStrStr 这种常用切片类型的别名。
	MapStrInt数组  = []MapStrInt  // ListStrInt 是 []MapStrInt 这一常用切片类型的别名。
	MapIntAny数组  = []MapIntAny  // ListIntAny 是频繁使用的切片类型 []MapIntAny 的别名。
	MapIntStr数组  = []MapIntStr  // ListIntStr 是频繁使用的切片类型 []MapIntStr 的别名。
	MapIntInt数组  = []MapIntInt  // ListIntInt 是 []MapIntInt 这种常用切片类型的别名。
	MapAnyBool数组 = []MapAnyBool // ListAnyBool 是 []MapAnyBool 这一常用切片类型的别名。
	MapStrBool数组 = []MapStrBool // ListStrBool 是 []MapStrBool 这种常用切片类型的别名。
	MapIntBool数组 = []MapIntBool // ListIntBool 是频繁使用的切片类型 []MapIntBool 的别名。
)

type (
	Slice别名    = []interface{} // Slice 是 []interface{} 类型的别名，常用于表示频繁使用的切片类型。
	SliceAny别名 = []interface{} // SliceAny 是频繁使用的切片类型 []interface{} 的别名。
	SliceStr别名 = []string      // SliceStr 是对频繁使用的切片类型 []string 的别名。
	SliceInt别名 = []int         // SliceInt 是对常用切片类型 []int 的别名。
)

type (
	X数组    = []interface{} // Array 是 []interface{} 类型的别名，用于表示常用的切片类型。
	Any数组 = []interface{} // ArrayAny 是 []interface{} 这一频繁使用的切片类型的别名。
	X文本数组 = []string      // ArrayStr 是对常用切片类型 []string 的别名。
	X整数数组 = []int         // ArrayInt 是频繁使用的切片类型 []int 的别名。
)
