package iteration

import "strings"

// Repeat takes a string and number of times to repeat and returns a repeated string
func Repeat(character string, count int) string {
	return strings.Repeat(character, count)

}
