package wordwrap

import (
	"strings"
)

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
