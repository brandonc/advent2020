package tools

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// Readlines is an iterator that returns one line of a file at a time.
func Readlines(f *os.File) (<-chan string, error) {
	scanner := bufio.NewScanner(f)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	chnl := make(chan string)
	go func() {
		for scanner.Scan() {
			chnl <- scanner.Text()
		}
		close(chnl)
	}()

	return chnl, nil
}

func ScanCSV(data []byte, atEOF bool) (advance int, token []byte, err error) {
	commaidx := bytes.IndexByte(data, ',')
	if commaidx > 0 {
		// we need to return the next position
		buffer := data[:commaidx]
		return commaidx + 1, bytes.TrimSpace(buffer), nil
	}

	// if we are at the end of the string, just return the entire buffer
	if atEOF {
		// but only do that when there is some data. If not, this might mean
		// that we've reached the end of our input CSV string
		if len(data) > 0 {
			return len(data), bytes.TrimSpace(data), nil
		}
	}

	// when 0, nil, nil is returned, this is a signal to the interface to read
	// more data in from the input reader. In this case, this input is our
	// string reader and this pretty much will never occur.
	return 0, nil, nil
}

func ReadlinesInts(f *os.File) ([]int, error) {
	scanner, err := Readlines(f)

	if err != nil {
		return nil, err
	}

	result := []int{}

	for line := range scanner {
		i, err := strconv.Atoi(line)

		if err != nil {
			return nil, err
		}

		result = append(result, i)
	}

	return result, nil
}

func ReadInts(f io.Reader) []int {
	scanner := bufio.NewScanner(f)
	scanner.Split(ScanCSV)

	var numbers []int
	for scanner.Scan() {
		numbers = append(numbers, ToInt(scanner.Text()))
	}
	return numbers
}

func ToInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// WriteTempFileOrDie creates and opens a file using the string provided as its contents
func WriteTempFileOrDie(example string) *os.File {
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if _, err := tmpfile.WriteString(example); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	file, err := os.Open(tmpfile.Name())

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return file
}