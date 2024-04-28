package logs

import (
	"regexp"
)

func FilterLogs(logContents []string, pattern string) []string {
	var filteredLogs []string
	regex := regexp.MustCompile(pattern)

	for _, log := range logContents {
		if regex.MatchString(log) {
			filteredLogs = append(filteredLogs, log)
		}
	}

	return filteredLogs
}
