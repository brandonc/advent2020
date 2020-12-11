package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/brandonc/advent2020/internal/passport"
	"github.com/brandonc/advent2020/pkg/tools"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "4 [input file]",
		Short: "Runs the day 4 challenge",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunWithArgs(args, day4)
		},
	})
}


func day4(file *os.File) error {
	scanner, err := tools.Readlines(file)

	if err != nil {
		return fmt.Errorf("could not read specified input: %w", err)
	}

	var document = make(map[string]string)
	validDocuments := 0
	validDocumentsWithValidFields := 0
	totalDocuments := 0

	for line := range scanner {
		if line == "" {
			totalDocuments++
			if passport.AllFieldsPresent(&document) {
				validDocuments++
			}
			if passport.IsValid(&document) {
				validDocumentsWithValidFields++
			}
			document = make(map[string]string)
			continue
		}
		
		fields := strings.Split(line, " ")

		for _, field := range fields {
			keyValue := strings.Split(field, ":")

			document[keyValue[0]] = keyValue[1]
		}
	}

	totalDocuments++
	if passport.AllFieldsPresent(&document) {
		validDocuments++
	}
	if passport.IsValid(&document) {
		validDocumentsWithValidFields++
	}

	fmt.Printf("There are %d valid documents out of %d (part one)\n", validDocuments, totalDocuments)
	fmt.Printf("There are %d valid documents out of %d (part two)\n", validDocumentsWithValidFields, totalDocuments)

	return nil
}