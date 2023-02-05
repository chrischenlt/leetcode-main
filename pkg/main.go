package main

import "fmt"

func main() {

	fmt.Println(+1)
}
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	var res string
	for i := 0; i < len(s); i++ {
		s2 := a(s, i, i)
		if len(s2) > len(res) {
			res = s2
		}
		s3 := a(s, i, i+1)
		if len(s3) > len(res) {
			res = s3
		}
	}
	return res
}

func a(s string, left, right int) string {
	var res string
	for left >= 0 && right < len(s) && s[left] == s[right] {
		res = s[left : right+1]
		left--
		right++
	}
	return res
}
