package cmd

import "testing"

func TestOccupancy(t *testing.T) {
	var noDir = [][]rune{
		{'.', '#', '#', '.', '#', '#', '.'},
		{'#', '.', '#', '.', '#', '.', '#'},
		{'#', '#', '.', '.', '.', '#', '#'},
		{'.', '.', '.', 'L', '.', '.', '.'},
		{'#', '#', '.', '.', '.', '#', '#'},
		{'#', '.', '#', '.', '#', '.', '#'},
		{'.', '#', '#', '.', '#', '#', '.'},
	}

	n := occupiedAdjacentAny(noDir, 3, 3)
	if n != 0 {
		t.Errorf("n = %d; want 0", n)
	}

	var blocked = [][]rune{
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', 'L', '.', 'L', '.', '#', '.', '#', '.', '#', '.', '#', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}

	n = occupiedAdjacentAny(blocked, 1, 1)
	if n != 0 {
		t.Errorf("n = %d; want 0", n)
	}

	var everyDir = [][]rune{
		{'.', '.', '.', '.', '.', '.', '.', '#', '.'},
		{'.', '.', '.', '#', '.', '.', '.', '.', '.'},
		{'.', '#', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '#', 'L', '.', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '#', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '#', '.', '.', '.', '.', '.'},
	}

	n = occupiedAdjacentAny(everyDir, 4, 3)
	if n != 8 {
		t.Errorf("n = %d; want 8", n)
	}
}