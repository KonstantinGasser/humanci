package humanci

type ValueFunc func(v interface{}) (string, interface{})

type NodeFunc func(Node) error

type Node interface {
	// Fill creates and adds a Fill-Word node
	// to the command which only exists so that
	// a sentence is grammarly correct. It has no
	// functionality
	NOP(keys ...string) Node

	// Value adds a stateful node to the command
	// in which CmdFunc additional values can be added to
	// the Node data map
	Value(fn ValueFunc, keys ...string) Node

	// Exec adds a node to the cmd which will trigger the mapped
	// function. Exec should be called on either of the
	// last nodes or if a node should perform a task in prior to the
	// actual requested command
	Exec(fn NodeFunc, keys ...string) Node

	// Regex returns a Node which key is a regex.
	// The node will allow to match every token which
	// matches the pattern.
	// The ValueFunc can be used to add the value with a custom key.
	Regex(fn ValueFunc, pattern string) Node

	// Int returns a Node which does not specifiys a key.
	// However, an IntNode matches every token which can be parsed as
	// and int64 and adds the value to the overall data.
	// The ValueFunc can be used to add the value with a custom key.
	Int(fn ValueFunc) Node

	// Str returns a Node which does not specifiys a key.
	// However, an StrNode matches EVERY next token and adds
	// the value to the overall data.
	// The ValueFunc can be used to add the value with a custom key.
	Str(fn ValueFunc) Node

	// Values returns the current state of the Cmd's data
	Values() map[string]interface{}
}

type NOPNode struct{}

type ValueNode struct{}

type ExecNode struct{}

type RegexNode struct{}

type IntNode struct{}

type StrNode struct{}
