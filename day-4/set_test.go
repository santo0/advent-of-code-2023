package main

import (
	"testing"
)

func TestSet(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(1)
	if set.Len() != 1 {
		t.Fatal("No idempotence")
	}
	if !set.In(1) {
		t.Fatal("No correct storing")
	}
	set.Remove(1)
	if set.Len() != 0 {
		t.Fatal("No removing correctly")
	}
	set.Add(1)
	set.Add(2)
	set.Add(1)
	if set.Len() != 2 {
		t.Fatal("Multiple elements not supported")
	}
}
func TestSetFromSLice(t *testing.T) {
	set := NewSetFromSlice[int]([]int{1, 2, 3, 3, 2, 1})
	if set.Len() != 3 {
		t.Fatal("No idempotence")
	}
	if !set.In(1) {
		t.Fatal("No correct storing")
	}
	set.Remove(1)
	if set.Len() != 2 {
		t.Fatal("No removing correctly")
	}
	set.Add(1)
	set.Add(2)
	set.Add(1)
	if set.Len() != 3 {
		t.Fatal("Multiple elements not supported")
	}
}
