package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	scanner := bufio.NewScanner(f)

	var input []int32

	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			panic(err.Error())
		}
		input = append(input, int32(number))
	}

	entries, err := LocateTwoEntries(input, 2020)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Answer 1: %v times %v is %v\n", entries[0], entries[1], entries[0]*entries[1])

	entries, err = LocateThreeEntries(input, 2020)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Answer 2: %v times %v times %v is %v\n", entries[0], entries[1], entries[2], entries[0]*entries[1]*entries[2])
}

// LocateTwoEntries finds two numbers in a list that
// equal a target value when added together.
func LocateTwoEntries(list []int32, target int32) ([]int32, error) {

	for firstIndex, firstValue := range list {
		for secondIndex, secondValue := range list {
			if firstIndex == secondIndex {
				break
			}

			if firstValue+secondValue == target {
				return []int32{firstValue, secondValue}, nil
			}
		}
	}

	return []int32{}, errors.New("no input values were found that sum to " + strconv.Itoa(int(target)))
}

// LocateThreeEntries finds three numbers in a list that
// equal a target value when added together.
func LocateThreeEntries(list []int32, target int32) ([]int32, error) {

	for firstIndex, firstValue := range list {
		for secondIndex, secondValue := range list {
			if firstIndex == secondIndex {
				break
			}

			for thirdIndex, thirdValue := range list {
				if thirdIndex == secondIndex {
					break
				}

				if firstValue+secondValue+thirdValue == target {
					return []int32{firstValue, secondValue, thirdValue}, nil
				}
			}
		}
	}

	return []int32{}, errors.New("no input values were found that sum to 2020")
}
