package main

import (
	"fmt"
	"log"
	"os"

	"github.com/brandonc/advent2020/pkg/day01"
	"github.com/brandonc/advent2020/pkg/day02"
	"github.com/brandonc/advent2020/pkg/day03"
	"github.com/brandonc/advent2020/pkg/day04"
	"github.com/brandonc/advent2020/pkg/day05"
	"github.com/brandonc/advent2020/pkg/day06"
	"github.com/brandonc/advent2020/pkg/day07"
)

func printUsage() {
	fmt.Println("Usage: advent nn [input file]")
}

func main() {
	file := os.Stdin

	if len(os.Args) > 2 {
		// Filename given
		var err error
		file, err = os.Open(os.Args[2])
		
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		
		defer file.Close()
	} else if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch(os.Args[1]) {
	case "01":
		day01.Run(file)
	case "02":
		day02.Run(file)
	case "03":
		day03.Run(file)
	case "04":
		day04.Run(file)
	case "05":
		day05.Run(file)
	case "06":
		day06.Run(file)
	case "07":
		day07.Run(file)
	default:
		printUsage()
	}

	os.Exit(0)
}