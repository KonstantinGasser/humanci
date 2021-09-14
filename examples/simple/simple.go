package main

import "github.com/KonstantinGasser/humanci"

func main() {
	cli := humanci.New()
	// cli.RootNOP("what").
	// 	WithNOP("time").
	// 	WithNOP("is").
	// 	WithNOP("it")

	cli.RootNOP("what", "which").
		WithNOP("date", "time").
		WithNOP("is").
		WithNOP("it")

	cli.RootNOP("what").
		WithNOP("date").
		WithNOP("and").
		WithNOP("year").
		WithNOP("is").
		WithNOP("it")

	cli.RootNOP("tell").
		WithNOP("me").
		WithNOP("the").
		WithNOP("time")

	// cli.RootNOP("tell").
	// 	WithNOP("me").
	// 	WithNOP("the").
	// 	WithNOP("date")

	cli.Help()
}

/*
{
	keys: [what]
	subs: {
		"time": -> Node_1,
		"date": -> Node_1,
		"car":  -> Node_2
	}
}

=> what time|date -> #1 command
=> what car	      -> #2 command

add to [what] time,date
=> do I have time | date in subs?
	=> [yes] ->


*/
