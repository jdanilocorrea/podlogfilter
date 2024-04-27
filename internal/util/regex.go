package util

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

func ReadInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func BuildRegexPattern(patterns []string) string {
	var regexParts []string
	for _, pattern := range patterns {
		regexParts = append(regexParts, regexp.QuoteMeta(pattern))
	}
	return strings.Join(regexParts, "|")
}
