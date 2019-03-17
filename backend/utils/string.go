package utils

import "strings"

func ComposeStrings(params ...string) string {
	var pathBuilder strings.Builder
	for _, str := range params {
		pathBuilder.WriteString(str)
	}
	return pathBuilder.String()
}