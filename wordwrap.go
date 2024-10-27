package wordwrap

import (
	"strings"
)

func WordWrap(s string, w int) string {
	var out []string

	for len(s) > w {
		lastIndex := strings.LastIndex(s[:w], " ")
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
