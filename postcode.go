package postcode

import (
	"strings"
)

type Node struct {
	letter string
	region int
	children [26 + 10]*Node
}

func buildNode(letter string, region int) (node *Node) {
	children := [26 + 10]*Node {}
	return &Node {letter, region, children}
}

func makeRootNode() (node *Node) {
	return buildNode("", 0)
}

func addChildNode(parent *Node, letter string) (node *Node) {
	up := strings.ToUpper(letter)
	node = buildNode(up, 0)
	index := strings.IndexAny("ABCDEFGHIJKLMNOPQRSTUVWXY0123456789", up)
	parent.children[index] = node
	return node
}