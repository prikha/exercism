package acronym

import (
	"regexp"
	"strings"
)

const testVersion = 2


// Abbreviate does an interesting job in abbreviating.
func Abbreviate(s string) (out string) {
	var re = regexp.MustCompile(`[a-zA-Z]([a-z]+|[A-Z]+)`)

	for _, word := range re.FindAllString(s, -1) {
		out += string(word[0])
	}
	return strings.ToUpper(out)
}