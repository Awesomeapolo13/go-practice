package longestPalindromicSubstr_test

import (
	"algorithm/longestPalindromicSubstr"
	"testing"
)

var testCases = []struct {
	name     string
	str      string
	expected string
}{
	{"ccd", "ccd", "cc"},
	{"bb", "bb", "bb"},
	{"a", "a", "a"},
	{"cbbd", "cbbd", "bb"},
	{"babad", "babad", "bab"},
}

func TestFindLongestPalindrome(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := longestPalindromicSubstr.FindLongestPalindrome(tc.str)
			if result != tc.expected {
				t.Errorf("TestCase %v failed. Got %v, expected %v", tc.name, result, tc.expected)
			}
		})
	}
}

func TestLongestPalindrome(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := longestPalindromicSubstr.FindLongestPalindromeExpandAroundCenter(tc.str)
			if result != tc.expected {
				t.Errorf("TestCase %v failed. Got %v, expected %v", tc.name, result, tc.expected)
			}
		})
	}
}
