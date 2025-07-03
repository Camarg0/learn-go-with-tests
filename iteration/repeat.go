package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		// strings are immutable in Go, overwriting the string with the strings lib to gain performance
		repeated.WriteString(character)
	}
	return repeated.String()
}
