package tools

import "testing"

func TestGraphEdges(t *testing.T) {
	g := NewGraph()

	oneA := g.Add("1st gen")
	oneB := g.Add("1st gen 2")

	twoAA := g.Add("2nd gen")
	twoAB := g.Add("2nd gen 2")
	
	threeAA := g.Add("3rd gen")
	threeAB := g.Add("3rd gen 2")
	threeBA := g.Add("3rd gen 3")
	threeBB := g.Add("3rd gen 4")

	g.AddEdge(oneA, twoAA, 1)
	g.AddEdge(oneB, twoAB, 1)
	g.AddEdge(twoAA, threeAA, 3)
	g.AddEdge(twoAA, threeAB, 4)
	g.AddEdge(twoAB, threeBA, 5)
	g.AddEdge(twoAB, threeBB, 6)

	if len(g.Edges[twoAA]) != 2 {
		t.Errorf("len(twoAA) = %d; want 2", len(g.Edges[twoAA]))
		return
	}
	
	edge1 := g.Edges[twoAA][0]
	edge2 := g.Edges[twoAA][1]

	if edge1.Data != 3 {
		t.Errorf("edge1.Data = %d; want 3", edge1.Data)
	}

	if edge2.Data != 4 {
		t.Errorf("edge2.Data = %d; want 3", edge2.Data)
	}
}

func TestGraphLookup(t *testing.T) {
	g := NewGraph()

	g.Add("one")
	two := g.Add("two")
	g.Add("three")

	check, ok := g.Lookup("two")

	if ok != true {
		t.Errorf("ok = %v; want true", ok)
		return;
	}

	if check.Name != "two" {
		t.Errorf("Name = %s; want two", check.Name)
		return
	}

	if check != two {
		t.Errorf("check = %v; want %v", check, two)
	}
}