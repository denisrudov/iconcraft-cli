package main

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func camelCaseFromDash(input string) string {
	words := strings.Split(input, "-")
	c := cases.Title(language.Und)
	for i := 0; i < len(words); i++ {
		words[i] = c.String(words[i])
	}
	return strings.Join(words, "")
}
