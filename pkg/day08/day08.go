package day08

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/brandonc/advent2020/pkg/handheld"
	"github.com/brandonc/advent2020/pkg/tools"
)

// Run runs the day 08 challenge on the specified input
func Run(file *os.File) {
	reader, err := tools.Readlines(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	lines := make([]string, 0, 16)
	for line := range reader {
		lines = append(lines, line)
	}

	h, err := handheld.NewHandheld(lines)

	if err != nil {
		log.Fatal(err)
		return
	}

	v, loop := h.RunUntilLoopDetected()

	if (loop) {
		fmt.Printf("the value of the accumulator after loop is %d (part one)\n", v)
	} else {
		log.Fatal("No loop was detected in this instruction")
		return
	}

	for change := 0; change < len(lines); change++ {
		modified := make([]string, len(lines))
		for i := 0; i < len(lines); i++ {
			if change == i {
				split := strings.Fields(lines[i])
				if split[0] == "jmp" {
					modified[i] = fmt.Sprintf("%s %s", "nop", split[1])
				} else if split[0] == "nop" {
					modified[i] = fmt.Sprintf("%s %s", "jmp", split[1])
				} else {
					modified[i] = lines[i]
				}
			} else {
				modified[i] = lines[i]
			}
		}
		variant, err := handheld.NewHandheld(modified)

		if err != nil {
			log.Fatal(err)
			return
		}

		variantAcc, variantLoop := variant.RunUntilLoopDetected()

		if !variantLoop {
			fmt.Printf("the value of the accumulator at the end is %d (part two)\n", variantAcc)
			return
		}
	}

	log.Fatal("no way to change the program to not end in a loop!")
}