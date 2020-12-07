package main

import (
	"regexp"
	"strconv"
	"strings"
)

// Passport represents passport data that has been scanned
type Passport struct {
	BYR string // Birth Year
	IYR string // Issue Year
	EYR string // Expiration Year
	HGT string // Height
	HCL string // Hair Color
	ECL string // Eye Color
	PID string // Passport ID
	CID string // Country ID
}

// ContainsRequiredKeys tells us if a passport contains
// values for all of the required information keys
func (p *Passport) ContainsRequiredKeys() bool {
	return len(p.BYR) > 0 &&
		len(p.IYR) > 0 &&
		len(p.EYR) > 0 &&
		len(p.HGT) > 0 &&
		len(p.HCL) > 0 &&
		len(p.ECL) > 0 &&
		len(p.PID) > 0
}

// IsValid tells us whether all of the information in the
// passport meet the required validation rules
func (p *Passport) IsValid() bool {
	return p.validBirthYear() &&
		p.validIssueYear() &&
		p.validExpirationYear() &&
		p.validHeight() &&
		p.validHairColor() &&
		p.validEyeColor() &&
		p.validPassportID()
}

// Problems returns an array of fields that don't
// meet the validation requirements.
func (p *Passport) Problems() []string {
	problems := []string{}

	if !p.validBirthYear() {
		problems = append(problems, "Invalid Birth Year")
	}

	if !p.validIssueYear() {
		problems = append(problems, "Invalid Issue Year")
	}

	if !p.validExpirationYear() {
		problems = append(problems, "Invalid Expiration Year")
	}

	if !p.validHeight() {
		problems = append(problems, "Invalid Height")
	}

	if !p.validHairColor() {
		problems = append(problems, "Invalid Hair Color")
	}

	if !p.validEyeColor() {
		problems = append(problems, "Invalid Eye Color")
	}

	if !p.validPassportID() {
		problems = append(problems, "Invalid Passport ID")
	}

	return problems
}

// BYR must be four digits; at least 1920 and at most 2002.
func (p *Passport) validBirthYear() bool {
	if len(p.BYR) < 1 || len(p.BYR) > 4 {
		return false
	}

	integer, err := strconv.Atoi(p.BYR)
	if err != nil {
		return false
	}

	return integer >= 1920 && integer < 2003
}

// IYR must be four digits; at least 2010 and at most 2020
func (p *Passport) validIssueYear() bool {
	if len(p.IYR) < 1 || len(p.IYR) > 4 {
		return false
	}

	integer, err := strconv.Atoi(p.IYR)
	if err != nil {
		return false
	}

	return integer >= 2010 && integer < 2021
}

// EYR must be four digits; at least 2020 and at most 2030
func (p *Passport) validExpirationYear() bool {
	if len(p.EYR) < 1 || len(p.EYR) > 4 {
		return false
	}

	integer, err := strconv.Atoi(p.EYR)
	if err != nil {
		return false
	}

	return integer >= 2020 && integer < 2031
}

// HGT must be a number followed by either cm or in:
// - If cm, the number must be at least 150 and at most 193.
// - If in, the number must be at least 59 and at most 76.
func (p *Passport) validHeight() bool {
	if strings.Contains(p.HGT, "cm") {
		integer, err := strconv.Atoi(strings.ReplaceAll(p.HGT, "cm", ""))
		if err != nil {
			return false
		}

		return integer >= 150 && integer < 194
	}

	if strings.Contains(p.HGT, "in") {
		integer, err := strconv.Atoi(strings.ReplaceAll(p.HGT, "in", ""))
		if err != nil {
			return false
		}

		return integer >= 59 && integer < 77
	}

	return false
}

// HCL must be a # followed by exactly six characters 0-9 or a-f.
func (p *Passport) validHairColor() bool {
	if len(p.HCL) < 1 || len(p.HCL) > 7 {
		return false
	}

	if !strings.HasPrefix(p.HCL, "#") {
		return false
	}

	hex := strings.ReplaceAll(p.HCL, "#", "")

	found, err := regexp.MatchString("([^a-f0-9])", hex)

	if err != nil {
		return false
	}

	return !found
}

// ECL must be exactly one of: amb blu brn gry grn hzl oth
func (p *Passport) validEyeColor() bool {
	options := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	found := false

	for _, opt := range options {
		if p.ECL == opt {
			found = true
		}
	}

	return found
}

// PID must be a nine-digit number, including leading zeroes
func (p *Passport) validPassportID() bool {
	if len(p.PID) != 9 {
		return false
	}

	_, err := strconv.Atoi(p.PID)
	if err != nil {
		return false
	}

	return true
}

// NewPassport creates a Passport struct from raw input
func NewPassport(input string) Passport {
	p := Passport{}
	input = strings.ReplaceAll(input, "\n", " ")
	raw := strings.Split(input, " ")

	for _, field := range raw {

		if len(field) == 0 {
			continue
		}

		kv := strings.Split(field, ":")

		key := strings.ToUpper(kv[0])
		value := strings.TrimSpace(kv[1])

		switch key {
		case "BYR":
			p.BYR = value
		case "IYR":
			p.IYR = value
		case "EYR":
			p.EYR = value
		case "HGT":
			p.HGT = value
		case "HCL":
			p.HCL = value
		case "ECL":
			p.ECL = value
		case "PID":
			p.PID = value
		case "CID":
			p.CID = value
		}
	}

	return p
}
