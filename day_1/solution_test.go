package main

import (
	"testing"
)

func TestSolvePart1(t *testing.T) {
	var tests = []struct {
		testname string
		lines    []string
		want     int
	}{
		{"AOCTestInputUnsorted", []string{
			"3   4",
			"4   3",
			"2   5",
			"1   3",
			"3   9",
			"3   3",
		}, 11},
		{"AOCTestInputSorted", []string{
			"1   3",
			"2   3",
			"3   3",
			"3   4",
			"3   5",
			"4   9",
		}, 11},
		{"RightSliceHasNegativeUnsorted", []string{
			"3   -2",
			"4   3",
			"2   5",
			"1   3",
			"3   9",
			"3   3",
		}, 5},
		{"LeftSliceHasNegativeUnsorted", []string{
			"-3   2",
			"4   3",
			"2   5",
			"1   3",
			"3   9",
			"3   3",
		}, 15},
		{"RightSliceHasNegativeSorted", []string{
			"1   -2",
			"2   3",
			"3   3",
			"3   3",
			"3   5",
			"4   9",
		}, 5},
		{"LeftSliceHasNegativeSorted", []string{
			"-3   2",
			"1   3",
			"2   5",
			"3   3",
			"3   9",
			"4   3",
		}, 15},
		{"BothSlicesHaveNegativeSorted", []string{
			"-3   -5",
			"1   2",
			"2   3",
			"3   3",
			"3   3",
			"4   9",
		}, 5},
		{"BothSlicesHaveNegativeUnsorted", []string{
			"-3   2",
			"4   3",
			"2   -5",
			"1   3",
			"3   9",
			"3   3",
		}, 5},
	}

	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			got := SolvePart1(test.lines)
			if got != test.want {
				t.Errorf("SolvePart1() = %v, want %v", got, test.want)
			}
		})
	}
}
