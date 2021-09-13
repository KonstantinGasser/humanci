package humanci

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

	// WithExec adds a node to the cmd which will trigger the mapped
	// function. WithExec should be called on either of the
	// last nodes or if a node should perform a task in prior to the
	// actual requested command. The NodeFunc gets the full context of
	// the run CMD with all the stored data
	WithExec(fn NodeFunc, keys ...string) Node

	// Values returns the current state of the Cmd's data
	Values() map[string]interface{}
}

type NOPNode struct{}

type ValueNode struct{}

type ExecNode struct{}

type RegexNode struct{}

type IntNode struct{}

type StrNode struct{}
