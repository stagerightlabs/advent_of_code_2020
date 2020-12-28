package main

import (
	"testing"
)

func TestNavigate(t *testing.T) {
	instructions := getTestInput()
	ship := NewShip()

	Navigate(&ship, instructions)

	if ship.X != 17 {
		t.Errorf("Expected the ship to be at X coordinate %v, got %v", 17, ship.X)
	}

	if ship.Y != -8 {
		t.Errorf("Expected the ship to be at Y coordinate %v, got %v", -8, ship.Y)
	}

	if ship.CoordinateSum() != 25 {
		t.Errorf("Expected coordinate sum to be %v, got %v", 25, ship.CoordinateSum())
	}
}

func TestNavigateAlt(t *testing.T) {
	instructions := getTestInput()
	ship := NewShip()

	NavigateAlt(&ship, instructions)

	if ship.X != 214 {
		t.Errorf("Expected the ship to be at X coordinate %v, got %v", 214, ship.X)
	}

	if ship.Y != -72 {
		t.Errorf("Expected the ship to be at Y coordinate %v, got %v", -72, ship.Y)
	}

	if ship.CoordinateSum() != 286 {
		t.Errorf("Expected coordinate sum to be %v, got %v", 286, ship.CoordinateSum())
	}
}

func TestRotateClockwise(t *testing.T) {
	ship := NewShip()

	ship.RotateWaypointClockwise()
	if ship.waypointX != 1 || ship.waypointY != -10 {
		t.Errorf("Expected rotated coordinates to be (%v, %v), got (%v, %v)", 1, -10, ship.waypointX, ship.waypointY)
	}

	ship.RotateWaypointClockwise()
	if ship.waypointX != -10 || ship.waypointY != -1 {
		t.Errorf("Expected rotated coordinates to be (%v, %v), got (%v, %v)", -10, -1, ship.waypointX, ship.waypointY)
	}

	ship.RotateWaypointClockwise()
	if ship.waypointX != -1 || ship.waypointY != 10 {
		t.Errorf("Expected rotated coordinates to be (%v, %v), got (%v, %v)", -10, -1, ship.waypointX, ship.waypointY)
	}

	ship.RotateWaypointClockwise()
	if ship.waypointX != 10 || ship.waypointY != 1 {
		t.Errorf("Expected rotated coordinates to be (%v, %v), got (%v, %v)", 10, 1, ship.waypointX, ship.waypointY)
	}
}

func TestRotateCounterClockwise(t *testing.T) {
	ship := NewShip()

	ship.RotateWaypointCounterClockwise()
	if ship.waypointX != -1 || ship.waypointY != 10 {
		t.Errorf("Expected rotated coordinates to be (%v, %v), got (%v, %v)", -10, 1, ship.waypointX, ship.waypointY)
	}

	ship.RotateWaypointCounterClockwise()
	if ship.waypointX != -10 || ship.waypointY != -1 {
		t.Errorf("Expected rotated coordinates to be (%v, %v), got (%v, %v)", -10, -1, ship.waypointX, ship.waypointY)
	}

	ship.RotateWaypointCounterClockwise()
	if ship.waypointX != 1 || ship.waypointY != -10 {
		t.Errorf("Expected rotated coordinates to be (%v, %v), got (%v, %v)", 10, -1, ship.waypointX, ship.waypointY)
	}

	ship.RotateWaypointCounterClockwise()
	if ship.waypointX != 10 || ship.waypointY != 1 {
		t.Errorf("Expected rotated coordinates to be (%v, %v), got (%v, %v)", 10, 1, ship.waypointX, ship.waypointY)
	}
}

func getTestInput() string {
	return `
F10
N3
F7
R90
F11
`
}
