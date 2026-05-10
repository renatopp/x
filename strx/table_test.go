package strx_test

import (
	"strings"
	"testing"

	"github.com/renatopp/x/strx"
	"github.com/renatopp/x/testx"
)

func makeTable(s ...string) string {
	return strings.Join(s, "\n")
}

func TestTableEmpty(t *testing.T) {
	b := strx.NewTable()
	testx.Equal(t, "", b.Render())
}

func TestTableDataOnly(t *testing.T) {
	expected := makeTable(
		"+---+---+---+",
		"| a | b | c |",
		"| d | e | f |",
		"+---+---+---+",
	)
	b := strx.NewTable()
	b.Data("a", "b", "c")
	b.Data("d", "e", "f")
	testx.Equal(t, expected, b.Render())
}

func TestTableMetaOnly(t *testing.T) {
	expected := makeTable(
		"+===+===+===+",
		"| a | b | c |",
		"+===+===+===+",
	)
	b := strx.NewTable()
	b.Meta("a", "b", "c")
	b.WithStyle(strx.TableStyleAsciiSeparated)
	testx.Equal(t, expected, b.Render())
}

func TestSectionOnly(t *testing.T) {
	expected := makeTable(
		"+=========+",
		"| SECTION |",
		"+=========+",
	)
	b := strx.NewTable()
	b.MetaSection("SECTION")
	b.WithStyle(strx.TableStyleAsciiSeparated)
	testx.Equal(t, expected, b.Render())
}

func TestMixed(t *testing.T) {
	expected := makeTable(
		"+=========+",
		"|  |",
		"+=========+",
		"| a | b | c |",
		"+---+---+---+",