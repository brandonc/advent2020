package cmd

import "testing"

func TestGame(t *testing.T) {
	s := []int{0, 3, 6}
	fourth := game(s, 4)
	if fourth != 0 {
		t.Errorf("fourth number was %d; wanted 0", fourth)
	}

	fifth := game(s, 5)
	if fifth != 3 {
		t.Errorf("fifth number was %d, wanted 3", fifth)
	}

	sixth := game(s, 6)
	if sixth != 3 {
		t.Errorf("sixth number was %d, wanted 3", sixth)
	}

	seventh := game(s, 7)
	if seventh != 1 {
		t.Errorf("seventh number was %d, wanted 1", seventh)
	}

	eighth := game(s, 8)
	if eighth != 0 {
		t.Errorf("eighth number was %d, wanted 0", eighth)
	}

	nineth := game(s, 9)
	if nineth != 4 {
		t.Errorf("nineth number was %d, wanted 4", nineth)
	}

	tenth := game(s, 10)
	if tenth != 0 {
		t.Errorf("tenth number was %d, wanted 0", tenth)
	}

	twothousandtwentieth := game(s, 2020)
	if twothousandtwentieth != 436 {
		t.Errorf("twothousandtwentieth number was %d, wanted 436", twothousandtwentieth)
	}

	
}