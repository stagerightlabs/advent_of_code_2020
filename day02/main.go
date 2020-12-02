package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	scanner := bufio.NewScanner(f)

	var input []string

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	validSledPasswordCount := 0
	for _, entry := range input {
		rule, password := ParseInput(entry)
		if ValidateSledPassword(password, rule) {
			validSledPasswordCount = validSledPasswordCount + 1
		}
	}

	fmt.Printf("Answer 1: %v passwords are valid according to the 'sled' interpretation. \n", validSledPasswordCount)

	validTobogganPasswordCount := 0
	for _, entry := range input {
		rule, password := ParseInput(entry)
		if ValidateTobogganPassword(password, rule) {
			validTobogganPasswordCount = validTobogganPasswordCount + 1
		}
	}

	fmt.Printf("Answer 2: %v passwords are valid according to the 'toboggan' interpretation. \n", validTobogganPasswordCount)
}

// Rule represents a password validation rule
type Rule struct {
	letter rune
	min    int
	max    int
}

// NewRule creates a rule from string input
func NewRule(input string) (Rule, error) {

	// Extract the letter
	split := strings.Split(input, " ")
	letter := []rune(split[1])

	// Extract the allowed span as min and max values
	span := strings.Split(split[0], "-")

	min, err := strconv.Atoi(span[0])
	if err != nil {
		return Rule{}, err
	}

	max, err := strconv.Atoi(span[1])
	if err != nil {
		return Rule{}, err
	}

	return Rule{letter: letter[0], min: min, max: max}, nil
}

// ParseInput separates an input string into a Rule and a password
func ParseInput(input string) (Rule, string) {
	split := strings.Split(input, ":")
	password := strings.TrimSpace(split[1])

	rule, err := NewRule(split[0])
	if err != nil {
		panic(err.Error())
	}

	return rule, password
}

// ValidateSledPassword checks a password string against the provided rule,
// according to the 'sled' interpretation.
func ValidateSledPassword(password string, rule Rule) bool {

	count := 0

	for _, letter := range password {
		if letter == rule.letter {
			count = count + 1
		}
	}

	return count >= rule.min && count <= rule.max
}

// ValidateTobogganPassword checks a password string against the provided rule,
// according to the 'toboggan' interpretation.
func ValidateTobogganPassword(password string, rule Rule) bool {

	first := rune(password[rule.min-1])
	second := rune(password[rule.max-1])

	if first == rule.letter && second == rule.letter {
		return false
	}

	return first == rule.letter || second == rule.letter
}
