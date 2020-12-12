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

	rules := ExtractRulesFromInput(string(input))
	index := createOwnershipIndexFromRules("shiny gold", rules)

	numberOfBagsThatUltimatelyContainShinyGold := 0
	for _, containsShinyGold := range index {
		if containsShinyGold {
			numberOfBagsThatUltimatelyContainShinyGold++
		}
	}

	fmt.Printf("Answer 1: %v bags can ultimately contain a shiny gold bag.\n", numberOfBagsThatUltimatelyContainShinyGold)

	containmentRules := createContainmentCountIndexFromRules(rules)

	fmt.Printf("Answer 2: %v individual bags are required in the shiny gold bag\n", containmentRules["shiny gold"])

}

func createOwnershipIndexFromRules(target string, rules []Rule) map[string]bool {
	ownership := make(map[string]bool)

	for _, rule := range rules {
		if rule.color == target {
			continue
		}

		ownership[rule.color] = doesBagContainTarget(rule.color, target, rules, ownership)
	}

	return ownership
}

func doesBagContainTarget(bag string, target string, rules []Rule, ownership map[string]bool) bool {
	// First we will check the existing ownership array to see
	// if we have already checked this bag.
	already, checked := ownership[bag]
	if already && checked {
		return already
	}

	// Find the rule for this bag type
	rule := findRule(bag, rules)

	// If this bag has no allowed contents we can return false
	if len(rule.contents) == 0 {
		ownership[bag] = false
		return false
	}

	// Go through each sub-bag that can be contained by this bag
	// and check to see if they qualify
	ultimatelyContainsTarget := false
	for subbag := range rule.contents {
		if subbag == target {
			ultimatelyContainsTarget = true
		} else if doesBagContainTarget(subbag, target, rules, ownership) {
			ultimatelyContainsTarget = true
			ownership[subbag] = true
		}
	}

	ownership[bag] = ultimatelyContainsTarget

	return ultimatelyContainsTarget
}

// Find a rule for a given bag in the rules slice
func findRule(bag string, rules []Rule) Rule {
	for _, r := range rules {
		if r.color == bag {
			return r
		}
	}

	panic("Could not find rule for " + bag)
}

func createContainmentCountIndexFromRules(rules []Rule) map[string]int {
	containmentCounts := make(map[string]int)

	for _, rule := range rules {
		containmentCounts[rule.color] = countBagContents(rule.color, rules, containmentCounts)
	}

	return containmentCounts
}

func countBagContents(bag string, rules []Rule, containmentCounts map[string]int) int {

	// Has this bag type already been counted?
	counted, already := containmentCounts[bag]
	if already {
		return counted
	}

	// Fetch the rule for this bag
	rule := findRule(bag, rules)

	// If there are no contents listed this bag contains no other bags
	if len(rule.contents) == 0 {
		containmentCounts[bag] = 0
		return 0
	}

	// Initialize our count for this bag type
	count := 0

	// Cycle through each sub-bag in this rule and count the contents within
	for subbag, number := range rule.contents {
		subcount := countBagContents(subbag, rules, containmentCounts)
		containmentCounts[subbag] = subcount
		count = count + number + (number * subcount)
	}

	return count
}
