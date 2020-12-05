package main

import (
	"errors"
	"strings"
)

// Grid represents our map of trees and open spaces
type Grid struct {
	template       [][]byte
	templateWidth  int
	templateHeight int
	tobogganX      int
	tobogganY      int
}

// AdvanceToboggan moves the toboggan coordinates across three positions
// and down one position.
func (g *Grid) AdvanceToboggan(right int, down int) error {

	g.tobogganY = g.tobogganY + down

	if g.tobogganY >= g.templateHeight {
		return errors.New("out of bounds")
	}

	g.tobogganX = g.tobogganX + right

	if g.tobogganX > g.templateWidth-1 {
		g.tobogganX = g.tobogganX - g.templateWidth
	}

	return nil
}

// TobogganHasMetATree tells us if the toboggan's current
// position is on top of a tree.
func (g *Grid) TobogganHasMetATree() bool {
	// fmt.Printf("[%v,%v]: %c\n", g.tobogganY, g.tobogganX, g.template[g.tobogganY][g.tobogganX])

	if g.tobogganY > g.templateHeight {
		return false
	}

	return g.template[g.tobogganY][g.tobogganX] == '#'
}

// NewGrid creates a new grid struct from an input string
func NewGrid(input string) Grid {
	g := Grid{}

	for _, line := range strings.Split(input, "\n") {

		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		g.template = append(g.template, []byte(line))
		g.templateWidth = len(line)
	}

	g.templateHeight = len(g.template)
	g.tobogganX = 0
	g.tobogganY = 0

	return g
}
