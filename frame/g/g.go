// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package g provides commonly used type/function defines and coupled calling for creating commonly used objects.
package g

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/util/gmeta"
)

type (
	Var  = gvar.Var        // Var is a universal variable interface, like generics.
	Ctx  = context.Context // Ctx is alias of frequently-used type context.Context.
	Meta = gmeta.Meta      // Meta is alias of frequently-used type gmeta.Meta.
)

type (
	Map        = map[string]interface{}      // Map is alias of frequently-used map type map[string]interface{}.
	MapAnyAny  = map[interface{}]interface{} // MapAnyAny is alias of frequently-used map type map[interface{}]interface{}.
	MapAnyStr  = map[interface{}]string      // MapAnyStr is alias of frequently-used map type map[interface{}]string.
	MapAnyInt  = map[interface{}]int         // MapAnyInt is alias of frequently-used map type map[interface{}]int.
	MapStrAny  = map[string]interface{}      // MapStrAny is alias of frequently-used map type map[string]interface{}.
	MapStrStr  = map[string]string           // MapStrStr is alias of frequently-used map type map[string]string.
	MapStrInt  = map[string]int              // MapStrInt is alias of frequently-used map type map[string]int.
	MapIntAny  = map[int]interface{}         // MapIntAny is alias of frequently-used map type map[int]interface{}.
	MapIntStr  = map[int]string              // MapIntStr is alias of frequently-used map type map[int]string.
	MapIntInt  = map[int]int                 // MapIntInt is alias of frequently-used map type map[int]int.
	MapAnyBool = map[interface{}]bool        // MapAnyBool is alias of frequently-used map type map[interface{}]bool.
	MapStrBool = map[string]bool             // MapStrBool is alias of frequently-used map type map[string]bool.
	MapIntBool = map[int]bool                // MapIntBool is alias of frequently-used map type map[int]bool.
)

type (
	List        = []Map         //qm:Map数组 cz:List = []Map     // List is alias of frequently-used slice type []Map.
	ListAnyAny  = []MapAnyAny   //qm:MapAnyAny数组 cz:ListAnyAny = []MapAnyAny     // ListAnyAny is alias of frequently-used slice type []MapAnyAny.
	ListAnyStr  = []MapAnyStr   //qm:MapAnyStr数组 cz:ListAnyStr = []MapAnyStr     // ListAnyStr is alias of frequently-used slice type []MapAnyStr.
	ListAnyInt  = []MapAnyInt   //qm:MapAnyInt数组 cz:ListAnyInt = []MapAnyInt     // ListAnyInt is alias of frequently-used slice type []MapAnyInt.
	ListStrAny  = []MapStrAny   //qm:MapStrAny数组 cz:ListStrAny = []MapStrAny     // ListStrAny is alias of frequently-used slice type []MapStrAny.
	ListStrStr  = []MapStrStr   //qm:MapStrStr数组 cz:ListStrStr = []MapStrStr     // ListStrStr is alias of frequently-used slice type []MapStrStr.
	ListStrInt  = []MapStrInt   //qm:MapStrInt数组 cz:ListStrInt = []MapStrInt     // ListStrInt is alias of frequently-used slice type []MapStrInt.
	ListIntAny  = []MapIntAny   //qm:MapIntAny数组 cz:ListIntAny = []MapIntAny     // ListIntAny is alias of frequently-used slice type []MapIntAny.
	ListIntStr  = []MapIntStr   //qm:MapIntStr数组 cz:ListIntStr = []MapIntStr     // ListIntStr is alias of frequently-used slice type []MapIntStr.
	ListIntInt  = []MapIntInt   //qm:MapIntInt数组 cz:ListIntInt = []MapIntInt     // ListIntInt is alias of frequently-used slice type []MapIntInt.
	ListAnyBool = []MapAnyBool  //qm:MapAnyBool数组 cz:ListAnyBool = []MapAnyBool     // ListAnyBool is alias of frequently-used slice type []MapAnyBool.
	ListStrBool = []MapStrBool  //qm:MapStrBool数组 cz:ListStrBool = []MapStrBool     // ListStrBool is alias of frequently-used slice type []MapStrBool.
	ListIntBool = []MapIntBool  //qm:MapIntBool数组 cz:ListIntBool = []MapIntBool     // ListIntBool is alias of frequently-used slice type []MapIntBool.
)

type (
	Slice    = []interface{}  //qm:Slice别名 cz:Slice = []interface{}     // Slice is alias of frequently-used slice type []interface{}.
	SliceAny = []interface{}  //qm:SliceAny别名 cz:SliceAny = []interface{}     // SliceAny is alias of frequently-used slice type []interface{}.
	SliceStr = []string       //qm:SliceStr别名 cz:SliceStr = []string     // SliceStr is alias of frequently-used slice type []string.
	SliceInt = []int          //qm:SliceInt别名 cz:SliceInt = []int     // SliceInt is alias of frequently-used slice type []int.
)

type (
	Array    = []interface{}  //qm:数组 cz:Array = []interface{}     // Array is alias of frequently-used slice type []interface{}.
	ArrayAny = []interface{}  //qm:Any数组 cz:ArrayAny = []interface{}     // ArrayAny is alias of frequently-used slice type []interface{}.
	ArrayStr = []string       //qm:文本数组 cz:ArrayStr = []string     // ArrayStr is alias of frequently-used slice type []string.
	ArrayInt = []int          //qm:整数数组 cz:ArrayInt = []int     // ArrayInt is alias of frequently-used slice type []int.
)
