package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/brandonc/advent2020/internal/handheld"
	"github.com/brandonc/advent2020/pkg/tools"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "8 [input file]",
		Short: "Runs the day 8 challenge",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunWithArgs(args, day8)
		},
	})
}

func day8(file *os.File) error {
	reader, err := tools.Readlines(file)
	if err != nil {
		return fmt.Errorf("could not read input file: %w", err)
	}

	lines := make([]string, 0, 16)
	for line := range reader {
		lines = append(lines, line)
	}

	h, err := handheld.NewHandheld(lines)

	if err != nil {
		return err
	}

	v, loop := h.RunUntilLoopDetected()

	if (loop) {
		fmt.Printf("the value of the accumulator after loop is %d (part one)\n", v)
	} else {
		return fmt.Errorf("No loop was detected in this instruction")
	}

	for change := 0; change < len(lines); change++ {
		modified := make([]string, len(lines))
		copy(modified, lines)
		
		split := strings.Fields(modified[change])
		if split[0] == "jmp" {
			modified[change] = fmt.Sprintf("%s %s", "nop", split[1])
		} else if split[0] == "nop" {
			modified[change] = fmt.Sprintf("%s %s", "jmp", split[1])
		} else {
			continue
		}

		variant, err := handheld.NewHandheld(modified)

		if err != nil {
			return err
		}

		variantAcc, variantLoop := variant.RunUntilLoopDetected()

		if !variantLoop {
			fmt.Printf("the value of the accumulator at the end is %d (part two)\n", variantAcc)
			return nil
		}
	}

	return fmt.Errorf("no way to change the program to not end in a loop")
}