package day03

import (
	"fmt"
	"log"
	"os"

	"github.com/brandonc/advent2020/pkg/tools"
)

type vector struct {
	y int
	x int
}

func treesAtAngle(treeMap *[]string, angle vector) int {
	treeCount := 0
	width := len((*treeMap)[0])

	for y, x := 0, 0; y < len((*treeMap)); y, x = y + angle.y, x + angle.x {
		if (*treeMap)[y][x % width] == '#' {
			treeCount++
		}
	}

	return treeCount
}

// Run runs the day 3 example using the specified input
func Run(file *os.File) {
	scanner, err := tools.Readlines(file)

	if err != nil {
		log.Fatal(err)
		return
	}

	// Start of part 1

	treeMap := make([]string, 0, 100)
	for line := range scanner {
		treeMap = append(treeMap, line)
	}

	treeCountPartOne := treesAtAngle(&treeMap, vector{1, 3})

	fmt.Printf("Encountered %d trees (first part)\n", treeCountPartOne)

	// Start of part 2

	angles := []vector{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}}
	multipliedTrees := 1

	for _, v := range angles {
		i := treesAtAngle(&treeMap, v)
		fmt.Printf("Encountered %d trees at angle %d, %d\n", i, v.y, v.x)
		multipliedTrees *= i
	}

	fmt.Printf("Multiplied trees at various angles %d (second part)\n", multipliedTrees)
}
