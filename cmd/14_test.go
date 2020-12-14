package cmd

import (
	"sort"
	"testing"
)

func TestMaskedValue(t *testing.T) {
	mask := [bitLength]byte{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', '1', 'X', 'X', 'X', 'X', '0', 'X' }
	
	v := getMaskedValue(&mask, 11)
	if v != 73 {
		t.Errorf("mask(11) == %d; wanted 73", v)
	}

	v = getMaskedValue(&mask, 101)
	if v != 101 {
		t.Errorf("mask(101) == %d; wanted 101", v)
	}

	v = getMaskedValue(&mask, 0)
	if v != 64 {
		t.Errorf("mask(0) == %d; wanted 64", v)
	}
}

func TestMaskedAddr(t *testing.T) {
	mask := [bitLength]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', 'X', '0', 'X', 'X' }

	addrs := make([]uint64, 0, 8)
	getMaskedAddresses(&mask, 26, &addrs)

	if len(addrs) != 8 {
		t.Errorf("len(addrs) == %d; wanted 8", len(addrs))
	}

	sort.Slice(addrs, func(i int, j int) bool {
		return addrs[i] < addrs[j]
	})

	if addrs[0] != 16 {
		t.Errorf("addrs[0] == %d; wanted 16", addrs[0])
	}

	if addrs[1] != 17 {
		t.Errorf("addrs[1] == %d; wanted 17", addrs[1])
	}

	if addrs[2] != 18 {
		t.Errorf("addrs[2] == %d; wanted 18", addrs[2])
	}

	if addrs[3] != 19 {
		t.Errorf("addrs[3] == %d; wanted 19", addrs[3])
	}

	if addrs[4] != 24 {
		t.Errorf("addrs[4] == %d; wanted 24", addrs[4])
	}

	if addrs[5] != 25 {
		t.Errorf("addrs[5] == %d; wanted 25", addrs[5])
	}

	if addrs[6] != 26 {
		t.Errorf("addrs[6] == %d; wanted 26", addrs[6])
	}

	if addrs[7] != 27 {
		t.Errorf("addrs[7] == %d; wanted 27", addrs[7])
	}
}