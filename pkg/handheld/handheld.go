package handheld

import (
	"strconv"
	"strings"
)

type instruction struct {
	instr string
	value int
}

// Handheld represents a bootable handheld computer
type Handheld struct {
	bootSeq []*instruction
	pointer int
}

// NewHandheld creates a new handheld struct
func NewHandheld(lines []string) (*Handheld, error) {
	result := Handheld{bootSeq: make([]*instruction, 0, 16), pointer: 0}

	for _, line := range lines {
		split := strings.Fields(line)
		value, err := strconv.Atoi(strings.TrimLeft(split[1], "+"))

		if err != nil {
			return nil, err
		}

		i := instruction{instr: split[0], value: value }
		result.bootSeq = append(result.bootSeq, &i)
	}

	return &result, nil
}

// Step executes the next step in the boot seq and increments the pointer
func (h *Handheld) Step() (int, bool) {
	next := h.bootSeq[h.pointer]

	if next.instr == "nop" {
		h.pointer++
	}
	
	if next.instr == "acc" {
		h.pointer++
		return next.value, h.pointer >= len(h.bootSeq)
	}

	if next.instr == "jmp" {
		h.pointer += next.value
	}

	return 0, h.pointer >= len(h.bootSeq)
}

// RunUntilLoopDetected returns the accumulated value at the moment an infinute loop is detected
func (h *Handheld) RunUntilLoopDetected() (int, bool) {
	instructions := make([]int, len(h.bootSeq))
	for i := 0; i < len(instructions); i++ {
		instructions[i] = 0
	}

	acc := 0
	for {
		if instructions[h.pointer] > 0 {
			break
		}
		instructions[h.pointer]++
		v, done := h.Step()
		acc += v

		if done {
			return acc, false
		}
	}
	return acc, true
}