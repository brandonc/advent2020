package day05

import (
	"fmt"
	"log"
	"os"

	"github.com/brandonc/advent2020/tools"
)

// Run runs the day 5 advent of code challenge using the specified input
func Run(file *os.File) {
	scanner, err := tools.Readlines(file);

	if err != nil {
		log.Fatal(err);
		return
	}

	highestID := 0
	seatMap := make(map[int]bool)

	for line := range scanner {
		row := 0
		col := 0
		// Traverses a pair of hypothetical complete binary tree using
		// control characters, F for left and B for right
		for index, c := range line {
			switch index {
			case 0, 1, 2, 3, 4, 5, 6:
				if c == 'F' {
					// Lower half
					row = row * 2
				}
				if c == 'B' {
					// Upper half
					row = row * 2 + 1
				}
			case 7, 8, 9:
				if c == 'L' {
					// Lower half
					col = col * 2
				}
				if c == 'R' {
					// Upper half
					col = col * 2 + 1
				}
			}
		}

		id := (row * 8) + col
		seatMap[id] = true
		if (id > highestID) {
			highestID = id
		}
	}

	fmt.Printf("Seat with the highest ID is %d (part one)\n", highestID)

	for row := 0; row < 128; row++ {
		for col := 0; col < 8; col++ {
			checkID := (row * 8) + col
			_, ok := seatMap[checkID]
			_, okNext := seatMap[checkID + 1]
			_, okPrev := seatMap[checkID - 1]
			if !ok && okNext && okPrev {
				fmt.Printf("Your seat ID is %d (part two)\n", checkID)
				return
			}
		}
	}
}