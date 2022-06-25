package util

import "strings"

// SplitString splits a string by commas or empty spaces.
func SplitString(str string) []string {
	if len(str) == 0 {
		return []string{}
	}

	return strings.Fields(strings.ReplaceAll(str, ",", " "))
}
