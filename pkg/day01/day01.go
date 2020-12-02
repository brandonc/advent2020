package day01

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/brandonc/advent2020/tools"
)

// Run runs day01 challenge using the specified input
func Run(file *os.File) {
	sorted, err := tools.ReadlinesInts(file)

	if err != nil {
		log.Fatal(err)
		return
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
					fmt.Printf("%d + %d + %d = 2020\n", sorted[aIndex], sorted[zIndex], sorted[xIndex])
					fmt.Printf("%d * %d * %d = %d\n", sorted[aIndex], sorted[zIndex], sorted[xIndex], sorted[aIndex] * sorted[zIndex] * sorted[xIndex])
					return
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

	fmt.Println("No 2020 sum found")
}