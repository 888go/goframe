
<原文开始>
// just basic is this working stuff
<原文结束>

# <翻译开始>
// 这只是基本的功能是否正常运作的测试. md5:649d3c44ae9d6089
# <翻译结束>


<原文开始>
// These tests test that all supported basic types are copied correctly.  This
// is done by copying a struct with fields of most of the basic types as []T.
<原文结束>

# <翻译开始>
// 这些测试检查所有支持的基本类型是否正确复制。这是通过复制包含大多数基本类型字段的结构体来实现的，这些字段表示为[]T。
// md5:b02f5e2f90167f01
# <翻译结束>


<原文开始>
// see if they point to the same location
<原文结束>

# <翻译开始>
	// 检查它们是否指向同一位置. md5:cde4a22cfeba39cc
# <翻译结束>


<原文开始>
// Go through each field and check to see it got copied properly
<原文结束>

# <翻译开始>
	// 遍历每个字段，检查它是否被正确复制. md5:d00bdddf588274c3
# <翻译结束>


<原文开始>
// not meant to be exhaustive
<原文结束>

# <翻译开始>
// 不旨在穷尽所有情况. md5:e42da1665679bd48
# <翻译结束>


<原文开始>
// the slice headers should point to different data
<原文结束>

# <翻译开始>
		// 切片的头部应该指向不同的数据. md5:f31a1f39c49f362e
# <翻译结束>


<原文开始>
// Note: this test will fail until https://github.com/golang/go/issues/15716 is
// fixed and the version it is part of gets released.
<原文结束>

# <翻译开始>
// 注意：此测试将在 https://github.com/golang/go/issues/15716 被修复并发布相应的版本之前失败。
// md5:7c10c88146b8e213
# <翻译结束>


<原文开始>
// check that the maps point to different locations
<原文结束>

# <翻译开始>
	// 检查映射是否指向不同的位置. md5:d54b1339f5f533b3
# <翻译结束>


<原文开始>
// test that map keys are deep copied
<原文结束>

# <翻译开始>
	// 测试映射键被深复制. md5:607218635a2fb46a
# <翻译结束>


<原文开始>
// make sure the lengths are the same
<原文结束>

# <翻译开始>
	// 确保长度相同. md5:979294d8e0d545af
# <翻译结束>


<原文开始>
	// check that everything was deep copied: since the key is a pointer, we check to
	// see if the pointers are different but the values being pointed to are the same.
<原文结束>

# <翻译开始>
	// 确保所有内容都进行了深拷贝：由于键是一个指针，我们检查指针是否不同，但被指向的值是否相同。
	// md5:6b3bfbe7f869c100
# <翻译结束>


<原文开始>
// check that the slices are the same but different
<原文结束>

# <翻译开始>
				// 检查切片是否相同但不同. md5:21adab07ec230dcc
# <翻译结束>


<原文开始>
// check that Foo points to different locations
<原文结束>

# <翻译开始>
				// 检查Foo指向不同的位置. md5:a6e3fb097ed803fc
# <翻译结束>


<原文开始>
// check for nesting values
<原文结束>

# <翻译开始>
	// 检查嵌套值. md5:649777c4742c42b6
# <翻译结束>

