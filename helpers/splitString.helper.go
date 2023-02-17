package helpers

import "strings"

func SplitString(str string, sep string) []string {
	var parts []string
	for {
		index := strings.Index(str, sep)
		if index == -1 {
			break
		}
		parts = append(parts, str[:index])
		str = str[index+len(sep):]
	}
	parts = append(parts, str)
	return parts
}
