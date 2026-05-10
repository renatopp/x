package main

import (
	"github.com/renatopp/x/strx"
)

func main() {
	t := strx.NewTableWriter(3)
	// t.MetaSection("TITLE")
	t.Meta("ID", "NAME", "AGE")
	t.Data("0", "renato", "39")
	t.Data("1", "maria", "28")
	t.Data("2", "joão", "3").WithAlignments(strx.TableWriterAlignmentRight, strx.TableWriterAlignmentCenter, strx.TableWriterAlignmentRight)
	// t.MetaSection("FOOTER")

	flavors := []strx.TableWriterFlavor{
		strx.TableWriterFlavorAscii,
		strx.TableWriterFlavorAsciiSeparated,
		strx.TableWriterFlavorAsciiCompact,
		strx.TableWriterFlavorAsciiDots,
		strx.TableWriterFlavorGithub,
		strx.TableWriterFlavorReddit,
		strx.TableWriterFlavorRestructuredGrid,
		strx.TableWriterFlavorRestructuredSimple,
		strx.TableWriterFlavorUnicode,
		strx.TableWriterFlavorUnicodeDouble,
	}

	for _, flavor := range flavors {
		t.WithFlavor(flavor)
		println()
		println(t.Render())
	}

	lorem := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."

	println()
	// println(strx.WrapWord(lorem, 23))
	// println(strx.WrapLetter(lorem, 23))
	println(strx.WrapHyphen(lorem, 11))
	println()
	// println(strx.WrapWord("apple, bananas and oranges", 10))
	// println(strx.WrapLetter("apple, bananas and oranges", 10))
	println(strx.WrapHyphen("apple, bananas and oranges", 9))
}
