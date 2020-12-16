package cmd

import (
	"fmt"
	"os"

	"github.com/brandonc/advent2020/pkg/tools"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "15 [input file]",
		Short: "Runs the day 15 challenge",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunWithArgs(args, day15)
		},
	})
}

var memo = make(map[int]*tools.LinkedListNode)

func game(start []int, turns int) int {
	if len(start) == 0 {
		return 0
	}

	if turns < len(start) {
		return start[turns - 1]
	}

	tail := tools.NewLinkedList(start[0], 1)
	memo[start[0]] = tail
	for si, n := range start[1:] {
		tail = tail.Insert(n, si + 2)
		memo[n] = tail
	}

	var i = len(start)
	for {
		if i == turns {
			break
		}
		lastNumber := tail.Value

		var lastSeen = 1
		if i % 50000 == 0 {
			fmt.Printf("...%f%%, keysize = %d\n", (float32(i) / float32(turns)) * 100.0, len(memo))
		}
		memoNode, exists := memo[lastNumber];
		if exists && memoNode != tail {
			lastSeen = tail.Weight - memoNode.Weight
		} else {
			if !exists {
				lastSeen = 0
			} else {
				searcher := tail.Prev
				for {
					if searcher == nil {
						lastSeen = 0
						break
					}
					if searcher.Value == lastNumber {
						break
					}
					
					searcher = searcher.Prev
					lastSeen++
				}
			}
		}
		tail = tail.Insert(lastSeen, i)
		memo[lastSeen] = tail
		i++
	}
	return tail.Value
}

func day15(file *os.File) error {
	start, err := tools.ReadlinesInts(file)

	if err != nil {
		return fmt.Errorf("could not read input file: %w", err)
	}

	fmt.Printf("the 2020th number spoken is %d (part one)\n", game(start, 2020))
	fmt.Printf("the 30000000th number spoken is %d (part two)\n", game(start, 30000000))

	return nil
}