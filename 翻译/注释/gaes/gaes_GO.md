
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
// Package gaes provides useful API for AES encryption/decryption algorithms.
<原文结束>

# <翻译开始>
// 包gaes提供了AES加密/解密算法的有用API。
# <翻译结束>


<原文开始>
// IVDefaultValue is the default value for IV.
<原文结束>

# <翻译开始>
// IVDefaultValue 是 IV 的默认值。
# <翻译结束>


<原文开始>
// Encrypt is alias of EncryptCBC.
<原文结束>

# <翻译开始>
// Encrypt 是 EncryptCBC 的别名。
# <翻译结束>


<原文开始>
// Decrypt is alias of DecryptCBC.
<原文结束>

# <翻译开始>
// Decrypt 是 DecryptCBC 的别名。
# <翻译结束>


<原文开始>
// EncryptCBC encrypts `plainText` using CBC mode.
// Note that the key must be 16/24/32 bit length.
// The parameter `iv` initialization vector is unnecessary.
<原文结束>

# <翻译开始>
// EncryptCBC 使用CBC模式加密`plainText`。
// 注意，密钥必须为16/24/32位长度。
// 参数`iv`（初始化向量）是不必要的。
# <翻译结束>


<原文开始>
// DecryptCBC decrypts `cipherText` using CBC mode.
// Note that the key must be 16/24/32 bit length.
// The parameter `iv` initialization vector is unnecessary.
<原文结束>

# <翻译开始>
// DecryptCBC 使用CBC模式解密`cipherText`。
// 注意，密钥必须为16/24/32比特长度。
// 参数`iv`初始化向量是不必要的。
# <翻译结束>


<原文开始>
// PKCS5Padding applies PKCS#5 padding to the source byte slice to match the given block size.
//
// If the block size is not provided, it defaults to 8.
<原文结束>

# <翻译开始>
// PKCS5Padding 对源字节切片应用PKCS#5填充，以匹配给定的块大小。
//
// 若未提供块大小，则默认为8。
# <翻译结束>


<原文开始>
// PKCS5UnPadding removes PKCS#5 padding from the source byte slice based on the given block size.
//
// If the block size is not provided, it defaults to 8.
<原文结束>

# <翻译开始>
// PKCS5UnPadding 根据给定的块大小，从源字节切片中移除PKCS#5填充。
//
// 若未提供块大小，则默认为8。
# <翻译结束>


<原文开始>
// PKCS7Padding applies PKCS#7 padding to the source byte slice to match the given block size.
<原文结束>

# <翻译开始>
// PKCS7Padding 对源字节切片应用PKCS#7填充，以匹配给定的块大小。
# <翻译结束>


<原文开始>
// PKCS7UnPadding removes PKCS#7 padding from the source byte slice based on the given block size.
<原文结束>

# <翻译开始>
// PKCS7UnPadding 根据给定的块大小，从源字节切片中移除PKCS#7填充。
# <翻译结束>


<原文开始>
// EncryptCFB encrypts `plainText` using CFB mode.
// Note that the key must be 16/24/32 bit length.
// The parameter `iv` initialization vector is unnecessary.
<原文结束>

# <翻译开始>
// EncryptCFB 使用CFB模式加密`plainText`。
// 注意，密钥必须为16/24/32比特长度。
// 参数`iv`（初始化向量）不是必需的。
# <翻译结束>


<原文开始>
// DecryptCFB decrypts `plainText` using CFB mode.
// Note that the key must be 16/24/32 bit length.
// The parameter `iv` initialization vector is unnecessary.
<原文结束>

# <翻译开始>
// DecryptCFB 使用CFB模式解密`plainText`。
// 注意，密钥必须为16/24/32位长度。
// 参数`iv`初始化向量是不必要的。
# <翻译结束>

