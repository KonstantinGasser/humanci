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
	nexts map[string]Node
}
type CLI interface {
	// Help prints all possible commands as given by the trie
	// as tree to the standrard output
	Help()

	// RootNOP creates and NOPNode as one root node for the
	// cli
	RootNOP(roots ...string) Node
}

// New create new CLI
func New() CLI {
	return &cli{
		nexts: make(map[string]Node),
	}
}

func (ci *cli) RootNOP(roots ...string) Node {
	// if node is already present as root node
	// return node
	var next Node
	for _, root := range roots {
		fmt.Printf("cli.add: [%v] | %v\n", roots, ci.nexts)
		if cached, ok := ci.nexts[root]; ok {
			next = cached
			break
		}
	}
	if next == nil {
		next = newNOPNode(roots...)
	}
	for _, root := range roots {
		ci.nexts[root] = next
	}
	return next
}

func (ci *cli) Help() {
	if len(ci.nexts) == 0 {
		fmt.Println("empty Grammar-Trie...")
	}
	tree := treeprint.NewWithRoot("*")

	for _, node := range ci.nexts {
		node.Print(tree)
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
