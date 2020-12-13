package main

import (
	"reflect"
	"testing"
)

func TestFindRuleBreakingNumber(t *testing.T) {
	xmas := NewXmas(getTestInput(), 5)

	errant, err := FindRuleBreakingNumber(xmas)

	if err != nil {
		t.Error(err.Error())
	}

	if errant != 127 {
		t.Errorf("Expected errant number to be %v, got %v", 127, errant)
	}
}

func TestFindContiguousSet(t *testing.T) {
	xmas := NewXmas(getTestInput(), 5)
	errant, _ := FindRuleBreakingNumber(xmas)
	got, _ := FindContiguousSet(xmas, errant)
	want := []int{15, 25, 47, 40}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected contiguous set to be %v, got %v", got, want)
	}

	min := sliceMin(got)
	if min != 15 {
		t.Errorf("Expected min slice value to be 15, got %v", min)
	}

	max := sliceMax(got)
	if max != 47 {
		t.Errorf("Expected max slice value to be 15, got %v", max)
	}
}

func getTestInput() string {
	return `
35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`
}
