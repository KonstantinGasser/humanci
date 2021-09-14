package humanci

import (
	"fmt"
	"testing"
)

func TestNextMapAppend(t *testing.T) {

	nm := NextMap{
		keys:  make([][]string, 0),
		nodes: make([]Node, 0),
	}

	nm.Append("key1", "key2")
	nm.Append("key3")
	fmt.Println(nm.keys)
}
