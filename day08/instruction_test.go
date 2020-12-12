package main

import "testing"

func TestNewInstruction(t *testing.T) {
	instruction := NewInstruction("nop +0")

	if instruction.operation != "nop" {
		t.Errorf("Expected operation %q, got %q", "nop", instruction.operation)
	}
	if instruction.argument != 0 {
		t.Errorf("Expected argument to be %v, got %v", 0, instruction.argument)
	}

	instruction = NewInstruction("acc +1")

	if instruction.operation != "acc" {
		t.Errorf("Expected operation %q, got %q", "nop", instruction.operation)
	}
	if instruction.argument != 1 {
		t.Errorf("Expected argument to be %v, got %v", 0, instruction.argument)
	}

	instruction = NewInstruction("jmp -3")

	if instruction.operation != "jmp" {
		t.Errorf("Expected operation %q, got %q", "nop", instruction.operation)
	}
	if instruction.argument != -3 {
		t.Errorf("Expected argument to be %v, got %v", 0, instruction.argument)
	}

	instruction = NewInstruction("acc -99")

	if instruction.operation != "acc" {
		t.Errorf("Expected operation %q, got %q", "nop", instruction.operation)
	}
	if instruction.argument != -99 {
		t.Errorf("Expected argument to be %v, got %v", 0, instruction.argument)
	}
}

func TestExtractInstructionsFromInput(t *testing.T) {
	instructions := ExtractInstructionsFromInput(getTestInput())

	if len(instructions) != 9 {
		t.Errorf("Expected %v instructions in the test input, got %v", 9, len(instructions))
	}
}
