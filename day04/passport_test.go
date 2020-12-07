package main

import "testing"

func TestNewPassport(t *testing.T) {
	passport := NewPassport(`
ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm
`)

	if passport.BYR != "1937" {
		t.Errorf("expected %q to be %q, got %q", "BYR", "1937", passport.BYR)
	}
	if passport.IYR != "2017" {
		t.Errorf("expected %q to be %q, got %q", "IYR", "2017", passport.IYR)
	}
	if passport.EYR != "2020" {
		t.Errorf("expected %q to be %q, got %q", "EYR", "2020", passport.EYR)
	}
	if passport.HGT != "183cm" {
		t.Errorf("expected %q to be %q, got %q", "HGT", "183cm", passport.HGT)
	}
	if passport.HCL != "#fffffd" {
		t.Errorf("expected %q to be %q, got %q", "HCL", "#fffffd", passport.HCL)
	}
	if passport.ECL != "gry" {
		t.Errorf("expected %q to be %q, got %q", "RCL", "gry", passport.ECL)
	}
	if passport.PID != "860033327" {
		t.Errorf("expected %q to be %q, got %q", "PID", "860033327", passport.PID)
	}
	if passport.CID != "147" {
		t.Errorf("expected %q to be %q, got %q", "CID", "147", passport.CID)
	}
}

func TestPassportKeyValidation(t *testing.T) {
	passport := NewPassport(`
ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm
`)
	if !passport.IsValid() {
		t.Errorf("Expected passport #1 to be valid")
	}

	passport = NewPassport(`
iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929
`)
	if passport.IsValid() {
		t.Errorf("Expected passport #2 to be invalid")
	}

	passport = NewPassport(`
hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm
`)
	if !passport.IsValid() {
		t.Errorf("Expected passport #3 to be valid")
	}

	passport = NewPassport(`
hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
`)
	if passport.IsValid() {
		t.Errorf("Expected passport #4 to be invalid")
	}
}

func TestPassportValidationWithKnownInvalids(t *testing.T) {
	passports := ParsePassports(getInvalidPassportSet())

	for _, passport := range passports {
		if passport.IsValid() {
			t.Errorf("Expected %s to be invalid", passport)
		}
	}
}

func getInvalidPassportSet() string {
	return `
eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007

pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1920
hcl:#623z2f
`
}

func TestPassportValidationWithKnownValids(t *testing.T) {
	passports := ParsePassports(getValidPassportSet())

	for _, passport := range passports {
		if !passport.IsValid() {
			t.Errorf("Expected %s to be valid. Problems: %v", passport, passport.Problems())
		}
	}
}

func getValidPassportSet() string {
	return `
pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719
`
}

func TestBYRValidation(t *testing.T) {
	testCases := map[string]bool{
		"2002": true,
		"2003": false,
	}

	for tc, expectation := range testCases {
		p := Passport{BYR: tc}

		if p.validBirthYear() != expectation {
			t.Errorf("Expected case %q to be %v, got %v", tc, expectation, p.validBirthYear())
		}
	}
}

func TestHGTValidation(t *testing.T) {
	testCases := map[string]bool{
		"60in":  true,
		"190cm": true,
		"190in": false,
		"190":   false,
	}

	for tc, expectation := range testCases {
		p := Passport{HGT: tc}

		if p.validHeight() != expectation {
			t.Errorf("Expected case %q to be %v, got %v", tc, expectation, p.validHeight())
		}
	}
}

func TestHCLValidation(t *testing.T) {
	testCases := map[string]bool{
		"#123abc": true,
		"#123abz": false,
		"123abc":  false,
	}

	for tc, expectation := range testCases {
		p := Passport{HCL: tc}

		if p.validHairColor() != expectation {
			t.Errorf("Expected case %q to be %v, got %v", tc, expectation, p.validHairColor())
		}
	}
}

func TestECLValidation(t *testing.T) {
	testCases := map[string]bool{
		"brn": true,
		"wat": false,
	}

	for tc, expectation := range testCases {
		p := Passport{ECL: tc}

		if p.validEyeColor() != expectation {
			t.Errorf("Expected case %q to be %v, got %v", tc, expectation, p.validEyeColor())
		}
	}
}

func TestPIDValidation(t *testing.T) {
	testCases := map[string]bool{
		"000000001":  true,
		"0123456789": false,
		"75827285":   false,
	}

	for tc, expectation := range testCases {
		p := Passport{PID: tc}

		if p.validPassportID() != expectation {
			t.Errorf("Expected case %q to be %v, got %v", tc, expectation, p.validPassportID())
		}
	}
}
