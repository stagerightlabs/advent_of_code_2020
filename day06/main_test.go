package main

import (
	"testing"
)

func TestGroupExtraction(t *testing.T) {
	groups := ExtractGroupsFromInput(getTestInput())

	if len(groups) != 5 {
		t.Errorf("Expected 5 groups in the input string, got %v", len(groups))
	}
}

func TestNewGroup(t *testing.T) {

	group := NewGroup("abc")
	got := group.totalAnsweredByAnyone
	want := 3
	if got != want {
		t.Errorf("Expected %q to  have an 'anyone' answer count of %v, got %v", "abc", want, got)
	}
	got = group.totalAnsweredByEveryone
	want = 3
	if got != want {
		t.Errorf("Expected %q to  have an 'everyone' answer count of %v, got %v", "abc", want, got)
	}
	got = group.members
	want = 1
	if got != want {
		t.Errorf("Expected %q to have a member count of %v, got %v", "abc", want, got)
	}

	group = NewGroup("a\nb\nc\n")
	got = group.totalAnsweredByAnyone
	want = 3
	if got != want {
		t.Errorf("Expected %q to  have an 'anyone' answer count of %v, got %v", "abc", want, got)
	}
	got = group.totalAnsweredByEveryone
	want = 0
	if got != want {
		t.Errorf("Expected %q to  have an 'everyone' answer count of %v, got %v", "abc", want, got)
	}
	got = group.members
	want = 3
	if got != want {
		t.Errorf("Expected %q to have a member count of %v, got %v", "abc", want, got)
	}

	group = NewGroup("ab\nac\n")
	got = group.totalAnsweredByAnyone
	want = 3
	if got != want {
		t.Errorf("Expected %q to  have an 'anyone' answer count of %v, got %v", "abc", want, got)
	}
	got = group.totalAnsweredByEveryone
	want = 1
	if got != want {
		t.Errorf("Expected %q to  have an 'everyone' answer count of %v, got %v", "abc", want, got)
	}
	got = group.members
	want = 2
	if got != want {
		t.Errorf("Expected %q to have a member count of %v, got %v", "abc", want, got)
	}

	group = NewGroup("a\na\na\na\n")
	got = group.totalAnsweredByAnyone
	want = 1
	if got != want {
		t.Errorf("Expected %q to  have an 'anyone' answer count of %v, got %v", "abc", want, got)
	}
	got = group.totalAnsweredByEveryone
	want = 1
	if got != want {
		t.Errorf("Expected %q to  have an 'everyone' answer count of %v, got %v", "abc", want, got)
	}
	got = group.members
	want = 4
	if got != want {
		t.Errorf("Expected %q to have a member count of %v, got %v", "abc", want, got)
	}

	group = NewGroup("b")
	got = group.totalAnsweredByAnyone
	want = 1
	if got != want {
		t.Errorf("Expected %q to  have an 'anyone' answer count of %v, got %v", "abc", want, got)
	}
	got = group.totalAnsweredByEveryone
	want = 1
	if got != want {
		t.Errorf("Expected %q to  have an 'everyone' answer count of %v, got %v", "abc", want, got)
	}
	got = group.members
	want = 1
	if got != want {
		t.Errorf("Expected %q to have a member count of %v, got %v", "abc", want, got)
	}
}

func getTestInput() string {
	return `
abc

a
b
c

ab
ac

a
a
a
a

b
`
}
