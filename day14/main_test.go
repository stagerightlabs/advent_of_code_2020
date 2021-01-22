package main

import (
	"reflect"
	"testing"
)

func TestVersionOneInstructions(t *testing.T) {
	instructions := ParseInput(getFirstTestInput())
	computer := NewComputer()
	computer.followVersionOneInstructions(instructions)
	sum := computer.Sum()
	want := int64(165)

	if sum != want {
		t.Errorf("Expected the sum to be %v, got %v", want, sum)
	}
}

func TestVersionTwoInstructions(t *testing.T) {
	instructions := ParseInput(getSecondTestInput())
	computer := NewComputer()
	computer.followVersionTwoInstructions(instructions)
	sum := computer.Sum()
	want := int64(208)

	if sum != want {
		t.Errorf("Expected the sum to be %v, got %v", want, sum)
	}
}

func TestVersionTwoMaskedAddresses(t *testing.T) {
	mask := NewMask("mask = 000000000000000000000000000000X1001X")
	address := int64(42)
	got := getVersionTwoMaskedAddresses(address, mask)
	want := []int64{26, 58, 27, 59}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected masked addresses to be %v, got %v", want, got)
	}
}

func getFirstTestInput() string {
	return `
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
`
}

func getSecondTestInput() string {
	return `
mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1
`
}
