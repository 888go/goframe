
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// Package gaes provides useful API for AES encryption/decryption algorithms.
<原文结束>

# <翻译开始>
// gaes 包提供了对AES加密/解密算法有用的API。. md5:a8dd4c4d404f7193
# <翻译结束>


<原文开始>
// IVDefaultValue is the default value for IV.
<原文结束>

# <翻译开始>
// IVDefaultValue 是初始向量(IV)的默认值。. md5:4a6e2194de451335
# <翻译结束>


<原文开始>
// Encrypt is alias of EncryptCBC.
<原文结束>

# <翻译开始>
// Encrypt 是 EncryptCBC 的别名。. md5:d1191baf4cd313b4
# <翻译结束>


<原文开始>
// Decrypt is alias of DecryptCBC.
<原文结束>

# <翻译开始>
// Decrypt 是DecryptCBC的别名。. md5:8cf4ecbfea3f2867
# <翻译结束>


<原文开始>
// EncryptCBC encrypts `plainText` using CBC mode.
// Note that the key must be 16/24/32 bit length.
// The parameter `iv` initialization vector is unnecessary.
<原文结束>

# <翻译开始>
// 使用CBC模式加密`plainText`。
// 注意密钥长度必须为16/24/32位。
// 参数`iv`初始化向量是不必要的。
// md5:1628ebc8e55608ea
# <翻译结束>


<原文开始>
// DecryptCBC decrypts `cipherText` using CBC mode.
// Note that the key must be 16/24/32 bit length.
// The parameter `iv` initialization vector is unnecessary.
<原文结束>

# <翻译开始>
// DecryptCBC 使用CBC模式解密`cipherText`。
// 请注意，密钥必须为16/24/32位长度。
// 参数`iv`初始化向量是可选的。
// md5:ffdc2bd43249f656
# <翻译结束>


<原文开始>
// PKCS5Padding applies PKCS#5 padding to the source byte slice to match the given block size.
//
// If the block size is not provided, it defaults to 8.
<原文结束>

# <翻译开始>
// PKCS5Padding 对源字节切片应用 PKCS#5 填充，以匹配给定的块大小。
//
// 如果未提供块大小，则默认为 8。
// md5:709e406aa572f106
# <翻译结束>


<原文开始>
// PKCS5UnPadding removes PKCS#5 padding from the source byte slice based on the given block size.
//
// If the block size is not provided, it defaults to 8.
<原文结束>

# <翻译开始>
// PKCS5UnPadding 根据给定的块大小，从源字节切片中移除PKCS#5填充。
// 
// 如果未提供块大小，则默认为8。
// md5:d61fd48ac346da68
# <翻译结束>


<原文开始>
// PKCS7Padding applies PKCS#7 padding to the source byte slice to match the given block size.
<原文结束>

# <翻译开始>
// PKCS7Padding 对源字节切片应用PKCS#7填充，以匹配给定的块大小。. md5:c93d69357ddcf364
# <翻译结束>


<原文开始>
// PKCS7UnPadding removes PKCS#7 padding from the source byte slice based on the given block size.
<原文结束>

# <翻译开始>
// PKCS7UnPadding 根据给定的块大小从源字节切片中移除PKCS#7填充。. md5:405becc32a6915c2
# <翻译结束>


<原文开始>
// EncryptCFB encrypts `plainText` using CFB mode.
// Note that the key must be 16/24/32 bit length.
// The parameter `iv` initialization vector is unnecessary.
<原文结束>

# <翻译开始>
// EncryptCFB 使用CFB模式对`plainText`进行加密。
// 注意，密钥必须是16/24/32位长度。
// 参数`iv`（初始化向量）是不必要的。
// md5:cdcc74633b342790
# <翻译结束>


<原文开始>
// DecryptCFB decrypts `plainText` using CFB mode.
// Note that the key must be 16/24/32 bit length.
// The parameter `iv` initialization vector is unnecessary.
<原文结束>

# <翻译开始>
// DecryptCFB 使用CFB模式解密`plainText`。
// 注意，密钥必须是16/24/32位长度。
// 参数`iv`（初始化向量）是不必要的。
// md5:f6a0b1655dd052b7
# <翻译结束>

