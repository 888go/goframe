// 版权所有 GoFrame gf 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一个。

// Package deepcopy 提供了通过反射实现深拷贝的功能。
//
// 该包来源于：https://github.com/mohae/deepcopy
package deepcopy
import (
	"reflect"
	"time"
	)
// 为将复制过程委托给类型定义的接口
type Interface interface {
	DeepCopy() interface{}
}

// Copy 函数接收任意参数，并创建其深度拷贝，然后以 interface{} 类型返回该拷贝。
// 返回的值需要断言为正确的类型。
func Copy(src interface{}) interface{} {
	if src == nil {
		return nil
	}

	// 通过类型断言进行复制。
	switch r := src.(type) {
	case
		int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64,
		complex64, complex128,
		string,
		bool:
		return r

	default:
		if v, ok := src.(Interface); ok {
			return v.DeepCopy()
		}
		var (
			original = reflect.ValueOf(src)                // 将接口转换为 reflect.Value 类型
			dst      = reflect.New(original.Type()).Elem() // 创建一个与原始类型相同的副本。
		)
		// 递归地复制原对象。
		copyRecursive(original, dst)
		// 返回作为接口的副本。
		return dst.Interface()
	}
}

// copyRecursive 实现了接口的实际复制操作。当前，它对所能处理的数据类型的支持有限。根据需要进行添加。
func copyRecursive(original, cpy reflect.Value) {
	// 检查是否实现了 deepcopy.Interface 接口
	if original.CanInterface() && original.IsValid() && !original.IsZero() {
		if copier, ok := original.Interface().(Interface); ok {
			cpy.Set(reflect.ValueOf(copier.DeepCopy()))
			return
		}
	}

	// 根据原始的 Kind 进行处理
	switch original.Kind() {
	case reflect.Ptr:
		// 获取所指向的实际值。
		originalValue := original.Elem()

		// 如果它无效，则返回。
		if !originalValue.IsValid() {
			return
		}
		cpy.Set(reflect.New(originalValue.Type()))
		copyRecursive(originalValue, cpy.Elem())

	case reflect.Interface:
		// 如果这是一个nil，那么什么也不做
		if original.IsNil() {
			return
		}
		// 获取接口的值，而不是指针。
		originalValue := original.Elem()

		// 通过调用 Elem() 获取值。
		copyValue := reflect.New(originalValue.Type()).Elem()
		copyRecursive(originalValue, copyValue)
		cpy.Set(copyValue)

	case reflect.Struct:
		t, ok := original.Interface().(time.Time)
		if ok {
			cpy.Set(reflect.ValueOf(t))
			return
		}
		// 遍历结构体中的每个字段并进行复制。
		for i := 0; i < original.NumField(); i++ {
// 对于给定的字段，检查Type的StructField以查看StructField.PkgPath是否已设置，从而确定该字段是否为导出字段。这是因为CanSet()对于可设置字段会返回false，我不确定具体原因。——mohae
			if original.Type().Field(i).PkgPath != "" {
				continue
			}
			copyRecursive(original.Field(i), cpy.Field(i))
		}

	case reflect.Slice:
		if original.IsNil() {
			return
		}
		// 创建一个新的切片并复制每个元素。
		cpy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i++ {
			copyRecursive(original.Index(i), cpy.Index(i))
		}

	case reflect.Map:
		if original.IsNil() {
			return
		}
		cpy.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			copyValue := reflect.New(originalValue.Type()).Elem()
			copyRecursive(originalValue, copyValue)
			copyKey := Copy(key.Interface())
			cpy.SetMapIndex(reflect.ValueOf(copyKey), copyValue)
		}

	default:
		cpy.Set(original)
	}
}
