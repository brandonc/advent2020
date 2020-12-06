package day06

import (
	"os"
	"testing"

	"github.com/brandonc/advent2020/tools"
)

// TestRun calls day06.Run with an example
func TestRun(t *testing.T) {
	file := tools.WriteTempFileOrDie(`abc

a
b
c

ab
ac

a
a
a
a

b`);

	defer file.Close()
	defer os.Remove(file.Name())

	Run(file)
}