package humanci

import (
	"fmt"
	"strings"

	"github.com/xlab/treeprint"
)

type NOPNode struct {
	keys  map[string]struct{}
	meta  map[string]interface{}
	nexts map[string]Node
}

func newNOPNode(keys ...string) *NOPNode {
	node := &NOPNode{
		keys:  make(map[string]struct{}),
		meta:  make(map[string]interface{}),
		nexts: make(map[string]Node),
	}
	for _, key := range keys {
		node.keys[key] = struct{}{}
	}
	return node
}

func (node *NOPNode) WithNOP(keys ...string) Node {

	var nextNode Node
	for _, key := range keys {

		nextNode.keys[key] = struct{}{}
	}
	if nextNode == nil {
		nextNode = &NOPNode{
			keys:  make(map[string]struct{}),
			meta:  make(map[string]interface{}),
			nexts: make([]Node, 0),
		}
	}
	node.nexts = append(node.nexts, nextNode)
	// fmt.Printf("[%v] with %v -> %v\n", &node, node.keys, node.nexts)
	return nextNode
}

func (node *NOPNode) Print(tree treeprint.Tree) {
	nodeKey := strings.Builder{}
	nodeKey.WriteString("[")
	keys := []string{}
	for key := range node.keys {
		keys = append(keys, key)
	}
	nodeKey.WriteString(strings.Join(keys, " | "))
	nodeKey.WriteString("]")

	branch := tree.AddBranch(nodeKey.String())
	fmt.Printf("%s -> %v\n", nodeKey.String(), node.nexts)
	for _, node := range node.nexts {
		node.Print(branch)
	}
}

func (node *NOPNode) Has(keys ...string) bool {
	for _, key := range keys {
		if _, ok := node.meta[key]; ok {
			return true
		}
	}
	return false
}

func (node *NOPNode) Values() map[string]interface{} {
	return node.meta
}

func (node *NOPNode) Nexts() []Node {
	return node.nexts
}

// Func on NOPNode does nothing since NOPNodes
// don't have any ValueFunc/NodeFunc mapped to it
func (node *NOPNode) Func() error {
	return nil
}

func (node *NOPNode) WithKeys(fn ValueFunc, keys ...string) Node {
	panic("not implemented")
}

func (node *NOPNode) WithRegex(fn ValueFunc, pattern string) Node {
	panic("not implemented")
}

func (node *NOPNode) WithInt(fn ValueFunc) Node {
	panic("not implemented")
}

func (node *NOPNode) WithStr(fn ValueFunc) Node {
	panic("not implemented")
}

func (node *NOPNode) WithExec(fn NodeFunc, keys ...string) Node {
	panic("not implemented")
}
