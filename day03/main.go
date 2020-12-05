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

	grid := NewGrid(string(input))

	// Calculate the answer to part one
	treeCount3x1 := CountTrees(grid, 3, 1)

	fmt.Printf("Answer 1: Encountered %v trees\n", treeCount3x1)

	// Calculate the answer to part two
	var treeCounts []int
	traversals := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	for _, t := range traversals {
		count := CountTrees(grid, t[0], t[1])
		treeCounts = append(treeCounts, count)
	}

	treeCountProduct := treeCounts[0] * treeCounts[1]
	for i := 2; i < len(treeCounts); i++ {
		treeCountProduct = treeCountProduct * treeCounts[i]
	}

	fmt.Printf("Answer 2: Product of all tree counts: %v\n", treeCountProduct)
}

// CountTrees counts the number of trees encountered
// by our toboggan when traversing the grid.
func CountTrees(grid Grid, right int, down int) int {

	// fmt.Printf("%+v\n", grid)

	count := 0

	for {

		err := grid.AdvanceToboggan(right, down)

		if err != nil {
			break
		}

		if grid.TobogganHasMetATree() {
			count++
		}

		// fmt.Println([]int{grid.tobogganY, len(grid.template)})
		// if grid.tobogganY >= grid.templateHeight {
		// 	break
		// }
	}
	return count
}
