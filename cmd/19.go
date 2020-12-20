package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/brandonc/advent2020/pkg/tools"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "19 [input file]",
		Short: "Runs the day 19 challenge",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunWithArgs(args, day19)
		},
	})
}

type rule struct {
	match *byte
	subrules [][]string
}

func day19(file *os.File) error {
	scanner, err := tools.Readlines(file)

	if err != nil {
		return fmt.Errorf("could not read input file: %w", err)
	}

	rules := make(map[string]*rule)
	for {
		line := <-scanner
		if line == "" {
			break
		}

		splitColon := strings.Split(line, ":")
		ruleName := splitColon[0]
		splitColon[1] = strings.TrimSpace(splitColon[1])

		if strings.HasPrefix(splitColon[1], "\"") {
			// match rule
			matchChar := splitColon[1][1]
			rules[ruleName] = &rule{
				match: &matchChar,
			}
		} else {
			// ref rule
			splitPipe := strings.Split(splitColon[1], "|")
			
			allsubrules := make([][]string, 0)
			for sr := 0; sr < len(splitPipe); sr++ {
				subrules := make([]string, 0)
				splitRules := strings.Split(splitPipe[sr], " ")
				for r := 0; r < len(splitRules); r++ {
					subrules = append(subrules, strings.TrimSpace(splitRules[r]))
				}
				allsubrules = append(allsubrules, subrules)
			}

			rules[ruleName] = &rule{
				subrules: allsubrules,
			}
		}
	}

	for name, r := range rules {
		fmt.Printf("rule %s: %v\n", name, r)
	}

	return nil
}