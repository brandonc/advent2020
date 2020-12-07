package tools

import "testing"

func TestGraphSearch(t *testing.T) {
	g := NewGraph("root", 0)

	oneA := g.Root.Add("1st gen", 1)
	oneB := g.Root.Add("1st gen 2", 1)

	twoAA := oneA.Add("2nd gen", 2)
	twoAB := oneA.Add("2nd gen 2", 2)

	twoBA := oneB.Add("2nd gen 3", 2)
	twoBB := oneB.Add("2nd gen 4", 2)

	twoAA.Add("3rd gen", 3)
	twoAB.Add("3rd gen", 3)
	twoBA.Add("3rd gen", 3)
	twoBB.Add("3rd gen", 3)

	found := []string{}
	for s := range g.Search("3rd gen") {
		found = append(found, s.Name)
	}
	
	if len(found) != 4 {
		t.Errorf("len(search) = %d; want 4", len(found))
	}
}