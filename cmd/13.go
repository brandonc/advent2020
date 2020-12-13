package cmd

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/brandonc/advent2020/pkg/tools"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "13 [input file]",
		Short: "Runs the day 13 challenge",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunWithArgs(args, day13)
		},
	})
}

type indexBus struct {
	index int
	id *int
}

func readInput(file *os.File) (int, []indexBus, error) {
	scanner, err := tools.Readlines(file)

	if err != nil {
		return 0, nil, fmt.Errorf("could not read input file: %w", err)
	}

	earliest, err := strconv.Atoi(<-scanner)
	
	buses := make([]indexBus, 0, 4)
	for i, busS := range strings.Split(<-scanner, ",") {
		var id *int
		bus, err := strconv.Atoi(busS)
		if err == nil {
			id = &bus
		}
		buses = append(buses, indexBus{index: i, id: id})
	}

	return earliest, buses, nil
}

func filterBuses(buses []indexBus) []int {
	result := make([]int, 0, 4)
	for _, b := range buses {
		if b.id == nil {
			continue
		}

		result = append(result, *b.id)
	}
	return result
}

func day13(file *os.File) error {
	earliest, buses, err := readInput(file)
	
	if err != nil {
		return err
	}

	scheduledBuses := filterBuses(buses)

	ts := earliest
	var depart *int
	for {
		for i, b := range scheduledBuses {
			if ts % b == 0 {
				depart = &scheduledBuses[i]
				break
			}
		}
		if depart != nil {
			break
		}
		ts++
	}

	fmt.Printf("had to wait %d minutes for bus %d, answer is %d (part one)\n", ts - earliest, *depart, (ts - earliest) * (*depart))

	sort.Slice(buses, func(i, j int) bool {
		if buses[i].id != nil && buses[j].id == nil {
			return true
		}
		if buses[i].id == nil && buses[j].id != nil {
			return false
		}
		if buses[i].id == nil && buses[j].id == nil {
			return true
		}
		
		return *buses[i].id > *buses[j].id
	})

	ts = *buses[0].id
	var answer *int
	for {
		found := true
		for _, b := range buses {
			if b.id == nil {
				break
			}
			if (ts - buses[0].index + b.index) % *b.id != 0 {
				found = false
				break
			}
		}
		if found {
			a := ts - buses[0].index
			answer = &a
			break
		}
		ts += *buses[0].id
	}

	if answer != nil {
		fmt.Printf("earliest timestamp is %d (part two)\n", *answer)
	} else {
		return fmt.Errorf("no answer found")
	}

	return nil
}