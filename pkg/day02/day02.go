package day02

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/brandonc/advent2020/tools"
)

// Run runs the day 02 challenge on the specified input
func Run(file *os.File) {
	regex := *regexp.MustCompile(`^(\d+)-(\d+) ([a-z]): ([a-z]+)$`)

	valid1st := 0
	valid2nd := 0

	scanner, err := tools.Readlines(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	
	for line := range scanner {
		matches := regex.FindStringSubmatch(line)
		
		min, err := strconv.Atoi(matches[1])

		if err != nil {
			log.Fatal(err)
			return
		}
		
		max, err := strconv.Atoi(matches[2])
		
		if err != nil {
			log.Fatal(err)
			return
		}

		exhibit := matches[3][0]
		passwd := matches[4]

		found := 0

		for _, c := range passwd {
			if c == rune(exhibit) {
				found++
			}
		}

		if found >= min && found <= max {
			valid1st++
		}

		if (passwd[min - 1] == exhibit || passwd[max - 1] == exhibit) && !(passwd[min - 1] == exhibit && passwd[max - 1] == exhibit) {
			valid2nd++
		}
	}

	fmt.Printf("There are %d valid passwords (1st part)\n", valid1st)
	fmt.Printf("There are %d valid passwords (2nd part)\n", valid2nd)
}