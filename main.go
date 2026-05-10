package main

import (
	"github.com/renatopp/x/strx"
)

func main() {
	t := strx.NewTable()
	t.MetaSection("THIS IS A VERY LARGE TITLE THAT SHOULD BE BROKEN INTO MULTIPLE LINES")
	t.Meta("ID", "NAME", "AGE")
	t.Data("0", "renato\nr2p.dev", 39).ToCenter()
	t.Data("1", "maria", 28).ToRight()
	t.DataSection("This is another very large section that should be broken into multiple lines")
	t.Data("2", "joão", 3).WithAlignments(strx.TableLeft, strx.TableRight, strx.TableCenter)
	t.MetaSection("FOOTER")
	t.WithLength(10, 20, 5)
	t.WithStyle(strx.TableStyleUnicodeDoubleGrid)
	println(t.Render())
}
