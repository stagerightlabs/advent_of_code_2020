package main

import (
	"reflect"
	"testing"
)

func TestNewJoltage(t *testing.T) {
	got := ArrangementFromString(getFirstTestInput())
	want := []int{0, 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, 22}

	if !reflect.DeepEqual(got.adapters, want) {
		t.Errorf("Joltage array does not match expectations, got %v, want %v", got, want)
	}
}

func TestDistribution(t *testing.T) {
	arrangement := ArrangementFromString(getFirstTestInput())
	got := Distribution(arrangement)
	want := map[int]int{
		1: 7,
		3: 5,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Joltage distribution does not match expectations, got %v, want %v", got, want)
	}

	arrangement = ArrangementFromString(getSecondTestInput())
	got = Distribution(arrangement)
	want = map[int]int{
		1: 22,
		3: 10,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Joltage distribution does not match expectations, got %v, want %v", got, want)
	}
}

func TestPossibilities(t *testing.T) {
	arrangement := ArrangementFromString(getFirstTestInput())
	possibilities := Possibilities(arrangement)

	if possibilities != 8 {
		t.Errorf("Expected the first input to have 8 possibilities, got %v", possibilities)
	}

	arrangement = ArrangementFromString(getSecondTestInput())
	possibilities = Possibilities(arrangement)

	if possibilities != 19208 {
		t.Errorf("Expected the first input to have 19208 possibilities, got %v", possibilities)
	}
}

// func TestPossibilities(t *testing.T) {
// 	arrangement := ArrangementFromString(getFirstTestInput())
// 	possibilities := Possibilities(arrangement)

// 	if possibilities != 8 {
// 		t.Errorf("Expected the first input string to have 8 possibilities, got %v", possibilities)
// 	}

// 	arrangement = ArrangementFromString(getSecondTestInput())
// 	possibilities = Possibilities(arrangement)

// 	if possibilities != 19208 {
// 		t.Errorf("Expected the first input string to have 19208 possibilities, got %v", possibilities)
// 	}
// }

func getFirstTestInput() string {
	return `
16
10
15
5
1
11
7
19
6
12
4
`
}

func getSecondTestInput() string {
	return `
28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
`
}
