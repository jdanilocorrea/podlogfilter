package util

import (
	"bufio"
	"fmt"
	"strings"
)

func ReadInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
