package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

// Binary represents a binary number as a map of bit flags
type Binary map[int64]bool

// Mask represents an input mask from the instruction file
type Mask map[int64]rune

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err.Error())
	}

	// Part One
	instructions := ParseInput(string(input))
	computer := NewComputer()
	computer.followVersionOneInstructions(instructions)
	sum := computer.Sum()
	fmt.Printf("Answer 1: Sum is %v\n", sum)

	// Part Two
	computer = NewComputer()
	computer.followVersionTwoInstructions(instructions)
	sum = computer.Sum()
	fmt.Printf("Answer 2: Sum is %v\n", sum)

}

// ParseInput retrieves the mask and the instructions from an input string
func ParseInput(input string) []string {
	var instructions []string

	for _, line := range strings.Split(input, "\n") {
		instructions = append(instructions, line)
	}

	return instructions
}

// NewMask creates a bit map from a "mask" input string
func NewMask(input string) Mask {

	parsed := strings.Split(input, " = ")

	if len(parsed[1]) != 36 {
		panic("invalid mask input")
	}

	mask := make(Mask)

	for i, r := range parsed[1] {
		power := int64(math.Pow(2, float64(35-i)))
		mask[power] = r
	}

	return mask
}

// IntegerToBinary converts a number to a binary map
func IntegerToBinary(number int64) Binary {
	binary := make(Binary)

	for i := 35; i >= 0; i-- {
		power := int64(math.Pow(2, float64(i)))

		if number >= power {
			number -= power
			binary[power] = true
		} else {
			binary[power] = false
		}
	}

	return binary
}

// BinaryToInteger converts a binary bit map to an integer
func BinaryToInteger(binary Binary) int64 {
	var value int64 = 0

	for i := 35; i >= 0; i-- {
		power := int64(math.Pow(2, float64(i)))

		used, exists := binary[power]

		if !exists {
			continue
		}

		if used {
			value += power
		}
	}

	return value
}

// ApplyMaskToBinary alters the binary by setting the mask bits to match the mask
func ApplyMaskToBinary(mask Mask, binary Binary) Binary {

	for index, value := range mask {
		switch value {
		case '1':
			binary[index] = true
		case '0':
			binary[index] = false
		default:
			continue
		}
	}

	return binary
}

func (m Mask) string() string {
	str := ""

	for i := 35; i >= 0; i-- {
		power := int64(math.Pow(2, float64(i)))
		str += fmt.Sprintf("%c", m[power])
	}

	return str
}
