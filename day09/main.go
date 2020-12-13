package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err.Error())
	}

	xmas := NewXmas(string(input), 25)
	errant, _ := FindRuleBreakingNumber(xmas)
	fmt.Printf("Answer 1: The rule breaking number is %v\n", errant)

	subset, _ := FindContiguousSet(xmas, errant)
	min := sliceMin(subset)
	max := sliceMax(subset)
	fmt.Printf("Answer 2: The encryption weakness is %v\n", min+max)

}

// FindRuleBreakingNumber finds the number in our
// array that does not meet the expectation that
// the previous two numbers in the array, when
// summed, will equal that number
func FindRuleBreakingNumber(xmas Xmas) (int, error) {

	for i := xmas.consider; i < len(xmas.arr); i++ {

		length := i - xmas.consider
		consideration := xmas.arr[length:i]
		number := xmas.arr[i]
		found := false

		for j := 0; j < len(consideration); j++ {
			for k := 0; k < len(consideration); k++ {
				if consideration[j] == consideration[k] {
					continue
				}

				if consideration[k]+consideration[j] == number {
					found = true
				}
			}
		}

		if !found {
			return number, nil
		}

	}

	return 0, errors.New("could not find a rule breaking number")
}

// FindContiguousSet finds a subslice of ints in xmas.arr that
// sub to equal a target number.
func FindContiguousSet(xmas Xmas, target int) ([]int, error) {
	for index := range xmas.arr {

		for offset := index; offset < len(xmas.arr); offset++ {
			subset := xmas.arr[index:offset]
			total := sliceTotal(subset)

			if total == target {
				return subset, nil
			}

			if total > target {
				break
			}
		}
	}

	return []int{}, fmt.Errorf("Could not find a contiguous subset with a total of %v", target)
}

func sliceTotal(list []int) int {
	total := 0
	for _, number := range list {
		total += number
	}

	return total
}

func sliceMin(list []int) int {
	min := list[0]
	for _, number := range list {
		if number < min {
			min = number
		}
	}
	return min
}

func sliceMax(list []int) int {
	min := list[0]
	for _, number := range list {
		if number > min {
			min = number
		}
	}
	return min
}
