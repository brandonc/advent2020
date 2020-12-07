package tools

// Graph represents a directed acyclical graph
type Graph struct {
	Nodes []*Node
	Edges map[*Node][]*Edge
}

// Node is a node of the graph
type Node struct {
	Name string
}

// Edge is an edge of the graph
type Edge struct {
	Parent *Node
	Child *Node
	Data int
}

// Add adds a new child to the graph
func (g *Graph) Add(name string) *Node {
	newNode := Node{Name: name}
	g.Nodes = append(g.Nodes, &newNode)
	return &newNode
}

// LookupOrAdd adds a node if it does not exist
func (g *Graph) LookupOrAdd(name string) *Node {
	known, ok := g.Lookup(name)
	if !ok {
		return g.Add(name)
	}
	return known
}

// AddEdge adds a new edge to a pair of nodes
func (g *Graph) AddEdge(n1 *Node, n2 *Node, data int) *Edge {
	newEdge := Edge{Parent: n1, Child: n2, Data: data}
	if g.Edges == nil {
		g.Edges = make(map[*Node][]*Edge)
	}

	if g.Edges[n1] == nil {
		g.Edges[n1] = make([]*Edge, 0, 2)
	}

	g.Edges[n1] = append(g.Edges[n1], &newEdge)
	return &newEdge
}

// Lookup returns a node by name
func (g *Graph) Lookup(name string) (*Node, bool) {
	for _, n := range g.Nodes {
		if n.Name == name {
			return n, true
		}
	}
	return nil, false
}

// NewGraph creates a new Graph object
func NewGraph() *Graph {
	return &Graph{Nodes: make([]*Node, 0, 16)}
}
