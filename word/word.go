package word

import "strings"

func Count(s string) int {
	words := strings.Fields(s)
	return len(words)
}
