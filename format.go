package main

import (
	"regexp"
	"strings"
)

func Format(text string) string {
	// Removes all special characters from string
	re, _ := regexp.Compile("/[^\\w\\s]/gi")

	text = re.FindString(text)

	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\r", " ")
	text = strings.ReplaceAll(text, "#", "")
	text = strings.ReplaceAll(text, "`", "")

	// Still more replacements left

	return text
}
