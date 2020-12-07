package day06

import (
	"fmt"
	"log"
	"os"

	"github.com/brandonc/advent2020/pkg/tools"
)

func affirmatives(lines []string) []int {
	answers := make([]int, 27)
	for i := 0; i < 27; i++ {
		answers[i] = 0
	}

	for _, line := range lines {
		for _, c := range line {
			answers[c - 'a']++
		}
	}

	return answers
}

func countAnyAffirmatives(answers *[]int) int {
	total := 0
	for _, ans := range *answers {
		if ans > 0 {
			total++
		}
	}
	return total
}

func countAllAffirmatives(answers *[]int, groupCount int) int {
	total := 0
	for _, ans := range *answers {
		if ans == groupCount {
			total++
		}
	}
	return total
}

// Run runs the day 6 challenge on the specified input
func Run(file *os.File) {
	scanner, err := tools.Readlines(file)

	if err != nil {
		log.Fatal(err)
		return
	}

	group := []string{}
	totalAnyAffirmatives := 0
	totalAllAffirmatives := 0
	for line := range scanner {
		if len(line) == 0 {
			aff := affirmatives(group)
			totalAnyAffirmatives += countAnyAffirmatives(&aff)
			totalAllAffirmatives += countAllAffirmatives(&aff, len(group))
			group = []string{}
			continue
		}
		group = append(group, line)
	}

	aff := affirmatives(group)
	totalAnyAffirmatives += countAnyAffirmatives(&aff)
	totalAllAffirmatives += countAllAffirmatives(&aff, len(group))

	fmt.Printf("Sum of any counts %d (part one)\n", totalAnyAffirmatives)
	fmt.Printf("Sum of every counts %d (part two)\n", totalAllAffirmatives)
}
