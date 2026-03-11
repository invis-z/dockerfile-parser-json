// main.go - Parse Dockerfiles using Docker's own parser
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// Node mirrors the full parser.Node tree for JSON serialization.
type Node struct {
	Value    string   `json:"value,omitempty"`
	Next     *Node    `json:"next,omitempty"`
	Children []*Node  `json:"children,omitempty"`
	Flags    []string `json:"flags,omitempty"`
}

func convertNode(n *parser.Node) *Node {
	if n == nil {
		return nil
	}
	out := &Node{
		Value: n.Value,
		Flags: n.Flags,
	}
	if n.Next != nil {
		out.Next = convertNode(n.Next)
	}
	for _, c := range n.Children {
		out.Children = append(out.Children, convertNode(c))
	}
	return out
}

func main() {
	result, err := parser.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	var nodes []*Node
	for _, n := range result.AST.Children {
		nodes = append(nodes, convertNode(n))
	}

	output, _ := json.Marshal(nodes)
	fmt.Println(string(output))
}
