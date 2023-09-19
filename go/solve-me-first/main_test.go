package main

import "testing"

func TestSolveMeFirst(t *testing.T) {
	result := solveMeFirst(7, 3)
	if result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}
}
