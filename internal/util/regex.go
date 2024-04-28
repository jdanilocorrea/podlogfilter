package util

import (
	"regexp"
	"strings"
)

func BuildRegexPattern(patterns []string) string {
	var regexParts []string
	for _, pattern := range patterns {
		regexParts = append(regexParts, regexp.QuoteMeta(pattern))
	}
	return strings.Join(regexParts, "|")
}
