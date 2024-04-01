package stringutils

import "strings"

func Upper(word string) string {
	return strings.ToUpper(word)
}

func Lower(word string) string {
	return strings.ToLower(word)
}

func Title(word string) string {
	return strings.ToTitle(word)
}
