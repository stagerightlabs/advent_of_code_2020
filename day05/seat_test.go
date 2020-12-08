package main

import "testing"

func TestNewSeat(t *testing.T) {
	testCases := []struct {
		code   string
		row    int
		column int
		id     int
	}{
		{code: "BFFFBBFRRR", row: 70, column: 7, id: 567},
		{code: "FFFBBBFRRR", row: 14, column: 7, id: 119},
		{code: "BBFFBBFRLL", row: 102, column: 4, id: 820},
	}

	for _, tc := range testCases {
		seat := NewSeat(tc.code)

		if seat.Row != tc.row {
			t.Errorf("Expected seat %q to have a row of %v, got %v", tc.code, tc.row, seat.Row)
		}

		if seat.Column != tc.column {
			t.Errorf("Expected seat %q to have a column of %v, got %v", tc.code, tc.row, seat.Column)
		}

		if seat.ID != tc.id {
			t.Errorf("Expected seat %q to have an ID of %v, got %v", tc.code, tc.id, seat.ID)
		}
	}
}
