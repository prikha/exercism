package pangram

import "strings"

const testVersion = 1

func IsPangram(s string) bool {
	alphabet := make(map[rune]bool)
  lower_in := []rune(strings.ToLower(s))
	for _, r := range lower_in {
		if r >= 'a' && r <= 'z' {
			alphabet[r] = true
		}
	}
	return len(alphabet) == (('z' - 'a') + 1)
}
