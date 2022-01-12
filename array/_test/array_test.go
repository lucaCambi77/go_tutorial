package _test_test

import (
	"fmt"
	"sort"
	"testing"
)

func TestSmallestEqual(t *testing.T) {

	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{0, 1, 2}, 0},
		{[]int{4, 3, 2, 1}, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, -1},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf("%v", tt.input)
		t.Run(testcase, func(t *testing.T) {
			ans := smallestEqual(tt.input)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

/*
return the smallest index i of nums such that i mod 10 == nums[i]
*/
func smallestEqual(nums []int) int {

	for i, num := range nums {
		if i%10 == num {
			return i
		}
	}

	return -1
}

func TestCanAttendMeetings(t *testing.T) {

	var tests = []struct {
		input [][]int
		want  bool
	}{
		{[][]int{{0, 30}, {5, 10}, {15, 20}}, false},
		{[][]int{{7, 10}, {2, 4}}, true},
		{[][]int{{6, 15}, {13, 20}, {6, 17}}, false},
	}

	for _, tt := range tests {
		testcase := fmt.Sprintf("%v", tt.input)
		t.Run(testcase, func(t *testing.T) {
			ans := canAttendMeetings(tt.input)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}

func canAttendMeetings(intervals [][]int) bool {

	sort.Slice(intervals, func(j, k int) bool {
		return intervals[j][0] < intervals[k][0]
	})

	for i := 1; i < len(intervals)-1; i++ {
		if intervals[i][1] > intervals[i][0] {
			return false
		}

	}

	return true
}
