package cmd

import (
	"fmt"
	"os"

	"github.com/brandonc/advent2020/pkg/tools"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "6 [input file]",
		Short: "Runs the day 6 challenge",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunWithArgs(args, day6)
		},
	})
}

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

func day6(file *os.File) error {
	scanner, err := tools.Readlines(file)

	if err != nil {
		return fmt.Errorf("could not read input file: %w", err)
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


	return nil
}
