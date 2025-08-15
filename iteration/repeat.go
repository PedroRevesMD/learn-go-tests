package iteration

import "strings"

const COUNTER = 5

func Repeat(character string) string {
	var finalStirng strings.Builder
	for range COUNTER {
		finalStirng.WriteString(character)
	}

	return finalStirng.String()
}
