package cmd

import (
	"fmt"
	"go/types"
	"os"
	"strconv"
	"strings"

	"github.com/brandonc/advent2020/pkg/tools"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "16 [input file]",
		Short: "Runs the day 16 challenge",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunWithArgs(args, day16)
		},
	})
}

type validRange struct {
	name string
	minA int
	maxA int
	minB int
	maxB int
}

type paperwork struct {
	unassigned []int
	assigned map[string]int;
}

func resolveField(index int, papers []*paperwork, name string) {
	for _, paper := range papers {
		paper.assigned[name] = paper.unassigned[index]
	}
}

func resolveFields(rules []*validRange, papers []*paperwork) map[int]string {
	result := make(map[int]string)
	var fieldCount = len(papers[0].unassigned)

	var noneResolved = false
	for {
		if noneResolved {
			fmt.Printf("could not deduct all fields!")
			break
		}

		if len(result) == len(rules) {
			break
		}

		noneResolved = true
		for i := 0; i < fieldCount; i++ {
			if _, alreadyResolved := result[i]; alreadyResolved {
				continue
			}
			var possible = make(map[*validRange]types.Nil)
			for _, rule := range rules {
				var alreadyResolved = false
				for _, resolvedRule := range result {
					if rule.name == resolvedRule {
						alreadyResolved = true
						break
					}
				}
				if !alreadyResolved {
					possible[rule] = types.Nil{}
				}
			}

			for _, paper := range papers {
				for _, rule := range rules {
					if (paper.unassigned[i] < rule.minA || paper.unassigned[i] > rule.maxA) && (paper.unassigned[i] < rule.minB || paper.unassigned[i] > rule.maxB) {
						delete(possible, rule)
					}
				}
				if len(possible) == 1 {
					noneResolved = false
					for rule := range possible {
						// Loop will be executed once
						result[i] = rule.name
						resolveField(i, papers, rule.name)
					}
					break
				}
			}
		}
	}

	return result
}

func decodePaperwork(rules []*validRange, fields []string) (*paperwork, []int) {
	newPaperwork := &paperwork{
		assigned: make(map[string]int),
		unassigned: make([]int, 0, 16),
	}
	invalidFields := make([]int, 0, 2)
	for _, field := range fields {
		f, _ := strconv.Atoi(field)
		newPaperwork.unassigned = append(newPaperwork.unassigned, f)
		
		fAnyValid := false
		for _, v := range rules {
			if (f >= v.minA && f <= v.maxA) || (f >= v.minB && f <= v.maxB) {
				fAnyValid = true
				break
			}
		}

		if !fAnyValid {
			invalidFields = append(invalidFields, f)
		}
	}

	return newPaperwork, invalidFields
}

func day16(file *os.File) error {
	scanner, _ := tools.Readlines(file)

	valid := make([]*validRange, 0, 4)

	for line := range scanner {
		if line == "" {
			break
		}

		splitLabel := strings.Split(line, ":")
		splitRanges := strings.Split(splitLabel[1], "or")
		
		splitRangeA := strings.Split(splitRanges[0], "-")
		splitRangeB := strings.Split(splitRanges[1], "-")

		minA, _ := strconv.Atoi(strings.TrimSpace(splitRangeA[0]))
		maxA, _ := strconv.Atoi(strings.TrimSpace(splitRangeA[1]))
		minB, _ := strconv.Atoi(strings.TrimSpace(splitRangeB[0]))
		maxB, _ := strconv.Atoi(strings.TrimSpace(splitRangeB[1]))

		valid = append(valid, &validRange{
			name: strings.TrimSpace(splitLabel[0]),
			minA: minA,
			maxA: maxA,
			minB: minB,
			maxB: maxB,
		})
	}

	_ = <-scanner
	myTicket, _ := decodePaperwork(valid, strings.Split(<-scanner, ","))
	_ = <-scanner
	_ = <-scanner

	var sumInvalidFields = 0
	validPaperworks := make([]*paperwork, 0)

	for line := range scanner {
		fields := strings.Split(line, ",")
		newPaperwork, invalidFields := decodePaperwork(valid, fields)

		if len(invalidFields) == 0 {
			validPaperworks = append(validPaperworks, newPaperwork)
		} else {
			for _, invalidField := range invalidFields {
				sumInvalidFields += invalidField
			}
		}
	}

	fmt.Printf("sum of invalid fields = %d (part one)\n", sumInvalidFields)

	resolvedFields := resolveFields(valid, validPaperworks)

	var product uint64 = 1
	for index, name := range resolvedFields {
		if strings.HasPrefix(name, "departure") {
			product *= uint64(myTicket.unassigned[index])
		}
	}

	fmt.Printf("product of departure fields = %d (part two)\n", product)

	return nil
}

