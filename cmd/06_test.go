package cmd

import (
	"os"
	"testing"

	"github.com/brandonc/advent2020/pkg/tools"
)

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

	if err := day6(file); err != nil {
		t.Error(err)
	}
}