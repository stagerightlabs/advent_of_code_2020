package main

import (
	"testing"
)

func TestFindBrokenInstruction(t *testing.T) {
	line, accumulated, original, altered := FindBrokenInstruction(getTestInput())

	if line+1 != 8 {
		t.Errorf("Expected the broken line in the test input to be %v, got %v", 8, line+1)
	}
	if accumulated != 8 {
		t.Errorf("Expected new accumulator value to be %v, got %v", 8, accumulated)
	}
	if original != "jmp" {
		t.Errorf("Expected the instruction to be altered to be %q, got %q", "jmp", original)
	}

	if altered != "nop" {
		t.Errorf("Expected the instruction to be altered to be %q, got %q", "nop", altered)
	}
}

func getTestInput() string {
	return `
nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`
}
