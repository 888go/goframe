
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>

















<原文开始>
// decrypt content too short error
<原文结束>

# <翻译开始>
// 解密内容过短错误
# <翻译结束>







<原文开始>
// PKCS5UnPadding blockSize zero
<原文结束>

# <翻译开始>
// PKCS5UnPadding：对于给定的 blockSize，执行PKCS#5填充解码（即去除填充部分）
// 此函数通常用于对使用PKCS#5填充方式加密的数据进行解密前的填充剥离
# <翻译结束>


<原文开始>
// PKCS5UnPadding src len zero
<原文结束>

# <翻译开始>
// PKCS5UnPadding 解密时去除PKCS5填充, src为待解密的字节切片, len为src的长度
// 当len为0时，表示处理特殊情况，可能由于数据异常或者已无填充
// 在密码学中，PKCS5Padding是对明文进行填充以满足块加密算法对输入长度要求的一种方法。在解密时需要将这些填充的部分去除。上述注释表明这是一个用于在解密后执行PKCS5填充去除操作的函数，并特别指出了当输入长度为0时的情况，这可能意味着原始数据不包含有效的填充或已经不是完整的块。
# <翻译结束>


<原文开始>
// PKCS5UnPadding src len > blockSize
<原文结束>

# <翻译开始>
// PKCS5UnPadding 当原始数据长度大于块大小时
// 这段注释表明这是一个Go语言函数，用于执行PKCS#5填充的解填充操作，前提是源数据的长度大于块大小。在密码学中，PKCS#5填充是一种用于确保明文数据长度适配分组密码（如AES）所需块大小的方法，在解密后需要进行解填充以恢复原始数据。
# <翻译结束>


<原文开始>
// PKCS7UnPadding blockSize zero
<原文结束>

# <翻译开始>
// PKCS7UnPadding 函数，用于移除PKCS7填充
// blockSize 指定块大小（以字节为单位），在解密时需要提供该参数以便正确地移除填充
// ```go
// PKCS7UnPadding func
// 参数:
//   - paddedData: 已经过PKCS7填充的原始数据
//   - blockSize: 块大小，以字节为单位
// func PKCS7UnPadding(paddedData []byte, blockSize int) ([]byte, error) {
    // ...
// }
# <翻译结束>


<原文开始>
// PKCS7UnPadding src len zero
<原文结束>

# <翻译开始>
// PKCS7UnPadding 对PKCS7填充方式进行解码，src为待解码的源数据，len表示源数据长度，需要注意的是这里的源数据长度可能为零。
# <翻译结束>


<原文开始>
// PKCS7UnPadding src len > blockSize
<原文结束>

# <翻译开始>
// PKCS7UnPadding 函数，用于处理src（源数据）的PKCS#7解填充，要求src的长度大于块大小（blockSize）
// 在密码学中，PKCS7Padding是一种常用的对称加密数据填充方式，确保加密后数据长度是块大小的整数倍。当解密时，需要通过PKCS7UnPadding函数去除末尾的填充数据。这段注释说明了该函数的作用和调用前的约束条件——待解填充的数据长度必须大于块大小。
# <翻译结束>


<原文开始>
// go test *.go -bench=".*"
<原文结束>

# <翻译开始>
// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试
# <翻译结束>


<原文开始>
// decrypt content size error
<原文结束>

# <翻译开始>
// 解密内容大小错误
# <翻译结束>

