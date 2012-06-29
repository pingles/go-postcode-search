package postcode

import "testing"

func TestNodeCreation(t *testing.T) {
	n := makeRootNode()
	if n == nil {
		t.Error("Couldn't create root node")
	}

	for i := 0; i < len(n.children); i++ {
		if n.children[i] != nil {
			t.Error("Children should be empty, contained: ", n.children[i]);
		}
	}
}

func TestAddingChildNodes(t *testing.T) {
	parent := makeRootNode()
	child := addChildNode(parent, "A")
	if parent.children[0] != child {
		t.Error("Should have added child at index 0 of children")
	}
}