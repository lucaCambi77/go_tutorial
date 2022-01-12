package _test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestMaximumValueAfterInsertion(t *testing.T) {

	var tests = []struct {
		s    string
		x    int
		want string
	}{
		{"89", 9, "989"},
		{"-13", 2, "-123"},
		{"28824579515", 8, "828824579515"},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf("%v %v", tt.s, tt.x)
		t.Run(testcase, func(t *testing.T) {
			ans := maximumValueAfterInsertion(tt.s, tt.x)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func maximumValueAfterInsertion(s string, x int) string {

	if s[0] != '-' {

		for i := 0; i < len(s); i++ {
			if x > int(s[i]-'0') {
				return s[0:i] + strconv.Itoa(x) + s[i:]
			}
		}

	} else {

		for i := 1; i < len(s); i++ {
			if x < int(s[i]-'0') {
				return s[0:i] + strconv.Itoa(x) + s[i:]
			}
		}

	}

	return s + strconv.Itoa(x)
}

// Given a string S, return the number of substrings that have only one distinct letter.
func TestCountLetters(t *testing.T) {

	var tests = []struct {
		input string
		want  int
	}{
		{"aaaba", 8},
		{"aaaaaaaaaa", 55},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf("%v", tt.input)
		t.Run(testcase, func(t *testing.T) {
			ans := countLetters(tt.input)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func countLetters(s string) int {
	count := 1
	sol := 0

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			sol += (count * (count + 1)) / 2
			count = 1
		}
	}
	sol += (count * (count + 1)) / 2

	return sol
}

// Given a binary string s without leading zeros, return true if s contains at most one contiguous segment of ones. Otherwise, return false.
func TestCheckBinaryStringHasAtMostOneSegOnes(t *testing.T) {

	var tests = []struct {
		input string
		want  bool
	}{
		{"1001", false},
		{"110", true},
		{"1", true},
		{"10", true},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf("%v", tt.input)
		t.Run(testcase, func(t *testing.T) {
			ans := checkBinaryStringHasAtMostOneSegOnes(tt.input)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func checkBinaryStringHasAtMostOneSegOnes(s string) bool {

	if strings.Contains(s, "01") {
		return false
	}

	return true
}
