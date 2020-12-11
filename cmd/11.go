package cmd

import (
	"fmt"
	"os"

	"github.com/brandonc/advent2020/pkg/tools"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "11 [input file]",
		Short: "Runs the day 11 challenge",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunWithArgs(args, day11)
		},
	})
}

func occupiedAdjacent(plan [][]rune, y int, x int) int {
	result := 0
	if y > 0 && x > 0 && plan[y-1][x-1] == '#' {
		result++
	}
	if y > 0 && plan[y-1][x] == '#' {
		result++
	}
	if y > 0 && x < len(plan[y-1]) - 1 && plan[y-1][x+1] == '#' {
		result++
	}
	if x > 0 && plan[y][x-1] == '#' {
		result++
	}
	if x < len(plan[y]) - 1 && plan[y][x+1] == '#' {
		result++
	}
	if y < len(plan) - 1 && x > 0 && plan[y+1][x-1] == '#' {
		result++
	}
	if y < len(plan) - 1 && plan[y+1][x] == '#' {
		result++
	}
	if y < len(plan) - 1 && x < len(plan[y+1]) - 1 && plan[y+1][x+1] == '#' {
		result++
	}

	return result
}

func occupiedAdjacentAny(plan [][]rune, y int, x int) int {
	count := 0

	nwy, nwx := y, x
	for {
		nwy--
		nwx--
		
		if nwy < 0 || nwx < 0 {
			break
		}

		if plan[nwy][nwx] != '.' {
			if plan[nwy][nwx] == '#' {
				count++
			}
			break
		}
	}

	n := y
	for {
		n--
		
		if n < 0 {
			break
		}

		if plan[n][x] != '.' {
			if plan[n][x] == '#' {
				count++
			}
			break
		}
	}

	ney, nex := y, x
	for {
		ney--
		nex++
		
		if ney < 0 || nex >= len(plan[0]) {
			break
		}

		if plan[ney][nex] != '.' {
			if plan[ney][nex] == '#' {
				count++
			}
			break
		}
	}

	e := x
	for {
		e++
		
		if e >= len(plan[0]) {
			break
		}

		if plan[y][e] != '.' {
			if plan[y][e] == '#' {
				count++
			}
			break
		}
	}

	sey, sex := y, x
	for {
		sey++
		sex++
		
		if sey >= len(plan) || sex >= len(plan[0]) {
			break
		}

		if plan[sey][sex] != '.' {
			if plan[sey][sex] == '#' {
				count++
			}
			break
		}
	}

	s := y
	for {
		s++
		
		if s >= len(plan) {
			break
		}

		if plan[s][x] != '.' {
			if plan[s][x] == '#' {
				count++
			}
			break
		}
	}

	swy, swx := y, x
	for {
		swy++
		swx--
		
		if swy >= len(plan) || swx < 0 {
			break
		}

		if plan[swy][swx] != '.' {
			if plan[swy][swx] == '#' {
				count++
			}
			break
		}
	}

	w := x
	for {
		w--
		
		if w < 0 {
			break
		}

		if plan[y][w] != '.' {
			if plan[y][w] == '#' {
				count++
			}
			break
		}
	}

	return count
}

func countOccupied(plan [][]rune) int {
	count := 0
	for y := 0; y < len(plan); y++ {
		for x := 0; x < len(plan[0]); x++ {
			if plan[y][x] == '#' {
				count++
			}
		}
	}
	return count
}

func copyPlan(plan [][]rune) [][]rune {
	result := make([][]rune, len(plan))

	for y := 0; y < len(plan); y++ {
		result[y] = make([]rune, len(plan[0]))
		for x := 0; x < len(plan[0]); x++ {
			result[y][x] = plan[y][x]
		}
	}
	return result
}

func makePlan(scanner <-chan string) [][]rune {
	result := make([][]rune, 0)

	for line := range scanner {
		rowLen := len(line)
		row := make([]rune, rowLen)
		for i, c := range line {
			row[i] = rune(c)
		}
		result = append(result, row)
	}

	return result
}

func day11(file *os.File) error {
	scanner, err := tools.Readlines(file)

	if err != nil {
		return fmt.Errorf("could not read input file: %w", err)
	}

	originalPlan := makePlan(scanner)

	if len(originalPlan) == 0 || len(originalPlan[0]) == 0 {
		return fmt.Errorf("no plan data was provided")
	}

	plan := copyPlan(originalPlan)
	for {
		nextPlan := copyPlan(plan)
		rulesApplied := false

		for y := 0; y < len(plan); y++ {
			for x := 0; x < len(plan[0]); x++ {
				if plan[y][x] == 'L' && occupiedAdjacent(plan, y, x) == 0 {
					nextPlan[y][x] = '#'
					rulesApplied = true
				} else if plan[y][x] == '#' && occupiedAdjacent(plan, y, x) >= 4 {
					nextPlan[y][x] = 'L'
					rulesApplied = true
				}
			}
		}

		if !rulesApplied {
			break
		}

		plan = nextPlan
	}

	fmt.Printf("There are %d occupied seats (part one)\n", countOccupied(plan))

	planPart2 := copyPlan(originalPlan)
	for {
		nextPlanPart2 := copyPlan(planPart2)
		rulesApplied := false

		for y := 0; y < len(planPart2); y++ {
			for x := 0; x < len(planPart2[0]); x++ {
				if planPart2[y][x] == 'L' && occupiedAdjacentAny(planPart2, y, x) == 0 {
					nextPlanPart2[y][x] = '#'
					rulesApplied = true
				} else if planPart2[y][x] == '#' && occupiedAdjacentAny(planPart2, y, x) >= 5 {
					nextPlanPart2[y][x] = 'L'
					rulesApplied = true
				}
			}
		}

		if !rulesApplied {
			break
		}

		planPart2 = nextPlanPart2
	}

	fmt.Printf("There are %d occupied seats (part two)\n", countOccupied(planPart2))

	return nil
}