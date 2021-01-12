package main

import (
	"math"
	"strconv"
	"strings"
)

// Computer represents the port computer system
type Computer struct {
	memory map[int64]int64
	mask   Mask
}

// NewComputer initializes a new memory map
func NewComputer() Computer {
	computer := Computer{}
	computer.memory = make(map[int64]int64)
	return computer
}

// Write to a position in memory, applying a mask to the input.
func (c *Computer) writeWithMask(address int64, value int64) {
	binary := IntegerToBinary(value)
	binary = ApplyMaskToBinary(c.mask, binary)
	c.memory[address] = BinaryToInteger(binary)
}

// Write to a position in memory without using a mask.
func (c *Computer) writeWithoutMask(address int64, value int64) {
	c.memory[address] = value
}

func (c *Computer) followVersionOneInstructions(instructions []string) {
	for _, instruction := range instructions {
		if len(instruction) == 0 {
			continue
		}

		// Parse the instruction
		parsed := strings.Split(instruction, " = ")

		// Is this a "mask" instruction or a "mem" instruction?
		if parsed[0] == "mask" {
			c.mask = NewMask(instruction)
		} else {
			// Get the instruction value
			number, err := strconv.ParseInt(parsed[1], 10, 64)
			if err != nil {
				panic(err.Error())
			}

			// Get the instruction address
			pivot := strings.Index(parsed[0], "]")
			address, err := strconv.ParseInt(parsed[0][4:pivot], 10, 64)
			if err != nil {
				panic(err.Error())
			}

			c.writeWithMask(address, number)
		}
	}
}

func (c *Computer) followVersionTwoInstructions(instructions []string) {
	for _, instruction := range instructions {
		if len(instruction) == 0 {
			continue
		}

		// Parse the instruction
		parsed := strings.Split(instruction, " = ")

		// Is this a "mask" instruction or a "mem" instruction?
		if parsed[0] == "mask" {
			c.mask = NewMask(instruction)
		} else {
			// Get the instruction value
			number, err := strconv.ParseInt(parsed[1], 10, 64)
			if err != nil {
				panic(err.Error())
			}

			// Get the instruction address
			pivot := strings.Index(parsed[0], "]")
			address, err := strconv.ParseInt(parsed[0][4:pivot], 10, 64)
			if err != nil {
				panic(err.Error())
			}

			// Get the addresses dictated by the address mask
			for _, addr := range getVersionTwoMaskedAddresses(address, c.mask) {
				c.writeWithoutMask(addr, number)
			}
		}
	}
}

func getVersionTwoMaskedAddresses(address int64, mask Mask) []int64 {

	// Apply the mask to the address
	binary := IntegerToBinary(address)
	masked := Mask{}
	for index, bit := range mask {
		switch bit {
		case '0':
			if binary[index] {
				masked[index] = '1'
			} else {
				masked[index] = '0'
			}
		case '1':
			masked[index] = '1'
		default:
			masked[index] = bit
		}

	}

	// Use the floating bits to create an array of addresses
	addresses := []int64{0}
	for i := 35; i >= 0; i-- {
		power := int64(math.Pow(2, float64(i)))

		switch masked[power] {
		case '1':
			for index := range addresses {
				addresses[index] += power
			}
		case 'X':
			duplicates := []int64{}
			for _, value := range addresses {
				duplicates = append(duplicates, value+power)
			}
			addresses = append(addresses, duplicates...)
		default:
			continue
		}
	}

	return addresses
}

// Sum returns the total value stored across all memory locations
func (c *Computer) Sum() int64 {
	total := int64(0)

	for _, value := range c.memory {
		total += value
	}

	return total
}
