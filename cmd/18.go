package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/brandonc/advent2020/pkg/tools"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "18 [input file]",
		Short: "Runs the day 18 challenge",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunWithArgs(args, day18)
		},
	})
}

func paren(ex tokens, exF func (ex tokens) uint) (uint, tokens, error) {
	nesting := 0
	for i := 0; i < len(ex); i++ {
		if ex[i] == tokenRparen {
			if nesting == 0 {
				return exF(ex[0:i]), ex[i+1:], nil
			}
			nesting--
		}
		if ex[i] == tokenLparen {
			nesting++
		}
	}

	return 0, nil, fmt.Errorf("expected rparen")
}

func expr(ex tokens) uint {
	var (
		result uint
		opstr string
		val uint
		err error
		next string
	)

	for {
		if len(ex) == 0 {
			return result
		}

		next, ex = shift(ex)
		switch next {
		case tokenLparen:
			val, ex, err = paren(ex, expr)
			if err != nil {
				fmt.Printf("invalid input: %s", err)
				return 0
			}
		case tokenAdd, tokenProduct:
			opstr = next
			continue
		default:
			// its a number
			ival, err := strconv.Atoi(next)
			if err != nil {
				fmt.Printf("invalid input: %s", err)
				return 0
			}
			val = uint(ival)
		}

		if opstr == "" {
			// lvalue
			result = uint(val)
		} else {
			if opstr == tokenAdd {
				result = result + uint(val)
			}
			if opstr == tokenProduct {
				result = result * uint(val)
			}
			opstr = ""
		}
	}
}

func exprPart2(ex tokens) uint {
	var (
		result uint
		opstr string
		val uint
		err error
		next string
	)

	for {
		if len(ex) == 0 {
			return result
		}

		next, ex = shift(ex)
		switch next {
		case tokenLparen:
			val, ex, err = paren(ex, exprPart2)
			if err != nil {
				fmt.Printf("invalid input: %s", err)
				return 0
			}
		case tokenProduct:
			return result * exprPart2(ex)
		case tokenAdd:
			opstr = tokenAdd
			continue
		default:
			// its a number
			ival, err := strconv.Atoi(next)
			if err != nil {
				fmt.Printf("invalid input: %s", err)
				return 0
			}
			val = uint(ival)
		}

		if opstr == "" {
			result = val
		}
		if opstr == tokenAdd {
			result = result + uint(val)
		}
	}
}

const tokenLparen = "("
const tokenRparen =  ")"
const tokenAdd = "+"
const tokenProduct = "*"

type tokens []string

func shift(t tokens) (string, tokens) {
	return t[0], t[1:]
}

func peek(t tokens) string {
	return t[0]
}

func tokenize(s string) tokens {
	result := make([]string, 0, len(s) / 2 + 1)
	for _, c := range s {
		var n string = ""
		switch c {
		case ' ':
			continue
		case '(':
			n = tokenLparen
		case ')':
			n = tokenRparen
		case '+':
			n = tokenAdd
		case '*':
			n = tokenProduct
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			n = string(c)
		}
		if n != "" {
			result = append(result, n)
		}
	}
	return result
}


func day18(file *os.File) error {
	scanner, err := tools.Readlines(file)

	if err != nil {
		return fmt.Errorf("could not read input file: %w", err)
	}

	var sum uint = 0
	var sumPart2 uint = 0
	for line := range scanner {
		sum += expr(tokenize(line))
		sumPart2 += exprPart2(tokenize(line))
	}

	fmt.Printf("sum of all lines is %d (part one)\n", sum)
	fmt.Printf("sum of all lines is %d (part two)\n", sumPart2)

	return nil
}