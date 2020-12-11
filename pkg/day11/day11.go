package day11

import (
	"log"
	"os"

	"github.com/brandonc/advent2020/pkg/tools"
)

// Run runs the day 11 challenge on the specified input
func Run(file *os.File) {
	_, err := tools.Readlines(file)

	if err != nil {
		log.Fatal(err)
		return
	}


}