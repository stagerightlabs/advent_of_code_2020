package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Rule represents the constraints on a type of luggage
type Rule struct {
	color    string
	contents map[string]int
}

const (
	ruleSeparator = " bags contain"
	noOtherBags   = "no other bags"
)

// ExtractRulesFromInput converts our string input into a slice of Rules
func ExtractRulesFromInput(input string) []Rule {
	rules := []Rule{}

	for _, r := range strings.Split(input, "\n") {
		rule, err := NewRule(r)
		if err == nil {
			rules = append(rules, rule)
		}
	}

	return rules
}

// NewRule creates a Rule from a string in
func NewRule(input string) (Rule, error) {
	rule := Rule{}

	if len(input) == 0 {
		return rule, errors.New("empty rule string")
	}

	// Determine the color of the bag
	color, err := extractBagColorFromRuleString(input)
	if err != nil {
		return rule, err
	}
	rule.color = color

	// Determine the allowed contents
	rule.contents = extractBagContentsFromRuleString(input)

	return rule, nil
}

func extractBagColorFromRuleString(input string) (string, error) {
	index := strings.Index(input, ruleSeparator)
	if index < 0 {
		return "", fmt.Errorf("Cannot parse: %q", input)
	}

	return input[:index], nil
}

func extractBagContentsFromRuleString(input string) map[string]int {
	contents := make(map[string]int)

	index := strings.Index(input, ruleSeparator)
	index = index + len(ruleSeparator)

	substring := strings.TrimSpace(input[index:])
	substring = strings.ReplaceAll(substring, ".", "")

	if substring == noOtherBags {
		return contents
	}

	// convert "1 bright white bag, 2 muted yellow bags."
	// to map[string]int{"bright white": 1, "muted yellow": 2}
	for _, bag := range strings.Split(substring, ",") {
		bag = strings.ReplaceAll(bag, "bags", "")
		bag = strings.ReplaceAll(bag, "bag", "")
		bag = strings.TrimSpace(bag)
		index := strings.Index(bag, " ")
		countString := bag[:index]
		countInt, err := strconv.Atoi(countString)
		if err != nil {
			fmt.Println(err.Error())
			panic("Invalid rule contents: " + bag)
		}
		bag = bag[index+1:]
		contents[bag] = countInt
	}

	return contents
}
