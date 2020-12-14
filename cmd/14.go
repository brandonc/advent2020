package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/brandonc/advent2020/pkg/tools"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "14 [input file]",
		Short: "Runs the day 14 challenge",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunWithArgs(args, day14)
		},
	})
}

const (
	bitLength = 36 // up to 64
	maskX = byte('X')
	mask1 = byte('1')
	mask0 = byte('0')
	maskforce0 = byte('!')
)

func setMask(mask *[bitLength]byte, repr string) error {
	for i, c := range repr {
		switch(c) {
		case 'X', '1', '0':
			mask[i] = byte(c)
		default:
			return fmt.Errorf("invalid mask char: %q", c)
		}
	}
	return nil
}

func getMaskedValue(mask *[bitLength]byte, value uint64) uint64 {
	var result = value
	for index, i := 0, bitLength - 1; i >= 0; index, i = index + 1, i - 1 {
		if mask[i] == maskX {
			continue
		}
		
		if mask[i] == mask0 {
			var zeromask uint64 = ^(1 << index)
			result &= zeromask
		}

		if mask[i] == mask1 {
			result |= (1 << index)
		}
	}
	return result
}

func printMask(mask *[bitLength]byte) string {
	var result = strings.Builder{}
	for _, b := range mask {
		result.WriteByte(b)
	}
	return result.String()
}

func getMaskedAddresses(mask *[bitLength]byte, addr uint64, result *[]uint64) {
	var value uint64 = addr
	for index, i := 0, bitLength - 1; i >= 0; index, i = index + 1, i - 1 {
		if mask[i] == maskX {
			// make two copies of the mask: one with a 0 and one with a 1 at position i
			// call getMaskedAddresses with each mask, append results to result
			maskVar1, maskVar0 := *mask, *mask
			maskVar1[i] = mask1
			maskVar0[i] = maskforce0

			getMaskedAddresses(&maskVar1, addr, result)
			getMaskedAddresses(&maskVar0, addr, result)
			return
		}
		
		if mask[i] == maskforce0 {
			var zeromask uint64 = ^(1 << index)
			value &= zeromask
		}

		if mask[i] == mask1 {
			value |= (1 << index)
		}
	}
	
	*result = append(*result, value)
}

func getSum(memory map[uint64]uint64) uint64 {
	var sum uint64 = 0
	for _, value := range memory {
		sum += value
	}
	return sum
}

func day14(file *os.File) error {
	scanner, err := tools.Readlines(file)

	if err != nil {
		return fmt.Errorf("could not read input file: %w", err)
	}

	// Initialize mask
	var mask [bitLength]byte
	for i := 0; i < bitLength; i++ {
		mask[i] = maskX
	}

	// Initialize memory
	var memoryPart1 map[uint64]uint64 = make(map[uint64]uint64)
	var memoryPart2 map[uint64]uint64 = make(map[uint64]uint64)
	memSetPattern := *regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

	// Run initialization
	for line := range scanner {
		if strings.HasPrefix(line, "mask = ") {
			setMask(&mask, line[len("mask = "):])
		}
		if strings.HasPrefix(line, "mem[") {
			match := memSetPattern.FindStringSubmatch(line)
			if match == nil || len(match) < 3 {
				return fmt.Errorf("invalid input line: '%s'", line)
			}

			// Already established these as \d+ in the pattern
			addr, _ := strconv.ParseUint(match[1], 10, 64)
			value, _ := strconv.ParseUint(match[2], 10, 64)

			maskedValue := getMaskedValue(&mask, value)

			memoryPart1[addr] = maskedValue
			var addrs []uint64 = make([]uint64, 0, 8)
			getMaskedAddresses(&mask, addr, &addrs)
			
			for _, addrPart2 := range addrs {
				memoryPart2[addrPart2] = value
			}
		}
	}

	fmt.Printf("sum of all values %d (part one)\n", getSum(memoryPart1))
	fmt.Printf("sum of all values %d (part two)\n", getSum(memoryPart2))

	return nil
}