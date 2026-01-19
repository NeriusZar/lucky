package cmd

import "strings"

func cleanInput(input string) []string {
	input = strings.ToLower(input)
	return strings.Fields(input)
}
