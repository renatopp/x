package runex_test

import (
	"testing"

	"github.com/renatopp/x/runex"
	"github.com/renatopp/x/testx"
)

func TestIsDigit(t *testing.T) {
	testx.True(t, runex.IsDigit('0'))
	testx.True(t, runex.IsDigit('1'))
	testx.True(t, runex.IsDigit('2'))
	testx.True(t, runex.IsDigit('3'))
	testx.True(t, runex.IsDigit('4'))
	testx.True(t, runex.IsDigit('5'))
	testx.True(t, runex.IsDigit('6'))
	testx.True(t, runex.IsDigit('7'))
	testx.True(t, runex.IsDigit('8'))
	testx.True(t, runex.IsDigit('9'))
	testx.False(t, runex.IsDigit('a'))
	testx.False(t, runex.IsDigit('A'))
	testx.False(t, runex.IsDigit(' '))
	testx.False(t, runex.IsDigit(';'))
}

func TestIsLetter(t *testing.T) {
	testx.True(t, runex.IsLetter('a'))
	testx.True(t, runex.IsLetter('b'))
	testx.True(t, runex.IsLetter('c'))
	testx.True(t, runex.IsLetter('A'))
	testx.True(t, runex.IsLetter('B'))
	testx.True(t, runex.IsLetter('W'))
	testx.True(t, runex.IsLetter('X'))
	testx.False(t, runex.IsLetter('0'))
	testx.False(t, runex.IsLetter('9'))
	testx.False(t, runex.IsLetter(' '))
	testx.False(t, runex.IsLetter(';'))
}

func TestIsWhitespace(t *testing.T) {
	testx.True(t, runex.IsWhitespace(' '))
	testx.True(t, runex.IsWhitespace('\n'))
	testx.True(t, runex.IsWhitespace('\r'))
	testx.False(t, runex.IsWhitespace('a'))
	testx.False(t, runex.IsWhitespace('A'))
	testx.False(t, runex.IsWhitespace('0'))
	testx.False(t, runex.IsWhitespace('9'))
	testx.False(t, runex.IsWhitespace(';'))
}

func TestIsEof(t *testing.T) {
	testx.True(t, runex.IsEof(0))
	testx.False(t, runex.IsEof('a'))
	testx.False(t, runex.IsEof('A'))
	testx.False(t, runex.IsEof('0'))
	testx.False(t, runex.IsEof('9'))
	testx.False(t, runex.IsEof(';'))
}

func TestIsAlphaNumeric(t *testing.T) {
	testx.True(t, runex.IsAlphaNumeric('a'))
	testx.True(t, runex.IsAlphaNumeric('A'))
	testx.True(t, runex.IsAlphaNumeric('0'))
	testx.True(t, runex.IsAlphaNumeric('9'))
	testx.False(t, runex.IsAlphaNumeric(' '))
	testx.False(t, runex.IsAlphaNumeric(';'))
}

func TestIsAnyOf(t *testing.T) {
	testx.True(t, runex.IsOneOf('a', 'a', 'b', 'c'))
	testx.True(t, runex.IsOneOf('b', 'a', 'b', 'c'))
	testx.True(t, runex.IsOneOf('c', 'a', 'b', 'c'))
	testx.False(t, runex.IsOneOf('d', 'a', 'b', 'c'))
	testx.False(t, runex.IsOneOf(' ', 'a', 'b', 'c'))
	testx.False(t, runex.IsOneOf(';', 'a', 'b', 'c'))
}

func TestIsHexadecimal(t *testing.T) {
	testx.True(t, runex.IsHexadecimal('0'))
	testx.True(t, runex.IsHexadecimal('1'))
	testx.True(t, runex.IsHexadecimal('2'))
	testx.True(t, runex.IsHexadecimal('3'))
	testx.True(t, runex.IsHexadecimal('4'))
	testx.True(t, runex.IsHexadecimal('5'))
	testx.True(t, runex.IsHexadecimal('6'))
	testx.True(t, runex.IsHexadecimal('7'))
	testx.True(t, runex.IsHexadecimal('8'))
	testx.True(t, runex.IsHexadecimal('9'))
	testx.True(t, runex.IsHexadecimal('a'))
	testx.True(t, runex.IsHexadecimal('b'))
	testx.True(t, runex.IsHexadecimal('c'))
	testx.True(t, runex.IsHexadecimal('d'))
	testx.True(t, runex.IsHexadecimal('e'))
	testx.True(t, runex.IsHexadecimal('f'))
	testx.True(t, runex.IsHexadecimal('A'))
	testx.True(t, runex.IsHexadecimal('B'))
	testx.True(t, runex.IsHexadecimal('C'))
	testx.True(t, runex.IsHexadecimal('D'))
	testx.True(t, runex.IsHexadecimal('E'))
	testx.True(t, runex.IsHexadecimal('F'))
	testx.False(t, runex.IsHexadecimal(' '))
	testx.False(t, runex.IsHexadecimal(';'))
}

func TestIsOctal(t *testing.T) {
	testx.True(t, runex.IsOctal('0'))
	testx.True(t, runex.IsOctal('1'))
	testx.True(t, runex.IsOctal('2'))
	testx.True(t, runex.IsOctal('3'))
	testx.True(t, runex.IsOctal('4'))
	testx.True(t, runex.IsOctal('5'))
	testx.True(t, runex.IsOctal('6'))
	testx.True(t, runex.IsOctal('7'))
	testx.False(t, runex.IsOctal('8'))
	testx.False(t, runex.IsOctal('9'))
	testx.False(t, runex.IsOctal('a'))
	testx.False(t, runex.IsOctal('A'))
	testx.False(t, runex.IsOctal(' '))
	testx.False(t, runex.IsOctal(';'))
}

func TestIsBinary(t *testing.T) {
	testx.True(t, runex.IsBinary('0'))
	testx.True(t, runex.IsBinary('1'))
	testx.False(t, runex.IsBinary('2'))
	testx.False(t, runex.IsBinary('3'))
	testx.False(t, runex.IsBinary('4'))
	testx.False(t, runex.IsBinary('5'))
	testx.False(t, runex.IsBinary('6'))
	testx.False(t, runex.IsBinary('7'))
	testx.False(t, runex.IsBinary('8'))
	testx.False(t, runex.IsBinary('9'))
	testx.False(t, runex.IsBinary('a'))
	testx.False(t, runex.IsBinary('A'))
	testx.False(t, runex.IsBinary(' '))
	testx.False(t, runex.IsBinary(';'))
}
