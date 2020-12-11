package cmd

import (
	"fmt"
	"go/types"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/brandonc/advent2020/pkg/tools"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "7 [input file]",
		Short: "Runs the day 7 challenge",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunWithArgs(args, day7)
		},
	})
}

type countedBag struct {
	number int
	name   string
}

type decoded struct {
	bag      string
	contains []countedBag
}

func decode(raw string) decoded {
	split := strings.Split(raw, " bags contain ")
	contentsBags := make([]countedBag, 0, 2)

	if split[1] != "no other bags." {
		contentsStrings := strings.Split(split[1], ", ")
		contentsSuffix := regexp.MustCompile(`bags?\.?`)

		for _, contentsString := range contentsStrings {
			contentsRaw := strings.SplitN(contentsString, " ", 2)
			count, err := strconv.Atoi(contentsRaw[0])
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			contentsBags = append(contentsBags, countedBag{number: count, name: strings.TrimSpace(contentsSuffix.ReplaceAllString(contentsRaw[1], ""))})
		}
	}

	return decoded{bag: split[0], contains: contentsBags}
}

func parseGraph(input <-chan string) *tools.Graph {
	g := tools.NewGraph()

	for line := range input {
		// shiny purple bags contain 2 pale blue bags, 1 wavy fuchsia bag, 5 pale salmon bags.
		decoded := decode(line)

		parent := g.LookupOrAdd(decoded.bag)

		for _, contents := range decoded.contains {
			child := g.LookupOrAdd(contents.name)
			g.AddEdge(parent, child, contents.number)
		}
	}

	return g
}

func descend(name string, g *tools.Graph, n *tools.Node, prev *map[*tools.Node]types.Nil) bool {
	if n.Name == name {
		return true
	}
	found := false
	for _, edge := range g.Edges[n] {
		if descend(name, g, edge.Child, prev) {
			(*prev)[edge.Parent] = types.Nil{}
			found = true
		}
	}
	return found
}

func sumContents(g *tools.Graph, parent *tools.Node) int {
	edges := g.Edges[parent]
	sum := 0
	for _, edge := range edges {
		sum += edge.Data + (edge.Data * sumContents(g, edge.Child))
	}
	return sum
}

func day7(file *os.File) error {
	scanner, err := tools.Readlines(file)

	if err != nil {
		return fmt.Errorf("could not read input file: %w", err)
	}

	graph := parseGraph(scanner)
	sg, ok := graph.Lookup("shiny gold")

	if !ok {
		return fmt.Errorf("Input does not contain a shiny gold bag")
	}

	found := make(map[*tools.Node]types.Nil)
	for _, node := range graph.Nodes {
		breadcrumbs := make(map[*tools.Node]types.Nil)
		if node.Name == "shiny gold" {
			continue
		}
		if descend("shiny gold", graph, node, &breadcrumbs) {
			for b := range breadcrumbs {
				found[b] = types.Nil{}
			}
		}
	}

	fmt.Printf("there are %d valid containers (part one)\n", len(found))
	fmt.Printf("shiny gold bag must contain %d other bags (part two)\n", sumContents(graph, sg))

	return nil
}
