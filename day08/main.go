package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err.Error())
	}

	instructions := ExtractInstructionsFromInput(string(input))

	// Part 1
	tabulator := NewTabulator(instructions)
	total := tabulator.Tabulate()
	fmt.Printf("Answer 1: The value in the accumulator is %v\n", total)

	// Part 2
	line, accumulated, original, altered := FindBrokenInstruction(string(input))
	fmt.Println("Answer 2:")
	fmt.Printf("Changing line %v from %q to %q allows the instruction set to complete.\n", line+1, original, altered)
	fmt.Printf("The new accumulator value is %v\n", accumulated)
}

// FindBrokenInstruction cycles through the entire instruction set
// to find the one instruction that can be altered to let
// the tabulation fully complete, rather than running
// in an infinite loop.
func FindBrokenInstruction(input string) (int, int, string, string) {
	instructions := ExtractInstructionsFromInput(input)
	alteredInstructions := []Instruction{}

	for i := 0; i < len(instructions); i++ {
		alteredInstructions = ExtractInstructionsFromInput(input)
		ins := &alteredInstructions[i]

		var original, altered string
		switch ins.operation {
		case "nop":
			original = "nop"
			ins.operation = "jpm"
			altered = "jmp"
		case "jmp":
			original = "jmp"
			ins.operation = "nop"
			altered = "nop"
		}

		tabulator := NewTabulator(alteredInstructions)

		accumulated := tabulator.Tabulate()

		if tabulator.completed {
			return i, accumulated, original, altered
		}
	}

	return 0, 0, "", ""
}
