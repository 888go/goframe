// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文本类

// Levenshtein 计算两个字符串之间的 Levenshtein 距离。
// costIns: 定义插入操作的成本。
// costRep: 定义替换操作的成本。
// costDel: 定义删除操作的成本。
// 参考：http://php.net/manual/en/function.levenshtein.php.
func Levenshtein(str1, str2 string, costIns, costRep, costDel int) int {
	var maxLen = 255
	l1 := len(str1)
	l2 := len(str2)
	if l1 == 0 {
		return l2 * costIns
	}
	if l2 == 0 {
		return l1 * costDel
	}
	if l1 > maxLen || l2 > maxLen {
		return -1
	}

	tmp := make([]int, l2+1)
	p1 := make([]int, l2+1)
	p2 := make([]int, l2+1)
	var c0, c1, c2 int
	var i1, i2 int
	for i2 := 0; i2 <= l2; i2++ {
		p1[i2] = i2 * costIns
	}
	for i1 = 0; i1 < l1; i1++ {
		p2[0] = p1[0] + costDel
		for i2 = 0; i2 < l2; i2++ {
			if str1[i1] == str2[i2] {
				c0 = p1[i2]
			} else {
				c0 = p1[i2] + costRep
			}
			c1 = p1[i2+1] + costDel
			if c1 < c0 {
				c0 = c1
			}
			c2 = p2[i2] + costIns
			if c2 < c0 {
				c0 = c2
			}
			p2[i2+1] = c0
		}
		tmp = p1
		p1 = p2
		p2 = tmp
	}
	c0 = p1[l2]

	return c0
}

// SimilarText 计算两个字符串之间的相似度。
// 参考：http://php.net/manual/en/function.similar-text.php.
func X取相似度(文本1, 文本2 string, 百分比 *float64) int {
	var similarText func(string, string, int, int) int
	similarText = func(str1, str2 string, len1, len2 int) int {
		var sum, max int
		pos1, pos2 := 0, 0

		// 在两个字符串中查找相同片段的最长长度
		for i := 0; i < len1; i++ {
			for j := 0; j < len2; j++ {
				for l := 0; (i+l < len1) && (j+l < len2) && (str1[i+l] == str2[j+l]); l++ {
					if l+1 > max {
						max = l + 1
						pos1 = i
						pos2 = j
					}
				}
			}
		}

		if sum = max; sum > 0 {
			if pos1 > 0 && pos2 > 0 {
				sum += similarText(str1, str2, pos1, pos2)
			}
			if (pos1+max < len1) && (pos2+max < len2) {
				s1 := []byte(str1)
				s2 := []byte(str2)
				sum += similarText(string(s1[pos1+max:]), string(s2[pos2+max:]), len1-pos1-max, len2-pos2-max)
			}
		}

		return sum
	}

	l1, l2 := len(文本1), len(文本2)
	if l1+l2 == 0 {
		return 0
	}
	sim := similarText(文本1, 文本2, l1, l2)
	if 百分比 != nil {
		*百分比 = float64(sim*200) / float64(l1+l2)
	}
	return sim
}

// Soundex 计算一个字符串的 soundex 码。
// 参见：http://php.net/manual/en/function.soundex.php.
func X取soundex码(文本 string) string {
	if 文本 == "" {
		panic("str: cannot be an empty string")
	}
	table := [26]rune{
		'0', '1', '2', '3', // A, B, C, D
		'0', '1', '2', // E, F, G
		'0',                          // H
		'0', '2', '2', '4', '5', '5', // I, J, K, L, M, N
		'0', '1', '2', '6', '2', '3', // O, P, Q, R, S, T
		'0', '1', // U, V
		'0', '2', // W, X
		'0', '2', // Y, Z
	}
	last, code, small := -1, 0, 0
	sd := make([]rune, 4)
	// 构建 soundex 字符串
	for i := 0; i < len(文本) && small < 4; i++ {
		// ToUpper
		char := 文本[i]
		if char < '\u007F' && 'a' <= char && char <= 'z' {
			code = int(char - 'a' + 'A')
		} else {
			code = int(char)
		}
		if code >= 'A' && code <= 'Z' {
			if small == 0 {
				sd[small] = rune(code)
				small++
				last = int(table[code-'A'])
			} else {
				code = int(table[code-'A'])
				if code != last {
					if code != 0 {
						sd[small] = rune(code)
						small++
					}
					last = code
				}
			}
		}
	}
	// pad with "0"
	for ; small < 4; small++ {
		sd[small] = '0'
	}
	return string(sd)
}
