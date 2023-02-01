package main

func main() {

}

// 示例 1：
//输入：haystack = "sadbutsad", needle = "sad"
//输出：0
//解释："sad" 在下标 0 和 6 处匹配。
//第一个匹配项的下标是 0 ，所以返回 0 。

//示例 2：
//输入：haystack = "leetcode", needle = "leeto"
//输出：-1
//解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。

func strStr(haystack string, needle string) int {
	var BASE = 100000

	needleLen := len(needle)
	if len(haystack) == 0 || needleLen == 0 {
		return -1
	}

	power := 1
	targetCode := 0
	for i := 0; i < needleLen; i++ {
		// 计算所需要的次方数
		power = (power * 31) % BASE
		// 计算目标hashCode
		targetCode = (targetCode*31 + int(needle[i])) % BASE
	}

	hashCode := 0
	for i := 0; i < len(haystack); i++ {
		hashCode = (hashCode*31 + int(haystack[i])) % BASE
		if i < needleLen-1 {
			continue
		}

		if i >= needleLen {
			hashCode -= (int(haystack[i-needleLen]) * power) % BASE
			if hashCode < 0 {
				hashCode += BASE
			}
		}

		if hashCode == targetCode {
			if haystack[i-needleLen+1:i+1] == needle {
				return i - needleLen + 1
			}
		}
	}
	return -1
}

func strStr1(haystack string, needle string) int {

	res := -1
	if len(haystack) <= 0 || len(needle) <= 0 {
		return res
	}

	needleLen := len(needle)
	for i := range haystack {
		if len(haystack)-i < needleLen {
			break
		}

		if haystack[i] == needle[0] {
			equal := true
			for j := 1; j < len(needle); j++ {
				if haystack[j+i] != needle[j] {
					equal = false
					break
				}
			}
			if equal {
				return i
			}
		}
	}
	return res
}
