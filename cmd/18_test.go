package cmd

import (
	"strings"
	"testing"
)

func TestExpr(t *testing.T) {
	a := expr(strings.Split("7+2", ""))
	if a != 9 {
		t.Errorf("a == %d; wanted 9", a)
	}

	b := expr(strings.Split("(7+2)*3", ""))
	if b != 27 {
		t.Errorf("b == %d; wanted 27", b)
	}

	c := expr(strings.Split("7+(2*3)", ""))
	if c != 13 {
		t.Errorf("c == %d; wanted 13", c)
	}

	d := expr(strings.Split("1+2*3+4*5+6", ""))
	if d != 71 {
		t.Errorf("d == %d; wanted 71", d)
	}

	e := expr(strings.Split("((2+4*9)*(6+9*8+6)+6)+2+4*2", ""))
	if e != 13632 {
		t.Errorf("e == %d; wanted 13632", d)
	}
}

func TestExprPart2(t *testing.T) {
	a := exprPart2(strings.Split("1+2*3+4*5+6", ""))
	if a != 231 {
		t.Errorf("a == %d; wanted 231", a)
	}

	b := exprPart2(strings.Split("((2+4*9)*(6+9*8+6)+6)+2+4*2", ""))
	if b != 23340 {
		t.Errorf("b = %d; wanted 23340", b)
	}
}