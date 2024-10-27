package wordwrap

import (
	"strings"
)

// WordWrap takes a string and wraps it to fit within a certain width.
//
// All instances of " " are replaced with newlines (\n) to wrap the text.
// If newlines (\n) are in the text, it will split the string and continue.
//
// The algorithm works by looping through the string and finding the last
// index of a space that is not beyond the width. If it finds a newline
// character before the last space, it will use the newline character
// instead. If it doesn't find a space or a newline, it will split the
// string at the width.
//
// The function will return a single string with all the sub-strings
// joined together with newlines (\n) in between.
func WordWrap(s string, w int) string {
	var out []string

	for len(s) > w {
		lastIndex := strings.LastIndex(s[:w], " ")
		newlineIndex := strings.LastIndex(s[:w], "\n")
		if newlineIndex < lastIndex && newlineIndex > 0 {
			out = append(out, s[:newlineIndex])
			s = s[newlineIndex+1:]
			continue
		}
		if lastIndex == -1 {
			lastIndex = w
		}

		out = append(out, s[:lastIndex])
		if lastIndex <= w {
			s = s[lastIndex+1:]
			continue
		}
		s = s[lastIndex:]
	}
	out = append(out, s)

	return strings.Join(out, "\n")
}
