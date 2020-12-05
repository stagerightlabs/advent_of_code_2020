package main

import "testing"

func TestCountTrees(t *testing.T) {
	grid := getTestCaseGrid()

	testCases := [][]int{
		{1, 1, 2},
		{3, 1, 7},
		{5, 1, 3},
		{7, 1, 4},
		{1, 2, 2},
	}

	for _, testCase := range testCases {
		got := CountTrees(grid, testCase[0], testCase[1])
		want := testCase[2]

		if got != want {
			t.Errorf("For case %v, expected a tree count of %v, got %v", testCase, want, got)
		}
	}

}
