package humanci

import (
	"strings"

	"github.com/xlab/treeprint"
)

type NOPNode struct {
	keys  map[string]struct{}
	meta  map[string]interface{}
	nexts []Node
}

func newNOPNode(keys ...string) *NOPNode {
	node := &NOPNode{
		keys:  make(map[string]struct{}),
		meta:  make(map[string]interface{}),
		nexts: make([]Node, 0),
	}
	for _, key := range keys {
		node.keys[key] = struct{}{}
	}
	return node
}

func (node *NOPNode) WithNOP(keys ...string) Node {
	nextNode := &NOPNode{
		keys:  make(map[string]struct{}),
		meta:  make(map[string]interface{}),
		nexts: make([]Node, 0),
	}
	for _, key := range keys {
		nextNode.keys[key] = struct{}{}
	}
	node.nexts = append(node.nexts, nextNode)
	return nextNode
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

func (node *NOPNode) Print(tree treeprint.Tree) {
	compositionKey := strings.Builder{}
	compositionKey.WriteString("[")
	for key := range node.keys {
		compositionKey.WriteString(" " + key + " ")
	}
	compositionKey.WriteString("]")

	branch := tree.AddBranch(compositionKey)
	if len(node.nexts) == 0 {
		return
	}

	for _, node := range node.nexts {
		node.Print(branch)
	}
}

// Func on NOPNode does nothing since NOPNodes
// don't have any ValueFunc/NodeFunc mapped to it
func (node *NOPNode) Func() error {
	return nil
}
