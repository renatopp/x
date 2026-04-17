// Package strx provides various string manipulation functions that are not
// included in the standard library's strings package. It also reexports all
// functions from the strings package for convenience.
package strx

import (
	"iter"
	"strings"
	"unicode"
)

// Ident returns a string with the specified number of spaces
// inserted before the string.
func Ident(s string, pad int) string {
	ident := strings.Repeat(" ", pad)
	return ident + strings.ReplaceAll(s, "\n", "\n"+ident)
}

// IdentWith returns a string with the specified number of characters
// inserted before the string. The character used for padding is
// specified by the second parameter.
func IdentWith(s string, pad int, with string) string {
	ident := strings.Repeat(with, pad)
	return ident + strings.ReplaceAll(s, "\n", "\n"+ident)
}

// Escape replaces newlines and tabs in a string with their escaped
// representations. It returns the modified string.
func Escape(s string) string {
	s = strings.ReplaceAll(s, "\r", "\\r")
	s = strings.ReplaceAll(s, "\n", "\\n")
	s = strings.ReplaceAll(s, "\t", "\\t")
	return s
}

// JoinFunc is a helper function that applies a function to each element of the
// slice and then joins the resulting strings with the specified separator. It
// takes a slice of items, a function that takes an item and returns a string,
// and a separator string. It returns a string that is the result of calling
// Join with the resulting strings.
func JoinFunc[T any](items []T, fn func(T) string, sep string) string {
	strs := make([]string, len(items))
	for i, item := range items {
		strs[i] = fn(item)
	}
	return strings.Join(strs, sep)
}

// HumanList joins a slice of strings with "and" or "or" depending on the value of
// the last parameter. If the slice is empty, it returns an empty string. If the
// slice has one element, it returns that element. If the slice has two elements,
// it joins them with "and" or "or". If the slice has more than two elements,
// it joins them with ", " and the last element with "and" or "or".
func HumanList(items []string, and string) string {
	if len(items) == 0 {
		return ""
	}

	if len(items) == 1 {
		return items[0]
	}

	res := strings.Join(items[:len(items)-1], ", ") + " " + and + " " + items[len(items)-1]
	return res
}

// HumanListFunc is a helper function that applies a function to each element of
// the slice and then calls HumanList with the resulting strings. It takes a
// slice of items, a function that takes an item and returns a string, and a
// string to use for the last element. It returns a string that is the result of
// calling HumanList with the resulting strings.
func HumanListFunc[T any](items []T, fn func(T) string, and string) string {
	strs := make([]string, len(items))
	for i, item := range items {
		strs[i] = fn(item)
	}
	return HumanList(strs, and)
}

// Pads the string with spaces to the left until the string reaches the desired
// length.
//
// If the string is already longer than the desired length, it will be returned
// as is.
func PadLeft(s string, n int) string {
	return PadLeftWith(s, n, " ")
}

// Pads the string with the specified character to the left until the string
// reaches the desired length.
//
// If the string is already longer than the desired length, it will be returned
// as is.
func PadLeftWith(s string, n int, with string) string {
	n = n - len(s)
	if n <= 0 {
		return s
	}
	return strings.Repeat(with, n) + s
}

// Pads the string with spaces to the right until the string reaches the
// desired
//
// If the string is already longer than the desired length, it will be returned
// as is.
func PadRight(s string, n int) string {
	return PadRightWith(s, n, " ")
}

// Pads the string with the specified character to the right until the string
// reaches the desired length.
//
// If the string is already longer than the desired length, it will be returned
// as is.
func PadRightWith(s string, n int, with string) string {
	n = n - len(s)
	if n <= 0 {
		return s
	}
	return s + strings.Repeat(with, n)
}

// Pads the string with spaces to the left and right until the string reaches
// the desired length. In case of an odd number of characters, the left side
// will have one more character than the right side.
//
// If the string is already longer than the desired length, it will be returned
// as is.
func PadCenter(s string, n int) string {
	return PadCenterWith(s, n, " ")
}

// Pads the string with the specified character to the left and right until the
// string reaches the desired length. In case of an odd number of characters,
// the left side will have one more character than the right side.
//
// If the string is already longer than the desired length, it will be returned
// as is.
func PadCenterWith(s string, n int, with string) string {
	n = n - len(s)
	if n <= 0 {
		return s
	}
	left := n / 2
	right := left
	if left%2 == 1 {
		left++
	}
	return strings.Repeat(with, left) + s + strings.Repeat(with, right)
}

// FirstUp capitalizes the first letter of the string and returns the modified
// string. If the string is empty, it will be returned as is.
func FirstUp(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

// FirstLow lowercases the first letter of the string and returns the modified
// string. If the string is empty, it will be returned as is.
func FirstLow(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(string(s[0])) + s[1:]
}

// TrimSpaces removes all spaces from the start and end of the string.
func TrimSpaces(s string) string {
	return strings.Trim(s, " ")
}

// IsBlank returns true if the string is empty or contains only whitespace
// characters.
func IsBlank(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// Eliipsis truncates the string to the specified maximum length and appends
// "..." if the string was longer than the maximum length. If the string is
// shorter than or equal to the maximum length, it will be returned as is.
func Eliipsis(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
}

// IterString returns a Seq that yields each character of the input string as a
// separate string. For example, IterString("abc") would yield "a", then "b", and finally "c".
func IterString(seq string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, r := range seq {
			if !yield(string(r)) {
				return
			}
		}
	}
}

// IterRunes returns a Seq that yields each character of the input string as a
// separate rune. For example, IterRunes("abc") would yield 'a', then 'b', and finally 'c'.
func IterRunes(seq string) iter.Seq[rune] {
	return func(yield func(rune) bool) {
		for _, r := range seq {
			if !yield(r) {
				return
			}
		}
	}
}
