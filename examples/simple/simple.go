package main

import "github.com/KonstantinGasser/humanci"

func main() {
	cli := humanci.New()
	// cli.RootNOP("what").
	// 	WithNOP("time").
	// 	WithNOP("is").
	// 	WithNOP("it")

	cli.RootNOP("what").
		WithNOP("date").
		WithNOP("is").
		WithNOP("it")

	cli.RootNOP("what").
		WithNOP("date")
	// 	WithNOP("and").
	// 	WithNOP("year").
	// 	WithNOP("is").
	// 	WithNOP("it")

	// cli.RootNOP("tell").
	// 	WithNOP("me").
	// 	WithNOP("the").
	// 	WithNOP("time")

	// cli.RootNOP("tell").
	// 	WithNOP("me").
	// 	WithNOP("the").
	// 	WithNOP("date")

	cli.Help()
}
