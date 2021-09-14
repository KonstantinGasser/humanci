package humanci

import (
	"strings"

	"github.com/xlab/treeprint"
)

type NOPNode struct {
	nexts map[Node][]string
}

func newNOPNode() *NOPNode {
	node := &NOPNode{
		nexts: make(map[Node][]string),
	}
	return node
}

func (node *NOPNode) WithNOP(keys ...string) Node {

	for n, edges := range node.nexts {
		for _, edge := range edges {
			for index, key := range keys {
				if edge == key {
					node.nexts[n] = append(node.nexts[n], keys[index+1:]...)
					return n
				}
			}
		}
	}
	n := newNOPNode()
	node.nexts[n] = keys
	return n
}

func (node *NOPNode) Print(tree treeprint.Tree) {
	for node, edges := range node.nexts {
		branch := tree.AddBranch(SliceToString(edges))
		node.Print(branch)
	}
}

func (node *NOPNode) Has(keys ...string) bool {
	// for _, key := range keys {
	// 	if _, ok := node.meta[key]; ok {
	// 		return true
	// 	}
	// }
	return false
}

func (node *NOPNode) Values() map[string]interface{} {
	return nil
}

func (node *NOPNode) Nexts() []Node {
	return nil
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

func keyString(edges map[string]Node) string {
	nodeKey := strings.Builder{}
	nodeKey.WriteString("[")
	keys := []string{}
	for key := range edges {
		keys = append(keys, key)
	}
	nodeKey.WriteString(strings.Join(keys, " | "))
	nodeKey.WriteString("]")
	return nodeKey.String()
}

func SliceToString(edges []string) string {
	nodeKey := strings.Builder{}
	nodeKey.WriteString("[")
	keys := []string{}
	for _, key := range edges {
		keys = append(keys, key)
	}
	nodeKey.WriteString(strings.Join(keys, " | "))
	nodeKey.WriteString("]")
	return nodeKey.String()
}
