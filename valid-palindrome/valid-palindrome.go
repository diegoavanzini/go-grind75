package validpalindrome

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	re := regexp.MustCompile(`[^a-zA-Z]+`)
	s = re.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	rightIndex := len(s) - 1
	for leftIndex := 0; leftIndex <= rightIndex; leftIndex++ {
		rightIndex := rightIndex - leftIndex
		if rightIndex == leftIndex {
			return true
		}
		if s[leftIndex] != s[rightIndex] {
			return false
		}
	}
	return true
}
