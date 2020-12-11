package cmd

import (
	"fmt"
	"os"

	"github.com/brandonc/advent2020/pkg/tools"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "3 [input file]",
		Short: "Runs the day 3 challenge",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunWithArgs(args, day3)
		},
	})
}

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

func day3(file *os.File) error {
	scanner, err := tools.Readlines(file)

	if err != nil {
		return fmt.Errorf("could not read input file: %w", err)
	}

	// Start of part 1

	treeMap := make([]string, 0, 100)
	for line := range scanner {
		treeMap = append(treeMap, line)
	}

	treeCountPartOne := treesAtAngle(&treeMap, vector{1, 3})

	fmt.Printf("Encountered %d trees (part one)\n", treeCountPartOne)

	// Start of part 2

	angles := [...]vector{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}}
	multipliedTrees := 1

	for _, v := range angles {
		i := treesAtAngle(&treeMap, v)
		multipliedTrees *= i
	}

	fmt.Printf("Multiplied trees at various angles %d (part two)\n", multipliedTrees)

	return nil
}
