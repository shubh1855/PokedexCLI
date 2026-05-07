package main

import "strings"

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	trimmed := strings.TrimSpace(lowered)

	if trimmed == "" {
		return []string{}
	}

	return strings.Fields(trimmed)
}
