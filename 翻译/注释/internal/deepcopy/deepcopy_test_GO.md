
<原文开始>
// just basic is this working stuff
<原文结束>

# <翻译开始>
// 这只是基本功能是否能正常运行的测试
# <翻译结束>


<原文开始>
// These tests test that all supported basic types are copied correctly.  This
// is done by copying a struct with fields of most of the basic types as []T.
<原文结束>

# <翻译开始>
// 这些测试用于验证所有支持的基本类型都能被正确复制。这是通过将包含多种基本类型的字段结构体以 []T 的形式进行复制来实现的。
# <翻译结束>


<原文开始>
// see if they point to the same location
<原文结束>

# <翻译开始>
// 检查它们是否指向同一位置
# <翻译结束>


<原文开始>
// Go through each field and check to see it got copied properly
<原文结束>

# <翻译开始>
// 遍历每个字段，并检查是否已正确复制
# <翻译结束>







<原文开始>
// the slice headers should point to different data
<原文结束>

# <翻译开始>
// 切片头应指向不同的数据
# <翻译结束>


<原文开始>
// Note: this test will fail until https://github.com/golang/go/issues/15716 is
// fixed and the version it is part of gets released.
<原文结束>

# <翻译开始>
// 注意：在 https://github.com/golang/go/issues/15716 问题得到修复且包含此修复的版本发布之前，该测试将失败。
# <翻译结束>







<原文开始>
// check that the maps point to different locations
<原文结束>

# <翻译开始>
// 检查两个映射是否指向不同的位置
# <翻译结束>


<原文开始>
// test that map keys are deep copied
<原文结束>

# <翻译开始>
// 测试map的键是否被深度复制
# <翻译结束>


<原文开始>
// make sure the lengths are the same
<原文结束>

# <翻译开始>
// 确保长度相同
# <翻译结束>


<原文开始>
	// check that everything was deep copied: since the key is a pointer, we check to
	// see if the pointers are different but the values being pointed to are the same.
<原文结束>

# <翻译开始>
// 检查所有内容是否已深度复制：由于键是指针，我们检查这些指针是否不同，但指向的值是相同的。
# <翻译结束>


<原文开始>
// check that the slices are the same but different
<原文结束>

# <翻译开始>
// 检查两个切片内容相同但指向不同
# <翻译结束>


<原文开始>
// check that Foo points to different locations
<原文结束>

# <翻译开始>
// 检查Foo指向不同的位置
# <翻译结束>







<原文开始>
// not meant to be exhaustive
<原文结束>

# <翻译开始>
// 并不旨在穷举所有情况
# <翻译结束>


<原文开始>
// simple pointer copy
<原文结束>

# <翻译开始>
// 简单指针复制
# <翻译结束>


<原文开始>
// check for nesting values
<原文结束>

# <翻译开始>
// 检查嵌套值
# <翻译结束>

