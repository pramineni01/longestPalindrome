package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	testcases := [][]string{
		// input, expected
		{"", ""},
		{"b", "b"},
		{"ba", "b"},
		{"bbb", "bbb"},
		{"bbbb", "bbbb"},
		{"bbbbb", "bbbbb"},
		{"bbbbbbbb", "bbbbbbbb"},
		{"bbbbbbbbb", "bbbbbbbbb"},
		{"bab", "bab"},
		{"xbbb", "bbb"},
		{"bbbx", "bbb"},
		{"baab", "baab"},
		{"baxab", "baxab"},
		{"baxyxab", "baxyxab"},
		{"bxcacbbcab", "acbbca"},
	}

	for _, tc := range testcases {
		assert.Equal(t, tc[1], longestPalindrome(tc[0]), "Should be equal.")
	}
}
