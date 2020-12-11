package day10

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/brandonc/advent2020/pkg/tools"
)

func validVariant(joltages []int) bool {
	current := 0
	if len(joltages) < 2 {
		return false
	}
	for _, j := range joltages {
		if current + 1 != j && current + 2 != j && current + 3 != j {
			return false
		}
		current = j
	}
	return true
}

func countVariants(start int, joltages []int) int {
	result := 1
	for i := start; i < len(joltages) - 1; i++ {
		if i + 2 < len(joltages) && joltages[i + 2] <= joltages[i] + 3 {
			result += countVariants(i + 2, joltages)
		}
		if i + 3 < len(joltages) && joltages[i + 3] <= joltages[i] + 3 {
			result += countVariants(i + 3, joltages)
		}
	}
	return result
}

// Run runs the day 10 challenge on the specified input
func Run(file *os.File) {
	joltages, err := tools.ReadlinesInts(file)

	if err != nil {
		log.Fatal(err)
		return
	}

	steps, step3s := 0, 1

	sort.Ints(joltages)

	current := 0
	for _, j := range joltages {
		if j == current + 1 {
			steps++
		} else if j == current + 3 {
			step3s++
		} else {
			log.Fatalf("next step %d is not 1 or 3 bigger than %d", j, current)
			return
		}

		current = j
	}

	fmt.Printf("%d (1-jolts) * %d (3-jolts) = %d (part one)\n", steps, step3s, steps * step3s)

	
	joltages = append([]int{0}, joltages...)
	joltages = append(joltages, joltages[len(joltages) - 1] + 3)

	fmt.Printf("%d variations are supported (part two)\n", countVariants(0, joltages))
}