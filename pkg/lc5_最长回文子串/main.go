package main

import (
	"fmt"
)

func main() {

	palindrome := longestPalindrome("baad")
	fmt.Println(palindrome)

}

// 示例 1：
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。

//示例 2：
//输入：s = "cbbd"
//输出："bb"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	newS := make([]rune, 0)
	newS = append(newS, '#')
	for _, c := range s {
		newS = append(newS, c)
		newS = append(newS, '#')
	}
	// dp[i]:    以预处理字符串下标 i 为中心的回文半径(奇数长度时不包括中心)
	// maxRight: 通过中心扩散的方式能够扩散的最右边的下标
	// center:   与 maxRight 对应的中心字符的下标
	// maxLen:   记录最长回文串的半径
	// begin:    记录最长回文串在起始串 s 中的起始下标
	dp, maxRight, center, maxLen, begin := make([]int, len(newS)), 0, 0, 1, 0
	for i := 0; i < len(newS); i++ {
		if i < maxRight {
			// 这一行代码是 Manacher 算法的关键所在
			dp[i] = min(maxRight-i, dp[2*center-i])
		}
		// 中心扩散法更新 dp[i]
		left, right := i-(1+dp[i]), i+(1+dp[i])
		for left >= 0 && right < len(newS) && newS[left] == newS[right] {
			dp[i]++
			left--
			right++
		}
		// 更新 maxRight，它是遍历过的 i 的 i + dp[i] 的最大者
		if i+dp[i] > maxRight {
			maxRight = i + dp[i]
			center = i
		}
		// 记录最长回文子串的长度和相应它在原始字符串中的起点
		if dp[i] > maxLen {
			maxLen = dp[i]
			begin = (i - maxLen) / 2 // 这里要除以 2 因为有我们插入的辅助字符 #
		}
	}
	return s[begin : begin+maxLen]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func longestPalindrome1(s string) string {

	if s == "" {
		return ""
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		curLen := end - start
		// add
		left, right := a(s, i, i, curLen)
		if curLen < right-left {
			start = left
			end = right
		}

		// even
		left, right = a(s, i, i+1, curLen)
		if curLen < right-left {
			start = left
			end = right
		}
	}
	return s[start : end+1]
}

func a(s string, left, right, curLength int) (start, end int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		if curLength < right-left {
			start = left
			end = right
		}
		left--
		right++
		continue
	}
	return
}
