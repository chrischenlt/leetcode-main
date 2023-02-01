package main

import (
	"strings"
	"unicode"
)

func main() {
	isPalindrome(" ")
}
func isPalindrome(s string) bool {
	if len(s) <= 0 {
		return true
	}

	left, right := 0, len(s)-1
	for left <= right {
		for left <= right && !unicode.IsDigit(rune(s[left])) && !unicode.IsLetter(rune(s[left])) {
			left++
		}
		for left <= right && !unicode.IsDigit(rune(s[right])) && !unicode.IsLetter(rune(s[right])) {
			right--
		}

		if left <= right && strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}
