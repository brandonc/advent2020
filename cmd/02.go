package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/brandonc/advent2020/pkg/tools"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "2 [input file]",
		Short: "Runs the day 2 challenge",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunWithArgs(args, day2)
		},
	})
}

func day2(file *os.File) error {
	regex := *regexp.MustCompile(`^(\d+)-(\d+) ([a-z]): ([a-z]+)$`)

	valid1st := 0
	valid2nd := 0

	scanner, err := tools.Readlines(file)
	if err != nil {
		return fmt.Errorf("could not read input file: %w", err)
	}
	
	for line := range scanner {
		matches := regex.FindStringSubmatch(line)
		
		min, err := strconv.Atoi(matches[1])

		if err != nil {
			return fmt.Errorf("could not parse min int: %w", err)
		}
		
		max, err := strconv.Atoi(matches[2])
		
		if err != nil {
			return fmt.Errorf("could not parse max int: %w", err)
		}

		exhibit := matches[3][0]
		passwd := matches[4]

		found := 0

		for _, c := range passwd {
			if c == rune(exhibit) {
				found++
			}
		}

		if found >= min && found <= max {
			valid1st++
		}

		if (passwd[min - 1] == exhibit || passwd[max - 1] == exhibit) && !(passwd[min - 1] == exhibit && passwd[max - 1] == exhibit) {
			valid2nd++
		}
	}

	fmt.Printf("There are %d valid passwords (part one)\n", valid1st)
	fmt.Printf("There are %d valid passwords (part two)\n", valid2nd)

	return nil
}