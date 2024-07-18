// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

// 包deepcopy使用反射来创建对象的深拷贝。
//
// 此包的维护源代码来自：https://github.com/mohae/deepcopy
// md5:5b2714e1c20589de
package deepcopy

import (
	"reflect"
	"time"
)

// 用于将复制过程委派给类型的接口. md5:1c1a67064d6703f2
type Interface interface {
	DeepCopy() interface{}
}

// Copy 创建一个深度复制的副本，将传递给它的内容返回为一个接口{}。返回的值需要进行类型断言以获取正确的类型。
// md5:e0d2ca231ef6877f
func Copy(src interface{}) interface{} {
	if src == nil {
		return nil
	}

	// 通过类型断言进行复制。 md5:71fbe824c64b5731
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
			original = reflect.ValueOf(src)                // 将接口转换为reflect.Value类型. md5:9201337756a02f33
			dst      = reflect.New(original.Type()).Elem() // 创建与原始类型相同的副本。 md5:5b5c1d4fae5a0eff
		)
		// 递归地复制原始内容。 md5:2beee6938a2c312f
		copyRecursive(original, dst)
		// 返回副本作为接口。 md5:50f142c98dfe40dc
		return dst.Interface()
	}
}

// copyRecursive 实现了接口的实际复制。目前，它支持的类型有限，根据需要添加更多的支持。
// md5:aa40c1acbba074ce
func copyRecursive(original, cpy reflect.Value) {
	// 检查是否实现了deepcopy.Interface接口. md5:52d685857bbf7b9b
	if original.CanInterface() && original.IsValid() && !original.IsZero() {
		if copier, ok := original.Interface().(Interface); ok {
			cpy.Set(reflect.ValueOf(copier.DeepCopy()))
			return
		}
	}

	// 根据原始Kind进行处理. md5:7aba57c6e2dbe8f3
	switch original.Kind() {
	case reflect.Ptr:
		// 获取指针实际指向的值。 md5:89e7bbb6f609f50b
		originalValue := original.Elem()

		// 如果它无效，就返回。 md5:925afc12cf8e9320
		if !originalValue.IsValid() {
			return
		}
		cpy.Set(reflect.New(originalValue.Type()))
		copyRecursive(originalValue, cpy.Elem())

	case reflect.Interface:
		// 如果这是一个空值（nil），则什么都不做. md5:35928cbf0bf1c8f7
		if original.IsNil() {
			return
		}
		// 获取接口的值，而不是指针的值。 md5:b5cb1b00bd69a260
		originalValue := original.Elem()

		// 通过调用Elem()获取值。 md5:bf65094ef26fb870
		copyValue := reflect.New(originalValue.Type()).Elem()
		copyRecursive(originalValue, copyValue)
		cpy.Set(copyValue)

	case reflect.Struct:
		t, ok := original.Interface().(time.Time)
		if ok {
			cpy.Set(reflect.ValueOf(t))
			return
		}
		// 遍历结构体的每个字段并复制它。 md5:9aa258862f5c519c
		for i := 0; i < original.NumField(); i++ {
// 对于给定字段，检查Type的StructField，看StructField.PkgPath是否设置，以确定字段是否已导出。因为CanSet()方法对可设置的字段返回false，我不确定原因是什么。-mohae
// md5:3b6525bda8d105bc
			if original.Type().Field(i).PkgPath != "" {
				continue
			}
			copyRecursive(original.Field(i), cpy.Field(i))
		}

	case reflect.Slice:
		if original.IsNil() {
			return
		}
		// 创建一个新的切片，并复制每个元素。 md5:5c3c5bd2aaf76a00
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
