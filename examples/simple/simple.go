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
		WithNOP("date", "lol", "date")

	cli.RootNOP("tell").
		WithNOP("me").
		WithNOP("the").
		WithNOP("date", "time")

	cli.Help()
}
