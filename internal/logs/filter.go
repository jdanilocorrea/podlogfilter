package logs

import (
	"regexp"
	"strings"

	"github.com/jdanilocorrea/podlogfilter/internal/util"
)

func FilterLogs(logContent, patterns string) string {
	regex := regexp.MustCompile(util.BuildRegexPattern(strings.Split(patterns, ",")))
	var filteredLogs strings.Builder
	lines := strings.Split(logContent, "\n")
	for _, line := range lines {
		if regex.MatchString(line) {
			filteredLogs.WriteString(line + "\n")
		}
	}
	return filteredLogs.String()
}
