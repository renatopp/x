package main

import (
	"github.com/renatopp/x/strx"
)

func main() {
	t := strx.NewTable()
	t.MetaSection("THIS IS A VERY LARGE TITLE THAT SHOULD BE BROKEN INTO MULTIPLE LINES")
	t.Meta("ID", "NAME", "AGE")
	t.Data("0", "renato\npereira", "39")
	t.Data("1", "maria", "28")
	t.DataSection("--")
	t.Data("2", "joão", "3")
	t.MetaSection("FOOTER")
	t.WithLength(10, 20, 5)
	// t.WithStyle(strx.TableStyleUnicode)
	println(t.Render())
}
