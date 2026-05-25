package main

import "github.com/renatopp/x/strx"

func main() {
	t := strx.NewTree("Root")
	a := t.Add("a")
	a.Add("b")
	a.Add("c").Add("e")
	t.Add("d")

	// println(a.Render())
}
