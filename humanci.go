package humanci

import (
	"fmt"

	"github.com/xlab/treeprint"
)

// ValueFunc describes any function which recevies some data
// from the current node it gets exectuted on.
// Example:
// 		-> RegexNode, matches regex: ValueFunc(matched_regex_value)
//		-> ValueNode, passed the node's token into the ValueFunc: ValueFunc("foo")
type ValueFunc func(v interface{}) (string, interface{})

// NodeFunc is a function which recevied the current Node as an argument
// allowing the caller to perform a veraity of action.
// Propably only useful for the "last" node to execute the
// final command, while having access to all the stored data
type NodeFunc func(Node) error

type Node interface {
	// WithNOP creates and adds a NOPNode
	// to the command which only exists so that
	// a sentence is grammarly correct. It has no
	// functionality, but to allow clean and correct english/german/world
	// sentences.
	WithNOP(keys ...string) Node

	// WithKeys returns a Node which ValueFunc will be executed when the
	// Node is reached in the trie. This allows to add custom data to
	// the overall context of the command.
	WithKeys(fn ValueFunc, keys ...string) Node

	// WithRegex returns a Node which key is a regex.
	// The node will allow to match every token which
	// matches the pattern.
	// The ValueFunc can be used to add the value with a custom key and
	// will be executed if the Node is reached in the trie.
	WithRegex(fn ValueFunc, pattern string) Node

	// WithInt returns a Node which does not specifiys a key.
	// However, an IntNode matches every token which can be parsed as
	// and int64 and adds the value to the overall data.
	// The ValueFunc can be used to add the value with a custom key and
	// will be executed if the Node is reached in the trie.
	WithInt(fn ValueFunc) Node

	// WithStr returns a Node which does not specifiys a key.
	// However, an StrNode matches EVERY next token and adds
	// the value to the overall data.
	// The ValueFunc can be used to add the value with a custom key and
	// will be executed if the Node is reached in the trie.
	WithStr(fn ValueFunc) Node

	// WithExec adds a ExecNode to the cmd which when reached will execute the
	// node's NodeFunc. The NodeFunc gets the full context of
	// the run a CMD with all the stored data.
	// Tipically this node would be used as the last node of an CMD
	WithExec(fn NodeFunc, keys ...string) Node

	// Values returns the current state of the Cmd's data
	Values() map[string]interface{}

	// Nexts returns all nodes which can be reached from a node
	Nexts() []Node

	// Func executes the mapped ValueFunc or NodeFunc on a node
	// NOPNodes don't have any function, a call to NOPNode.Func() will
	// to nothing
	Func() error

	// Has checks if a node has a given set of keys as identifier
	Has(key ...string) bool

	// Print appends the treeprint.Tree with the node's keys
	// in the form of [key | key]. If the node has nexts Print
	// will be called on each next node.
	Print(tree treeprint.Tree)
}

type cli struct {
	nexts map[Node][]string
}
type CLI interface {
	// Help prints all possible commands as given by the trie
	// as tree to the standrard output
	Help()

	// RootNOP creates and NOPNode as one root node for the
	// cli
	RootNOP(keys ...string) Node
}

// New create new CLI
func New() CLI {
	return &cli{
		nexts: make(map[Node][]string),
	}
}

func (ci *cli) RootNOP(keys ...string) Node {
	// var cached Node
	// for _, key := range keys {
	// 	if cached != nil {
	// 		ci.edges[key] = cached
	// 	}
	// 	if n, ok := ci.edges[key]; ok {
	// 		cached = n
	// 	}
	// }
	// if cached == nil {
	// 	cached = newNOPNode()
	// 	for _, key := range keys {
	// 		ci.edges[key] = cached
	// 	}
	// }

	// Node_1: ["what"]
	// return cached
	for n, edges := range ci.nexts {
		for _, edge := range edges {
			for _, key := range keys {
				if edge == key {
					return n
				}
			}
		}
	}
	node := newNOPNode()
	ci.nexts[node] = keys
	return node
}

func (ci *cli) Help() {

	tree := treeprint.NewWithRoot("*")
	// fmt.Printf("[print]-[%v]\n", ci.edges)
	// key := keyString(ci.nexts)
	for node, edges := range ci.nexts {
		branch := tree.AddBranch(SliceToString(edges))
		node.Print(branch)
	}
	fmt.Println(tree.String())
}

/*
Just here so I can understand how the API should be and the naming convention.
Each type eventually will be moved to its own file + tests :)
*/

type ValueNode struct{}

type ExecNode struct{}

type RegexNode struct{}

type IntNode struct{}

type StrNode struct{}
