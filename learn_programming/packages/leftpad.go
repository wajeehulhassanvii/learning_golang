package leftpad

import (
	"unicode/utf8"
)

var defaultChar = ' '

// Format func
func Format(s string, size int) string {
	return FormatRune(s, size, defaultChar)
}

// FormatRune func
func FormatRune(s string, size int, r rune) string {
	l := utf8.RuneCountInString(s)
	if l >= size {
		return s
	}
}

// in Go we have two visibility, public and package level
// public items can be accessed by other packages, or packages imported
// we capitalize the first character for making it public
// to make it private, we start with lower letter
