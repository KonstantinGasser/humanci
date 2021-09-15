package humanci

import (
	"regexp"

	"github.com/xlab/treeprint"
)

type RegexNode struct {
	meta  MetaData
	fn    ValueFunc
	nexts map[Node][]string
}

func NewRegexNode(fn ValueFunc) func(MetaData) Node {
	return func(meta MetaData) Node {
		return &RegexNode{
			meta:  meta,
			fn:    fn,
			nexts: make(map[Node][]string),
		}
	}
}
func (node *RegexNode) Add(addFn func(MetaData) Node, patterns ...string) Node {
	for n, edges := range node.nexts {
		for _, edge := range edges {
			for index, pattern := range patterns {
				if edge == pattern {
					node.nexts[n] = append(node.nexts[n], patterns[index+1:]...)
					return n
				}
			}
		}
	}
	return addFn(node.Ctx())
}

func (node *RegexNode) WithNOP(keys ...string) Node {
	n := node.Add(NewNOPNode, keys...)
	node.nexts[n] = keys
	return n
}

func (node *RegexNode) WithRegex(fn func(MetaData) error, patterns ...string) Node {
	n := node.Add(NewRegexNode(fn), patterns...)
	node.nexts[n] = patterns
	return n
}

func (node *RegexNode) Exec(cmd []string) error {
	if err := node.fn(node.Ctx()); err != nil {
		return err
	}
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

func (node *RegexNode) Ctx() MetaData {
	return node.meta
}

func (node *RegexNode) WithKeys(fn ValueFunc, keys ...string) Node {
	panic("not implemented")
}

func (node *RegexNode) WithInt(fn ValueFunc) Node {
	panic("not implemented")
}

func (node *RegexNode) WithStr(fn ValueFunc) Node {
	panic("not implemented")
}

func (node *RegexNode) WithExec(fn NodeFunc, keys ...string) Node {
	panic("not implemented")
}

func (node *RegexNode) Print(tree treeprint.Tree) {
	for node, edges := range node.nexts {
		branch := tree.AddBranch(SliceToString(edges))
		node.Print(branch)
	}
}

func (node *RegexNode) Values() map[string]interface{} {
	return nil
}

func (node *RegexNode) Nexts() map[Node][]string {
	return node.nexts
}

func (node *RegexNode) AddToken(v string) {
	// fmt.Println("=========", node.meta["token"])
	// defer fmt.Println("=====", node.meta["token"])

	tokens, ok := node.meta["token"].([]interface{})
	if !ok {
		node.meta["token"] = []string{v}
	}
	node.meta["token"] = append(tokens, v)
}
func (node *RegexNode) Token() []string {
	var s []string
	for _, token := range node.meta["token"].([]interface{}) {
		s = append(s, token.(string))
	}
	return s
}

// Func on RegexNode does nothing since RegexNodes
// don't have any ValueFunc/NodeFunc mapped to it
func (node *RegexNode) Func(v interface{}, token string) error {
	return node.fn(node.Ctx())
}
