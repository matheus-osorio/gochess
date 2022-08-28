package fen

import (
	"regexp"
	"strings"
)

type FenStr string

//
func (f FenStr) Validate() bool {
	regex := `^[\/rnbkqpRNBKQ1-8P]+ (w|b) (-|[kqKQ]){1,4}$`
	matched, _ := regexp.Match(regex, []byte(f))
	return matched
}

func (f FenStr) Separate() (string, string, string) {
	parts := strings.Split(string(f), " ")

	return parts[0], parts[1], parts[2]
}
