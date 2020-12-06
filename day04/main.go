package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err.Error())
	}

	passports := ParsePassports(string(input))

	validPassportCount := 0

	for _, passport := range passports {
		if passport.ContainsRequiredKeys() {
			validPassportCount++
		}
	}

	fmt.Printf("Answer 1: the input file contains %v passports containting the necessary keys.\n", validPassportCount)

	validPassportCount = 0

	for _, passport := range passports {
		if passport.IsValid() && passport.ContainsRequiredKeys() {
			validPassportCount++
		}
	}

	fmt.Printf("Answer 2: the input file contains %v valid passports\n", validPassportCount)
}

// ParsePassports creates a slice of passports from our input text.
func ParsePassports(input string) []Passport {
	raw := strings.Split(input, "\n\n")
	var passports []Passport

	for _, p := range raw {
		passports = append(passports, NewPassport(p))
	}

	return passports
}
