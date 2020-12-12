package main

import (
	"strconv"
	"strings"
)

// Instruction represents a single machine instruction
// from our input
type Instruction struct {
	operation string
	argument  int
}

// NewInstruction creates a new Instruction from an input string
func NewInstruction(input string) Instruction {
	instruction := Instruction{}

	// Operation
	instruction.operation = input[:3]
	validOperation := false
	for _, ins := range []string{"acc", "jmp", "nop"} {
		if instruction.operation == ins {
			validOperation = true
		}
	}
	if !validOperation {
		panic("invalid instruction operation")
	}

	// Argument
	argument, err := strconv.Atoi(input[4:])
	if err != nil {
		panic("invalid instruction operation")
	}
	instruction.argument = argument

	return instruction
}

// ExtractInstructionsFromInput creates a slice of []Instruction from a string input
func ExtractInstructionsFromInput(input string) []Instruction {
	instructions := []Instruction{}

	for _, row := range strings.Split(input, "\n") {
		if len(row) == 0 {
			continue
		}

		instructions = append(instructions, NewInstruction(row))
	}

	return instructions
}
