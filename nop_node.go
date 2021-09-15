package humanci

import (
	"regexp"
	"strings"

	"github.com/xlab/treeprint"
)

type NOPNode struct {
	meta  MetaData
	nexts map[Node][]string
}

func NewNOPNode(meta MetaData) Node {
	return &NOPNode{
		meta:  meta,
		nexts: make(map[Node][]string),
	}
}

func (node *NOPNode) Add(addFn func(meta MetaData) Node, keys ...string) Node {
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
	return addFn(node.Ctx())
}

func (node *NOPNode) WithNOP(keys ...string) Node {
	n := node.Add(NewNOPNode, keys...)
	node.nexts[n] = keys
	return n
}

func (node *NOPNode) WithKeys(fn ValueFunc, keys ...string) Node {
	panic("not implemented")
}

func (node *NOPNode) WithRegex(fn func(MetaData) error, patterns ...string) Node {
	n := node.Add(NewRegexNode(fn), patterns...)
	node.nexts[n] = patterns
	return n
}

func (node *NOPNode) Ctx() MetaData {
	return node.meta
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

func (node *NOPNode) Print(tree treeprint.Tree) {
	for node, edges := range node.nexts {
		branch := tree.AddBranch(SliceToString(edges))
		node.Print(branch)
	}
}

func (node *NOPNode) Values() map[string]interface{} {
	return nil
}

func (node *NOPNode) Nexts() map[Node][]string {
	return node.nexts
}

// Func on NOPNode does nothing since NOPNodes
// don't have any ValueFunc/NodeFunc mapped to it
func (node *NOPNode) Func(v interface{}, token string) error {
	return nil
}

func (node *NOPNode) AddToken(v string) {
	// fmt.Println("=========", node.meta["token"])
	// defer fmt.Println("=====", node.meta["token"])

	tokens, ok := node.meta["token"].([]interface{})
	if !ok {
		node.meta["token"] = []string{v}
	}
	node.meta["token"] = append(tokens, v)
}
func (node *NOPNode) Token() []string {
	var s []string
	for _, token := range node.meta["token"].([]interface{}) {
		s = append(s, token.(string))
	}
	return s
}

func (node *NOPNode) Exec(cmd []string) error {
	for n, edges := range node.nexts {
		for _, edge := range edges {
			if regexp.MustCompile(edge).MatchString(cmd[0]) {
				n.AddToken(cmd[0])
				return n.Exec(cmd[1:])
			}
		}
	}
	return nil
}

func SliceToString(edges []string) string {
	nodeKey := strings.Builder{}
	nodeKey.WriteString("[")

	var keys []string
	keys = append(keys, edges...)

	nodeKey.WriteString(strings.Join(keys, " | "))
	nodeKey.WriteString("]")
	return nodeKey.String()
}
