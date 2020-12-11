package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "advent",
	Short: "runs a named advent of code challenge",
	Long: "an implementation in go of https://adventofcode.com/2020",
}

// Execute executes the specified CLI command
func Execute() error {
	return rootCmd.Execute()
}

// DayCommand describes a function that takes a file and returns an error
type DayCommand func(file *os.File) error

// RunWithArgs runs the specified day with the specified cobra arguments
func RunWithArgs(args []string, cmd DayCommand) error {
	file := os.Stdin

	if len(args) >= 1 {
		// Filename given
		var err error
		file, err = os.Open(args[0])
		
		if err != nil {
			return fmt.Errorf("could not open specified file: %w", err)
		}
	}

	defer file.Close()
	return cmd(file)
}