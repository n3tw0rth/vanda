package vanda

import (
	"strings"
)

func Contains(s []string, target string) (bool, string) {
	for _, pattern := range s {
		subCommand := strings.Split(pattern, " ")[0]
		subCommand = strings.Trim(subCommand, "<>")

		if subCommand == target {
			return true, pattern
		}
	}
	return false, ""
}
