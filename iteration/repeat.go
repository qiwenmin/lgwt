package iteration

import "strings"

const repeatCount = 5

// Repeat repeats the character five times.
func Repeat(character string) string {
	return strings.Repeat(character, 5)
}
