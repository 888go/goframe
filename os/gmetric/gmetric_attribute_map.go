// 版权所有 (c) GoFrame (https://goframe.org)，保留所有权利。
//
// 本源代码遵循MIT许可协议。若未随此文件分发MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:c14c707c81272457

package gmetric

// AttributeMap 是一个属性键值对的映射，便于轻松过滤。. md5:77c3a3707b06188f
type AttributeMap map[string]any

// Sets 将给定的属性映射添加到当前映射中。. md5:a21a166cb89d74a0
func (m AttributeMap) Sets(attrMap map[string]any) {
	for k, v := range attrMap {
		m[k] = v
	}
}

// Pick根据给定的属性键选择和返回属性。. md5:345fd91cbaf4ea48
func (m AttributeMap) Pick(keys ...string) Attributes {
	var attrs = make(Attributes, 0)
	for _, key := range keys {
		value, ok := m[key]
		if !ok {
			continue
		}
		attrs = append(attrs, NewAttribute(key, value))
	}
	return attrs
}

// PickEx 选取并返回那些属性键不在给定`keys`中的属性。. md5:cf773d7747da56d7
func (m AttributeMap) PickEx(keys ...string) Attributes {
	var (
		exKeyMap = make(map[string]struct{})
		attrs    = make(Attributes, 0)
	)
	for _, key := range keys {
		exKeyMap[key] = struct{}{}
	}
	for k, v := range m {
		_, ok := exKeyMap[k]
		if ok {
			continue
		}
		attrs = append(attrs, NewAttribute(k, v))
	}
	return attrs
}
