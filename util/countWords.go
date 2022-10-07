package util

import (
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
)

func CountWords(input string) int {
	stripped := strip.StripTags(input)
	totalWords := len(strings.Fields(stripped))
	return totalWords
}
