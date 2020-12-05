package main

import "testing"

func TestNewGrid(t *testing.T) {
	grid := getTestCaseGrid()

	if grid.tobogganX != 0 || grid.tobogganY != 0 {
		t.Errorf("grid coordinates were not initialized correctly")
	}

	if grid.templateHeight != 11 {
		t.Errorf("expected grid template height of %v, got %v", 11, grid.templateHeight)
	}

	if grid.templateWidth != 11 {
		t.Errorf("expected grid template width of %v, got %v", 11, grid.templateWidth)
	}

	if len(grid.template) != 11 {
		t.Errorf("grid template has not been initialized correctly")
	}
}

func TestAdvanceToboggan(t *testing.T) {
	grid := getTestCaseGrid()

	for i := 1; i < grid.templateHeight; i++ {
		right := 3
		down := 1

		y := grid.tobogganY

		grid.AdvanceToboggan(right, down)

		if grid.tobogganY != y+down {
			t.Errorf("At move %v, expected tobboganY to be %v, got %v", i, i, grid.tobogganY)
		}

		if grid.tobogganX > grid.templateWidth {
			t.Errorf("At move %v, the tobogganX position exceeded the template width of %v", i, grid.templateWidth)
		}
	}
}

func TestTobogganHasMetATree(t *testing.T) {
	grid := getTestCaseGrid()

	grid.tobogganX = 2
	grid.tobogganY = 2

	if grid.TobogganHasMetATree() {
		t.Errorf("reading a tree at position [%v, %v] when it is not actually there.", 2, 2)
	}

	grid.tobogganX = 2
	grid.tobogganY = 3

	if !grid.TobogganHasMetATree() {
		t.Errorf("reading an empty space at position [%v, %v] when it should be a tree.", 2, 2)
	}
}

func getTestCaseGrid() Grid {
	return NewGrid(`
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`)
}
