package day09

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/brandonc/advent2020/pkg/tools"
)


type fixedQueue struct {
	Items []int
	Size int
}

func (q *fixedQueue) Append(n int) *int {
	var discarded *int = nil
	if len(q.Items) >= q.Size {
		discarded = &q.Items[0]
		q.Items = q.Items[1:]
	}

	q.Items = append(q.Items, n)
	return discarded
}

func newFixedQueue(size int) *fixedQueue {
	return &fixedQueue{
		Size: size,
		Items: make([]int, 0, size),
	}
}

func canSum(q *fixedQueue, n int) bool {
	for x := 0; x < len(q.Items); x++ {
		for y := 1; y < len(q.Items); y++ {
			if y == x {
				continue
			}

			if q.Items[x] + q.Items[y] == n {
				return true
			}
		}
	}
	return false
}

func minMax(numbers []int) (int, int, error) {
	if len(numbers) == 0 {
		return 0, 0, errors.New("range error: slice was empty")
	}

	min, max := numbers[0], numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max, nil
}

// PrambleSize is the size of the encoding preamble
var PrambleSize = 25

func findFirstNumberWithoutSum(file *os.File) (int, error) {
	scanner, err := tools.Readlines(file)

	if err != nil {
		return 0, err
	}

	queue := newFixedQueue(PrambleSize)

	i := 0
	for line := range scanner {
		num, err := strconv.Atoi(line)

		if err != nil {
			return 0, err
		}

		i++

		if i > queue.Size {
			// Ensure number can be made from a sum of the previous items
			if !canSum(queue, num) {
				return num, nil
			}
		}

		queue.Append(num)
	}

	return 0, errors.New("no answer found")
}

// Run runs the day 9 challenge on the specified input
func Run(file *os.File) {
	answer, err := findFirstNumberWithoutSum(file)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("cannot sum %d from previous %d items (part one)\n", answer, PrambleSize)

	all, err := tools.ReadlinesInts(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	// find longest contiguous sum
	longX := 0
	longY := 0
	for x := 0; x < len(all); x++ {
		curCont := all[x]
		for y := x + 1; y < len(all); y++ {
			curCont += all[y]

			if curCont == answer {
				if y - x > longY - longX {
					longY, longX = y, x
				}
				break
			}

			if curCont > answer {
				// There can be no valid answer beyond this range
				break
			}
		}
	}

	min, max, err := minMax(all[longX:longY + 1])

	if err != nil {
		log.Fatalf("no contiguous sum found in this range: %s", err)
		return
	}	

	fmt.Printf("min %d + max %d = %d (part two)\n", min, max, min + max)
}