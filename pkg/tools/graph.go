package tools

// Graph represents a directed acyclical graph
type Graph struct {
	Root Node
}

// Node is a node of the graph
type Node struct {
	Name string
	Data int
	Parent *Node
	Children []*Node
}

// Add adds a new child to the node
func (n *Node) Add(name string, data int) *Node {
	newNode := Node{
		Name: name,
		Data: data,
		Parent: n,
		Children: make([]*Node, 0, 2),
	}

	n.Children = append(
		n.Children,
		&newNode,
	)

	return &newNode
}

// AddNoData adds a new child to the node with no data
func (n *Node) AddNoData(name string) *Node {
	return n.Add(name, 0)
}

// NewGraph creates a new Graph object
func NewGraph(name string, data int) Graph {
	return Graph{Node{Name: name, Data: data, Children: make([]*Node, 0, 2)}}
}

// Search performs a depth-first search on the Graph and returns
// a collection of matching Nodes
func (g Graph) Search(name string) (<-chan *Node) {
	chnl := make(chan *Node)
	go func() {
		g.Root.search(name, chnl)
		close(chnl)
	}()
	return chnl
}

func (n *Node) search(name string, chnl chan *Node) {
	for _, child := range n.Children {
		child.search(name, chnl)
	}
	if n.Name == name {
		chnl <- n
	}
}