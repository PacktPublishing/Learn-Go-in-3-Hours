package leftpad

import (
	"strings"
	"unicode/utf8"
)

var default_char = ' '

// Format takes in a string and an int and returns the string
// left-padded with spaces to the length of the int. If the
// string is already longer than the specified length, the
// original string is returned.
func Format(s string, size int) string {
	return FormatRune(s, size, default_char)
}

// FormatRune takes in a string, an int, and a rune and returns the string
// left-padded with the specifed rune to the length of the int. If the
// string is already longer than the specified length, the
// original string is returned.
func FormatRune(s string, size int, r rune) string {
	l := utf8.RuneCountInString(s)
	if l >= size {
		return s
	}
	out := strings.Repeat(string(r), size-l) + s
	return out
}
