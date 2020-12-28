package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err.Error())
	}

	// Part One
	first := NewShip()
	Navigate(&first, string(input))
	fmt.Printf("Answer 1: Manhattan distance is %v\n", first.CoordinateSum())

	// Part Two
	second := NewShip()
	NavigateAlt(&second, string(input))
	fmt.Printf("Answer 2: Manhattan distance is %v\n", second.CoordinateSum())

}

// Ship is a struct that represents the state of our ship
type Ship struct {
	X         int
	Y         int
	waypointX int
	waypointY int
	degrees   int
}

// CoordinateSum returns the total of the absolute values of a ship's coordinates
func (s *Ship) CoordinateSum() int {
	return int(math.Abs(float64(s.X))) + int(math.Abs(float64(s.Y)))
}

// RotateWaypointClockwise moves the waypoint
// around the ship in a clockwise direction
func (s *Ship) RotateWaypointClockwise() {
	if s.waypointX == 0 && s.waypointY > 0 {
		s.waypointX = s.waypointY
		s.waypointY = 0
	} else if s.waypointX == 0 && s.waypointY < 0 {
		s.waypointX = s.waypointY * -1
		s.waypointY = 0
	} else if s.waypointY == 0 && s.waypointX > 0 {
		s.waypointY = s.waypointX
		s.waypointX = 0
	} else if s.waypointY == 0 && s.waypointX < 0 {
		s.waypointY = s.waypointX * -1
		s.waypointX = 0
	} else if s.waypointX > 0 && s.waypointY > 0 {
		y := s.waypointY
		s.waypointY = s.waypointX * -1
		s.waypointX = y
	} else if s.waypointX > 0 && s.waypointY < 0 {
		x := s.waypointX
		s.waypointX = s.waypointY
		s.waypointY = x * -1
	} else if s.waypointX < 0 && s.waypointY < 0 {
		x := s.waypointX
		s.waypointX = s.waypointY
		s.waypointY = x * -1
	} else if s.waypointX < 0 && s.waypointY > 0 {
		y := s.waypointY
		s.waypointY = s.waypointX * -1
		s.waypointX = y
	}
}

// RotateWaypointCounterClockwise moves the waypoint
// around the ship in a counter clockwise direction
func (s *Ship) RotateWaypointCounterClockwise() {
	if s.waypointX == 0 && s.waypointY > 0 {
		s.waypointX = s.waypointY * -1
		s.waypointY = 0
	} else if s.waypointX == 0 && s.waypointY < 0 {
		s.waypointX = s.waypointY
		s.waypointY = 0
	} else if s.waypointY == 0 && s.waypointX > 0 {
		s.waypointY = s.waypointX * -1
		s.waypointX = 0
	} else if s.waypointY == 0 && s.waypointX < 0 {
		s.waypointY = s.waypointX
		s.waypointX = 0
	} else if s.waypointX > 0 && s.waypointY > 0 {
		y := s.waypointY
		s.waypointY = s.waypointX
		s.waypointX = y * -1
	} else if s.waypointX > 0 && s.waypointY < 0 {
		y := s.waypointY
		s.waypointY = s.waypointX
		s.waypointX = y * -1
	} else if s.waypointX < 0 && s.waypointY < 0 {
		y := s.waypointY
		s.waypointY = s.waypointX
		s.waypointX = y * -1
	} else if s.waypointX < 0 && s.waypointY > 0 {
		y := s.waypointY
		s.waypointY = s.waypointX
		s.waypointX = y * -1
	}
}

// NewShip creates a new ship struct
func NewShip() Ship {
	ship := Ship{}
	ship.X = 0
	ship.Y = 0
	ship.waypointX = 10
	ship.waypointY = 1
	ship.degrees = 90

	return ship
}

// Navigate a ship by following all of the provided instructions
func Navigate(ship *Ship, instructions string) {
	for _, instruction := range strings.Split(instructions, "\n") {
		DoTick(ship, instruction)
	}
}

// DoTick applies a single instruction to our ship
func DoTick(ship *Ship, instruction string) {
	if len(instruction) == 0 {
		return
	}

	action := rune(instruction[0])
	value, err := strconv.Atoi(instruction[1:])
	if err != nil {
		panic("Invalid instruction: " + instruction)
	}

	directions := map[int]rune{
		0:   'N',
		90:  'E',
		180: 'S',
		270: 'W',
	}

	if action == 'F' {
		action = directions[ship.degrees]
	}

	switch action {
	case 'N':
		ship.Y += value
	case 'S':
		ship.Y -= value
	case 'E':
		ship.X += value
	case 'W':
		ship.X -= value
	case 'L':
		ship.degrees -= value
	case 'R':
		ship.degrees += value
	}

	if ship.degrees >= 360 {
		ship.degrees = ship.degrees - 360
	}
	if ship.degrees < 0 {
		ship.degrees += 360
	}
}

// NavigateAlt a ship by following all of the provided instructions
// using the "correct" interpretation
func NavigateAlt(ship *Ship, instructions string) {
	for _, instruction := range strings.Split(instructions, "\n") {
		DoAlternateTick(ship, instruction)
	}
}

// DoAlternateTick applies a single "correct" instruction to our ship
func DoAlternateTick(ship *Ship, instruction string) {
	if len(instruction) == 0 {
		return
	}

	action := rune(instruction[0])
	value, err := strconv.Atoi(instruction[1:])
	if err != nil {
		panic("Invalid instruction: " + instruction)
	}

	switch action {
	case 'N':
		ship.waypointY += value
	case 'S':
		ship.waypointY -= value
	case 'E':
		ship.waypointX += value
	case 'W':
		ship.waypointX -= value
	case 'L':
		// https://github.com/alexchao26/advent-of-code-go/blob/main/2020/day12/main.go#L98
		turns := value / 90
		for turns > 0 {
			ship.waypointX, ship.waypointY = -ship.waypointY, ship.waypointX
			turns--
		}
	case 'R':
		turns := value / 90
		for turns > 0 {
			ship.waypointX, ship.waypointY = ship.waypointY, -ship.waypointX
			turns--
		}
	case 'F':
		ship.X += (ship.waypointX * value)
		ship.Y += (ship.waypointY * value)
	}
}
