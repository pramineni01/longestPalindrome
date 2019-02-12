package main

import (
	"sync"
)

func longestPalindrome(s string) string {

	// if empty return same
	if len(s) == 0 {
		return s
	}

	// start with first char as ans
	ans := s[0:1]

	maxLen := len(s)
	// l (as left), r (as right)
	for j := 0; j < maxLen; j++ {
		left, right := j, j

		for ; left >= 0 && right < maxLen && (s[left] == s[right]); left, right = left-1, right+1 {
			var once sync.Once
			once.Do(func() {
				if (right < maxLen-1) && (s[right] == s[right+1]) {
					right = right + 1
				}
				for (left > 0) && (s[left] == s[left-1]) {
					left = left - 1
				}
			})
		}

		if len(ans) < len(s[left+1:right]) {
			ans = s[left+1 : right]
		}
	}

	return ans
}
