// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 泛型类_test

import (
	"fmt"

	gvar "github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/json"
)

// New
func ExampleVarNew() {
	v := gvar.X创建(400)
	fmt.Println(v)

	// Output:
	// 400
}

// Clone
func ExampleVar_Clone() {
	tmp := "fisrt hello"
	v := gvar.X创建(tmp)
	g.X调试输出并带类型(v.X浅拷贝())
	fmt.Println(v == v.X浅拷贝())

	// Output:
	// *泛型类.Var(11) "fisrt hello"
	// false
}

// Set
func ExampleVar_Set() {
	var v = gvar.X创建(100.00)
	g.X调试输出(v.X设置值(200.00))
	g.X调试输出(v)

	// Output:
	// 100
	// "200"
}

// Val
func ExampleVar_Val() {
	var v = gvar.X创建(100.00)
	g.X调试输出并带类型(v.X取值())

	// Output:
	// float64(100)
}

// Interface
func ExampleVar_Interface() {
	var v = gvar.X创建(100.00)
	g.X调试输出并带类型(v.Interface())

	// Output:
	// float64(100)
}

// Bytes
func ExampleVar_Bytes() {
	var v = gvar.X创建("GoFrame")
	g.X调试输出并带类型(v.X取字节集())

	// Output:
	// []byte(7) "GoFrame"
}

// String
func ExampleVar_String() {
	var v = gvar.X创建("GoFrame")
	g.X调试输出并带类型(v.String())

	// Output:
	// string(7) "GoFrame"
}

// Bool
func ExampleVar_Bool() {
	var v = gvar.X创建(true)
	g.X调试输出并带类型(v.X取布尔())

	// Output:
	// bool(true)
}

// Int
func ExampleVar_Int() {
	var v = gvar.X创建(-1000)
	g.X调试输出并带类型(v.X取整数())

	// Output:
	// int(-1000)
}

// Uint
func ExampleVar_Uint() {
	var v = gvar.X创建(1000)
	g.X调试输出并带类型(v.X取正整数())

	// Output:
	// uint(1000)
}

// Float32
func ExampleVar_Float32() {
	var price = gvar.X创建(100.00)
	g.X调试输出并带类型(price.X取小数32位())

	// Output:
	// float32(100)
}

// Time
func ExampleVar_Time() {
	var v = gvar.X创建("2021-11-11 00:00:00")
	g.X调试输出并带类型(v.X取时间类())

	// Output:
	// time.Time(29) "2021-11-11 00:00:00 +0800 CST"
}

// GTime
func ExampleVar_GTime() {
	var v = gvar.X创建("2021-11-11 00:00:00")
	g.X调试输出并带类型(v.X取gtime时间类())

	// Output:
	// *时间类.Time(19) "2021-11-11 00:00:00"
}

// Duration
func ExampleVar_Duration() {
	var v = gvar.X创建("300s")
	g.X调试输出并带类型(v.X取时长())

	// Output:
	// time.Duration(4) "5m0s"
}

// MarshalJSON
func ExampleVar_MarshalJSON() {
	testMap := g.Map{
		"code":  "0001",
		"name":  "Golang",
		"count": 10,
	}

	var v = gvar.X创建(testMap)
	res, err := json.Marshal(&v)
	if err != nil {
		panic(err)
	}
	g.X调试输出并带类型(res)

	// Output:
	// []byte(42) "{"code":"0001","count":10,"name":"Golang"}"
}

// UnmarshalJSON
func ExampleVar_UnmarshalJSON() {
	tmp := []byte(`{
	     "Code":          "0003",
	     "Name":          "Golang Book3",
	     "Quantity":      3000,
	     "Price":         300,
	     "OnSale":        true
	}`)
	var v = gvar.X创建(map[string]interface{}{})
	if err := json.Unmarshal(tmp, &v); err != nil {
		panic(err)
	}

	g.X调试输出(v)

	// Output:
	// "{\"Code\":\"0003\",\"Name\":\"Golang Book3\",\"OnSale\":true,\"Price\":300,\"Quantity\":3000}"
}

// UnmarshalValue
func ExampleVar_UnmarshalValue() {
	tmp := g.Map{
		"code":  "00002",
		"name":  "GoFrame",
		"price": 100,
		"sale":  true,
	}

	var v = gvar.X创建(map[string]interface{}{})
	if err := v.UnmarshalValue(tmp); err != nil {
		panic(err)
	}
	g.X调试输出(v)

	// Output:
	// "{\"code\":\"00002\",\"name\":\"GoFrame\",\"price\":100,\"sale\":true}"
}

// IsNil
func ExampleVar_IsNil() {
	g.X调试输出(gvar.X创建(0).X是否为Nil())
	g.X调试输出(gvar.X创建(0.1).X是否为Nil())
	// true
	g.X调试输出(gvar.X创建(nil).X是否为Nil())
	g.X调试输出(gvar.X创建("").X是否为Nil())

	// Output:
	// false
	// false
	// true
	// false
}

// IsEmpty
func ExampleVar_IsEmpty() {
	g.X调试输出(gvar.X创建(0).X是否为空())
	g.X调试输出(gvar.X创建(nil).X是否为空())
	g.X调试输出(gvar.X创建("").X是否为空())
	g.X调试输出(gvar.X创建(g.Map{"k": "v"}).X是否为空())

	// Output:
	// true
	// true
	// true
	// false
}

// IsInt
func ExampleVar_IsInt() {
	g.X调试输出(gvar.X创建(0).X是否为整数())
	g.X调试输出(gvar.X创建(0.1).X是否为整数())
	g.X调试输出(gvar.X创建(nil).X是否为整数())
	g.X调试输出(gvar.X创建("").X是否为整数())

	// Output:
	// true
	// false
	// false
	// false
}

// IsUint
func ExampleVar_IsUint() {
	g.X调试输出(gvar.X创建(0).X是否为正整数())
	g.X调试输出(gvar.X创建(uint8(8)).X是否为正整数())
	g.X调试输出(gvar.X创建(nil).X是否为正整数())

	// Output:
	// false
	// true
	// false
}

// IsFloat
func ExampleVar_IsFloat() {
	g.X调试输出(g.X泛型类(uint8(8)).X是否为小数())
	g.X调试输出(g.X泛型类(float64(8)).X是否为小数())
	g.X调试输出(g.X泛型类(0.1).X是否为小数())

	// Output:
	// false
	// true
	// true
}

// IsSlice
func ExampleVar_IsSlice() {
	g.X调试输出(g.X泛型类(0).X是否为切片())
	g.X调试输出(g.X泛型类(g.Slice别名{0}).X是否为切片())

	// Output:
	// false
	// true
}

// IsMap
func ExampleVar_IsMap() {
	g.X调试输出(g.X泛型类(0).X是否为Map())
	g.X调试输出(g.X泛型类(g.Map{"k": "v"}).X是否为Map())
	g.X调试输出(g.X泛型类(g.Slice别名{}).X是否为Map())

	// Output:
	// false
	// true
	// false
}

// IsStruct
func ExampleVar_IsStruct() {
	g.X调试输出(g.X泛型类(0).X是否为结构())
	g.X调试输出(g.X泛型类(g.Map{"k": "v"}).X是否为结构())

	a := struct{}{}
	g.X调试输出(g.X泛型类(a).X是否为结构())
	g.X调试输出(g.X泛型类(&a).X是否为结构())

	// Output:
	// false
	// false
	// true
	// true
}

// ListItemValues
func ExampleVar_ListItemValues() {
	var goods1 = g.Map切片{
		g.Map{"id": 1, "price": 100.00},
		g.Map{"id": 2, "price": 0},
		g.Map{"id": 3, "price": nil},
	}
	var v = gvar.X创建(goods1)
	fmt.Println(v.X取结构切片或Map切片值("id"))
	fmt.Println(v.X取结构切片或Map切片值("price"))

	// Output:
	// [1 2 3]
	// [100 0 <nil>]
}

// ListItemValuesUnique
func ExampleVar_ListItemValuesUnique() {
	var (
		goods1 = g.Map切片{
			g.Map{"id": 1, "price": 100.00},
			g.Map{"id": 2, "price": 100.00},
			g.Map{"id": 3, "price": nil},
		}
		v = gvar.X创建(goods1)
	)

	fmt.Println(v.X取结构切片或Map切片值并去重("id"))
	fmt.Println(v.X取结构切片或Map切片值并去重("price"))

	// Output:
	// [1 2 3]
	// [100 <nil>]
}

func ExampleVar_Struct() {
	params1 := g.Map{
		"uid":  1,
		"Name": "john",
	}
	v := gvar.X创建(params1)
	type tartget struct {
		Uid  int
		Name string
	}
	t := new(tartget)
	if err := v.Struct(&t); err != nil {
		panic(err)
	}
	g.X调试输出(t)

	// Output:
	// {
	//     Uid:  1,
	//     Name: "john",
	// }
}

func ExampleVar_Structs() {
	paramsArray := []g.Map{}
	params1 := g.Map{
		"uid":  1,
		"Name": "golang",
	}
	params2 := g.Map{
		"uid":  2,
		"Name": "java",
	}

	paramsArray = append(paramsArray, params1, params2)
	v := gvar.X创建(paramsArray)
	type tartget struct {
		Uid  int
		Name string
	}
	var t []tartget
	if err := v.Structs(&t); err != nil {
		panic(err)
	}
	g.X调试输出并带类型(t)

	// Output:
	// []泛型类_test.tartget(2) [
	//     泛型类_test.tartget(2) {
	//         Uid:  int(1),
	//         Name: string(6) "golang",
	//     },
	//     泛型类_test.tartget(2) {
	//         Uid:  int(2),
	//         Name: string(4) "java",
	//     },
	// ]
}

// Ints
func ExampleVar_Ints() {
	var (
		arr = []int{1, 2, 3, 4, 5}
		obj = gvar.X创建(arr)
	)

	fmt.Println(obj.X取整数切片())

	// Output:
	// [1 2 3 4 5]
}

// Int64s
func ExampleVar_Int64s() {
	var (
		arr = []int64{1, 2, 3, 4, 5}
		obj = gvar.X创建(arr)
	)

	fmt.Println(obj.X取整数64位切片())

	// Output:
	// [1 2 3 4 5]
}

// Uints
func ExampleVar_Uints() {
	var (
		arr = []uint{1, 2, 3, 4, 5}
		obj = gvar.X创建(arr)
	)
	fmt.Println(obj.X取正整数切片())

	// Output:
	// [1 2 3 4 5]
}

// Uint64s
func ExampleVar_Uint64s() {
	var (
		arr = []uint64{1, 2, 3, 4, 5}
		obj = gvar.X创建(arr)
	)

	fmt.Println(obj.X取正整数64位切片())

	// Output:
	// [1 2 3 4 5]
}

// Floats
func ExampleVar_Floats() {
	var (
		arr = []float64{1, 2, 3, 4, 5}
		obj = gvar.X创建(arr)
	)

	fmt.Println(obj.X取小数切片())

	// Output:
	// [1 2 3 4 5]
}

// Float32s
func ExampleVar_Float32s() {
	var (
		arr = []float32{1, 2, 3, 4, 5}
		obj = gvar.X创建(arr)
	)

	fmt.Println(obj.X取小数32位切片())

	// Output:
	// [1 2 3 4 5]
}

// Float64s
func ExampleVar_Float64s() {
	var (
		arr = []float64{1, 2, 3, 4, 5}
		obj = gvar.X创建(arr)
	)

	fmt.Println(obj.X取小数64位切片())

	// Output:
	// [1 2 3 4 5]
}

// Strings
func ExampleVar_Strings() {
	var (
		arr = []string{"GoFrame", "Golang"}
		obj = gvar.X创建(arr)
	)
	fmt.Println(obj.X取文本切片())

	// Output:
	// [GoFrame Golang]
}

// Interfaces
func ExampleVar_Interfaces() {
	var (
		arr = []string{"GoFrame", "Golang"}
		obj = gvar.X创建(arr)
	)

	fmt.Println(obj.X取any切片())

	// Output:
	// [GoFrame Golang]
}

// Slice
func ExampleVar_Slice() {
	var (
		arr = []string{"GoFrame", "Golang"}
		obj = gvar.X创建(arr)
	)

	fmt.Println(obj.Slice别名())

	// Output:
	// [GoFrame Golang]
}

// Array
func ExampleVar_Array() {
	var (
		arr = []string{"GoFrame", "Golang"}
		obj = gvar.X创建(arr)
	)
	fmt.Println(obj.Array别名())

	// Output:
	// [GoFrame Golang]
}

// Vars
func ExampleVar_Vars() {
	var (
		arr = []string{"GoFrame", "Golang"}
		obj = gvar.X创建(arr)
	)

	fmt.Println(obj.X取泛型类切片())

	// Output:
	// [GoFrame Golang]
}

// Map
func ExampleVar_Map() {
	var (
		m   = g.Map{"id": 1, "price": 100.00}
		v   = gvar.X创建(m)
		res = v.X取Map()
	)

	fmt.Println(res["id"], res["price"])

	// Output:
	// 1 100
}

// MapStrAny
func ExampleVar_MapStrAny() {
	var (
		m1 = g.Map{"id": 1, "price": 100}
		v  = gvar.X创建(m1)
		v2 = v.X取MapStrAny()
	)

	fmt.Println(v2["price"], v2["id"])

	// Output:
	// 100 1
}

// MapStrStr
func ExampleVar_MapStrStr() {
	var (
		m1 = g.Map{"id": 1, "price": 100}
		v  = gvar.X创建(m1)
		v2 = v.X取文本Map()
	)

	fmt.Println(v2["price"] + "$")

	// Output:
	// 100$
}

// MapStrVar
func ExampleVar_MapStrVar() {
	var (
		m1 = g.Map{"id": 1, "price": 100}
		v  = gvar.X创建(m1)
		v2 = v.X取泛型类Map()
	)

	fmt.Println(v2["price"].X取小数64位() * 100)

	// Output:
	// 10000
}

// MapDeep
func ExampleVar_MapDeep() {
	var (
		m1 = g.Map{"id": 1, "price": 100}
		m2 = g.Map{"product": m1}
		v  = gvar.X创建(m2)
		v2 = v.MapDeep弃用()
	)

	fmt.Println(v2["product"])

	// Output:
	// map[id:1 price:100]
}

// MapStrStrDeep
func ExampleVar_MapStrStrDeep() {
	var (
		m1 = g.Map{"id": 1, "price": 100}
		m2 = g.Map{"product": m1}
		v  = gvar.X创建(m2)
		v2 = v.MapStrStrDeep弃用()
	)

	fmt.Println(v2["product"])

	// Output:
	// {"id":1,"price":100}
}

// MapStrVarDeep
func ExampleVar_MapStrVarDeep() {
	var (
		m1 = g.Map{"id": 1, "price": 100}
		m2 = g.Map{"product": m1}
		m3 = g.Map{}
		v  = gvar.X创建(m2)
		v2 = v.X取泛型类Map_递归()
		v3 = gvar.X创建(m3).X取泛型类Map_递归()
	)

	fmt.Println(v2["product"])
	fmt.Println(v3)

	// Output:
	// {"id":1,"price":100}
	// map[]
}

// Maps
func ExampleVar_Maps() {
	var m = gvar.X创建(g.MapIntInt切片{g.MapIntInt{0: 100, 1: 200}, g.MapIntInt{0: 300, 1: 400}})
	fmt.Printf("%#v", m.X取Map切片())

	// Output:
	// []map[string]interface {}{map[string]interface {}{"0":100, "1":200}, map[string]interface {}{"0":300, "1":400}}
}

// MapsDeep
func ExampleVar_MapsDeep() {
	var (
		p1 = g.MapStrAny{"product": g.Map{"id": 1, "price": 100}}
		p2 = g.MapStrAny{"product": g.Map{"id": 2, "price": 200}}
		v  = gvar.X创建(g.MapStrAny切片{p1, p2})
		v2 = v.MapsDeep弃用()
	)

	fmt.Printf("%#v", v2)

	// Output:
	// []map[string]interface {}{map[string]interface {}{"product":map[string]interface {}{"id":1, "price":100}}, map[string]interface {}{"product":map[string]interface {}{"id":2, "price":200}}}
}

// MapToMap
func ExampleVar_MapToMap() {
	var (
		m1 = gvar.X创建(g.MapIntInt{0: 100, 1: 200})
		m2 = g.MapStrStr{}
	)

	err := m1.MapToMap(&m2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", m2)

	// Output:
	// map[string]string{"0":"100", "1":"200"}
}

// MapToMaps
func ExampleVar_MapToMaps() {
	var (
		p1 = g.MapStrAny{"product": g.Map{"id": 1, "price": 100}}
		p2 = g.MapStrAny{"product": g.Map{"id": 2, "price": 200}}
		v  = gvar.X创建(g.MapStrAny切片{p1, p2})
		v2 []g.MapStrStr
	)

	err := v.MapToMaps(&v2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", v2)

	// Output:
	// []map[string]string{map[string]string{"product":"{\"id\":1,\"price\":100}"}, map[string]string{"product":"{\"id\":2,\"price\":200}"}}
}

// MapToMapsDeep
func ExampleVar_MapToMapsDeep() {
	var (
		p1 = g.MapStrAny{"product": g.Map{"id": 1, "price": 100}}
		p2 = g.MapStrAny{"product": g.Map{"id": 2, "price": 200}}
		v  = gvar.X创建(g.MapStrAny切片{p1, p2})
		v2 []g.MapStrStr
	)

	err := v.MapToMapsDeep(&v2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", v2)

	// Output:
	// []map[string]string{map[string]string{"product":"{\"id\":1,\"price\":100}"}, map[string]string{"product":"{\"id\":2,\"price\":200}"}}
}

// Scan
func ExampleVar_Scan() {
	type Student struct {
		Id     *g.Var
		Name   *g.Var
		Scores *g.Var
	}
	var (
		s Student
		m = g.Map{
			"Id":     1,
			"Name":   "john",
			"Scores": []int{100, 99, 98},
		}
	)
	v := gvar.X创建(m)
	if err := v.X取结构体指针(&s); err == nil {
		g.X调试输出并带类型(s)
	}

	// Output:
	// 泛型类_test.Student(3) {
	//     Id:     *泛型类.Var(1) "1",
	//     Name:   *泛型类.Var(4) "john",
	//     Scores: *泛型类.Var(11) "[100,99,98]",
	// }
}
