package main

// Tabulator follows instructions to calculate a total
type Tabulator struct {
	instructions []Instruction
	position     int
	accumulator  int
	history      []int
	completed    bool
}

// Tabulate uses an instruction set to calculate a total
func (t *Tabulator) Tabulate() int {

	for {

		// Have we reached the end of the instruction list?
		if t.position >= len(t.instructions) {
			t.completed = true
			break
		}

		// Check to see if we are executing an instruction for the second time.
		// If so, halt operation.
		if contains(t.position, t.history) {
			break
		}

		t.DoTick()
	}

	return t.accumulator
}

// DoTick executes a single instruction
func (t *Tabulator) DoTick() {
	instruction := t.instructions[t.position]
	t.history = append(t.history, t.position)

	switch instruction.operation {
	case "nop":
		t.position++
	case "acc":
		t.accumulator = t.accumulator + instruction.argument
		t.position++
	case "jmp":
		t.position = t.position + instruction.argument
	}
}

// Determine if a number contained in a slice
func contains(number int, set []int) bool {
	for i := 0; i < len(set); i++ {
		if set[i] == number {
			return true
		}
	}
	return false
}

// NewTabulator creates a tabulator with an instruction set
// generated from the input string
func NewTabulator(instructions []Instruction) Tabulator {
	tabulator := Tabulator{}

	tabulator.instructions = instructions
	tabulator.position = 0
	tabulator.accumulator = 0
	tabulator.completed = false

	return tabulator
}
