package day04

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/brandonc/advent2020/tools"
)

// Run runs the day 4 advent of code challenge
func Run(file *os.File) {
	scanner, err := tools.Readlines(file)

	if err != nil {
		log.Fatal(err)
		return
	}

	var document = make(map[string]string)
	validDocuments := 0
	validDocumentsWithValidFields := 0
	totalDocuments := 0

	for line := range scanner {
		if line == "" {
			totalDocuments++
			if AllFieldsPresent(&document) {
				validDocuments++
			}
			if IsValid(&document) {
				validDocumentsWithValidFields++
			}
			document = make(map[string]string)
			continue
		}
		
		fields := strings.Split(line, " ")

		for _, field := range fields {
			keyValue := strings.Split(field, ":")

			document[keyValue[0]] = keyValue[1]
		}
	}

	totalDocuments++
	if AllFieldsPresent(&document) {
		validDocuments++
	}
	if IsValid(&document) {
		validDocumentsWithValidFields++
	}

	fmt.Printf("There are %d valid documents out of %d (first part)\n", validDocuments, totalDocuments)
	fmt.Printf("There are %d valid documents out of %d (second part)\n", validDocumentsWithValidFields, totalDocuments)
}