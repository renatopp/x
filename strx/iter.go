package strx

import (
	"iter"
	"strings"
)

// IterString returns a Seq that yields each character of the input string as a
// separate string, along with its index. For example, IterString("abc")
// would yield (0, "a"), then (1, "b"), and finally (2, "c").
func IterString(seq string) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		for i, r := range seq {
			if !yield(i, string(r)) {
				return
			}
		}
	}
}

// IterRunes returns a Seq that yields each character of the input string as a
// separate rune, along with its index. For example, IterRunes("abc") would
// yield (0, 'a'), then (1, 'b'), and finally (2, 'c').
func IterRunes(seq string) iter.Seq2[int, rune] {
	return func(yield func(int, rune) bool) {
		for i, r := range seq {
			if !yield(i, r) {
				return
			}
		}
	}
}

func IterFields(s string) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		i := 0
		for word := range strings.FieldsSeq(s) {
			if !yield(i, word) {
				return
			}
			i++
		}
	}
}

func IterFieldsFunc(s string, f func(rune) bool) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		i := 0
		for word := range strings.FieldsFuncSeq(s, f) {
			if !yield(i, word) {
				return
			}
			i++
		}
	}
}

func IterLines(s string) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		i := 0
		for line := range strings.Lines(s) {
			if !yield(i, Trim(line, "\r\n")) {
				return
			}
			i++
		}
	}
}

func IterSplit(s, sep string) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		i := 0
		for part := range strings.SplitSeq(s, sep) {
			if !yield(i, part) {
				return
			}
			i++
		}
	}
}

func IterSplitAfter(s, sep string) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		i := 0
		for part := range strings.SplitAfterSeq(s, sep) {
			if !yield(i, part) {
				return
			}
			i++
		}
	}
}
