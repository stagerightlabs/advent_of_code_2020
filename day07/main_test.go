package main

import (
	"testing"
)

func TestOwnershipIndex(t *testing.T) {
	rules := ExtractRulesFromInput(getFirstTestInput())
	index := createOwnershipIndexFromRules("shiny gold", rules)

	count := 0

	for _, containsShinyGold := range index {
		if containsShinyGold {
			count++
		}
	}

	if count != 4 {
		t.Errorf("Expected 4 bags that can contain shiny gold bags, got %v", count)
	}
}

func TestBagContainmentCount(t *testing.T) {
	rules := ExtractRulesFromInput(getSecondTestInput())
	containmentCounts := createContainmentCountIndexFromRules(rules)

	if containmentCounts["shiny gold"] != 126 {
		t.Errorf("Expected %q to contain 126 bags, got %v", "shiny gold", containmentCounts["shiny gold"])
	}
}
