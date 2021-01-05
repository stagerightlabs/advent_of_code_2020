package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	earliest, busses := ParseInputForAnswerOne(getTestInput())
	expectedBusses := []int{7, 13, 59, 31, 19}

	if earliest != 939 {
		t.Errorf("Expected to get a departure time of %v, got %v", 939, earliest)
	}

	if !reflect.DeepEqual(busses, expectedBusses) {
		t.Errorf("Expected to see bus Ids %v, got %v", expectedBusses, busses)
	}
}

func TestFindDepartureOption(t *testing.T) {
	earliest, busses := ParseInputForAnswerOne(getTestInput())

	busID, departure := FindDepartureOption(earliest, busses)

	if busID != 59 {
		t.Errorf("Expected to see bus ID %v, got %v", 59, busID)
	}

	if departure != 944 {
		t.Errorf("Expected departure time to be %v, got %v", 944, departure)
	}
}

func TestParseInputForAnswerTwo(t *testing.T) {
	busses := ParseInputForAnswerTwo(getTestInput())
	expectedBusses := map[int]int{
		0: 7,
		1: 13,
		4: 59,
		6: 31,
		7: 19,
	}

	if !reflect.DeepEqual(busses, expectedBusses) {
		t.Errorf("Expected to see bus Ids %v, got %v", expectedBusses, busses)
	}
}

func TestFindDepartureSequenceTimestamp(t *testing.T) {
	busses := ParseInputForAnswerTwo(getTestInput())

	timestamp := FindDepartureSequenceTimestamp(busses)
	want := uint64(1068781)

	if timestamp != want {
		t.Errorf("Expected departure timestamp to be %v, got %v", want, timestamp)
	}

	// Test Case 2
	busses = map[int]int{
		0: 17,
		2: 13,
		3: 19,
	}

	timestamp = FindDepartureSequenceTimestamp(busses)
	want = uint64(3417)

	if timestamp != want {
		t.Errorf("Expected departure timestamp to be %v, got %v", want, timestamp)
	}

	// Test Case 3
	busses = map[int]int{
		0: 67,
		1: 7,
		2: 59,
		3: 61,
	}

	timestamp = FindDepartureSequenceTimestamp(busses)
	want = uint64(754018)

	if timestamp != want {
		t.Errorf("Expected departure timestamp to be %v, got %v", want, timestamp)
	}

	// Test Case 4
	busses = map[int]int{
		0: 67,
		2: 7,
		3: 59,
		4: 61,
	}

	timestamp = FindDepartureSequenceTimestamp(busses)
	want = uint64(779210)

	if timestamp != want {
		t.Errorf("Expected departure timestamp to be %v, got %v", want, timestamp)
	}

	// Test Case 5
	busses = map[int]int{
		0: 67,
		1: 7,
		3: 59,
		4: 61,
	}

	timestamp = FindDepartureSequenceTimestamp(busses)
	want = uint64(1261476)

	if timestamp != want {
		t.Errorf("Expected departure timestamp to be %v, got %v", want, timestamp)
	}

	// Test Case 5
	busses = map[int]int{
		0: 1789,
		1: 37,
		2: 47,
		3: 1889,
	}

	timestamp = FindDepartureSequenceTimestamp(busses)
	want = uint64(1202161486)

	if timestamp != want {
		t.Errorf("Expected departure timestamp to be %v, got %v", want, timestamp)
	}
}

func getTestInput() string {
	return `
939
7,13,x,x,59,x,31,19
`
}
