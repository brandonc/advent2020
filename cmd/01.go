package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/brandonc/advent2020/pkg/tools"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "1 [input file]",
		Short: "Runs the day 1 challenge",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunWithArgs(args, day1)
		},
	})
}

func day1(file *os.File) error {
	sorted, err := tools.ReadlinesInts(file)

	if err != nil {
		return err
	}

	sort.Ints(sorted)
	loops := 0

	for aIndex := 0; aIndex < len(sorted); aIndex++ {
		for zIndex := 1; zIndex < len(sorted); zIndex++ {
			for xIndex := 2; xIndex < len(sorted); xIndex++ {
				loops++
				if (aIndex == zIndex || aIndex == xIndex || zIndex == xIndex) {
					continue
				}

				sum := sorted[aIndex] + sorted[zIndex] + sorted[xIndex]

				if sum == 2020 {
					fmt.Printf("%d + %d + %d = 2020 (part one)\n", sorted[aIndex], sorted[zIndex], sorted[xIndex])
					fmt.Printf("%d * %d * %d = %d (part two)\n", sorted[aIndex], sorted[zIndex], sorted[xIndex], sorted[aIndex] * sorted[zIndex] * sorted[xIndex])
					return nil
				}

				if sum > 2020 {
					break
				}
			}

			if sorted[aIndex] + sorted[zIndex] > 2020 {
				break
			}
		}
		if sorted[aIndex] > 2020 {
			break
		}
	}

	return fmt.Errorf("No 2020 sum found")
}
